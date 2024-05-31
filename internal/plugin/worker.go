package plugin

import (
	"fmt"
	"sort"
	"strings"

	g "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// genWorkerBuilderFunction generates a build<Workflow> function that converts
// a constructor function or method into a valid workflow function
func (svc *Manifest) genWorkerBuilderFunction(f *g.File, workflow protoreflect.FullName) {
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
						args.Op("*").Id(svc.toCamel("%sInput", workflow))
					} else {
						args.Op("*").Id(svc.toCamel("%sWorkflowInput", workflow))
					}
				}).
				Params(
					g.Id(svc.toCamel("%sWorkflow", workflow)),
					g.Error(),
				),
		).
		Params(
			g.Func().
				ParamsFunc(func(args *g.Group) {
					args.Qual(workflowPkg, "Context")
					if hasInput {
						args.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
					}
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
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
							args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
						}
					}).
					ParamsFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
						}
						returnVals.Error()
					}).
					BlockFunc(func(fn *g.Group) {
						// build input struct
						inputType := svc.toCamel("%sWorkflowInput", workflow)
						if svc.cfg.DisableWorkflowInputRename {
							inputType = svc.toCamel("%sInput", workflow)
						}
						fn.Id("input").Op(":=").Op("&").Id(inputType).BlockFunc(func(fields *g.Group) {
							if hasInput {
								fields.Id("Req").Op(":").Id("req").Op(",")
							}
							for _, s := range opts.GetSignal() {
								signal := getFullyQualifiedRef(workflow, s.GetRef())
								fields.Id(svc.methods[signal].GoName).Op(":").Op("&").Add(svc.Qual(signal, svc.toCamel("%sSignal", signal))).Block(
									g.Id("Channel").Op(":").Qual(workflowPkg, "GetSignalChannel").Call(
										g.Id("ctx"), svc.Qual(signal, svc.toCamel("%sSignalName", signal)),
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
									g.Id("ctx"), g.Id(svc.toCamel("%sQueryName", query)), g.Id("wf").Dot(query),
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
							update := getFullyQualifiedRef(workflow, u.GetRef())
							updateOpts := svc.updates[update]

							fn.BlockFunc(func(b *g.Group) {
								// build UpdateHandlerOptions
								var updateHandlerOptions []g.Code
								if updateOpts.GetValidate() {
									updateHandlerOptions = append(updateHandlerOptions, g.Id("Validator").Op(":").Id("wf").Dot(svc.toCamel("Validate%s", update)))
								}
								b.Id("opts").Op(":=").Qual(workflowPkg, "UpdateHandlerOptions").Values(updateHandlerOptions...)

								b.If(
									g.Err().Op(":=").Qual(workflowPkg, "SetUpdateHandlerWithOptions").Call(
										g.Id("ctx"), g.Id(svc.toCamel("%sUpdateName", update)), g.Id("wf").Dot(svc.methods[update].GoName), g.Id("opts"),
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
func (svc *Manifest) genWorkerRegisterWorkflow(f *g.File, workflow protoreflect.FullName) {
	method := svc.methods[workflow]
	opts := svc.workflows[workflow]
	builderName := fmt.Sprintf("build%s", method.GoName)
	varName := svc.toCamel("%sFunction", workflow)

	// generate Register<Workflow> function
	f.Commentf("Register%sWorkflow registers a %s workflow with the given worker", svc.methods[workflow].GoName, method.Desc.FullName())
	f.Func().
		Id(svc.toCamel("Register%sWorkflow", workflow)).
		Params(
			g.Id("r").Qual(workerPkg, "WorkflowRegistry"),
			g.Id("wf").
				Func().
				ParamsFunc(func(args *g.Group) {
					args.Qual(workflowPkg, "Context")
					if svc.cfg.DisableWorkflowInputRename {
						args.Op("*").Id(svc.toCamel("%sInput", workflow))
					} else {
						args.Op("*").Id(svc.toCamel("%sWorkflowInput", workflow))
					}
				}).
				Params(
					g.Id(svc.toCamel("%sWorkflow", workflow)),
					g.Error(),
				),
		).
		BlockFunc(func(fn *g.Group) {
			fn.Id(varName).Op("=").Id(builderName).Call(g.Id("wf"))
			names := []g.Code{g.Id(svc.toCamel("%sWorkflowName", workflow))}
			aliases := opts.GetAliases()
			sort.Strings(aliases)
			for _, alias := range aliases {
				names = append(names, g.Lit(alias))
			}
			for _, name := range names {
				fn.Id("r").Dot("RegisterWorkflowWithOptions").Call(
					g.Id(varName),
					g.Qual(workflowPkg, "RegisterOptions").Values(
						g.Id("Name").Op(":").Add(name),
					),
				)
			}
		})
}

// genWorkerRegisterWorkflows generates a public RegisterWorkflows method for a given service
func (svc *Manifest) genWorkerRegisterWorkflows(f *g.File) {
	f.Commentf("Register%sWorkflows registers %s workflows with the given worker", svc.Service.GoName, svc.Service.Desc.FullName())
	f.Func().
		Id(svc.toCamel("Register%sWorkflows", svc.Service.GoName)).
		Params(
			g.Id("r").Qual(workerPkg, "WorkflowRegistry"),
			g.Id("workflows").Id(svc.toCamel("%sWorkflows", svc.Service.GoName)),
		).
		BlockFunc(func(fn *g.Group) {
			for _, workflow := range svc.workflowsOrdered {
				if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
					continue
				}
				fn.Id(svc.toCamel("Register%sWorkflow", workflow)).Call(
					g.Id("r"), g.Id("workflows").Dot(svc.methods[workflow].GoName),
				)
			}
		})
}

// genWorkerSignal generates a worker signal struct
func (svc *Manifest) genWorkerSignal(f *g.File, signal protoreflect.FullName) {
	typeName := svc.toCamel("%sSignal", signal)

	f.Commentf("%s describes a(n) %s signal", typeName, svc.methods[signal].Desc.FullName())
	f.Type().Id(typeName).Struct(
		g.Id("Channel").Qual(workflowPkg, "ReceiveChannel"),
	)
}

func (svc *Manifest) genWorkerSignalConstructor(f *g.File, signal protoreflect.FullName) {
	typeName := svc.toCamel("%sSignal", signal)
	funcName := svc.toCamel("New%sSignal", signal)

	f.Commentf("%s initializes a new %s signal wrapper", funcName, svc.fqnForSignal(signal))
	f.Func().
		Id(funcName).
		Params(g.Id("ctx").Qual(workflowPkg, "Context")).
		Op("*").Id(typeName).
		Block(
			g.Return(
				g.Op("&").Id(typeName).ValuesFunc(func(fields *g.Group) {
					fields.Id("Channel").Op(":").Qual(workflowPkg, "GetSignalChannel").Call(g.Id("ctx"), g.Id(svc.toCamel("%sSignalName", signal)))
				}),
			),
		)
}

// genWorkerSignalExternal generates a <Signal>External public function
func (svc *Manifest) genWorkerSignalExternal(f *g.File, signal protoreflect.FullName) {
	functionName := svc.toCamel("%sExternal", signal)
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)

	commentWithDefaultf(f, methodSet(method), "%s sends a(n) %s signal to an existing workflow", functionName, svc.fqnForSignal(signal))
	f.Func().Id(functionName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
		}).
		Params(g.Error()).
		Block(
			g.Return(
				g.Id(svc.toCamel("%sAsync", functionName)).CallFunc(func(args *g.Group) {
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
func (svc *Manifest) genWorkerSignalExternalAsync(f *g.File, signal protoreflect.FullName) {
	functionName := svc.toCamel("%sExternalAsync", signal)
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)

	commentWithDefaultf(f, methodSet(method), "%s sends a(n) %s signal to an existing workflow", functionName, svc.fqnForSignal(signal))
	f.Func().Id(functionName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
		}).
		Params(g.Qual(workflowPkg, "Future")).
		Block(
			g.Return(
				g.Qual(workflowPkg, "SignalExternalWorkflow").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("workflowID")
					args.Id("runID")
					args.Id(svc.toCamel("%sSignalName", signal))
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
func (svc *Manifest) genWorkerSignalReceive(f *g.File, signal protoreflect.FullName) {
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)

	f.Commentf("Receive blocks until a(n) %s signal is received", method.Desc.FullName())
	f.Func().
		Params(g.Id("s").Op("*").Id(svc.toCamel("%sSignal", signal))).
		Id("Receive").
		Params(g.Id("ctx").Qual(workflowPkg, "Context")).
		ParamsFunc(func(returnVals *g.Group) {
			if hasInput {
				returnVals.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
			returnVals.Bool()
		}).
		BlockFunc(func(b *g.Group) {
			if hasInput {
				b.Var().Id("resp").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
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
func (svc *Manifest) genWorkerSignalReceiveAsync(f *g.File, signal protoreflect.FullName) {
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)

	f.Commentf("ReceiveAsync checks for a %s signal without blocking", method.Desc.FullName())
	f.Func().
		Params(g.Id("s").Op("*").Id(svc.toCamel("%sSignal", signal))).
		Id("ReceiveAsync").
		Params().
		ParamsFunc(func(returnVals *g.Group) {
			if hasInput {
				returnVals.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			} else {
				returnVals.Bool()
			}
		}).
		BlockFunc(func(b *g.Group) {
			if hasInput {
				b.Var().Id("resp").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
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
func (svc *Manifest) genWorkerSignalReceiveWithTimeout(f *g.File, signal protoreflect.FullName) {
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)

	f.Commentf("ReceiveWithTimeout blocks until a(n) %s signal is received or timeout expires.", method.Desc.FullName())
	f.Comment("Returns more value of false when Channel is closed.")
	f.Comment("Returns ok value of false when no value was found in the channel for the duration of timeout or the ctx was canceled.")
	if hasInput {
		f.Comment("resp will be nil if ok is false.")
	}
	f.Func().
		Params(g.Id("s").Op("*").Id(svc.toCamel("%sSignal", signal))).
		Id("ReceiveWithTimeout").
		Params(
			g.Id("ctx").Qual(workflowPkg, "Context"),
			g.Id("timeout").Qual("time", "Duration"),
		).
		ParamsFunc(func(returnVals *g.Group) {
			if hasInput {
				returnVals.Id("resp").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
			returnVals.Id("ok").Bool()
			returnVals.Id("more").Bool()
		}).
		BlockFunc(func(b *g.Group) {
			if hasInput {
				b.Id("resp").Op("=").Op("&").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input)).Values()
			}
			b.If(
				g.List(g.Id("ok"), g.Id("more")).Op("=").Id("s").Dot("Channel").Dot("ReceiveWithTimeout").CallFunc(func(args *g.Group) {
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
func (svc *Manifest) genWorkerSignalSelect(f *g.File, signal protoreflect.FullName) {
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)

	f.Commentf("Select checks for a(n) %s signal without blocking", method.Desc.FullName())
	f.Func().
		Params(g.Id("s").Op("*").Id(svc.toCamel("%sSignal", signal))).
		Id("Select").
		Params(
			g.Id("sel").Qual(workflowPkg, "Selector"),
			g.Id("fn").Func().ParamsFunc(func(args *g.Group) {
				if hasInput {
					args.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
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
func (svc *Manifest) genWorkerWorkflowChild(f *g.File, workflow protoreflect.FullName) {
	functionName := svc.toCamel("%sChild", workflow)
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	commentWithDefaultf(f, methodSet(method), "%s executes a child %s workflow and blocks until error or response received", functionName, svc.fqnForWorkflow(workflow))
	f.Func().
		Id(functionName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
			args.Id("options").Op("...").Op("*").Id(svc.toCamel("%sChildOptions", workflow))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			fn.List(g.Id("childRun"), g.Err()).Op(":=").Id(svc.toCamel("%sChildAsync", workflow)).CallFunc(func(args *g.Group) {
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
func (svc *Manifest) genWorkerWorkflowChildAsync(f *g.File, workflow protoreflect.FullName) {
	functionName := svc.toCamel("%sChildAsync", workflow)
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)

	commentWithDefaultf(f, methodSet(method), "%s starts a child %s workflow and returns a handle to the child workflow run", functionName, svc.fqnForWorkflow(workflow))
	f.Func().
		Id(functionName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
			args.Id("options").Op("...").Op("*").Id(svc.toCamel("%sChildOptions", workflow))
		}).
		Params(
			g.Op("*").Id(svc.toCamel("%sChildRun", workflow)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			// initialize options
			fn.Var().Id("o").Op("*").Id(svc.toCamel("%sChildOptions", workflow))
			fn.If(g.Len(g.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(g.Lit(0)).Op("!=").Nil()).Block(
				g.Id("o").Op("=").Id("options").Index(g.Lit(0)),
			).Else().Block(
				g.Id("o").Op("=").Id(svc.toCamel("New%sChildOptions", workflow)).Call(),
			)

			// initialize client.StartWorkfowOptions
			fn.List(g.Id("opts"), g.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(args *g.Group) {
				args.Id("ctx")
				if hasInput {
					args.Id("req").Dot("ProtoReflect").Call()
				} else {
					args.Nil()
				}
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error initializing workflow.ChildWorkflowOptions: %w"), g.Err())),
			)

			fn.Id("ctx").Op("=").Qual(workflowPkg, "WithChildOptions").Call(g.Id("ctx"), g.Id("opts"))
			fn.Return(
				g.Op("&").Id(svc.toCamel("%sChildRun", workflow)).Values(
					g.Id("Future").Op(":").Qual(workflowPkg, "ExecuteChildWorkflow").CallFunc(func(args *g.Group) {
						args.Id("ctx")
						args.Id(svc.toCamel("%sWorkflowName", workflow))
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

// genWorkerWorkflowChildRun generates a <Workflow>ChildRun struct
func (svc *Manifest) genWorkerWorkflowChildRun(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toCamel("%sChildRun", workflow)

	f.Commentf("%s describes a child %s workflow run", typeName, svc.methods[workflow].GoName)
	f.Type().Id(typeName).StructFunc(func(fields *g.Group) {
		fields.Add(g.Id("Future").Qual(workflowPkg, "ChildWorkflowFuture"))
	})
}

// genWorkerWorkflowChildRunGet generates a <Workflow>ChildRun Get method
func (svc *Manifest) genWorkerWorkflowChildRunGet(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toCamel("%sChildRun", workflow)
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
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			if hasOutput {
				fn.Var().Id("resp").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
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
func (svc *Manifest) genWorkerWorkflowChildRunSelect(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toCamel("%sChildRun", workflow)

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
func (svc *Manifest) genWorkerWorkflowChildRunSelectStart(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toCamel("%sChildRun", workflow)

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
func (svc *Manifest) genWorkerWorkflowChildRunSignals(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toCamel("%sChildRun", workflow)
	opts := svc.workflows[workflow]

	for _, signalOpts := range opts.GetSignal() {
		signal := getFullyQualifiedRef(workflow, signalOpts.GetRef())
		handler := svc.methods[signal]
		hasInput := !isEmpty(handler.Input)
		asyncName := svc.toCamel("%sAsync", signal)

		f.Commentf("%s sends a(n) %q signal request to the child workflow", svc.methods[signal].GoName, svc.fqnForSignal(signal))
		f.Func().
			Params(g.Id("r").Op("*").Id(typeName)).
			Id(svc.methods[signal].GoName).
			ParamsFunc(func(params *g.Group) {
				params.Id("ctx").Qual(workflowPkg, "Context")
				if hasInput {
					params.Id("input").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
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
					params.Id("input").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
				}
			}).
			Params(g.Qual(workflowPkg, "Future")).
			Block(
				g.Return(g.Id("r").Dot("Future").Dot("SignalChildWorkflow").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Add(svc.Qual(signal, svc.toCamel("%sSignalName", signal)))
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
func (svc *Manifest) genWorkerWorkflowChildRunWaitStart(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toCamel("%sChildRun", workflow)

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
func (svc *Manifest) genWorkerWorkflowFunctionVars(f *g.File) {
	if len(svc.workflowsOrdered) == 0 {
		return
	}
	f.Commentf("Reference to generated workflow functions")
	f.Var().DefsFunc(func(defs *g.Group) {
		for _, workflow := range svc.workflowsOrdered {
			if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
				continue
			}
			method := svc.methods[workflow]
			hasInput := !isEmpty(method.Input)
			hasOutput := !isEmpty(method.Output)
			varName := svc.toCamel("%sFunction", workflow)

			commentWithDefaultf(defs, methodSet(method), "%s implements a %q workflow", varName, svc.fqnForWorkflow(workflow))
			defs.Id(varName).
				Func().
				ParamsFunc(func(args *g.Group) {
					args.Qual(workflowPkg, "Context")
					if hasInput {
						args.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
					}
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
					}
					returnVals.Error()
				})
		}
	})

	typeName := svc.toCamel("%sWorkflowFunctions", svc.GoName)
	implName := svc.toLowerCamel("%sWorkflowFunctions", svc.GoName)
	f.Commentf("%s describes a mockable dependency for inlining workflows within other workflows", typeName)
	f.Type().Defs(
		g.Commentf("%s describes a mockable dependency for inlining workflows within other workflows", typeName),
		g.Id(typeName).InterfaceFunc(func(methods *g.Group) {
			for _, workflow := range svc.workflowsOrdered {
				if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
					continue
				}
				methodName := svc.toCamel("%s", workflow)
				method := svc.methods[workflow]
				hasInput := !isEmpty(method.Input)
				hasOutput := !isEmpty(method.Output)
				commentWithDefaultf(methods, methodSet(method), "%s executes a %q workflow inline", methodName, svc.fqnForWorkflow(workflow))
				methods.Id(methodName).
					ParamsFunc(func(args *g.Group) {
						args.Qual(workflowPkg, "Context")
						if hasInput {
							args.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
						}
					}).
					ParamsFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
						}
						returnVals.Error()
					})
			}
		}),

		g.Commentf("%s provides an internal %s implementation", implName, typeName),
		g.Id(implName).Struct(),
	)

	f.Func().Id(svc.toCamel("New%s", typeName)).Params().Id(typeName).Block(
		g.Return(g.Op("&").Id(implName).Values()),
	)

	for _, workflow := range svc.workflowsOrdered {
		if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
			continue
		}
		methodName := svc.toCamel("%s", workflow)
		method := svc.methods[workflow]
		varName := svc.toCamel("%sFunction", workflow)
		hasInput := !isEmpty(method.Input)
		hasOutput := !isEmpty(method.Output)
		commentWithDefaultf(f, methodSet(method), "%s executes a %q workflow inline", methodName, svc.fqnForWorkflow(workflow))
		f.Func().
			Params(g.Id("f").Op("*").Id(implName)).
			Id(methodName).
			ParamsFunc(func(args *g.Group) {
				args.Id("ctx").Qual(workflowPkg, "Context")
				if hasInput {
					args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
				}
			}).
			ParamsFunc(func(returnVals *g.Group) {
				if hasOutput {
					returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
				}
				returnVals.Error()
			}).
			Block(
				g.If(g.Id(varName).Op("==").Nil()).Block(
					g.ReturnFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Qual("errors", "New").Call(g.Lit(fmt.Sprintf("%s requires workflow registration via %s or %s", methodName, svc.toCamel("Register%sWorkflows", svc.GoName), svc.toCamel("Register%sWorkflow", workflow))))
					}),
				),
				g.Return(g.Id(varName).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					if hasInput {
						args.Id("req")
					}
				})),
			)
	}
}

// genWorkerWorkflowInput generates a <Workflow>Input struct
func (svc *Manifest) genWorkerWorkflowInput(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toCamel("%sWorkflowInput", workflow)
	if svc.cfg.DisableWorkflowInputRename {
		typeName = svc.toCamel("%sInput", workflow)
	}
	opts := svc.workflows[workflow]
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)

	f.Commentf("%s describes the input to a(n) %s workflow constructor", typeName, svc.fqnForWorkflow(workflow))
	f.Type().Id(typeName).StructFunc(func(fields *g.Group) {
		if hasInput {
			fields.Id("Req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
		}

		// add workflow signals
		for _, signalOpts := range opts.GetSignal() {
			signal := getFullyQualifiedRef(workflow, signalOpts.GetRef())
			fields.Id(svc.methods[signal].GoName).Op("*").Add(svc.Qual(signal, svc.toCamel("%sSignal", signal)))
		}
	})
}

// genWorkerWorkflowInterface generates a <Workflow> interface
func (svc *Manifest) genWorkerWorkflowInterface(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toCamel("%sWorkflow", workflow)
	opts := svc.workflows[workflow]
	method := svc.methods[workflow]
	hasOutput := !isEmpty(method.Output)

	// generate workflow interface
	var details []string
	if comments := method.Comments.Leading.String(); comments != "" && !strings.Contains(comments, svc.fqnForWorkflow(workflow)) {
		details = append(details, fmt.Sprintf("name: \"%s\"", svc.fqnForWorkflow(workflow)))
	}
	if id := opts.GetId(); id != "" {
		details = append(details, fmt.Sprintf("id: \"%s\"", id))
	}

	commentWithDefaultf(f, methodSet(method), "%s describes a(n) %s workflow implementation", typeName, svc.fqnForWorkflow(workflow))
	if len(details) > 0 {
		f.Comment(" ")
		f.Commentf("workflow details: (%s)", strings.Join(details, ", "))
	}
	f.Type().Id(typeName).InterfaceFunc(func(methods *g.Group) {
		commentf(methods, methodSet(method), "Execute defines the entrypoint to a(n) %s workflow", svc.fqnForWorkflow(workflow))
		methods.
			Id("Execute").
			Params(
				g.Id("ctx").Qual(workflowPkg, "Context"),
			).
			ParamsFunc(func(returnVals *g.Group) {
				if hasOutput {
					returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
				}
				returnVals.Error()
			}).
			Line()

		// add workflow query methods
		for _, queryOpts := range opts.GetQuery() {
			query := getFullyQualifiedRef(workflow, queryOpts.GetRef())
			handler := svc.methods[query]
			hasInput := !isEmpty(handler.Input)

			commentWithDefaultf(methods, methodSet(handler), "%s implements a(n) %s query handler", query, svc.fqnForQuery(query))
			methods.Id(svc.toCamel("%s", query)).
				ParamsFunc(func(args *g.Group) {
					if hasInput {
						args.Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
					}
				}).
				Params(
					g.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), svc.getMessageName(handler.Output)),
					g.Error(),
				).
				Line()
		}

		// add workflow update methods
		for _, updateOpts := range opts.GetUpdate() {
			update := getFullyQualifiedRef(workflow, updateOpts.GetRef())
			handler := svc.methods[update]
			handlerOpts := svc.updates[update]
			hasInput := !isEmpty(handler.Input)
			hasOutput := !isEmpty(handler.Output)

			// add Validate<Update> method if enabled
			if handlerOpts.GetValidate() {
				validatorName := svc.toCamel("Validate%s", update)
				methods.Commentf("%s validates a(n) %s update", validatorName, svc.fqnForUpdate(update))
				methods.Id(validatorName).
					ParamsFunc(func(args *g.Group) {
						args.Qual(workflowPkg, "Context")
						if hasInput {
							args.Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
						}
					}).
					Params(g.Error()).
					Line()
			}

			// add <Update> method
			commentWithDefaultf(methods, methodSet(handler), "%s implements a(n) %s update handler", update, svc.fqnForQuery(update))
			methods.Id(svc.toCamel("%s", update)).
				ParamsFunc(func(args *g.Group) {
					args.Qual(workflowPkg, "Context")
					if hasInput {
						args.Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
					}
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), svc.getMessageName(handler.Output))
					}
					returnVals.Error()
				}).
				Line()
		}
	})
}

// genWorkerWorkflowsInterface generates a Workflows interface for a given service
func (svc *Manifest) genWorkerWorkflowsInterface(f *g.File) {
	typeName := svc.toCamel("%sWorkflows", svc.Service.GoName)
	// generate workflows interface
	f.Commentf("%s provides methods for initializing new %s workflow values", typeName, svc.Service.Desc.FullName())
	f.Type().Id(typeName).InterfaceFunc(func(methods *g.Group) {
		for _, workflow := range svc.workflowsOrdered {
			if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
				continue
			}
			handler := svc.methods[workflow]

			commentWithDefaultf(methods, methodSet(handler), "%s initializes a new a(n) %s implementation", svc.toCamel("%s", workflow), svc.toCamel("%sWorkflow", workflow))
			methods.
				Id(svc.methods[workflow].GoName).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual(workflowPkg, "Context")
					if svc.cfg.DisableWorkflowInputRename {
						args.Id("input").Op("*").Id(svc.toCamel("%sInput", workflow))
					} else {
						args.Id("input").Op("*").Id(svc.toCamel("%sWorkflowInput", workflow))
					}
				}).
				Params(
					g.Id(svc.toCamel("%sWorkflow", workflow)),
					g.Error(),
				).
				Line()
		}
	})
}
