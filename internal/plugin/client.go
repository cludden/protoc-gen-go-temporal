package plugin

import (
	"fmt"
	"strconv"
	"strings"

	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	g "github.com/dave/jennifer/jen"
)

// genClientImpl generates a <service>Client implementation
func (svc *Service) genClientImpl(f *g.File) {
	typeName := toLowerCamel("%sClient", svc.Service.GoName)

	f.Commentf("%s implements a temporal client for a %s service", typeName, svc.Service.Desc.FullName())
	f.Type().
		Id(typeName).
		StructFunc(func(fields *g.Group) {
			fields.Id("client").Qual(clientPkg, "Client")
		})
}

// genClientImplConstructor generates a New<Service>Client function
func (svc *Service) genClientImplConstructor(f *g.File) {
	methodName := toCamel("New%sClient", svc.Service.GoName)
	implName := toLowerCamel("%sClient", svc.Service.GoName)
	interfaceName := toCamel("%sClient", svc.Service.GoName)

	f.Commentf("%s initializes a new %s client", methodName, svc.Service.Desc.FullName())
	f.Func().
		Id(methodName).
		Params(
			g.Id("c").Qual(clientPkg, "Client"),
		).
		Params(
			g.Id(interfaceName),
		).
		Block(
			g.Return(
				g.Op("&").Id(implName).Values(
					g.Id("client").Op(":").Id("c"),
				),
			),
		)

	methodName += "WithOptions"
	f.Commentf("%s initializes a new %s client with the given options", methodName, svc.Service.GoName)
	f.Func().
		Id(methodName).
		Params(
			g.Id("c").Qual(clientPkg, "Client"),
			g.Id("opts").Qual(clientPkg, "Options"),
		).
		Params(
			g.Id(interfaceName),
			g.Error(),
		).
		Block(
			g.Var().Err().Error(),
			g.List(g.Id("c"), g.Err()).Op("=").Qual(clientPkg, "NewClientFromExisting").Call(g.Id("c"), g.Id("opts")),
			g.If().Err().Op("!=").Nil().Block(
				g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error initializing client with options: %w"), g.Err())),
			),
			g.Return(
				g.Op("&").Id(implName).Values(
					g.Id("client").Op(":").Id("c"),
				),
				g.Nil(),
			),
		)
}

// genClientImplQueryMethod adds a <Query> method to a workflowClient
func (svc *Service) genClientImplQueryMethod(f *g.File, query string) {
	clientType := toLowerCamel("%sClient", svc.Service.GoName)
	method := svc.methods[query]
	hasInput := !isEmpty(method.Input)

	f.Commentf("%s sends a(n) %s query to an existing workflow", query, svc.fqnForQuery(query))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(query).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("query").Op("*").Id(method.Input.GoIdent.GoName)
			}
		}).
		Params(
			g.Op("*").Id(method.Output.GoIdent.GoName),
			g.Error(),
		).
		Block(
			g.Var().Id("resp").Id(method.Output.GoIdent.GoName),
			g.If(
				g.List(g.Id("val"), g.Err()).Op(":=").Id("c").Dot("client").Dot("QueryWorkflow").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("workflowID")
					args.Id("runID")
					args.Id(fmt.Sprintf("%sQueryName", query))
					if hasInput {
						args.Id("query")
					}
				}),
				g.Err().Op("!=").Nil(),
			).Block(
				g.Return(g.Nil(), g.Err()),
			).Else().If(
				g.Err().Op("=").Id("val").Dot("Get").Call(
					g.Op("&").Id("resp"),
				),
				g.Err().Op("!=").Nil(),
			).Block(
				g.Return(g.Nil(), g.Err()),
			),
			g.Return(
				g.Op("&").Id("resp"), g.Nil(),
			),
		)
}

// genClientImplSignalMethod adds a <Signal> method to a workflowClient
func (svc *Service) genClientImplSignalMethod(f *g.File, signal string) {
	clientType := toLowerCamel("%sClient", svc.Service.GoName)
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)

	f.Commentf("%s sends a(n) %s signal to an existing workflow", signal, svc.fqnForSignal(signal))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(signal).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("signal").Op("*").Id(method.Input.GoIdent.GoName)
			}
		}).
		Params(g.Error()).
		Block(
			g.Return(
				g.Id("c").Dot("client").Dot("SignalWorkflow").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("workflowID")
					args.Id("runID")
					args.Id(fmt.Sprintf("%sSignalName", signal))
					if hasInput {
						args.Id("signal")
					} else {
						args.Nil()
					}
				}),
			),
		)
}

// genClientImplSignalWithStartAsyncMethod adds a <Workflow>With<Signal>Async client method
func (svc *Service) genClientImplSignalWithStartAsyncMethod(f *g.File, workflow, signal string) {
	clientType := toLowerCamel("%sClient", svc.Service.GoName)
	method := svc.methods[workflow]
	handler := svc.methods[signal]
	name := toCamel("%sWith%sAsync", workflow, signal)
	runName := toLowerCamel("%sRun", workflow)
	hasWorkflowInput := !isEmpty(method.Input)
	hasSignalInput := !isEmpty(handler.Input)

	f.Commentf("%s starts a(n) %s workflow and sends a(n) %s signal in a transaction", name, svc.fqnForWorkflow(workflow), svc.fqnForSignal(signal))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(name).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasWorkflowInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
			if hasSignalInput {
				args.Id("signal").Op("*").Id(handler.Input.GoIdent.GoName)
			}
			args.Id("options").Op("...").Op("*").Id(toCamel("%sOptions", workflow))
		}).
		Params(
			g.Id(toCamel("%sRun", workflow)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			// initialize StartWorkflowOptions
			svc.genClientStartWorkflowOptions(fn, workflow, false)

			// signal with start workflow
			fn.Id("run").Op(",").Err().Op(":=").Id("c").Dot("client").Dot("SignalWithStartWorkflow").CallFunc(func(args *g.Group) {
				args.Id("ctx")
				args.Id("opts").Dot("ID")
				args.Id(toCamel("%sSignalName", signal))
				if hasSignalInput {
					args.Id("signal")
				} else {
					args.Nil()
				}
				args.Op("*").Id("opts")
				args.Id(toCamel("%sWorkflowName", workflow))
				if hasWorkflowInput {
					args.Id("req")
				}
			})
			fn.If(g.Id("run").Op("==").Nil().Op("||").Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Err()),
			)
			fn.Return(
				g.Op("&").Id(runName).Block(
					g.Id("client").Op(":").Id("c").Op(","),
					g.Id("run").Op(":").Id("run").Op(","),
				),
				g.Nil(),
			)
		})
}

