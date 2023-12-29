package plugin

import (
	"fmt"
	"strings"

	g "github.com/dave/jennifer/jen"
)

// genWorkerBuilderFunction generates a build<Workflow> function that converts
// a constructor function or method into a valid workflow function
func (svc *Service) genWorkerBuilderFunction(f *g.File, workflow string) {
	method := svc.methods[workflow]
	opts := svc.workflows[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	builderName := fmt.Sprintf("build%s", method.GoName)

	// generate Build<Workflow> function
	f.Commentf("%s converts a %s workflow struct into a valid workflow function", builderName, method.GoName)
	f.Func().
		Id(builderName).
		Params(
			g.Id("ctor").
				Func().
				ParamsFunc(func(args *g.Group) {
					args.Qual(workflowPkg, "Context")
					if svc.cfg.DisableWorkflowInputRename {
						args.Op("*").Id(toCamel("%sInput", workflow))
					} else {
						args.Op("*").Id(toCamel("%sWorkflowInput", workflow))
					}
				}).
				Params(
					g.Id(toCamel("%sWorkflow", workflow)),
					g.Error(),
				),
		).
		Params(
			g.Func().
				ParamsFunc(func(args *g.Group) {
					args.Qual(workflowPkg, "Context")
					if hasInput {
						args.Op("*").Id(svc.getMessageName(method.Input))
					}
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Id(svc.getMessageName(method.Output))
					}
					returnVals.Error()
				}),
		).
		Block(
			g.Return(
				//g.Parens(g.Op("&").Id(workerName).Values(g.Id("wf"))).Dot(method.GoName),
				// generate <Workflow> method for worker struct
				g.Func().
					ParamsFunc(func(args *g.Group) {
						args.Id("ctx").Qual(workflowPkg, "Context")
						if hasInput {
							args.Id("req").Op("*").Id(svc.getMessageName(method.Input))
						}
					}).
					ParamsFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Op("*").Id(svc.getMessageName(method.Output))
						}
						returnVals.Error()
					}).
					BlockFunc(func(fn *g.Group) {
						// build input struct
						inputType := toCamel("%sWorkflowInput", workflow)
						if svc.cfg.DisableWorkflowInputRename {
							inputType = toCamel("%sInput", workflow)
						}
						fn.Id("input").Op(":=").Op("&").Id(inputType).BlockFunc(func(fields *g.Group) {
							if hasInput {
								fields.Id("Req").Op(":").Id("req").Op(",")
							}
							for _, s := range opts.GetSignal() {
								signal := s.GetRef()
								fields.Id(signal).Op(":").Op("&").Id(toCamel("%sSignal", signal)).Block(
									g.Id("Channel").Op(":").Qual(workflowPkg, "GetSignalChannel").Call(
										g.Id("ctx"), g.Id(toCamel("%sSignalName", signal)),
									).Op(","),
								).Op(",")
							}
						})

						// call constructor to get workflow implementation
						fn.List(g.Id("wf"), g.Err()).Op(":=").Id("ctor").Call(
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

						fn.If(
							g.List(g.Id("initializable"), g.Id("ok")).Op(":=").Id("wf").Op(".").Parens(g.Qual(helpersPkg, "Initializable")),
							g.Id("ok"),
						).Block(
							g.If(g.Err().Op(":=").Id("initializable").Dot("Initialize").Call(g.Id("ctx")), g.Err().Op("!=").Nil()).Block(
								g.ReturnFunc(func(returnVals *g.Group) {
									if hasOutput {
										returnVals.Nil()
									}
									returnVals.Err()
								}),
							),
						)

						// register query handlers
						for _, q := range opts.GetQuery() {
							query := q.GetRef()
							fn.If(
								g.Err().Op(":=").Qual(workflowPkg, "SetQueryHandler").Call(
									g.Id("ctx"), g.Id(toCamel("%sQueryName", query)), g.Id("wf").Dot(query),
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

						// register update handlers
						for _, u := range opts.GetUpdate() {
							update := u.GetRef()
							updateOpts := svc.updates[update]

							fn.BlockFunc(func(b *g.Group) {
								// build UpdateHandlerOptions
								var updateHandlerOptions []g.Code
								if updateOpts.GetValidate() {
									updateHandlerOptions = append(updateHandlerOptions, g.Id("Validator").Op(":").Id("wf").Dot(fmt.Sprintf("Validate%s", update)))
								}
								b.Id("opts").Op(":=").Qual(workflowPkg, "UpdateHandlerOptions").Values(updateHandlerOptions...)

								b.If(
									g.Err().Op(":=").Qual(workflowPkg, "SetUpdateHandlerWithOptions").Call(
										g.Id("ctx"), g.Id(fmt.Sprintf("%sUpdateName", update)), g.Id("wf").Dot(update), g.Id("opts"),
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
							})
						}

						// execute workflow
						fn.Return(
							g.Id("wf").Dot("Execute").Call(g.Id("ctx")),
						)
					}),
			),
		)
}

// genWorkerRegisterWorkflow generates a Register<Workflow> public function
func (svc *Service) genWorkerRegisterWorkflow(f *g.File, workflow string) {
	method := svc.methods[workflow]
	builderName := fmt.Sprintf("build%s", method.GoName)
	varName := toCamel("%sFunction", workflow)

	// generate Register<Workflow> function
	f.Commentf("Register%sWorkflow registers a %s workflow with the given worker", workflow, method.Desc.FullName())
	f.Func().
		Id(fmt.Sprintf("Register%sWorkflow", workflow)).
		Params(
			g.Id("r").Qual(workerPkg, "WorkflowRegistry"),
			g.Id("wf").
				Func().
				ParamsFunc(func(args *g.Group) {
					args.Qual(workflowPkg, "Context")
					if svc.cfg.DisableWorkflowInputRename {
						args.Op("*").Id(toCamel("%sInput", workflow))
					} else {
						args.Op("*").Id(toCamel("%sWorkflowInput", workflow))
					}
				}).
				Params(
					g.Id(toCamel("%sWorkflow", workflow)),
					g.Error(),
				),
		).
		Block(
			g.Id(varName).Op("=").Id(builderName).Call(g.Id("wf")),
			g.Id("r").Dot("RegisterWorkflowWithOptions").Call(
				g.Id(varName),
				g.Qual(workflowPkg, "RegisterOptions").Values(
					g.Id("Name").Op(":").Id(toCamel("%sWorkflowName", workflow)),
				),
			),
		)
}

// genWorkerRegisterWorkflows generates a public RegisterWorkflows method for a given service
func (svc *Service) genWorkerRegisterWorkflows(f *g.File) {
	f.Commentf("Register%sWorkflows registers %s workflows with the given worker", svc.Service.GoName, svc.Service.Desc.FullName())
	f.Func().
		Id(toCamel("Register%sWorkflows", svc.Service.GoName)).
		Params(
			g.Id("r").Qual(workerPkg, "WorkflowRegistry"),
			g.Id("workflows").Id(toCamel("%sWorkflows", svc.Service.GoName)),
		).
		BlockFunc(func(fn *g.Group) {
			for _, workflow := range svc.workflowsOrdered {
				fn.Id(toCamel("Register%sWorkflow", workflow)).Call(
					g.Id("r"), g.Id("workflows").Dot(workflow),
				)
			}
		})
}

// genWorkerSignal generates a worker signal struct
func (svc *Service) genWorkerSignal(f *g.File, signal string) {
	typeName := toCamel("%sSignal", signal)

	f.Commentf("%s describes a(n) %s signal", typeName, svc.methods[signal].Desc.FullName())
	f.Type().Id(typeName).Struct(
		g.Id("Channel").Qual(workflowPkg, "ReceiveChannel"),
	)
}

// genWorkerSignalExternal generates a <Signal>External public function
func (svc *Service) genWorkerSignalExternal(f *g.File, signal string) {
	functionName := toCamel("%sExternal", signal)
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)

	f.Commentf("%s sends a(n) %s signal to an existing workflow", functionName, method.Desc.FullName())
	f.Func().Id(functionName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Id(svc.getMessageName(method.Input))
			}
		}).
		Params(g.Error()).
		Block(
			g.Return(
				g.Id(toCamel("%sAsync", functionName)).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("workflowID")
					args.Id("runID")
					if hasInput {
						args.Id("req")
					}
				}).Dot("Get").Call(g.Id("ctx"), g.Nil()),
			),
		)
}

