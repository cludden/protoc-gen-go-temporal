package plugin

import (
	"fmt"
	"sort"
	"strings"

	j "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (n *names) workflowIDExpression(workflow protoreflect.FullName) string {
	return n.toCamel("%sIDExpression", workflow)
}

// genWorkerBuilderFunction generates a build<Workflow> function that converts
// a constructor function or method into a valid workflow function
func (m *Manifest) genWorkerBuilderFunction(f *j.File, workflow protoreflect.FullName) {
	method := m.methods[workflow]
	opts := m.workflows[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	builderName := fmt.Sprintf("build%s", method.GoName)

	// generate Build<Workflow> function
	f.Commentf("%s converts a %s workflow struct into a valid workflow function", builderName, method.GoName)
	f.Func().
		Id(builderName).
		Params(
			j.Id("ctor").
				Func().
				ParamsFunc(func(args *j.Group) {
					args.Qual(workflowPkg, "Context")
					if m.cfg.DisableWorkflowInputRename {
						args.Op("*").Id(m.toCamel("%sInput", workflow))
					} else {
						args.Op("*").Id(m.toCamel("%sWorkflowInput", workflow))
					}
				}).
				Params(
					j.Id(m.toCamel("%sWorkflow", workflow)),
					j.Error(),
				),
		).
		Params(
			j.Func().
				ParamsFunc(func(args *j.Group) {
					args.Qual(workflowPkg, "Context")
					if hasInput {
						args.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
					}
				}).
				ParamsFunc(func(returnVals *j.Group) {
					if hasOutput {
						returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
					}
					returnVals.Error()
				}),
		).
		Block(
			j.Return(
				//g.Parens(g.Op("&").Id(workerName).Values(g.Id("wf"))).Dot(method.GoName),
				// generate <Workflow> method for worker struct
				j.Func().
					ParamsFunc(func(args *j.Group) {
						args.Id("ctx").Qual(workflowPkg, "Context")
						if hasInput {
							args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
						}
					}).
					ParamsFunc(func(returnVals *j.Group) {
						if hasOutput {
							returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
						}
						returnVals.Error()
					}).
					BlockFunc(func(fn *j.Group) {
						// build input struct
						inputType := m.toCamel("%sWorkflowInput", workflow)
						if m.cfg.DisableWorkflowInputRename {
							inputType = m.toCamel("%sInput", workflow)
						}
						fn.Id("input").Op(":=").Op("&").Id(inputType).BlockFunc(func(fields *j.Group) {
							if hasInput {
								fields.Id("Req").Op(":").Id("req").Op(",")
							}
							for _, s := range opts.GetSignal() {
								signal := getFullyQualifiedRef(workflow, s.GetRef())
								fields.Id(m.methods[signal].GoName).Op(":").Op("&").Add(m.Qual(signal, m.toCamel("%sSignal", signal))).Block(
									j.Id("Channel").Op(":").Qual(workflowPkg, "GetSignalChannel").Call(
										j.Id("ctx"), m.Qual(signal, m.toCamel("%sSignalName", signal)),
									).Op(","),
								).Op(",")
							}
						})

						// call constructor to get workflow implementation
						fn.List(j.Id("wf"), j.Err()).Op(":=").Id("ctor").Call(
							j.Id("ctx"), j.Id("input"),
						)
						fn.If(j.Err().Op("!=").Nil()).Block(
							j.ReturnFunc(func(returnVals *j.Group) {
								if hasOutput {
									returnVals.Nil()
								}
								returnVals.Err()
							}),
						)

						fn.If(
							j.List(j.Id("initializable"), j.Id("ok")).Op(":=").Id("wf").Op(".").Parens(j.Qual(helpersPkg, "Initializable")),
							j.Id("ok"),
						).Block(
							j.If(j.Err().Op(":=").Id("initializable").Dot("Initialize").Call(j.Id("ctx")), j.Err().Op("!=").Nil()).Block(
								j.ReturnFunc(func(returnVals *j.Group) {
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
								j.Err().Op(":=").Qual(workflowPkg, "SetQueryHandler").Call(
									j.Id("ctx"), j.Id(m.toCamel("%sQueryName", query)), j.Id("wf").Dot(query),
								),
								j.Err().Op("!=").Nil(),
							).Block(
								j.ReturnFunc(func(returnVals *j.Group) {
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
							updateOpts := m.updates[update]

							fn.BlockFunc(func(b *j.Group) {
								// build UpdateHandlerOptions
								var updateHandlerOptions []j.Code
								if updateOpts.GetValidate() {
									updateHandlerOptions = append(updateHandlerOptions, j.Id("Validator").Op(":").Id("wf").Dot(m.toCamel("Validate%s", update)))
								}
								b.Id("opts").Op(":=").Qual(workflowPkg, "UpdateHandlerOptions").Values(updateHandlerOptions...)

								b.If(
									j.Err().Op(":=").Qual(workflowPkg, "SetUpdateHandlerWithOptions").Call(
										j.Id("ctx"), j.Id(m.toCamel("%sUpdateName", update)), j.Id("wf").Dot(m.methods[update].GoName), j.Id("opts"),
									),
									j.Err().Op("!=").Nil(),
								).Block(
									j.ReturnFunc(func(returnVals *j.Group) {
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
							j.Id("wf").Dot("Execute").Call(j.Id("ctx")),
						)
					}),
			),
		)
}

// genWorkerRegisterWorkflow generates a Register<Workflow> public function
func (m *Manifest) genWorkerRegisterWorkflow(f *j.File, workflow protoreflect.FullName) {
	method := m.methods[workflow]
	opts := m.workflows[workflow]
	builderName := fmt.Sprintf("build%s", method.GoName)
	varName := m.toCamel("%sFunction", workflow)

	// generate Register<Workflow> function
	f.Commentf("Register%sWorkflow registers a %s workflow with the given worker", m.methods[workflow].GoName, method.Desc.FullName())
	f.Func().
		Id(m.toCamel("Register%sWorkflow", workflow)).
		Params(
			j.Id("r").Qual(workerPkg, "WorkflowRegistry"),
			j.Id("wf").
				Func().
				ParamsFunc(func(args *j.Group) {
					args.Qual(workflowPkg, "Context")
					if m.cfg.DisableWorkflowInputRename {
						args.Op("*").Id(m.toCamel("%sInput", workflow))
					} else {
						args.Op("*").Id(m.toCamel("%sWorkflowInput", workflow))
					}
				}).
				Params(
					j.Id(m.toCamel("%sWorkflow", workflow)),
					j.Error(),
				),
		).
		BlockFunc(func(fn *j.Group) {
			fn.Id(varName).Op("=").Id(builderName).Call(j.Id("wf"))
			names := []j.Code{j.Id(m.toCamel("%sWorkflowName", workflow))}
			aliases := opts.GetAliases()
			sort.Strings(aliases)
			for _, alias := range aliases {
				names = append(names, j.Lit(alias))
			}
			for _, name := range names {
				fn.Id("r").Dot("RegisterWorkflowWithOptions").Call(
					j.Id(varName),
					j.Qual(workflowPkg, "RegisterOptions").Values(
						j.Id("Name").Op(":").Add(name),
					),
				)
			}
		})
}

// genWorkerRegisterWorkflows generates a public RegisterWorkflows method for a given service
func (m *Manifest) genWorkerRegisterWorkflows(f *j.File) {
	f.Commentf("Register%sWorkflows registers %s workflows with the given worker", m.Service.GoName, m.Service.Desc.FullName())
	f.Func().
		Id(m.toCamel("Register%sWorkflows", m.Service.GoName)).
		Params(
			j.Id("r").Qual(workerPkg, "WorkflowRegistry"),
			j.Id("workflows").Id(m.toCamel("%sWorkflows", m.Service.GoName)),
		).
		BlockFunc(func(fn *j.Group) {
			for _, workflow := range m.workflowsOrdered {
				if m.methods[workflow].Desc.Parent() != m.Service.Desc {
					continue
				}
				fn.Id(m.toCamel("Register%sWorkflow", workflow)).Call(
					j.Id("r"), j.Id("workflows").Dot(m.methods[workflow].GoName),
				)
			}
		})
}

// genWorkerSignal generates a worker signal struct
func (m *Manifest) genWorkerSignal(f *j.File, signal protoreflect.FullName) {
	typeName := m.toCamel("%sSignal", signal)

	f.Commentf("%s describes a(n) %s signal", typeName, m.methods[signal].Desc.FullName())
	f.Type().Id(typeName).Struct(
		j.Id("Channel").Qual(workflowPkg, "ReceiveChannel"),
	)
}

func (m *Manifest) genWorkerSignalConstructor(f *j.File, signal protoreflect.FullName) {
	typeName := m.toCamel("%sSignal", signal)
	funcName := m.toCamel("New%sSignal", signal)

	f.Commentf("%s initializes a new %s signal wrapper", funcName, m.fqnForSignal(signal))
	f.Func().
		Id(funcName).
		Params(j.Id("ctx").Qual(workflowPkg, "Context")).
		Op("*").Id(typeName).
		Block(
			j.Return(
				j.Op("&").Id(typeName).ValuesFunc(func(fields *j.Group) {
					fields.Id("Channel").Op(":").Qual(workflowPkg, "GetSignalChannel").Call(j.Id("ctx"), j.Id(m.toCamel("%sSignalName", signal)))
				}),
			),
		)
}

// genWorkerSignalExternal generates a <Signal>External public function
func (m *Manifest) genWorkerSignalExternal(f *j.File, signal protoreflect.FullName) {
	functionName := m.toCamel("%sExternal", signal)
	method := m.methods[signal]
	hasInput := !isEmpty(method.Input)

	commentWithDefaultf(f, methodSet(method), "%s sends a(n) %s signal to an existing workflow", functionName, m.fqnForSignal(signal))
	f.Func().Id(functionName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
		}).
		Params(j.Error()).
		Block(
			j.Return(
				j.Id(m.toCamel("%sAsync", functionName)).CallFunc(func(args *j.Group) {
					args.Id("ctx")
					args.Id("workflowID")
					args.Id("runID")
					if hasInput {
						args.Id("req")
					}
				}).Dot("Get").Call(j.Id("ctx"), j.Nil()),
			),
		)
}

// genWorkerSignalExternalAsync generates a <Signal>ExternalAsync public function
func (m *Manifest) genWorkerSignalExternalAsync(f *j.File, signal protoreflect.FullName) {
	functionName := m.toCamel("%sExternalAsync", signal)
	method := m.methods[signal]
	hasInput := !isEmpty(method.Input)

	commentWithDefaultf(f, methodSet(method), "%s sends a(n) %s signal to an existing workflow", functionName, m.fqnForSignal(signal))
	f.Func().Id(functionName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
		}).
		Params(j.Qual(workflowPkg, "Future")).
		Block(
			j.Return(
				j.Qual(workflowPkg, "SignalExternalWorkflow").CallFunc(func(args *j.Group) {
					args.Id("ctx")
					args.Id("workflowID")
					args.Id("runID")
					args.Id(m.toCamel("%sSignalName", signal))
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
func (m *Manifest) genWorkerSignalReceive(f *j.File, signal protoreflect.FullName) {
	method := m.methods[signal]
	hasInput := !isEmpty(method.Input)

	f.Commentf("Receive blocks until a(n) %s signal is received", method.Desc.FullName())
	f.Func().
		Params(j.Id("s").Op("*").Id(m.toCamel("%sSignal", signal))).
		Id("Receive").
		Params(j.Id("ctx").Qual(workflowPkg, "Context")).
		ParamsFunc(func(returnVals *j.Group) {
			if hasInput {
				returnVals.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			returnVals.Bool()
		}).
		BlockFunc(func(b *j.Group) {
			if hasInput {
				b.Var().Id("resp").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			b.Id("more").Op(":=").Id("s").Dot("Channel").Dot("Receive").CallFunc(func(args *j.Group) {
				args.Id("ctx")
				if hasInput {
					args.Op("&").Id("resp")
				} else {
					args.Nil()
				}
			})
			b.ReturnFunc(func(returnVals *j.Group) {
				if hasInput {
					returnVals.Op("&").Id("resp")
				}
				returnVals.Id("more")
			})
		})
}

// genWorkerSignalReceiveAsync generates a worker signal ReceiveAsync method
func (m *Manifest) genWorkerSignalReceiveAsync(f *j.File, signal protoreflect.FullName) {
	method := m.methods[signal]
	hasInput := !isEmpty(method.Input)

	f.Commentf("ReceiveAsync checks for a %s signal without blocking", method.Desc.FullName())
	f.Func().
		Params(j.Id("s").Op("*").Id(m.toCamel("%sSignal", signal))).
		Id("ReceiveAsync").
		Params().
		ParamsFunc(func(returnVals *j.Group) {
			if hasInput {
				returnVals.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			} else {
				returnVals.Bool()
			}
		}).
		BlockFunc(func(b *j.Group) {
			if hasInput {
				b.Var().Id("resp").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
				b.If(
					j.Id("ok").Op(":=").Id("s").Dot("Channel").Dot("ReceiveAsync").Call(
						j.Op("&").Id("resp"),
					),
					j.Op("!").Id("ok"),
				).Block(
					j.Return(j.Nil()),
				)
				b.Return(j.Op("&").Id("resp"))
			} else {
				b.Return(j.Id("s").Dot("Channel").Dot("ReceiveAsync").Call(j.Nil()))
			}
		})
}

// genWorkerSignalReceiveWithTimeout generates a worker signal ReceiveWithTimeout method
func (m *Manifest) genWorkerSignalReceiveWithTimeout(f *j.File, signal protoreflect.FullName) {
	method := m.methods[signal]
	hasInput := !isEmpty(method.Input)

	f.Commentf("ReceiveWithTimeout blocks until a(n) %s signal is received or timeout expires.", method.Desc.FullName())
	f.Comment("Returns more value of false when Channel is closed.")
	f.Comment("Returns ok value of false when no value was found in the channel for the duration of timeout or the ctx was canceled.")
	if hasInput {
		f.Comment("resp will be nil if ok is false.")
	}
	f.Func().
		Params(j.Id("s").Op("*").Id(m.toCamel("%sSignal", signal))).
		Id("ReceiveWithTimeout").
		Params(
			j.Id("ctx").Qual(workflowPkg, "Context"),
			j.Id("timeout").Qual("time", "Duration"),
		).
		ParamsFunc(func(returnVals *j.Group) {
			if hasInput {
				returnVals.Id("resp").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			returnVals.Id("ok").Bool()
			returnVals.Id("more").Bool()
		}).
		BlockFunc(func(b *j.Group) {
			if hasInput {
				b.Id("resp").Op("=").Op("&").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input)).Values()
			}
			b.If(
				j.List(j.Id("ok"), j.Id("more")).Op("=").Id("s").Dot("Channel").Dot("ReceiveWithTimeout").CallFunc(func(args *j.Group) {
					args.Id("ctx")
					args.Id("timeout")
					if hasInput {
						args.Op("&").Id("resp")
					} else {
						args.Nil()
					}
				}),
				j.Op("!").Id("ok"),
			).Block(
				j.ReturnFunc(func(returnVals *j.Group) {
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
func (m *Manifest) genWorkerSignalSelect(f *j.File, signal protoreflect.FullName) {
	method := m.methods[signal]
	hasInput := !isEmpty(method.Input)

	f.Commentf("Select checks for a(n) %s signal without blocking", method.Desc.FullName())
	f.Func().
		Params(j.Id("s").Op("*").Id(m.toCamel("%sSignal", signal))).
		Id("Select").
		Params(
			j.Id("sel").Qual(workflowPkg, "Selector"),
			j.Id("fn").Func().ParamsFunc(func(args *j.Group) {
				if hasInput {
					args.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
				}
			}),
		).
		Params(
			j.Qual(workflowPkg, "Selector"),
		).
		Block(
			j.Return(
				j.Id("sel").Dot("AddReceive").Call(
					j.Id("s").Dot("Channel"),
					j.Func().
						Params(
							j.Qual(workflowPkg, "ReceiveChannel"),
							j.Bool(),
						).
						BlockFunc(func(fn *j.Group) {
							if hasInput {
								fn.Id("req").Op(":=").Id("s").Dot("ReceiveAsync").Call()
							} else {
								fn.Id("s").Dot("ReceiveAsync").Call()
							}
							fn.If(j.Id("fn").Op("!=").Nil()).Block(
								j.Id("fn").CallFunc(func(args *j.Group) {
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
func (m *Manifest) genWorkerWorkflowChild(f *j.File, workflow protoreflect.FullName) {
	functionName := m.toCamel("%sChild", workflow)
	method := m.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	commentWithDefaultf(f, methodSet(method), "%s executes a child %s workflow and blocks until error or response received", functionName, m.fqnForWorkflow(workflow))
	f.Func().
		Id(functionName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			args.Id("options").Op("...").Op("*").Id(m.toCamel("%sChildOptions", workflow))
		}).
		ParamsFunc(func(returnVals *j.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *j.Group) {
			fn.List(j.Id("childRun"), j.Err()).Op(":=").Id(m.toCamel("%sChildAsync", workflow)).CallFunc(func(args *j.Group) {
				args.Id("ctx")
				if hasInput {
					args.Id("req")
				}
				args.Id("options").Op("...")
			})
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.ReturnFunc(func(returnVals *j.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)
			fn.Return(
				j.Id("childRun").Dot("Get").Call(j.Id("ctx")),
			)
		})
}

// genWorkerWorkflowChildAsync generates a public <Workflow>Child function
func (m *Manifest) genWorkerWorkflowChildAsync(f *j.File, workflow protoreflect.FullName) {
	functionName := m.toCamel("%sChildAsync", workflow)
	method := m.methods[workflow]
	hasInput := !isEmpty(method.Input)

	commentWithDefaultf(f, methodSet(method), "%s starts a child %s workflow and returns a handle to the child workflow run", functionName, m.fqnForWorkflow(workflow))
	f.Func().
		Id(functionName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			args.Id("options").Op("...").Op("*").Id(m.toCamel("%sChildOptions", workflow))
		}).
		Params(
			j.Op("*").Id(m.toCamel("%sChildRun", workflow)),
			j.Error(),
		).
		BlockFunc(func(fn *j.Group) {
			// initialize options
			fn.Var().Id("o").Op("*").Id(m.toCamel("%sChildOptions", workflow))
			fn.If(j.Len(j.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(j.Lit(0)).Op("!=").Nil()).Block(
				j.Id("o").Op("=").Id("options").Index(j.Lit(0)),
			).Else().Block(
				j.Id("o").Op("=").Id(m.toCamel("New%sChildOptions", workflow)).Call(),
			)

			// initialize client.StartWorkfowOptions
			fn.List(j.Id("opts"), j.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(args *j.Group) {
				args.Id("ctx")
				if hasInput {
					args.Id("req").Dot("ProtoReflect").Call()
				} else {
					args.Nil()
				}
			})
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing workflow.ChildWorkflowOptions: %w"), j.Err())),
			)

			fn.Id("ctx").Op("=").Qual(workflowPkg, "WithChildOptions").Call(j.Id("ctx"), j.Id("opts"))
			fn.Return(
				j.Op("&").Id(m.toCamel("%sChildRun", workflow)).Values(
					j.Id("Future").Op(":").Qual(workflowPkg, "ExecuteChildWorkflow").CallFunc(func(args *j.Group) {
						args.Id("ctx")
						args.Id(m.toCamel("%sWorkflowName", workflow))
						if hasInput {
							args.Id("req")
						} else {
							args.Nil()
						}
					}),
				),
				j.Nil(),
			)
		})
}

// genWorkerWorkflowChildRun generates a <Workflow>ChildRun struct
func (m *Manifest) genWorkerWorkflowChildRun(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toCamel("%sChildRun", workflow)

	f.Commentf("%s describes a child %s workflow run", typeName, m.methods[workflow].GoName)
	f.Type().Id(typeName).StructFunc(func(fields *j.Group) {
		fields.Add(j.Id("Future").Qual(workflowPkg, "ChildWorkflowFuture"))
	})
}

// genWorkerWorkflowChildRunGet generates a <Workflow>ChildRun Get method
func (m *Manifest) genWorkerWorkflowChildRunGet(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toCamel("%sChildRun", workflow)
	method := m.methods[workflow]
	hasOutput := !isEmpty(method.Output)

	f.Comment("Get blocks until the workflow is completed, returning the response value")
	f.Func().
		Params(
			j.Id("r").Op("*").Id(typeName),
		).
		Id("Get").
		Params(
			j.Id("ctx").Qual(workflowPkg, "Context"),
		).
		ParamsFunc(func(returnVals *j.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *j.Group) {
			if hasOutput {
				fn.Var().Id("resp").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			fn.If(
				j.Err().Op(":=").Id("r").Dot("Future").Dot("Get").CallFunc(func(args *j.Group) {
					args.Id("ctx")
					if hasOutput {
						args.Op("&").Id("resp")
					} else {
						args.Nil()
					}
				}),
				j.Err().Op("!=").Nil(),
			).BlockFunc(func(b *j.Group) {
				if hasOutput {
					b.Return(j.Nil(), j.Err())
				} else {
					b.Return(j.Err())
				}
			})
			if hasOutput {
				fn.Return(j.Op("&").Id("resp"), j.Nil())
			} else {
				fn.Return(j.Nil())
			}
		})
}

// genWorkerWorkflowChildRunSelect generates a <Workflow>ChildRun Select method
func (m *Manifest) genWorkerWorkflowChildRunSelect(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toCamel("%sChildRun", workflow)

	f.Comment("Select adds this completion to the selector. Callback can be nil.")
	f.Func().
		Params(
			j.Id("r").Op("*").Id(typeName),
		).
		Id("Select").
		Params(
			j.Id("sel").Qual(workflowPkg, "Selector"),
			j.Id("fn").Func().Params(j.Op("*").Id(typeName)),
		).
		Params(
			j.Qual(workflowPkg, "Selector"),
		).
		Block(
			j.Return(
				j.Id("sel").Dot("AddFuture").Call(
					j.Id("r").Dot("Future"),
					j.Func().Params(j.Qual(workflowPkg, "Future")).Block(
						j.If(j.Id("fn").Op("!=").Nil()).Block(
							j.Id("fn").Call(j.Id("r")),
						),
					),
				),
			),
		)
}

// genWorkerWorkflowChildRunSelectStart generates a <Workflow>ChildRun SelectStart method
func (m *Manifest) genWorkerWorkflowChildRunSelectStart(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toCamel("%sChildRun", workflow)

	f.Comment("SelectStart adds waiting for start to the selector. Callback can be nil.")
	f.Func().
		Params(
			j.Id("r").Op("*").Id(typeName),
		).
		Id("SelectStart").
		Params(
			j.Id("sel").Qual(workflowPkg, "Selector"),
			j.Id("fn").Func().Params(j.Op("*").Id(typeName)),
		).
		Params(
			j.Qual(workflowPkg, "Selector"),
		).
		Block(
			j.Return(
				j.Id("sel").Dot("AddFuture").Call(
					j.Id("r").Dot("Future").Dot("GetChildWorkflowExecution").Call(),
					j.Func().Params(j.Qual(workflowPkg, "Future")).Block(
						j.If(j.Id("fn").Op("!=").Nil()).Block(
							j.Id("fn").Call(j.Id("r")),
						),
					),
				),
			),
		)
}

// genWorkerWorkflowChildRunSignals generates <Workflow>ChildRun signal methods
func (m *Manifest) genWorkerWorkflowChildRunSignals(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toCamel("%sChildRun", workflow)
	opts := m.workflows[workflow]

	for _, signalOpts := range opts.GetSignal() {
		signal := getFullyQualifiedRef(workflow, signalOpts.GetRef())
		handler := m.methods[signal]
		hasInput := !isEmpty(handler.Input)
		asyncName := m.toCamel("%sAsync", signal)

		f.Commentf("%s sends a(n) %q signal request to the child workflow", m.methods[signal].GoName, m.fqnForSignal(signal))
		f.Func().
			Params(j.Id("r").Op("*").Id(typeName)).
			Id(m.methods[signal].GoName).
			ParamsFunc(func(params *j.Group) {
				params.Id("ctx").Qual(workflowPkg, "Context")
				if hasInput {
					params.Id("input").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
				}
			}).
			Params(j.Error()).
			Block(
				j.Return(j.Id("r").Dot(asyncName).CallFunc(func(args *j.Group) {
					args.Id("ctx")
					if hasInput {
						args.Id("input")
					}
				})).Dot("Get").Call(j.Id("ctx"), j.Nil()),
			)

		f.Commentf("%s sends a(n) %q signal request to the child workflow", asyncName, m.fqnForSignal(signal))
		f.Func().
			Params(j.Id("r").Op("*").Id(typeName)).
			Id(asyncName).
			ParamsFunc(func(params *j.Group) {
				params.Id("ctx").Qual(workflowPkg, "Context")
				if hasInput {
					params.Id("input").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
				}
			}).
			Params(j.Qual(workflowPkg, "Future")).
			Block(
				j.Return(j.Id("r").Dot("Future").Dot("SignalChildWorkflow").CallFunc(func(args *j.Group) {
					args.Id("ctx")
					args.Add(m.Qual(signal, m.toCamel("%sSignalName", signal)))
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
func (m *Manifest) genWorkerWorkflowChildRunWaitStart(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toCamel("%sChildRun", workflow)

	f.Comment("WaitStart waits for the child workflow to start")
	f.Func().
		Params(
			j.Id("r").Op("*").Id(typeName),
		).
		Id("WaitStart").
		Params(
			j.Id("ctx").Qual(workflowPkg, "Context"),
		).
		Params(
			j.Op("*").Qual(workflowPkg, "Execution"),
			j.Error(),
		).
		Block(
			j.Var().Id("exec").Qual(workflowPkg, "Execution"),
			j.If(
				j.Err().Op(":=").Id("r").Dot("Future").Dot("GetChildWorkflowExecution").Call().Dot("Get").Call(
					j.Id("ctx"),
					j.Op("&").Id("exec"),
				),
				j.Err().Op("!=").Nil(),
			).Block(
				j.Return(j.Nil(), j.Err()),
			),
			j.Return(j.Op("&").Id("exec"), j.Nil()),
		)
}

// genWorkerWorkflowFunctionVars generates a <Workflow>Function var for each workflow that are
// initialized on registration
func (m *Manifest) genWorkerWorkflowFunctionVars(f *j.File) {
	if len(m.workflowsOrdered) == 0 {
		return
	}
	f.Commentf("Reference to generated workflow functions")
	f.Var().DefsFunc(func(defs *j.Group) {
		for _, workflow := range m.workflowsOrdered {
			if m.methods[workflow].Desc.Parent() != m.Service.Desc {
				continue
			}
			method := m.methods[workflow]
			hasInput := !isEmpty(method.Input)
			hasOutput := !isEmpty(method.Output)
			varName := m.toCamel("%sFunction", workflow)

			commentWithDefaultf(defs, methodSet(method), "%s implements a %q workflow", varName, m.fqnForWorkflow(workflow))
			defs.Id(varName).
				Func().
				ParamsFunc(func(args *j.Group) {
					args.Qual(workflowPkg, "Context")
					if hasInput {
						args.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
					}
				}).
				ParamsFunc(func(returnVals *j.Group) {
					if hasOutput {
						returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
					}
					returnVals.Error()
				})
		}
	})

	typeName := m.toCamel("%sWorkflowFunctions", m.GoName)
	implName := m.toLowerCamel("%sWorkflowFunctions", m.GoName)
	f.Commentf("%s describes a mockable dependency for inlining workflows within other workflows", typeName)
	f.Type().Defs(
		j.Commentf("%s describes a mockable dependency for inlining workflows within other workflows", typeName),
		j.Id(typeName).InterfaceFunc(func(methods *j.Group) {
			for _, workflow := range m.workflowsOrdered {
				if m.methods[workflow].Desc.Parent() != m.Service.Desc {
					continue
				}
				methodName := m.toCamel("%s", workflow)
				method := m.methods[workflow]
				hasInput := !isEmpty(method.Input)
				hasOutput := !isEmpty(method.Output)
				commentWithDefaultf(methods, methodSet(method), "%s executes a %q workflow inline", methodName, m.fqnForWorkflow(workflow))
				methods.Id(methodName).
					ParamsFunc(func(args *j.Group) {
						args.Qual(workflowPkg, "Context")
						if hasInput {
							args.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
						}
					}).
					ParamsFunc(func(returnVals *j.Group) {
						if hasOutput {
							returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
						}
						returnVals.Error()
					})
			}
		}),

		j.Commentf("%s provides an internal %s implementation", implName, typeName),
		j.Id(implName).Struct(),
	)

	f.Func().Id(m.toCamel("New%s", typeName)).Params().Id(typeName).Block(
		j.Return(j.Op("&").Id(implName).Values()),
	)

	for _, workflow := range m.workflowsOrdered {
		if m.methods[workflow].Desc.Parent() != m.Service.Desc {
			continue
		}
		methodName := m.toCamel("%s", workflow)
		method := m.methods[workflow]
		varName := m.toCamel("%sFunction", workflow)
		hasInput := !isEmpty(method.Input)
		hasOutput := !isEmpty(method.Output)
		commentWithDefaultf(f, methodSet(method), "%s executes a %q workflow inline", methodName, m.fqnForWorkflow(workflow))
		f.Func().
			Params(j.Id("f").Op("*").Id(implName)).
			Id(methodName).
			ParamsFunc(func(args *j.Group) {
				args.Id("ctx").Qual(workflowPkg, "Context")
				if hasInput {
					args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
				}
			}).
			ParamsFunc(func(returnVals *j.Group) {
				if hasOutput {
					returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
				}
				returnVals.Error()
			}).
			Block(
				j.If(j.Id(varName).Op("==").Nil()).Block(
					j.ReturnFunc(func(returnVals *j.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Qual("errors", "New").Call(j.Lit(fmt.Sprintf("%s requires workflow registration via %s or %s", methodName, m.toCamel("Register%sWorkflows", m.GoName), m.toCamel("Register%sWorkflow", workflow))))
					}),
				),
				j.Return(j.Id(varName).CallFunc(func(args *j.Group) {
					args.Id("ctx")
					if hasInput {
						args.Id("req")
					}
				})),
			)
	}
}

// genWorkerWorkflowInput generates a <Workflow>Input struct
func (m *Manifest) genWorkerWorkflowInput(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toCamel("%sWorkflowInput", workflow)
	if m.cfg.DisableWorkflowInputRename {
		typeName = m.toCamel("%sInput", workflow)
	}
	opts := m.workflows[workflow]
	method := m.methods[workflow]
	hasInput := !isEmpty(method.Input)

	f.Commentf("%s describes the input to a(n) %s workflow constructor", typeName, m.fqnForWorkflow(workflow))
	f.Type().Id(typeName).StructFunc(func(fields *j.Group) {
		if hasInput {
			fields.Id("Req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
		}

		// add workflow signals
		for _, signalOpts := range opts.GetSignal() {
			signal := getFullyQualifiedRef(workflow, signalOpts.GetRef())
			fields.Id(m.methods[signal].GoName).Op("*").Add(m.Qual(signal, m.toCamel("%sSignal", signal)))
		}
	})
}

// genWorkerWorkflowInterface generates a <Workflow> interface
func (m *Manifest) genWorkerWorkflowInterface(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toCamel("%sWorkflow", workflow)
	opts := m.workflows[workflow]
	method := m.methods[workflow]
	hasOutput := !isEmpty(method.Output)

	// generate workflow interface
	var details []string
	if comments := method.Comments.Leading.String(); comments != "" && !strings.Contains(comments, m.fqnForWorkflow(workflow)) {
		details = append(details, fmt.Sprintf("name: \"%s\"", m.fqnForWorkflow(workflow)))
	}
	if id := opts.GetId(); id != "" {
		details = append(details, fmt.Sprintf("id: \"%s\"", id))
	}

	commentWithDefaultf(f, methodSet(method), "%s describes a(n) %s workflow implementation", typeName, m.fqnForWorkflow(workflow))
	if len(details) > 0 {
		f.Comment(" ")
		f.Commentf("workflow details: (%s)", strings.Join(details, ", "))
	}
	f.Type().Id(typeName).InterfaceFunc(func(methods *j.Group) {
		commentf(methods, methodSet(method), "Execute defines the entrypoint to a(n) %s workflow", m.fqnForWorkflow(workflow))
		methods.
			Id("Execute").
			Params(
				j.Id("ctx").Qual(workflowPkg, "Context"),
			).
			ParamsFunc(func(returnVals *j.Group) {
				if hasOutput {
					returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
				}
				returnVals.Error()
			}).
			Line()

		// add workflow query methods
		for _, queryOpts := range opts.GetQuery() {
			query := getFullyQualifiedRef(workflow, queryOpts.GetRef())
			handler := m.methods[query]
			hasInput := !isEmpty(handler.Input)

			commentWithDefaultf(methods, methodSet(handler), "%s implements a(n) %s query handler", query, m.fqnForQuery(query))
			methods.Id(m.toCamel("%s", query)).
				ParamsFunc(func(args *j.Group) {
					if hasInput {
						args.Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
				}).
				Params(
					j.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output)),
					j.Error(),
				).
				Line()
		}

		// add workflow update methods
		for _, updateOpts := range opts.GetUpdate() {
			update := getFullyQualifiedRef(workflow, updateOpts.GetRef())
			handler := m.methods[update]
			handlerOpts := m.updates[update]
			hasInput := !isEmpty(handler.Input)
			hasOutput := !isEmpty(handler.Output)

			// add Validate<Update> method if enabled
			if handlerOpts.GetValidate() {
				validatorName := m.toCamel("Validate%s", update)
				methods.Commentf("%s validates a(n) %s update", validatorName, m.fqnForUpdate(update))
				methods.Id(validatorName).
					ParamsFunc(func(args *j.Group) {
						args.Qual(workflowPkg, "Context")
						if hasInput {
							args.Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
						}
					}).
					Params(j.Error()).
					Line()
			}

			// add <Update> method
			commentWithDefaultf(methods, methodSet(handler), "%s implements a(n) %s update handler", update, m.fqnForQuery(update))
			methods.Id(m.toCamel("%s", update)).
				ParamsFunc(func(args *j.Group) {
					args.Qual(workflowPkg, "Context")
					if hasInput {
						args.Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
				}).
				ParamsFunc(func(returnVals *j.Group) {
					if hasOutput {
						returnVals.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
					}
					returnVals.Error()
				}).
				Line()
		}
	})
}

// genWorkerWorkflowsInterface generates a Workflows interface for a given service
func (m *Manifest) genWorkerWorkflowsInterface(f *j.File) {
	typeName := m.toCamel("%sWorkflows", m.Service.GoName)
	// generate workflows interface
	f.Commentf("%s provides methods for initializing new %s workflow values", typeName, m.Service.Desc.FullName())
	f.Type().Id(typeName).InterfaceFunc(func(methods *j.Group) {
		for _, workflow := range m.workflowsOrdered {
			if m.methods[workflow].Desc.Parent() != m.Service.Desc {
				continue
			}
			handler := m.methods[workflow]

			commentWithDefaultf(methods, methodSet(handler), "%s initializes a new a(n) %s implementation", m.toCamel("%s", workflow), m.toCamel("%sWorkflow", workflow))
			methods.
				Id(m.methods[workflow].GoName).
				ParamsFunc(func(args *j.Group) {
					args.Id("ctx").Qual(workflowPkg, "Context")
					if m.cfg.DisableWorkflowInputRename {
						args.Id("input").Op("*").Id(m.toCamel("%sInput", workflow))
					} else {
						args.Id("input").Op("*").Id(m.toCamel("%sWorkflowInput", workflow))
					}
				}).
				Params(
					j.Id(m.toCamel("%sWorkflow", workflow)),
					j.Error(),
				).
				Line()
		}
	})
}