// genClientImplSignalWithStartMethod adds a Start<Workflow>With<Signal> client method
func (svc *Service) genClientImplSignalWithStartMethod(f *g.File, workflow, signal string) {
	clientType := toLowerCamel("%sClient", svc.Service.GoName)
	method := svc.methods[workflow]
	handler := svc.methods[signal]
	name := toCamel("%sWith%s", workflow, signal)
	hasWorkflowInput := !isEmpty(method.Input)
	hasWorkflowOutput := !isEmpty(method.Output)
	hasSignalInput := !isEmpty(handler.Input)

	f.Commentf("%s starts a(n) %s workflow and sends a(n) %s signal in a transaction", name, svc.fqnForWorkflow(workflow), svc.fqnForSignal(signal))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(name).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasWorkflowInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
			if hasSignalInput {
				args.Id("signal").Op("*").Id(handler.Input.GoIdent.GoName)
			}
			args.Id("options").Op("...").Op("*").Id(toCamel("%sOptions", workflow))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasWorkflowOutput {
				returnVals.Op("*").Id(method.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			// signal with start workflow
			fn.Id("run").Op(",").Err().Op(":=").Id("c").Dot(toCamel("%sWith%sAsync", workflow, signal)).CallFunc(func(args *g.Group) {
				args.Id("ctx")
				if hasWorkflowInput {
					args.Id("req")
				}
				if hasSignalInput {
					args.Id("signal")
				}
				args.Id("options").Op("...")
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasWorkflowOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)
			fn.Return(
				g.Id("run").Dot("Get").Call(g.Id("ctx")),
			)
		})
}

// genClientImplUpdateMethod adds an <Update> method to a workflowClient
func (svc *Service) genClientImplUpdateMethod(f *g.File, update string) {
	clientType := toLowerCamel("%sClient", svc.Service.GoName)
	handler := svc.methods[update]
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)

	f.Commentf("%s sends a(n) %s update to an existing workflow", update, svc.fqnForUpdate(update))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(update).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
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
		BlockFunc(func(method *g.Group) {
			// initialize update request options
			method.Id("options").Op(":=").Id(toCamel("New%sOptions", update)).Call()
			method.If(g.Len(g.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(g.Lit(0)).Dot("opts").Op("!=").Nil()).Block(
				g.Id("options").Op("=").Id("opts").Index(g.Lit(0)),
			)

			// override wait policy
			method.Id("options").Dot("opts").Dot("WaitPolicy").Op("=").Op("&").Qual(updatePkg, "WaitPolicy").Values(
				g.Id("LifecycleStage").Op(":").Qual(enumsPkg, "UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_COMPLETED"),
			)

			// call async method
			method.List(g.Id("handle"), g.Err()).Op(":=").Id("c").Dot(fmt.Sprintf("%sAsync", update)).CallFunc(func(args *g.Group) {
				args.Id("ctx")
				args.Id("workflowID")
				args.Id("runID")
				if hasInput {
					args.Id("req")
				}
				args.Id("options")
			})
			method.If(g.Err().Op("!=").Nil()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)

			// call handle get
			method.Return(g.Id("handle").Dot("Get").Call(g.Id("ctx")))
		})
}

// genClientImplUpdateMethodAsync adds an <Update>Async method to a workflowClient
func (svc *Service) genClientImplUpdateMethodAsync(f *g.File, update string) {
	clientType := toLowerCamel("%sClient", svc.Service.GoName)
	handler := svc.methods[update]
	handleName := toLowerCamel("%sHandle", update)
	hasInput := !isEmpty(handler.Input)
	methodName := toCamel("%sAsync", update)

	f.Commentf("%s sends a(n) %s update to an existing workflow", methodName, svc.fqnForUpdate(update))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
			}
			args.Id("opts").Op("...").Op("*").Id(toCamel("%sOptions", update))
		}).
		Params(
			g.Id(toCamel("%sHandle", update)),
			g.Error(),
		).
		BlockFunc(func(method *g.Group) {
			svc.genClientUpdateWorkflowOptions(method, update)

			// update workflow
			method.List(g.Id("handle"), g.Err()).Op(":=").Id("c").Dot("client").Dot("UpdateWorkflowWithOptions").Call(g.Id("ctx"), g.Id("options"))
			method.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Err()),
			)

			method.Return(
				g.Op("&").Id(handleName).Values(
					g.Id("client").Op(":").Id("c"),
					g.Id("handle").Op(":").Id("handle"),
				),
				g.Nil(),
			)
		})
}

// genClientImplWorkflowAsyncMethod generates an <Workflow>Async client method
func (svc *Service) genClientImplWorkflowAsyncMethod(f *g.File, workflow string) {
	clientType := toLowerCamel("%sClient", svc.Service.GoName)
	method := svc.methods[workflow]
	methodName := toCamel("%sAsync", workflow)
	runImplType := toLowerCamel("%sRun", workflow)
	runInterfaceType := toCamel("%sRun", workflow)
	hasInput := !isEmpty(method.Input)

	f.Commentf("%s starts a(n) %s workflow", methodName, svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
			args.Id("options").Op("...").Op("*").Id(toCamel("%sOptions", workflow))
		}).
		Params(
			g.Id(runInterfaceType),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			// initialize StartWorkflowOptions with defaults
			svc.genClientStartWorkflowOptions(fn, workflow, false)

			// execute workflow
			fn.Id("run").Op(",").Err().Op(":=").Id("c").Dot("client").Dot("ExecuteWorkflow").CallFunc(func(args *g.Group) {
				args.Id("ctx")
				args.Op("*").Id("opts")
				args.Id(fmt.Sprintf("%sWorkflowName", workflow))
				if hasInput {
					args.Id("req")
				}
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Err()),
			)
			fn.If(g.Id("run").Op("==").Nil()).Block(
				g.Return(g.Nil(), g.Qual("errors", "New").Call(g.Lit("execute workflow returned nil run"))),
			)
			fn.Return(
				g.Op("&").Id(runImplType).Block(
					g.Id("client").Op(":").Id("c").Op(","),
					g.Id("run").Op(":").Id("run").Op(","),
				),
				g.Nil(),
			)
		})
}

