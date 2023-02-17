package main

import "fmt"

func (g *gen) genClient() {
	g.genClientInterface()
	g.genClientOptions()
	g.genCallResponseHandler()
	g.genClientImpl()
	for _, w := range g.svc.workflows {
		g.genWorkflowRun(w)
	}
}

func (g *gen) genClientInterface() {
	g.P(g.svc.Comments.Leading, "type ", g.prefix(), "Client interface {")
	for _, w := range g.svc.workflows {
		g.P()
		g.P(w.Comments.Leading, g.clientExecuteWorkflowSignature(w))
		g.P()
		g.P("// Get", w.GoName, " returns an existing run started by Execute", w.GoName, ".")
		g.P(g.clientGetWorkflowSignature(w))
	}
	for _, q := range g.svc.queries {
		g.P()
		g.P(q.Comments.Leading, g.clientQuerySignature(q))
	}
	for _, s := range g.svc.signals {
		g.P()
		g.P(s.Comments.Leading, g.clientSignalSignature(s))
	}
	for _, c := range g.svc.calls {
		// Only if there is a response task queue field
		if c.inputResponseTaskQueueField != nil {
			g.P()
			g.P(c.Comments.Leading, g.clientCallSignature(c))
		}
	}
	g.P("}")
	g.P()
}

func (g *gen) genClientOptions() {
	g.P("// ", g.prefix(), "ClientOptions are used for New", g.prefix(), "Client.")
	g.P("type ", g.prefix(), "ClientOptions struct {")
	g.P("// Required client.")
	g.P("Client ", g.QualifiedGoIdent(clientPackage.Ident("Client")))
	g.P("// Handler that must be present for client calls to succeed.")
	g.P("CallResponseHandler ", g.prefix(), "CallResponseHandler")
	g.P("}")
	g.P("")
}

func (g *gen) genCallResponseHandler() {
	g.P("// ", g.prefix(), "CallResponseHandler handles activity responses.")
	g.P("type ", g.prefix(), "CallResponseHandler interface {")
	g.P("// TaskQueue returns the task queue for response activities.")
	g.P("TaskQueue() string")
	g.P()
	g.P("// PrepareCall creates a new ID and channels to receive response/error.")
	g.P("// Each channel only has a buffer of one and are never closed and only one is ever sent to.")
	g.P("// If context is closed, the context error is returned on error channel.")
	g.P("PrepareCall(ctx ", g.goContext(), ") (id string, chOk <-chan interface{}, chErr <-chan error)")
	g.P()
	g.P("// AddResponseType adds an activity for the given type and ID field.")
	g.P("// Does not error if activity name already exists for the same params.")
	g.P("AddResponseType(activityName string, typ ", g.QualifiedGoIdent(reflectPackage.Ident("Type")),
		", idField string) error")
	g.P("}")
	g.P()
}

func (g *gen) genClientImpl() {
	g.genClientImplStruct()
	g.genClientImplConstructor()
	g.genClientImplMethods()
}

func (g *gen) genClientImplStruct() {
	g.P("type ", g.privatePrefixed("ClientImpl"), " struct {")
	g.P("client ", g.QualifiedGoIdent(clientPackage.Ident("Client")))
	g.P("callResponseHandler ", g.prefix(), "CallResponseHandler")
	g.P("}")
	g.P()
}

func (g *gen) genClientImplConstructor() {
	g.P("// New", g.prefix(), "Client creates a new ", g.prefix(), "Client.")
	g.P("func New", g.prefix(), "Client(opts ", g.prefix(), "ClientOptions) ", g.prefix(), "Client {")
	g.P(`if opts.Client == nil { panic("missing client") }`)
	g.P("c := &", g.privatePrefixed("clientImpl"), "{client: opts.Client, callResponseHandler: opts.CallResponseHandler}")
	// Add each call's response type
	first := true
	for _, c := range g.svc.calls {
		if c.inputResponseTaskQueueField == nil {
			continue
		}
		if first {
			g.P("if opts.CallResponseHandler != nil {")
			first = false
		}
		g.P("if err := opts.CallResponseHandler.AddResponseType(", g.prefix(), c.GoName, "ResponseName, ",
			g.QualifiedGoIdent(reflectPackage.Ident("TypeOf")), "((*", g.QualifiedGoIdent(c.Output.GoIdent), ")(nil)), ",
			fmt.Sprintf("%q", c.outputIDField.GoName), "); err != nil { panic(err) }")
	}
	if !first {
		g.P("}")
	}
	g.P("return c")
	g.P("}")
	g.P()
}

