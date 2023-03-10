package plugin

import (
	"fmt"

	g "github.com/dave/jennifer/jen"
	pgs "github.com/lyft/protoc-gen-star/v2"
)

// genWorkflowsInterface generates a Workflows interface for a given service
func (svc *Service) genWorkflowsInterface(f *g.File) {
	// generate workflows interface
	f.Commentf("Workflows provides methods for initializing new %s workflow values", svc.GoName)
	f.Type().Id("Workflows").InterfaceFunc(func(methods *g.Group) {
		for workflow := range svc.workflows {
			// method := svc.methods[workflow]
			methods.Commentf("%s initializes a new %sWorkflow value", workflow, workflow).Line().
				Id(workflow).
				Params(
					g.Id("ctx").Qual(workflowPkg, "Context"),
					g.Id("input").Op("*").Id(fmt.Sprintf("%sInput", workflow)),
				).
				Params(
					g.Id(workflow),
					g.Error(),
				)
		}
	})
}

// genRegisterWorkflows generates a public RegisterWorkflows method for a given service
func (svc *Service) genRegisterWorkflows(f *g.File) {
	// generate workflow registration function for service
	f.Commentf("RegisterWorkflows registers %s workflows with the given worker", svc.GoName)
	f.Func().
		Id("RegisterWorkflows").
		Params(
			g.Id("r").Qual(workerPkg, "Registry"),
			g.Id("workflows").Id("Workflows"),
		).
		BlockFunc(func(fn *g.Group) {
			for workflow := range svc.workflows {
				fn.Id(fmt.Sprintf("Register%s", workflow)).Call(
					g.Id("r"), g.Id("workflows").Dot(workflow),
				)
			}
		})
}

// genWorkflowWorker generates a <Workflow>Worker struct
func (svc *Service) genWorkflowWorker(f *g.File, workflow string) {
	method := svc.methods[workflow]
	privateName := pgs.Name(method.GoName).LowerCamelCase().String()
	workerName := privateName
	f.Commentf("%s provides an %s method for calling the user's implementation", workerName, method.GoName)
	f.Type().
		Id(workerName).
		Struct(
			g.Id("ctor").
				Func().
				Params(
					g.Qual(workflowPkg, "Context"),
					g.Op("*").Id(fmt.Sprintf("%sInput", method.GoName)),
				).
				Params(
					g.Id(method.GoName),
					g.Error(),
				),
		)
}

// genWorkflowWorkerExecuteMethod generates a <Workflow>Worker's <Workflow> method
func (svc *Service) genWorkflowWorkerExecuteMethod(f *g.File, workflow string) {
	method := svc.methods[workflow]
	opts := svc.workflows[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	privateName := pgs.Name(method.GoName).LowerCamelCase().String()
	workerName := privateName

	// generate <Workflow> method for worker struct
	f.Commentf("%s constructs a new %s value and executes it", method.GoName, method.GoName)
	f.Func().
		Params(g.Id("w").Op("*").Id(workerName)).
		Id(method.GoName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(method.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			// build input struct
			fn.Id("input").Op(":=").Op("&").Id(fmt.Sprintf("%sInput", method.GoName)).BlockFunc(func(fields *g.Group) {
				if hasInput {
					fields.Id("Req").Op(":").Id("req").Op(",")
				}
				for _, s := range opts.GetSignal() {
					signal := s.GetRef()
					fields.Id(signal).Op(":").Op("&").Id(signal).Block(
						g.Id("Channel").Op(":").Qual(workflowPkg, "GetSignalChannel").Call(
							g.Id("ctx"), g.Id(fmt.Sprintf("%sName", signal)),
						).Op(","),
					).Op(",")
				}
			})

			// call constructor to get workflow implementation
			fn.List(g.Id("wf"), g.Err()).Op(":=").Id("w").Dot("ctor").Call(
				g.Id("ctx"), g.Id("input"),
			)
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)

			// register query handlers
			for _, q := range opts.GetQuery() {
				query := q.GetRef()
				fn.If(
					g.Err().Op(":=").Qual(workflowPkg, "SetQueryHandler").Call(
						g.Id("ctx"), g.Id(fmt.Sprintf("%sName", query)), g.Id("wf").Dot(query),
					),
					g.Err().Op("!=").Nil(),
				).Block(
					g.ReturnFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Err()
					}),
				)
			}

			// execute workflow
			fn.Return(
				g.Id("wf").Dot("Execute").Call(g.Id("ctx")),
			)
		})
}

