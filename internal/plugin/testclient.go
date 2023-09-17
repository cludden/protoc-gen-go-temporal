package plugin

import (
	"fmt"

	g "github.com/dave/jennifer/jen"
)

const (
	testsuitePkg = "go.temporal.io/sdk/testsuite"
	testutilPkg  = "github.com/cludden/protoc-gen-go-temporal/pkg/testutil"
)

// genTestClientImpl generates a TestClient struct
func (svc *Service) genTestClientImpl(f *g.File) {
	f.Comment("TestClient provides a testsuite-compatible Client")
	f.Type().Id(toCamel("Test%sClient", svc.Service.GoName)).Struct(
		g.Id("env").Op("*").Qual(testsuitePkg, "TestWorkflowEnvironment"),
		g.Id("workflows").Id(toCamel("%sWorkflows", svc.Service.GoName)),
	)
}

// genTestClientImplNewMethod generates a NewTestClient constructor function
func (svc *Service) genTestClientImplNewMethod(f *g.File) {
	interfaceName := toCamel("%sClient", svc.Service.GoName)
	typeName := toCamel("Test%sClient", svc.Service.GoName)
	functionName := "New" + typeName

	f.Var().Id("_").Id(interfaceName).Op("=").Op("&").Id(typeName).Values()
	f.Commentf("%s initializes a new %s value", functionName, typeName)
	f.Func().Id(functionName).
		Params(
			g.Id("env").Op("*").Qual(testsuitePkg, "TestWorkflowEnvironment"),
			g.Id("workflows").Id(toCamel("%sWorkflows", svc.Service.GoName)),
			g.Id("activities").Id(toCamel("%sActivities", svc.Service.GoName)),
		).
		Params(
			g.Op("*").Id(typeName),
		).
		Block(
			g.Id(toCamel("Register%sWorkflows", svc.Service.GoName)).Call(g.Id("env"), g.Id("workflows")),
			g.If(g.Id("activities").Op("!=").Nil()).Block(
				g.Id(toCamel("Register%sActivities", svc.Service.GoName)).Call(g.Id("env"), g.Id("activities")),
			),
			g.Return(g.Op("&").Id(typeName).Values(g.Id("env"), g.Id("workflows"))),
		)
}

// genTestClientImplQueryMethod genereates a TestClient <Query> method
func (svc *Service) genTestClientImplQueryMethod(f *g.File, query string) {
	handler := svc.methods[query]
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)

	f.Commentf("%s executes a %s query", query, svc.fqnForQuery(query))
	f.Func().
		Params(g.Id("c").Op("*").Id(toCamel("Test%sClient", svc.Service.GoName))).
		Id(query).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
			}
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(handler.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			fn.List(g.Id("val"), g.Err()).Op(":=").Id("c").Dot("env").Dot("QueryWorkflow").CallFunc(func(args *g.Group) {
				args.Id(fmt.Sprintf("%sQueryName", query))
				if hasInput {
					args.Id("req")
				}
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			).Else().If(g.Op("!").Id("val").Dot("HasValue").Call()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Nil()
				}),
			).Else().BlockFunc(func(bl *g.Group) {
				if !hasOutput {
					bl.Return(g.Nil())
				} else {
					bl.Var().Id("result").Id(handler.Output.GoIdent.GoName)
					bl.If(g.Err().Op(":=").Id("val").Dot("Get").Call(g.Op("&").Id("result")), g.Err().Op("!=").Nil()).Block(
						g.Return(
							g.Nil(),
							g.Err(),
						),
					)
					bl.Return(g.Op("&").Id("result"), g.Nil())
				}
			})
		})
}