// genWorkerSignalExternalAsync generates a <Signal>ExternalAsync public function
func (svc *Service) genWorkerSignalExternalAsync(f *g.File, signal string) {
	functionName := toCamel("%sExternalAsync", signal)
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)

	f.Commentf("%s sends a(n) %s signal to an existing workflow", functionName, method.Desc.FullName())
	f.Func().Id(functionName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Id(svc.getMessageName(method.Input))
			}
		}).
		Params(g.Qual(workflowPkg, "Future")).
		Block(
			g.Return(
				g.Qual(workflowPkg, "SignalExternalWorkflow").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("workflowID")
					args.Id("runID")
					args.Id(toCamel("%sSignalName", signal))
					if hasInput {
						args.Id("req")
					} else {
						args.Nil()
					}
				}),
			),
		)
}

// genWorkerSignalReceive generates a worker signal Receive method
func (svc *Service) genWorkerSignalReceive(f *g.File, signal string) {
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)

	f.Commentf("Receive blocks until a(n) %s signal is received", method.Desc.FullName())
	f.Func().
		Params(g.Id("s").Op("*").Id(toCamel("%sSignal", signal))).
		Id("Receive").
		Params(g.Id("ctx").Qual(workflowPkg, "Context")).
		ParamsFunc(func(returnVals *g.Group) {
			if hasInput {
				returnVals.Op("*").Id(svc.getMessageName(method.Input))
			}
			returnVals.Bool()
		}).
		BlockFunc(func(b *g.Group) {
			if hasInput {
				b.Var().Id("resp").Id(svc.getMessageName(method.Input))
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

	f.Commentf("ReceiveAsync checks for a %s signal without blocking", method.Desc.FullName())
	f.Func().
		Params(g.Id("s").Op("*").Id(toCamel("%sSignal", signal))).
		Id("ReceiveAsync").
		Params().
		ParamsFunc(func(returnVals *g.Group) {
			if hasInput {
				returnVals.Op("*").Id(svc.getMessageName(method.Input))
			} else {
				returnVals.Bool()
			}
		}).
		BlockFunc(func(b *g.Group) {
			if hasInput {
				b.Var().Id("resp").Id(svc.getMessageName(method.Input))
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

// genWorkerSignalReceiveWithTimeout generates a worker signal ReceiveWithTimeout method
func (svc *Service) genWorkerSignalReceiveWithTimeout(f *g.File, signal string) {
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)

	f.Commentf("ReceiveWithTimeout blocks until a(n) %s signal is received or timeout expires.", method.Desc.FullName())
	f.Comment("Returns more value of false when Channel is closed.")
	f.Comment("Returns ok value of false when no value was found in the channel for the duration of timeout or the ctx was canceled.")
	if hasInput {
		f.Comment("resp will be nil if ok is false.")
	}
	f.Func().
		Params(g.Id("s").Op("*").Id(toCamel("%sSignal", signal))).
		Id("ReceiveWithTimeout").
		Params(
			g.Id("ctx").Qual(workflowPkg, "Context"),
			g.Id("timeout").Qual("time", "Duration"),
		).
		ParamsFunc(func(returnVals *g.Group) {
			if hasInput {
				returnVals.Id("resp").Op("*").Id(svc.getMessageName(method.Input))
			}
			returnVals.Id("ok").Bool()
			returnVals.Id("more").Bool()
		}).
		BlockFunc(func(b *g.Group) {
			if hasInput {
				b.Id("resp").Op("=").Op("&").Id(svc.getMessageName(method.Input)).Values()
			}
			b.If(
				b.List(g.Id("ok"), g.Id("more")).Op("=").Id("s").Dot("Channel").Dot("ReceiveWithTimeout").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("timeout")
					if hasInput {
						args.Op("&").Id("resp")
					} else {
						args.Nil()
					}
				}),
				g.Op("!").Id("ok"),
			).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasInput {
						returnVals.Nil()
					}
					returnVals.False()
					returnVals.Id("more")
				}),
			)
			b.Return()
		})
}