// genWorkflowWorkerBuilderFunction generates a build<Workflow> function that converts
// a constructor function or method into a valid workflow function
func (svc *Service) genWorkflowWorkerBuilderFunction(f *g.File, workflow string) {
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	privateName := pgs.Name(method.GoName).LowerCamelCase().String()
	workerName := privateName
	builderName := fmt.Sprintf("build%s", method.GoName)

	// generate Build<Workflow> function
	f.Commentf("%s converts a %s workflow struct into a valid workflow function", builderName, method.GoName)
	f.Func().
		Id(builderName).
		Params(
			g.Id("wf").
				Func().
				Params(
					g.Qual(workflowPkg, "Context"),
					g.Op("*").Id(fmt.Sprintf("%sInput", method.GoName)),
				).
				Params(
					g.Id(method.GoName),
					g.Error(),
				),
		).
		Params(
			g.Func().
				ParamsFunc(func(args *g.Group) {
					args.Qual(workflowPkg, "Context")
					if hasInput {
						args.Op("*").Id(method.Input.GoIdent.GoName)
					}
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Id(method.Output.GoIdent.GoName)
					}
					returnVals.Error()
				}),
		).
		Block(
			g.Return(
				g.Parens(g.Op("&").Id(workerName).Values(g.Id("wf"))).Dot(method.GoName),
			),
		)
}

// genRegisterWorkflow generates a Register<Workflow> public function
func (svc *Service) genRegisterWorkflow(f *g.File, workflow string) {
	method := svc.methods[workflow]
	builderName := fmt.Sprintf("build%s", method.GoName)

	// generate Register<Workflow> function
	f.Commentf("Register%s registers a %s workflow with the given worker", method.GoName, method.GoName)
	f.Func().
		Id(fmt.Sprintf("Register%s", method.GoName)).
		Params(
			g.Id("r").Qual(workerPkg, "Registry"),
			g.Id("wf").
				Func().
				Params(
					g.Qual(workflowPkg, "Context"),
					g.Op("*").Id(fmt.Sprintf("%sInput", method.GoName)),
				).
				Params(
					g.Id(method.GoName),
					g.Error(),
				),
			// g.Id("wf").Id("Workflows"),
		).
		Block(
			g.Id("r").Dot("RegisterWorkflowWithOptions").Call(
				g.Id(builderName).Call(g.Id("wf")),
				g.Qual(workflowPkg, "RegisterOptions").Values(
					g.Id("Name").Op(":").Id(fmt.Sprintf("%sName", method.GoName)),
				),
			),
		)
}

// genWorkflowInterface generates a <Workflow> interface
func (svc *Service) genWorkflowInterface(f *g.File, workflow string) {
	opts := svc.workflows[workflow]
	method := svc.methods[workflow]
	hasOutput := !isEmpty(method.Output)
	// generate workflow interface
	f.Commentf("%s describes a %s workflow implementation", workflow, workflow)
	f.Type().Id(workflow).InterfaceFunc(func(methods *g.Group) {
		methods.Commentf("Execute a %s workflow", workflow).Line().
			Id("Execute").
			Params(
				g.Id("ctx").Qual(workflowPkg, "Context"),
			).
			ParamsFunc(func(returnVals *g.Group) {
				if hasOutput {
					returnVals.Op("*").Id(method.Output.GoIdent.GoName)
				}
				returnVals.Error()
			})

		// add workflow query methods
		for _, queryOpts := range opts.GetQuery() {
			query := queryOpts.GetRef()
			handler := svc.methods[query]
			hasInput := !isEmpty(handler.Input)
			methods.Commentf("%s query handler", query)
			methods.Id(query).
				ParamsFunc(func(args *g.Group) {
					if hasInput {
						args.Op("*").Id(handler.Input.GoIdent.GoName)
					}
				}).
				Params(
					g.Op("*").Id(handler.Output.GoIdent.GoName),
					g.Error(),
				)
		}
	})
}

