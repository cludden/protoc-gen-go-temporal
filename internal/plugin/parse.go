package plugin

import (
	"cmp"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/alta/protopatch/lint"
	"github.com/alta/protopatch/patch/gopb"
	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	j "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// imported packages
const (
	activityPkg      = "go.temporal.io/sdk/activity"
	atomicPkg        = "sync/atomic"
	clientPkg        = "go.temporal.io/sdk/client"
	cliPkg           = "github.com/urfave/cli/v2"
	cliV3Pkg         = "github.com/urfave/cli/v3"
	converterPkg     = "go.temporal.io/sdk/converter"
	convertPkg       = "github.com/cludden/protoc-gen-go-temporal/pkg/convert"
	durationpbPkg    = "google.golang.org/protobuf/types/known/durationpb"
	enumsPkg         = "go.temporal.io/api/enums/v1"
	expressionPkg    = "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	helpersPkg       = "github.com/cludden/protoc-gen-go-temporal/pkg/helpers"
	homedirPkg       = "github.com/mitchellh/go-homedir"
	nexusPkg         = "github.com/nexus-rpc/sdk-go/nexus"
	protojsonPkg     = "google.golang.org/protobuf/encoding/protojson"
	protoreflectPkg  = "google.golang.org/protobuf/reflect/protoreflect"
	serviceerrorPkg  = "go.temporal.io/api/serviceerror"
	temporalnexusPkg = "go.temporal.io/sdk/temporalnexus"
	temporalPkg      = "go.temporal.io/sdk/temporal"
	temporalv1Pkg    = "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	updatePkg        = "go.temporal.io/api/update/v1"
	uuidPkg          = "github.com/google/uuid"
	workerPkg        = "go.temporal.io/sdk/worker"
	workflowPkg      = "go.temporal.io/sdk/workflow"
)

const (
	modeWorkflow int = 1 << iota
	modeActivity
	modeQuery
	modeSignal
	modeUpdate
)

var aliases = map[string]string{
	cliV3Pkg:      "cliv3",
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

	opts                 *temporalv1.ServiceOptions
	activitiesOrdered    []protoreflect.FullName
	activities           map[protoreflect.FullName]*temporalv1.ActivityOptions
	cliFlagUnmarshallers map[string]map[string]struct{}
	commands             map[protoreflect.FullName]*temporalv1.CommandOptions
	files                map[protoreflect.FullName]*protogen.File
	methods              map[protoreflect.FullName]*protogen.Method
	patches              map[temporalv1.Patch_Version]temporalv1.Patch_Mode
	patchesByRef         map[protoreflect.FullName]map[temporalv1.Patch_Version]temporalv1.Patch_Mode
	queriesOrdered       []protoreflect.FullName
	queries              map[protoreflect.FullName]*temporalv1.QueryOptions
	serviceFiles         map[protoreflect.FullName]*protogen.File
	serviceDetails       map[protoreflect.FullName]renderServiceDetails
	serviceOptions       map[protoreflect.FullName]*temporalv1.ServiceOptions
	signalsOrdered       []protoreflect.FullName
	signals              map[protoreflect.FullName]*temporalv1.SignalOptions
	updatesOrdered       []protoreflect.FullName
	updates              map[protoreflect.FullName]*temporalv1.UpdateOptions
	workflowsOrdered     []protoreflect.FullName
	workflows            map[protoreflect.FullName]*temporalv1.WorkflowOptions
}

// parse extracts a Service from a protogen.Service value
func parse(p *Plugin) (*Manifest, error) {
	m := Manifest{
		Plugin:               p,
		activities:           make(map[protoreflect.FullName]*temporalv1.ActivityOptions),
		cliFlagUnmarshallers: make(map[string]map[string]struct{}),
		commands:             make(map[protoreflect.FullName]*temporalv1.CommandOptions),
		files:                make(map[protoreflect.FullName]*protogen.File),
		methods:              make(map[protoreflect.FullName]*protogen.Method),
		patches:              make(map[temporalv1.Patch_Version]temporalv1.Patch_Mode),
		patchesByRef:         make(map[protoreflect.FullName]map[temporalv1.Patch_Version]temporalv1.Patch_Mode),
		queries:              make(map[protoreflect.FullName]*temporalv1.QueryOptions),
		serviceDetails:       make(map[protoreflect.FullName]renderServiceDetails),
		serviceFiles:         make(map[protoreflect.FullName]*protogen.File),
		serviceOptions:       make(map[protoreflect.FullName]*temporalv1.ServiceOptions),
		signals:              make(map[protoreflect.FullName]*temporalv1.SignalOptions),
		updates:              make(map[protoreflect.FullName]*temporalv1.UpdateOptions),
		workflows:            make(map[protoreflect.FullName]*temporalv1.WorkflowOptions),
	}

	// index global patch settings
	for _, p := range strings.Split(m.cfg.Patches, ";") {
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
		m.patches[pv] = pvm
	}

	for _, file := range p.Files {
		if !file.Generate {
			continue
		}

		for _, service := range file.Services {
			details := renderServiceDetails{}
			m.serviceFiles[service.Desc.FullName()] = file
			if opts, ok := proto.GetExtension(service.Desc.Options(), temporalv1.E_Service).(*temporalv1.ServiceOptions); ok && opts != nil {
				m.opts = opts
				m.serviceOptions[service.Desc.FullName()] = opts

				// index service level patch settings
				if len(opts.GetPatches()) > 0 {
					patchIndex := make(map[temporalv1.Patch_Version]temporalv1.Patch_Mode)
					for _, p := range opts.GetPatches() {
						patchIndex[p.GetVersion()] = p.GetMode()
					}
					m.patchesByRef[service.Desc.FullName()] = patchIndex
				}
			}

			for _, method := range service.Methods {
				name := method.Desc.FullName()
				m.methods[name] = method
				m.files[name] = file

				var mode int
				var patches []*temporalv1.Patch
				if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Workflow).(*temporalv1.WorkflowOptions); ok && opts != nil {
					m.workflows[name] = opts
					m.workflowsOrdered = append(m.workflowsOrdered, name)
					details.workflows = append(details.workflows, name)
					mode |= modeWorkflow
					patches = opts.GetPatches()
				}

				if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Activity).(*temporalv1.ActivityOptions); ok && opts != nil {
					m.activities[name] = opts
					m.activitiesOrdered = append(m.activitiesOrdered, name)
					details.activities = append(details.activities, name)
					mode |= modeActivity
				}

				if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Query).(*temporalv1.QueryOptions); ok && opts != nil {
					m.queries[name] = opts
					m.queriesOrdered = append(m.queriesOrdered, name)
					details.queries = append(details.queries, name)
					mode |= modeQuery
					patches = opts.GetPatches()
				}

				if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Signal).(*temporalv1.SignalOptions); ok && opts != nil {
					m.signals[name] = opts
					m.signalsOrdered = append(m.signalsOrdered, name)
					details.signals = append(details.signals, name)
					mode |= modeSignal
					patches = opts.GetPatches()
				}

				if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Update).(*temporalv1.UpdateOptions); ok && opts != nil {
					m.updates[name] = opts
					m.updatesOrdered = append(m.updatesOrdered, name)
					details.updates = append(details.updates, name)
					mode |= modeUpdate
					patches = opts.GetPatches()
				}

				if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Command).(*temporalv1.CommandOptions); ok && opts != nil { //nolint
					m.commands[name] = opts
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
					m.patchesByRef[method.Desc.FullName()] = patchIndex
				}
			}
			slices.SortFunc(details.workflows, func(a, b protoreflect.FullName) int {
				return strings.Compare(string(a), string(b))
			})
			slices.SortFunc(details.activities, func(a, b protoreflect.FullName) int {
				return strings.Compare(string(a), string(b))
			})
			slices.SortFunc(details.queries, func(a, b protoreflect.FullName) int {
				return strings.Compare(string(a), string(b))
			})
			slices.SortFunc(details.signals, func(a, b protoreflect.FullName) int {
				return strings.Compare(string(a), string(b))
			})
			slices.SortFunc(details.updates, func(a, b protoreflect.FullName) int {
				return strings.Compare(string(a), string(b))
			})
			m.serviceDetails[service.Desc.FullName()] = details
		}
	}

	sortFullNames(m.activitiesOrdered)
	sortFullNames(m.queriesOrdered)
	sortFullNames(m.signalsOrdered)
	sortFullNames(m.updatesOrdered)
	sortFullNames(m.workflowsOrdered)

	var errs error
	for _, workflow := range m.workflowsOrdered {
		opts := m.workflows[workflow]

		// ensure workflow queries are defined
		for _, queryOpts := range opts.GetQuery() {
			query := getFullyQualifiedRef(workflow, queryOpts.GetRef())
			if _, ok := m.queries[query]; !ok {
				errs = errors.Join(errs, fmt.Errorf("workflow  %q references undefined query: %q", workflow, query))
			}
		}

		// ensure workflow signals are defined
		for _, signalOpts := range opts.GetSignal() {
			signal := getFullyQualifiedRef(workflow, signalOpts.GetRef())
			if _, ok := m.signals[signal]; !ok {
				errs = errors.Join(errs, fmt.Errorf("workflow  %q references undefined signal: %q", workflow, signal))
			}
		}

		// ensure workflow updates are defined
		for _, updateOpts := range opts.GetUpdate() {
			update := getFullyQualifiedRef(workflow, updateOpts.GetRef())
			if _, ok := m.updates[update]; !ok {
				errs = errors.Join(errs, fmt.Errorf("workflow  %q references undefined update: %q", workflow, update))
			}
		}
	}

	// ensure that signals return no value, unless signal method is also an activity, query, and/or workflow
	for _, signal := range m.signalsOrdered {
		handler := m.methods[signal]
		_, isActivity := m.activities[signal]
		_, isQuery := m.queries[signal]
		_, isUpdate := m.updates[signal]
		_, isWorkflow := m.workflows[signal]
		if !isActivity && !isQuery && !isUpdate && !isWorkflow && !isEmpty(handler.Output) {
			errs = errors.Join(errs, fmt.Errorf("expected signal %q output to be google.protobuf.Empty, got: %s", signal, handler.Output.GoIdent.GoName))
		}
	}
	return &m, errs
}

