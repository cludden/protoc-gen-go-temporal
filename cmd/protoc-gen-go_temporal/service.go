package main

import (
	"fmt"

	temporalpb "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type service struct {
	*protogen.Service
	workflows  []*workflowMethod
	activities []*activityMethod
	queries    []*queryMethod
	signals    []*signalMethod
	calls      []*callMethod
}

// Only checks whether workflow-referenced items are present. Items are
// intentionally left in definition order.
func newService(svc *protogen.Service) (s *service, errs []error) {
	s = &service{Service: svc}
	for _, m := range svc.Methods {
		opts := m.Desc.Options()
		if v := proto.GetExtension(opts, temporalpb.E_Workflow).(*temporalpb.WorkflowOptions); v != nil {
			s.workflows = append(s.workflows, &workflowMethod{Method: m, WorkflowOptions: v})
		}
		if v := proto.GetExtension(opts, temporalpb.E_Activity).(*temporalpb.ActivityOptions); v != nil {
			s.activities = append(s.activities, &activityMethod{m, v})
		}
		if v := proto.GetExtension(opts, temporalpb.E_Query).(*temporalpb.QueryOptions); v != nil {
			s.queries = append(s.queries, &queryMethod{m, v})
		}
		if v := proto.GetExtension(opts, temporalpb.E_Signal).(*temporalpb.SignalOptions); v != nil {
			s.signals = append(s.signals, &signalMethod{m, v})
		}
		if v := proto.GetExtension(opts, temporalpb.E_Call).(*temporalpb.CallOptions); v != nil {
			s.calls = append(s.calls, &callMethod{Method: m, CallOptions: v})
		}
	}

	// Validate all items
	for _, w := range s.workflows {
		errs = append(errs, w.applyExternalsAndValidate(s)...)
	}
	for _, a := range s.activities {
		errs = append(errs, a.validate()...)
	}
	for _, q := range s.queries {
		errs = append(errs, q.validate()...)
	}
	for _, s := range s.signals {
		errs = append(errs, s.validate()...)
	}
	for _, c := range s.calls {
		errs = append(errs, c.applyExternalsAndValidate()...)
	}
	return
}

func (s *service) query(name string) *queryMethod {
	for _, query := range s.queries {
		if string(query.Desc.Name()) == name {
			return query
		}
	}
	return nil
}

func (s *service) signal(name string) *signalMethod {
	for _, signal := range s.signals {
		if string(signal.Desc.Name()) == name {
			return signal
		}
	}
	return nil
}

func (s *service) call(name string) *callMethod {
	for _, call := range s.calls {
		if string(call.Desc.Name()) == name {
			return call
		}
	}
	return nil
}

type workflowMethod struct {
	*protogen.Method
	*temporalpb.WorkflowOptions
	queries         []*queryMethod
	signals         []*signalMethod
	calls           []*callMethod
	signalStart     *signalMethod
	workflowIDField *protogen.Field
}

func (w *workflowMethod) applyExternalsAndValidate(s *service) (errs []error) {
	for _, ref := range w.Query {
		if query := s.query(ref.Ref); query == nil {
			errs = append(errs, fmt.Errorf("query %q on workflow %q not found in service", ref.Ref, w.Desc.Name()))
		} else {
			w.queries = append(w.queries, query)
		}
	}
	for _, ref := range w.Signal {
		if signal := s.signal(ref.Ref); signal == nil {
			errs = append(errs, fmt.Errorf("signal %q on workflow %q not found in service", ref.Ref, w.Desc.Name()))
		} else {
			w.signals = append(w.signals, signal)
		}
	}
	for _, ref := range w.Call {
		if call := s.call(ref.Ref); call == nil {
			errs = append(errs, fmt.Errorf("call %q on workflow %q not found in service", ref.Ref, w.Desc.Name()))
		} else {
			w.calls = append(w.calls, call)
		}
	}
	// Must be in the normal list
	if w.SignalStart != nil {
		for _, signal := range w.signals {
			if string(signal.Desc.Name()) == w.SignalStart.Ref {
				w.signalStart = signal
				break
			}
		}
		if w.signalStart == nil {
			errs = append(errs, fmt.Errorf("signal start %q on workflow %q not in list of signals on workflow",
				w.SignalStart.Ref, w.Desc.Name()))
		}
	}
	// ID must be on the request object
	if w.WorkflowIdField != "" {
		for _, field := range w.Input.Fields {
			if string(field.Desc.Name()) == w.WorkflowIdField {
				w.workflowIDField = field
				break
			}
		}
		if w.workflowIDField == nil {
			errs = append(errs, fmt.Errorf("workflow ID field %q on workflow %q not on request type %v",
				w.WorkflowIdField, w.Desc.Name(), w.Input.Desc.Name()))
		} else if w.workflowIDField.Desc.Kind() != protoreflect.StringKind {
			errs = append(errs, fmt.Errorf("workflow ID field %q on workflow %q is a %v type, not string",
				w.WorkflowIdField, w.Desc.Name(), w.workflowIDField.Desc.Kind()))
		}
	}
	return
}

