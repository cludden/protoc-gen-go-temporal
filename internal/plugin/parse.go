package plugin

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/alta/protopatch/lint"
	"github.com/alta/protopatch/patch/gopb"
	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	g "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// imported packages
const (
	activityPkg     = "go.temporal.io/sdk/activity"
	clientPkg       = "go.temporal.io/sdk/client"
	durationpbPkg   = "google.golang.org/protobuf/types/known/durationpb"
	enumsPkg        = "go.temporal.io/api/enums/v1"
	expressionPkg   = "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	helpersPkg      = "github.com/cludden/protoc-gen-go-temporal/pkg/helpers"
	protoreflectPkg = "google.golang.org/protobuf/reflect/protoreflect"
	serviceerrorPkg = "go.temporal.io/api/serviceerror"
	temporalPkg     = "go.temporal.io/sdk/temporal"
	temporalv1Pkg   = "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	timestamppbPkg  = "google.golang.org/protobuf/types/known/timestamppb"
	updatePkg       = "go.temporal.io/api/update/v1"
	uuidPkg         = "github.com/google/uuid"
	workflowPkg     = "go.temporal.io/sdk/workflow"
	workerPkg       = "go.temporal.io/sdk/worker"
)

const (
	modeWorkflow int = 1 << iota
	modeActivity
	modeQuery
	modeSignal
	modeUpdate
)

var aliases = map[string]string{
	enumsPkg:      "enumsv1",
	temporalv1Pkg: "temporalv1",
	updatePkg:     "updatev1",
	xnsv1Pkg:      "xnsv1",
}

// Manifest describes temporal artifacts parsed from source protos
type Manifest struct {
	*Plugin
	*protogen.Service
	*protogen.File

	opts              *temporalv1.ServiceOptions
	activitiesOrdered []protoreflect.FullName
	activities        map[protoreflect.FullName]*temporalv1.ActivityOptions
	commands          map[protoreflect.FullName]*temporalv1.CommandOptions
	files             map[protoreflect.FullName]*protogen.File
	methods           map[protoreflect.FullName]*protogen.Method
	patches           map[temporalv1.Patch_Version]temporalv1.Patch_Mode
	patchesByRef      map[protoreflect.FullName]map[temporalv1.Patch_Version]temporalv1.Patch_Mode
	queriesOrdered    []protoreflect.FullName
	queries           map[protoreflect.FullName]*temporalv1.QueryOptions
	serviceFiles      map[protoreflect.FullName]*protogen.File
	serviceOptions    map[protoreflect.FullName]*temporalv1.ServiceOptions
	signalsOrdered    []protoreflect.FullName
	signals           map[protoreflect.FullName]*temporalv1.SignalOptions
	updatesOrdered    []protoreflect.FullName
	updates           map[protoreflect.FullName]*temporalv1.UpdateOptions
	workflowsOrdered  []protoreflect.FullName
	workflows         map[protoreflect.FullName]*temporalv1.WorkflowOptions
}