// genWorkerSignalSelect generates a worker signal Select method
func (svc *Service) genWorkerSignalSelect(f *g.File, signal string) {
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)

	f.Commentf("Select checks for a(n) %s signal without blocking", method.Desc.FullName())
	f.Func().
		Params(g.Id("s").Op("*").Id(toCamel("%sSignal", signal))).
		Id("Select").
		Params(
			g.Id("sel").Qual(workflowPkg, "Selector"),
			g.Id("fn").Func().ParamsFunc(func(args *g.Group) {
				if hasInput {
					args.Op("*").Id(svc.getMessageName(method.Input))
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

// genWorkerWorkflowChild generates a public <Workflow>Child function
func (svc *Service) genWorkerWorkflowChild(f *g.File, workflow string) {
	functionName := toCamel("%sChild", workflow)
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s executes a child %s workflow", functionName, svc.fqnForWorkflow(workflow))
	f.Func().
		Id(functionName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				args.Id("req").Op("*").Id(svc.getMessageName(method.Input))
			}
			args.Id("options").Op("...").Op("*").Id(toCamel("%sChildOptions", workflow))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			fn.List(g.Id("childRun"), g.Err()).Op(":=").Id(toCamel("%sChildAsync", workflow)).CallFunc(func(args *g.Group) {
				args.Id("ctx")
				if hasInput {
					args.Id("req")
				}
				args.Id("options").Op("...")
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)
			fn.Return(
				g.Id("childRun").Dot("Get").Call(g.Id("ctx")),
			)
		})
}

// genWorkerWorkflowChildAsync generates a public <Workflow>Child function
func (svc *Service) genWorkerWorkflowChildAsync(f *g.File, workflow string) {
	functionName := toCamel("%sChildAsync", workflow)
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)

	f.Commentf("%s executes a child %s workflow", functionName, svc.fqnForWorkflow(workflow))
	f.Func().
		Id(functionName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				args.Id("req").Op("*").Id(svc.getMessageName(method.Input))
			}
			args.Id("options").Op("...").Op("*").Id(toCamel("%sChildOptions", workflow))
		}).
		Params(
			g.Op("*").Id(toCamel("%sChildRun", workflow)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			// initialize child workflow options with default values
			svc.genClientStartWorkflowOptions(fn, workflow, true)

			fn.Id("ctx").Op("=").Qual(workflowPkg, "WithChildOptions").Call(g.Id("ctx"), g.Op("*").Id("opts"))
			fn.Return(
				g.Op("&").Id(toCamel("%sChildRun", workflow)).Values(
					g.Id("Future").Op(":").Qual(workflowPkg, "ExecuteChildWorkflow").CallFunc(func(args *g.Group) {
						args.Id("ctx")
						args.Id(fmt.Sprintf("%sWorkflowName", workflow))
						if hasInput {
							args.Id("req")
						} else {
							args.Nil()
						}
					}),
				),
				g.Nil(),
			)
		})
}