type activityMethod struct {
	*protogen.Method
	*temporalpb.ActivityOptions
}

func (a *activityMethod) validate() []error {
	// TODO(cretz): This
	return nil
}

type queryMethod struct {
	*protogen.Method
	*temporalpb.QueryOptions
}

func (q *queryMethod) validate() []error {
	// Query must have a response
	if isEmpty(q.Output) {
		return []error{fmt.Errorf("query %q must have return type", q.Desc.Name())}
	}
	return nil
}

type signalMethod struct {
	*protogen.Method
	*temporalpb.SignalOptions
}

func (s *signalMethod) validate() []error {
	// Signal cannot have a response
	if !isEmpty(s.Output) {
		return []error{fmt.Errorf("signal %q must have return type of google.protobuf.Empty", s.Desc.Name())}
	}
	return nil
}

type callMethod struct {
	*protogen.Method
	*temporalpb.CallOptions
	inputIDField                 *protogen.Field
	inputResponseTaskQueueField  *protogen.Field
	inputResponseWorkflowIDField *protogen.Field
	outputIDField                *protogen.Field
}

func (c *callMethod) applyExternalsAndValidate() (errs []error) {
	// Check input
	if isEmpty(c.Input) {
		return []error{fmt.Errorf("call %q must have request type", c.Desc.Name())}
	} else {
		// TODO(cretz): Allow these to be customizable
		const idFieldName = "id"
		const responseTaskQueueFieldName = "response_task_queue"
		const responseWorkflowIDFieldName = "response_workflow_id"
		// Collect fields
		for _, field := range c.Input.Fields {
			if n := string(field.Desc.Name()); n == idFieldName {
				c.inputIDField = field
			} else if n == responseTaskQueueFieldName {
				c.inputResponseTaskQueueField = field
			} else if n == responseWorkflowIDFieldName {
				c.inputResponseWorkflowIDField = field
			}
		}
		if c.inputIDField == nil {
			errs = append(errs, fmt.Errorf("field %q on call %q not on request type %v",
				idFieldName, c.Desc.Name(), c.Input.Desc.Name()))
		} else if c.inputIDField.Desc.Kind() != protoreflect.StringKind {
			errs = append(errs, fmt.Errorf("field %q on call %q is a %v type, not string",
				idFieldName, c.Desc.Name(), c.inputIDField.Desc.Kind()))
		}
		if c.inputResponseTaskQueueField != nil {
			if c.inputResponseTaskQueueField.Desc.Kind() != protoreflect.StringKind {
				errs = append(errs, fmt.Errorf("field %q on call %q is a %v type, not string",
					responseTaskQueueFieldName, c.Desc.Name(), c.inputResponseTaskQueueField.Desc.Kind()))
			}
		} else if c.inputResponseWorkflowIDField == nil {
			errs = append(errs, fmt.Errorf("call %q must have at least one of %q or %q on request type %v",
				c.Desc.Name(), responseTaskQueueFieldName, responseWorkflowIDFieldName, c.Input.Desc.Name()))
		}
		if c.inputResponseWorkflowIDField != nil {
			if c.inputResponseWorkflowIDField.Desc.Kind() != protoreflect.StringKind {
				errs = append(errs, fmt.Errorf("field %q on call %q is a %v type, not string",
					responseWorkflowIDFieldName, c.Desc.Name(), c.inputResponseWorkflowIDField.Desc.Kind()))
			}
		}
	}

	// Check output
	if isEmpty(c.Output) {
		return []error{fmt.Errorf("call %q must have return type", c.Desc.Name())}
	} else {
		// TODO(cretz): Allow this to be customizable
		const idFieldName = "id"
		// Collect fields
		for _, field := range c.Input.Fields {
			if string(field.Desc.Name()) == idFieldName {
				c.outputIDField = field
				break
			}
		}
		if c.outputIDField == nil {
			errs = append(errs, fmt.Errorf("field %q on call %q not on response type %v",
				idFieldName, c.Desc.Name(), c.Input.Desc.Name()))
		} else if c.outputIDField.Desc.Kind() != protoreflect.StringKind {
			errs = append(errs, fmt.Errorf("field %q on call %q is a %v type, not string",
				idFieldName, c.Desc.Name(), c.outputIDField.Desc.Kind()))
		}
	}
	return
}
