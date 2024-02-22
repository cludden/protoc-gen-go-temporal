package plugin

import (
	"errors"
	"fmt"
	"sort"

	"github.com/alta/protopatch/lint"
	"github.com/alta/protopatch/patch/gopb"
	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
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
	commands          map[string]*temporalv1.CommandOptions
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
		commands:   make(map[string]*temporalv1.CommandOptions),
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
		name := strcase.ToCamel(method.GoName)
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

		if opts, ok := proto.GetExtension(method.Desc.Options(), temporalv1.E_Command).(*temporalv1.CommandOptions); ok && opts != nil {
			svc.commands[name] = opts
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

// getFieldName returns the name of the go field associated with the field
func (svc *Service) getFieldName(field *protogen.Field) string {
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
func (svc *Service) getMessageName(msg *protogen.Message) string {
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