// genWorkerWorkflowChildOptions generates a <Workflow>ChildOptions struct
func (svc *Service) genWorkerWorkflowChildOptions(f *g.File, workflow string) {
	typeName := toCamel("%sChildOptions", workflow)
	constructorName := "New" + typeName

	f.Commentf("%s provides configuration for a %s workflow operation", typeName, svc.fqnForWorkflow(workflow))
	f.Type().Id(typeName).Struct(
		g.Id("opts").Op("*").Qual(workflowPkg, "ChildWorkflowOptions"),
	)

	f.Commentf("%s initializes a new %s value", constructorName, typeName)
	f.Func().Id(constructorName).Params().Op("*").Id(typeName).Block(
		g.Return(g.Op("&").Id(typeName).Values()),
	)

	f.Comment("WithChildWorkflowOptions sets the initial client.StartWorkflowOptions")
	f.Func().
		Params(g.Id("opts").Op("*").Id(typeName)).
		Id("WithChildWorkflowOptions").
		Params(g.Id("options").Qual(workflowPkg, "ChildWorkflowOptions")).
		Op("*").Id(typeName).
		Block(
			g.Id("opts").Dot("opts").Op("=").Op("&").Id("options"),
			g.Return(g.Id("opts")),
		)
}

// genWorkerWorkflowChildRun generates a <Workflow>ChildRun struct
func (svc *Service) genWorkerWorkflowChildRun(f *g.File, workflow string) {
	typeName := toCamel("%sChildRun", workflow)

	f.Commentf("%s describes a child %s workflow run", typeName, svc.methods[workflow].Desc.FullName())
	f.Type().Id(typeName).StructFunc(func(fields *g.Group) {
		fields.Add(g.Id("Future").Qual(workflowPkg, "ChildWorkflowFuture"))
	})
}

