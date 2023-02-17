package main

import (
	"fmt"

	"google.golang.org/protobuf/types/known/durationpb"
)

func (g *gen) genWorker() {
	for _, w := range g.svc.workflows {
		g.genWorkflowImpl(w)
	}
	for _, s := range g.svc.signals {
		g.genSignal(s)
	}
	for _, c := range g.svc.calls {
		g.genCall(c)
	}
	for _, w := range g.svc.workflows {
		g.genWorkflowChild(w)
	}
	for _, s := range g.svc.signals {
		g.genSignalExternal(s)
	}
	for _, c := range g.svc.calls {
		// Only if there is a response workflow ID
		if c.inputResponseWorkflowIDField == nil {
			continue
		}
		g.genCallExternal(c)
	}
	g.genActivities()
}

func (g *gen) genWorkflowImpl(w *workflowMethod) {
	g.genWorkflowImplInterface(w)
	g.genWorkflowInput(w)
	g.genWorkflowWorker(w)
}

func (g *gen) genWorkflowImplInterface(w *workflowMethod) {
	g.P(w.Comments.Leading, "type ", g.prefix(), w.GoName, "Impl interface {")
	if isEmpty(w.Output) {
		g.P("Run(", g.workflowContext(), ") error")
	} else {
		g.P("Run(", g.workflowContext(), ") (*", g.QualifiedGoIdent(w.Output.GoIdent), ", error)")
	}
	g.P()
	for _, q := range w.queries {
		arg := ""
		if !isEmpty(q.Input) {
			arg = "*" + g.QualifiedGoIdent(q.Input.GoIdent)
		}
		g.P(q.GoName, "(", arg, ") (*", g.QualifiedGoIdent(q.Output.GoIdent), ", error)")
		g.P()
	}
	g.P("}")
	g.P()
}

func (g *gen) genWorkflowInput(w *workflowMethod) {
	g.P("// ", g.prefix(), w.GoName, "Input is input provided to ", g.prefix(), w.GoName, "Impl.Run.")
	g.P("type ", g.prefix(), w.GoName, "Input struct {")
	if !isEmpty(w.Input) {
		g.P("Req *", g.QualifiedGoIdent(w.Input.GoIdent))
	}
	for _, s := range w.signals {
		g.P(s.GoName, " ", g.prefix(), s.GoName)
	}
	for _, c := range w.calls {
		g.P(c.GoName, " ", g.prefix(), c.GoName)
	}
	g.P("}")
	g.P()
}

func (g *gen) genWorkflowWorker(w *workflowMethod) {
	newImpl := "newImpl func(" + g.workflowContext() + ", *" +
		g.prefix() + w.GoName + "Input) (" + g.prefix() + w.GoName + "Impl, error)"
	g.P("type ", g.privatePrefixed(w.GoName), "Worker struct{ ", newImpl, " }")
	g.P()
	sig := "(ctx " + g.workflowContext()
	structReq := ""
	if !isEmpty(w.Input) {
		sig += ", req *" + g.QualifiedGoIdent(w.Input.GoIdent)
		structReq = "Req: req"
	}
	sig += ") ("
	if !isEmpty(w.Output) {
		sig += "*" + g.QualifiedGoIdent(w.Output.GoIdent) + ", "
	}
	sig += "error)"
	g.P("func (w ", g.privatePrefixed(w.GoName), "Worker) ", w.GoName, sig, " {")
	g.P("in := &", g.prefix(), w.GoName, "Input{", structReq, "}")
	for _, s := range w.signals {
		g.P("in.", s.GoName, ".Channel = ", g.QualifiedGoIdent(workflowPackage.Ident("GetSignalChannel")), "(ctx, ",
			g.prefix(), s.GoName, "Name)")
	}
	for _, c := range w.calls {
		g.P("in.", c.GoName, ".Channel = ", g.QualifiedGoIdent(workflowPackage.Ident("GetSignalChannel")), "(ctx, ",
			g.prefix(), c.GoName, "SignalName)")
	}
	g.P("impl, err := w.newImpl(ctx, in)")
	if isEmpty(w.Output) {
		g.P("if err != nil { return err }")
	} else {
		g.P("if err != nil { return nil, err }")
	}
	for _, q := range w.queries {
		returnErr := "return err"
		if !isEmpty(w.Output) {
			returnErr = "return nil, err"
		}
		g.P("if err := ", g.QualifiedGoIdent(workflowPackage.Ident("SetQueryHandler")), "(ctx, ", g.prefix(),
			q.GoName, "Name, impl.", q.GoName, "); err != nil { ", returnErr, " }")
	}
	g.P("return impl.Run(ctx)")
	g.P("}")
	g.P()
	g.P("// Build", g.prefix(), w.GoName, " returns a function for the given impl.")
	g.P("func Build", g.prefix(), w.GoName, "(", newImpl, ") func", sig, " {")
	g.P("return ", g.privatePrefixed(w.GoName), "Worker{newImpl}.", w.GoName)
	g.P("}")
	g.P()
	g.P("// Register", g.prefix(), w.GoName, " registers a workflow with the given impl.")
	g.P("func Register", g.prefix(), w.GoName, "(r ", g.QualifiedGoIdent(workerPackage.Ident("WorkflowRegistry")),
		", ", newImpl, ") {")
	g.P("r.RegisterWorkflowWithOptions(Build", g.prefix(), w.GoName, "(newImpl), ",
		g.QualifiedGoIdent(workflowPackage.Ident("RegisterOptions")), "{Name: ", g.prefix(), w.GoName, "Name})")
	g.P("}")
	g.P()
}