func (g *gen) genClientImplMethods() {
	for _, w := range g.svc.workflows {
		g.genClientImplWorkflowMethods(w)
	}
	for _, q := range g.svc.queries {
		g.genClientImplQueryMethod(q)
	}
	for _, s := range g.svc.signals {
		g.genClientImplSignalMethod(s)
	}
	for _, c := range g.svc.calls {
		g.genClientImplCallMethod(c)
	}
}

func (g *gen) genClientImplWorkflowMethods(w *workflowMethod) {
	g.P("func (c *", g.privatePrefixed("ClientImpl"), ") ", g.clientExecuteWorkflowSignature(w), " {")
	g.P("if opts == nil { opts = &", g.QualifiedGoIdent(clientPackage.Ident("StartWorkflowOptions")), "{} }")
	// TODO(cretz): Support required and always-overwrite ID options
	if w.workflowIDField != nil {
		g.P(`if opts.ID == "" && req.`, w.workflowIDField.GoName, ` != "" { opts.ID = req.`, w.workflowIDField.GoName, " }")
	}
	// TODO(cretz): More options
	if w.DefaultOptions.GetTaskQueue() != "" {
		g.P(`if opts.TaskQueue == "" { opts.TaskQueue = `, fmt.Sprintf("%q", w.DefaultOptions.GetTaskQueue()), " }")
	}
	reqParam := ""
	if !isEmpty(w.Input) {
		reqParam = ", req"
	}
	if w.signalStart == nil {
		g.P("run, err := c.client.ExecuteWorkflow(ctx, *opts, ", g.prefix(), w.GoName, "Name", reqParam, ")")
	} else {
		g.P("var run ", g.QualifiedGoIdent(clientPackage.Ident("WorkflowRun")))
		g.P("var err error")
		if isEmpty(w.signalStart.Input) {
			g.P("if signalStart {")
			g.P("run, err = c.client.SignalWithStartWorkflow(ctx, opts.ID, ", g.prefix(), w.signalStart.GoName,
				"Name, nil, *opts, ", g.prefix(), w.GoName, "Name", reqParam, ")")
		} else {
			g.P("if signalStart != nil {")
			g.P("run, err = c.client.SignalWithStartWorkflow(ctx, opts.ID, ", g.prefix(), w.signalStart.GoName,
				"Name, signalStart, *opts, ", g.prefix(), w.GoName, "Name", reqParam, ")")
		}
		g.P("} else {")
		g.P("run, err = c.client.ExecuteWorkflow(ctx, *opts, ", g.prefix(), w.GoName, "Name", reqParam, ")")
		g.P("}")
	}
	g.P("if run == nil || err != nil { return nil, err }")
	g.P("return &", g.privatePrefixed(w.GoName+"Run"), "{c, run}, nil")
	g.P("}")
	g.P()
	g.P("func (c *", g.privatePrefixed("ClientImpl"), ") ", g.clientGetWorkflowSignature(w), " {")
	g.P("return &", g.privatePrefixed(w.GoName+"Run"), "{c, c.client.GetWorkflow(ctx, workflowID, runID)}, nil")
	g.P("}")
	g.P()
}

func (g *gen) genClientImplQueryMethod(q *queryMethod) {
	g.P("func (c *", g.privatePrefixed("ClientImpl"), ") ", g.clientQuerySignature(q), " {")
	g.P("var resp ", g.QualifiedGoIdent(q.Output.GoIdent))
	reqParam := ""
	if !isEmpty(q.Input) {
		reqParam = ", req"
	}
	g.P("if val, err := c.client.QueryWorkflow(ctx, workflowID, runID, ", g.prefix(), q.GoName, "Name",
		reqParam, "); err != nil {")
	g.P("return nil, err")
	g.P("} else if err = val.Get(&resp); err != nil {")
	g.P("return nil, err")
	g.P("}")
	g.P("return &resp, nil")
	g.P("}")
	g.P()
}

