package plugin

import (
	"errors"
	"fmt"
	"sort"

	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	g "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

// imported packages
const (
	activityPkg   = "go.temporal.io/sdk/activity"
	clientPkg     = "go.temporal.io/sdk/client"
	enumsPkg      = "go.temporal.io/api/enums/v1"
	expressionPkg = "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	helpersPkg    = "github.com/cludden/protoc-gen-go-temporal/pkg/helpers"
	temporalPkg   = "go.temporal.io/sdk/temporal"
	updatePkg     = "go.temporal.io/api/update/v1"
	uuidPkg       = "github.com/google/uuid"
	workflowPkg   = "go.temporal.io/sdk/workflow"
	workerPkg     = "go.temporal.io/sdk/worker"
)

const (
	modeWorkflow int = 1 << iota
	modeActivity
	modeQuery
	modeSignal
	modeUpdate
)

// Service describes a temporal protobuf service definition
type Service struct {
	*protogen.Plugin
	*protogen.Service
	*protogen.File
	cfg               *Config
	opts              *temporalv1.ServiceOptions
	activitiesOrdered []string
	activities        map[string]*temporalv1.ActivityOptions
	methods           map[string]*protogen.Method
	queriesOrdered    []string
	queries           map[string]*temporalv1.QueryOptions
	signalsOrdered    []string
	signals           map[string]*temporalv1.SignalOptions
	updatesOrdered    []string
	updates           map[string]*temporalv1.UpdateOptions
	workflowsOrdered  []string
	workflows         map[string]*temporalv1.WorkflowOptions
}

// parseService extracts a Service from a protogen.Service value
func parseService(p *protogen.Plugin, cfg *Config, file *protogen.File, service *protogen.Service) (*Service, error) {
	svc := Service{
		Plugin:     p,
		cfg:        cfg,
		Service:    service,
		File:       file,
		activities: make(map[string]*temporalv1.ActivityOptions),
		methods:    make(map[string]*protogen.Method),
		queries:    make(map[string]*temporalv1.QueryOptions),
		signals:    make(map[string]*temporalv1.SignalOptions),
		updates:    make(map[string]*temporalv1.UpdateOptions),
		workflows:  make(map[string]*temporalv1.WorkflowOptions),
	}

	if opts, ok := proto.GetExtension(service.Desc.Options(), temporalv1.E_Service).(*temporalv1.ServiceOptions); ok && opts != nil {
		svc.opts = opts
	}

	for _, method := range service.Methods {
		name := toCamel(method.GoName)
		svc.methods[name] = method

		var mode int
		if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Workflow).(*temporalv1.WorkflowOptions); ok && opts != nil {
			svc.workflows[name] = opts
			svc.workflowsOrdered = append(svc.workflowsOrdered, name)
			mode |= modeWorkflow
		}

		if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Activity).(*temporalv1.ActivityOptions); ok && opts != nil {
			svc.activities[name] = opts
			svc.activitiesOrdered = append(svc.activitiesOrdered, name)
			mode |= modeActivity
		}

		if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Query).(*temporalv1.QueryOptions); ok && opts != nil {
			svc.queries[name] = opts
			svc.queriesOrdered = append(svc.queriesOrdered, name)
			mode |= modeQuery
		}

		if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Signal).(*temporalv1.SignalOptions); ok && opts != nil {
			svc.signals[name] = opts
			svc.signalsOrdered = append(svc.signalsOrdered, name)
			mode |= modeSignal
		}

		if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Update).(*temporalv1.UpdateOptions); ok && opts != nil {
			if !svc.cfg.WorkflowUpdateEnabled {
				return nil, fmt.Errorf("method %q includes an update configuration, but workflow updates are not enabled: enable them with \"workflow-update-enabled=true\" plugin option", name)
			}
			svc.updates[name] = opts
			svc.updatesOrdered = append(svc.updatesOrdered, name)
			mode |= modeUpdate
		}

		switch mode {
		case 0:
		case modeActivity:
		case modeQuery:
		case modeSignal, modeSignal | modeActivity:
		case modeUpdate, modeUpdate | modeActivity:
		case modeWorkflow, modeWorkflow | modeActivity:
		default:
			p.Error(fmt.Errorf("invalid method options for method %q", method.Desc.FullName()))
		}
	}

	sort.Strings(svc.activitiesOrdered)
	sort.Strings(svc.queriesOrdered)
	sort.Strings(svc.signalsOrdered)
	sort.Strings(svc.updatesOrdered)
	sort.Strings(svc.workflowsOrdered)

	var errs error
	for _, workflow := range svc.workflowsOrdered {
		opts := svc.workflows[workflow]

		// ensure workflow queries are defined
		for _, queryOpts := range opts.GetQuery() {
			query := queryOpts.GetRef()
			if _, ok := svc.queries[query]; !ok {
				errs = errors.Join(errs, fmt.Errorf("workflow  %q references undefined query: %q", workflow, query))
			}
		}

		// ensure workflow signals are defined
		for _, signalOpts := range opts.GetSignal() {
			signal := signalOpts.GetRef()
			if _, ok := svc.signals[signal]; !ok {
				errs = errors.Join(errs, fmt.Errorf("workflow  %q references undefined signal: %q", workflow, signal))
			}
		}

		// ensure workflow updates are defined
		for _, updateOpts := range opts.GetUpdate() {
			update := updateOpts.GetRef()
			if _, ok := svc.updates[update]; !ok {
				errs = errors.Join(errs, fmt.Errorf("workflow  %q references undefined update: %q", workflow, update))
			}
		}
	}

	// ensure that signals return no value, unless signal method is also an activity, query, and/or workflow
	for _, signal := range svc.signalsOrdered {
		handler := svc.methods[signal]
		_, isActivity := svc.activities[signal]
		_, isQuery := svc.queries[signal]
		_, isUpdate := svc.updates[signal]
		_, isWorkflow := svc.workflows[signal]
		if !isActivity && !isQuery && !isUpdate && !isWorkflow && !isEmpty(handler.Output) {
			errs = errors.Join(errs, fmt.Errorf("expected signal %q output to be google.protobuf.Empty, got: %s", signal, handler.Output.GoIdent.GoName))
		}
	}
	return &svc, errs
}