// parse extracts a Service from a protogen.Service value
func parse(p *Plugin) (*Manifest, error) {
	svc := Manifest{
		Plugin:         p,
		activities:     make(map[protoreflect.FullName]*temporalv1.ActivityOptions),
		commands:       make(map[protoreflect.FullName]*temporalv1.CommandOptions),
		files:          make(map[protoreflect.FullName]*protogen.File),
		methods:        make(map[protoreflect.FullName]*protogen.Method),
		patches:        make(map[temporalv1.Patch_Version]temporalv1.Patch_Mode),
		patchesByRef:   make(map[protoreflect.FullName]map[temporalv1.Patch_Version]temporalv1.Patch_Mode),
		queries:        make(map[protoreflect.FullName]*temporalv1.QueryOptions),
		serviceFiles:   make(map[protoreflect.FullName]*protogen.File),
		serviceOptions: make(map[protoreflect.FullName]*temporalv1.ServiceOptions),
		signals:        make(map[protoreflect.FullName]*temporalv1.SignalOptions),
		updates:        make(map[protoreflect.FullName]*temporalv1.UpdateOptions),
		workflows:      make(map[protoreflect.FullName]*temporalv1.WorkflowOptions),
	}

	// index global patch settings
	for _, p := range strings.Split(svc.cfg.Patches, ";") {
		if p == "" {
			continue
		}
		fields := strings.Split(p, "_")
		if len(fields) > 2 {
			return nil, fmt.Errorf("invalid patches option")
		}
		pv := temporalv1.Patch_Version(temporalv1.Patch_Version_value[fmt.Sprintf("PV_%s", fields[0])])
		pvm := temporalv1.Patch_PVM_ENABLED
		if len(fields) > 0 {
			pvm = temporalv1.Patch_Mode(temporalv1.Patch_Mode_value[fmt.Sprintf("PVM_%s", fields[1])])
		}
		svc.patches[pv] = pvm
	}

	for _, file := range p.Files {
		if !file.Generate {
			continue
		}

		for _, service := range file.Services {
			svc.serviceFiles[service.Desc.FullName()] = file
			if opts, ok := proto.GetExtension(service.Desc.Options(), temporalv1.E_Service).(*temporalv1.ServiceOptions); ok && opts != nil {
				svc.opts = opts
				svc.serviceOptions[service.Desc.FullName()] = opts

				// index service level patch settings
				if len(opts.GetPatches()) > 0 {
					patchIndex := make(map[temporalv1.Patch_Version]temporalv1.Patch_Mode)
					for _, p := range opts.GetPatches() {
						patchIndex[p.GetVersion()] = p.GetMode()
					}
					svc.patchesByRef[service.Desc.FullName()] = patchIndex
				}
			}

			for _, method := range service.Methods {
				name := method.Desc.FullName()
				svc.methods[name] = method
				svc.files[name] = file

				var mode int
				var patches []*temporalv1.Patch
				if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Workflow).(*temporalv1.WorkflowOptions); ok && opts != nil {
					svc.workflows[name] = opts
					svc.workflowsOrdered = append(svc.workflowsOrdered, name)
					mode |= modeWorkflow
					patches = opts.GetPatches()
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
					patches = opts.GetPatches()
				}

				if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Signal).(*temporalv1.SignalOptions); ok && opts != nil {
					svc.signals[name] = opts
					svc.signalsOrdered = append(svc.signalsOrdered, name)
					mode |= modeSignal
					patches = opts.GetPatches()
				}

				if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Update).(*temporalv1.UpdateOptions); ok && opts != nil {
					if !svc.cfg.WorkflowUpdateEnabled {
						return nil, fmt.Errorf("method %q includes an update configuration, but workflow updates are not enabled: enable them with \"workflow-update-enabled=true\" plugin option", name)
					}
					svc.updates[name] = opts
					svc.updatesOrdered = append(svc.updatesOrdered, name)
					mode |= modeUpdate
					patches = opts.GetPatches()
				}

				if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Command).(*temporalv1.CommandOptions); ok && opts != nil {
					svc.commands[name] = opts
				}

				// validate method option combinations
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

				// index method level patch settings
				if len(patches) > 0 {
					patchIndex := make(map[temporalv1.Patch_Version]temporalv1.Patch_Mode)
					for _, p := range patches {
						patchIndex[p.GetVersion()] = p.GetMode()
					}
					svc.patchesByRef[method.Desc.FullName()] = patchIndex
				}
			}
		}
	}

	sortFullNames(svc.activitiesOrdered)
	sortFullNames(svc.queriesOrdered)
	sortFullNames(svc.signalsOrdered)
	sortFullNames(svc.updatesOrdered)
	sortFullNames(svc.workflowsOrdered)

	var errs error
	for _, workflow := range svc.workflowsOrdered {
		opts := svc.workflows[workflow]

		// ensure workflow queries are defined
		for _, queryOpts := range opts.GetQuery() {
			query := getFullyQualifiedRef(workflow, queryOpts.GetRef())
			if _, ok := svc.queries[query]; !ok {
				errs = errors.Join(errs, fmt.Errorf("workflow  %q references undefined query: %q", workflow, query))
			}
		}

		// ensure workflow signals are defined
		for _, signalOpts := range opts.GetSignal() {
			signal := getFullyQualifiedRef(workflow, signalOpts.GetRef())
			if _, ok := svc.signals[signal]; !ok {
				errs = errors.Join(errs, fmt.Errorf("workflow  %q references undefined signal: %q", workflow, signal))
			}
		}

		// ensure workflow updates are defined
		for _, updateOpts := range opts.GetUpdate() {
			update := getFullyQualifiedRef(workflow, updateOpts.GetRef())
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

func (svc *Manifest) fqnForActivity(activity protoreflect.FullName) string {
	if fqn := svc.activities[activity].GetName(); fqn != "" {
		return fqn
	}
	return string(svc.methods[activity].Desc.FullName())
}

func (svc *Manifest) fqnForQuery(query protoreflect.FullName) string {
	if fqn := svc.queries[query].GetName(); fqn != "" {
		return fqn
	}
	return string(svc.methods[query].Desc.FullName())
}

func (svc *Manifest) fqnForSignal(signal protoreflect.FullName) string {
	if fqn := svc.signals[signal].GetName(); fqn != "" {
		return fqn
	}
	return string(svc.methods[signal].Desc.FullName())
}

func (svc *Manifest) fqnForUpdate(update protoreflect.FullName) string {
	if fqn := svc.updates[update].GetName(); fqn != "" {
		return fqn
	}
	return string(svc.methods[update].Desc.FullName())
}

func (svc *Manifest) fqnForWorkflow(workflow protoreflect.FullName) string {
	if fqn := svc.workflows[workflow].GetName(); fqn != "" {
		return fqn
	}
	return string(svc.methods[workflow].Desc.FullName())
}

// genConstants generates constants
func (svc *Manifest) genConstants(f *g.File) {
	// add task queue
	if taskQueue := svc.opts.GetTaskQueue(); taskQueue != "" {
		name := svc.toCamel("%sTaskQueue", svc.Service.GoName)
		f.Commentf("%s is the default task-queue for a %s worker", name, svc.Service.Desc.FullName())
		f.Const().Id(name).Op("=").Lit(taskQueue)
	}

	// add workflow names
	var workflows []protoreflect.FullName
	for _, workflow := range svc.workflowsOrdered {
		if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
			continue
		}
		workflows = append(workflows, workflow)
	}
	if len(workflows) > 0 {
		f.Commentf("%s workflow names", svc.Service.Desc.FullName())
		f.Const().DefsFunc(func(defs *g.Group) {
			for _, workflow := range workflows {
				method := svc.methods[workflow]
				opts := svc.workflows[workflow]
				name := opts.GetName()
				if name == "" {
					name = string(method.Desc.FullName())
				}
				defs.Id(svc.toCamel("%sWorkflowName", workflow)).Op("=").Lit(name)
			}
		})
	}

	// add workflow id expressions
	workflowIdExpressions := [][]string{}
	for _, workflow := range svc.workflowsOrdered {
		if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
			continue
		}
		opts := svc.workflows[workflow]
		if expr := opts.GetId(); expr != "" {
			workflowIdExpressions = append(workflowIdExpressions, []string{svc.methods[workflow].GoName, expr})
		}
	}
	if len(workflowIdExpressions) > 0 {
		f.Commentf("%s workflow id expressions", svc.Service.Desc.FullName())
		f.Var().DefsFunc(func(defs *g.Group) {
			for _, pair := range workflowIdExpressions {
				defs.Id(svc.toCamel("%sIDExpression", pair[0])).Op("=").Qual(expressionPkg, "MustParseExpression").Call(g.Lit(pair[1]))
			}
		})
	}

	// add workflow search attribute mappings
	workflowSearchAttributes := [][]string{}
	for _, workflow := range svc.workflowsOrdered {
		if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
			continue
		}
		opts := svc.workflows[workflow]
		if mapping := opts.GetSearchAttributes(); mapping != "" {
			workflowSearchAttributes = append(workflowSearchAttributes, []string{svc.methods[workflow].GoName, mapping})
		}
	}
	if len(workflowSearchAttributes) > 0 {
		f.Commentf("%s workflow search attribute mappings", svc.Service.Desc.FullName())
		f.Var().DefsFunc(func(defs *g.Group) {
			for _, pair := range workflowSearchAttributes {
				defs.Id(svc.toCamel("%sSearchAttributesMapping", pair[0])).Op("=").Qual(expressionPkg, "MustParseMapping").Call(g.Lit(pair[1]))
			}
		})
	}

	// add activity names
	var activities []protoreflect.FullName
	for _, activity := range svc.activitiesOrdered {
		if svc.methods[activity].Desc.Parent() != svc.Service.Desc {
			continue
		}
		activities = append(activities, activity)
	}
	if len(activities) > 0 {
		f.Commentf("%s activity names", svc.Service.Desc.FullName())
		f.Const().DefsFunc(func(defs *g.Group) {
			for _, activity := range activities {
				method := svc.methods[activity]
				opts := svc.activities[activity]
				name := opts.GetName()
				if name == "" {
					name = string(method.Desc.FullName())
				}
				defs.Id(svc.toCamel("%sActivityName", activity)).Op("=").Lit(name)
			}
		})
	}

	// add query names
	var queries []protoreflect.FullName
	for _, query := range svc.queriesOrdered {
		if svc.methods[query].Desc.Parent() != svc.Service.Desc {
			continue
		}
		queries = append(queries, query)
	}
	if len(queries) > 0 {
		f.Commentf("%s query names", svc.Service.Desc.FullName())
		f.Const().DefsFunc(func(defs *g.Group) {
			for _, query := range queries {
				method := svc.methods[query]
				opts := svc.queries[query]
				name := opts.GetName()
				if name == "" {
					name = string(method.Desc.FullName())
				}
				defs.Id(svc.toCamel("%sQueryName", query)).Op("=").Lit(name)
			}
		})
	}

	// add signal names
	var signals []protoreflect.FullName
	for _, signal := range svc.signalsOrdered {
		if svc.methods[signal].Desc.Parent() != svc.Service.Desc {
			continue
		}
		signals = append(signals, signal)
	}
	if len(signals) > 0 {
		f.Commentf("%s signal names", svc.Service.Desc.FullName())
		f.Const().DefsFunc(func(defs *g.Group) {
			for _, signal := range signals {
				method := svc.methods[signal]
				opts := svc.signals[signal]
				name := opts.GetName()
				if name == "" {
					name = string(method.Desc.FullName())
				}
				defs.Id(svc.toCamel("%sSignalName", signal)).Op("=").Lit(name)
			}
		})
	}

	// add update names
	var updates []protoreflect.FullName
	for _, update := range svc.updatesOrdered {
		if svc.methods[update].Desc.Parent() != svc.Service.Desc {
			continue
		}
		updates = append(updates, update)
	}
	if len(updates) > 0 {
		f.Commentf("%s update names", svc.Service.Desc.FullName())
		f.Const().DefsFunc(func(defs *g.Group) {
			for _, update := range updates {
				method := svc.methods[update]
				opts := svc.updates[update]
				name := opts.GetName()
				if name == "" {
					name = string(method.Desc.FullName())
				}
				defs.Id(svc.toCamel("%sUpdateName", update)).Op("=").Lit(name)
			}
		})
	}

	// add update id expressions
	updateIdExpressions := [][]string{}
	for _, update := range svc.updatesOrdered {
		if svc.methods[update].Desc.Parent() != svc.Service.Desc {
			continue
		}
		opts := svc.updates[update]
		if expr := opts.GetId(); expr != "" {
			updateIdExpressions = append(updateIdExpressions, []string{svc.methods[update].GoName, expr})
		}
	}
	if len(updateIdExpressions) > 0 {
		f.Commentf("%s update id expressions", svc.Service.Desc.FullName())
		f.Var().DefsFunc(func(defs *g.Group) {
			for _, pair := range updateIdExpressions {
				defs.Id(svc.toCamel("%sIDExpression", pair[0])).Op("=").Qual(expressionPkg, "MustParseExpression").Call(g.Lit(pair[1]))
			}
		})
	}
}