func (g *gen) genClientImplSignalMethod(s *signalMethod) {
	g.P("func (c *", g.privatePrefixed("ClientImpl"), ") ", g.clientSignalSignature(s), " {")
	reqParam := ", nil"
	if !isEmpty(s.Input) {
		reqParam = ", req"
	}
	g.P("return c.client.SignalWorkflow(ctx, workflowID, runID, ", g.prefix(), s.GoName, "Name", reqParam, ")")
	g.P("}")
	g.P()
}

func (g *gen) genClientImplCallMethod(c *callMethod) {
	g.P("func (c *", g.privatePrefixed("ClientImpl"), ") ", g.clientCallSignature(c), " {")
	g.P("if c.callResponseHandler == nil { return nil, ", g.QualifiedGoIdent(fmtPackage.Ident("Errorf")),
		`("missing response handler") }`)
	g.P("ctx, cancel := ", g.QualifiedGoIdent(contextPackage.Ident("WithCancel")), "(ctx)")
	g.P("defer cancel()")
	g.P("id, chOk, chErr := c.callResponseHandler.PrepareCall(ctx)")
	g.P("req.", c.inputIDField.GoName, " = id")
	g.P("req.", c.inputResponseTaskQueueField.GoName, " = c.callResponseHandler.TaskQueue()")
	g.P("if err := c.client.SignalWorkflow(ctx, workflowID, runID, ", g.prefix(), c.GoName,
		"SignalName, req); err != nil { return nil, err }")
	g.P("select {")
	g.P("case resp := <-chOk: return resp.(*", g.QualifiedGoIdent(c.Output.GoIdent), "), nil")
	g.P("case err := <-chErr: return nil, err")
	g.P("}")
	g.P("}")
	g.P()
}

func (g *gen) clientExecuteWorkflowSignature(w *workflowMethod) string {
	str := "Execute" + w.GoName + "(ctx " + g.goContext() +
		", opts *" + g.QualifiedGoIdent(clientPackage.Ident("StartWorkflowOptions"))
	if !isEmpty(w.Input) {
		str += ", req *" + g.QualifiedGoIdent(w.Input.GoIdent)
	}
	if w.signalStart != nil {
		if isEmpty(w.signalStart.Input) {
			str += ", signalStart bool"
		} else {
			str += ", signalStart *" + g.QualifiedGoIdent(w.signalStart.Input.GoIdent)
		}
	}
	return str + ") (" + w.GoName + "Run, error)"
}

func (g *gen) clientGetWorkflowSignature(w *workflowMethod) string {
	return "Get" + w.GoName + "(ctx " + g.goContext() + ", workflowID, runID string) (" + w.GoName + "Run, error)"
}

func (g *gen) clientQuerySignature(q *queryMethod) string {
	str := q.GoName + "(ctx " + g.goContext() + ", workflowID, runID string"
	if !isEmpty(q.Input) {
		str += ", req *" + g.QualifiedGoIdent(q.Input.GoIdent)
	}
	return str + ") (*" + g.QualifiedGoIdent(q.Output.GoIdent) + ", error)"
}

func (g *gen) clientSignalSignature(s *signalMethod) string {
	str := s.GoName + "(ctx " + g.goContext() + ", workflowID, runID string"
	if !isEmpty(s.Input) {
		str += ", req *" + g.QualifiedGoIdent(s.Input.GoIdent)
	}
	return str + ") error"
}

func (g *gen) clientCallSignature(c *callMethod) string {
	return c.GoName + "(ctx " + g.goContext() +
		", workflowID, runID string, req *" + g.QualifiedGoIdent(c.Input.GoIdent) + ") (*" +
		g.QualifiedGoIdent(c.Output.GoIdent) + ", error)"
}

func (g *gen) genWorkflowRun(w *workflowMethod) {
	g.genWorkflowRunInterface(w)
	g.genWorkflowRunImpl(w)
}

