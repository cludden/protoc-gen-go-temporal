package plugin

import (
	"fmt"
	"sort"

	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	g "github.com/dave/jennifer/jen"
	"github.com/hashicorp/go-multierror"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

// imported packages
const (
	activityPkg   = "go.temporal.io/sdk/activity"
	clientPkg     = "go.temporal.io/sdk/client"
	enumsPkg      = "go.temporal.io/api/enums/v1"
	expressionPkg = "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	temporalPkg   = "go.temporal.io/sdk/temporal"
	uuidPkg       = "github.com/google/uuid"
	workflowPkg   = "go.temporal.io/sdk/workflow"
	workerPkg     = "go.temporal.io/sdk/worker"
)

// Service describes a temporal protobuf service definition
type Service struct {
	*protogen.Plugin
	*protogen.Service
	opts              *temporalv1.ServiceOptions
	activitiesOrdered []string
	activities        map[string]*temporalv1.ActivityOptions
	methods           map[string]*protogen.Method
	queriesOrdered    []string
	queries           map[string]*temporalv1.QueryOptions
	signalsOrdered    []string
	signals           map[string]*temporalv1.SignalOptions
	workflowsOrdered  []string
	workflows         map[string]*temporalv1.WorkflowOptions
}

// parseService extracts a Service from a protogen.Service value
func parseService(p *protogen.Plugin, service *protogen.Service) (*Service, error) {
	svc := Service{
		Plugin:     p,
		Service:    service,
		activities: make(map[string]*temporalv1.ActivityOptions),
		methods:    make(map[string]*protogen.Method),
		queries:    make(map[string]*temporalv1.QueryOptions),
		signals:    make(map[string]*temporalv1.SignalOptions),
		workflows:  make(map[string]*temporalv1.WorkflowOptions),
	}

	if opts, ok := proto.GetExtension(service.Desc.Options(), temporalv1.E_Service).(*temporalv1.ServiceOptions); ok && opts != nil {
		svc.opts = opts
	}

	for _, method := range service.Methods {
		name := method.GoName
		svc.methods[name] = method

		if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Activity).(*temporalv1.ActivityOptions); ok && opts != nil {
			svc.activities[name] = opts
			svc.activitiesOrdered = append(svc.activitiesOrdered, name)
		}

		if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Query).(*temporalv1.QueryOptions); ok && opts != nil {
			svc.queries[name] = opts
			svc.queriesOrdered = append(svc.queriesOrdered, name)
		}

		if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Signal).(*temporalv1.SignalOptions); ok && opts != nil {
			svc.signals[name] = opts
			svc.signalsOrdered = append(svc.signalsOrdered, name)
		}

		if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Workflow).(*temporalv1.WorkflowOptions); ok && opts != nil {
			svc.workflows[name] = opts
			svc.workflowsOrdered = append(svc.workflowsOrdered, name)
		}
	}
	sort.Strings(svc.activitiesOrdered)
	sort.Strings(svc.queriesOrdered)
	sort.Strings(svc.signalsOrdered)
	sort.Strings(svc.workflowsOrdered)

	errs := multierror.Append(nil)
	for _, workflow := range svc.workflowsOrdered {
		opts := svc.workflows[workflow]

		// ensure workflow signals are defined
		for _, signalOpts := range opts.GetSignal() {
			signal := signalOpts.GetRef()
			if _, ok := svc.signals[signal]; !ok {
				errs = multierror.Append(errs, fmt.Errorf("workflow  %q references undefined signal: %q", workflow, signal))
			}
		}

		// ensure workflow queries are defined
		for _, queryOpts := range opts.GetQuery() {
			query := queryOpts.GetRef()
			if _, ok := svc.queries[query]; !ok {
				errs = multierror.Append(errs, fmt.Errorf("workflow  %q references undefined query: %q", workflow, query))
			}
		}
	}
	for _, signal := range svc.signalsOrdered {
		handler := svc.methods[signal]
		_, isActivity := svc.activities[signal]
		_, isWorkflow := svc.workflows[signal]
		_, isQuery := svc.queries[signal]
		if !isActivity && !isWorkflow && !isQuery && !isEmpty(handler.Output) {
			errs = multierror.Append(errs, fmt.Errorf("expected signal %q output to be google.protobuf.Empty, got: %s", signal, handler.Output.GoIdent.GoName))
		}
	}
	return &svc, errs.ErrorOrNil()
}

