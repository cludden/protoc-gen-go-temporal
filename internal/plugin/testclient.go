package plugin

import (
	j "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	testsuitePkg = "go.temporal.io/sdk/testsuite"
	testutilPkg  = "github.com/cludden/protoc-gen-go-temporal/pkg/testutil"
)

func (n *names) testClient() string {
	return n.toCamel("Test%sClient", n.Manifest.Service.GoName)
}

func (n *names) testClientUpdateHandle(update protoreflect.FullName) string {
	return n.toLowerCamel("test%sHandle", update)
}

// genTestClientImpl generates a TestClient struct
func (m *Manifest) genTestClientImpl(f *j.File) {
	f.Comment("TestClient provides a testsuite-compatible Client")
	f.Type().Id(m.toCamel("Test%sClient", m.Service.GoName)).Struct(
		j.Id("env").Op("*").Qual(testsuitePkg, "TestWorkflowEnvironment"),
		j.Id("workflows").Id(m.toCamel("%sWorkflows", m.Service.GoName)),
	)
}

// genTestClientImplNewMethod generates a NewTestClient constructor function
func (m *Manifest) genTestClientImplNewMethod(f *j.File) {
	interfaceName := m.toCamel("%sClient", m.Service.GoName)
	typeName := m.toCamel("Test%sClient", m.Service.GoName)
	functionName := "New" + typeName

	f.Var().Id("_").Id(interfaceName).Op("=").Op("&").Id(typeName).Values()
	f.Commentf("%s initializes a new %s value", functionName, typeName)
	f.Func().Id(functionName).
		Params(
			j.Id("env").Op("*").Qual(testsuitePkg, "TestWorkflowEnvironment"),
			j.Id("workflows").Id(m.toCamel("%sWorkflows", m.Service.GoName)),
			j.Id("activities").Id(m.toCamel("%sActivities", m.Service.GoName)),
		).
		Params(
			j.Op("*").Id(typeName),
		).
		Block(
			j.If(j.Id("workflows").Op("!=").Nil()).Block(
				j.Id(m.toCamel("Register%sWorkflows", m.Service.GoName)).Call(j.Id("env"), j.Id("workflows")),
			),
			j.If(j.Id("activities").Op("!=").Nil()).Block(
				j.Id(m.toCamel("Register%sActivities", m.Service.GoName)).Call(j.Id("env"), j.Id("activities")),
			),
			j.Return(j.Op("&").Id(typeName).Values(j.Id("env"), j.Id("workflows"))),
		)
}

// genTestClientImplQueryMethod genereates a TestClient <Query> method
func (m *Manifest) genTestClientImplQueryMethod(f *j.File, query protoreflect.FullName) {
	handler := m.methods[query]
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)

	f.Commentf("%s executes a %s query", m.methods[query].GoName, m.fqnForQuery(query))
	f.Func().
		Params(j.Id("c").Op("*").Id(m.toCamel("Test%sClient", m.Service.GoName))).
		Id(m.methods[query].GoName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			g.Id("workflowID").String()
			g.Id("runID").String()
			if hasInput {
				g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
		}).
		ParamsFunc(func(g *j.Group) {
			if hasOutput {
				g.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
			}
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			g.List(j.Id("val"), j.Err()).Op(":=").Id("c").Dot("env").Dot("QueryWorkflow").CallFunc(func(g *j.Group) {
				g.Id(m.toCamel("%sQueryName", query))
				if hasInput {
					g.Id("req")
				}
			})
			g.If(j.Err().Op("!=").Nil()).Block(
				j.ReturnFunc(func(g *j.Group) {
					if hasOutput {
						g.Nil()
					}
					g.Err()
				}),
			).Else().If(j.Op("!").Id("val").Dot("HasValue").Call()).Block(
				j.ReturnFunc(func(g *j.Group) {
					if hasOutput {
						g.Nil()
					}
					g.Nil()
				}),
			).Else().BlockFunc(func(g *j.Group) {
				if !hasOutput {
					g.Return(j.Nil())
				} else {
					g.Var().Id("result").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
					g.If(j.Err().Op(":=").Id("val").Dot("Get").Call(j.Op("&").Id("result")), j.Err().Op("!=").Nil()).Block(
						j.Return(
							j.Nil(),
							j.Err(),
						),
					)
					g.Return(j.Op("&").Id("result"), j.Nil())
				}
			})
		})
}

// genTestClientImplSignalMethod genereates a TestClient <Signal> method
func (m *Manifest) genTestClientImplSignalMethod(f *j.File, signal protoreflect.FullName) {
	handler := m.methods[signal]
	hasInput := !isEmpty(handler.Input)

	f.Commentf("%s executes a %s signal", m.methods[signal].GoName, m.fqnForSignal(signal))
	f.Func().
		Params(j.Id("c").Op("*").Id(m.toCamel("Test%sClient", m.Service.GoName))).
		Id(m.methods[signal].GoName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			g.Id("workflowID").String()
			g.Id("runID").String()
			if hasInput {
				g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
		}).
		Params(
			j.Error(),
		).
		Block(
			j.Id("c").Dot("env").Dot("SignalWorkflow").CallFunc(func(g *j.Group) {
				g.Id(m.toCamel("%sSignalName", signal))
				if hasInput {
					g.Id("req")
				} else {
					g.Nil()
				}
			}),
			j.Return(j.Nil()),
		)
}

func (m *Manifest) genTestClientImplUpdateGetMethod(f *j.File, update protoreflect.FullName) {
	methodName := m.toCamel("Get%s", update)

	f.Commentf("%s retrieves a handle to an existing %s update", methodName, m.fqnForUpdate(update))
	f.Func().
		Params(j.Id("c").Op("*").Id(m.toCamel("Test%sClient", m.Service.GoName))).
		Id(methodName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			g.Id("req").Qual(clientPkg, "GetWorkflowUpdateHandleOptions")
		}).
		Params(
			j.Id(m.toCamel("%sHandle", update)),
			j.Error(),
		).
		Block(
			j.Return(
				j.Nil(),
				j.Qual("errors", "New").Call(j.Lit("unimplemented")),
			),
		)
}