func (g *gen) genWorkflowRunInterface(w *workflowMethod) {
	g.P("// ", g.prefix(), w.GoName, "Run represents an execution of ", g.prefix(), w.GoName, ".")
	g.P("type ", g.prefix(), w.GoName, "Run interface {")
	g.P("// ID is the workflow ID.")
	g.P("ID() string")
	g.P()
	g.P("// RunID is the workflow run ID.")
	g.P("RunID() string")
	g.P()
	g.P("// Get returns the completed workflow value, waiting if necessary.")
	if isEmpty(w.Output) {
		g.P("Get(ctx ", g.goContext(), ") error")
	} else {
		g.P("Get(ctx ", g.goContext(), ") (*", g.QualifiedGoIdent(w.Output.GoIdent), ", error)")
	}
	g.P()
	for _, q := range w.queries {
		g.P()
		g.P(q.Comments.Leading, g.workflowRunQuerySignature(q))
	}
	for _, s := range w.signals {
		g.P()
		g.P(s.Comments.Leading, g.workflowRunSignalSignature(s))
	}
	for _, c := range w.calls {
		// Only if there is a response task queue field
		if c.inputResponseTaskQueueField != nil {
			g.P()
			g.P(c.Comments.Leading, g.workflowRunCallSignature(c))
		}
	}
	g.P("}")
	g.P()
}

func (g *gen) genWorkflowRunImpl(w *workflowMethod) {
	typ := g.privatePrefixed(w.GoName + "Run")
	g.P("type ", typ, " struct {")
	g.P("client *", g.privatePrefixed("ClientImpl"))
	g.P("run ", g.QualifiedGoIdent(clientPackage.Ident("WorkflowRun")))
	g.P("}")
	g.P()
	g.P("func (r *", typ, ") ID() string { return r.run.GetID() }")
	g.P()
	g.P("func (r *", typ, ") RunID() string { return r.run.GetRunID() }")
	g.P()
	if isEmpty(w.Output) {
		g.P("func (r *", typ, ") Get(ctx ", g.goContext(), ") error {")
		g.P("return r.run.Get(ctx, nil)")
	} else {
		g.P("func (r *", typ, ") Get(ctx ", g.goContext(), ") (*", g.QualifiedGoIdent(w.Output.GoIdent), ", error) {")
		g.P("var resp ", g.QualifiedGoIdent(w.Output.GoIdent))
		g.P("if err := r.run.Get(ctx, &resp); err != nil { return nil, err }")
		g.P("return &resp, nil")
	}
	g.P("}")
	g.P()
	for _, q := range w.queries {
		g.P("func (r *", typ, ") ", g.workflowRunQuerySignature(q), " {")
		reqParam := ""
		if !isEmpty(q.Input) {
			reqParam = ", req"
		}
		g.P("return r.client.", q.GoName, `(ctx, r.ID(), ""`, reqParam, ")")
		g.P("}")
		g.P()
	}
	for _, s := range w.signals {
		g.P("func (r *", typ, ") ", g.workflowRunSignalSignature(s), " {")
		reqParam := ""
		if !isEmpty(s.Input) {
			reqParam = ", req"
		}
		g.P("return r.client.", s.GoName, `(ctx, r.ID(), ""`, reqParam, ")")
		g.P("}")
		g.P()
	}
	for _, c := range w.calls {
		g.P("func (r *", typ, ") ", g.workflowRunCallSignature(c), " {")
		g.P("return r.client.", c.GoName, `(ctx, r.ID(), "", req)`)
		g.P("}")
		g.P()
	}
}

func (g *gen) workflowRunQuerySignature(q *queryMethod) string {
	str := q.GoName + "(ctx " + g.goContext()
	if !isEmpty(q.Input) {
		str += ", req *" + g.QualifiedGoIdent(q.Input.GoIdent)
	}
	return str + ") (*" + g.QualifiedGoIdent(q.Output.GoIdent) + ", error)"
}

func (g *gen) workflowRunSignalSignature(s *signalMethod) string {
	str := s.GoName + "(ctx " + g.goContext()
	if !isEmpty(s.Input) {
		str += ", req *" + g.QualifiedGoIdent(s.Input.GoIdent)
	}
	return str + ") error"
}

func (g *gen) workflowRunCallSignature(c *callMethod) string {
	return c.GoName + "(ctx " + g.goContext() + ", req *" + g.QualifiedGoIdent(c.Input.GoIdent) +
		") (*" + g.QualifiedGoIdent(c.Output.GoIdent) + ", error)"
}