// genWorkflowInput generates a <Workflow>Input struct
func (svc *Service) genWorkflowInput(f *g.File, workflow string) {
	opts := svc.workflows[workflow]
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	f.Commentf("%sInput describes the input to a %s workflow constructor", workflow, workflow)
	f.Type().Id(fmt.Sprintf("%sInput", workflow)).StructFunc(func(fields *g.Group) {
		if hasInput {
			fields.Id("Req").Op("*").Id(method.Input.GoIdent.GoName)
		}

		// add workflow signals
		for _, signalOpts := range opts.GetSignal() {
			signal := signalOpts.GetRef()
			fields.Id(signal).Op("*").Id(signal)
		}
	})
}

// genExecuteChildWorkflow generates a public <Workflow>Child function
func (svc *Service) genExecuteChildWorkflow(f *g.File, workflow string) {
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	f.Commentf("%sChild executes a child %s workflow", workflow, workflow)
	f.Func().
		Id(fmt.Sprintf("%sChild", workflow)).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("opts").Op("*").Qual(workflowPkg, "ChildWorkflowOptions")
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
		}).
		Id(fmt.Sprintf("%sChildRun", workflow)).
		BlockFunc(func(fn *g.Group) {
			// initialize child workflow options with default values
			svc.genStartWorkflowOptions(fn, workflow, true)

			fn.Id("ctx").Op("=").Qual(workflowPkg, "WithChildOptions").Call(g.Id("ctx"), g.Op("*").Id("opts"))
			fn.Return(
				g.Id(fmt.Sprintf("%sChildRun", workflow)).Block(
					g.Id("Future").Op(":").Qual(workflowPkg, "ExecuteChildWorkflow").CallFunc(func(args *g.Group) {
						args.Id("ctx")
						args.Lit(fmt.Sprintf("%sName", workflow))
						if hasInput {
							args.Id("req")
						} else {
							args.Nil()
						}
					}).Op(","),
				),
			)
		})
}

// genWorkflowChildRun generates a <Workflow>ChildRun struct
func (svc *Service) genWorkflowChildRun(f *g.File, workflow string) {
	// generate child workflow run struct
	f.Commentf("%sChildRun describes a child %s workflow run", workflow, workflow)
	f.Type().Id(fmt.Sprintf("%sChildRun", workflow)).StructFunc(func(fields *g.Group) {
		fields.Add(g.Id("Future").Qual(workflowPkg, "ChildWorkflowFuture"))
	})
}

// genWorkflowChildRunGet generates a <Workflow>ChildRun Get method
func (svc *Service) genWorkflowChildRunGet(f *g.File, workflow string) {
	method := svc.methods[workflow]
	hasOutput := !isEmpty(method.Output)
	f.Comment("Get blocks until the workflow is completed, returning the response value")
	f.Func().
		Params(
			g.Id("r").Op("*").Id(fmt.Sprintf("%sChildRun", workflow)),
		).
		Id("Get").
		Params(
			g.Id("ctx").Qual(workflowPkg, "Context"),
		).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(method.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			if hasOutput {
				fn.Var().Id("resp").Id(method.Output.GoIdent.GoName)
			}
			fn.If(
				g.Err().Op(":=").Id("r").Dot("Future").Dot("Get").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					if hasOutput {
						args.Op("&").Id("resp")
					} else {
						args.Nil()
					}
				}),
				g.Err().Op("!=").Nil(),
			).BlockFunc(func(b *g.Group) {
				if hasOutput {
					b.Return(g.Nil(), g.Err())
				} else {
					b.Return(g.Err())
				}
			})
			if hasOutput {
				fn.Return(g.Op("&").Id("resp"), g.Nil())
			} else {
				fn.Return(g.Nil())
			}
		})
}

// genWorkflowChildRunSelect generates a <Workflow>ChildRun Select method
func (svc *Service) genWorkflowChildRunSelect(f *g.File, workflow string) {
	f.Comment("Select adds this completion to the selector. Callback can be nil.")
	f.Func().
		Params(
			g.Id("r").Op("*").Id(fmt.Sprintf("%sChildRun", workflow)),
		).
		Id("Select").
		Params(
			g.Id("sel").Qual(workflowPkg, "Selector"),
			g.Id("fn").Func().Params(g.Id(fmt.Sprintf("%sChildRun", workflow))),
		).
		Params(
			g.Qual(workflowPkg, "Selector"),
		).
		Block(
			g.Return(
				g.Id("sel").Dot("AddFuture").Call(
					g.Id("r").Dot("Future"),
					g.Func().Params(g.Qual(workflowPkg, "Future")).Block(
						g.If(g.Id("fn").Op("!=").Nil()).Block(
							g.Id("fn").Call(g.Op("*").Id("r")),
						),
					),
				),
			),
		)
}