// genTestClientImplUpdateMethod genereates a TestClient <Update> method
func (m *Manifest) genTestClientImplUpdateMethod(f *j.File, update protoreflect.FullName) {
	method := m.methods[update]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	asyncName := m.toCamel("%sAsync", update)

	f.Commentf("%s executes a(n) %s update in the test environment", m.methods[update].GoName, m.fqnForUpdate(update))
	f.Func().
		Params(j.Id("c").Op("*").Id(m.toCamel("Test%sClient", m.Service.GoName))).
		Id(m.methods[update].GoName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			g.Id("workflowID").String()
			g.Id("runID").String()
			if hasInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			g.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", update))
		}).
		ParamsFunc(func(g *j.Group) {
			if hasOutput {
				g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			g.Error()
		}).
		Block(
			// initialize update request options
			j.Id("options").Op(":=").Id(m.toCamel("New%sOptions", update)).Call(),
			j.If(j.Len(j.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(j.Lit(0)).Dot("Options").Op("!=").Nil()).Block(
				j.Id("options").Op("=").Id("opts").Index(j.Lit(0)),
			),

			// override wait policy
			j.Id("options").Dot("Options").Dot("WaitForStage").Op("=").Qual(clientPkg, "WorkflowUpdateStageCompleted"),
			j.List(j.Id("handle"), j.Err()).Op(":=").Id("c").Dot(asyncName).CallFunc(func(g *j.Group) {
				g.Id("ctx")
				g.Id("workflowID")
				g.Id("runID")
				if hasInput {
					g.Id("req")
				}
				g.Id("options")
			}),
			j.If(j.Err().Op("!=").Nil()).Block(
				j.ReturnFunc(func(g *j.Group) {
					if hasOutput {
						g.Nil()
					}
					g.Err()
				}),
			),
			j.Return(j.Id("handle").Dot("Get").Call(j.Id("ctx"))),
		)
}

// genTestClientImplUpdateMethodAsync genereates a TestClient <UpdateAsync> method
func (m *Manifest) genTestClientImplUpdateMethodAsync(f *j.File, update protoreflect.FullName) {
	method := m.methods[update]
	hasInput := !isEmpty(method.Input)
	methodName := m.toCamel("%sAsync", update)

	f.Commentf("%s executes a(n) %s update in the test environment", methodName, m.fqnForUpdate(update))
	f.Func().
		Params(j.Id("c").Op("*").Id(m.toCamel("Test%sClient", m.Service.GoName))).
		Id(methodName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			g.Id("workflowID").String()
			g.Id("runID").String()
			if hasInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			g.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", update))
		}).
		Params(
			j.Id(m.toCamel("%sHandle", update)),
			j.Error(),
		).
		BlockFunc(func(g *j.Group) {
			// initialize options
			g.Var().Id("o").Op("*").Id(m.toCamel("%sOptions", update))
			g.If(j.Len(j.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(j.Lit(0)).Op("!=").Nil()).Block(
				j.Id("o").Op("=").Id("opts").Index(j.Lit(0)),
			).Else().Block(
				j.Id("o").Op("=").Id(m.toCamel("New%sOptions", update)).Call(),
			)

			// build UpdateWorkflowWithOptions
			g.List(j.Id("options"), j.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(g *j.Group) {
				g.Id("workflowID")
				g.Id("runID")
				if hasInput {
					g.Id("req")
				}
			})
			g.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing UpdateWorkflowWithOptions: %w"), j.Err())),
			).Line()

			g.If(j.Id("options").Dot("UpdateID").Op("==").Lit("")).Block(
				j.Id("options").Dot("UpdateID").Op("=").Qual(uuidPkg, "New").Call().Dot("String").Call(),
			).Line()

			// update workflow
			g.Id("uc").Op(":=").Qual(testutilPkg, "NewUpdateCallbacks").Call()
			g.Id("c").Dot("env").Dot("UpdateWorkflow").CallFunc(func(g *j.Group) {
				g.Id(m.toCamel("%sUpdateName", update))
				g.Id("options").Dot("UpdateID")
				g.Id("uc")
				if hasInput {
					g.Id("req")
				}
			})

			g.Return(
				j.Op("&").Id(m.toLowerCamel("test%sHandle", update)).CustomFunc(multiLineValues, func(g *j.Group) {
					g.Id("callbacks").Op(":").Id("uc")
					g.Id("env").Op(":").Id("c").Dot("env")
					g.Id("opts").Op(":").Id("options")
					g.Id("runID").Op(":").Id("runID")
					g.Id("workflowID").Op(":").Id("workflowID")
					if hasInput {
						g.Id("req").Op(":").Id("req")
					}
				}),
				j.Nil(),
			)
		})
}