// genTestClientImplSignalMethod genereates a TestClient <Signal> method
func (svc *Service) genTestClientImplSignalMethod(f *g.File, signal string) {
	handler := svc.methods[signal]
	hasInput := !isEmpty(handler.Input)

	f.Commentf("%s executes a %s signal", signal, signal)
	f.Func().
		Params(g.Id("c").Op("*").Id(toCamel("Test%sClient", svc.Service.GoName))).
		Id(signal).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
			}
		}).
		Params(
			g.Error(),
		).
		Block(
			g.Id("c").Dot("env").Dot("SignalWorkflow").CallFunc(func(args *g.Group) {
				args.Id(fmt.Sprintf("%sSignalName", signal))
				if hasInput {
					args.Id("req")
				} else {
					args.Nil()
				}
			}),
			g.Return(g.Nil()),
		)
}

// genTestClientImplUpdateMethod genereates a TestClient <Update> method
func (svc *Service) genTestClientImplUpdateMethod(f *g.File, update string) {
	method := svc.methods[update]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	asyncName := toCamel("%sAsync", update)

	f.Commentf("%s executes a(n) %s update in the test environment", update, svc.fqnForUpdate(update))
	f.Func().
		Params(g.Id("c").Op("*").Id(toCamel("Test%sClient", svc.Service.GoName))).
		Id(update).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sOptions", update))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(method.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		Block(
			// initialize update request options
			g.Id("options").Op(":=").Id(toCamel("New%sOptions", update)).Call(),
			g.If(g.Len(g.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(g.Lit(0)).Dot("opts").Op("!=").Nil()).Block(
				g.Id("options").Op("=").Id("opts").Index(g.Lit(0)),
			),

			// override wait policy
			g.Id("options").Dot("opts").Dot("WaitPolicy").Op("=").Op("&").Qual(updatePkg, "WaitPolicy").Values(
				g.Id("LifecycleStage").Op(":").Qual(enumsPkg, "UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_COMPLETED"),
			),
			g.List(g.Id("handle"), g.Err()).Op(":=").Id("c").Dot(asyncName).CallFunc(func(args *g.Group) {
				args.Id("ctx")
				args.Id("workflowID")
				args.Id("runID")
				if hasInput {
					args.Id("req")
				}
				args.Id("options")
			}),
			g.If(g.Err().Op("!=").Nil()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			),
			g.Return(g.Id("handle").Dot("Get").Call(g.Id("ctx"))),
		)
}

// genTestClientImplUpdateAsyncMethod genereates a TestClient <UpdateAsync> method
func (svc *Service) genTestClientImplUpdateAsyncMethod(f *g.File, update string) {
	method := svc.methods[update]
	hasInput := !isEmpty(method.Input)
	methodName := toCamel("%sAsync", update)

	f.Commentf("%s executes a(n) %s update in the test environment", methodName, svc.fqnForUpdate(update))
	f.Func().
		Params(g.Id("c").Op("*").Id(toCamel("Test%sClient", svc.Service.GoName))).
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sOptions", update))
		}).
		Params(
			g.Id(toCamel("%sHandle", update)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			// generate UpdateWorkflowWithOptionsRequest with defaults
			svc.genClientUpdateWorkflowOptions(fn, update)

			// update workflow
			fn.Id("uc").Op(":=").Qual(testutilPkg, "NewUpdateCallbacks").Call()
			fn.Id("c").Dot("env").Dot("UpdateWorkflow").CallFunc(func(args *g.Group) {
				args.Id(toCamel("%sUpdateName", update))
				args.Id("workflowID")
				args.Id("uc")
				if hasInput {
					args.Id("req")
				}
			})

			fn.Return(
				g.Op("&").Id(fmt.Sprintf("test%sHandle", update)).CustomFunc(multiLineValues, func(fields *g.Group) {
					fields.Id("callbacks").Op(":").Id("uc")
					fields.Id("env").Op(":").Id("c").Dot("env")
					fields.Id("opts").Op(":").Id("options")
					fields.Id("runID").Op(":").Id("runID")
					fields.Id("workflowID").Op(":").Id("workflowID")
					if hasInput {
						fields.Id("req").Op(":").Id("req")
					}
				}),
				g.Nil(),
			)
		})
}