// genWorkerWorkflowChildRunGet generates a <Workflow>ChildRun Get method
func (svc *Service) genWorkerWorkflowChildRunGet(f *g.File, workflow string) {
	typeName := toCamel("%sChildRun", workflow)
	method := svc.methods[workflow]
	hasOutput := !isEmpty(method.Output)

	f.Comment("Get blocks until the workflow is completed, returning the response value")
	f.Func().
		Params(
			g.Id("r").Op("*").Id(typeName),
		).
		Id("Get").
		Params(
			g.Id("ctx").Qual(workflowPkg, "Context"),
		).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			if hasOutput {
				fn.Var().Id("resp").Id(svc.getMessageName(method.Output))
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

// genWorkerWorkflowChildRunSelect generates a <Workflow>ChildRun Select method
func (svc *Service) genWorkerWorkflowChildRunSelect(f *g.File, workflow string) {
	typeName := toCamel("%sChildRun", workflow)

	f.Comment("Select adds this completion to the selector. Callback can be nil.")
	f.Func().
		Params(
			g.Id("r").Op("*").Id(typeName),
		).
		Id("Select").
		Params(
			g.Id("sel").Qual(workflowPkg, "Selector"),
			g.Id("fn").Func().Params(g.Op("*").Id(typeName)),
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
							g.Id("fn").Call(g.Id("r")),
						),
					),
				),
			),
		)
}

// genWorkerWorkflowChildRunSelectStart generates a <Workflow>ChildRun SelectStart method
func (svc *Service) genWorkerWorkflowChildRunSelectStart(f *g.File, workflow string) {
	typeName := toCamel("%sChildRun", workflow)

	f.Comment("SelectStart adds waiting for start to the selector. Callback can be nil.")
	f.Func().
		Params(
			g.Id("r").Op("*").Id(typeName),
		).
		Id("SelectStart").
		Params(
			g.Id("sel").Qual(workflowPkg, "Selector"),
			g.Id("fn").Func().Params(g.Op("*").Id(typeName)),
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
							g.Id("fn").Call(g.Id("r")),
						),
					),
				),
			),
		)
}