func (m *Manifest) genTestClientImplUpdateWithStartMethod(f *j.File, workflow, update protoreflect.FullName) {
	method := m.methods[workflow]
	hasWorkflowInput := !isEmpty(method.Input)
	handler := m.methods[update]
	hasUpdateInput := !isEmpty(handler.Input)
	hasUpdateOutput := !isEmpty(handler.Output)

	asyncName := m.Names().clientUpdateWithStartAsync(workflow, update)
	client := m.Names().testClient()
	options := m.Names().clientUpdateWithStartOptions(workflow, update)
	optionsCtor := m.Names().clientUpdateWithStartOptionsCtor(workflow, update)
	run := m.Names().clientWorkflowRun(workflow)

	methodName := m.Names().clientUpdateWithStart(workflow, update)
	f.Commentf("%s executes a(n) %s workflow and a(n) %s update in the test environment", methodName, m.fqnForWorkflow(workflow), m.fqnForUpdate(update))
	f.Func().
		Params(j.Id("c").Op("*").Id(client)).
		Id(methodName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasWorkflowInput {
				g.Id("input").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if hasUpdateInput {
				g.Id("update").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			g.Id("options").Op("...").Op("*").Id(options)
		}).
		ParamsFunc(func(g *j.Group) {
			if hasUpdateOutput {
				g.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
			}
			g.Id(run)
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			// initialize method options
			g.Var().Id("o").Op("*").Id(options)
			g.IfFunc(func(g *j.Group) {
				g.Len(j.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(j.Lit(0)).Op("!=").Nil()
			}).BlockFunc(func(g *j.Group) {
				g.Id("o").Op("=").Id("options").Index(j.Lit(0))
			}).Else().BlockFunc(func(g *j.Group) {
				g.Id("o").Op("=").Id(optionsCtor).Call()
			})

			// invoke async method
			g.List(j.Id("handle"), j.Id("run"), j.Err()).Op(":=").Id("c").Dot(asyncName).CallFunc(func(g *j.Group) {
				g.Id("ctx")
				if hasWorkflowInput {
					g.Id("input")
				}
				if hasUpdateInput {
					g.Id("update")
				}
				g.Id("o")
			})
			g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
				g.ReturnFunc(func(g *j.Group) {
					if hasUpdateOutput {
						g.Nil()
					}
					g.Id("run")
					g.Err()
				})
			})

			// execute workflow to prevent deadlock on update callback
			g.Id("run").Dot("Get").Call(j.Id("ctx"))

			// await update result
			if hasUpdateOutput {
				g.List(j.Id("out"), j.Err()).Op(":=").Id("handle").Dot("Get").Call(j.Id("ctx"))
				g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
					g.Return(j.Nil(), j.Id("run"), j.Err())
				})
			} else {
				g.IfFunc(func(g *j.Group) {
					g.Err().Op(":=").Id("handle").Dot("Get").Call(j.Id("ctx"))
					g.Err().Op("!=").Nil()
				}).BlockFunc(func(g *j.Group) {
					g.Return(j.Id("run"), j.Err())
				})
			}

			g.ReturnFunc(func(g *j.Group) {
				if hasUpdateOutput {
					g.Id("out")
				}
				g.Id("run")
				g.Nil()
			})
		})
}

func (m *Manifest) genTestClientImplUpdateWithStartMethodAsync(f *j.File, workflow, update protoreflect.FullName) {
	method := m.methods[workflow]
	hasWorkflowInput := !isEmpty(method.Input)
	handler := m.methods[update]
	hasUpdateInput := !isEmpty(handler.Input)

	async := m.Names().clientUpdateWithStartAsync(workflow, update)
	client := m.Names().testClient()
	options := m.Names().clientUpdateWithStartOptions(workflow, update)
	optionsCtor := m.Names().clientUpdateWithStartOptionsCtor(workflow, update)
	workflowOptionsCtor := m.Names().clientWorkflowOptionsCtor(workflow)
	updateOptionsCtor := m.Names().clientUpdateOptionsCtor(update)
	run := m.Names().clientWorkflowRun(workflow)
	handle := m.Names().clientUpdateHandleIface(update)
	testHandle := m.Names().testClientUpdateHandle(update)
	updateName := m.Names().updateName(update)

	f.Commentf("%s executes a(n) %s workflow and a(n) %s update in the test environment", async, m.fqnForWorkflow(workflow), m.fqnForUpdate(update))
	f.Func().
		Params(j.Id("c").Op("*").Id(client)).
		Id(async).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasWorkflowInput {
				g.Id("input").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if hasUpdateInput {
				g.Id("update").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			g.Id("options").Op("...").Op("*").Id(options)
		}).
		Params(
			j.Id(handle),
			j.Id(run),
			j.Error(),
		).
		BlockFunc(func(g *j.Group) {
			// initialize method options
			g.Var().Id("o").Op("*").Id(options)
			g.If(j.Len(j.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(j.Lit(0)).Op("!=").Nil()).BlockFunc(func(g *j.Group) {
				g.Id("o").Op("=").Id("options").Index(j.Lit(0))
			}).Else().BlockFunc(func(g *j.Group) {
				g.Id("o").Op("=").Id(optionsCtor).Call()
			})

			// initialize start workflow options
			g.If(j.Id("o").Dot("workflowOptions").Op("==").Nil()).BlockFunc(func(g *j.Group) {
				g.Id("o").Dot("workflowOptions").Op("=").Id(workflowOptionsCtor).Call()
			})
			g.List(j.Id("swo"), j.Err()).Op(":=").Id("o").Dot("workflowOptions").Dot("Build").CallFunc(func(g *j.Group) {
				if hasWorkflowInput {
					g.Id("input").Dot("ProtoReflect").Call()
				} else {
					g.Nil()
				}
			})
			g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
				g.Return(j.Nil(), j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing workflowOptions: %w"), j.Err()))
			})

			// initialize update workflow options
			g.If(j.Id("o").Dot("updateOptions").Op("==").Nil()).BlockFunc(func(g *j.Group) {
				g.Id("o").Dot("updateOptions").Op("=").Id(updateOptionsCtor).Call()
			})
			g.List(j.Id("uo"), j.Err()).Op(":=").Id("o").Dot("updateOptions").Dot("Build").CallFunc(func(g *j.Group) {
				g.Id("swo").Dot("ID")
				g.Lit("")
				if hasUpdateInput {
					g.Id("update")
				}
			})
			g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
				g.Return(j.Nil(), j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing updateOptions: %w"), j.Err()))
			})

			// initialize test run
			g.List(j.Id("run"), j.Err()).Op(":=").Id("c").Dot(m.Names().clientWorkflowAsync(workflow)).CallFunc(func(g *j.Group) {
				g.Id("ctx")
				if hasWorkflowInput {
					g.Id("input")
				}
			})
			g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
				g.Return(j.Nil(), j.Nil(), j.Err())
			})

			// initialize update callbacks
			g.Id("uc").Op(":=").Qual(testutilPkg, "NewUpdateCallbacks").Call()

			// register delayed callback
			g.Id("c").Dot("env").Dot("RegisterDelayedCallback").CallFunc(func(g *j.Group) {
				g.Func().Params().BlockFunc(func(g *j.Group) {
					g.Id("c").Dot("env").Dot("UpdateWorkflow").CallFunc(func(g *j.Group) {
						g.Id(updateName)
						g.Id("uo").Dot("UpdateID")
						g.Id("uc")
						if hasUpdateInput {
							g.Id("update")
						} else {
							g.Nil()
						}
					})
				})
				g.Lit(0)
			})

			g.ReturnFunc(func(g *j.Group) {
				g.Op("&").Id(testHandle).Values(j.DictFunc(func(d j.Dict) {
					d[j.Id("callbacks")] = j.Id("uc")
					d[j.Id("env")] = j.Id("c").Dot("env")
					d[j.Id("opts")] = j.Id("uo")
					d[j.Id("runID")] = j.Lit("")
					d[j.Id("workflowID")] = j.Id("swo").Dot("ID")
					if hasUpdateInput {
						d[j.Id("req")] = j.Id("update")
					}
				}))
				g.Id("run")
				g.Nil()
			})
		})

}

