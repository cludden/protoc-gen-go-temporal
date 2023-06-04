package plugin

import (
	"fmt"
	"strconv"
	"strings"

	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	g "github.com/dave/jennifer/jen"
	pgs "github.com/lyft/protoc-gen-star/v2"
)

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
				methods.Commentf("%s executes a %s workflow and blocks until error or response received", workflow, workflow)
			}
			methods.Id(workflow).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("opts").Op("*").Qual(clientPkg, "StartWorkflowOptions")
					if hasInput {
						args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
					}
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Id(method.Output.GoIdent.GoName)
					}
					returnVals.Error()
				})

			// generate Execute<Workflow> method
			methods.Commentf("Execute%s executes a %s workflow", workflow, workflow)
			methods.Id(fmt.Sprintf("Execute%s", workflow)).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("opts").Op("*").Qual(clientPkg, "StartWorkflowOptions")
					if hasInput {
						args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
					}
				}).
				Params(
					g.Id(fmt.Sprintf("%sRun", workflow)),
					g.Error(),
				)

			// generate Get<Workflow> method
			methods.Commentf("Get%s retrieves a %s workflow execution", workflow, workflow)
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

			// add Start<Workflow>With<Signal> method
			for _, signalOpts := range opts.GetSignal() {
				if !signalOpts.GetStart() {
					continue
				}
				method := svc.methods[workflow]
				signal := signalOpts.GetRef()
				handler := svc.methods[signal]
				hasWorkflowInput := !isEmpty(method.Input)
				hasSignalInput := !isEmpty(handler.Input)

				methods.Commentf("Start%sWith%s sends a %s signal to a %s workflow, starting it if not present", workflow, signal, signal, workflow)
				methods.Id(fmt.Sprintf("Start%sWith%s", workflow, signal)).
					ParamsFunc(func(args *g.Group) {
						args.Id("ctx").Qual("context", "Context")
						args.Id("opts").Op("*").Qual(clientPkg, "StartWorkflowOptions")
						if hasWorkflowInput {
							args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
						}
						if hasSignalInput {
							args.Id("signal").Op("*").Id(handler.Input.GoIdent.GoName)
						}
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
			methods.Commentf("Query%s sends a %s query to an existing workflow", query, query)
			methods.Id(fmt.Sprintf("Query%s", query)).
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
			methods.Commentf("Signal%s sends a %s signal to an existing workflow", signal, signal)
			methods.Id(fmt.Sprintf("Signal%s", signal)).
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
	})
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
	})
}