// render writes the temporal service to the given File
func (svc *Service) render(f *g.File) {
	svc.genConstants(f)

	// generate client interface and implementation
	svc.genClientInterface(f)
	svc.genClient(f)
	svc.genClientConstructor(f)

	// generate client workflow methods
	for _, workflow := range svc.workflowsOrdered {
		opts := svc.workflows[workflow]
		svc.genClientWorkflow(f, workflow)
		svc.genClientWorkflowExecute(f, workflow)
		svc.genClientWorkflowGet(f, workflow)
		for _, signal := range opts.GetSignal() {
			if signal.GetStart() {
				svc.genClientSignalWithStart(f, workflow, signal.GetRef())
			}
		}
	}

	// generate client query methods
	for _, query := range svc.queriesOrdered {
		svc.genClientQueryMethod(f, query)
	}

	// generate client signal methods
	for _, signal := range svc.signalsOrdered {
		svc.genClientSignalMethod(f, signal)
	}

	// generate <Workflow>Run interfaces and implementations used by client
	for _, workflow := range svc.workflowsOrdered {
		opts := svc.workflows[workflow]
		svc.genClientWorkflowRunInterface(f, workflow)
		svc.genClientWorkflowRun(f, workflow)
		svc.genClientWorkflowRunIDMethod(f, workflow)
		svc.genClientWorkflowRunRunIDMethod(f, workflow)
		svc.genClientWorkflowRunGetMethod(f, workflow)

		// generate query methods
		for _, queryOpts := range opts.GetQuery() {
			svc.genClientWorkflowRunQueryMethod(f, workflow, queryOpts.GetRef())
		}

		// generate signal methods
		for _, signalOpts := range opts.GetSignal() {
			svc.genClientWorkflowRunSignalMethod(f, workflow, signalOpts.GetRef())
		}
	}

	// generate workflows interface and registration helper
	svc.genWorkflowsInterface(f)
	svc.genRegisterWorkflows(f)

	// generate workflow types, methods, functions
	for _, workflow := range svc.workflowsOrdered {
		svc.genRegisterWorkflow(f, workflow)
		svc.genWorkflowWorkerBuilderFunction(f, workflow)
		svc.genWorkflowWorker(f, workflow)
		svc.genWorkflowWorkerExecuteMethod(f, workflow)
		svc.genWorkflowInput(f, workflow)
		svc.genWorkflowInterface(f, workflow)
		svc.genExecuteChildWorkflow(f, workflow)
		svc.genWorkflowChildRun(f, workflow)
		svc.genWorkflowChildRunGet(f, workflow)
		svc.genWorkflowChildRunSelect(f, workflow)
		svc.genWorkflowChildRunSelectStart(f, workflow)
		svc.genWorkflowChildRunWaitStart(f, workflow)
		svc.genWorkflowChildRunSignals(f, workflow)
	}

	// generate signal types, methods, functions
	for _, signal := range svc.signalsOrdered {
		svc.genWorkerSignal(f, signal)
		svc.genWorkerSignalReceive(f, signal)
		svc.genWorkerSignalReceiveAsync(f, signal)
		svc.genWorkerSignalSelect(f, signal)
		svc.genWorkerSignalExternal(f, signal)
	}

	// generate activities
	svc.genActivitiesInterface(f)
	svc.genRegisterActivities(f)
	for _, activity := range svc.activitiesOrdered {
		svc.genRegisterActivity(f, activity)
		svc.genActivityFuture(f, activity)
		svc.genActivityFutureGetMethod(f, activity)
		svc.genActivityFutureSelectMethod(f, activity)
		svc.genActivityFunction(f, activity, false)
		svc.genActivityFunction(f, activity, true)
	}
}

// genConstants generates constants
func (svc *Service) genConstants(f *g.File) {
	// add task queue
	if taskQueue := svc.opts.GetTaskQueue(); taskQueue != "" {
		f.Commentf("%sTaskQueue is the default task-queue for a %s worker", svc.GoName, svc.GoName)
		f.Const().Id(fmt.Sprintf("%sTaskQueue", svc.GoName)).Op("=").Lit(taskQueue)
	}

	// add workflow names
	if len(svc.workflows) > 0 {
		f.Commentf("%s workflow names", svc.GoName)
		f.Const().DefsFunc(func(defs *g.Group) {
			for _, workflow := range svc.workflowsOrdered {
				method := svc.methods[workflow]
				defs.Id(fmt.Sprintf("%sWorkflowName", workflow)).Op("=").Lit(fmt.Sprintf("%sWorkflow", string(method.Desc.FullName())))
			}
		})
	}

	// add id expressions
	workflowIdExpressions := [][]string{}
	for _, workflow := range svc.workflowsOrdered {
		opts := svc.workflows[workflow]
		if expr := opts.GetDefaultOptions().GetId(); expr != "" {
			workflowIdExpressions = append(workflowIdExpressions, []string{workflow, expr})
		}
	}
	if len(workflowIdExpressions) > 0 {
		f.Commentf("%s id expressions", svc.GoName)
		f.Var().DefsFunc(func(defs *g.Group) {
			for _, pair := range workflowIdExpressions {
				defs.Id(fmt.Sprintf("%sIDExpression", pair[0])).Op("=").Qual(expressionPkg, "MustParseExpression").Call(g.Lit(pair[1]))
			}
		})
	}

	// add query names
	if len(svc.queries) > 0 {
		f.Commentf("%s query names", svc.GoName)
		f.Const().DefsFunc(func(defs *g.Group) {
			for _, query := range svc.queriesOrdered {
				method := svc.methods[query]
				defs.Id(fmt.Sprintf("%sQueryName", query)).Op("=").Lit(fmt.Sprintf("%sQuery", string(method.Desc.FullName())))
			}
		})
	}

	// add signal names
	if len(svc.signals) > 0 {
		f.Commentf("%s signal names", svc.GoName)
		f.Const().DefsFunc(func(defs *g.Group) {
			for _, signal := range svc.signalsOrdered {
				method := svc.methods[signal]
				defs.Id(fmt.Sprintf("%sSignalName", signal)).Op("=").Lit(fmt.Sprintf("%sSignal", string(method.Desc.FullName())))
			}
		})
	}

	// add activity names
	if len(svc.activities) > 0 {
		f.Commentf("%s activity names", svc.GoName)
		f.Const().DefsFunc(func(defs *g.Group) {
			for _, activity := range svc.activitiesOrdered {
				method := svc.methods[activity]
				defs.Id(fmt.Sprintf("%sActivityName", activity)).Op("=").Lit(fmt.Sprintf("%sActivity", string(method.Desc.FullName())))
			}
		})
	}
}