// genClientImplWorkflowCancelMethod generates a Cancel<Workflow> client method
func (m *Manifest) genTestClientImplWorkflowCancelMethod(f *j.File) {
	clientType := m.toCamel("Test%sClient", m.Service.GoName)
	methodName := "CancelWorkflow"

	f.Commentf("%s requests cancellation of an existing workflow execution", methodName)
	f.Func().
		Params(j.Id("c").Op("*").Id(clientType)).
		Id(methodName).
		Params(
			j.Id("ctx").Qual("context", "Context"),
			j.Id("workflowID").String(),
			j.Id("runID").String(),
		).
		Params(
			j.Error(),
		).
		Block(
			j.Id("c").Dot("env").Dot("CancelWorkflow").Call(),
			j.Return(j.Nil()),
		)
}

// genTestClientImplWorkflowGetMethod generates a TestClient's Get<workflow> method
func (m *Manifest) genTestClientImplWorkflowGetMethod(f *j.File, workflow protoreflect.FullName) {
	f.Commentf("%s is a noop", m.toCamel("Get%s", workflow))
	f.Func().
		Params(j.Id("c").Op("*").Id(m.toCamel("Test%sClient", m.Service.GoName))).
		Id(m.toCamel("Get%s", workflow)).
		Params(
			j.Id("ctx").Qual("context", "Context"),
			j.Id("workflowID").String(),
			j.Id("runID").String(),
		).
		Params(
			j.Id(m.toCamel("%sRun", workflow)),
		).
		Block(
			j.Return(
				j.Op("&").Id(m.toLowerCamel("test%sRun", workflow)).Values(
					j.Id("env").Op(":").Id("c").Dot("env"),
					j.Id("workflows").Op(":").Id("c").Dot("workflows"),
				),
			),
		)
}

// genTestClientImplWorkflowMethod generates a TestClient <Workflow> method
func (m *Manifest) genTestClientImplWorkflowMethod(f *j.File, workflow protoreflect.FullName) {
	method := m.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s executes a(n) %s workflow in the test environment", m.methods[workflow].GoName, m.fqnForWorkflow(workflow))
	f.Func().
		Params(j.Id("c").Op("*").Id(m.toCamel("Test%sClient", m.Service.GoName))).
		Id(m.methods[workflow].GoName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			g.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
		}).
		ParamsFunc(func(g *j.Group) {
			if hasOutput {
				g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			g.Error()
		}).
		Block(
			j.List(j.Id("run"), j.Err()).Op(":=").Id("c").Dot(m.toCamel("%sAsync", workflow)).CallFunc(func(g *j.Group) {
				g.Id("ctx")
				if hasInput {
					g.Id("req")
				}
				g.Id("opts").Op("...")
			}),
			j.If(j.Err().Op("!=").Nil()).Block(
				j.ReturnFunc(func(g *j.Group) {
					if hasOutput {
						g.Nil()
					}
					g.Err()
				}),
			),
			j.Return(j.Id("run").Dot("Get").Call(j.Id("ctx"))),
		)
}

// genTestClientImplWorkflowMethodAsync generates a TestClient's <workflow>Async method
func (m *Manifest) genTestClientImplWorkflowMethodAsync(f *j.File, workflow protoreflect.FullName) {
	method := m.methods[workflow]
	hasInput := !isEmpty(method.Input)

	f.Commentf("%sAsync executes a(n) %s workflow in the test environment", m.methods[workflow].GoName, m.fqnForWorkflow(workflow))
	f.Func().
		Params(j.Id("c").Op("*").Id(m.toCamel("Test%sClient", m.Service.GoName))).
		Id(m.toCamel("%sAsync", workflow)).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			g.Id("options").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
		}).
		Params(
			j.Id(m.toCamel("%sRun", workflow)),
			j.Error(),
		).
		BlockFunc(func(g *j.Group) {
			// initialize options
			g.Var().Id("o").Op("*").Id(m.toCamel("%sOptions", workflow))
			g.If(j.Len(j.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(j.Lit(0)).Op("!=").Nil()).Block(
				j.Id("o").Op("=").Id("options").Index(j.Lit(0)),
			).Else().Block(
				j.Id("o").Op("=").Id(m.toCamel("New%sOptions", workflow)).Call(),
			)

			// initialize client.StartWorkfowOptions
			g.List(j.Id("opts"), j.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(g *j.Group) {
				if hasInput {
					g.Id("req").Dot("ProtoReflect").Call()
				} else {
					g.Nil()
				}
			})
			g.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing client.StartWorkflowOptions: %w"), j.Err())),
			)

			g.Return(
				j.Op("&").Id(m.toLowerCamel("test%sRun", workflow)).ValuesFunc(func(g *j.Group) {
					g.Id("client").Op(":").Id("c")
					g.Id("env").Op(":").Id("c").Dot("env")
					g.Id("opts").Op(":").Op("&").Id("opts")
					if hasInput {
						g.Id("req").Op(":").Id("req")
					}
					g.Id("workflows").Op(":").Id("c").Dot("workflows")
				}),
				j.Nil(),
			)
		})
}