// genClientImplWorkflowCancelMethod generates a Cancel<Workflow> client method
func (svc *Service) genClientImplWorkflowCancelMethod(f *g.File) {
	clientType := toLowerCamel("%sClient", svc.Service.GoName)
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
			g.Return(
				g.Id("c").Dot("client").Dot(methodName).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("workflowID")
					args.Id("runID")
				}),
			),
		)
}

// genClientImplWorkflowGetMethod generates a Get<Workflow> client method
func (svc *Service) genClientImplWorkflowGetMethod(f *g.File, workflow string) {
	clientType := toLowerCamel("%sClient", svc.Service.GoName)
	methodName := toCamel("Get%s", workflow)
	runImplType := toLowerCamel("%sRun", workflow)
	runInterfaceType := toCamel("%sRun", workflow)

	f.Commentf("%s fetches an existing %s execution", methodName, svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(methodName).
		Params(
			g.Id("ctx").Qual("context", "Context"),
			g.Id("workflowID").String(),
			g.Id("runID").String(),
		).
		Params(
			g.Id(runInterfaceType),
		).
		Block(
			g.Return(
				g.Op("&").Id(runImplType).Block(
					g.Id("client").Op(":").Id("c").Op(","),
					g.Id("run").Op(":").Id("c").Dot("client").Dot("GetWorkflow").Call(
						g.Id("ctx"), g.Id("workflowID"), g.Id("runID"),
					).Op(","),
				),
			),
		)
}

// genClientImplWorkflowTerminateMethod generates a TerminateWorkflow client method
func (svc *Service) genClientImplWorkflowTerminateMethod(f *g.File) {
	clientType := toLowerCamel("%sClient", svc.Service.GoName)
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
				g.Id("c").Dot("client").Dot(methodName).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("workflowID")
					args.Id("runID")
					args.Id("reason")
					args.Id("details").Op("...")
				}),
			),
		)
}

// genClientImplWorkflowMethod generates an <Workflow> client method
func (svc *Service) genClientImplWorkflowMethod(f *g.File, workflow string) {
	clientType := toLowerCamel("%sClient", svc.Service.GoName)
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s executes a %s workflow and blocks until error or response received", workflow, svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(workflow).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
			args.Id("options").Op("...").Op("*").Id(toCamel("%sOptions", workflow))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(method.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			// execute workflow
			fn.Id("run").Op(",").Err().Op(":=").Id("c").Dot(toCamel("%sAsync", workflow)).CallFunc(func(args *g.Group) {
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
				g.Id("run").Dot("Get").Call(g.Id("ctx")),
			)
		})
}

