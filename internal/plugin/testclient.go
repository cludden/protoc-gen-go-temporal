package plugin

import (
	g "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	testsuitePkg = "go.temporal.io/sdk/testsuite"
	testutilPkg  = "github.com/cludden/protoc-gen-go-temporal/pkg/testutil"
)

// genTestClientImpl generates a TestClient struct
func (svc *Manifest) genTestClientImpl(f *g.File) {
	f.Comment("TestClient provides a testsuite-compatible Client")
	f.Type().Id(svc.toCamel("Test%sClient", svc.Service.GoName)).Struct(
		g.Id("env").Op("*").Qual(testsuitePkg, "TestWorkflowEnvironment"),
		g.Id("workflows").Id(svc.toCamel("%sWorkflows", svc.Service.GoName)),
	)
}

// genTestClientImplNewMethod generates a NewTestClient constructor function
func (svc *Manifest) genTestClientImplNewMethod(f *g.File) {
	interfaceName := svc.toCamel("%sClient", svc.Service.GoName)
	typeName := svc.toCamel("Test%sClient", svc.Service.GoName)
	functionName := "New" + typeName

	f.Var().Id("_").Id(interfaceName).Op("=").Op("&").Id(typeName).Values()
	f.Commentf("%s initializes a new %s value", functionName, typeName)
	f.Func().Id(functionName).
		Params(
			g.Id("env").Op("*").Qual(testsuitePkg, "TestWorkflowEnvironment"),
			g.Id("workflows").Id(svc.toCamel("%sWorkflows", svc.Service.GoName)),
			g.Id("activities").Id(svc.toCamel("%sActivities", svc.Service.GoName)),
		).
		Params(
			g.Op("*").Id(typeName),
		).
		Block(
			g.If(g.Id("workflows").Op("!=").Nil()).Block(
				g.Id(svc.toCamel("Register%sWorkflows", svc.Service.GoName)).Call(g.Id("env"), g.Id("workflows")),
			),
			g.If(g.Id("activities").Op("!=").Nil()).Block(
				g.Id(svc.toCamel("Register%sActivities", svc.Service.GoName)).Call(g.Id("env"), g.Id("activities")),
			),
			g.Return(g.Op("&").Id(typeName).Values(g.Id("env"), g.Id("workflows"))),
		)
}