// genClientImplWorkflowTerminateMethod generates a Terminate<Workflow> client method
func (m *Manifest) genTestClientImplWorkflowTerminateMethod(f *j.File) {
	clientType := m.toCamel("Test%sClient", m.Service.GoName)
	methodName := "TerminateWorkflow"

	f.Commentf("%s terminates an existing workflow execution", methodName)
	f.Func().
		Params(j.Id("c").Op("*").Id(clientType)).
		Id(methodName).
		Params(
			j.Id("ctx").Qual("context", "Context"),
			j.Id("workflowID").String(),
			j.Id("runID").String(),
			j.Id("reason").String(),
			j.Id("details").Op("...").Interface(),
		).
		Params(
			j.Error(),
		).
		Block(
			j.Return(
				j.Id("c").Dot("CancelWorkflow").Call(
					j.Id("ctx"),
					j.Id("workflowID"),
					j.Id("runID"),
				),
			),
		)
}

// genTestClientImplWorkflowWithSignalMethod generates a TestClient's <workflow>With<signal> method
func (m *Manifest) genTestClientImplWorkflowWithSignalMethod(f *j.File, workflow, signal protoreflect.FullName) {
	method := m.methods[workflow]
	methodName := m.toCamel("%sWith%s", workflow, signal)
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	handler := m.methods[signal]
	hasSignalInput := !isEmpty(handler.Input)

	f.Commentf("%s sends a(n) %s signal to a(n) %s workflow, starting it if necessary", methodName, m.fqnForSignal(signal), m.fqnForWorkflow(workflow))
	f.Func().
		Params(j.Id("c").Op("*").Id(m.toCamel("Test%sClient", m.Service.GoName))).
		Id(methodName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if hasSignalInput {
				g.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			g.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
		}).
		ParamsFunc(func(g *j.Group) {
			if hasOutput {
				g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			g.Error()
		}).
		Block(
			j.Id("c").Dot("env").Dot("RegisterDelayedCallback").Call(
				j.Func().Params().Block(
					j.Id("c").Dot("env").Dot("SignalWorkflow").CallFunc(func(g *j.Group) {
						g.Qual(m.goImportPathForMethod(signal), m.toCamel("%sSignalName", signal))
						if hasSignalInput {
							g.Id("signal")
						} else {
							g.Nil()
						}
					}),
				),
				j.Lit(0),
			),
			j.Return(
				j.Id("c").Dot(m.methods[workflow].GoName).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					if hasInput {
						g.Id("req")
					}
					g.Id("opts").Op("...")
				}),
			),
		)
}

// genTestClientImplWorkflowWithSignalMethodAsync generates a TestClient's <workflow>With<signal>Async method
func (m *Manifest) genTestClientImplWorkflowWithSignalMethodAsync(f *j.File, workflow, signal protoreflect.FullName) {
	method := m.methods[workflow]
	methodName := m.toCamel("%sWith%sAsync", workflow, signal)
	hasInput := !isEmpty(method.Input)
	handler := m.methods[signal]
	hasSignalInput := !isEmpty(handler.Input)

	f.Commentf("%s sends a(n) %s signal to a(n) %s workflow, starting it if necessary", methodName, m.fqnForSignal(signal), m.fqnForWorkflow(workflow))
	f.Func().
		Params(j.Id("c").Op("*").Id(m.toCamel("Test%sClient", m.Service.GoName))).
		Id(methodName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if hasSignalInput {
				g.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			g.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
		}).
		Params(
			j.Id(m.toCamel("%sRun", workflow)),
			j.Error(),
		).
		Block(
			j.Id("c").Dot("env").Dot("RegisterDelayedCallback").Call(
				j.Func().Params().Block(
					j.Id("c").Dot("env").Dot("SignalWorkflow").CallFunc(func(g *j.Group) {
						g.Qual(m.goImportPathForMethod(signal), m.toCamel("%sSignalName", signal))
						if hasSignalInput {
							g.Id("signal")
						} else {
							g.Nil()
						}
					}),
				),
				j.Lit(0),
			),
			j.Return(
				j.Id("c").Dot(m.toCamel("%sAsync", workflow)).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					if hasInput {
						g.Id("req")
					}
					g.Id("opts").Op("...")
				}),
			),
		)
}

func (m *Manifest) genTestClientUpdateHandleImpl(f *j.File, update protoreflect.FullName) {
	typeName := m.toLowerCamel("test%sHandle", update)
	interfaceName := m.toCamel("%sHandle", update)
	handler := m.methods[update]
	hasInput := !isEmpty(handler.Input)

	// generate struct
	f.Var().Id("_").Id(interfaceName).Op("=").Op("&").Id(typeName).Values()
	f.Commentf("%s provides an internal implementation of a(n) %s", typeName, interfaceName)
	f.Type().
		Id(typeName).
		StructFunc(func(g *j.Group) {
			g.Id("callbacks").Op("*").Qual(testutilPkg, "UpdateCallbacks")
			g.Id("env").Op("*").Qual(testsuitePkg, "TestWorkflowEnvironment")
			g.Id("opts").Op("*").Qual(clientPkg, "UpdateWorkflowOptions")
			if hasInput {
				g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			g.Id("runID").String()
			g.Id("workflowID").String()
		})
}

func (m *Manifest) genTestClientUpdateHandleImplGetMethod(f *j.File, update protoreflect.FullName) {
	method := m.methods[update]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("Get retrieves a test %s update result", m.fqnForUpdate(update))
	f.Func().
		Params(j.Id("h").Op("*").Id(m.toLowerCamel("test%sHandle", update))).
		Id("Get").
		Params(j.Id("ctx").Qual("context", "Context")).
		ParamsFunc(func(g *j.Group) {
			if hasOutput {
				g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			g.Error()
		}).
		Block(
			j.If(j.ListFunc(func(g *j.Group) {
				if hasOutput {
					g.Id("resp")
				} else {
					g.Id("_")
				}
				g.Err()
			}).Op(":=").Id("h").Dot("callbacks").Dot("Get").Call(j.Id("ctx")), j.Err().Op("!=").Nil()).
				Block(
					j.ReturnFunc(func(g *j.Group) {
						if hasOutput {
							g.Nil()
						}
						g.Err()
					}),
				).
				Else().
				Block(
					j.ReturnFunc(func(g *j.Group) {
						if hasOutput {
							g.Id("resp").Op(".").Parens(j.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output)))
						}
						g.Nil()
					}),
				),
		)
}