// genWorkerWorkflowChildRunSignals generates <Workflow>ChildRun signal methods
func (svc *Service) genWorkerWorkflowChildRunSignals(f *g.File, workflow string) {
	typeName := toCamel("%sChildRun", workflow)
	opts := svc.workflows[workflow]

	for _, signalOpts := range opts.GetSignal() {
		signal := signalOpts.GetRef()
		handler := svc.methods[signal]
		hasInput := !isEmpty(handler.Input)
		asyncName := toCamel("%sAsync", signal)

		f.Commentf("%s sends a(n) %q signal request to the child workflow", signal, svc.fqnForSignal(signal))
		f.Func().
			Params(g.Id("r").Op("*").Id(typeName)).
			Id(signal).
			ParamsFunc(func(params *g.Group) {
				params.Id("ctx").Qual(workflowPkg, "Context")
				if hasInput {
					params.Id("input").Op("*").Id(svc.getMessageName(handler.Input))
				}
			}).
			Params(g.Error()).
			Block(
				g.Return(g.Id("r").Dot(asyncName).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					if hasInput {
						args.Id("input")
					}
				})).Dot("Get").Call(g.Id("ctx"), g.Nil()),
			)

		f.Commentf("%s sends a(n) %q signal request to the child workflow", asyncName, svc.fqnForSignal(signal))
		f.Func().
			Params(g.Id("r").Op("*").Id(typeName)).
			Id(asyncName).
			ParamsFunc(func(params *g.Group) {
				params.Id("ctx").Qual(workflowPkg, "Context")
				if hasInput {
					params.Id("input").Op("*").Id(svc.getMessageName(handler.Input))
				}
			}).
			Params(g.Qual(workflowPkg, "Future")).
			Block(
				g.Return(g.Id("r").Dot("Future").Dot("SignalChildWorkflow").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id(toCamel("%sSignalName", signal))
					if hasInput {
						args.Id("input")
					} else {
						args.Nil()
					}
				})),
			)
	}
}

// genWorkerWorkflowChildRunWaitStart generates a <Workflow>ChildRun WaitStart method
func (svc *Service) genWorkerWorkflowChildRunWaitStart(f *g.File, workflow string) {
	typeName := toCamel("%sChildRun", workflow)

	f.Comment("WaitStart waits for the child workflow to start")
	f.Func().
		Params(
			g.Id("r").Op("*").Id(typeName),
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

// genWorkerWorkflowFunctionVars generates a <Workflow>Function var for each workflow that are
// initialized on registration
func (svc *Service) genWorkerWorkflowFunctionVars(f *g.File) {
	f.Commentf("Reference to generated workflow functions")
	f.Var().DefsFunc(func(defs *g.Group) {
		for _, workflow := range svc.workflowsOrdered {
			method := svc.methods[workflow]
			hasInput := !isEmpty(method.Input)
			hasOutput := !isEmpty(method.Output)
			varName := toCamel("%sFunction", workflow)

			if method.Comments.Leading.String() != "" {
				defs.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
			} else {
				defs.Commentf("%s implements a %q workflow", varName, toCamel("%sWorkflow", workflow))
			}
			defs.Id(varName).
				Func().
				ParamsFunc(func(args *g.Group) {
					args.Qual(workflowPkg, "Context")
					if hasInput {
						args.Op("*").Id(svc.getMessageName(method.Input))
					}
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Id(svc.getMessageName(method.Output))
					}
					returnVals.Error()
				})
		}
	})
}

// genWorkerWorkflowInput generates a <Workflow>Input struct
func (svc *Service) genWorkerWorkflowInput(f *g.File, workflow string) {
	typeName := toCamel("%sWorkflowInput", workflow)
	if svc.cfg.DisableWorkflowInputRename {
		typeName = toCamel("%sInput", workflow)
	}
	opts := svc.workflows[workflow]
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)

	f.Commentf("%s describes the input to a(n) %s workflow constructor", typeName, method.Desc.FullName())
	f.Type().Id(typeName).StructFunc(func(fields *g.Group) {
		if hasInput {
			fields.Id("Req").Op("*").Id(svc.getMessageName(method.Input))
		}

		// add workflow signals
		for _, signalOpts := range opts.GetSignal() {
			signal := signalOpts.GetRef()
			fields.Id(signal).Op("*").Id(fmt.Sprintf("%sSignal", signal))
		}
	})
}

