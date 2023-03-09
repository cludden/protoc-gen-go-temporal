package plugin

import (
	"fmt"

	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	g "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

// imported packages
const (
	activityPkg = "go.temporal.io/sdk/activity"
	clientPkg   = "go.temporal.io/sdk/client"
	enumsPkg    = "go.temporal.io/api/enums/v1"
	temporalPkg = "go.temporal.io/sdk/temporal"
	uuidPkg     = "github.com/google/uuid"
	workflowPkg = "go.temporal.io/sdk/workflow"
	workerPkg   = "go.temporal.io/sdk/worker"
)

// Service describes a temporal protobuf service definition
type Service struct {
	*protogen.Plugin
	*protogen.Service
	opts       *temporalv1.ServiceOptions
	activities map[string]*temporalv1.ActivityOptions
	methods    map[string]*protogen.Method
	queries    map[string]*temporalv1.QueryOptions
	signals    map[string]*temporalv1.SignalOptions
	workflows  map[string]*temporalv1.WorkflowOptions
}

// parseService extracts a Service from a protogen.Service value
func parseService(p *protogen.Plugin, service *protogen.Service) *Service {
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
		}

		if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Query).(*temporalv1.QueryOptions); ok && opts != nil {
			svc.queries[name] = opts
		}

		if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Signal).(*temporalv1.SignalOptions); ok && opts != nil {
			svc.signals[name] = opts
		}

		if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Workflow).(*temporalv1.WorkflowOptions); ok && opts != nil {
			svc.workflows[name] = opts
		}
	}
	return &svc
}

// render writes the temporal service to the given File
func (svc *Service) render(f *g.File) {
	svc.genConstants(f)

	// generate client interface and implementation
	svc.genClientInterface(f)
	svc.genClient(f)
	svc.genClientConstructor(f)

	// generate client workflow methods
	for workflow, opts := range svc.workflows {
		svc.genClientWorkflowExecute(f, workflow)
		svc.genClientWorkflowGet(f, workflow)
		for _, signal := range opts.GetSignal() {
			if signal.GetStart() {
				svc.genClientSignalWithStart(f, workflow, signal.GetRef())
			}
		}
	}

	// generate client query methods
	for query := range svc.queries {
		svc.genClientQueryMethod(f, query)
	}

	// generate client signal methods
	for signal := range svc.signals {
		svc.genClientSignalMethod(f, signal)
	}

	// generate <Workflow>Run interfaces and implementations used by client
	for workflow, opts := range svc.workflows {
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
	for workflow := range svc.workflows {
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
	for signal := range svc.signals {
		svc.genWorkerSignal(f, signal)
		svc.genWorkerSignalReceive(f, signal)
		svc.genWorkerSignalReceiveAsync(f, signal)
		svc.genWorkerSignalSelect(f, signal)
		svc.genWorkerSignalExternal(f, signal)
	}

	// generate activities
	svc.genActivitiesInterface(f)
	svc.genRegisterActivities(f)
	for activity := range svc.activities {
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
	// add workflow names
	if len(svc.workflows) > 0 {
		f.Commentf("%s workflow names", svc.GoName)
		f.Const().DefsFunc(func(defs *g.Group) {
			for workflow := range svc.workflows {
				method := svc.methods[workflow]
				defs.Id(fmt.Sprintf("%sName", workflow)).Op("=").Lit(string(method.Desc.FullName()))
			}
		})
	}

	// add id prefixes
	workflowsIdPrefixes := map[string]string{}
	for workflow, opts := range svc.workflows {
		if prefix := opts.GetDefaultOptions().GetIdPrefix(); prefix != "" {
			workflowsIdPrefixes[workflow] = prefix
		}
	}
	if len(workflowsIdPrefixes) > 0 {
		f.Commentf("%s id prefixes", svc.GoName)
		f.Const().DefsFunc(func(defs *g.Group) {
			for workflow, prefix := range workflowsIdPrefixes {
				defs.Id(fmt.Sprintf("%sIDPrefix", workflow)).Op("=").Lit(prefix)
			}
		})
	}

	// add query names
	if len(svc.queries) > 0 {
		f.Commentf("%s query names", svc.GoName)
		f.Const().DefsFunc(func(defs *g.Group) {
			for query := range svc.queries {
				method := svc.methods[query]
				defs.Id(fmt.Sprintf("%sName", query)).Op("=").Lit(string(method.Desc.FullName()))
			}
		})
	}

	// add signal names
	if len(svc.signals) > 0 {
		f.Commentf("%s signal names", svc.GoName)
		f.Const().DefsFunc(func(defs *g.Group) {
			for signal := range svc.signals {
				method := svc.methods[signal]
				defs.Id(fmt.Sprintf("%sName", signal)).Op("=").Lit(string(method.Desc.FullName()))
			}
		})
	}

	// add activity names
	if len(svc.activities) > 0 {
		f.Commentf("%s activity names", svc.GoName)
		f.Const().DefsFunc(func(defs *g.Group) {
			for activity := range svc.activities {
				method := svc.methods[activity]
				defs.Id(fmt.Sprintf("%sName", activity)).Op("=").Lit(string(method.Desc.FullName()))
			}
		})
	}
}