func (m *Manifest) genTestClientUpdateHandleImplRunIDMethod(f *j.File, update protoreflect.FullName) {
	f.Comment("RunID implementation")
	f.Func().
		Params(j.Id("h").Op("*").Id(m.toLowerCamel("test%sHandle", update))).
		Id("RunID").
		Params().
		Params(j.String()).
		Block(
			j.Return(j.Id("h").Dot("runID")),
		)
}

func (m *Manifest) genTestClientUpdateHandleImplUpdateIDMethod(f *j.File, update protoreflect.FullName) {
	f.Comment("UpdateID implementation")
	f.Func().
		Params(j.Id("h").Op("*").Id(m.toLowerCamel("test%sHandle", update))).
		Id("UpdateID").
		Params().
		Params(j.String()).
		Block(
			j.If(j.Id("h").Dot("opts").Op("!=").Nil()).Block(
				j.Return(j.Id("h").Dot("opts").Dot("UpdateID")),
			),
			j.Return(j.Lit("")),
		)
}

func (m *Manifest) genTestClientUpdateHandleImplWorkflowIDMethod(f *j.File, update protoreflect.FullName) {
	f.Comment("WorkflowID implementation")
	f.Func().
		Params(j.Id("h").Op("*").Id(m.toLowerCamel("test%sHandle", update))).
		Id("WorkflowID").
		Params().
		Params(j.String()).
		Block(
			j.Return(j.Id("h").Dot("workflowID")),
		)
}

// genTestClientWorkflowRunImpl generates a test<Workflow>Run struct
func (m *Manifest) genTestClientWorkflowRunImpl(f *j.File, workflow protoreflect.FullName) {
	method := m.methods[workflow]
	methodName := m.toLowerCamel("test%sRun", workflow)
	hasInput := !isEmpty(method.Input)
	// generate test<Workflow>Run struct
	f.Var().Id("_").Id(m.toCamel("%sRun", workflow)).Op("=").Op("&").Id(m.toLowerCamel("test%sRun", workflow)).Values()
	f.Commentf("%s provides convenience methods for interacting with a(n) %s workflow in the test environment", methodName, m.fqnForWorkflow(workflow))
	f.Type().Id(methodName).StructFunc(func(g *j.Group) {
		g.Id("client").Op("*").Id(m.toCamel("Test%sClient", m.Service.GoName))
		g.Id("env").Op("*").Qual(testsuitePkg, "TestWorkflowEnvironment")
		g.Id("isStarted").Qual(atomicPkg, "Bool")
		g.Id("opts").Op("*").Qual(clientPkg, "StartWorkflowOptions")
		if hasInput {
			g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
		}
		g.Id("workflows").Id(m.toCamel("%sWorkflows", m.Service.GoName))
	})
}

// genClientWorkflowRunImplCancelMethod generates a <Workflow>Run's Cancel method
func (m *Manifest) genTestClientWorkflowRunImplCancelMethod(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toLowerCamel("Test%sRun", workflow)

	f.Comment("Cancel requests cancellation of a workflow in execution, returning an error if applicable")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("Cancel").
		Params(j.Id("ctx").Qual("context", "Context")).
		Params(j.Error()).
		Block(
			j.Return(
				j.Id("r").Dot("client").Dot("CancelWorkflow").CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("r").Dot("ID").Call()
					g.Id("r").Dot("RunID").Call()
				}),
			),
		)
}