// genWorkerWorkflowInterface generates a <Workflow> interface
func (svc *Service) genWorkerWorkflowInterface(f *g.File, workflow string) {
	typeName := toCamel("%sWorkflow", workflow)
	opts := svc.workflows[workflow]
	method := svc.methods[workflow]
	hasOutput := !isEmpty(method.Output)

	// generate workflow interface
	if method.Comments.Leading.String() != "" {
		f.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
	} else {
		f.Commentf("%s describes a(n) %s workflow implementation", typeName, method.Desc.FullName())
	}
	f.Type().Id(typeName).InterfaceFunc(func(methods *g.Group) {
		if method.Comments.Leading.String() != "" {
			methods.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
		} else {
			methods.Commentf("Execute defines the entrypoint to a(n) %s workflow", method.Desc.FullName())
		}
		methods.
			Id("Execute").
			Params(
				g.Id("ctx").Qual(workflowPkg, "Context"),
			).
			ParamsFunc(func(returnVals *g.Group) {
				if hasOutput {
					returnVals.Op("*").Id(svc.getMessageName(method.Output))
				}
				returnVals.Error()
			})

		// add workflow query methods
		for _, queryOpts := range opts.GetQuery() {
			query := queryOpts.GetRef()
			handler := svc.methods[query]
			hasInput := !isEmpty(handler.Input)

			if handler.Comments.Leading.String() != "" {
				methods.Comment(strings.TrimSuffix(handler.Comments.Leading.String(), "\n"))
			} else {
				methods.Commentf("%s implements a(n) %s query handler", query, handler.Desc.FullName())
			}
			methods.Id(query).
				ParamsFunc(func(args *g.Group) {
					if hasInput {
						args.Op("*").Id(svc.getMessageName(handler.Input))
					}
				}).
				Params(
					g.Op("*").Id(svc.getMessageName(handler.Output)),
					g.Error(),
				)
		}

		// add workflow update methods
		for _, updateOpts := range opts.GetUpdate() {
			update := updateOpts.GetRef()
			handler := svc.methods[update]
			handlerOpts := svc.updates[update]
			hasInput := !isEmpty(handler.Input)
			hasOutput := !isEmpty(handler.Output)

			// add Validate<Update> method if enabled
			if handlerOpts.GetValidate() {
				validatorName := toCamel("Validate%s", update)
				methods.Commentf("%s validates a(n) %s update", validatorName, handler.Desc.FullName())
				methods.Id(validatorName).
					ParamsFunc(func(args *g.Group) {
						args.Qual(workflowPkg, "Context")
						if hasInput {
							args.Op("*").Id(svc.getMessageName(handler.Input))
						}
					}).
					Params(g.Error())
			}

			// add <Update> method
			if handler.Comments.Leading.String() != "" {
				methods.Comment(strings.TrimSuffix(handler.Comments.Leading.String(), "\n"))
			} else {
				methods.Commentf("%s implements a(n) %s update handler", update, handler.Desc.FullName())
			}
			methods.Id(update).
				ParamsFunc(func(args *g.Group) {
					args.Qual(workflowPkg, "Context")
					if hasInput {
						args.Op("*").Id(svc.getMessageName(handler.Input))
					}
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Id(svc.getMessageName(handler.Output))
					}
					returnVals.Error()
				})
		}
	})
}

// genWorkerWorkflowsInterface generates a Workflows interface for a given service
func (svc *Service) genWorkerWorkflowsInterface(f *g.File) {
	typeName := toCamel("%sWorkflows", svc.Service.GoName)
	// generate workflows interface
	f.Commentf("%s provides methods for initializing new %s workflow values", typeName, svc.Service.Desc.FullName())
	f.Type().Id(typeName).InterfaceFunc(func(methods *g.Group) {
		for _, workflow := range svc.workflowsOrdered {
			handler := svc.methods[workflow]

			if handler.Comments.Leading.String() != "" {
				f.Comment(strings.TrimSuffix(handler.Comments.Leading.String(), "\n"))
			} else {
				f.Commentf("%s initializes a new a(n) %s implementation", workflow, toCamel("%sWorkflow", workflow))
			}
			methods.
				Id(workflow).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual(workflowPkg, "Context")
					if svc.cfg.DisableWorkflowInputRename {
						args.Id("input").Op("*").Id(toCamel("%sInput", workflow))
					} else {
						args.Id("input").Op("*").Id(toCamel("%sWorkflowInput", workflow))
					}
				}).
				Params(
					g.Id(toCamel("%sWorkflow", workflow)),
					g.Error(),
				)
		}
	})
}