// genWorkflowChildRunSelectStart generates a <Workflow>ChildRun SelectStart method
func (svc *Service) genWorkflowChildRunSelectStart(f *g.File, workflow string) {
	f.Comment("SelectStart adds waiting for start to the selector. Callback can be nil.")
	f.Func().
		Params(
			g.Id("r").Op("*").Id(fmt.Sprintf("%sChildRun", workflow)),
		).
		Id("SelectStart").
		Params(
			g.Id("sel").Qual(workflowPkg, "Selector"),
			g.Id("fn").Func().Params(g.Id(fmt.Sprintf("%sChildRun", workflow))),
		).
		Params(
			g.Qual(workflowPkg, "Selector"),
		).
		Block(
			g.Return(
				g.Id("sel").Dot("AddFuture").Call(
					g.Id("r").Dot("Future").Dot("GetChildWorkflowExecution").Call(),
					g.Func().Params(g.Qual(workflowPkg, "Future")).Block(
						g.If(g.Id("fn").Op("!=").Nil()).Block(
							g.Id("fn").Call(g.Op("*").Id("r")),
						),
					),
				),
			),
		)
}

// genWorkflowChildRunWaitStart generates a <Workflow>ChildRun WaitStart method
func (svc *Service) genWorkflowChildRunWaitStart(f *g.File, workflow string) {
	f.Comment("WaitStart waits for the child workflow to start")
	f.Func().
		Params(
			g.Id("r").Op("*").Id(fmt.Sprintf("%sChildRun", workflow)),
		).
		Id("WaitStart").
		Params(
			g.Id("ctx").Qual(workflowPkg, "Context"),
		).
		Params(
			g.Op("*").Qual(workflowPkg, "Execution"),
			g.Error(),
		).
		Block(
			g.Var().Id("exec").Qual(workflowPkg, "Execution"),
			g.If(
				g.Err().Op(":=").Id("r").Dot("Future").Dot("GetChildWorkflowExecution").Call().Dot("Get").Call(
					g.Id("ctx"),
					g.Op("&").Id("exec"),
				),
				g.Err().Op("!=").Nil(),
			).Block(
				g.Return(g.Nil(), g.Err()),
			),
			g.Return(g.Op("&").Id("exec"), g.Nil()),
		)
}

// genWorkflowChildRunSignals generates <Workflow>ChildRun signal methods
func (svc *Service) genWorkflowChildRunSignals(f *g.File, workflow string) {
	opts := svc.workflows[workflow]
	for _, signalOpts := range opts.GetSignal() {
		signal := signalOpts.GetRef()
		handler := svc.methods[signal]
		hasInput := !isEmpty(handler.Input)
		f.Commentf("%s sends the corresponding signal request to the child workflow", signal)
		f.Func().
			Params(g.Id("r").Op("*").Id(fmt.Sprintf("%sChildRun", workflow))).
			Id(signal).
			ParamsFunc(func(params *g.Group) {
				params.Id("ctx").Qual(workflowPkg, "Context")
				if hasInput {
					params.Id("input").Op("*").Id(handler.Input.GoIdent.GoName)
				}
			}).
			Params(g.Qual(workflowPkg, "Future")).
			Block(
				g.Return(g.Id("r").Dot("Future").Dot("SignalChildWorkflow").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id(fmt.Sprintf("%sName", signal))
					if hasInput {
						args.Id("input")
					} else {
						args.Nil()
					}
				})),
			)
	}
}

// genWorkerSignal generates a worker signal struct
func (svc *Service) genWorkerSignal(f *g.File, signal string) {
	f.Commentf("%s describes a %s signal", signal, signal)
	f.Type().Id(signal).Struct(
		g.Id("Channel").Qual(workflowPkg, "ReceiveChannel"),
	)
}