// genTestClientImplWorkflowMethod generates a TestClient <Workflow> method
func (svc *Service) genTestClientImplWorkflowMethod(f *g.File, workflow string) {
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s executes a(n) %s workflow in the test environment", workflow, workflow)
	f.Func().
		Params(g.Id("c").Op("*").Id(toCamel("Test%sClient", svc.Service.GoName))).
		Id(workflow).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sOptions", workflow))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(method.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		Block(
			g.List(g.Id("run"), g.Err()).Op(":=").Id("c").Dot(fmt.Sprintf("%sAsync", workflow)).CallFunc(func(args *g.Group) {
				args.Id("ctx")
				if hasInput {
					args.Id("req")
				}
				args.Id("opts").Op("...")
			}),
			g.If(g.Err().Op("!=").Nil()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			),
			g.Return(g.Id("run").Dot("Get").Call(g.Id("ctx"))),
		)
}

// genTestClientImplWorkflowAsyncMethod generates a TestClient's <workflow>Async method
func (svc *Service) genTestClientImplWorkflowAsyncMethod(f *g.File, workflow string) {
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)

	f.Commentf("%sAsync executes a(n) %s workflow in the test environment", workflow, workflow)
	f.Func().
		Params(g.Id("c").Op("*").Id(toCamel("Test%sClient", svc.Service.GoName))).
		Id(fmt.Sprintf("%sAsync", workflow)).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
			args.Id("options").Op("...").Op("*").Id(toCamel("%sOptions", workflow))
		}).
		Params(
			g.Id(fmt.Sprintf("%sRun", workflow)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			svc.genClientStartWorkflowOptions(fn, workflow, false)
			fn.Return(
				g.Op("&").Id(fmt.Sprintf("test%sRun", workflow)).ValuesFunc(func(fields *g.Group) {
					fields.Id("client").Op(":").Id("c")
					fields.Id("env").Op(":").Id("c").Dot("env")
					fields.Id("opts").Op(":").Id("opts")
					if hasInput {
						fields.Id("req").Op(":").Id("req")
					}
					fields.Id("workflows").Op(":").Id("c").Dot("workflows")
				}),
				g.Nil(),
			)
		})
}

// genClientImplWorkflowCancelMethod generates a Cancel<Workflow> client method
func (svc *Service) genTestClientImplWorkflowCancelMethod(f *g.File) {
	clientType := toCamel("Test%sClient", svc.Service.GoName)
	methodName := "CancelWorkflow"

	f.Commentf("%s requests cancellation of an existing workflow execution", methodName)
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(methodName).
		Params(
			g.Id("ctx").Qual("context", "Context"),
			g.Id("workflowID").String(),
			g.Id("runID").String(),
		).
		Params(
			g.Error(),
		).
		Block(
			g.Id("c").Dot("env").Dot("CancelWorkflow").Call(),
			g.Return(g.Nil()),
		)
}

// genTestClientImplWorkflowGetMethod generates a TestClient's Get<workflow> method
func (svc *Service) genTestClientImplWorkflowGetMethod(f *g.File, workflow string) {
	f.Commentf("Get%s is a noop", workflow)
	f.Func().
		Params(g.Id("c").Op("*").Id(toCamel("Test%sClient", svc.Service.GoName))).
		Id(fmt.Sprintf("Get%s", workflow)).
		Params(
			g.Id("ctx").Qual("context", "Context"),
			g.Id("workflowID").String(),
			g.Id("runID").String(),
		).
		Params(
			g.Id(fmt.Sprintf("%sRun", workflow)),
		).
		Block(
			g.Return(
				g.Op("&").Id(fmt.Sprintf("test%sRun", workflow)).Values(
					g.Id("env").Op(":").Id("c").Dot("env"),
					g.Id("workflows").Op(":").Id("c").Dot("workflows"),
				),
			),
		)
}