func (svc *Service) fqnForActivity(activity string) string {
	if fqn := svc.activities[activity].GetName(); fqn != "" {
		return fqn
	}
	return string(svc.methods[activity].Desc.FullName())
}

func (svc *Service) fqnForQuery(query string) string {
	if fqn := svc.activities[query].GetName(); fqn != "" {
		return fqn
	}
	return string(svc.methods[query].Desc.FullName())
}

func (svc *Service) fqnForSignal(signal string) string {
	if fqn := svc.activities[signal].GetName(); fqn != "" {
		return fqn
	}
	return string(svc.methods[signal].Desc.FullName())
}

func (svc *Service) fqnForUpdate(update string) string {
	if fqn := svc.activities[update].GetName(); fqn != "" {
		return fqn
	}
	return string(svc.methods[update].Desc.FullName())
}

func (svc *Service) fqnForWorkflow(workflow string) string {
	if fqn := svc.activities[workflow].GetName(); fqn != "" {
		return fqn
	}
	return string(svc.methods[workflow].Desc.FullName())
}

// genConstants generates constants
func (svc *Service) genConstants(f *g.File) {
	// add task queue
	if taskQueue := svc.opts.GetTaskQueue(); taskQueue != "" {
		name := toCamel("%sTaskQueue", svc.Service.GoName)
		f.Commentf("%s= is the default task-queue for a %s worker", name, svc.Service.Desc.FullName())
		f.Const().Id(name).Op("=").Lit(taskQueue)
	}

	// add workflow names
	if len(svc.workflows) > 0 {
		f.Commentf("%s workflow names", svc.Service.Desc.FullName())
		f.Const().DefsFunc(func(defs *g.Group) {
			for _, workflow := range svc.workflowsOrdered {
				method := svc.methods[workflow]
				opts := svc.workflows[workflow]
				name := opts.GetName()
				if name == "" {
					name = string(method.Desc.FullName())
				}
				defs.Id(toCamel("%sWorkflowName", workflow)).Op("=").Lit(name)
			}
		})
	}

	// add workflow id expressions
	workflowIdExpressions := [][]string{}
	for _, workflow := range svc.workflowsOrdered {
		opts := svc.workflows[workflow]
		if expr := opts.GetId(); expr != "" {
			workflowIdExpressions = append(workflowIdExpressions, []string{workflow, expr})
		}
	}
	if len(workflowIdExpressions) > 0 {
		f.Commentf("%s workflow id expressions", svc.Service.Desc.FullName())
		f.Var().DefsFunc(func(defs *g.Group) {
			for _, pair := range workflowIdExpressions {
				defs.Id(toCamel("%sIDExpression", pair[0])).Op("=").Qual(expressionPkg, "MustParseExpression").Call(g.Lit(pair[1]))
			}
		})
	}

	// add workflow search attribute mappings
	workflowSearchAttributes := [][]string{}
	for _, workflow := range svc.workflowsOrdered {
		opts := svc.workflows[workflow]
		if mapping := opts.GetSearchAttributes(); mapping != "" {
			workflowSearchAttributes = append(workflowSearchAttributes, []string{workflow, mapping})
		}
	}
	if len(workflowSearchAttributes) > 0 {
		f.Commentf("%s workflow search attribute mappings", svc.Service.Desc.FullName())
		f.Var().DefsFunc(func(defs *g.Group) {
			for _, pair := range workflowSearchAttributes {
				defs.Id(toCamel("%sSearchAttributesMapping", pair[0])).Op("=").Qual(expressionPkg, "MustParseMapping").Call(g.Lit(pair[1]))
			}
		})
	}

	// add activity names
	if len(svc.activities) > 0 {
		f.Commentf("%s activity names", svc.Service.Desc.FullName())
		f.Const().DefsFunc(func(defs *g.Group) {
			for _, activity := range svc.activitiesOrdered {
				method := svc.methods[activity]
				opts := svc.activities[activity]
				name := opts.GetName()
				if name == "" {
					name = string(method.Desc.FullName())
				}
				defs.Id(toCamel("%sActivityName", activity)).Op("=").Lit(name)
			}
		})
	}

	// add query names
	if len(svc.queries) > 0 {
		f.Commentf("%s query names", svc.Service.Desc.FullName())
		f.Const().DefsFunc(func(defs *g.Group) {
			for _, query := range svc.queriesOrdered {
				method := svc.methods[query]
				opts := svc.queries[query]
				name := opts.GetName()
				if name == "" {
					name = string(method.Desc.FullName())
				}
				defs.Id(toCamel("%sQueryName", query)).Op("=").Lit(name)
			}
		})
	}

	// add signal names
	if len(svc.signals) > 0 {
		f.Commentf("%s signal names", svc.Service.Desc.FullName())
		f.Const().DefsFunc(func(defs *g.Group) {
			for _, signal := range svc.signalsOrdered {
				method := svc.methods[signal]
				opts := svc.signals[signal]
				name := opts.GetName()
				if name == "" {
					name = string(method.Desc.FullName())
				}
				defs.Id(toCamel("%sSignalName", signal)).Op("=").Lit(name)
			}
		})
	}

	// add update names
	if len(svc.updates) > 0 {
		f.Commentf("%s update names", svc.Service.Desc.FullName())
		f.Const().DefsFunc(func(defs *g.Group) {
			for _, update := range svc.updatesOrdered {
				method := svc.methods[update]
				opts := svc.updates[update]
				name := opts.GetName()
				if name == "" {
					name = string(method.Desc.FullName())
				}
				defs.Id(toCamel("%sUpdateName", update)).Op("=").Lit(name)
			}
		})
	}

	// add update id expressions
	updateIdExpressions := [][]string{}
	for _, update := range svc.updatesOrdered {
		opts := svc.updates[update]
		if expr := opts.GetId(); expr != "" {
			updateIdExpressions = append(updateIdExpressions, []string{update, expr})
		}
	}
	if len(updateIdExpressions) > 0 {
		f.Commentf("%s update id expressions", svc.Service.Desc.FullName())
		f.Var().DefsFunc(func(defs *g.Group) {
			for _, pair := range updateIdExpressions {
				defs.Id(toCamel("%sIDExpression", pair[0])).Op("=").Qual(expressionPkg, "MustParseExpression").Call(g.Lit(pair[1]))
			}
		})
	}
}