// genTestClientWorkflowRunImplGetMethod generates a test<Workflow>Run's Get method
func (m *Manifest) genTestClientWorkflowRunImplGetMethod(f *j.File, workflow protoreflect.FullName) {
	method := m.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	f.Commentf("Get retrieves a test %s workflow result", m.fqnForWorkflow(workflow))
	f.Func().
		Params(j.Id("r").Op("*").Id(m.toLowerCamel("test%sRun", workflow))).
		Id("Get").
		Params(j.Qual("context", "Context")).
		ParamsFunc(func(g *j.Group) {
			if hasOutput {
				g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			// execute workflow
			g.IfFunc(func(g *j.Group) {
				g.Id("r").Dot("isStarted").Dot("CompareAndSwap").Call(j.False(), j.True())
			}).BlockFunc(func(g *j.Group) {
				g.Id("r").Dot("env").Dot("ExecuteWorkflow").CallFunc(func(g *j.Group) {
					g.Id(m.toCamel("%sWorkflowName", workflow))
					if hasInput {
						g.Id("r").Dot("req")
					}
				})
			})
			// ensure completed
			g.If(j.Op("!").Id("r").Dot("env").Dot("IsWorkflowCompleted").Call()).Block(
				j.ReturnFunc(func(g *j.Group) {
					if hasOutput {
						g.Nil()
					}
					g.Qual("errors", "New").Call(j.Lit("workflow in progress"))
				}),
			)
			// return workflow error if applicable
			g.If(j.Err().Op(":=").Id("r").Dot("env").Dot("GetWorkflowError").Call(), j.Err().Op("!=").Nil()).Block(
				j.ReturnFunc(func(g *j.Group) {
					if hasOutput {
						g.Nil()
					}
					g.Err()
				}),
			)
			// return workflow result
			if hasOutput {
				g.Var().Id("result").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
				g.If(j.Err().Op(":=").Id("r").Dot("env").Dot("GetWorkflowResult").Call(j.Op("&").Id("result")), j.Err().Op("!=").Nil()).Block(
					j.Return(j.Nil(), j.Err()),
				)
				g.Return(j.Op("&").Id("result"), j.Nil())
			} else {
				g.Return(j.Nil())
			}
		})
}

// genTestClientWorkflowRunImplIDMethod generates a test<Workflow>Run's workflow ID
func (m *Manifest) genTestClientWorkflowRunImplIDMethod(f *j.File, workflow protoreflect.FullName) {
	f.Commentf("ID returns a test %s workflow run's workflow ID", m.fqnForWorkflow(workflow))
	f.Func().
		Params(j.Id("r").Op("*").Id(m.toLowerCamel("test%sRun", workflow))).
		Id("ID").
		Params().
		Params(j.String()).
		Block(
			j.If(j.Id("r").Dot("opts").Op("!=").Nil()).Block(
				j.Return(j.Id("r").Dot("opts").Dot("ID")),
			),
			j.Return(j.Lit("")),
		)
}

// genTestClientWorkflowRunImplQueryMethod generates a test<Workflow>Run's <Query> method
func (m *Manifest) genTestClientWorkflowRunImplQueryMethod(f *j.File, workflow, query protoreflect.FullName) {
	handler := m.methods[query]
	handlerName := m.methods[query].GoName
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)

	f.Commentf("%s executes a %s query against a test %s workflow", handlerName, m.fqnForQuery(query), m.fqnForWorkflow(workflow))
	f.Func().
		Params(j.Id("r").Op("*").Id(m.toLowerCamel("test%sRun", workflow))).
		Id(handlerName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasInput {
				g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
		}).
		ParamsFunc(func(g *j.Group) {
			if hasOutput {
				g.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
			}
			g.Error()
		}).Block(
		j.Return(
			j.Id("r").Dot("client").Dot(m.methods[query].GoName).CallFunc(func(g *j.Group) {
				g.Id("ctx")
				g.Id("r").Dot("ID").Call()
				g.Id("r").Dot("RunID").Call()
				if hasInput {
					g.Id("req")
				}
			}),
		),
	)
}

// genTestClientWorkflowRunImplRunMethod generates a test<Workflow>Run's Run method
func (m *Manifest) genTestClientWorkflowRunImplRunMethod(f *j.File, workflow protoreflect.FullName) {
	f.Comment("Run noop implementation")
	f.Func().
		Params(j.Id("r").Op("*").Id(m.toLowerCamel("test%sRun", workflow))).
		Id("Run").
		Params().
		Qual(clientPkg, "WorkflowRun").
		Block(
			j.Return(j.Nil()),
		)
}

// genTestClientWorkflowRunImplRunIDMethod generates a test<Workflow>Run's RunID method
func (m *Manifest) genTestClientWorkflowRunImplRunIDMethod(f *j.File, workflow protoreflect.FullName) {
	f.Comment("RunID noop implementation")
	f.Func().
		Params(j.Id("r").Op("*").Id(m.toLowerCamel("test%sRun", workflow))).
		Id("RunID").
		Params().
		Params(j.String()).
		Block(
			j.Return(j.Lit("")),
		)
}

// genTestClientWorkflowRunImplQueryMethod generates a test<Workflow>Run's <Signal> method
func (m *Manifest) genTestClientWorkflowRunImplSignalMethod(f *j.File, workflow, signal protoreflect.FullName) {
	handler := m.methods[signal]
	handlerName := m.methods[signal].GoName
	hasInput := !isEmpty(handler.Input)

	f.Commentf("%s executes a %s signal against a test %s workflow", handlerName, m.fqnForSignal(signal), m.fqnForWorkflow(workflow))
	f.Func().
		Params(j.Id("r").Op("*").Id(m.toLowerCamel("test%sRun", workflow))).
		Id(handlerName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasInput {
				g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
		}).
		Params(
			j.Error(),
		).
		BlockFunc(func(g *j.Group) {
			if m.methodsFromSameService(signal, workflow) {
				g.Return().Id("r").Dot("client").Dot(handlerName).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("r").Dot("ID").Call()
					g.Id("r").Dot("RunID").Call()
					if hasInput {
						g.Id("req")
					}
				})
			} else {
				g.Id("r").Dot("env").Dot("SignalWorkflow").CallFunc(func(g *j.Group) {
					g.Add(m.Qual(signal, m.toCamel("%sSignalName", signal)))
					if hasInput {
						g.Id("req")
					} else {
						g.Nil()
					}
				})
				g.Return(j.Nil())
			}
		})
}

// genClientWorkflowRunImplTerminateMethod generates a <Workflow>Run's Terminate method
func (m *Manifest) genTestClientWorkflowRunImplTerminateMethod(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toLowerCamel("Test%sRun", workflow)

	f.Comment("Terminate terminates a workflow in execution, returning an error if applicable")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("Terminate").
		Params(
			j.Id("ctx").Qual("context", "Context"),
			j.Id("reason").String(),
			j.Id("details").Op("...").Interface(),
		).
		Params(j.Error()).
		Block(
			j.Return(
				j.Id("r").Dot("client").Dot("TerminateWorkflow").CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("r").Dot("ID").Call()
					g.Id("r").Dot("RunID").Call()
					g.Id("reason")
					g.Id("details").Op("...")
				}),
			),
		)
}