// genTestClientImplWorkflowWithSignalMethod generates a TestClient's <workflow>With<signal> method
func (svc *Service) genTestClientImplWorkflowWithSignalMethod(f *g.File, workflow, signal string) {
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	handler := svc.methods[signal]
	hasSignalInput := !isEmpty(handler.Input)

	f.Commentf("%sWith%s sends a(n) %s signal to a(n) %s workflow, starting it if necessary", workflow, signal, signal, workflow)
	f.Func().
		Params(g.Id("c").Op("*").Id(toCamel("Test%sClient", svc.Service.GoName))).
		Id(fmt.Sprintf("%sWith%s", workflow, signal)).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
			if hasSignalInput {
				args.Id("signal").Op("*").Id(handler.Input.GoIdent.GoName)
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sOptions", workflow))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(method.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		Block(
			g.Id("c").Dot("env").Dot("RegisterDelayedCallback").Call(
				g.Func().Params().Block(
					g.Id("c").Dot("env").Dot("SignalWorkflow").CallFunc(func(args *g.Group) {
						args.Id(fmt.Sprintf("%sSignalName", signal))
						if hasSignalInput {
							args.Id("signal")
						} else {
							args.Nil()
						}
					}),
				),
				g.Lit(0),
			),
			g.Return(
				g.Id("c").Dot(workflow).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					if hasInput {
						args.Id("req")
					}
					args.Id("opts").Op("...")
				}),
			),
		)
}

// genTestClientImplWorkflowWithSignalAsyncMethod generates a TestClient's <workflow>With<signal>Async method
func (svc *Service) genTestClientImplWorkflowWithSignalAsyncMethod(f *g.File, workflow, signal string) {
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	handler := svc.methods[signal]
	hasSignalInput := !isEmpty(handler.Input)

	f.Commentf("%sWith%sAsync sends a(n) %s signal to a(n) %s workflow, starting it if necessary", workflow, signal, signal, workflow)
	f.Func().
		Params(g.Id("c").Op("*").Id(toCamel("Test%sClient", svc.Service.GoName))).
		Id(fmt.Sprintf("%sWith%sAsync", workflow, signal)).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
			if hasSignalInput {
				args.Id("signal").Op("*").Id(handler.Input.GoIdent.GoName)
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sOptions", workflow))
		}).
		Params(
			g.Id(toCamel("%sRun", workflow)),
			g.Error(),
		).
		Block(
			g.Id("c").Dot("env").Dot("RegisterDelayedCallback").Call(
				g.Func().Params().Block(
					g.Id("_").Op("=").Id("c").Dot(signal).CallFunc(func(args *g.Group) {
						args.Id("ctx")
						args.Lit("")
						args.Lit("")
						if hasInput {
							args.Id("signal")
						}
					}),
				),
				g.Lit(0),
			),
			g.Return(
				g.Id("c").Dot(fmt.Sprintf("%sAsync", workflow)).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					if hasInput {
						args.Id("req")
					}
					args.Id("opts").Op("...")
				}),
			),
		)
}

// genClientImplWorkflowTerminateMethod generates a Terminate<Workflow> client method
func (svc *Service) genTestClientImplWorkflowTerminateMethod(f *g.File) {
	clientType := toCamel("Test%sClient", svc.Service.GoName)
	methodName := "TerminateWorkflow"

	f.Commentf("%s terminates an existing workflow execution", methodName)
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(methodName).
		Params(
			g.Id("ctx").Qual("context", "Context"),
			g.Id("workflowID").String(),
			g.Id("runID").String(),
			g.Id("reason").String(),
			g.Id("details").Op("...").Interface(),
		).
		Params(
			g.Error(),
		).
		Block(
			g.Return(
				g.Id("c").Dot("CancelWorkflow").Call(
					g.Id("ctx"),
					g.Id("workflowID"),
					g.Id("runID"),
				),
			),
		)
}