// render writes the temporal service to the given File
func (svc *Service) render(f *g.File) {
	svc.genConstants(f)

	// generate client interface and implementation
	svc.genClientInterface(f)
	svc.genClientImpl(f)
	svc.genClientImplConstructor(f)

	// generate client workflow methods
	for _, workflow := range svc.workflowsOrdered {
		opts := svc.workflows[workflow]
		svc.genClientImplWorkflowMethod(f, workflow)
		svc.genClientImplWorkflowAsyncMethod(f, workflow)
		svc.genClientImplWorkflowGetMethod(f, workflow)
		for _, signal := range opts.GetSignal() {
			if signal.GetStart() {
				svc.genClientImplSignalWithStartMethod(f, workflow, signal.GetRef())
				svc.genClientImplSignalWithStartAsyncMethod(f, workflow, signal.GetRef())
			}
		}
	}
	svc.genClientImplWorkflowCancelMethod(f)
	svc.genClientImplWorkflowTerminateMethod(f)

	// generate client query methods
	for _, query := range svc.queriesOrdered {
		svc.genClientImplQueryMethod(f, query)
	}

	// generate client signal methods
	for _, signal := range svc.signalsOrdered {
		svc.genClientImplSignalMethod(f, signal)
	}

	// generate client update methods
	for _, update := range svc.updatesOrdered {
		svc.genClientImplUpdateMethod(f, update)
		svc.genClientImplUpdateMethodAsync(f, update)
	}

	// generate <Workflow>Options, <Workflow>Run interfaces and implementations used by client
	for _, workflow := range svc.workflowsOrdered {
		opts := svc.workflows[workflow]
		svc.genClientWorkflowOptions(f, workflow)
		svc.genClientWorkflowRunInterface(f, workflow)
		svc.genClientWorkflowRunImpl(f, workflow)
		svc.genClientWorkflowRunImplIDMethod(f, workflow)
		svc.genClientWorkflowRunImplRunIDMethod(f, workflow)
		svc.genClientWorkflowRunImplCancelMethod(f, workflow)
		svc.genClientWorkflowRunImplGetMethod(f, workflow)
		svc.genClientWorkflowRunImplTerminateMethod(f, workflow)

		// generate query methods
		for _, queryOpts := range opts.GetQuery() {
			svc.genClientWorkflowRunImplQueryMethod(f, workflow, queryOpts.GetRef())
		}

		// generate signal methods
		for _, signalOpts := range opts.GetSignal() {
			svc.genClientWorkflowRunImplSignalMethod(f, workflow, signalOpts.GetRef())
		}

		// generate update methods
		for _, updateOpts := range opts.GetUpdate() {
			svc.genClientWorkflowRunImplUpdateMethod(f, workflow, updateOpts.GetRef())
			svc.genClientWorkflowRunImplUpdateAsyncMethod(f, workflow, updateOpts.GetRef())
		}
	}

	// generate <Update>Handle interfaces and implementations used by client
	for _, update := range svc.updatesOrdered {
		svc.genClientUpdateHandleInterface(f, update)
		svc.genClientUpdateHandleImpl(f, update)
		svc.genClientUpdateHandleImplWorkflowIDMethod(f, update)
		svc.genClientUpdateHandleImplRunIDMethod(f, update)
		svc.genClientUpdateHandleImplUpdateIDMethod(f, update)
		svc.genClientUpdateHandleImplGetMethod(f, update)
		svc.genClientUpdateOptions(f, update)
	}

	// generate workflows interface and registration helper
	svc.genWorkerWorkflowFunctionVars(f)
	svc.genWorkerWorkflowsInterface(f)
	svc.genWorkerRegisterWorkflows(f)

	// generate workflow types, methods, functions
	for _, workflow := range svc.workflowsOrdered {
		svc.genWorkerRegisterWorkflow(f, workflow)
		svc.genWorkerBuilderFunction(f, workflow)
		svc.genWorkerWorkflowInput(f, workflow)
		svc.genWorkerWorkflowInterface(f, workflow)
		svc.genWorkerWorkflowChild(f, workflow)
		svc.genWorkerWorkflowChildAsync(f, workflow)
		svc.genWorkerWorkflowChildOptions(f, workflow)
		svc.genWorkerWorkflowChildRun(f, workflow)
		svc.genWorkerWorkflowChildRunGet(f, workflow)
		svc.genWorkerWorkflowChildRunSelect(f, workflow)
		svc.genWorkerWorkflowChildRunSelectStart(f, workflow)
		svc.genWorkerWorkflowChildRunWaitStart(f, workflow)
		svc.genWorkerWorkflowChildRunSignals(f, workflow)
	}

	// generate signal types, methods, functions
	for _, signal := range svc.signalsOrdered {
		svc.genWorkerSignal(f, signal)
		svc.genWorkerSignalReceive(f, signal)
		svc.genWorkerSignalReceiveAsync(f, signal)
		svc.genWorkerSignalSelect(f, signal)
		svc.genWorkerSignalExternal(f, signal)
		svc.genWorkerSignalExternalAsync(f, signal)
	}

	// generate activities
	svc.genActivitiesInterface(f)
	svc.genActivityRegisterAllFunction(f)
	for _, activity := range svc.activitiesOrdered {
		svc.genActivityRegisterOneFunction(f, activity)
		svc.genActivityFuture(f, activity)
		svc.genActivityFutureGetMethod(f, activity)
		svc.genActivityFutureSelectMethod(f, activity)
		svc.genActivityFunction(f, activity, false, false)
		svc.genActivityFunction(f, activity, false, true)
		svc.genActivityFunction(f, activity, true, false)
		svc.genActivityFunction(f, activity, true, true)
		svc.genActivityLocalOptions(f, activity)
		svc.genActivityOptions(f, activity)
	}
}