// genTestClientWorkflowRunImplQueryMethod generates a test<Workflow>Run's <Update> method
func (m *Manifest) genTestClientWorkflowRunImplUpdateMethod(f *j.File, workflow, update protoreflect.FullName) {
	handler := m.methods[update]
	handlerName := m.methods[update].GoName
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)

	f.Commentf("%s executes a(n) %s update against a test %s workflow", handlerName, m.fqnForUpdate(update), m.fqnForWorkflow(workflow))
	f.Func().
		Params(j.Id("r").Op("*").Id(m.toLowerCamel("test%sRun", workflow))).
		Id(handlerName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasInput {
				g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			g.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", update))
		}).
		ParamsFunc(func(g *j.Group) {
			if hasOutput {
				g.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
			}
			g.Error()
		}).
		Block(
			j.Return(
				j.Id("r").Dot("client").Dot(m.methods[update].GoName).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("r").Dot("ID").Call()
					g.Id("r").Dot("RunID").Call()
					if hasInput {
						g.Id("req")
					}
					g.Id("opts").Op("...")
				}),
			),
		)
}

// genTestClientWorkflowRunImplQueryMethod generates a test<Workflow>Run's <Update>Async method
func (m *Manifest) genTestClientWorkflowRunImplUpdateMethodAsync(f *j.File, workflow, update protoreflect.FullName) {
	handler := m.methods[update]
	hasInput := !isEmpty(handler.Input)
	methodName := m.toCamel("%sAsync", update)

	f.Commentf("%s executes a(n) %s update against a test %s workflow", methodName, m.fqnForUpdate(update), m.fqnForWorkflow(workflow))
	f.Func().
		Params(j.Id("r").Op("*").Id(m.toLowerCamel("test%sRun", workflow))).
		Id(methodName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasInput {
				g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			g.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", update))
		}).
		Params(
			j.Id(m.toCamel("%sHandle", update)),
			j.Error(),
		).
		Block(
			j.Return(
				j.Id("r").Dot("client").Dot(methodName).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("r").Dot("ID").Call()
					g.Id("r").Dot("RunID").Call()
					if hasInput {
						g.Id("req")
					}
					g.Id("opts").Op("...")
				}),
			),
		)
}

func (m *Manifest) renderTestClient(f *j.File) {
	// generate test client
	m.genTestClientImpl(f)
	m.genTestClientImplNewMethod(f)
	for _, workflow := range m.workflowsOrdered {
		if m.methods[workflow].Desc.Parent() != m.Service.Desc {
			continue
		}
		m.genTestClientImplWorkflowMethod(f, workflow)
		m.genTestClientImplWorkflowMethodAsync(f, workflow)
		m.genTestClientImplWorkflowGetMethod(f, workflow)
		for _, signal := range m.workflows[workflow].GetSignal() {
			if !signal.GetStart() {
				continue
			}
			m.genTestClientImplWorkflowWithSignalMethod(f, workflow, getFullyQualifiedRef(workflow, signal.GetRef()))
			m.genTestClientImplWorkflowWithSignalMethodAsync(f, workflow, getFullyQualifiedRef(workflow, signal.GetRef()))
		}
		for _, update := range m.workflows[workflow].GetUpdate() {
			if !update.GetStart() {
				continue
			}
			m.genTestClientImplUpdateWithStartMethod(f, workflow, getFullyQualifiedRef(workflow, update.GetRef()))
			m.genTestClientImplUpdateWithStartMethodAsync(f, workflow, getFullyQualifiedRef(workflow, update.GetRef()))
		}
	}

	m.genTestClientImplWorkflowCancelMethod(f)
	m.genTestClientImplWorkflowTerminateMethod(f)

	// generate test client query methods
	for _, query := range m.queriesOrdered {
		if m.methods[query].Desc.Parent() != m.Service.Desc {
			continue
		}
		m.genTestClientImplQueryMethod(f, query)
	}

	// generate test client signal methods
	for _, signal := range m.signalsOrdered {
		if m.methods[signal].Desc.Parent() != m.Service.Desc {
			continue
		}
		m.genTestClientImplSignalMethod(f, signal)
	}

	// generate test client update methods
	for _, update := range m.updatesOrdered {
		if m.methods[update].Desc.Parent() != m.Service.Desc {
			continue
		}
		m.genTestClientImplUpdateMethod(f, update)
		m.genTestClientImplUpdateMethodAsync(f, update)
		m.genTestClientImplUpdateGetMethod(f, update)

		m.genTestClientUpdateHandleImpl(f, update)
		m.genTestClientUpdateHandleImplGetMethod(f, update)
		m.genTestClientUpdateHandleImplRunIDMethod(f, update)
		m.genTestClientUpdateHandleImplUpdateIDMethod(f, update)
		m.genTestClientUpdateHandleImplWorkflowIDMethod(f, update)
	}

	// generate workflow test runs
	for _, workflow := range m.workflowsOrdered {
		if m.methods[workflow].Desc.Parent() != m.Service.Desc {
			continue
		}
		opts := m.workflows[workflow]
		m.genTestClientWorkflowRunImpl(f, workflow)
		m.genTestClientWorkflowRunImplCancelMethod(f, workflow)
		m.genTestClientWorkflowRunImplGetMethod(f, workflow)
		m.genTestClientWorkflowRunImplIDMethod(f, workflow)
		m.genTestClientWorkflowRunImplRunMethod(f, workflow)
		m.genTestClientWorkflowRunImplRunIDMethod(f, workflow)
		m.genTestClientWorkflowRunImplTerminateMethod(f, workflow)

		// generate query methods
		for _, queryOpts := range opts.GetQuery() {
			m.genTestClientWorkflowRunImplQueryMethod(f, workflow, getFullyQualifiedRef(workflow, queryOpts.GetRef()))
		}

		// generate signal methods
		for _, signalOpts := range opts.GetSignal() {
			m.genTestClientWorkflowRunImplSignalMethod(f, workflow, getFullyQualifiedRef(workflow, signalOpts.GetRef()))
		}

		// generate update methods
		for _, updateOpts := range opts.GetUpdate() {
			m.genTestClientWorkflowRunImplUpdateMethod(f, workflow, getFullyQualifiedRef(workflow, updateOpts.GetRef()))
			m.genTestClientWorkflowRunImplUpdateMethodAsync(f, workflow, getFullyQualifiedRef(workflow, updateOpts.GetRef()))
		}
	}
}