func (g *gen) genSignal(s *signalMethod) {
	g.P(s.Comments.Leading, "type ", g.prefix(), s.GoName, " struct{ Channel ",
		g.QualifiedGoIdent(workflowPackage.Ident("ReceiveChannel")), " }")
	g.P()
	if isEmpty(s.Input) {
		g.P("// Receive blocks until signal is received.")
		g.P("func (s ", g.prefix(), s.GoName, ") Receive(ctx ", g.workflowContext(), ") {")
		g.P("s.Channel.Receive(ctx, nil)")
		g.P("}")
		g.P()
		g.P("// ReceiveAsync returns true if signal received or false if not.")
		g.P("func (s ", g.prefix(), s.GoName, ") ReceiveAsync() (received bool) {")
		g.P("return s.Channel.ReceiveAsync(nil)")
		g.P("}")
		g.P()
		g.P("// Select adds the callback to the selector to be invoked when signal received. Callback can be nil.")
		g.P("func (s ", g.prefix(), s.GoName, ") Select(sel ", g.QualifiedGoIdent(workflowPackage.Ident("Selector")),
			", fn func()) ", g.QualifiedGoIdent(workflowPackage.Ident("Selector")), " {")
		g.P("return sel.AddReceive(s.Channel, func(", g.QualifiedGoIdent(workflowPackage.Ident("ReceiveChannel")),
			", bool) {")
		g.P("s.ReceiveAsync()")
		g.P("if fn != nil { fn() }")
		g.P("})")
		g.P("}")
		g.P()
	} else {
		g.P("// Receive blocks until signal is received.")
		g.P("func (s ", g.prefix(), s.GoName, ") Receive(ctx ", g.workflowContext(), ") *",
			g.QualifiedGoIdent(s.Input.GoIdent), " {")
		g.P("var resp ", g.QualifiedGoIdent(s.Input.GoIdent))
		g.P("s.Channel.Receive(ctx, &resp)")
		g.P("return &resp")
		g.P("}")
		g.P()
		g.P("// ReceiveAsync returns received signal or nil if none.")
		g.P("func (s ", g.prefix(), s.GoName, ") ReceiveAsync() *", g.QualifiedGoIdent(s.Input.GoIdent), " {")
		g.P("var resp ", g.QualifiedGoIdent(s.Input.GoIdent))
		g.P("if !s.Channel.ReceiveAsync(&resp) { return nil }")
		g.P("return &resp")
		g.P("}")
		g.P()
		g.P("// Select adds the callback to the selector to be invoked when signal received. Callback can be nil.")
		g.P("func (s ", g.prefix(), s.GoName, ") Select(sel ", g.QualifiedGoIdent(workflowPackage.Ident("Selector")),
			", fn func(*", g.QualifiedGoIdent(s.Input.GoIdent), ")) ",
			g.QualifiedGoIdent(workflowPackage.Ident("Selector")), " {")
		g.P("return sel.AddReceive(s.Channel, func(", g.QualifiedGoIdent(workflowPackage.Ident("ReceiveChannel")),
			", bool) {")
		g.P("req := s.ReceiveAsync()")
		g.P("if fn != nil { fn(req) }")
		g.P("})")
		g.P("}")
		g.P()
	}
}