// genTestClientImplQueryMethod genereates a TestClient <Query> method
func (svc *Manifest) genTestClientImplQueryMethod(f *g.File, query protoreflect.FullName) {
	handler := svc.methods[query]
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)

	f.Commentf("%s executes a %s query", svc.methods[query].GoName, svc.fqnForQuery(query))
	f.Func().
		Params(g.Id("c").Op("*").Id(svc.toCamel("Test%sClient", svc.Service.GoName))).
		Id(svc.methods[query].GoName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
			}
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), svc.getMessageName(handler.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			fn.List(g.Id("val"), g.Err()).Op(":=").Id("c").Dot("env").Dot("QueryWorkflow").CallFunc(func(args *g.Group) {
				args.Id(svc.toCamel("%sQueryName", query))
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
					bl.Var().Id("result").Qual(string(handler.Output.GoIdent.GoImportPath), svc.getMessageName(handler.Output))
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
func (svc *Manifest) genTestClientImplSignalMethod(f *g.File, signal protoreflect.FullName) {
	handler := svc.methods[signal]
	hasInput := !isEmpty(handler.Input)

	f.Commentf("%s executes a %s signal", svc.methods[signal].GoName, svc.fqnForSignal(signal))
	f.Func().
		Params(g.Id("c").Op("*").Id(svc.toCamel("Test%sClient", svc.Service.GoName))).
		Id(svc.methods[signal].GoName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
			}
		}).
		Params(
			g.Error(),
		).
		Block(
			g.Id("c").Dot("env").Dot("SignalWorkflow").CallFunc(func(args *g.Group) {
				args.Id(svc.toCamel("%sSignalName", signal))
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
func (svc *Manifest) genTestClientImplUpdateMethod(f *g.File, update protoreflect.FullName) {
	method := svc.methods[update]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	asyncName := svc.toCamel("%sAsync", update)

	f.Commentf("%s executes a(n) %s update in the test environment", svc.methods[update].GoName, svc.fqnForUpdate(update))
	f.Func().
		Params(g.Id("c").Op("*").Id(svc.toCamel("Test%sClient", svc.Service.GoName))).
		Id(svc.methods[update].GoName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(svc.toCamel("%sOptions", update))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		Block(
			// initialize update request options
			g.Id("options").Op(":=").Id(svc.toCamel("New%sOptions", update)).Call(),
			g.If(g.Len(g.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(g.Lit(0)).Dot("Options").Op("!=").Nil()).Block(
				g.Id("options").Op("=").Id("opts").Index(g.Lit(0)),
			),

			// override wait policy
			g.Id("options").Dot("Options").Dot("WaitForStage").Op("=").Qual(clientPkg, "WorkflowUpdateStageCompleted"),
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
func (svc *Manifest) genTestClientImplUpdateAsyncMethod(f *g.File, update protoreflect.FullName) {
	method := svc.methods[update]
	hasInput := !isEmpty(method.Input)
	methodName := svc.toCamel("%sAsync", update)

	f.Commentf("%s executes a(n) %s update in the test environment", methodName, svc.fqnForUpdate(update))
	f.Func().
		Params(g.Id("c").Op("*").Id(svc.toCamel("Test%sClient", svc.Service.GoName))).
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(svc.toCamel("%sOptions", update))
		}).
		Params(
			g.Id(svc.toCamel("%sHandle", update)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			// initialize options
			fn.Var().Id("o").Op("*").Id(svc.toCamel("%sOptions", update))
			fn.If(g.Len(g.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(g.Lit(0)).Op("!=").Nil()).Block(
				g.Id("o").Op("=").Id("opts").Index(g.Lit(0)),
			).Else().Block(
				g.Id("o").Op("=").Id(svc.toCamel("New%sOptions", update)).Call(),
			)

			// build UpdateWorkflowWithOptions
			fn.List(g.Id("options"), g.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(args *g.Group) {
				args.Id("workflowID")
				args.Id("runID")
				if hasInput {
					args.Id("req")
				}
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error initializing UpdateWorkflowWithOptions: %w"), g.Err())),
			).Line()

			fn.If(g.Id("options").Dot("UpdateID").Op("==").Lit("")).Block(
				g.Id("options").Dot("UpdateID").Op("=").Id("workflowID"),
			).Line()

			// update workflow
			fn.Id("uc").Op(":=").Qual(testutilPkg, "NewUpdateCallbacks").Call()
			fn.Id("c").Dot("env").Dot("UpdateWorkflow").CallFunc(func(args *g.Group) {
				args.Id(svc.toCamel("%sUpdateName", update))
				args.Id("options").Dot("UpdateID")
				args.Id("uc")
				if hasInput {
					args.Id("req")
				}
			})

			fn.Return(
				g.Op("&").Id(svc.toLowerCamel("test%sHandle", update)).CustomFunc(multiLineValues, func(fields *g.Group) {
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

func (svc *Manifest) genTestClientImplUpdateGetMethod(f *g.File, update protoreflect.FullName) {
	methodName := svc.toCamel("Get%s", update)

	f.Commentf("%s retrieves a handle to an existing %s update", methodName, svc.fqnForUpdate(update))
	f.Func().
		Params(g.Id("c").Op("*").Id(svc.toCamel("Test%sClient", svc.Service.GoName))).
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("req").Qual(clientPkg, "GetWorkflowUpdateHandleOptions")
		}).
		Params(
			g.Id(svc.toCamel("%sHandle", update)),
			g.Error(),
		).
		Block(
			g.Return(
				g.Nil(),
				g.Qual("errors", "New").Call(g.Lit("unimplemented")),
			),
		)
}

// genTestClientImplWorkflowMethod generates a TestClient <Workflow> method
func (svc *Manifest) genTestClientImplWorkflowMethod(f *g.File, workflow protoreflect.FullName) {
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s executes a(n) %s workflow in the test environment", svc.methods[workflow].GoName, svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("c").Op("*").Id(svc.toCamel("Test%sClient", svc.Service.GoName))).
		Id(svc.methods[workflow].GoName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(svc.toCamel("%sOptions", workflow))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		Block(
			g.List(g.Id("run"), g.Err()).Op(":=").Id("c").Dot(svc.toCamel("%sAsync", workflow)).CallFunc(func(args *g.Group) {
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
func (svc *Manifest) genTestClientImplWorkflowAsyncMethod(f *g.File, workflow protoreflect.FullName) {
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)

	f.Commentf("%sAsync executes a(n) %s workflow in the test environment", svc.methods[workflow].GoName, svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("c").Op("*").Id(svc.toCamel("Test%sClient", svc.Service.GoName))).
		Id(svc.toCamel("%sAsync", workflow)).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
			args.Id("options").Op("...").Op("*").Id(svc.toCamel("%sOptions", workflow))
		}).
		Params(
			g.Id(svc.toCamel("%sRun", workflow)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			// initialize options
			fn.Var().Id("o").Op("*").Id(svc.toCamel("%sOptions", workflow))
			fn.If(g.Len(g.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(g.Lit(0)).Op("!=").Nil()).Block(
				g.Id("o").Op("=").Id("options").Index(g.Lit(0)),
			).Else().Block(
				g.Id("o").Op("=").Id(svc.toCamel("New%sOptions", workflow)).Call(),
			)

			// initialize client.StartWorkfowOptions
			fn.List(g.Id("opts"), g.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(args *g.Group) {
				if hasInput {
					args.Id("req").Dot("ProtoReflect").Call()
				} else {
					args.Nil()
				}
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error initializing client.StartWorkflowOptions: %w"), g.Err())),
			)

			fn.Return(
				g.Op("&").Id(svc.toLowerCamel("test%sRun", workflow)).ValuesFunc(func(fields *g.Group) {
					fields.Id("client").Op(":").Id("c")
					fields.Id("env").Op(":").Id("c").Dot("env")
					fields.Id("opts").Op(":").Op("&").Id("opts")
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
func (svc *Manifest) genTestClientImplWorkflowCancelMethod(f *g.File) {
	clientType := svc.toCamel("Test%sClient", svc.Service.GoName)
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
func (svc *Manifest) genTestClientImplWorkflowGetMethod(f *g.File, workflow protoreflect.FullName) {
	f.Commentf("%s is a noop", svc.toCamel("Get%s", workflow))
	f.Func().
		Params(g.Id("c").Op("*").Id(svc.toCamel("Test%sClient", svc.Service.GoName))).
		Id(svc.toCamel("Get%s", workflow)).
		Params(
			g.Id("ctx").Qual("context", "Context"),
			g.Id("workflowID").String(),
			g.Id("runID").String(),
		).
		Params(
			g.Id(svc.toCamel("%sRun", workflow)),
		).
		Block(
			g.Return(
				g.Op("&").Id(svc.toLowerCamel("test%sRun", workflow)).Values(
					g.Id("env").Op(":").Id("c").Dot("env"),
					g.Id("workflows").Op(":").Id("c").Dot("workflows"),
				),
			),
		)
}

// genTestClientImplWorkflowWithSignalMethod generates a TestClient's <workflow>With<signal> method
func (svc *Manifest) genTestClientImplWorkflowWithSignalMethod(f *g.File, workflow, signal protoreflect.FullName) {
	method := svc.methods[workflow]
	methodName := svc.toCamel("%sWith%s", workflow, signal)
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)
	handler := svc.methods[signal]
	hasSignalInput := !isEmpty(handler.Input)

	f.Commentf("%s sends a(n) %s signal to a(n) %s workflow, starting it if necessary", methodName, svc.fqnForSignal(signal), svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("c").Op("*").Id(svc.toCamel("Test%sClient", svc.Service.GoName))).
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
			if hasSignalInput {
				args.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(svc.toCamel("%sOptions", workflow))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		Block(
			g.Id("c").Dot("env").Dot("RegisterDelayedCallback").Call(
				g.Func().Params().Block(
					g.Id("c").Dot("env").Dot("SignalWorkflow").CallFunc(func(args *g.Group) {
						args.Qual(svc.goImportPathForMethod(signal), svc.toCamel("%sSignalName", signal))
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
				g.Id("c").Dot(svc.methods[workflow].GoName).CallFunc(func(args *g.Group) {
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
func (svc *Manifest) genTestClientImplWorkflowWithSignalAsyncMethod(f *g.File, workflow, signal protoreflect.FullName) {
	method := svc.methods[workflow]
	methodName := svc.toCamel("%sWith%sAsync", workflow, signal)
	hasInput := !isEmpty(method.Input)
	handler := svc.methods[signal]
	hasSignalInput := !isEmpty(handler.Input)

	f.Commentf("%s sends a(n) %s signal to a(n) %s workflow, starting it if necessary", methodName, svc.fqnForSignal(signal), svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("c").Op("*").Id(svc.toCamel("Test%sClient", svc.Service.GoName))).
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
			if hasSignalInput {
				args.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(svc.toCamel("%sOptions", workflow))
		}).
		Params(
			g.Id(svc.toCamel("%sRun", workflow)),
			g.Error(),
		).
		Block(
			g.Id("c").Dot("env").Dot("RegisterDelayedCallback").Call(
				g.Func().Params().Block(
					g.Id("c").Dot("env").Dot("SignalWorkflow").CallFunc(func(args *g.Group) {
						args.Qual(svc.goImportPathForMethod(signal), svc.toCamel("%sSignalName", signal))
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
				g.Id("c").Dot(svc.toCamel("%sAsync", workflow)).CallFunc(func(args *g.Group) {
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
func (svc *Manifest) genTestClientImplWorkflowTerminateMethod(f *g.File) {
	clientType := svc.toCamel("Test%sClient", svc.Service.GoName)
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

func (svc *Manifest) genTestClientUpdateHandleImpl(f *g.File, update protoreflect.FullName) {
	typeName := svc.toLowerCamel("test%sHandle", update)
	interfaceName := svc.toCamel("%sHandle", update)
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
			fields.Id("opts").Op("*").Qual(clientPkg, "UpdateWorkflowOptions")
			if hasInput {
				fields.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
			}
			fields.Id("runID").String()
			fields.Id("workflowID").String()
		})
}

func (svc *Manifest) genTestClientUpdateHandleImplGetMethod(f *g.File, update protoreflect.FullName) {
	method := svc.methods[update]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("Get retrieves a test %s update result", svc.fqnForUpdate(update))
	f.Func().
		Params(g.Id("h").Op("*").Id(svc.toLowerCamel("test%sHandle", update))).
		Id("Get").
		Params(g.Id("ctx").Qual("context", "Context")).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		Block(
			g.If(g.ListFunc(func(ls *g.Group) {
				if hasOutput {
					ls.Id("resp")
				} else {
					ls.Id("_")
				}
				ls.Err()
			}).Op(":=").Id("h").Dot("callbacks").Dot("Get").Call(g.Id("ctx")), g.Err().Op("!=").Nil()).
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
							returnVals.Id("resp").Op(".").Parens(g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output)))
						}
						returnVals.Nil()
					}),
				),
		)
}

func (svc *Manifest) genTestClientUpdateHandleImplRunIDMethod(f *g.File, update protoreflect.FullName) {
	f.Comment("RunID implementation")
	f.Func().
		Params(g.Id("h").Op("*").Id(svc.toLowerCamel("test%sHandle", update))).
		Id("RunID").
		Params().
		Params(g.String()).
		Block(
			g.Return(g.Id("h").Dot("runID")),
		)
}

func (svc *Manifest) genTestClientUpdateHandleImplUpdateIDMethod(f *g.File, update protoreflect.FullName) {
	f.Comment("UpdateID implementation")
	f.Func().
		Params(g.Id("h").Op("*").Id(svc.toLowerCamel("test%sHandle", update))).
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

func (svc *Manifest) genTestClientUpdateHandleImplWorkflowIDMethod(f *g.File, update protoreflect.FullName) {
	f.Comment("WorkflowID implementation")
	f.Func().
		Params(g.Id("h").Op("*").Id(svc.toLowerCamel("test%sHandle", update))).
		Id("WorkflowID").
		Params().
		Params(g.String()).
		Block(
			g.Return(g.Id("h").Dot("workflowID")),
		)
}

// genTestClientWorkflowRunImpl generates a test<Workflow>Run struct
func (svc *Manifest) genTestClientWorkflowRunImpl(f *g.File, workflow protoreflect.FullName) {
	method := svc.methods[workflow]
	methodName := svc.toLowerCamel("test%sRun", workflow)
	hasInput := !isEmpty(method.Input)
	// generate test<Workflow>Run struct
	f.Var().Id("_").Id(svc.toCamel("%sRun", workflow)).Op("=").Op("&").Id(svc.toLowerCamel("test%sRun", workflow)).Values()
	f.Commentf("%s provides convenience methods for interacting with a(n) %s workflow in the test environment", methodName, svc.fqnForWorkflow(workflow))
	f.Type().Id(methodName).StructFunc(func(fields *g.Group) {
		fields.Id("client").Op("*").Id(svc.toCamel("Test%sClient", svc.Service.GoName))
		fields.Id("env").Op("*").Qual(testsuitePkg, "TestWorkflowEnvironment")
		fields.Id("opts").Op("*").Qual(clientPkg, "StartWorkflowOptions")
		if hasInput {
			fields.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
		}
		fields.Id("workflows").Id(svc.toCamel("%sWorkflows", svc.Service.GoName))
	})
}

// genClientWorkflowRunImplCancelMethod generates a <Workflow>Run's Cancel method
func (svc *Manifest) genTestClientWorkflowRunImplCancelMethod(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toLowerCamel("Test%sRun", workflow)

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
func (svc *Manifest) genTestClientWorkflowRunImplGetMethod(f *g.File, workflow protoreflect.FullName) {
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	f.Commentf("Get retrieves a test %s workflow result", svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("r").Op("*").Id(svc.toLowerCamel("test%sRun", workflow))).
		Id("Get").
		Params(g.Qual("context", "Context")).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			// execute workflow
			fn.Id("r").Dot("env").Dot("ExecuteWorkflow").CallFunc(func(args *g.Group) {
				args.Id(svc.toCamel("%sWorkflowName", workflow))
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
				fn.Var().Id("result").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
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
func (svc *Manifest) genTestClientWorkflowRunImplIDMethod(f *g.File, workflow protoreflect.FullName) {
	f.Commentf("ID returns a test %s workflow run's workflow ID", svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("r").Op("*").Id(svc.toLowerCamel("test%sRun", workflow))).
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
func (svc *Manifest) genTestClientWorkflowRunImplQueryMethod(f *g.File, workflow, query protoreflect.FullName) {
	handler := svc.methods[query]
	handlerName := svc.methods[query].GoName
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)

	f.Commentf("%s executes a %s query against a test %s workflow", handlerName, svc.fqnForQuery(query), svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("r").Op("*").Id(svc.toLowerCamel("test%sRun", workflow))).
		Id(handlerName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
			}
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), svc.getMessageName(handler.Output))
			}
			returnVals.Error()
		}).Block(
		g.Return(
			g.Id("r").Dot("client").Dot(svc.methods[query].GoName).CallFunc(func(args *g.Group) {
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

// genTestClientWorkflowRunImplRunMethod generates a test<Workflow>Run's Run method
func (svc *Manifest) genTestClientWorkflowRunImplRunMethod(f *g.File, workflow protoreflect.FullName) {
	f.Comment("Run noop implementation")
	f.Func().
		Params(g.Id("r").Op("*").Id(svc.toLowerCamel("test%sRun", workflow))).
		Id("Run").
		Params().
		Qual(clientPkg, "WorkflowRun").
		Block(
			g.Return(g.Nil()),
		)
}

// genTestClientWorkflowRunImplRunIDMethod generates a test<Workflow>Run's RunID method
func (svc *Manifest) genTestClientWorkflowRunImplRunIDMethod(f *g.File, workflow protoreflect.FullName) {
	f.Comment("RunID noop implementation")
	f.Func().
		Params(g.Id("r").Op("*").Id(svc.toLowerCamel("test%sRun", workflow))).
		Id("RunID").
		Params().
		Params(g.String()).
		Block(
			g.Return(g.Lit("")),
		)
}

// genTestClientWorkflowRunImplQueryMethod generates a test<Workflow>Run's <Signal> method
func (svc *Manifest) genTestClientWorkflowRunImplSignalMethod(f *g.File, workflow, signal protoreflect.FullName) {
	handler := svc.methods[signal]
	handlerName := svc.methods[signal].GoName
	hasInput := !isEmpty(handler.Input)

	f.Commentf("%s executes a %s signal against a test %s workflow", handlerName, svc.fqnForSignal(signal), svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("r").Op("*").Id(svc.toLowerCamel("test%sRun", workflow))).
		Id(handlerName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
			}
		}).
		Params(
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			if svc.methodsFromSameService(signal, workflow) {
				fn.Return().Id("r").Dot("client").Dot(handlerName).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("r").Dot("ID").Call()
					args.Id("r").Dot("RunID").Call()
					if hasInput {
						args.Id("req")
					}
				})
			} else {
				fn.Id("r").Dot("env").Dot("SignalWorkflow").CallFunc(func(args *g.Group) {
					args.Add(svc.Qual(signal, svc.toCamel("%sSignalName", signal)))
					if hasInput {
						args.Id("req")
					} else {
						args.Nil()
					}
				})
				fn.Return(g.Nil())
			}
		})
}

// genClientWorkflowRunImplTerminateMethod generates a <Workflow>Run's Terminate method
func (svc *Manifest) genTestClientWorkflowRunImplTerminateMethod(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toLowerCamel("Test%sRun", workflow)

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
func (svc *Manifest) genTestClientWorkflowRunImplUpdateMethod(f *g.File, workflow, update protoreflect.FullName) {
	handler := svc.methods[update]
	handlerName := svc.methods[update].GoName
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)

	f.Commentf("%s executes a(n) %s update against a test %s workflow", handlerName, svc.fqnForUpdate(update), svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("r").Op("*").Id(svc.toLowerCamel("test%sRun", workflow))).
		Id(handlerName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(svc.toCamel("%sOptions", update))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), svc.getMessageName(handler.Output))
			}
			returnVals.Error()
		}).
		Block(
			g.Return(
				g.Id("r").Dot("client").Dot(svc.methods[update].GoName).CallFunc(func(args *g.Group) {
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
func (svc *Manifest) genTestClientWorkflowRunImplUpdateAsyncMethod(f *g.File, workflow, update protoreflect.FullName) {
	handler := svc.methods[update]
	hasInput := !isEmpty(handler.Input)
	methodName := svc.toCamel("%sAsync", update)

	f.Commentf("%s executes a(n) %s update against a test %s workflow", methodName, svc.fqnForUpdate(update), svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("r").Op("*").Id(svc.toLowerCamel("test%sRun", workflow))).
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(svc.toCamel("%sOptions", update))
		}).
		Params(
			g.Id(svc.toCamel("%sHandle", update)),
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

func (svc *Manifest) renderTestClient(f *g.File) {
	// generate test client
	svc.genTestClientImpl(f)
	svc.genTestClientImplNewMethod(f)
	for _, workflow := range svc.workflowsOrdered {
		if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
			continue
		}
		svc.genTestClientImplWorkflowMethod(f, workflow)
		svc.genTestClientImplWorkflowAsyncMethod(f, workflow)
		svc.genTestClientImplWorkflowGetMethod(f, workflow)
		for _, signal := range svc.workflows[workflow].GetSignal() {
			if !signal.GetStart() {
				continue
			}
			svc.genTestClientImplWorkflowWithSignalMethod(f, workflow, getFullyQualifiedRef(workflow, signal.GetRef()))
			svc.genTestClientImplWorkflowWithSignalAsyncMethod(f, workflow, getFullyQualifiedRef(workflow, signal.GetRef()))
		}
	}

	svc.genTestClientImplWorkflowCancelMethod(f)
	svc.genTestClientImplWorkflowTerminateMethod(f)

	// generate test client query methods
	for _, query := range svc.queriesOrdered {
		if svc.methods[query].Desc.Parent() != svc.Service.Desc {
			continue
		}
		svc.genTestClientImplQueryMethod(f, query)
	}

	// generate test client signal methods
	for _, signal := range svc.signalsOrdered {
		if svc.methods[signal].Desc.Parent() != svc.Service.Desc {
			continue
		}
		svc.genTestClientImplSignalMethod(f, signal)
	}

	// generate test client update methods
	for _, update := range svc.updatesOrdered {
		if svc.methods[update].Desc.Parent() != svc.Service.Desc {
			continue
		}
		svc.genTestClientImplUpdateMethod(f, update)
		svc.genTestClientImplUpdateAsyncMethod(f, update)
		svc.genTestClientImplUpdateGetMethod(f, update)

		svc.genTestClientUpdateHandleImpl(f, update)
		svc.genTestClientUpdateHandleImplGetMethod(f, update)
		svc.genTestClientUpdateHandleImplRunIDMethod(f, update)
		svc.genTestClientUpdateHandleImplUpdateIDMethod(f, update)
		svc.genTestClientUpdateHandleImplWorkflowIDMethod(f, update)
	}

	// generate workflow test runs
	for _, workflow := range svc.workflowsOrdered {
		if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
			continue
		}
		opts := svc.workflows[workflow]
		svc.genTestClientWorkflowRunImpl(f, workflow)
		svc.genTestClientWorkflowRunImplCancelMethod(f, workflow)
		svc.genTestClientWorkflowRunImplGetMethod(f, workflow)
		svc.genTestClientWorkflowRunImplIDMethod(f, workflow)
		svc.genTestClientWorkflowRunImplRunMethod(f, workflow)
		svc.genTestClientWorkflowRunImplRunIDMethod(f, workflow)
		svc.genTestClientWorkflowRunImplTerminateMethod(f, workflow)

		// generate query methods
		for _, queryOpts := range opts.GetQuery() {
			svc.genTestClientWorkflowRunImplQueryMethod(f, workflow, getFullyQualifiedRef(workflow, queryOpts.GetRef()))
		}

		// generate signal methods
		for _, signalOpts := range opts.GetSignal() {
			svc.genTestClientWorkflowRunImplSignalMethod(f, workflow, getFullyQualifiedRef(workflow, signalOpts.GetRef()))
		}

		// generate update methods
		for _, updateOpts := range opts.GetUpdate() {
			svc.genTestClientWorkflowRunImplUpdateMethod(f, workflow, getFullyQualifiedRef(workflow, updateOpts.GetRef()))
			svc.genTestClientWorkflowRunImplUpdateAsyncMethod(f, workflow, getFullyQualifiedRef(workflow, updateOpts.GetRef()))
		}
	}
}