// genClientInterface generates a Client interface for a given service
func (svc *Service) genClientInterface(f *g.File) {
	typeName := toCamel("%sClient", svc.Service.GoName)

	f.Commentf("%s describes a client for a(n) %s worker", typeName, svc.Service.Desc.FullName())
	f.Type().Id(typeName).InterfaceFunc(func(methods *g.Group) {
		for _, workflow := range svc.workflowsOrdered {
			opts := svc.workflows[workflow]

			method := svc.methods[workflow]
			hasInput := !isEmpty(method.Input)
			hasOutput := !isEmpty(method.Output)
			runInterfaceType := toCamel("%sRun", workflow)

			// generate <Workflow> method
			methodName := workflow
			if method.Comments.Leading.String() != "" {
				methods.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
			} else {
				methods.Commentf("%s executes a(n) %s workflow and blocks until error or response received", methodName, svc.fqnForWorkflow(workflow))
			}
			methods.Id(workflow).
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
				})

			// generate <Workflow>Async method
			methodName = toCamel("%sAsync", workflow)
			methods.Commentf("%s executes a(n) %s workflow asynchronously", methodName, svc.fqnForWorkflow(workflow))
			methods.Id(methodName).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
					}
					args.Id("opts").Op("...").Op("*").Id(toCamel("%sOptions", workflow))
				}).
				Params(
					g.Id(runInterfaceType),
					g.Error(),
				)

			// generate Get<Workflow> method
			methodName = toCamel("Get%s", workflow)
			methods.Commentf("%s retrieves a handle to an existing %s workflow execution", methodName, svc.fqnForWorkflow(workflow))
			methods.Id(toCamel("Get%s", workflow)).
				Params(
					g.Id("ctx").Qual("context", "Context"),
					g.Id("workflowID").String(),
					g.Id("runID").String(),
				).
				Params(
					g.Id(runInterfaceType),
				)

			// add <Workflow>With<Signal> methods
			for _, signalOpts := range opts.GetSignal() {
				if !signalOpts.GetStart() {
					continue
				}
				method := svc.methods[workflow]
				signal := signalOpts.GetRef()
				handler := svc.methods[signal]
				hasWorkflowInput := !isEmpty(method.Input)
				hasWorkflowOutput := !isEmpty(method.Output)
				hasSignalInput := !isEmpty(handler.Input)

				// add synchronous flavor
				methodName := toCamel("%sWith%s", workflow, signal)
				if desc := handler.Comments.Leading.String(); desc != "" {
					methods.Comment(strings.ReplaceAll(strings.TrimPrefix(desc, "//"), "\n//", ""))
				} else {
					methods.Commentf("%s sends a(n) %s signal to a(n) %s workflow, starting it if necessary, and blocks until workflow completion", methodName, svc.fqnForSignal(signal), svc.fqnForWorkflow(workflow))
				}
				methods.Id(methodName).
					ParamsFunc(func(args *g.Group) {
						args.Id("ctx").Qual("context", "Context")
						if hasWorkflowInput {
							args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
						}
						if hasSignalInput {
							args.Id("signal").Op("*").Id(handler.Input.GoIdent.GoName)
						}
						args.Id("opts").Op("...").Op("*").Id(toCamel("%sOptions", workflow))
					}).
					ParamsFunc(func(returnVals *g.Group) {
						if hasWorkflowOutput {
							returnVals.Op("*").Id(method.Output.GoIdent.GoName)
						}
						returnVals.Error()
					})

				// add async flavor
				methodName += "Async"
				if desc := handler.Comments.Leading.String(); desc != "" {
					methods.Comment(strings.ReplaceAll(strings.TrimPrefix(desc, "//"), "\n//", ""))
				} else {
					methods.Commentf("%s sends a(n) %s signal to a(n) %s workflow, starting it if necessary, and returns a handle to the workflow execution", methodName, svc.fqnForSignal(signal), svc.fqnForWorkflow(workflow))
				}
				methods.Id(methodName).
					ParamsFunc(func(args *g.Group) {
						args.Id("ctx").Qual("context", "Context")
						if hasWorkflowInput {
							args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
						}
						if hasSignalInput {
							args.Id("signal").Op("*").Id(handler.Input.GoIdent.GoName)
						}
						args.Id("opts").Op("...").Op("*").Id(toCamel("%sOptions", workflow))
					}).
					Params(
						g.Id(runInterfaceType),
						g.Error(),
					)
			}
		}

		// generate CancelWorkflow method
		methodName := "CancelWorkflow"
		methods.Commentf("%s requests cancellation of an existing workflow execution", methodName)
		methods.Id(methodName).
			Params(
				g.Id("ctx").Qual("context", "Context"),
				g.Id("workflowID").String(),
				g.Id("runID").String(),
			).
			Params(g.Error())

		// generate TerminateWorkflow method
		methodName = "TerminateWorkflow"
		methods.Commentf("%s an existing workflow execution", methodName)
		methods.Id(methodName).
			Params(
				g.Id("ctx").Qual("context", "Context"),
				g.Id("workflowID").String(),
				g.Id("runID").String(),
				g.Id("reason").String(),
				g.Id("details").Op("...").Interface(),
			).
			Params(g.Error())

		// add <Query> methods
		for _, query := range svc.queriesOrdered {
			handler := svc.methods[query]
			hasInput := !isEmpty(handler.Input)
			if desc := handler.Comments.Leading.String(); desc != "" {
				methods.Comment(strings.ReplaceAll(strings.TrimPrefix(desc, "//"), "\n//", ""))
			} else {
				methods.Commentf("%s executes a(n) %s query", query, svc.fqnForQuery(query))
			}
			methods.Id(query).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("workflowID").String()
					args.Id("runID").String()
					if hasInput {
						args.Id("query").Op("*").Id(handler.Input.GoIdent.GoName)
					}
				}).
				Params(
					g.Op("*").Id(handler.Output.GoIdent.GoName),
					g.Error(),
				)
		}

		// add <Signal> methods
		for _, signal := range svc.signalsOrdered {
			handler := svc.methods[signal]
			hasInput := !isEmpty(handler.Input)
			if desc := handler.Comments.Leading.String(); desc != "" {
				methods.Comment(strings.ReplaceAll(strings.TrimPrefix(desc, "//"), "\n//", ""))
			} else {
				methods.Commentf("%s sends a(n) %s signal", signal, svc.fqnForSignal(signal))
			}
			methods.Id(signal).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("workflowID").String()
					args.Id("runID").String()
					if hasInput {
						args.Id("signal").Op("*").Id(handler.Input.GoIdent.GoName)
					}
				}).
				Params(g.Error())
		}

		// add <Update> methods
		for _, update := range svc.updatesOrdered {
			handler := svc.methods[update]
			hasInput := !isEmpty(handler.Input)
			hasOutput := !isEmpty(handler.Output)

			// add synchronous flavor
			if desc := handler.Comments.Leading.String(); desc != "" {
				methods.Comment(strings.ReplaceAll(strings.TrimPrefix(desc, "//"), "\n//", ""))
			} else {
				methods.Commentf("%s executes a(n) %s update and blocks until update completion", update, svc.fqnForUpdate(update))
			}
			methods.Id(update).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("workflowID").String()
					args.Id("runID").String()
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
				})

			// add async flavor
			methodName := toCamel("%sAsync", update)
			if desc := handler.Comments.Leading.String(); desc != "" {
				methods.Comment(strings.ReplaceAll(strings.TrimPrefix(desc, "//"), "\n//", ""))
			} else {
				methods.Commentf("%s executes a(n) %s update and blocks until update completion", update, svc.fqnForUpdate(update))
			}
			methods.Id(methodName).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("workflowID").String()
					args.Id("runID").String()
					if hasInput {
						args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
					}
					args.Id("opts").Op("...").Op("*").Id(toCamel("%sOptions", update))
				}).
				Params(
					g.Id(toCamel("%sHandle", update)),
					g.Error(),
				)
		}
	})
}

