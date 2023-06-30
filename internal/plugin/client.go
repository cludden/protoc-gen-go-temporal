package plugin

import (
	"fmt"
	"strconv"
	"strings"

	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	g "github.com/dave/jennifer/jen"
	pgs "github.com/lyft/protoc-gen-star/v2"
)

// genClientImpl generates the client implementation
func (svc *Service) genClientImpl(f *g.File) {
	f.Commentf("workflowClient implements a temporal client for a %s service", svc.GoName)
	f.Type().
		Id("workflowClient").
		StructFunc(func(fields *g.Group) {
			fields.Id("client").Qual(clientPkg, "Client")
		})
}

func (svc *Service) genClientImplConstructor(f *g.File) {
	f.Commentf("NewClient initializes a new %s client", svc.GoName)
	f.Func().
		Id("NewClient").
		Params(
			g.Id("c").Qual(clientPkg, "Client"),
		).
		Params(
			g.Id("Client"),
		).
		Block(
			g.Return(
				g.Op("&").Id("workflowClient").Values(
					g.Id("client").Op(":").Id("c"),
				),
			),
		)

	f.Commentf("NewClientWithOptions initializes a new %s client with the given options", svc.GoName)
	f.Func().
		Id("NewClientWithOptions").
		Params(
			g.Id("c").Qual(clientPkg, "Client"),
			g.Id("opts").Qual(clientPkg, "Options"),
		).
		Params(
			g.Id("Client"),
			g.Error(),
		).
		Block(
			g.Var().Err().Error(),
			g.List(g.Id("c"), g.Err()).Op("=").Qual(clientPkg, "NewClientFromExisting").Call(g.Id("c"), g.Id("opts")),
			g.If().Err().Op("!=").Nil().Block(
				g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error initializing client with options: %w"), g.Err())),
			),
			g.Return(
				g.Op("&").Id("workflowClient").Values(
					g.Id("client").Op(":").Id("c"),
				),
				g.Nil(),
			),
		)
}

// genClientImplQueryMethod adds a <Query> method to a workflowClient
func (svc *Service) genClientImplQueryMethod(f *g.File, query string) {
	method := svc.methods[query]
	hasInput := !isEmpty(method.Input)
	f.Commentf("%s sends a %s query to an existing workflow", query, query)
	f.Func().
		Params(g.Id("c").Op("*").Id("workflowClient")).
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
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)
	f.Commentf("%s sends a %s signal to an existing workflow", signal, signal)
	f.Func().
		Params(g.Id("c").Op("*").Id("workflowClient")).
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
	method := svc.methods[workflow]
	handler := svc.methods[signal]
	name := fmt.Sprintf("%sWith%sAsync", workflow, signal)
	runName := pgs.Name(method.GoName).LowerCamelCase().String()
	hasWorkflowInput := !isEmpty(method.Input)
	hasSignalInput := !isEmpty(handler.Input)
	f.Commentf("%s starts a %s workflow and sends a %s signal in a transaction", name, workflow, signal)
	f.Func().
		Params(g.Id("c").Op("*").Id("workflowClient")).
		Id(name).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasWorkflowInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
			if hasSignalInput {
				args.Id("signal").Op("*").Id(handler.Input.GoIdent.GoName)
			}
			args.Id("options").Op("...").Op("*").Qual(clientPkg, "StartWorkflowOptions")
		}).
		Params(
			g.Id(fmt.Sprintf("%sRun", workflow)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			// initialize StartWorkflowOptions
			svc.genClientStartWorkflowOptions(fn, workflow, false)

			// signal with start workflow
			fn.Id("run").Op(",").Err().Op(":=").Id("c").Dot("client").Dot("SignalWithStartWorkflow").CallFunc(func(args *g.Group) {
				args.Id("ctx")
				args.Id("opts").Dot("ID")
				args.Id(fmt.Sprintf("%sSignalName", signal))
				if hasSignalInput {
					args.Id("signal")
				} else {
					args.Nil()
				}
				args.Op("*").Id("opts")
				args.Id(fmt.Sprintf("%sWorkflowName", workflow))
				if hasWorkflowInput {
					args.Id("req")
				}
			})
			fn.If(g.Id("run").Op("==").Nil().Op("||").Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Err()),
			)
			fn.Return(
				g.Op("&").Id(fmt.Sprintf("%sRun", runName)).Block(
					g.Id("client").Op(":").Id("c").Op(","),
					g.Id("run").Op(":").Id("run").Op(","),
				),
				g.Nil(),
			)
		})
}