// getFieldName returns the name of the go field associated with the field
func (svc *Manifest) getFieldName(field *protogen.Field) string {
	var lintOpts *gopb.LintOptions
	if opts, ok := proto.GetExtension(field.Desc.ParentFile().Options(), gopb.E_Lint).(*gopb.LintOptions); ok && opts != nil {
		lintOpts = opts
	}

	var fieldOpts *gopb.Options
	if opts, ok := proto.GetExtension(field.Desc.Options(), gopb.E_Field).(*gopb.Options); ok && opts != nil {
		fieldOpts = opts
	}

	goName := field.GoName
	if svc.cfg.EnablePatchSupport {
		if n := fieldOpts.GetName(); n != "" {
			goName = n
		}
		if lintOpts.GetAll() || lintOpts.GetFields() {
			goName = lint.Name(goName, lintOpts.InitialismsMap())
		}
	}
	return goName
}

// getMessageName returns the name of the go type associated with the message
func (svc *Manifest) getMessageName(msg *protogen.Message) string {
	var lintOpts *gopb.LintOptions
	if opts, ok := proto.GetExtension(msg.Desc.ParentFile().Options(), gopb.E_Lint).(*gopb.LintOptions); ok && opts != nil {
		lintOpts = opts
	}

	var msgOpts *gopb.Options
	if opts, ok := proto.GetExtension(msg.Desc.Options(), gopb.E_Message).(*gopb.Options); ok && opts != nil {
		msgOpts = opts
	}

	name := msg.GoIdent.GoName
	if svc.cfg.EnablePatchSupport {
		if n := msgOpts.GetName(); n != "" {
			name = n
		}
		if lintOpts.GetAll() || lintOpts.GetMessages() {
			name = lint.Name(name, lintOpts.InitialismsMap())
		}
	}
	return name
}