// genClientStartWorkflowOptions adds logic for initializing StartWorkflowOptions with default values
func (svc *Service) genClientStartWorkflowOptions(fn *g.Group, workflow string, child bool) {
	method := svc.methods[workflow]
	opts := svc.workflows[workflow]
	hasInput := !isEmpty(method.Input)

	// initialize options if nil
	if child {
		fn.Var().Id("opts").Op("*").Qual(workflowPkg, "ChildWorkflowOptions")
		fn.If(g.Len(g.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(g.Lit(0)).Dot("opts").Op("!=").Nil()).
			Block(
				g.Id("opts").Op("=").Id("options").Index(g.Lit(0)).Dot("opts"),
			).
			Else().
			Block(
				g.Id("childOpts").Op(":=").Qual(workflowPkg, "GetChildWorkflowOptions").Call(g.Id("ctx")),
				g.Id("opts").Op("=").Op("&").Id("childOpts"),
			)
	} else {
		fn.Id("opts").Op(":=").Op("&").Qual(clientPkg, "StartWorkflowOptions").Values()
		fn.If(g.Len(g.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(g.Lit(0)).Dot("opts").Op("!=").Nil()).Block(
			g.Id("opts").Op("=").Id("options").Index(g.Lit(0)).Dot("opts"),
		)
	}

	// set task queue if unset and default available
	var taskQueue g.Code
	if tq := opts.GetTaskQueue(); tq != "" {
		taskQueue = g.Lit(tq)
	} else if tq = svc.opts.GetTaskQueue(); tq != "" {
		taskQueue = g.Id(toCamel("%sTaskQueue", svc.Service.GoName))
	}
	if taskQueue != nil {
		fn.If(g.Id("opts").Dot("TaskQueue").Op("==").Lit("")).Block(
			g.Id("opts").Dot("TaskQueue").Op("=").Add(taskQueue),
		)
	}

	idFieldName := "ID"
	if child {
		idFieldName = "WorkflowID"
	}

	// set workflow id if unset and  id field and/or prefix defined
	if idExpr := opts.GetId(); idExpr != "" {
		fn.If(g.Id("opts").Dot(idFieldName).Op("==").Lit("")).BlockFunc(func(b *g.Group) {
			b.List(g.Id("id"), g.Err()).Op(":=").Qual(expressionPkg, "EvalExpression").CallFunc(func(args *g.Group) {
				args.Id(toCamel("%sIDExpression", workflow))
				if hasInput {
					args.Id("req").Dot("ProtoReflect").Call()
				} else {
					args.Nil()
				}
			})
			b.If(g.Err().Op("!=").Nil()).BlockFunc(func(returnVals *g.Group) {
				if child {
					returnVals.Panic(g.Err())
				} else {
					returnVals.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit(fmt.Sprintf("error evaluating id expression for %q workflow: %%w", workflow)), g.Err()))
				}
			})
			b.Id("opts").Dot(idFieldName).Op("=").Id("id")
		})
	}

	// set default id reuse policy
	var idReusePolicy string
	switch opts.GetIdReusePolicy() {
	case temporalv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE:
		idReusePolicy = "WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE"
	case temporalv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY:
		idReusePolicy = "WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY"
	case temporalv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE:
		idReusePolicy = "WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE"
	case temporalv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING:
		idReusePolicy = "WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING"
	}

	if idReusePolicy != "" {
		fn.If(g.Id("opts").Dot("WorkflowIDReusePolicy").Op("==").Qual(enumsPkg, "WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED")).
			Block(
				g.Id("opts").Dot("WorkflowIDReusePolicy").Op("=").Qual(enumsPkg, idReusePolicy),
			)
	}

	if policy := opts.GetRetryPolicy(); policy != nil {
		fn.If(g.Id("opts").Dot("RetryPolicy").Op("==").Nil()).Block(
			g.Id("opts").Dot("RetryPolicy").Op("=").Op("&").Qual(temporalPkg, "RetryPolicy").ValuesFunc(func(fields *g.Group) {
				if d := policy.GetInitialInterval(); d.IsValid() {
					fields.Id("InitialInterval").Op(":").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10))
				}
				if d := policy.GetMaxInterval(); d.IsValid() {
					fields.Id("MaximumInterval").Op(":").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10))
				}
				if n := policy.GetBackoffCoefficient(); n != 0 {
					fields.Id("BackoffCoefficient").Op(":").Lit(n)
				}
				if n := policy.GetMaxAttempts(); n != 0 {
					fields.Id("MaximumAttempts").Op(":").Lit(n)
				}
				if errs := policy.GetNonRetryableErrorTypes(); len(errs) > 0 {
					fields.Id("NonRetryableErrorTypes").Op(":").Lit(errs)
				}
			}),
		)
	}

	if timeout := opts.GetExecutionTimeout(); timeout.IsValid() {
		fn.If(g.Id("opts").Dot("WorkflowExecutionTimeout").Op("==").Lit(0)).
			Block(
				g.Id("opts").Dot("WorkflowRunTimeout").Op("=").Id(strconv.FormatInt(timeout.AsDuration().Nanoseconds(), 10)).Comment(timeout.AsDuration().String()),
			)
	}

	if timeout := opts.GetRunTimeout(); timeout.IsValid() {
		fn.If(g.Id("opts").Dot("WorkflowRunTimeout").Op("==").Lit(0)).
			Block(
				g.Id("opts").Dot("WorkflowRunTimeout").Op("=").Id(strconv.FormatInt(timeout.AsDuration().Nanoseconds(), 10)).Comment(timeout.AsDuration().String()),
			)
	}

	if timeout := opts.GetTaskTimeout(); timeout.IsValid() {
		fn.If(g.Id("opts").Dot("WorkflowTaskTimeout").Op("==").Lit(0)).
			Block(
				g.Id("opts").Dot("WorkflowRunTimeout").Op("=").Id(strconv.FormatInt(timeout.AsDuration().Nanoseconds(), 10)).Comment(timeout.AsDuration().String()),
			)
	}

	if mapping := opts.GetSearchAttributes(); mapping != "" {
		fn.If(g.Id("opts").Dot("SearchAttributes").Op("==").Nil()).
			BlockFunc(func(bl *g.Group) {
				// initalize mapping input
				if hasInput {
					bl.List(g.Id("structured"), g.Err()).Op(":=").Qual(expressionPkg, "ToStructured").Call(g.Id("req").Dot("ProtoReflect").Call())
					bl.If(g.Err().Op("!=").Nil()).Block(
						g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit(fmt.Sprintf("error serializing input for %q search attribute mapping: %%v", workflow)), g.Err())),
					)
				} else {
					bl.Var().Id("structured").Any()
				}

				bl.List(g.Id("result"), g.Err()).Op(":=").Id(toCamel("%sSearchAttributesMapping", workflow)).Dot("Query").Call(g.Id("structured"))
				bl.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit(fmt.Sprintf("error executing %q search attribute mapping: %%v", workflow)), g.Err())),
				)
				bl.List(g.Id("searchAttributes"), g.Id("ok")).Op(":=").Id("result").Op(".").Parens(g.Map(g.String()).Interface())
				bl.If(g.Op("!").Id("ok")).Block(
					g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit(fmt.Sprintf("expected %q search attribute mapping to return map[string]any, got: %%T", workflow)), g.Id("result"))),
				)
				bl.Id("opts").Dot("SearchAttributes").Op("=").Id("searchAttributes")
			})
	}

	// add child workflow default options
	if child {
		ns := opts.GetNamespace()
		if ns == "" {
			ns = svc.opts.GetNamespace()
		}
		if ns != "" {
			fn.If(g.Id("opts").Dot("Namespace").Op("==").Lit("")).Block(
				g.Id("opts").Dot("Namespace").Op("=").Lit(ns),
			)
		}

		var parentClosePolicy string
		switch opts.GetParentClosePolicy() {
		case temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_ABANDON:
			parentClosePolicy = "PARENT_CLOSE_POLICY_ABANDON"
		case temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL:
			parentClosePolicy = "PARENT_CLOSE_POLICY_REQUEST_CANCEL"
		case temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE:
			parentClosePolicy = "PARENT_CLOSE_POLICY_TERMINATE"
		}
		if parentClosePolicy != "" {
			fn.If(g.Id("opts").Dot("ParentClosePolicy").Op("==").Qual(enumsPkg, "PARENT_CLOSE_POLICY_UNSPECIFIED")).Block(
				g.Id("opts").Dot("ParentClosePolicy").Op("=").Qual(enumsPkg, parentClosePolicy),
			)
		}

		if opts.GetWaitForCancellation() {
			fn.Id("opts").Dot("WaitForCancellation").Op("=").Lit(true)
		}
	}
}