func (g *gen) genCall(c *callMethod) {
	g.P(c.Comments.Leading, "type ", g.prefix(), c.GoName, " struct{ Channel ",
		g.QualifiedGoIdent(workflowPackage.Ident("ReceiveChannel")), " }")
	g.P()
	g.P("// Receive blocks until call is received.")
	g.P("func (s ", g.prefix(), c.GoName, ") Receive(ctx ", g.workflowContext(), ") *",
		g.QualifiedGoIdent(c.Input.GoIdent), " {")
	g.P("var resp ", g.QualifiedGoIdent(c.Input.GoIdent))
	g.P("s.Channel.Receive(ctx, &resp)")
	g.P("return &resp")
	g.P("}")
	g.P()
	g.P("// ReceiveAsync returns received signal or nil if none.")
	g.P("func (s ", g.prefix(), c.GoName, ") ReceiveAsync() *", g.QualifiedGoIdent(c.Input.GoIdent), " {")
	g.P("var resp ", g.QualifiedGoIdent(c.Input.GoIdent))
	g.P("if !s.Channel.ReceiveAsync(&resp) { return nil }")
	g.P("return &resp")
	g.P("}")
	g.P()
	g.P("// Select adds the callback to the selector to be invoked when signal received. Callback can be nil")
	g.P("func (s ", g.prefix(), c.GoName, ") Select(sel ", g.QualifiedGoIdent(workflowPackage.Ident("Selector")),
		", fn func(*", g.QualifiedGoIdent(c.Input.GoIdent), ")) ",
		g.QualifiedGoIdent(workflowPackage.Ident("Selector")), " {")
	g.P("return sel.AddReceive(s.Channel, func(", g.QualifiedGoIdent(workflowPackage.Ident("ReceiveChannel")),
		", bool) {")
	g.P("req := s.ReceiveAsync()")
	g.P("if fn != nil { fn(req) }")
	g.P("})")
	g.P("}")
	g.P()
	g.P("// Respond sends a response. Activity options not used if request received via")
	g.P("// another workflow. If activity options needed and not present, they are taken")
	g.P("// from the context.")
	g.P("func (s ", g.prefix(), c.GoName, ") Respond(ctx ", g.workflowContext(), ", opts *",
		g.QualifiedGoIdent(workflowPackage.Ident("ActivityOptions")), ", req *",
		g.QualifiedGoIdent(c.Input.GoIdent), ", resp *", g.QualifiedGoIdent(c.Output.GoIdent), ") ",
		g.QualifiedGoIdent(workflowPackage.Ident("Future")), " {")
	g.P("resp.", c.outputIDField.GoName, " = req.", c.inputIDField.GoName)
	if c.inputResponseTaskQueueField != nil && c.inputResponseWorkflowIDField != nil {
		g.P("if req.", c.inputResponseWorkflowIDField.GoName, ` != "" {`)
	}
	if c.inputResponseWorkflowIDField != nil {
		g.P("return ", g.QualifiedGoIdent(workflowPackage.Ident("SignalExternalWorkflow")), "(ctx, req.",
			c.inputResponseWorkflowIDField.GoName, `, "", `, g.prefix(), c.GoName, `ResponseName+"-"+req.`,
			c.inputIDField.GoName, ", resp)")
	}
	if c.inputResponseTaskQueueField != nil && c.inputResponseWorkflowIDField != nil {
		g.P("}")
	}
	if c.inputResponseTaskQueueField != nil {
		g.P("newOpts := ", g.QualifiedGoIdent(workflowPackage.Ident("GetActivityOptions")), "(ctx)")
		g.P("if opts != nil { newOpts = *opts }")
		g.P("newOpts.TaskQueue = req.", c.inputResponseTaskQueueField.GoName)
		g.P("ctx = ", g.QualifiedGoIdent(workflowPackage.Ident("WithActivityOptions")), "(ctx, newOpts)")
		g.P("return ", g.QualifiedGoIdent(workflowPackage.Ident("ExecuteActivity")), "(ctx, ",
			g.prefix(), c.GoName, "ResponseName, resp)")
	}
	g.P("}")
	g.P()
}