func (svc *Manifest) goImportPathForMethod(name protoreflect.FullName) string {
	method := svc.methods[name]
	file := svc.serviceFiles[method.Parent.Desc.FullName()]
	if file == nil {
		return ""
	}
	return string(file.GoImportPath)
}

func (svc *Manifest) patchMode(pv temporalv1.Patch_Version, ref protoreflect.FullName) temporalv1.Patch_Mode {
	patchIndex, ok := svc.patchesByRef[ref]
	if ok {
		return patchIndex[pv]
	}
	patchIndex, ok = svc.patchesByRef[ref.Parent()]
	if ok {
		return patchIndex[pv]
	}
	return svc.patches[pv]
}

func (svc *Manifest) render() error {
	for _, file := range svc.Plugin.Files {
		if !file.Generate {
			continue
		}

		f := g.NewFilePathName(string(file.GoImportPath), string(file.GoPackageName))
		genCodeGenerationHeader(svc.Plugin, f, file)
		for pkg, alias := range aliases {
			f.ImportAlias(pkg, alias)
		}

		var xns *g.File
		var xnsGoPackageName, xnsFilePath string
		var hasXNS bool
		if svc.cfg.EnableXNS {
			xnsGoPackageName = fmt.Sprintf("%sxns", file.GoPackageName)
			xnsGoImportPath := path.Join(string(file.GoImportPath), xnsGoPackageName)
			xns = g.NewFilePathName(xnsGoImportPath, xnsGoPackageName)
			genCodeGenerationHeader(svc.Plugin, xns, file)
			for pkg, alias := range aliases {
				xns.ImportAlias(pkg, alias)
			}

			prefixToSlash := filepath.ToSlash(file.GeneratedFilenamePrefix)
			xnsFilePath = path.Join(
				path.Dir(prefixToSlash),
				xnsGoPackageName,
				path.Base(prefixToSlash),
			)
		}

		var hasContent bool
		for _, service := range file.Services {
			var hasTemporalResources bool
			for _, m := range service.Methods {
				if _, ok := svc.activities[m.Desc.FullName()]; ok {
					hasTemporalResources = true
					break
				}
				if _, ok := svc.workflows[m.Desc.FullName()]; ok {
					hasTemporalResources = true
					break
				}
				if _, ok := svc.queries[m.Desc.FullName()]; ok {
					hasTemporalResources = true
					break
				}
				if _, ok := svc.signals[m.Desc.FullName()]; ok {
					hasTemporalResources = true
					break
				}
				if _, ok := svc.updates[m.Desc.FullName()]; ok {
					hasTemporalResources = true
					break
				}
			}
			if !hasTemporalResources {
				continue
			}

			svc.renderService(f, file, service)
			svc.renderTestClient(f)
			if svc.cfg.CliEnabled {
				svc.renderCLI(f)
			}
			if svc.cfg.EnableXNS {
				svc.renderXNS(xns)
				hasXNS = true
			}
			if svc.cfg.EnableCodec {
				svc.renderCodec(f)
			}
			hasContent = true
		}

		if !hasContent {
			continue
		}

		if err := f.Render(svc.Plugin.NewGeneratedFile(fmt.Sprintf("%s_temporal.pb.go", file.GeneratedFilenamePrefix), file.GoImportPath)); err != nil {
			return fmt.Errorf("error rendering file: %w", err)
		}
		if hasXNS {
			if err := xns.Render(svc.Plugin.NewGeneratedFile(
				fmt.Sprintf("%s_xns_temporal.pb.go", xnsFilePath),
				protogen.GoImportPath(path.Join(
					string(file.GoImportPath),
					xnsGoPackageName,
				)),
			)); err != nil {
				return fmt.Errorf("error rendering file: %w", err)
			}
		}
	}

	if svc.cfg.DocsOut != "" {
		if err := renderDocs(svc.Plugin.Plugin, svc.cfg); err != nil {
			fmt.Fprintf(os.Stderr, "error rendering docs: %v", err)
		}
	}
	return nil
}