func (m *Manifest) fqnForActivity(activity protoreflect.FullName) string {
	if fqn := m.activities[activity].GetName(); fqn != "" {
		return fqn
	}
	return string(m.methods[activity].Desc.FullName())
}

func (m *Manifest) fqnForQuery(query protoreflect.FullName) string {
	if fqn := m.queries[query].GetName(); fqn != "" {
		return fqn
	}
	return string(m.methods[query].Desc.FullName())
}

func (m *Manifest) fqnForSignal(signal protoreflect.FullName) string {
	if fqn := m.signals[signal].GetName(); fqn != "" {
		return fqn
	}
	return string(m.methods[signal].Desc.FullName())
}

func (m *Manifest) fqnForUpdate(update protoreflect.FullName) string {
	if fqn := m.updates[update].GetName(); fqn != "" {
		return fqn
	}
	return string(m.methods[update].Desc.FullName())
}

func (m *Manifest) fqnForWorkflow(workflow protoreflect.FullName) string {
	if fqn := m.workflows[workflow].GetName(); fqn != "" {
		return fqn
	}
	return string(m.methods[workflow].Desc.FullName())
}

// genConstants generates constants
func (m *Manifest) genConstants(f *j.File) {
	// add task queue
	if taskQueue := m.opts.GetTaskQueue(); taskQueue != "" {
		name := m.toCamel("%sTaskQueue", m.Service.GoName)
		f.Commentf("%s is the default task-queue for a %s worker", name, m.Service.Desc.FullName())
		f.Var().Id(name).Op("=").Lit(taskQueue)
	}

	// add workflow names
	var workflows []protoreflect.FullName
	for _, workflow := range m.workflowsOrdered {
		if m.methods[workflow].Desc.Parent() != m.Service.Desc {
			continue
		}
		workflows = append(workflows, workflow)
	}
	if len(workflows) > 0 {
		f.Commentf("%s workflow names", m.Service.Desc.FullName())
		f.Const().DefsFunc(func(defs *j.Group) {
			for _, workflow := range workflows {
				method := m.methods[workflow]
				opts := m.workflows[workflow]
				name := opts.GetName()
				if name == "" {
					name = string(method.Desc.FullName())
				}
				defs.Id(m.toCamel("%sWorkflowName", workflow)).Op("=").Lit(name)
			}
		})
	}

	// add workflow id expressions
	workflowIdExpressions := [][]string{}
	for _, workflow := range m.workflowsOrdered {
		if m.methods[workflow].Desc.Parent() != m.Service.Desc {
			continue
		}
		opts := m.workflows[workflow]
		if expr := opts.GetId(); expr != "" {
			workflowIdExpressions = append(workflowIdExpressions, []string{m.methods[workflow].GoName, expr})
		}
	}
	if len(workflowIdExpressions) > 0 {
		f.Commentf("%s workflow id expressions", m.Service.Desc.FullName())
		f.Var().DefsFunc(func(defs *j.Group) {
			for _, pair := range workflowIdExpressions {
				defs.Id(m.toCamel("%sIDExpression", pair[0])).Op("=").Qual(expressionPkg, "MustParseExpression").Call(j.Lit(pair[1]))
			}
		})
	}

	// add workflow search attribute mappings
	workflowSearchAttributes := [][]string{}
	for _, workflow := range m.workflowsOrdered {
		if m.methods[workflow].Desc.Parent() != m.Service.Desc {
			continue
		}
		opts := m.workflows[workflow]
		if mapping := opts.GetSearchAttributes(); mapping != "" {
			workflowSearchAttributes = append(workflowSearchAttributes, []string{m.methods[workflow].GoName, mapping})
		}
	}
	if len(workflowSearchAttributes) > 0 {
		f.Commentf("%s workflow search attribute mappings", m.Service.Desc.FullName())
		f.Var().DefsFunc(func(defs *j.Group) {
			for _, pair := range workflowSearchAttributes {
				defs.Id(m.toCamel("%sSearchAttributesMapping", pair[0])).Op("=").Qual(expressionPkg, "MustParseMapping").Call(j.Lit(pair[1]))
			}
		})
	}

	// add activity names
	var activities []protoreflect.FullName
	for _, activity := range m.activitiesOrdered {
		if m.methods[activity].Desc.Parent() != m.Service.Desc {
			continue
		}
		activities = append(activities, activity)
	}
	if len(activities) > 0 {
		f.Commentf("%s activity names", m.Service.Desc.FullName())
		f.Const().DefsFunc(func(defs *j.Group) {
			for _, activity := range activities {
				method := m.methods[activity]
				opts := m.activities[activity]
				name := opts.GetName()
				if name == "" {
					name = string(method.Desc.FullName())
				}
				defs.Id(m.toCamel("%sActivityName", activity)).Op("=").Lit(name)
			}
		})
	}

	// add query names
	var queries []protoreflect.FullName
	for _, query := range m.queriesOrdered {
		if m.methods[query].Desc.Parent() != m.Service.Desc {
			continue
		}
		queries = append(queries, query)
	}
	if len(queries) > 0 {
		f.Commentf("%s query names", m.Service.Desc.FullName())
		f.Const().DefsFunc(func(defs *j.Group) {
			for _, query := range queries {
				method := m.methods[query]
				opts := m.queries[query]
				name := opts.GetName()
				if name == "" {
					name = string(method.Desc.FullName())
				}
				defs.Id(m.toCamel("%sQueryName", query)).Op("=").Lit(name)
			}
		})
	}

	// add signal names
	var signals []protoreflect.FullName
	for _, signal := range m.signalsOrdered {
		if m.methods[signal].Desc.Parent() != m.Service.Desc {
			continue
		}
		signals = append(signals, signal)
	}
	if len(signals) > 0 {
		f.Commentf("%s signal names", m.Service.Desc.FullName())
		f.Const().DefsFunc(func(defs *j.Group) {
			for _, signal := range signals {
				method := m.methods[signal]
				opts := m.signals[signal]
				name := opts.GetName()
				if name == "" {
					name = string(method.Desc.FullName())
				}
				defs.Id(m.toCamel("%sSignalName", signal)).Op("=").Lit(name)
			}
		})
	}

	// add update names
	var updates []protoreflect.FullName
	for _, update := range m.updatesOrdered {
		if m.methods[update].Desc.Parent() != m.Service.Desc {
			continue
		}
		updates = append(updates, update)
	}
	if len(updates) > 0 {
		f.Commentf("%s update names", m.Service.Desc.FullName())
		f.Const().DefsFunc(func(defs *j.Group) {
			for _, update := range updates {
				method := m.methods[update]
				opts := m.updates[update]
				name := opts.GetName()
				if name == "" {
					name = string(method.Desc.FullName())
				}
				defs.Id(m.toCamel("%sUpdateName", update)).Op("=").Lit(name)
			}
		})
	}

	// add update id expressions
	updateIdExpressions := [][]string{}
	for _, update := range m.updatesOrdered {
		if m.methods[update].Desc.Parent() != m.Service.Desc {
			continue
		}
		opts := m.updates[update]
		if expr := opts.GetId(); expr != "" {
			updateIdExpressions = append(updateIdExpressions, []string{m.methods[update].GoName, expr})
		}
	}
	if len(updateIdExpressions) > 0 {
		f.Commentf("%s update id expressions", m.Service.Desc.FullName())
		f.Var().DefsFunc(func(defs *j.Group) {
			for _, pair := range updateIdExpressions {
				defs.Id(m.toCamel("%sIDExpression", pair[0])).Op("=").Qual(expressionPkg, "MustParseExpression").Call(j.Lit(pair[1]))
			}
		})
	}
}