func (svc *Service) genTestClientUpdateHandleImpl(f *g.File, update string) {
	typeName := toLowerCamel("test%sHandle", update)
	interfaceName := toCamel("%sHandle", update)
	handler := svc.methods[update]
	hasInput := !isEmpty(handler.Input)

	// generate struct
	f.Var().Id("_").Id(interfaceName).Op("=").Op("&").Id(typeName).Values()
	f.Commentf("%s provides an internal implementation of a(n) %s", typeName, interfaceName)
	f.Type().
		Id(typeName).
		StructFunc(func(fields *g.Group) {
			fields.Id("callbacks").Op("*").Qual(testutilPkg, "UpdateCallbacks")
			fields.Id("env").Op("*").Qual(testsuitePkg, "TestWorkflowEnvironment")
			fields.Id("opts").Op("*").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest")
			if hasInput {
				fields.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
			}
			fields.Id("runID").String()
			fields.Id("workflowID").String()
		})
}

func (svc *Service) genTestClientUpdateHandleImplGetMethod(f *g.File, update string) {
	method := svc.methods[update]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("Get retrieves a test %s update result", svc.fqnForUpdate(update))
	f.Func().
		Params(g.Id("h").Op("*").Id(toLowerCamel("test%sHandle", update))).
		Id("Get").
		Params(g.Id("ctx").Qual("context", "Context")).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(method.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		Block(
			g.If(g.List(g.Id("resp"), g.Err()).Op(":=").Id("h").Dot("callbacks").Dot("Get").Call(g.Id("ctx")), g.Err().Op("!=").Nil()).
				Block(
					g.ReturnFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Err()
					}),
				).
				Else().
				Block(
					g.ReturnFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Id("resp").Op(".").Parens(g.Op("*").Id(method.Output.GoIdent.GoName))
						}
						returnVals.Nil()
					}),
				),
		)
}

func (svc *Service) genTestClientUpdateHandleImplRunIDMethod(f *g.File, update string) {
	f.Comment("RunID implementation")
	f.Func().
		Params(g.Id("h").Op("*").Id(fmt.Sprintf("test%sHandle", update))).
		Id("RunID").
		Params().
		Params(g.String()).
		Block(
			g.Return(g.Id("h").Dot("runID")),
		)
}

func (svc *Service) genTestClientUpdateHandleImplUpdateIDMethod(f *g.File, update string) {
	f.Comment("UpdateID implementation")
	f.Func().
		Params(g.Id("h").Op("*").Id(fmt.Sprintf("test%sHandle", update))).
		Id("UpdateID").
		Params().
		Params(g.String()).
		Block(
			g.If(g.Id("h").Dot("opts").Op("!=").Nil()).Block(
				g.Return(g.Id("h").Dot("opts").Dot("UpdateID")),
			),
			g.Return(g.Lit("")),
		)
}

func (svc *Service) genTestClientUpdateHandleImplWorkflowIDMethod(f *g.File, update string) {
	f.Comment("WorkflowID implementation")
	f.Func().
		Params(g.Id("h").Op("*").Id(fmt.Sprintf("test%sHandle", update))).
		Id("WorkflowID").
		Params().
		Params(g.String()).
		Block(
			g.Return(g.Id("h").Dot("workflowID")),
		)
}

// genTestClientWorkflowGetMethod generates a noop TestClient Get<Workflow> method
func (svc *Service) genTestClientWorkflowGetMethod(f *g.File, workflow string) {
	f.Commentf("Get%s retrieves a test %sRun", workflow, workflow)
	f.Func().
		Params(g.Id("c").Op("*").Id(toCamel("Test%sClient", svc.Service.GoName))).
		Params(
			g.Qual("context", "Context"),
			g.String(),
			g.String(),
		).
		Params(
			g.Id(fmt.Sprintf("%sRun", workflow)),
			g.Error(),
		).
		Block(
			g.Return(g.Op("&").Id(fmt.Sprintf("test%sRun", workflow)).Values(g.Id("env").Op(":").Id("c").Dot("env"))),
		)
}