// genClientWorkflowRun generates a <Workflow>Run struct
func (svc *Service) genClientWorkflowRun(f *g.File, workflow string) {
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

// genClientWorkflowRunIDMethod generates a <Workflow>Run's ID method
func (svc *Service) genClientWorkflowRunIDMethod(f *g.File, workflow string) {
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

// genClientWorkflowRunRunIDMethod generates a <Workflow>Run's RunID method
func (svc *Service) genClientWorkflowRunRunIDMethod(f *g.File, workflow string) {
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

// genClientWorkflowRunGetMethod generates a <Workflow>Run's Get method
func (svc *Service) genClientWorkflowRunGetMethod(f *g.File, workflow string) {
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

// genClientWorkflowRunQueryMethod generates a <WOrkflow>Run's <Query> method
func (svc *Service) genClientWorkflowRunQueryMethod(f *g.File, workflow string, query string) {
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
				g.Id("r").Dot("client").Dot(fmt.Sprintf("Query%s", query)).CallFunc(func(args *g.Group) {
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

// genClientWorkflowRunSignalMethod generates a <Workflow>Run's <Signal> method
func (svc *Service) genClientWorkflowRunSignalMethod(f *g.File, workflow string, signal string) {
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
				g.Id("r").Dot("client").Dot(fmt.Sprintf("Signal%s", signal)).CallFunc(func(args *g.Group) {
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

// genClient generates the client implementation
func (svc *Service) genClient(f *g.File) {
	f.Comment("Compile-time check that workflowClient satisfies Client")
	f.Var().Op("_").Id("Client").Op("=").Op("&").Id("workflowClient").Block()

	f.Commentf("workflowClient implements a temporal client for a %s service", svc.GoName)
	f.Type().
		Id("workflowClient").
		StructFunc(func(fields *g.Group) {
			fields.Id("client").Qual(clientPkg, "Client")
		})
}

func (svc *Service) genClientConstructor(f *g.File) {
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

// genClientWorkflow generates an <Workflow> client method
func (svc *Service) genClientWorkflow(f *g.File, workflow string) {
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
			args.Id("opts").Op("*").Qual(clientPkg, "StartWorkflowOptions")
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
			// execute workflow
			fn.Id("run").Op(",").Err().Op(":=").Id("c").Dot(fmt.Sprintf("Execute%s", workflow)).CallFunc(func(args *g.Group) {
				args.Id("ctx")
				args.Id("opts")
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
			)
			fn.Return(
				g.Id("run").Dot("Get").Call(g.Id("ctx")),
			)
		})
}

// genClientWorkflowExecute generates an Execute<Workflow> client method
func (svc *Service) genClientWorkflowExecute(f *g.File, workflow string) {
	method := svc.methods[workflow]
	name := pgs.Name(method.GoName).LowerCamelCase().String()
	hasInput := !isEmpty(method.Input)
	f.Commentf("Execute%s starts a %s workflow", workflow, workflow)
	f.Func().
		Params(g.Id("c").Op("*").Id("workflowClient")).
		Id(fmt.Sprintf("Execute%s", workflow)).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("opts").Op("*").Qual(clientPkg, "StartWorkflowOptions")
			if hasInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
		}).
		Params(
			g.Id(fmt.Sprintf("%sRun", workflow)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			// initialize StartWorkflowOptions with defaults
			svc.genStartWorkflowOptions(fn, workflow, false)

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

// genClientWorkflowGet generates a Get<Workflow> client method
func (svc *Service) genClientWorkflowGet(f *g.File, workflow string) {
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

// genClientSignalWithStart adds a Start<Workflow>With<Signal> client method
func (svc *Service) genClientSignalWithStart(f *g.File, workflow, signal string) {
	method := svc.methods[workflow]
	handler := svc.methods[signal]
	name := fmt.Sprintf("Start%sWith%s", workflow, signal)
	runName := pgs.Name(method.GoName).LowerCamelCase().String()
	hasWorkflowInput := !isEmpty(method.Input)
	hasSignalInput := !isEmpty(handler.Input)
	f.Commentf("%s starts a %s workflow and sends a %s signal in a transaction", name, workflow, signal)
	f.Func().
		Params(g.Id("c").Op("*").Id("workflowClient")).
		Id(name).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("opts").Op("*").Qual(clientPkg, "StartWorkflowOptions")
			if hasWorkflowInput {
				args.Id("req").Op("*").Id(method.Input.GoIdent.GoName)
			}
			if hasSignalInput {
				args.Id("signal").Op("*").Id(handler.Input.GoIdent.GoName)
			}
		}).
		Params(
			g.Id(fmt.Sprintf("%sRun", workflow)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			// initialize StartWorkflowOptions
			svc.genStartWorkflowOptions(fn, workflow, false)

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

// genClientQueryMethod adds a <Query> method to a workflowClient
func (svc *Service) genClientQueryMethod(f *g.File, query string) {
	method := svc.methods[query]
	hasInput := !isEmpty(method.Input)
	f.Commentf("Query%s sends a %s query to an existing workflow", query, query)
	f.Func().
		Params(g.Id("c").Op("*").Id("workflowClient")).
		Id(fmt.Sprintf("Query%s", query)).
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

// genClientSignalMethod adds a <Signal> method to a workflowClient
func (svc *Service) genClientSignalMethod(f *g.File, signal string) {
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)
	f.Commentf("Signal%s sends a %s signal to an existing workflow", signal, signal)
	f.Func().
		Params(g.Id("c").Op("*").Id("workflowClient")).
		Id(fmt.Sprintf("Signal%s", signal)).
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

// genStartWorkflowOptions adds logic for initializing StartWorkflowOptions with default values
func (svc *Service) genStartWorkflowOptions(fn *g.Group, workflow string, child bool) {
	method := svc.methods[workflow]
	opts := svc.workflows[workflow]
	hasInput := !isEmpty(method.Input)

	// initialize options if nil
	fn.If(g.Id("opts").Op("==").Nil()).BlockFunc(func(bl *g.Group) {
		if child {
			bl.Id("childOpts").Op(":=").Qual(workflowPkg, "GetChildWorkflowOptions").Call(g.Id("ctx"))
			bl.Id("opts").Op("=").Op("&").Id("childOpts")
		} else {
			bl.Id("opts").Op("=").Op("&").Qual(clientPkg, "StartWorkflowOptions").Block()
		}
	})

	// set task queue if unset and default available
	taskQueue := opts.GetDefaultOptions().GetTaskQueue()
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
	if idExpr := opts.GetDefaultOptions().GetId(); idExpr != "" {
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
	switch opts.GetDefaultOptions().GetIdReusePolicy() {
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

	if policy := opts.GetDefaultOptions().GetRetryPolicy(); policy != nil {
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

	if timeout := opts.GetDefaultOptions().GetExecutionTimeout(); timeout.IsValid() {
		fn.If(g.Id("opts").Dot("WorkflowExecutionTimeout").Op("==").Lit(0)).
			Block(
				g.Id("opts").Dot("WorkflowRunTimeout").Op("=").Id(strconv.FormatInt(timeout.AsDuration().Nanoseconds(), 10)).Comment(timeout.AsDuration().String()),
			)
	}

	if timeout := opts.GetDefaultOptions().GetRunTimeout(); timeout.IsValid() {
		fn.If(g.Id("opts").Dot("WorkflowRunTimeout").Op("==").Lit(0)).
			Block(
				g.Id("opts").Dot("WorkflowRunTimeout").Op("=").Id(strconv.FormatInt(timeout.AsDuration().Nanoseconds(), 10)).Comment(timeout.AsDuration().String()),
			)
	}

	if timeout := opts.GetDefaultOptions().GetTaskTimeout(); timeout.IsValid() {
		fn.If(g.Id("opts").Dot("WorkflowTaskTimeout").Op("==").Lit(0)).
			Block(
				g.Id("opts").Dot("WorkflowRunTimeout").Op("=").Id(strconv.FormatInt(timeout.AsDuration().Nanoseconds(), 10)).Comment(timeout.AsDuration().String()),
			)
	}

	// add child workflow default options
	if child {
		ns := opts.GetDefaultOptions().GetNamespace()
		if ns == "" {
			ns = svc.opts.GetNamespace()
		}
		if ns != "" {
			fn.If(g.Id("opts").Dot("Namespace").Op("==").Lit("")).Block(
				g.Id("opts").Dot("Namespace").Op("=").Lit(ns),
			)
		}

		var parentClosePolicy string
		switch opts.GetDefaultOptions().GetParentClosePolicy() {
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

		if opts.GetDefaultOptions().GetWaitForCancellation() {
			fn.Id("opts").Dot("WaitForCancellation").Op("=").Lit(true)
		}
	}
}