// getFieldName returns the name of the go field associated with the field
func (m *Manifest) getFieldName(field *protogen.Field) string {
	var lintOpts *gopb.LintOptions
	if opts, ok := proto.GetExtension(field.Desc.ParentFile().Options(), gopb.E_Lint).(*gopb.LintOptions); ok && opts != nil {
		lintOpts = opts
	}

	var fieldOpts *gopb.Options
	if opts, ok := proto.GetExtension(field.Desc.Options(), gopb.E_Field).(*gopb.Options); ok && opts != nil {
		fieldOpts = opts
	}

	goName := field.GoName
	if m.cfg.EnablePatchSupport {
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
func (m *Manifest) getMessageName(msg *protogen.Message) string {
	var lintOpts *gopb.LintOptions
	if opts, ok := proto.GetExtension(msg.Desc.ParentFile().Options(), gopb.E_Lint).(*gopb.LintOptions); ok && opts != nil {
		lintOpts = opts
	}

	var msgOpts *gopb.Options
	if opts, ok := proto.GetExtension(msg.Desc.Options(), gopb.E_Message).(*gopb.Options); ok && opts != nil {
		msgOpts = opts
	}

	name := msg.GoIdent.GoName
	if m.cfg.EnablePatchSupport {
		if n := msgOpts.GetName(); n != "" {
			name = n
		}
		if lintOpts.GetAll() || lintOpts.GetMessages() {
			name = lint.Name(name, lintOpts.InitialismsMap())
		}
	}
	return name
}

func (m *Manifest) goImportPathForMethod(name protoreflect.FullName) string {
	method := m.methods[name]
	file := m.serviceFiles[method.Parent.Desc.FullName()]
	if file == nil {
		return ""
	}
	return string(file.GoImportPath)
}

func (m *Manifest) patchMode(pv temporalv1.Patch_Version, ref protoreflect.FullName) temporalv1.Patch_Mode {
	patchIndex, ok := m.patchesByRef[ref]
	if ok {
		return patchIndex[pv]
	}
	patchIndex, ok = m.patchesByRef[ref.Parent()]
	if ok {
		return patchIndex[pv]
	}
	return m.patches[pv]
}

func (m *Manifest) render() error {
	for _, file := range m.Plugin.Files {
		if !file.Generate {
			continue
		}

		f := j.NewFilePathName(string(file.GoImportPath), string(file.GoPackageName))
		genCodeGenerationHeader(m.Plugin, f, file)
		for pkg, alias := range aliases {
			f.ImportAlias(pkg, alias)
		}

		var xns *j.File
		var xnsGoPackageName, xnsFilePath string
		var hasXNS bool
		if m.cfg.EnableXNS {
			xnsGoPackageName = fmt.Sprintf("%sxns", file.GoPackageName)
			xnsGoImportPath := path.Join(string(file.GoImportPath), xnsGoPackageName)
			xns = j.NewFilePathName(xnsGoImportPath, xnsGoPackageName)
			genCodeGenerationHeader(m.Plugin, xns, file)
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

		var temporalF *j.File
		var temporalGoPackageName, temporalFilePath string
		var hasTemporal bool
		if m.cfg.NexusEnabled {
			temporalGoPackageName = fmt.Sprintf("%stemporal", file.GoPackageName)
			temporalGoImportPath := path.Join(string(file.GoImportPath), temporalGoPackageName)
			temporalF = j.NewFilePathName(temporalGoImportPath, temporalGoPackageName)
			genCodeGenerationHeader(m.Plugin, temporalF, file)
			for pkg, alias := range aliases {
				temporalF.ImportAlias(pkg, alias)
			}

			prefixToSlash := filepath.ToSlash(file.GeneratedFilenamePrefix)
			temporalFilePath = path.Join(
				path.Dir(prefixToSlash),
				temporalGoPackageName,
				path.Base(prefixToSlash),
			)
		}

		var hasContent bool
		for _, service := range file.Services {
			details, ok := m.serviceDetails[service.Desc.FullName()]
			if !ok || details.IsEmpty() {
				continue
			}
			m.renderService(f, file, service)

			if !details.ActivitiesOnly() {
				m.renderTestClient(f)
				if m.cfg.CliEnabled {
					if m.cfg.CliV3Enabled {
						m.renderCLIV3(f)
					} else {
						m.renderCLI(f)
					}
				}

				if m.cfg.EnableXNS {
					m.renderXNS(xns)
					hasXNS = true
				}

				if m.cfg.NexusEnabled {
					hasTemporal = cmp.Or(m.renderNexus(temporalF, file, service), hasTemporal)
				}
			}

			if m.cfg.EnableCodec {
				m.renderCodec(f)
			}
			hasContent = true
		}

		if !hasContent {
			continue
		}

		if err := f.Render(m.Plugin.NewGeneratedFile(fmt.Sprintf("%s_temporal.pb.go", file.GeneratedFilenamePrefix), file.GoImportPath)); err != nil {
			return fmt.Errorf("error rendering file: %w", err)
		}
		if hasXNS {
			if err := xns.Render(m.Plugin.NewGeneratedFile(
				fmt.Sprintf("%s_xns_temporal.pb.go", xnsFilePath),
				protogen.GoImportPath(path.Join(
					string(file.GoImportPath),
					xnsGoPackageName,
				)),
			)); err != nil {
				return fmt.Errorf("error rendering file: %w", err)
			}
		}
		if hasTemporal {
			if err := temporalF.Render(m.Plugin.NewGeneratedFile(
				fmt.Sprintf("%s_temporal.pb.go", temporalFilePath),
				protogen.GoImportPath(path.Join(
					string(file.GoImportPath),
					temporalGoPackageName,
				)),
			)); err != nil {
				return fmt.Errorf("error rendering file: %w", err)
			}
		}
	}

	if m.cfg.DocsOut != "" {
		if err := renderDocs(m.Plugin.Plugin, m.cfg); err != nil {
			fmt.Fprintf(os.Stderr, "error rendering docs: %v", err)
		}
	}
	return nil
}

type renderServiceDetails struct {
	workflows  []protoreflect.FullName
	activities []protoreflect.FullName
	queries    []protoreflect.FullName
	signals    []protoreflect.FullName
	updates    []protoreflect.FullName
}

func (d *renderServiceDetails) IsEmpty() bool {
	return len(d.workflows) == 0 && len(d.activities) == 0 && len(d.queries) == 0 && len(d.signals) == 0 && len(d.updates) == 0
}

func (d *renderServiceDetails) ActivitiesOnly() bool {
	return len(d.activities) > 0 && len(d.workflows) == 0 && len(d.queries) == 0 && len(d.signals) == 0 && len(d.updates) == 0
}

// renderService writes the temporal service to the given File
func (m *Manifest) renderService(f *j.File, file *protogen.File, service *protogen.Service) {
	m.File, m.Service = file, service
	m.opts = m.serviceOptions[service.Desc.FullName()]
	details := m.serviceDetails[service.Desc.FullName()]
	m.genConstants(f)

	// generate client interface and implementation
	if !details.ActivitiesOnly() {
		m.genClientInterface(f)
		m.genClientImpl(f)
		m.genClientImplConstructor(f)
		m.genClientOptions(f)

		// generate client workflow methods
		for _, workflow := range m.workflowsOrdered {
			if m.methods[workflow].Desc.Parent() != m.Service.Desc {
				continue
			}
			opts := m.workflows[workflow]
			m.genClientImplWorkflowMethod(f, workflow)
			m.genClientImplWorkflowMethodAsync(f, workflow)
			m.genClientImplWorkflowGetMethod(f, workflow)
			for _, signal := range opts.GetSignal() {
				if signal.GetStart() {
					m.genClientImplSignalWithStartMethod(f, workflow, getFullyQualifiedRef(workflow, signal.GetRef()))
					m.genClientImplSignalWithStartMethodAsync(f, workflow, getFullyQualifiedRef(workflow, signal.GetRef()))
				}
			}
			for _, update := range opts.GetUpdate() {
				if update.GetStart() {
					m.genClientUpdateWithStartOptions(f, workflow, getFullyQualifiedRef(workflow, update.GetRef()))
					m.genClientImplUpdateWithStartMethod(f, workflow, getFullyQualifiedRef(workflow, update.GetRef()))
					m.genClientImplUpdateWithStartMethodAsync(f, workflow, getFullyQualifiedRef(workflow, update.GetRef()))
				}
			}
		}
		m.genClientImplWorkflowCancelMethod(f)
		m.genClientImplWorkflowTerminateMethod(f)

		// generate client query methods
		for _, query := range m.queriesOrdered {
			if m.methods[query].Desc.Parent() != m.Service.Desc {
				continue
			}
			m.genClientImplQueryMethod(f, query)
		}

		// generate client signal methods
		for _, signal := range m.signalsOrdered {
			if m.methods[signal].Desc.Parent() != m.Service.Desc {
				continue
			}
			m.genClientImplSignalMethod(f, signal)
		}

		// generate client update methods
		for _, update := range m.updatesOrdered {
			if m.methods[update].Desc.Parent() != m.Service.Desc {
				continue
			}
			m.genClientImplUpdateMethod(f, update)
			m.genClientImplUpdateMethodAsync(f, update)
			m.genClientImplUpdateGetMethod(f, update)
		}

		// generate <Workflow>Options, <Workflow>Run interfaces and implementations used by client
		for _, workflow := range m.workflowsOrdered {
			if m.methods[workflow].Desc.Parent() != m.Service.Desc {
				continue
			}
			opts := m.workflows[workflow]
			m.genWorkflowOptions(f, workflow, false)
			m.genClientWorkflowRunInterface(f, workflow)
			m.genClientWorkflowRunImpl(f, workflow)
			m.genClientWorkflowRunImplIDMethod(f, workflow)
			m.genClientWorkflowRunImplRunMethod(f, workflow)
			m.genClientWorkflowRunImplRunIDMethod(f, workflow)
			m.genClientWorkflowRunImplCancelMethod(f, workflow)
			m.genClientWorkflowRunImplGetMethod(f, workflow)
			m.genClientWorkflowRunImplTerminateMethod(f, workflow)

			// generate query methods
			for _, queryOpts := range opts.GetQuery() {
				m.genClientWorkflowRunImplQueryMethod(f, workflow, getFullyQualifiedRef(workflow, queryOpts.GetRef()))
			}

			// generate signal methods
			for _, signalOpts := range opts.GetSignal() {
				m.genClientWorkflowRunImplSignalMethod(f, workflow, getFullyQualifiedRef(workflow, signalOpts.GetRef()))
			}

			// generate update methods
			for _, updateOpts := range opts.GetUpdate() {
				m.genClientWorkflowRunImplUpdateMethod(f, workflow, getFullyQualifiedRef(workflow, updateOpts.GetRef()))
				m.genClientWorkflowRunImplUpdateAsyncMethod(f, workflow, getFullyQualifiedRef(workflow, updateOpts.GetRef()))
			}
		}

		// generate <Update>Handle interfaces and implementations used by client
		for _, update := range m.updatesOrdered {
			if m.methods[update].Desc.Parent() != m.Service.Desc {
				continue
			}
			m.genClientUpdateHandleInterface(f, update)
			m.genClientUpdateHandleImpl(f, update)
			m.genClientUpdateHandleImplWorkflowIDMethod(f, update)
			m.genClientUpdateHandleImplRunIDMethod(f, update)
			m.genClientUpdateHandleImplUpdateIDMethod(f, update)
			m.genClientUpdateHandleImplGetMethod(f, update)
			m.genClientUpdateOptions(f, update)
		}
	}

	if len(details.workflows) > 0 {
		// generate workflows interface and registration helper
		m.genWorkerWorkflowFunctionVars(f)
		m.genWorkerWorkflowsInterface(f)
		m.genWorkerRegisterWorkflows(f)

		// generate workflow types, methods, functions
		for _, workflow := range details.workflows {
			m.genWorkerRegisterWorkflow(f, workflow)
			m.genWorkerBuilderFunction(f, workflow)
			m.genWorkerWorkflowInput(f, workflow)
			m.genWorkerWorkflowInterface(f, workflow)
			m.genWorkerWorkflowChild(f, workflow)
			m.genWorkerWorkflowChildAsync(f, workflow)
			m.genWorkflowOptions(f, workflow, true)
			m.genWorkerWorkflowChildRun(f, workflow)
			m.genWorkerWorkflowChildRunGet(f, workflow)
			m.genWorkerWorkflowChildRunSelect(f, workflow)
			m.genWorkerWorkflowChildRunSelectStart(f, workflow)
			m.genWorkerWorkflowChildRunWaitStart(f, workflow)
			m.genWorkerWorkflowChildRunSignals(f, workflow)
		}
	}

	// generate signal types, methods, functions
	for _, signal := range details.signals {
		m.genWorkerSignal(f, signal)
		m.genWorkerSignalConstructor(f, signal)
		m.genWorkerSignalReceive(f, signal)
		m.genWorkerSignalReceiveAsync(f, signal)
		m.genWorkerSignalReceiveWithTimeout(f, signal)
		m.genWorkerSignalSelect(f, signal)
		m.genWorkerSignalExternal(f, signal)
		m.genWorkerSignalExternalAsync(f, signal)
	}

	// generate activities
	m.genActivitiesInterface(f)
	m.genActivityRegisterAllFunction(f)
	for _, activity := range details.activities {
		m.genActivityRegisterOneFunction(f, activity)
		m.genActivityFuture(f, activity)
		m.genActivityFutureGetMethod(f, activity)
		m.genActivityFutureSelectMethod(f, activity)
		m.genActivityFunction(f, activity, false, false)
		m.genActivityFunction(f, activity, false, true)
		m.genActivityFunction(f, activity, true, false)
		m.genActivityFunction(f, activity, true, true)
		m.genActivityOptions(f, activity, false)
		m.genActivityOptions(f, activity, true)
	}
}

func getFullyQualifiedRef(parent protoreflect.FullName, ref string) protoreflect.FullName {
	if strings.Contains(ref, ".") {
		return protoreflect.FullName(ref)
	}
	return parent.Parent().Append(protoreflect.Name(ref))
}

func (m *Manifest) methodGoImportPath(method protoreflect.FullName) string {
	return string(m.files[method].GoImportPath)
}

func (m *Manifest) methodGoPackageName(method protoreflect.FullName) string {
	return string(m.files[method].GoPackageName)
}

func (m *Manifest) methodXNSPackage(method protoreflect.FullName) string {
	goPackageName := m.methodGoPackageName(method)
	goImportPath := m.methodGoImportPath(method)
	xnsGoPackageName := fmt.Sprintf("%sxns", goPackageName)
	return path.Join(goImportPath, xnsGoPackageName)
}

func (m *Manifest) methodsFromSamePackage(a, b protoreflect.FullName) bool {
	return m.methods[a].Desc.ParentFile().Package() == m.methods[b].Desc.ParentFile().Package()
}

func (m *Manifest) methodsFromSameService(a, b protoreflect.FullName) bool {
	return m.methods[a].Desc.Parent().FullName() == m.methods[b].Desc.Parent().FullName()
}

func (m *Manifest) Qual(method protoreflect.FullName, name string) *j.Statement {
	return j.Qual(m.methodGoImportPath(method), name)
}

func sortFullNames(s []protoreflect.FullName) {
	sort.Slice(s, func(i, j int) bool {
		return string(s[i]) < string(s[j])
	})
}