// genClientUpdateHandleImpl generates a <Update>Handle struct
func (svc *Service) genClientUpdateHandleImpl(f *g.File, update string) {
	clientImplType := toLowerCamel("%sClient", svc.Service.GoName)
	typeName := toLowerCamel("%sHandle", update)
	interfaceName := toCamel("%sHandle", update)

	// generate struct
	f.Commentf("%s provides an internal implementation of a(n) %s", typeName, interfaceName)
	f.Type().
		Id(typeName).
		Struct(
			g.Id("client").Op("*").Id(clientImplType),
			g.Id("handle").Qual(clientPkg, "WorkflowUpdateHandle"),
		)
}

// genClientUpdateHandleImplGetMethod generates a <UpdateHandle>'s Get method
func (svc *Service) genClientUpdateHandleImplGetMethod(f *g.File, update string) {
	typeName := toLowerCamel("%sHandle", update)
	method := svc.methods[update]
	hasOutput := !isEmpty(method.Output)

	f.Comment("Get blocks until the update wait policy is met, returning the result if applicable")
	f.Func().
		Params(g.Id("h").Op("*").Id(typeName)).
		Id("Get").
		Params(g.Id("ctx").Qual("context", "Context")).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(method.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			if hasOutput {
				fn.Var().Id("resp").Id(method.Output.GoIdent.GoName)
				fn.If(
					g.Err().Op(":=").Id("h").Dot("handle").Dot("Get").Call(
						g.Id("ctx"),
						g.Op("&").Id("resp"),
					),
					g.Err().Op("!=").Nil(),
				).Block(
					g.Return(
						g.Nil(), g.Err(),
					),
				)
				fn.Return(
					g.Op("&").Id("resp"), g.Nil(),
				)
			} else {
				fn.Return(
					g.Id("h").Dot("handle").Dot("Get").Call(
						g.Id("ctx"),
						g.Nil(),
					),
				)
			}
		})
}

// genClientUpdateHandleImplRunIDMethod generates a <UpdateHandle>'s RunID method
func (svc *Service) genClientUpdateHandleImplRunIDMethod(f *g.File, update string) {
	typeName := toLowerCamel("%sHandle", update)

	f.Comment("RunID returns the execution ID")
	f.Func().
		Params(g.Id("h").Op("*").Id(typeName)).
		Id("RunID").
		Params().
		String().
		Block(
			g.Return(g.Id("h").Dot("handle").Dot("RunID").Call()),
		)
}

// genClientUpdateHandleImplUpdateIDMethod generates a <UpdateHandle>'s UpdateID method
func (svc *Service) genClientUpdateHandleImplUpdateIDMethod(f *g.File, update string) {
	typeName := toLowerCamel("%sHandle", update)

	f.Comment("UpdateID returns the update ID")
	f.Func().
		Params(g.Id("h").Op("*").Id(typeName)).
		Id("UpdateID").
		Params().
		String().
		Block(
			g.Return(g.Id("h").Dot("handle").Dot("UpdateID").Call()),
		)
}

// genClientUpdateHandleImplWorkflowIDMethod generates a <UpdateHandle>'s WorkflowID method
func (svc *Service) genClientUpdateHandleImplWorkflowIDMethod(f *g.File, update string) {
	typeName := toLowerCamel("%sHandle", update)

	f.Comment("WorkflowID returns the workflow ID")
	f.Func().
		Params(g.Id("h").Op("*").Id(typeName)).
		Id("WorkflowID").
		Params().
		String().
		Block(
			g.Return(g.Id("h").Dot("handle").Dot("WorkflowID").Call()),
		)
}

// genClientUpdateHandleInterface generates a <Workflow>Run interface
func (svc *Service) genClientUpdateHandleInterface(f *g.File, update string) {
	typeName := toCamel("%sHandle", update)
	method := svc.methods[update]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s describes a(n) %s update handle", typeName, svc.fqnForUpdate(update))
	f.Type().Id(typeName).InterfaceFunc(func(methods *g.Group) {
		methods.Comment("WorkflowID returns the workflow ID")
		methods.Id("WorkflowID").Params().String()

		methods.Comment("RunID returns the workflow instance ID")
		methods.Id("RunID").Params().String()

		methods.Comment("UpdateID returns the update ID")
		methods.Id("UpdateID").Params().String()

		methods.Comment("Get blocks until the workflow is complete and returns the result")
		methods.Id("Get").
			Params(g.Id("ctx").Qual("context", "Context")).
			ParamsFunc(func(returnVals *g.Group) {
				if hasOutput {
					returnVals.Op("*").Id(method.Output.GoIdent.GoName)
				}
				returnVals.Error()
			})
	})
}