// genClientImplSignalWithStartMethod adds a Start<Workflow>With<Signal> client method
func (svc *Service) genClientImplSignalWithStartMethod(f *g.File, workflow, signal string) {
	method := svc.methods[workflow]
	handler := svc.methods[signal]
	name := fmt.Sprintf("%sWith%s", workflow, signal)
	hasWorkflowInput := !isEmpty(method.Input)
	hasWorkflowOutput := !isEmpty(method.Output)
	hasSignalInput := !isEmpty(handler.Input)
	f.Commentf("%s starts a %s workflow and sends a %s signal in a transaction", name, workflow, signal)
	f.Func().
		Params(g.Id("c").Op("*").Id("workflowClient")).
		Id(name).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasWorkflowInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
			if hasSignalInput {
				args.Id("signal").Op("*").Id(handler.Input.GoIdent.GoName)
			}
			args.Id("options").Op("...").Op("*").Qual(clientPkg, "StartWorkflowOptions")
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasWorkflowOutput {
				returnVals.Op("*").Id(method.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			// signal with start workflow
			fn.Id("run").Op(",").Err().Op(":=").Id("c").Dot(fmt.Sprintf("%sWith%sAsync", workflow, signal)).CallFunc(func(args *g.Group) {
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
	handler := svc.methods[update]
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)
	f.Commentf("%s sends a(n) %s update to an existing workflow", update, update)
	f.Func().
		Params(g.Id("c").Op("*").Id("workflowClient")).
		Id(update).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
			}
			args.Id("opts").Op("...").Op("*").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest")
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(handler.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		BlockFunc(func(method *g.Group) {
			// initialize update request options
			method.Id("options").Op(":=").Op("&").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest").Values()
			method.If(g.Len(g.Id("opts")).Op(">").Lit(0)).Block(
				g.Id("options").Op("=").Id("opts").Index(g.Lit(0)),
			)

			// override wait policy
			method.Id("options").Dot("WaitPolicy").Op("=").Op("&").Qual(updatePkg, "WaitPolicy").Values(
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
	handler := svc.methods[update]
	handleName := fmt.Sprintf("%sHandle", pgs.Name(handler.GoName).LowerCamelCase().String())
	updateOpts := svc.updates[update]
	hasInput := !isEmpty(handler.Input)
	f.Commentf("%sAsync sends a(n) %s update to an existing workflow", update, update)
	f.Func().
		Params(g.Id("c").Op("*").Id("workflowClient")).
		Id(fmt.Sprintf("%sAsync", update)).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
			}
			args.Id("opts").Op("...").Op("*").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest")
		}).
		Params(
			g.Id(fmt.Sprintf("%sHandle", update)),
			g.Error(),
		).
		BlockFunc(func(method *g.Group) {
			// initialize update request options
			method.Id("options").Op(":=").Op("&").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest").Values()
			method.If(g.Len(g.Id("opts")).Op(">").Lit(0)).Block(
				g.Id("options").Op("=").Id("opts").Index(g.Lit(0)),
			)

			// add request args if update has inpute
			if hasInput {
				method.Id("options").Dot("Args").Op("=").Index().Any().Values(g.Id("req"))
			}
			// add run id if specified
			method.Id("options").Dot("RunID").Op("=").Id("runID")
			method.Id("options").Dot("UpdateName").Op("=").Id(fmt.Sprintf("%sUpdateName", update))
			method.Id("options").Dot("WorkflowID").Op("=").Id("workflowID")

			// add update id if specified
			if idExpr := updateOpts.GetId(); idExpr != "" {
				method.If(g.Id("options").Dot("UpdateID").Op("==").Lit("")).BlockFunc(func(b *g.Group) {
					b.List(g.Id("id"), g.Err()).Op(":=").Qual(expressionPkg, "EvalExpression").CallFunc(func(args *g.Group) {
						args.Id(fmt.Sprintf("%sIDExpression", update))
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
				method.If(g.Id("options").Dot("WaitPolicy").Dot("GetLifecycleStage").Call().Op("==").Qual(enumsPkg, "UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_UNSPECIFIED")).Block(
					g.Id("options").Dot("WaitPolicy").Op("=").Op("&").Qual(updatePkg, "WaitPolicy").Values(
						g.Id("LifecycleStage").Op(":").Qual(enumsPkg, stage),
					),
				)
			}

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
	method := svc.methods[workflow]
	name := pgs.Name(method.GoName).LowerCamelCase().String()
	hasInput := !isEmpty(method.Input)
	f.Commentf("%sAsync starts a %s workflow", workflow, workflow)
	f.Func().
		Params(g.Id("c").Op("*").Id("workflowClient")).
		Id(fmt.Sprintf("%sAsync", workflow)).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
			args.Id("options").Op("...").Op("*").Qual(clientPkg, "StartWorkflowOptions")
		}).
		Params(
			g.Id(fmt.Sprintf("%sRun", workflow)),
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
				g.Op("&").Id(fmt.Sprintf("%sRun", name)).Block(
					g.Id("client").Op(":").Id("c").Op(","),
					g.Id("run").Op(":").Id("run").Op(","),
				),
				g.Nil(),
			)
		})
}

// genClientImplWorkflowGetMethod generates a Get<Workflow> client method
func (svc *Service) genClientImplWorkflowGetMethod(f *g.File, workflow string) {
	method := svc.methods[workflow]
	name := pgs.Name(method.GoName).LowerCamelCase().String()
	f.Commentf("Get%s fetches an existing %s execution", workflow, workflow)
	f.Func().
		Params(g.Id("c").Op("*").Id("workflowClient")).
		Id(fmt.Sprintf("Get%s", workflow)).
		Params(
			g.Id("ctx").Qual("context", "Context"),
			g.Id("workflowID").String(),
			g.Id("runID").String(),
		).
		Params(
			g.Id(fmt.Sprintf("%sRun", workflow)),
			g.Error(),
		).
		Block(
			g.Return(
				g.Op("&").Id(fmt.Sprintf("%sRun", name)).Block(
					g.Id("client").Op(":").Id("c").Op(","),
					g.Id("run").Op(":").Id("c").Dot("client").Dot("GetWorkflow").Call(
						g.Id("ctx"), g.Id("workflowID"), g.Id("runID"),
					).Op(","),
				),
				g.Nil(),
			),
		)
}

// genClientImplWorkflowMethod generates an <Workflow> client method
func (svc *Service) genClientImplWorkflowMethod(f *g.File, workflow string) {
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	if method.Comments.Leading.String() != "" {
		f.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
	} else {
		f.Commentf("%s executes a %s workflow and blocks until error or response received", workflow, workflow)
	}
	f.Func().
		Params(g.Id("c").Op("*").Id("workflowClient")).
		Id(workflow).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
			args.Id("options").Op("...").Op("*").Qual(clientPkg, "StartWorkflowOptions")
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Id(method.Output.GoIdent.GoName)
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			// execute workflow
			fn.Id("run").Op(",").Err().Op(":=").Id("c").Dot(fmt.Sprintf("%sAsync", workflow)).CallFunc(func(args *g.Group) {
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
	f.Commentf("Client describes a client for a %s worker", svc.GoName)
	f.Type().Id("Client").InterfaceFunc(func(methods *g.Group) {
		for _, workflow := range svc.workflowsOrdered {
			opts := svc.workflows[workflow]

			method := svc.methods[workflow]
			hasInput := !isEmpty(method.Input)
			hasOutput := !isEmpty(method.Output)

			// generate <Workflow> method
			if method.Comments.Leading.String() != "" {
				methods.Comment(strings.TrimSuffix(method.Comments.Leading.String(), "\n"))
			} else {
				methods.Commentf("%s executes a(n) %s workflow and blocks until error or response received", workflow, workflow)
			}
			methods.Id(workflow).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
					}
					args.Id("opts").Op("...").Op("*").Qual(clientPkg, "StartWorkflowOptions")
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Id(method.Output.GoIdent.GoName)
					}
					returnVals.Error()
				})

			// generate <Workflow>Async method
			methods.Commentf("%sAsync starts a(n) %s workflow", workflow, workflow)
			methods.Id(fmt.Sprintf("%sAsync", workflow)).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
					}
					args.Id("opts").Op("...").Op("*").Qual(clientPkg, "StartWorkflowOptions")
				}).
				Params(
					g.Id(fmt.Sprintf("%sRun", workflow)),
					g.Error(),
				)

			// generate Get<Workflow> method
			methods.Commentf("Get%s retrieves a(n) existing %s workflow execution", workflow, workflow)
			methods.Id(fmt.Sprintf("Get%s", workflow)).
				Params(
					g.Id("ctx").Qual("context", "Context"),
					g.Id("workflowID").String(),
					g.Id("runID").String(),
				).
				Params(
					g.Id(fmt.Sprintf("%sRun", workflow)),
					g.Error(),
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

				// add synchronous favor
				methods.Commentf("%sWith%s sends a(n) %s signal to a %s workflow, starting it if not present", workflow, signal, signal, workflow)
				methods.Id(fmt.Sprintf("%sWith%s", workflow, signal)).
					ParamsFunc(func(args *g.Group) {
						args.Id("ctx").Qual("context", "Context")
						if hasWorkflowInput {
							args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
						}
						if hasSignalInput {
							args.Id("signal").Op("*").Id(handler.Input.GoIdent.GoName)
						}
						args.Id("opts").Op("...").Op("*").Qual(clientPkg, "StartWorkflowOptions")
					}).
					ParamsFunc(func(returnVals *g.Group) {
						if hasWorkflowOutput {
							returnVals.Op("*").Id(method.Output.GoIdent.GoName)
						}
						returnVals.Error()
					})

				// add async flavor
				methods.Commentf("%sWith%sAsync sends a(n) %s signal to a %s workflow, starting it if not present", workflow, signal, signal, workflow)
				methods.Id(fmt.Sprintf("%sWith%sAsync", workflow, signal)).
					ParamsFunc(func(args *g.Group) {
						args.Id("ctx").Qual("context", "Context")
						if hasWorkflowInput {
							args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
						}
						if hasSignalInput {
							args.Id("signal").Op("*").Id(handler.Input.GoIdent.GoName)
						}
						args.Id("opts").Op("...").Op("*").Qual(clientPkg, "StartWorkflowOptions")
					}).
					Params(
						g.Id(fmt.Sprintf("%sRun", workflow)),
						g.Error(),
					)
			}
		}

		// add <Query> methods
		for _, query := range svc.queriesOrdered {
			handler := svc.methods[query]
			hasInput := !isEmpty(handler.Input)
			methods.Commentf("%s sends a %s query to an existing workflow", query, query)
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
			methods.Commentf("%s sends a %s signal to an existing workflow", signal, signal)
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
			methods.Commentf("%s sends a(n) %s update to an existing workflow", update, update)
			methods.Id(update).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("workflowID").String()
					args.Id("runID").String()
					if hasInput {
						args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
					}
					args.Id("opts").Op("...").Op("*").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest")
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Id(handler.Output.GoIdent.GoName)
					}
					returnVals.Error()
				})

			// add async flavor
			methods.Commentf("%sAsync sends a(n) %s update to an existing workflow", update, update)
			methods.Id(fmt.Sprintf("%sAsync", update)).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("workflowID").String()
					args.Id("runID").String()
					if hasInput {
						args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
					}
					args.Id("opts").Op("...").Op("*").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest")
				}).
				Params(
					g.Id(fmt.Sprintf("%sHandle", update)),
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
		fn.If(g.Len(g.Id("options")).Op(">").Lit(0)).
			Block(
				g.Id("opts").Op("=").Id("options").Index(g.Lit(0)),
			).
			Else().
			Block(
				g.Id("childOpts").Op(":=").Qual(workflowPkg, "GetChildWorkflowOptions").Call(g.Id("ctx")),
				g.Id("opts").Op("=").Op("&").Id("childOpts"),
			)
	} else {
		fn.Id("opts").Op(":=").Op("&").Qual(clientPkg, "StartWorkflowOptions").Values()
		fn.If(g.Len(g.Id("options")).Op(">").Lit(0)).Block(
			g.Id("opts").Op("=").Id("options").Index(g.Lit(0)),
		)
	}

	// set task queue if unset and default available
	taskQueue := opts.GetTaskQueue()
	if taskQueue == "" {
		taskQueue = svc.opts.GetTaskQueue()
	}
	if taskQueue != "" {
		fn.If(g.Id("opts").Dot("TaskQueue").Op("==").Lit("")).Block(
			g.Id("opts").Dot("TaskQueue").Op("=").Lit(taskQueue),
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
				args.Id(fmt.Sprintf("%sIDExpression", workflow))
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
					returnVals.Return(g.Nil(), g.Err())
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
	method := svc.methods[update]
	name := pgs.Name(method.GoName).LowerCamelCase().String()

	// generate struct
	f.Commentf("%sHandle provides an internal implementation of a %sHandle", name, update)
	f.Type().
		Id(fmt.Sprintf("%sHandle", name)).
		Struct(
			g.Id("client").Op("*").Id("workflowClient"),
			g.Id("handle").Qual(clientPkg, "WorkflowUpdateHandle"),
		)
}

// genClientUpdateHandleImplGetMethod generates a <UpdateHandle>'s Get method
func (svc *Service) genClientUpdateHandleImplGetMethod(f *g.File, workflow string) {
	method := svc.methods[workflow]
	name := pgs.Name(method.GoName).LowerCamelCase().String()
	hasOutput := !isEmpty(method.Output)
	f.Comment("Get blocks until the update wait policy is met, returning the result if applicable")
	f.Func().
		Params(g.Id("h").Op("*").Id(fmt.Sprintf("%sHandle", name))).
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
func (svc *Service) genClientUpdateHandleImplRunIDMethod(f *g.File, workflow string) {
	method := svc.methods[workflow]
	name := pgs.Name(method.GoName).LowerCamelCase().String()

	f.Comment("RunID returns the execution ID")
	f.Func().
		Params(g.Id("h").Op("*").Id(fmt.Sprintf("%sHandle", name))).
		Id("RunID").
		Params().
		String().
		Block(
			g.Return(g.Id("h").Dot("handle").Dot("RunID").Call()),
		)
}

// genClientUpdateHandleImplUpdateIDMethod generates a <UpdateHandle>'s UpdateID method
func (svc *Service) genClientUpdateHandleImplUpdateIDMethod(f *g.File, workflow string) {
	method := svc.methods[workflow]
	name := pgs.Name(method.GoName).LowerCamelCase().String()

	f.Comment("UpdateID returns the update ID")
	f.Func().
		Params(g.Id("h").Op("*").Id(fmt.Sprintf("%sHandle", name))).
		Id("UpdateID").
		Params().
		String().
		Block(
			g.Return(g.Id("h").Dot("handle").Dot("UpdateID").Call()),
		)
}

// genClientUpdateHandleImplWorkflowIDMethod generates a <UpdateHandle>'s WorkflowID method
func (svc *Service) genClientUpdateHandleImplWorkflowIDMethod(f *g.File, workflow string) {
	method := svc.methods[workflow]
	name := pgs.Name(method.GoName).LowerCamelCase().String()
	f.Comment("WorkflowID returns the workflow ID")
	f.Func().
		Params(g.Id("h").Op("*").Id(fmt.Sprintf("%sHandle", name))).
		Id("WorkflowID").
		Params().
		String().
		Block(
			g.Return(g.Id("h").Dot("handle").Dot("WorkflowID").Call()),
		)
}

// genClientUpdateHandleInterface generates a <Workflow>Run interface
func (svc *Service) genClientUpdateHandleInterface(f *g.File, update string) {
	method := svc.methods[update]
	hasOutput := !isEmpty(method.Output)
	f.Commentf("%sHandle describes a(n) %s update handle", update, update)
	f.Type().Id(fmt.Sprintf("%sHandle", update)).InterfaceFunc(func(methods *g.Group) {
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

// genClientWorkflowRunImpl generates a <Workflow>Run struct
func (svc *Service) genClientWorkflowRunImpl(f *g.File, workflow string) {
	method := svc.methods[workflow]
	name := pgs.Name(method.GoName).LowerCamelCase().String()

	// generate struct
	f.Commentf("%sRun provides an internal implementation of a %sRun", name, workflow)
	f.Type().
		Id(fmt.Sprintf("%sRun", name)).
		Struct(
			g.Id("client").Op("*").Id("workflowClient"),
			g.Id("run").Qual(clientPkg, "WorkflowRun"),
		)
}

// genClientWorkflowRunImplGetMethod generates a <Workflow>Run's Get method
func (svc *Service) genClientWorkflowRunImplGetMethod(f *g.File, workflow string) {
	method := svc.methods[workflow]
	name := pgs.Name(method.GoName).LowerCamelCase().String()
	hasOutput := !isEmpty(method.Output)
	f.Comment("Get blocks until the workflow is complete, returning the result if applicable")
	f.Func().
		Params(g.Id("r").Op("*").Id(fmt.Sprintf("%sRun", name))).
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
	method := svc.methods[workflow]
	name := pgs.Name(method.GoName).LowerCamelCase().String()
	f.Comment("ID returns the workflow ID")
	f.Func().
		Params(g.Id("r").Op("*").Id(fmt.Sprintf("%sRun", name))).
		Id("ID").
		Params().
		String().
		Block(
			g.Return(g.Id("r").Dot("run").Dot("GetID").Call()),
		)
}

// genClientWorkflowRunImplQueryMethod generates a <WOrkflow>Run's <Query> method
func (svc *Service) genClientWorkflowRunImplQueryMethod(f *g.File, workflow string, query string) {
	method := svc.methods[workflow]
	name := pgs.Name(method.GoName).LowerCamelCase().String()
	handler := svc.methods[query]
	hasInput := !isEmpty(handler.Input)

	f.Commentf("%s executes a %s query against the workflow", query, query)
	f.Func().
		Params(g.Id("r").Op("*").Id(fmt.Sprintf("%sRun", name))).
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
	method := svc.methods[workflow]
	name := pgs.Name(method.GoName).LowerCamelCase().String()

	f.Comment("RunID returns the execution ID")
	f.Func().
		Params(g.Id("r").Op("*").Id(fmt.Sprintf("%sRun", name))).
		Id("RunID").
		Params().
		String().
		Block(
			g.Return(g.Id("r").Dot("run").Dot("GetRunID").Call()),
		)
}

// genClientWorkflowRunImplSignalMethod generates a <Workflow>Run's <Signal> method
func (svc *Service) genClientWorkflowRunImplSignalMethod(f *g.File, workflow string, signal string) {
	method := svc.methods[workflow]
	name := pgs.Name(method.GoName).LowerCamelCase().String()
	handler := svc.methods[signal]
	hasInput := !isEmpty(handler.Input)

	// generate get method
	f.Commentf("%s sends a %s signal to the workflow", signal, signal)
	f.Func().
		Params(g.Id("r").Op("*").Id(fmt.Sprintf("%sRun", name))).
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

// genClientWorkflowRunImplUpdateAsyncMethod generates a <Workflow>Run's <Update>Async method
func (svc *Service) genClientWorkflowRunImplUpdateAsyncMethod(f *g.File, workflow string, update string) {
	method := svc.methods[workflow]
	name := pgs.Name(method.GoName).LowerCamelCase().String()
	handler := svc.methods[update]
	hasInput := !isEmpty(handler.Input)

	// generate get method
	f.Commentf("%sAsync sends a(n) %s update to the workflow", update, update)
	f.Func().
		Params(g.Id("r").Op("*").Id(fmt.Sprintf("%sRun", name))).
		Id(fmt.Sprintf("%sAsync", update)).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
			}
			args.Id("opts").Op("...").Op("*").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest")
		}).
		Params(
			g.Id(fmt.Sprintf("%sHandle", update)),
			g.Error(),
		).
		Block(
			g.Return(
				g.Id("r").Dot("client").Dot(fmt.Sprintf("%sAsync", update)).CallFunc(func(args *g.Group) {
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
	method := svc.methods[workflow]
	name := pgs.Name(method.GoName).LowerCamelCase().String()
	handler := svc.methods[update]
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)

	// generate get method
	f.Commentf("%s sends a(n) %s update to the workflow", update, update)
	f.Func().
		Params(g.Id("r").Op("*").Id(fmt.Sprintf("%sRun", name))).
		Id(update).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
			}
			args.Id("opts").Op("...").Op("*").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest")
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
	opts := svc.workflows[workflow]
	method := svc.methods[workflow]
	hasOutput := !isEmpty(method.Output)
	f.Commentf("%sRun describes a %s workflow run", workflow, workflow)
	f.Type().Id(fmt.Sprintf("%sRun", workflow)).InterfaceFunc(func(methods *g.Group) {
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

		for _, queryOpts := range opts.GetQuery() {
			query := queryOpts.GetRef()
			handler := svc.methods[query]
			hasInput := !isEmpty(handler.Input)
			methods.Commentf("%s runs the %s query against the workflow", query, query)
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
			methods.Commentf("%s sends a %s signal to the workflow", signal, signal)
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
			methods.Commentf("%s sends a(n) %s update to the workflow", update, update)
			methods.Id(update).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
					}
					args.Id("opts").Op("...").Op("*").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest")
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Id(handler.Output.GoIdent.GoName)
					}
					returnVals.Error()
				})

			// add async flavor
			methods.Commentf("%sAsync sends a(n) %s update to the workflow", update, update)
			methods.Id(fmt.Sprintf("%sAsync", update)).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Id(handler.Input.GoIdent.GoName)
					}
					args.Id("opts").Op("...").Op("*").Qual(clientPkg, "UpdateWorkflowWithOptionsRequest")
				}).
				Params(
					g.Id(fmt.Sprintf("%sHandle", update)),
					g.Error(),
				)
		}
	})
}