// genWorkerSignalReceive generates a worker signal Receive method
func (svc *Service) genWorkerSignalReceive(f *g.File, signal string) {
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)
	f.Commentf("Receive blocks until a %s signal is received", signal)
	f.Func().
		Params(g.Id("s").Op("*").Id(signal)).
		Id("Receive").
		Params(g.Id("ctx").Qual(workflowPkg, "Context")).
		ParamsFunc(func(returnVals *g.Group) {
			if hasInput {
				returnVals.Op("*").Id(method.Input.GoIdent.GoName)
			}
			returnVals.Bool()
		}).
		BlockFunc(func(b *g.Group) {
			if hasInput {
				b.Var().Id("resp").Id(method.Input.GoIdent.GoName)
			}
			b.Id("more").Op(":=").Id("s").Dot("Channel").Dot("Receive").CallFunc(func(args *g.Group) {
				args.Id("ctx")
				if hasInput {
					args.Op("&").Id("resp")
				} else {
					args.Nil()
				}
			})
			b.ReturnFunc(func(returnVals *g.Group) {
				if hasInput {
					returnVals.Op("&").Id("resp")
				}
				returnVals.Id("more")
			})
		})
}

// genWorkerSignalReceiveAsync generates a worker signal ReceiveAsync method
func (svc *Service) genWorkerSignalReceiveAsync(f *g.File, signal string) {
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)
	f.Commentf("ReceiveAsync checks for a %s signal without blocking", signal)
	f.Func().
		Params(g.Id("s").Op("*").Id(signal)).
		Id("ReceiveAsync").
		Params().
		ParamsFunc(func(returnVals *g.Group) {
			if hasInput {
				returnVals.Op("*").Id(method.Input.GoIdent.GoName)
			} else {
				returnVals.Bool()
			}
		}).
		BlockFunc(func(b *g.Group) {
			if hasInput {
				b.Var().Id("resp").Id(method.Input.GoIdent.GoName)
				b.If(
					g.Id("ok").Op(":=").Id("s").Dot("Channel").Dot("ReceiveAsync").Call(
						g.Op("&").Id("resp"),
					),
					g.Op("!").Id("ok"),
				).Block(
					g.Return(g.Nil()),
				)
				b.Return(g.Op("&").Id("resp"))
			} else {
				b.Return(g.Id("s").Dot("Channel").Dot("ReceiveAsync").Call(g.Nil()))
			}
		})
}

// genWorkerSignalSelect generates a worker signal Select method
func (svc *Service) genWorkerSignalSelect(f *g.File, signal string) {
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)
	f.Commentf("Select checks for a %s signal without blocking", signal)
	f.Func().
		Params(g.Id("s").Op("*").Id(signal)).
		Id("Select").
		Params(
			g.Id("sel").Qual(workflowPkg, "Selector"),
			g.Id("fn").Func().ParamsFunc(func(args *g.Group) {
				if hasInput {
					args.Op("*").Id(method.Input.GoIdent.GoName)
				}
			}),
		).
		Params(
			g.Qual(workflowPkg, "Selector"),
		).
		Block(
			g.Return(
				g.Id("sel").Dot("AddReceive").Call(
					g.Id("s").Dot("Channel"),
					g.Func().
						Params(
							g.Qual(workflowPkg, "ReceiveChannel"),
							g.Bool(),
						).
						BlockFunc(func(fn *g.Group) {
							if hasInput {
								fn.Id("req").Op(":=").Id("s").Dot("ReceiveAsync").Call()
							} else {
								fn.Id("s").Dot("ReceiveAsync").Call()
							}
							fn.If(g.Id("fn").Op("!=").Nil()).Block(
								g.Id("fn").CallFunc(func(args *g.Group) {
									if hasInput {
										args.Id("req")
									}
								}),
							)
						}),
				),
			),
		)
}

// genWorkerSignalExternal generates a <Signal>External public function
func (svc *Service) genWorkerSignalExternal(f *g.File, signal string) {
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)
	f.Commentf("%sExternal sends a %s signal to an existing workflow", signal, signal)
	f.Func().Id(fmt.Sprintf("%sExternal", signal)).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
		}).
		Params(g.Qual(workflowPkg, "Future")).
		Block(
			g.Return(
				g.Qual(workflowPkg, "SignalExternalWorkflow").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("workflowID")
					args.Id("runID")
					args.Id(fmt.Sprintf("%sName", signal))
					if hasInput {
						args.Id("req")
					} else {
						args.Nil()
					}
				}),
			),
		)
}