// genClientUpdateOptions generates a <Update>Options struct
func (svc *Service) genClientUpdateOptions(f *g.File, update string) {
	typeName := toCamel("%sOptions", update)
	constructorName := "New" + typeName

	f.Commentf("%s provides configuration for a %s update operation", typeName, svc.fqnForUpdate(update))
	f.Type().Id(typeName).Struct(
		g.Id("opts").Op("*").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest"),
	)

	f.Commentf("%s initializes a new %s value", constructorName, typeName)
	f.Func().Id(constructorName).Params().Op("*").Id(typeName).Block(
		g.Return(g.Op("&").Id(typeName).Values(
			g.Id("opts").Op(":").Op("&").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest").Values(),
		)),
	)

	f.Comment("WithUpdateWorkflowOptions sets the initial client.UpdateWorkflowWithOptionsRequest")
	f.Func().
		Params(g.Id("opts").Op("*").Id(typeName)).
		Id("WithUpdateWorkflowOptions").
		Params(g.Id("options").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest")).
		Op("*").Id(typeName).
		Block(
			g.Id("opts").Dot("opts").Op("=").Op("&").Id("options"),
			g.Return(g.Id("opts")),
		)
}

func (svc *Service) genClientUpdateWorkflowOptions(fn *g.Group, update string) {
	updateOpts := svc.updates[update]
	handler := svc.methods[update]
	hasInput := !isEmpty(handler.Input)

	// initialize update request options
	fn.Id("options").Op(":=").Op("&").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest").Values()
	fn.If(g.Len(g.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(g.Lit(0)).Dot("opts").Op("!=").Nil()).Block(
		g.Id("options").Op("=").Id("opts").Index(g.Lit(0)).Dot("opts"),
	)

	// add request args if update has inpute
	if hasInput {
		fn.Id("options").Dot("Args").Op("=").Index().Any().Values(g.Id("req"))
	}
	// add run id if specified
	fn.Id("options").Dot("RunID").Op("=").Id("runID")
	fn.Id("options").Dot("UpdateName").Op("=").Id(toCamel("%sUpdateName", update))
	fn.Id("options").Dot("WorkflowID").Op("=").Id("workflowID")

	// add update id if specified
	if idExpr := updateOpts.GetId(); idExpr != "" {
		fn.If(g.Id("options").Dot("UpdateID").Op("==").Lit("")).BlockFunc(func(b *g.Group) {
			b.List(g.Id("id"), g.Err()).Op(":=").Qual(expressionPkg, "EvalExpression").CallFunc(func(args *g.Group) {
				args.Id(toCamel("%sIDExpression", update))
				if hasInput {
					args.Id("req").Dot("ProtoReflect").Call()
				} else {
					args.Nil()
				}
			})
			b.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error evaluating %s id expression: %w"), g.Id(fmt.Sprintf("%sUpdateName", update)), g.Err())),
			)
			b.Id("options").Dot("UpdateID").Op("=").Id("id")
		})
	}

	// add default wait policy if specified
	if wp := updateOpts.GetWaitPolicy(); wp != temporalv1.WaitPolicy_WAIT_POLICY_UNSPECIFIED {
		var stage string
		switch wp {
		case temporalv1.WaitPolicy_WAIT_POLICY_ACCEPTED:
			stage = "UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ADMITTED"
		case temporalv1.WaitPolicy_WAIT_POLICY_ADMITTED:
			stage = "UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ACCEPTED"
		case temporalv1.WaitPolicy_WAIT_POLICY_COMPLETED:
			stage = "UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_COMPLETED"
		}
		fn.If(g.Id("options").Dot("WaitPolicy").Dot("GetLifecycleStage").Call().Op("==").Qual(enumsPkg, "UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_UNSPECIFIED")).Block(
			g.Id("options").Dot("WaitPolicy").Op("=").Op("&").Qual(updatePkg, "WaitPolicy").Values(
				g.Id("LifecycleStage").Op(":").Qual(enumsPkg, stage),
			),
		)
	}
}

// genClientWorkflowOptions generates a <Workflow>Options struct
func (svc *Service) genClientWorkflowOptions(f *g.File, workflow string) {
	typeName := toCamel("%sOptions", workflow)
	constructorName := "New" + typeName

	f.Commentf("%s provides configuration for a %s workflow operation", typeName, svc.fqnForWorkflow(workflow))
	f.Type().Id(typeName).Struct(
		g.Id("opts").Op("*").Qual(clientPkg, "StartWorkflowOptions"),
	)

	f.Commentf("%s initializes a new %s value", constructorName, typeName)
	f.Func().Id(constructorName).Params().Op("*").Id(typeName).Block(
		g.Return(g.Op("&").Id(typeName).Values()),
	)

	f.Comment("WithStartWorkflowOptions sets the initial client.StartWorkflowOptions")
	f.Func().
		Params(g.Id("opts").Op("*").Id(typeName)).
		Id("WithStartWorkflowOptions").
		Params(g.Id("options").Qual(clientPkg, "StartWorkflowOptions")).
		Op("*").Id(typeName).
		Block(
			g.Id("opts").Dot("opts").Op("=").Op("&").Id("options"),
			g.Return(g.Id("opts")),
		)
}

// genClientWorkflowRunImpl generates a <Workflow>Run struct
func (svc *Service) genClientWorkflowRunImpl(f *g.File, workflow string) {
	clientType := toLowerCamel("%sClient", svc.Service.GoName)
	typeName := toLowerCamel("%sRun", workflow)
	interfaceName := toCamel("%sRun", workflow)

	// generate struct
	f.Commentf("%s provides an internal implementation of a(n) %sRun", typeName, interfaceName)
	f.Type().
		Id(typeName).
		Struct(
			g.Id("client").Op("*").Id(clientType),
			g.Id("run").Qual(clientPkg, "WorkflowRun"),
		)
}

// genClientWorkflowRunImplCancelMethod generates a <Workflow>Run's Cancel method
func (svc *Service) genClientWorkflowRunImplCancelMethod(f *g.File, workflow string) {
	typeName := toLowerCamel("%sRun", workflow)

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

// genClientWorkflowRunImplGetMethod generates a <Workflow>Run's Get method
func (svc *Service) genClientWorkflowRunImplGetMethod(f *g.File, workflow string) {
	typeName := toLowerCamel("%sRun", workflow)
	method := svc.methods[workflow]
	hasOutput := !isEmpty(method.Output)

	f.Comment("Get blocks until the workflow is complete, returning the result if applicable")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Get").
		Params(g.Id("ctx").Qual("context", "Context")).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(method.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			if hasOutput {
				fn.Var().Id("resp").Id(method.Output.GoIdent.GoName)
				fn.If(
					g.Err().Op(":=").Id("r").Dot("run").Dot("Get").Call(
						g.Id("ctx"),
						g.Op("&").Id("resp"),
					),
					g.Err().Op("!=").Nil(),
				).Block(
					g.Return(
						g.Nil(), g.Err(),
					),
				)
				fn.Return(
					g.Op("&").Id("resp"), g.Nil(),
				)
			} else {
				fn.Return(
					g.Id("r").Dot("run").Dot("Get").Call(
						g.Id("ctx"),
						g.Nil(),
					),
				)
			}
		})
}

// genClientWorkflowRunImplIDMethod generates a <Workflow>Run's ID method
func (svc *Service) genClientWorkflowRunImplIDMethod(f *g.File, workflow string) {
	typeName := toLowerCamel("%sRun", workflow)

	f.Comment("ID returns the workflow ID")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("ID").
		Params().
		String().
		Block(
			g.Return(g.Id("r").Dot("run").Dot("GetID").Call()),
		)
}

// genClientWorkflowRunImplQueryMethod generates a <WOrkflow>Run's <Query> method
func (svc *Service) genClientWorkflowRunImplQueryMethod(f *g.File, workflow string, query string) {
	typeName := toLowerCamel("%sRun", workflow)
	handler := svc.methods[query]
	hasInput := !isEmpty(handler.Input)

	f.Commentf("%s executes a(n) %s query", query, svc.fqnForQuery(query))
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id(query).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
			}
		}).
		Params(
			g.Op("*").Id(handler.Output.GoIdent.GoName),
			g.Error(),
		).
		Block(
			g.Return(
				g.Id("r").Dot("client").Dot(query).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("r").Dot("ID").Call()
					args.Lit("")
					if hasInput {
						args.Id("req")
					}
				}),
			),
		)
}