// renderService writes the temporal service to the given File
func (svc *Manifest) renderService(f *g.File, file *protogen.File, service *protogen.Service) {
	svc.File, svc.Service = file, service
	svc.opts = svc.serviceOptions[service.Desc.FullName()]
	svc.genConstants(f)

	// generate client interface and implementation
	svc.genClientInterface(f)
	svc.genClientImpl(f)
	svc.genClientImplConstructor(f)
	svc.genClientOptions(f)

	// generate client workflow methods
	for _, workflow := range svc.workflowsOrdered {
		if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
			continue
		}
		opts := svc.workflows[workflow]
		svc.genClientImplWorkflowMethod(f, workflow)
		svc.genClientImplWorkflowAsyncMethod(f, workflow)
		svc.genClientImplWorkflowGetMethod(f, workflow)
		for _, signal := range opts.GetSignal() {
			if signal.GetStart() {
				svc.genClientImplSignalWithStartMethod(f, workflow, getFullyQualifiedRef(workflow, signal.GetRef()))
				svc.genClientImplSignalWithStartAsyncMethod(f, workflow, getFullyQualifiedRef(workflow, signal.GetRef()))
			}
		}
	}
	svc.genClientImplWorkflowCancelMethod(f)
	svc.genClientImplWorkflowTerminateMethod(f)

	// generate client query methods
	for _, query := range svc.queriesOrdered {
		if svc.methods[query].Desc.Parent() != svc.Service.Desc {
			continue
		}
		svc.genClientImplQueryMethod(f, query)
	}

	// generate client signal methods
	for _, signal := range svc.signalsOrdered {
		if svc.methods[signal].Desc.Parent() != svc.Service.Desc {
			continue
		}
		svc.genClientImplSignalMethod(f, signal)
	}

	// generate client update methods
	for _, update := range svc.updatesOrdered {
		if svc.methods[update].Desc.Parent() != svc.Service.Desc {
			continue
		}
		svc.genClientImplUpdateMethod(f, update)
		svc.genClientImplUpdateMethodAsync(f, update)
		svc.genClientImplUpdateGetMethod(f, update)
	}

	// generate <Workflow>Options, <Workflow>Run interfaces and implementations used by client
	for _, workflow := range svc.workflowsOrdered {
		if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
			continue
		}
		opts := svc.workflows[workflow]
		svc.genWorkflowOptions(f, workflow, false)
		svc.genClientWorkflowRunInterface(f, workflow)
		svc.genClientWorkflowRunImpl(f, workflow)
		svc.genClientWorkflowRunImplIDMethod(f, workflow)
		svc.genClientWorkflowRunImplRunMethod(f, workflow)
		svc.genClientWorkflowRunImplRunIDMethod(f, workflow)
		svc.genClientWorkflowRunImplCancelMethod(f, workflow)
		svc.genClientWorkflowRunImplGetMethod(f, workflow)
		svc.genClientWorkflowRunImplTerminateMethod(f, workflow)

		// generate query methods
		for _, queryOpts := range opts.GetQuery() {
			svc.genClientWorkflowRunImplQueryMethod(f, workflow, getFullyQualifiedRef(workflow, queryOpts.GetRef()))
		}

		// generate signal methods
		for _, signalOpts := range opts.GetSignal() {
			svc.genClientWorkflowRunImplSignalMethod(f, workflow, getFullyQualifiedRef(workflow, signalOpts.GetRef()))
		}

		// generate update methods
		for _, updateOpts := range opts.GetUpdate() {
			svc.genClientWorkflowRunImplUpdateMethod(f, workflow, getFullyQualifiedRef(workflow, updateOpts.GetRef()))
			svc.genClientWorkflowRunImplUpdateAsyncMethod(f, workflow, getFullyQualifiedRef(workflow, updateOpts.GetRef()))
		}
	}

	// generate <Update>Handle interfaces and implementations used by client
	for _, update := range svc.updatesOrdered {
		if svc.methods[update].Desc.Parent() != svc.Service.Desc {
			continue
		}
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
		if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
			continue
		}
		svc.genWorkerRegisterWorkflow(f, workflow)
		svc.genWorkerBuilderFunction(f, workflow)
		svc.genWorkerWorkflowInput(f, workflow)
		svc.genWorkerWorkflowInterface(f, workflow)
		svc.genWorkerWorkflowChild(f, workflow)
		svc.genWorkerWorkflowChildAsync(f, workflow)
		svc.genWorkflowOptions(f, workflow, true)
		svc.genWorkerWorkflowChildRun(f, workflow)
		svc.genWorkerWorkflowChildRunGet(f, workflow)
		svc.genWorkerWorkflowChildRunSelect(f, workflow)
		svc.genWorkerWorkflowChildRunSelectStart(f, workflow)
		svc.genWorkerWorkflowChildRunWaitStart(f, workflow)
		svc.genWorkerWorkflowChildRunSignals(f, workflow)
	}

	// generate signal types, methods, functions
	for _, signal := range svc.signalsOrdered {
		if svc.methods[signal].Desc.Parent() != svc.Service.Desc {
			continue
		}
		svc.genWorkerSignal(f, signal)
		svc.genWorkerSignalConstructor(f, signal)
		svc.genWorkerSignalReceive(f, signal)
		svc.genWorkerSignalReceiveAsync(f, signal)
		svc.genWorkerSignalReceiveWithTimeout(f, signal)
		svc.genWorkerSignalSelect(f, signal)
		svc.genWorkerSignalExternal(f, signal)
		svc.genWorkerSignalExternalAsync(f, signal)
	}

	// generate activities
	svc.genActivitiesInterface(f)
	svc.genActivityRegisterAllFunction(f)
	for _, activity := range svc.activitiesOrdered {
		if svc.methods[activity].Desc.Parent() != svc.Service.Desc {
			continue
		}
		svc.genActivityRegisterOneFunction(f, activity)
		svc.genActivityFuture(f, activity)
		svc.genActivityFutureGetMethod(f, activity)
		svc.genActivityFutureSelectMethod(f, activity)
		svc.genActivityFunction(f, activity, false, false)
		svc.genActivityFunction(f, activity, false, true)
		svc.genActivityFunction(f, activity, true, false)
		svc.genActivityFunction(f, activity, true, true)
		svc.genActivityOptions(f, activity, false)
		svc.genActivityOptions(f, activity, true)
	}
}