// genTestClientWorkflowRunImpl generates a test<Workflow>Run struct
func (svc *Service) genTestClientWorkflowRunImpl(f *g.File, workflow string) {
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	// generate test<Workflow>Run struct
	f.Var().Id("_").Id(fmt.Sprintf("%sRun", workflow)).Op("=").Op("&").Id(fmt.Sprintf("test%sRun", workflow)).Values()
	f.Commentf("test%sRun provides convenience methods for interacting with a(n) %s workflow in the test environment", workflow, workflow)
	f.Type().Id(fmt.Sprintf("test%sRun", workflow)).StructFunc(func(fields *g.Group) {
		fields.Id("client").Op("*").Id(toCamel("Test%sClient", svc.Service.GoName))
		fields.Id("env").Op("*").Qual(testsuitePkg, "TestWorkflowEnvironment")
		fields.Id("opts").Op("*").Qual(clientPkg, "StartWorkflowOptions")
		if hasInput {
			fields.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
		}
		fields.Id("workflows").Id(toCamel("%sWorkflows", svc.Service.GoName))
	})
}

// genClientWorkflowRunImplCancelMethod generates a <Workflow>Run's Cancel method
func (svc *Service) genTestClientWorkflowRunImplCancelMethod(f *g.File, workflow string) {
	typeName := toLowerCamel("Test%sRun", workflow)

	f.Comment("Cancel requests cancellation of a workflow in execution, returning an error if applicable")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Cancel").
		Params(g.Id("ctx").Qual("context", "Context")).
		Params(g.Error()).
		Block(
			g.Return(
				g.Id("r").Dot("client").Dot("CancelWorkflow").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("r").Dot("ID").Call()
					args.Id("r").Dot("RunID").Call()
				}),
			),
		)
}

// genTestClientWorkflowRunImplGetMethod generates a test<Workflow>Run's Get method
func (svc *Service) genTestClientWorkflowRunImplGetMethod(f *g.File, workflow string) {
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	f.Commentf("Get retrieves a test %s workflow result", workflow)
	f.Func().
		Params(g.Id("r").Op("*").Id(fmt.Sprintf("test%sRun", workflow))).
		Id("Get").
		Params(g.Qual("context", "Context")).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(method.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			// execute workflow
			fn.Id("r").Dot("env").Dot("ExecuteWorkflow").CallFunc(func(args *g.Group) {
				args.Id(toCamel("%sWorkflowName", workflow))
				if hasInput {
					args.Id("r").Dot("req")
				}
			})
			// ensure completed
			fn.If(g.Op("!").Id("r").Dot("env").Dot("IsWorkflowCompleted").Call()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Qual("errors", "New").Call(g.Lit("workflow in progress"))
				}),
			)
			// return workflow error if applicable
			fn.If(g.Err().Op(":=").Id("r").Dot("env").Dot("GetWorkflowError").Call(), g.Err().Op("!=").Nil()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)
			// return workflow result
			if hasOutput {
				fn.Var().Id("result").Id(method.Output.GoIdent.GoName)
				fn.If(g.Err().Op(":=").Id("r").Dot("env").Dot("GetWorkflowResult").Call(g.Op("&").Id("result")), g.Err().Op("!=").Nil()).Block(
					g.Return(g.Nil(), g.Err()),
				)
				fn.Return(g.Op("&").Id("result"), g.Nil())
			} else {
				fn.Return(g.Nil())
			}
		})
}

// genTestClientWorkflowRunImplIDMethod generates a test<Workflow>Run's workflow ID
func (svc *Service) genTestClientWorkflowRunImplIDMethod(f *g.File, workflow string) {
	f.Commentf("ID returns a test %s workflow run's workflow ID", workflow)
	f.Func().
		Params(g.Id("r").Op("*").Id(fmt.Sprintf("test%sRun", workflow))).
		Id("ID").
		Params().
		Params(g.String()).
		Block(
			g.If(g.Id("r").Dot("opts").Op("!=").Nil()).Block(
				g.Return(g.Id("r").Dot("opts").Dot("ID")),
			),
			g.Return(g.Lit("")),
		)
}

