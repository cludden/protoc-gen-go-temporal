package plugin

import (
	"fmt"
	"sort"
	"strings"

	j "github.com/dave/jennifer/jen"
	"go.temporal.io/api/enums/v1"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

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
				ParamsFunc(func(g *j.Group) {
					g.Qual(workflowPkg, "Context")
					if m.cfg.DisableWorkflowInputRename {
						g.Op("*").Id(m.toCamel("%sInput", workflow))
					} else {
						g.Op("*").Id(m.toCamel("%sWorkflowInput", workflow))
					}
				}).
				Params(
					j.Id(m.toCamel("%sWorkflow", workflow)),
					j.Error(),
				),
		).
		Params(
			j.Func().
				ParamsFunc(func(g *j.Group) {
					g.Qual(workflowPkg, "Context")
					if hasInput {
						g.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
					}
				}).
				ParamsFunc(func(g *j.Group) {
					if hasOutput {
						g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
					}
					g.Error()
				}),
		).
		Block(
			j.Return(
				//g.Parens(g.Op("&").Id(workerName).Values(g.Id("wf"))).Dot(method.GoName),
				// generate <Workflow> method for worker struct
				j.Func().
					ParamsFunc(func(g *j.Group) {
						g.Id("ctx").Qual(workflowPkg, "Context")
						if hasInput {
							g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
						}
					}).
					ParamsFunc(func(g *j.Group) {
						if hasOutput {
							g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
						}
						g.Error()
					}).
					BlockFunc(func(g *j.Group) {
						// build input struct
						inputType := m.toCamel("%sWorkflowInput", workflow)
						if m.cfg.DisableWorkflowInputRename {
							inputType = m.toCamel("%sInput", workflow)
						}
						g.Id("input").Op(":=").Op("&").Id(inputType).BlockFunc(func(g *j.Group) {
							if hasInput {
								g.Id("Req").Op(":").Id("req").Op(",")
							}
							for _, s := range opts.GetSignal() {
								signal := getFullyQualifiedRef(workflow, s.GetRef())
								g.Id(m.methods[signal].GoName).Op(":").Op("&").Add(m.Qual(signal, m.toCamel("%sSignal", signal))).Block(
									j.Id("Channel").Op(":").Qual(workflowPkg, "GetSignalChannel").Call(
										j.Id("ctx"), m.Qual(signal, m.toCamel("%sSignalName", signal)),
									).Op(","),
								).Op(",")
							}
						})

						// call constructor to get workflow implementation
						g.List(j.Id("wf"), j.Err()).Op(":=").Id("ctor").Call(
							j.Id("ctx"), j.Id("input"),
						)
						g.If(j.Err().Op("!=").Nil()).Block(
							j.ReturnFunc(func(g *j.Group) {
								if hasOutput {
									g.Nil()
								}
								g.Err()
							}),
						)

						g.If(
							j.List(j.Id("initializable"), j.Id("ok")).Op(":=").Id("wf").Op(".").Parens(j.Qual(helpersPkg, "Initializable")),
							j.Id("ok"),
						).Block(
							j.If(j.Err().Op(":=").Id("initializable").Dot("Initialize").Call(j.Id("ctx")), j.Err().Op("!=").Nil()).Block(
								j.ReturnFunc(func(g *j.Group) {
									if hasOutput {
										g.Nil()
									}
									g.Err()
								}),
							),
						)

						// register query handlers
						for _, q := range opts.GetQuery() {
							query := q.GetRef()
							g.If(
								j.Err().Op(":=").Qual(workflowPkg, "SetQueryHandler").Call(
									j.Id("ctx"), j.Id(m.toCamel("%sQueryName", query)), j.Id("wf").Dot(query),
								),
								j.Err().Op("!=").Nil(),
							).Block(
								j.ReturnFunc(func(g *j.Group) {
									if hasOutput {
										g.Nil()
									}
									g.Err()
								}),
							)
						}

						// register update handlers
						for _, u := range opts.GetUpdate() {
							update := getFullyQualifiedRef(workflow, u.GetRef())
							updateOpts := m.updates[update]

							validate := updateOpts.GetValidate()
							if u.Validate != nil {
								validate = u.GetValidate()
							}

							g.BlockFunc(func(g *j.Group) {
								// build UpdateHandlerOptions
								var updateHandlerOptions []j.Code
								if validate {
									updateHandlerOptions = append(updateHandlerOptions, j.Id("Validator").Op(":").Id("wf").Dot(m.toCamel("Validate%s", update)))
								}
								g.Id("opts").Op(":=").Qual(workflowPkg, "UpdateHandlerOptions").Values(updateHandlerOptions...)

								g.If(
									j.Err().Op(":=").Qual(workflowPkg, "SetUpdateHandlerWithOptions").Call(
										j.Id("ctx"), j.Id(m.toCamel("%sUpdateName", update)), j.Id("wf").Dot(m.methods[update].GoName), j.Id("opts"),
									),
									j.Err().Op("!=").Nil(),
								).Block(
									j.ReturnFunc(func(g *j.Group) {
										if hasOutput {
											g.Nil()
										}
										g.Err()
									}),
								)
							})
						}

						// execute workflow
						g.Return(
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
				ParamsFunc(func(g *j.Group) {
					g.Qual(workflowPkg, "Context")
					if m.cfg.DisableWorkflowInputRename {
						g.Op("*").Id(m.toCamel("%sInput", workflow))
					} else {
						g.Op("*").Id(m.toCamel("%sWorkflowInput", workflow))
					}
				}).
				Params(
					j.Id(m.toCamel("%sWorkflow", workflow)),
					j.Error(),
				),
		).
		BlockFunc(func(g *j.Group) {
			registrationMutexName := m.Names().serviceRegistrationMutex(m.Service)
			g.Id(registrationMutexName).Dot("Lock").Call()
			g.Defer().Id(registrationMutexName).Dot("Unlock").Call()

			g.Id(varName).Op("=").Id(builderName).Call(j.Id("wf"))
			names := []j.Code{j.Id(m.toCamel("%sWorkflowName", workflow))}
			aliases := opts.GetAliases()
			sort.Strings(aliases)
			for _, alias := range aliases {
				names = append(names, j.Lit(alias))
			}
			for _, name := range names {
				g.Id("r").Dot("RegisterWorkflowWithOptions").Call(
					j.Id(varName),
					j.Qual(workflowPkg, "RegisterOptions").Values(
						j.DictFunc(func(d j.Dict) {
							d[j.Id("Name")] = name
							versioningBehavior := "VersioningBehaviorUnspecified"
							switch opts.GetVersioningBehavior() {
							case enums.VERSIONING_BEHAVIOR_AUTO_UPGRADE:
								versioningBehavior = "VersioningBehaviorAutoUpgrade"
							case enums.VERSIONING_BEHAVIOR_PINNED:
								versioningBehavior = "VersioningBehaviorPinned"
							}
							if versioningBehavior != "VersioningBehaviorUnspecified" {
								d[j.Id("VersioningBehavior")] = j.Qual(workflowPkg, versioningBehavior)
							}
						}),
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
		BlockFunc(func(g *j.Group) {
			for _, workflow := range m.workflowsOrdered {
				if m.methods[workflow].Desc.Parent() != m.Service.Desc {
					continue
				}
				g.Id(m.toCamel("Register%sWorkflow", workflow)).Call(
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
				j.Op("&").Id(typeName).ValuesFunc(func(g *j.Group) {
					g.Id("Channel").Op(":").Qual(workflowPkg, "GetSignalChannel").Call(j.Id("ctx"), j.Id(m.toCamel("%sSignalName", signal)))
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
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual(workflowPkg, "Context")
			g.Id("workflowID").String()
			g.Id("runID").String()
			if hasInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
		}).
		Params(j.Error()).
		Block(
			j.Return(
				j.Id(m.toCamel("%sAsync", functionName)).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("workflowID")
					g.Id("runID")
					if hasInput {
						g.Id("req")
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
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual(workflowPkg, "Context")
			g.Id("workflowID").String()
			g.Id("runID").String()
			if hasInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
		}).
		Params(j.Qual(workflowPkg, "Future")).
		Block(
			j.Return(
				j.Qual(workflowPkg, "SignalExternalWorkflow").CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("workflowID")
					g.Id("runID")
					g.Id(m.toCamel("%sSignalName", signal))
					if hasInput {
						g.Id("req")
					} else {
						g.Nil()
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
		ParamsFunc(func(g *j.Group) {
			if hasInput {
				g.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			g.Bool()
		}).
		BlockFunc(func(g *j.Group) {
			if hasInput {
				g.Var().Id("resp").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			g.Id("more").Op(":=").Id("s").Dot("Channel").Dot("Receive").CallFunc(func(g *j.Group) {
				g.Id("ctx")
				if hasInput {
					g.Op("&").Id("resp")
				} else {
					g.Nil()
				}
			})
			g.ReturnFunc(func(g *j.Group) {
				if hasInput {
					g.Op("&").Id("resp")
				}
				g.Id("more")
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
		ParamsFunc(func(g *j.Group) {
			if hasInput {
				g.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			} else {
				g.Bool()
			}
		}).
		BlockFunc(func(g *j.Group) {
			if hasInput {
				g.Var().Id("resp").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
				g.If(
					j.Id("ok").Op(":=").Id("s").Dot("Channel").Dot("ReceiveAsync").Call(
						j.Op("&").Id("resp"),
					),
					j.Op("!").Id("ok"),
				).Block(
					j.Return(j.Nil()),
				)
				g.Return(j.Op("&").Id("resp"))
			} else {
				g.Return(j.Id("s").Dot("Channel").Dot("ReceiveAsync").Call(j.Nil()))
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
		ParamsFunc(func(g *j.Group) {
			if hasInput {
				g.Id("resp").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			g.Id("ok").Bool()
			g.Id("more").Bool()
		}).
		BlockFunc(func(g *j.Group) {
			if hasInput {
				g.Id("resp").Op("=").Op("&").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input)).Values()
			}
			g.If(
				j.List(j.Id("ok"), j.Id("more")).Op("=").Id("s").Dot("Channel").Dot("ReceiveWithTimeout").CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("timeout")
					if hasInput {
						g.Op("&").Id("resp")
					} else {
						g.Nil()
					}
				}),
				j.Op("!").Id("ok"),
			).Block(
				j.ReturnFunc(func(g *j.Group) {
					if hasInput {
						g.Nil()
					}
					g.False()
					g.Id("more")
				}),
			)
			g.Return()
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
			j.Id("fn").Func().ParamsFunc(func(g *j.Group) {
				if hasInput {
					g.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
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
						BlockFunc(func(g *j.Group) {
							if hasInput {
								g.Id("req").Op(":=").Id("s").Dot("ReceiveAsync").Call()
							} else {
								g.Id("s").Dot("ReceiveAsync").Call()
							}
							g.If(j.Id("fn").Op("!=").Nil()).Block(
								j.Id("fn").CallFunc(func(g *j.Group) {
									if hasInput {
										g.Id("req")
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
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			g.Id("options").Op("...").Op("*").Id(m.toCamel("%sChildOptions", workflow))
		}).
		ParamsFunc(func(g *j.Group) {
			if hasOutput {
				g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			g.List(j.Id("childRun"), j.Err()).Op(":=").Id(m.toCamel("%sChildAsync", workflow)).CallFunc(func(g *j.Group) {
				g.Id("ctx")
				if hasInput {
					g.Id("req")
				}
				g.Id("options").Op("...")
			})
			g.If(j.Err().Op("!=").Nil()).Block(
				j.ReturnFunc(func(g *j.Group) {
					if hasOutput {
						g.Nil()
					}
					g.Err()
				}),
			)
			g.Return(
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
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual(workflowPkg, "Context")
			if hasInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			g.Id("options").Op("...").Op("*").Id(m.toCamel("%sChildOptions", workflow))
		}).
		Params(
			j.Op("*").Id(m.toCamel("%sChildRun", workflow)),
			j.Error(),
		).
		BlockFunc(func(g *j.Group) {
			// initialize options
			g.Var().Id("o").Op("*").Id(m.toCamel("%sChildOptions", workflow))
			g.If(j.Len(j.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(j.Lit(0)).Op("!=").Nil()).Block(
				j.Id("o").Op("=").Id("options").Index(j.Lit(0)),
			).Else().Block(
				j.Id("o").Op("=").Id(m.toCamel("New%sChildOptions", workflow)).Call(),
			)

			// initialize client.StartWorkfowOptions
			g.List(j.Id("opts"), j.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(g *j.Group) {
				g.Id("ctx")
				if hasInput {
					g.Id("req").Dot("ProtoReflect").Call()
				} else {
					g.Nil()
				}
			})
			g.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing workflow.ChildWorkflowOptions: %w"), j.Err())),
			)

			g.Id("ctx").Op("=").Qual(workflowPkg, "WithChildOptions").Call(j.Id("ctx"), j.Id("opts"))
			g.If(
				j.Id("o").Dot("dc").Op("!=").Nil(),
			).Block(
				j.Id("ctx").Op("=").Qual(workflowPkg, "WithDataConverter").Call(j.Id("ctx"), j.Id("o").Dot("dc")),
			)
			g.Return(
				j.Op("&").Id(m.toCamel("%sChildRun", workflow)).Values(
					j.Id("Future").Op(":").Qual(workflowPkg, "ExecuteChildWorkflow").CallFunc(func(g *j.Group) {
						g.Id("ctx")
						g.Id(m.toCamel("%sWorkflowName", workflow))
						if hasInput {
							g.Id("req")
						} else {
							g.Nil()
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
	f.Type().Id(typeName).StructFunc(func(g *j.Group) {
		g.Add(j.Id("Future").Qual(workflowPkg, "ChildWorkflowFuture"))
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
		ParamsFunc(func(g *j.Group) {
			if hasOutput {
				g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			if hasOutput {
				g.Var().Id("resp").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			g.If(
				j.Err().Op(":=").Id("r").Dot("Future").Dot("Get").CallFunc(func(g *j.Group) {
					g.Id("ctx")
					if hasOutput {
						g.Op("&").Id("resp")
					} else {
						g.Nil()
					}
				}),
				j.Err().Op("!=").Nil(),
			).BlockFunc(func(g *j.Group) {
				if hasOutput {
					g.Return(j.Nil(), j.Err())
				} else {
					g.Return(j.Err())
				}
			})
			if hasOutput {
				g.Return(j.Op("&").Id("resp"), j.Nil())
			} else {
				g.Return(j.Nil())
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
			ParamsFunc(func(g *j.Group) {
				g.Id("ctx").Qual(workflowPkg, "Context")
				if hasInput {
					g.Id("input").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
				}
			}).
			Params(j.Error()).
			Block(
				j.Return(j.Id("r").Dot(asyncName).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					if hasInput {
						g.Id("input")
					}
				})).Dot("Get").Call(j.Id("ctx"), j.Nil()),
			)

		f.Commentf("%s sends a(n) %q signal request to the child workflow", asyncName, m.fqnForSignal(signal))
		f.Func().
			Params(j.Id("r").Op("*").Id(typeName)).
			Id(asyncName).
			ParamsFunc(func(g *j.Group) {
				g.Id("ctx").Qual(workflowPkg, "Context")
				if hasInput {
					g.Id("input").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
				}
			}).
			Params(j.Qual(workflowPkg, "Future")).
			Block(
				j.Return(j.Id("r").Dot("Future").Dot("SignalChildWorkflow").CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Add(m.Qual(signal, m.toCamel("%sSignalName", signal)))
					if hasInput {
						g.Id("input")
					} else {
						g.Nil()
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
	f.Var().DefsFunc(func(g *j.Group) {
		registrationMutexName := m.Names().serviceRegistrationMutex(m.Service)
		g.Commentf("%s is a mutex for registering %s workflows", registrationMutexName, m.Service.Desc.FullName())
		g.Id(registrationMutexName).Qual("sync", "Mutex")

		for _, workflow := range m.workflowsOrdered {
			if m.methods[workflow].Desc.Parent() != m.Service.Desc {
				continue
			}
			method := m.methods[workflow]
			hasInput := !isEmpty(method.Input)
			hasOutput := !isEmpty(method.Output)
			varName := m.toCamel("%sFunction", workflow)

			commentWithDefaultf(g, methodSet(method), "%s implements a %q workflow", varName, m.fqnForWorkflow(workflow))
			g.Id(varName).
				Func().
				ParamsFunc(func(g *j.Group) {
					g.Qual(workflowPkg, "Context")
					if hasInput {
						g.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
					}
				}).
				ParamsFunc(func(g *j.Group) {
					if hasOutput {
						g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
					}
					g.Error()
				})
		}
	})

	typeName := m.toCamel("%sWorkflowFunctions", m.GoName)
	implName := m.toLowerCamel("%sWorkflowFunctions", m.GoName)
	f.Commentf("%s describes a mockable dependency for inlining workflows within other workflows", typeName)
	f.Type().Defs(
		j.Commentf("%s describes a mockable dependency for inlining workflows within other workflows", typeName),
		j.Id(typeName).InterfaceFunc(func(g *j.Group) {
			for _, workflow := range m.workflowsOrdered {
				if m.methods[workflow].Desc.Parent() != m.Service.Desc {
					continue
				}
				methodName := m.toCamel("%s", workflow)
				method := m.methods[workflow]
				hasInput := !isEmpty(method.Input)
				hasOutput := !isEmpty(method.Output)
				commentWithDefaultf(g, methodSet(method), "%s executes a %q workflow inline", methodName, m.fqnForWorkflow(workflow))
				g.Id(methodName).
					ParamsFunc(func(g *j.Group) {
						g.Qual(workflowPkg, "Context")
						if hasInput {
							g.Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
						}
					}).
					ParamsFunc(func(g *j.Group) {
						if hasOutput {
							g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
						}
						g.Error()
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
			ParamsFunc(func(g *j.Group) {
				g.Id("ctx").Qual(workflowPkg, "Context")
				if hasInput {
					g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
				}
			}).
			ParamsFunc(func(g *j.Group) {
				if hasOutput {
					g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
				}
				g.Error()
			}).
			Block(
				j.If(j.Id(varName).Op("==").Nil()).Block(
					j.ReturnFunc(func(g *j.Group) {
						if hasOutput {
							g.Nil()
						}
						g.Qual("errors", "New").Call(j.Lit(fmt.Sprintf("%s requires workflow registration via %s or %s", methodName, m.toCamel("Register%sWorkflows", m.GoName), m.toCamel("Register%sWorkflow", workflow))))
					}),
				),
				j.Return(j.Id(varName).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					if hasInput {
						g.Id("req")
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
	f.Type().Id(typeName).StructFunc(func(g *j.Group) {
		if hasInput {
			g.Id("Req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
		}

		// add workflow signals
		for _, signalOpts := range opts.GetSignal() {
			signal := getFullyQualifiedRef(workflow, signalOpts.GetRef())
			g.Id(m.methods[signal].GoName).Op("*").Add(m.Qual(signal, m.toCamel("%sSignal", signal)))
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
	f.Type().Id(typeName).InterfaceFunc(func(g *j.Group) {
		commentf(g, methodSet(method), "Execute defines the entrypoint to a(n) %s workflow", m.fqnForWorkflow(workflow))
		g.
			Id("Execute").
			Params(
				j.Id("ctx").Qual(workflowPkg, "Context"),
			).
			ParamsFunc(func(g *j.Group) {
				if hasOutput {
					g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
				}
				g.Error()
			}).
			Line()

		// add workflow query methods
		for _, queryOpts := range opts.GetQuery() {
			query := getFullyQualifiedRef(workflow, queryOpts.GetRef())
			handler := m.methods[query]
			hasInput := !isEmpty(handler.Input)

			commentWithDefaultf(g, methodSet(handler), "%s implements a(n) %s query handler", query, m.fqnForQuery(query))
			g.Id(m.toCamel("%s", query)).
				ParamsFunc(func(g *j.Group) {
					if hasInput {
						g.Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
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

			validate := handlerOpts.GetValidate()
			if updateOpts.Validate != nil {
				validate = updateOpts.GetValidate()
			}

			// add Validate<Update> method if enabled
			if validate {
				validatorName := m.toCamel("Validate%s", update)
				g.Commentf("%s validates a(n) %s update", validatorName, m.fqnForUpdate(update))
				g.Id(validatorName).
					ParamsFunc(func(g *j.Group) {
						g.Qual(workflowPkg, "Context")
						if hasInput {
							g.Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
						}
					}).
					Params(j.Error()).
					Line()
			}

			// add <Update> method
			commentWithDefaultf(g, methodSet(handler), "%s implements a(n) %s update handler", update, m.fqnForQuery(update))
			g.Id(m.toCamel("%s", update)).
				ParamsFunc(func(g *j.Group) {
					g.Qual(workflowPkg, "Context")
					if hasInput {
						g.Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
				}).
				ParamsFunc(func(g *j.Group) {
					if hasOutput {
						g.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
					}
					g.Error()
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
	f.Type().Id(typeName).InterfaceFunc(func(g *j.Group) {
		for _, workflow := range m.workflowsOrdered {
			if m.methods[workflow].Desc.Parent() != m.Service.Desc {
				continue
			}
			handler := m.methods[workflow]

			commentWithDefaultf(g, methodSet(handler), "%s initializes a new a(n) %s implementation", m.toCamel("%s", workflow), m.toCamel("%sWorkflow", workflow))
			g.
				Id(m.methods[workflow].GoName).
				ParamsFunc(func(g *j.Group) {
					g.Id("ctx").Qual(workflowPkg, "Context")
					if m.cfg.DisableWorkflowInputRename {
						g.Id("input").Op("*").Id(m.toCamel("%sInput", workflow))
					} else {
						g.Id("input").Op("*").Id(m.toCamel("%sWorkflowInput", workflow))
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

func (n *names) serviceRegistrationMutex(service *protogen.Service) string {
	return n.toLowerCamel("%sRegistrationMutex", service.GoName)
}

func (n *names) workflowIDExpression(workflow protoreflect.FullName) string {
	return n.toCamel("%sIDExpression", workflow)
}