func (g *gen) genWorkflowChild(w *workflowMethod) {
	var reqInParam, reqOutParam string
	if !isEmpty(w.Input) {
		reqInParam = ", req *" + g.QualifiedGoIdent(w.Input.GoIdent)
		reqOutParam = ", req"
	}
	g.P("// ", g.prefix(), w.GoName, "Child executes a child workflow.")
	g.P("// If options not present, they are taken from the context.")
	childRun := g.prefix() + w.GoName + "ChildRun"
	g.P("func ", g.prefix(), w.GoName, "Child(ctx ", g.workflowContext(), ", opts *",
		g.QualifiedGoIdent(workflowPackage.Ident("ChildWorkflowOptions")), reqInParam, ") ", childRun, " {")
	g.P("if opts == nil {")
	g.P("ctxOpts := ", g.QualifiedGoIdent(workflowPackage.Ident("GetChildWorkflowOptions")), "(ctx)")
	g.P("opts = &ctxOpts")
	g.P("}")
	if w.workflowIDField != nil {
		g.P(`if opts.WorkflowID == "" && req.`, w.workflowIDField.GoName, ` != "" {`)
		g.P("opts.WorkflowID = req.", w.workflowIDField.GoName)
		g.P("}")
	}
	g.P("ctx = ", g.QualifiedGoIdent(workflowPackage.Ident("WithChildOptions")), "(ctx, *opts)")
	g.P("return ", childRun, "{", g.QualifiedGoIdent(workflowPackage.Ident("ExecuteChildWorkflow")),
		"(ctx, ", g.prefix(), w.GoName, "Name", reqOutParam, ")}")
	g.P("}")
	g.P()
	g.P("// ", childRun, " is a future for the child workflow.")
	g.P("type ", childRun, " struct{ Future ",
		g.QualifiedGoIdent(workflowPackage.Ident("ChildWorkflowFuture")), " }")
	g.P()
	g.P("// WaitStart waits for the child workflow to start.")
	g.P("func (r ", childRun, ") WaitStart(ctx ", g.workflowContext(), ") (*",
		g.QualifiedGoIdent(workflowPackage.Ident("Execution")), ", error) {")
	g.P("var exec ", g.QualifiedGoIdent(workflowPackage.Ident("Execution")))
	g.P("if err := r.Future.GetChildWorkflowExecution().Get(ctx, &exec); err != nil { return nil, err }")
	g.P("return &exec, nil")
	g.P("}")
	g.P()
	g.P("// SelectStart adds waiting for start to the selector. Callback can be nil.")
	g.P("func (r ", childRun, ") SelectStart(sel ", g.QualifiedGoIdent(workflowPackage.Ident("Selector")),
		", fn func(", childRun, ")) ", g.QualifiedGoIdent(workflowPackage.Ident("Selector")), " {")
	g.P("return sel.AddFuture(r.Future.GetChildWorkflowExecution(), func(",
		g.QualifiedGoIdent(workflowPackage.Ident("Future")), ") {")
	g.P("if fn != nil { fn(r) }")
	g.P("})")
	g.P("}")
	g.P()
	g.P("// Get returns the completed workflow value, waiting if necessary.")
	if isEmpty(w.Output) {
		g.P("func (r ", childRun, ") Get(ctx ", g.workflowContext(), ") error {")
		g.P("return r.Future.Get(ctx, nil)")
		g.P("}")
	} else {
		g.P("func (r ", childRun, ") Get(ctx ", g.workflowContext(), ") (*", g.QualifiedGoIdent(w.Output.GoIdent),
			", error) {")
		g.P("var resp ", g.QualifiedGoIdent(w.Output.GoIdent))
		g.P("if err := r.Future.Get(ctx, &resp); err != nil { return nil, err }")
		g.P("return &resp, nil")
		g.P("}")
	}
	g.P()
	g.P("// Select adds this completion to the selector. Callback can be nil.")
	g.P("func (r ", childRun, ") Select(sel ", g.QualifiedGoIdent(workflowPackage.Ident("Selector")),
		", fn func(", g.prefix(), w.GoName, "ChildRun)) ", g.QualifiedGoIdent(workflowPackage.Ident("Selector")), " {")
	g.P("return sel.AddFuture(r.Future, func(", g.QualifiedGoIdent(workflowPackage.Ident("Future")), ") {")
	g.P("if fn != nil { fn(r) }")
	g.P("})")
	g.P("}")
	g.P()
	for _, s := range w.signals {
		reqInParam, reqOutParam := "", ", nil"
		if !isEmpty(s.Input) {
			reqInParam = ", req *" + g.QualifiedGoIdent(s.Input.GoIdent)
			reqOutParam = ", req"
		}
		g.P(s.Comments.Leading, "func (r ", childRun, ") ", s.GoName, "(ctx ", g.workflowContext(), reqInParam, ") ",
			g.QualifiedGoIdent(workflowPackage.Ident("Future")), " {")
		g.P("return r.Future.SignalChildWorkflow(ctx, ", g.prefix(), s.GoName, "Name", reqOutParam, ")")
		g.P("}")
		g.P()
	}
	for _, c := range w.calls {
		// Only if there is a response workflow ID
		if c.inputResponseWorkflowIDField == nil {
			continue
		}
		g.P(c.Comments.Leading, "func (r ", childRun, ") ", c.GoName, "(ctx ", g.workflowContext(), ", req *",
			g.QualifiedGoIdent(c.Input.GoIdent), ") (", g.prefix(), c.GoName, "ResponseExternal, error) {")
		g.P("var resp ", g.prefix(), c.GoName, "ResponseExternal")
		g.P("if req.", c.inputIDField.GoName, ` == "" { return resp, `, g.QualifiedGoIdent(fmtPackage.Ident("Errorf")),
			`("missing request ID") }`)
		if c.inputResponseTaskQueueField != nil {
			g.P("if req.", c.inputResponseTaskQueueField.GoName, ` != "" { return resp, `,
				g.QualifiedGoIdent(fmtPackage.Ident("Errorf")), `("cannot have task queue for child") }`)
		}
		g.P("req.", c.inputResponseWorkflowIDField.GoName, " = ", g.QualifiedGoIdent(workflowPackage.Ident("GetInfo")),
			"(ctx).WorkflowExecution.ID")
		g.P("resp.Channel = ", g.QualifiedGoIdent(workflowPackage.Ident("GetSignalChannel(ctx, ")), g.prefix(),
			c.GoName, `ResponseName+"-"+req.`, c.inputIDField.GoName, ")")
		g.P("resp.Future = r.Future.SignalChildWorkflow(ctx, ", g.prefix(), c.GoName, "SignalName, req)")
		g.P("return resp, nil")
		g.P("}")
		g.P()
	}
}