// genTestClientWorkflowRunImplQueryMethod generates a test<Workflow>Run's <Query> method
func (svc *Service) genTestClientWorkflowRunImplQueryMethod(f *g.File, workflow, query string) {
	handler := svc.methods[query]
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)

	f.Commentf("%s executes a %s query against a test %s workflow", query, query, workflow)
	f.Func().
		Params(g.Id("r").Op("*").Id(fmt.Sprintf("test%sRun", workflow))).
		Id(query).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
			}
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(handler.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).Block(
		g.Return(
			g.Id("r").Dot("client").Dot(query).CallFunc(func(args *g.Group) {
				args.Id("ctx")
				args.Id("r").Dot("ID").Call()
				args.Id("r").Dot("RunID").Call()
				if hasInput {
					args.Id("req")
				}
			}),
		),
	)
}

// genTestClientWorkflowRunImplRunIDMethod generates a test<Workflow>Run's RunID method
func (svc *Service) genTestClientWorkflowRunImplRunIDMethod(f *g.File, workflow string) {
	f.Comment("RunID noop implementation")
	f.Func().
		Params(g.Id("r").Op("*").Id(fmt.Sprintf("test%sRun", workflow))).
		Id("RunID").
		Params().
		Params(g.String()).
		Block(
			g.Return(g.Lit("")),
		)
}

// genTestClientWorkflowRunImplQueryMethod generates a test<Workflow>Run's <Signal> method
func (svc *Service) genTestClientWorkflowRunImplSignalMethod(f *g.File, workflow, signal string) {
	handler := svc.methods[signal]
	hasInput := !isEmpty(handler.Input)

	f.Commentf("%s executes a %s signal against a test %s workflow", signal, signal, workflow)
	f.Func().
		Params(g.Id("r").Op("*").Id(fmt.Sprintf("test%sRun", workflow))).
		Id(signal).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
			}
		}).
		Params(
			g.Error(),
		).
		Block(
			g.Return(
				g.Id("r").Dot("client").Dot(signal).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("r").Dot("ID").Call()
					args.Id("r").Dot("RunID").Call()
					if hasInput {
						args.Id("req")
					}
				}),
			),
		)
}

// genClientWorkflowRunImplTerminateMethod generates a <Workflow>Run's Terminate method
func (svc *Service) genTestClientWorkflowRunImplTerminateMethod(f *g.File, workflow string) {
	typeName := toLowerCamel("Test%sRun", workflow)

	f.Comment("Terminate terminates a workflow in execution, returning an error if applicable")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Terminate").
		Params(
			g.Id("ctx").Qual("context", "Context"),
			g.Id("reason").String(),
			g.Id("details").Op("...").Interface(),
		).
		Params(g.Error()).
		Block(
			g.Return(
				g.Id("r").Dot("client").Dot("TerminateWorkflow").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("r").Dot("ID").Call()
					args.Id("r").Dot("RunID").Call()
					args.Id("reason")
					args.Id("details").Op("...")
				}),
			),
		)
}

// genTestClientWorkflowRunImplQueryMethod generates a test<Workflow>Run's <Update> method
func (svc *Service) genTestClientWorkflowRunImplUpdateMethod(f *g.File, workflow, update string) {
	handler := svc.methods[update]
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)

	f.Commentf("%s executes a(n) %s update against a test %s workflow", update, svc.fqnForUpdate(update), svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("r").Op("*").Id(fmt.Sprintf("test%sRun", workflow))).
		Id(update).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sOptions", update))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(handler.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		Block(
			g.Return(
				g.Id("r").Dot("client").Dot(update).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("r").Dot("ID").Call()
					args.Id("r").Dot("RunID").Call()
					if hasInput {
						args.Id("req")
					}
					args.Id("opts").Op("...")
				}),
			),
		)
}