func getFullyQualifiedRef(parent protoreflect.FullName, ref string) protoreflect.FullName {
	if strings.Contains(ref, ".") {
		return protoreflect.FullName(ref)
	}
	return parent.Parent().Append(protoreflect.Name(ref))
}

func (svc *Manifest) methodGoImportPath(m protoreflect.FullName) string {
	return string(svc.files[m].GoImportPath)
}

func (svc *Manifest) methodGoPackageName(m protoreflect.FullName) string {
	return string(svc.files[m].GoPackageName)
}

func (svc *Manifest) methodXNSPackage(m protoreflect.FullName) string {
	goPackageName := svc.methodGoPackageName(m)
	goImportPath := svc.methodGoImportPath(m)
	xnsGoPackageName := fmt.Sprintf("%sxns", goPackageName)
	return path.Join(goImportPath, xnsGoPackageName)
}

func (svc *Manifest) methodsFromSamePackage(a, b protoreflect.FullName) bool {
	return svc.methods[a].Desc.ParentFile().Package() == svc.methods[b].Desc.ParentFile().Package()
}

func (svc *Manifest) methodsFromSameService(a, b protoreflect.FullName) bool {
	return svc.methods[a].Desc.Parent().FullName() == svc.methods[b].Desc.Parent().FullName()
}

func (svc *Manifest) Qual(m protoreflect.FullName, name string) *g.Statement {
	return g.Qual(svc.methodGoImportPath(m), name)
}

func sortFullNames(s []protoreflect.FullName) {
	sort.Slice(s, func(i, j int) bool {
		return string(s[i]) < string(s[j])
	})
}