// genClientWorkflowRunImplRunIDMethod generates a <Workflow>Run's RunID method
func (svc *Service) genClientWorkflowRunImplRunIDMethod(f *g.File, workflow string) {
	typeName := toLowerCamel("%sRun", workflow)

	f.Comment("RunID returns the execution ID")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("RunID").
		Params().
		String().
		Block(
			g.Return(g.Id("r").Dot("run").Dot("GetRunID").Call()),
		)
}

// genClientWorkflowRunImplSignalMethod generates a <Workflow>Run's <Signal> method
func (svc *Service) genClientWorkflowRunImplSignalMethod(f *g.File, workflow string, signal string) {
	typeName := toLowerCamel("%sRun", workflow)
	handler := svc.methods[signal]
	hasInput := !isEmpty(handler.Input)

	// generate get method
	f.Commentf("%s sends a(n) %s signal", signal, svc.fqnForSignal(signal))
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id(signal).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
			}
		}).
		Params(g.Error()).
		Block(
			g.Return(
				g.Id("r").Dot("client").Dot(signal).CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("r").Dot("ID").Call()
					args.Lit("")
					if hasInput {
						args.Id("req")
					}
				}),
			),
		)
}

// genClientWorkflowRunImplTerminateMethod generates a <Workflow>Run's Terminate method
func (svc *Service) genClientWorkflowRunImplTerminateMethod(f *g.File, workflow string) {
	typeName := toLowerCamel("%sRun", workflow)

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

// genClientWorkflowRunImplUpdateAsyncMethod generates a <Workflow>Run's <Update>Async method
func (svc *Service) genClientWorkflowRunImplUpdateAsyncMethod(f *g.File, workflow string, update string) {
	typeName := toLowerCamel("%sRun", workflow)
	methodName := toCamel("%sAsync", update)
	handler := svc.methods[update]
	hasInput := !isEmpty(handler.Input)

	// generate get method
	f.Commentf("%s sends a(n) %s update to the workflow", methodName, svc.fqnForUpdate(update))
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
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
				g.Id("r").Dot("client").Dot(toCamel("%sAsync", update)).CallFunc(func(args *g.Group) {
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

// genClientWorkflowRunImplUpdateMethod generates a <Workflow>Run's <Update> method
func (svc *Service) genClientWorkflowRunImplUpdateMethod(f *g.File, workflow string, update string) {
	typeName := toLowerCamel("%sRun", workflow)
	handler := svc.methods[update]
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)

	// generate get method
	f.Commentf("%s executes a(n) %s update", update, svc.fqnForUpdate(update))
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
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

// genClientWorkflowRunInterface generates a <Workflow>Run interface
func (svc *Service) genClientWorkflowRunInterface(f *g.File, workflow string) {
	typeName := toCamel("%sRun", workflow)
	opts := svc.workflows[workflow]
	method := svc.methods[workflow]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s describes a(n) %s workflow run", typeName, svc.fqnForWorkflow(workflow))
	f.Type().Id(typeName).InterfaceFunc(func(methods *g.Group) {
		methods.Comment("ID returns the workflow ID")
		methods.Id("ID").Params().String()

		methods.Comment("RunID returns the workflow instance ID")
		methods.Id("RunID").Params().String()

		methods.Comment("Get blocks until the workflow is complete and returns the result")
		methods.Id("Get").
			Params(g.Id("ctx").Qual("context", "Context")).
			ParamsFunc(func(returnVals *g.Group) {
				if hasOutput {
					returnVals.Op("*").Id(method.Output.GoIdent.GoName)
				}
				returnVals.Error()
			})

		methods.Comment("Cancel requests cancellation of a workflow in execution, returning an error if applicable")
		methods.Id("Cancel").
			Params(g.Id("ctx").Qual("context", "Context")).
			Params(g.Error())

		methods.Comment("Terminate terminates a workflow in execution, returning an error if applicable")
		methods.Id("Terminate").
			Params(
				g.Id("ctx").Qual("context", "Context"),
				g.Id("reason").String(),
				g.Id("details").Op("...").Interface(),
			).
			Params(g.Error())

		for _, queryOpts := range opts.GetQuery() {
			query := queryOpts.GetRef()
			handler := svc.methods[query]
			hasInput := !isEmpty(handler.Input)

			if desc := handler.Comments.Leading.String(); desc != "" {
				methods.Comment(strings.ReplaceAll(strings.TrimPrefix(desc, "//"), "\n//", ""))
			} else {
				methods.Commentf("%s executes a(n) %s query", query, svc.fqnForQuery(query))
			}
			methods.Id(query).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
					}
				}).
				Params(
					g.Op("*").Id(handler.Output.GoIdent.GoName),
					g.Error(),
				)
		}

		for _, signalOpts := range opts.GetSignal() {
			signal := signalOpts.GetRef()
			handler := svc.methods[signal]
			hasInput := !isEmpty(handler.Input)

			if desc := handler.Comments.Leading.String(); desc != "" {
				methods.Comment(strings.ReplaceAll(strings.TrimPrefix(desc, "//"), "\n//", ""))
			} else {
				methods.Commentf("%s sends a(n) %s signal", signal, svc.fqnForSignal(signal))
			}
			methods.Id(signal).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
					}
				}).
				Params(g.Error())
		}

		for _, updateOpts := range opts.GetUpdate() {
			update := updateOpts.GetRef()
			handler := svc.methods[update]
			hasInput := !isEmpty(handler.Input)
			hasOutput := !isEmpty(handler.Output)

			// add synchronous flavor
			if desc := handler.Comments.Leading.String(); desc != "" {
				methods.Comment(strings.ReplaceAll(strings.TrimPrefix(desc, "//"), "\n//", ""))
			} else {
				methods.Commentf("%s executes a(n) %s update", update, svc.fqnForUpdate(update))
			}
			methods.Id(update).
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
				})

			// add async flavor
			methods.Commentf("%sAsync sends a(n) %s update to the workflow", update, svc.fqnForUpdate(update))
			methods.Id(toCamel("%sAsync", update)).
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
				)
		}
	})
}