func (g *gen) genSignalExternal(s *signalMethod) {
	reqInParam, reqOutParam := "", ", nil"
	if !isEmpty(s.Input) {
		reqInParam = ", req *" + g.QualifiedGoIdent(s.Input.GoIdent)
		reqOutParam = ", req"
	}
	g.P(s.Comments.Leading, "func ", g.prefix(), s.GoName, "External(ctx ", g.workflowContext(),
		", workflowID, runID string", reqInParam, ") ", g.QualifiedGoIdent(workflowPackage.Ident("Future")), " {")
	g.P("return ", g.QualifiedGoIdent(workflowPackage.Ident("SignalExternalWorkflow")), "(ctx, workflowID, runID, ",
		g.prefix(), s.GoName, "Name", reqOutParam, ")")
	g.P("}")
	g.P()
}

func (g *gen) genCallExternal(c *callMethod) {
	g.P(c.Comments.Leading, "func ", g.prefix(), c.GoName, "External(ctx ", g.workflowContext(),
		", workflowID, runID string, req *", g.QualifiedGoIdent(c.Input.GoIdent), ") (",
		g.prefix(), c.GoName, "ResponseExternal, error) {")
	g.P("var resp ", g.prefix(), c.GoName, "ResponseExternal")
	g.P("if req.", c.inputIDField.GoName, ` == "" { return resp, `, g.QualifiedGoIdent(fmtPackage.Ident("Errorf")),
		`("missing request ID") }`)
	if c.inputResponseTaskQueueField != nil {
		g.P("if req.", c.inputResponseTaskQueueField.GoName, ` != "" { return resp, `,
			g.QualifiedGoIdent(fmtPackage.Ident("Errorf")), `("cannot have task queue for child") }`)
	}
	g.P("req.", c.inputResponseWorkflowIDField.GoName, " = ", g.QualifiedGoIdent(workflowPackage.Ident("GetInfo")),
		"(ctx).WorkflowExecution.ID")
	g.P("resp.Channel = ", g.QualifiedGoIdent(workflowPackage.Ident("GetSignalChannel(ctx, ")), g.prefix(),
		c.GoName, `ResponseName+"-"+req.`, c.inputIDField.GoName, ")")
	g.P("resp.Future = ", g.QualifiedGoIdent(workflowPackage.Ident("SignalExternalWorkflow")),
		"(ctx, workflowID, runID, ", g.prefix(), c.GoName, "SignalName, req)")
	g.P("return resp, nil")
	g.P("}")
	g.P()
	g.P("// ", g.prefix(), c.GoName, "ResponseExternal represents a call response.")
	g.P("type ", g.prefix(), c.GoName, "ResponseExternal struct {")
	g.P("Future ", g.QualifiedGoIdent(workflowPackage.Ident("Future")))
	g.P("Channel ", g.QualifiedGoIdent(workflowPackage.Ident("ReceiveChannel")))
	g.P("}")
	g.P()
	g.P("// WaitSent blocks until the request is sent.")
	g.P("func (e ", g.prefix(), c.GoName, "ResponseExternal) WaitSent(ctx ", g.workflowContext(), ") error {")
	g.P("return e.Future.Get(ctx, nil)")
	g.P("}")
	g.P()
	g.P("// SelectSent adds when a request is sent to the selector. Callback can be nil.")
	g.P("func (e ", g.prefix(), c.GoName, "ResponseExternal) SelectSent(sel ",
		g.QualifiedGoIdent(workflowPackage.Ident("Selector")), ", fn func(", g.prefix(), c.GoName, "ResponseExternal)) ",
		g.QualifiedGoIdent(workflowPackage.Ident("Selector")), " {")
	g.P("return sel.AddFuture(e.Future, func(", g.QualifiedGoIdent(workflowPackage.Ident("Future")), ") {")
	g.P("if fn != nil { fn(e) }")
	g.P("})")
	g.P("}")
	g.P()
	g.P("// Receive blocks until response is received.")
	g.P("func (e ", g.prefix(), c.GoName, "ResponseExternal) Receive(ctx ", g.workflowContext(), ") *",
		g.QualifiedGoIdent(c.Output.GoIdent), " {")
	g.P("var resp ", g.QualifiedGoIdent(c.Output.GoIdent))
	g.P("e.Channel.Receive(ctx, &resp)")
	g.P("return &resp")
	g.P("}")
	g.P()
	g.P("// ReceiveAsync returns response or nil if none.")
	g.P("func (e ", g.prefix(), c.GoName, "ResponseExternal) ReceiveAsync() *",
		g.QualifiedGoIdent(c.Output.GoIdent), " {")
	g.P("var resp ", g.QualifiedGoIdent(c.Output.GoIdent))
	g.P("if !e.Channel.ReceiveAsync(&resp) { return nil }")
	g.P("return &resp")
	g.P("}")
	g.P()
	g.P("// Select adds the callback to the selector to be invoked when response received. Callback can be nil")
	g.P("func (e ", g.prefix(), c.GoName, "ResponseExternal) Select(sel ",
		g.QualifiedGoIdent(workflowPackage.Ident("Selector")), ", fn func(*", g.QualifiedGoIdent(c.Output.GoIdent), ")) ",
		g.QualifiedGoIdent(workflowPackage.Ident("Selector")), " {")
	g.P("return sel.AddReceive(e.Channel, func(", g.QualifiedGoIdent(workflowPackage.Ident("ReceiveChannel")),
		", bool) {")
	g.P("req := e.ReceiveAsync()")
	g.P("if fn != nil { fn(req) }")
	g.P("})")
	g.P("}")
	g.P()
}