// genTestClientWorkflowRunImplQueryMethod generates a test<Workflow>Run's <Update>Async method
func (svc *Service) genTestClientWorkflowRunImplUpdateAsyncMethod(f *g.File, workflow, update string) {
	handler := svc.methods[update]
	hasInput := !isEmpty(handler.Input)
	methodName := toCamel("%sAsync", update)

	f.Commentf("%s executes a(n) %s update against a test %s workflow", methodName, svc.fqnForUpdate(update), svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("r").Op("*").Id(fmt.Sprintf("test%sRun", workflow))).
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sOptions", update))
		}).
		Params(
			g.Id(toCamel("%sHandle", update)),
			g.Error(),
		).
		Block(
			g.Return(
				g.Id("r").Dot("client").Dot(methodName).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("r").Dot("ID").Call()
					args.Id("r").Dot("RunID").Call()
					if hasInput {
						args.Id("req")
					}
					args.Id("opts").Op("...")
				}),
			),
		)
}

func (svc *Service) renderTestClient(f *g.File) {
	// generate test client
	svc.genTestClientImpl(f)
	svc.genTestClientImplNewMethod(f)
	for _, workflow := range svc.workflowsOrdered {
		svc.genTestClientImplWorkflowMethod(f, workflow)
		svc.genTestClientImplWorkflowAsyncMethod(f, workflow)
		svc.genTestClientImplWorkflowGetMethod(f, workflow)
		for _, signal := range svc.workflows[workflow].GetSignal() {
			if !signal.GetStart() {
				continue
			}
			svc.genTestClientImplWorkflowWithSignalMethod(f, workflow, signal.GetRef())
			svc.genTestClientImplWorkflowWithSignalAsyncMethod(f, workflow, signal.GetRef())
		}
	}

	svc.genTestClientImplWorkflowCancelMethod(f)
	svc.genTestClientImplWorkflowTerminateMethod(f)

	// generate test client query methods
	for _, query := range svc.queriesOrdered {
		svc.genTestClientImplQueryMethod(f, query)
	}

	// generate test client signal methods
	for _, signal := range svc.signalsOrdered {
		svc.genTestClientImplSignalMethod(f, signal)
	}

	// generate test client update methods
	for _, update := range svc.updatesOrdered {
		svc.genTestClientImplUpdateMethod(f, update)
		svc.genTestClientImplUpdateAsyncMethod(f, update)

		svc.genTestClientUpdateHandleImpl(f, update)
		svc.genTestClientUpdateHandleImplGetMethod(f, update)
		svc.genTestClientUpdateHandleImplRunIDMethod(f, update)
		svc.genTestClientUpdateHandleImplUpdateIDMethod(f, update)
		svc.genTestClientUpdateHandleImplWorkflowIDMethod(f, update)
	}

	// generate workflow test runs
	for _, workflow := range svc.workflowsOrdered {
		opts := svc.workflows[workflow]
		svc.genTestClientWorkflowRunImpl(f, workflow)
		svc.genTestClientWorkflowRunImplCancelMethod(f, workflow)
		svc.genTestClientWorkflowRunImplGetMethod(f, workflow)
		svc.genTestClientWorkflowRunImplIDMethod(f, workflow)
		svc.genTestClientWorkflowRunImplRunIDMethod(f, workflow)
		svc.genTestClientWorkflowRunImplTerminateMethod(f, workflow)

		// generate query methods
		for _, queryOpts := range opts.GetQuery() {
			svc.genTestClientWorkflowRunImplQueryMethod(f, workflow, queryOpts.GetRef())
		}

		// generate signal methods
		for _, signalOpts := range opts.GetSignal() {
			svc.genTestClientWorkflowRunImplSignalMethod(f, workflow, signalOpts.GetRef())
		}

		// generate update methods
		for _, updateOpts := range opts.GetUpdate() {
			svc.genTestClientWorkflowRunImplUpdateMethod(f, workflow, updateOpts.GetRef())
			svc.genTestClientWorkflowRunImplUpdateAsyncMethod(f, workflow, updateOpts.GetRef())
		}
	}
}