func (g *gen) genActivities() {
	g.genActivitiesImpl()
	for _, a := range g.svc.activities {
		g.genActivity(a)
	}
}

func (g *gen) genActivitiesImpl() {
	if len(g.svc.activities) == 0 {
		return
	}
	g.P("// ", g.prefix(), "ActivitiesImpl is an interface for activity implementations.")
	g.P("type ", g.prefix(), "ActivitiesImpl interface {")
	for _, a := range g.svc.activities {
		g.P()
		g.P(a.Comments.Leading, a.GoName, g.activitySignature(a))
	}
	g.P("}")
	g.P()
	g.P("// Register", g.prefix(), "Activities registers all activities in the interface.")
	g.P("func Register", g.prefix(), "Activities(r ", g.QualifiedGoIdent(workerPackage.Ident("ActivityRegistry")),
		", a ", g.prefix(), "ActivitiesImpl) {")
	for _, a := range g.svc.activities {
		g.P("Register", g.prefix(), a.GoName, "(r, a.", a.GoName, ")")
	}
	g.P("}")
	g.P()
}

func (g *gen) genActivity(a *activityMethod) {
	g.P("// Register", g.prefix(), a.GoName, " registers the single activity.")
	g.P("func Register", g.prefix(), a.GoName, "(r ", g.QualifiedGoIdent(workerPackage.Ident("ActivityRegistry")),
		", impl func", g.activitySignature(a), ") {")
	g.P("r.RegisterActivityWithOptions(impl, ", g.QualifiedGoIdent(activityPackage.Ident("RegisterOptions")),
		"{Name: ", g.prefix(), a.GoName, "Name})")
	g.P("}")
	g.P()
	var reqInParam, reqOutParam string
	if !isEmpty(a.Input) {
		reqInParam = ", req *" + g.QualifiedGoIdent(a.Input.GoIdent)
		reqOutParam = ", req"
	}
	g.P(a.Comments.Leading, "func ", g.prefix(), a.GoName, "(ctx ", g.workflowContext(), ", opts *",
		g.QualifiedGoIdent(workflowPackage.Ident("ActivityOptions")), reqInParam, ") ", g.prefix(), a.GoName, "Future {")
	g.applyActivityOptions(a, false)
	g.P("ctx = ", g.QualifiedGoIdent(workflowPackage.Ident("WithActivityOptions")), "(ctx, *opts)")
	g.P("return ", g.prefix(), a.GoName, "Future{", g.QualifiedGoIdent(workflowPackage.Ident("ExecuteActivity")),
		"(ctx, ", g.prefix(), a.GoName, "Name", reqOutParam, ")}")
	g.P("}")
	g.P()
	g.P(a.Comments.Leading, "func ", g.prefix(), a.GoName, "Local(ctx ", g.workflowContext(), ", opts *",
		g.QualifiedGoIdent(workflowPackage.Ident("LocalActivityOptions")), ", fn func", g.activitySignature(a),
		reqInParam, ") ", g.prefix(), a.GoName, "Future {")
	g.applyActivityOptions(a, true)
	g.P("ctx = ", g.QualifiedGoIdent(workflowPackage.Ident("WithLocalActivityOptions")), "(ctx, *opts)")
	g.P("return ", g.prefix(), a.GoName, "Future{", g.QualifiedGoIdent(workflowPackage.Ident("ExecuteLocalActivity")),
		"(ctx, fn", reqOutParam, ")}")
	g.P("}")
	g.P()
	g.P("// ", g.prefix(), a.GoName, "Future represents completion of the activity.")
	g.P("type ", g.prefix(), a.GoName, "Future struct{ Future ",
		g.QualifiedGoIdent(workflowPackage.Ident("Future")), " }")
	g.P()
	g.P("// Get waits for completion.")
	if isEmpty(a.Output) {
		g.P("func (f ", g.prefix(), a.GoName, "Future) Get(ctx ", g.workflowContext(), ") error {")
		g.P("return f.Future.Get(ctx, nil)")
	} else {
		g.P("func (f ", g.prefix(), a.GoName, "Future) Get(ctx ", g.workflowContext(), ") (*",
			g.QualifiedGoIdent(a.Output.GoIdent), ", error) {")
		g.P("var resp ", g.QualifiedGoIdent(a.Output.GoIdent))
		g.P("if err := f.Future.Get(ctx, &resp); err != nil { return nil, err }")
		g.P("return &resp, nil")
	}
	g.P("}")
	g.P()
	g.P("// Select adds the completion to the selector. Callback can be nil.")
	g.P("func (f ", g.prefix(), a.GoName, "Future) Select(sel ", g.QualifiedGoIdent(workflowPackage.Ident("Selector")),
		", fn func(", g.prefix(), a.GoName, "Future)) ", g.QualifiedGoIdent(workflowPackage.Ident("Selector")), " {")
	g.P("return sel.AddFuture(f.Future, func(", g.QualifiedGoIdent(workflowPackage.Ident("Future")), ") {")
	g.P("if fn != nil { fn(f) }")
	g.P("})")
	g.P("}")
	g.P()
}

func (g *gen) activitySignature(a *activityMethod) string {
	str := "(" + g.goContext()
	if !isEmpty(a.Input) {
		str += ", *" + g.QualifiedGoIdent(a.Input.GoIdent)
	}
	str += ") "
	if !isEmpty(a.Output) {
		str += "(*" + g.QualifiedGoIdent(a.Output.GoIdent) + ", error)"
	} else {
		str += "error"
	}
	return str
}

func (g *gen) applyActivityOptions(a *activityMethod, local bool) {
	g.P("if opts == nil {")
	if local {
		g.P("ctxOpts := ", g.QualifiedGoIdent(workflowPackage.Ident("GetLocalActivityOptions")), "(ctx)")
	} else {
		g.P("ctxOpts := ", g.QualifiedGoIdent(workflowPackage.Ident("GetActivityOptions")), "(ctx)")
	}
	g.P("opts = &ctxOpts")
	g.P("}")
	if v := a.DefaultOptions.GetTaskQueue(); v != "" && !local {
		g.P(`if opts.TaskQueue == "" { opts.TaskQueue = `, fmt.Sprintf("%q", v), " }")
	}
	if v := a.DefaultOptions.GetScheduleToCloseTimeout(); v != nil {
		g.applyActivityDurationOption("ScheduleToCloseTimeout", v)
	}
	if v := a.DefaultOptions.GetScheduleToStartTimeout(); v != nil && !local {
		g.applyActivityDurationOption("ScheduleToStartTimeout", v)
	}
	if v := a.DefaultOptions.GetStartToCloseTimeout(); v != nil {
		g.applyActivityDurationOption("StartToCloseTimeout", v)
	}
	if v := a.DefaultOptions.GetHeartbeatTimeout(); v != nil && !local {
		g.applyActivityDurationOption("HeartbeatTimeout", v)
	}
}

func (g *gen) applyActivityDurationOption(name string, dur *durationpb.Duration) {
	g.P("if opts.", name, " == 0 {")
	g.P("opts.", name, " = ", dur.AsDuration().Nanoseconds(), " // ", dur.AsDuration().String())
	g.P("}")
}
