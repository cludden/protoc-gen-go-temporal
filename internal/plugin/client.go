package plugin

import (
	"fmt"
	"strconv"

	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	g "github.com/dave/jennifer/jen"
	"github.com/hako/durafmt"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// genClientImpl generates a <service>Client implementation
func (svc *Manifest) genClientImpl(f *g.File) {
	typeName := svc.toLowerCamel("%sClient", svc.Service.GoName)

	f.Commentf("%s implements a temporal client for a %s service", typeName, svc.Service.Desc.FullName())
	f.Type().
		Id(typeName).
		StructFunc(func(fields *g.Group) {
			fields.Id("client").Qual(clientPkg, "Client")
			fields.Id("log").Op("*").Qual("log/slog", "Logger")
		})
}

// genClientImplConstructor generates a New<Service>Client function
func (svc *Manifest) genClientImplConstructor(f *g.File) {
	methodName := svc.toCamel("New%sClient", svc.Service.GoName)
	implName := svc.toLowerCamel("%sClient", svc.Service.GoName)
	interfaceName := svc.toCamel("%sClient", svc.Service.GoName)
	optionsName := svc.toLowerCamel("%sClientOptions", svc.Service.GoName)

	f.Commentf("%s initializes a new %s client", methodName, svc.Service.Desc.FullName())
	f.Func().
		Id(methodName).
		Params(
			g.Id("c").Qual(clientPkg, "Client"),
			g.Id("options").Op("...").Op("*").Id(optionsName),
		).
		Params(
			g.Id(interfaceName),
		).
		Block(
			g.Var().Id("cfg").Op("*").Id(optionsName),
			g.If(g.Len(g.Id("options")).Op(">").Lit(0)).
				Block(
					g.Id("cfg").Op("=").Id("options").Index(g.Lit(0)),
				).
				Else().
				Block(
					g.Id("cfg").Op("=").Id(svc.toCamel("New%sClientOptions", svc.Service.GoName)).Call(),
				),
			g.Return(
				g.Op("&").Id(implName).Custom(
					multiLineValues,
					g.Id("client").Op(":").Id("c"),
					g.Id("log").Op(":").Id("cfg").Dot("getLogger").Call(),
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
			g.Id("options").Op("...").Op("*").Id(optionsName),
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
			g.Var().Id("cfg").Op("*").Id(optionsName),
			g.If(g.Len(g.Id("options")).Op(">").Lit(0)).
				Block(
					g.Id("cfg").Op("=").Id("options").Index(g.Lit(0)),
				).
				Else().
				Block(
					g.Id("cfg").Op("=").Id(svc.toCamel("New%sClientOptions", svc.Service.GoName)).Call(),
				),
			g.Return(
				g.Op("&").Id(implName).Custom(
					multiLineValues,
					g.Id("client").Op(":").Id("c"),
					g.Id("log").Op(":").Id("cfg").Dot("getLogger").Call(),
				),
				g.Nil(),
			),
		)
}

// genClientImplQueryMethod adds a <Query> method to a workflowClient
func (svc *Manifest) genClientImplQueryMethod(f *g.File, query protoreflect.FullName) {
	clientType := svc.toLowerCamel("%sClient", svc.Service.GoName)
	method := svc.methods[query]
	hasInput := !isEmpty(method.Input)

	commentWithDefaultf(f, methodSet(method), "%s sends a(n) %s query to an existing workflow", query, svc.fqnForQuery(query))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(svc.methods[query].GoName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("query").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
		}).
		Params(
			g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			if isDeprecated(method) {
				fn.Id("c").Dot("log").Dot("WarnContext").Call(g.Id("ctx"), g.Lit("use of deprecated client method detected"), g.Lit("method"), g.Lit(svc.toCamel("%s", query)), g.Lit("query"), g.Id(svc.toCamel("%sQueryName", query))).Line()
			}
			fn.Var().Id("resp").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
			fn.If(
				g.List(g.Id("val"), g.Err()).Op(":=").Id("c").Dot("client").Dot("QueryWorkflow").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("workflowID")
					args.Id("runID")
					args.Id(svc.toCamel("%sQueryName", query))
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
			)
			fn.Return(
				g.Op("&").Id("resp"), g.Nil(),
			)
		})
}

// genClientImplSignalMethod adds a <Signal> method to a workflowClient
func (svc *Manifest) genClientImplSignalMethod(f *g.File, signal protoreflect.FullName) {
	clientType := svc.toLowerCamel("%sClient", svc.Service.GoName)
	method := svc.methods[signal]
	hasInput := !isEmpty(method.Input)

	commentWithDefaultf(f, methodSet(method), "%s sends a(n) %s signal to an existing workflow", signal, svc.fqnForSignal(signal))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(svc.methods[signal].GoName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("signal").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
		}).
		Params(g.Error()).
		BlockFunc(func(fn *g.Group) {
			if isDeprecated(method) {
				fn.Id("c").Dot("log").Dot("WarnContext").Call(g.Id("ctx"), g.Lit("use of deprecated client method detected"), g.Lit("method"), g.Lit(svc.toCamel("%s", signal)), g.Lit("signal"), g.Id(svc.toCamel("%sSignalName", signal)))
				fn.Line()
			}
			fn.Return(
				g.Id("c").Dot("client").Dot("SignalWorkflow").CallFunc(func(args *g.Group) {
					args.Id("ctx")
					args.Id("workflowID")
					args.Id("runID")
					args.Id(svc.toCamel("%sSignalName", signal))
					if hasInput {
						args.Id("signal")
					} else {
						args.Nil()
					}
				}),
			)
		})
}

// genClientImplSignalWithStartAsyncMethod adds a <Workflow>With<Signal>Async client method
func (svc *Manifest) genClientImplSignalWithStartAsyncMethod(f *g.File, workflow, signal protoreflect.FullName) {
	clientType := svc.toLowerCamel("%sClient", svc.Service.GoName)
	method := svc.methods[workflow]
	handler := svc.methods[signal]
	name := svc.toCamel("%sWith%sAsync", workflow, signal)
	runName := svc.toLowerCamel("%sRun", workflow)
	hasWorkflowInput := !isEmpty(method.Input)
	hasSignalInput := !isEmpty(handler.Input)

	commentWithDefaultf(f, methodSet(method, handler), "%s starts a(n) %s workflow and sends a(n) %s signal in a transaction", name, svc.fqnForWorkflow(workflow), svc.fqnForSignal(signal))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(name).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasWorkflowInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
			if hasSignalInput {
				args.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
			}
			args.Id("options").Op("...").Op("*").Id(svc.toCamel("%sOptions", workflow))
		}).
		Params(
			g.Id(svc.toCamel("%sRun", workflow)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			if workflowDeprecated, signalDeprecated := isDeprecated(method), isDeprecated(handler); workflowDeprecated || signalDeprecated {
				if workflowDeprecated {
					fn.Id("c").Dot("log").Dot("WarnContext").Call(g.Id("ctx"), g.Lit("use of deprecated client method detected"), g.Lit("method"), g.Lit(name), g.Lit("workflow"), g.Id(svc.toCamel("%sWorkflowName", workflow)))
				}
				if signalDeprecated {
					fn.Id("c").Dot("log").Dot("WarnContext").Call(g.Id("ctx"), g.Lit("use of deprecated client method detected"), g.Lit("method"), g.Lit(name), g.Lit("signal"), g.Id(svc.toCamel("%sSignalName", signal)))
				}
				fn.Line()
			}

			// initialize options
			fn.Var().Id("o").Op("*").Id(svc.toCamel("%sOptions", workflow))
			fn.If(g.Len(g.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(g.Lit(0)).Op("!=").Nil()).Block(
				g.Id("o").Op("=").Id("options").Index(g.Lit(0)),
			).Else().Block(
				g.Id("o").Op("=").Id(svc.toCamel("New%sOptions", workflow)).Call(),
			)

			// initialize client.StartWorkfowOptions
			fn.List(g.Id("opts"), g.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(args *g.Group) {
				if hasWorkflowInput {
					args.Id("req").Dot("ProtoReflect").Call()
				} else {
					args.Nil()
				}
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error initializing client.StartWorkflowOptions: %w"), g.Err())),
			)

			// signal with start workflow
			fn.Id("run").Op(",").Err().Op(":=").Id("c").Dot("client").Dot("SignalWithStartWorkflow").CallFunc(func(args *g.Group) {
				args.Id("ctx")
				args.Id("opts").Dot("ID")
				args.Id(svc.toCamel("%sSignalName", signal))
				if hasSignalInput {
					args.Id("signal")
				} else {
					args.Nil()
				}
				args.Id("opts")
				args.Id(svc.toCamel("%sWorkflowName", workflow))
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
func (svc *Manifest) genClientImplSignalWithStartMethod(f *g.File, workflow, signal protoreflect.FullName) {
	clientType := svc.toLowerCamel("%sClient", svc.Service.GoName)
	method := svc.methods[workflow]
	handler := svc.methods[signal]
	name := svc.toCamel("%sWith%s", workflow, signal)
	hasWorkflowInput := !isEmpty(method.Input)
	hasWorkflowOutput := !isEmpty(method.Output)
	hasSignalInput := !isEmpty(handler.Input)

	commentWithDefaultf(f, methodSet(method, handler), "%s starts a(n) %s workflow and sends a(n) %s signal in a transaction", name, svc.fqnForWorkflow(workflow), svc.fqnForSignal(signal))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(name).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasWorkflowInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
			if hasSignalInput {
				args.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
			}
			args.Id("options").Op("...").Op("*").Id(svc.toCamel("%sOptions", workflow))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasWorkflowOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			if workflowDeprecated, signalDeprecated := isDeprecated(method), isDeprecated(handler); workflowDeprecated || signalDeprecated {
				if workflowDeprecated {
					fn.Id("c").Dot("log").Dot("WarnContext").Call(g.Id("ctx"), g.Lit("use of deprecated client method detected"), g.Lit("method"), g.Lit(name), g.Lit("workflow"), g.Id(svc.toCamel("%sWorkflowName", workflow)))
				}
				if signalDeprecated {
					fn.Id("c").Dot("log").Dot("WarnContext").Call(g.Id("ctx"), g.Lit("use of deprecated client method detected"), g.Lit("method"), g.Lit(name), g.Lit("signal"), g.Id(svc.toCamel("%sSignalName", signal)))
				}
				fn.Line()
			}
			// signal with start workflow
			fn.Id("run").Op(",").Err().Op(":=").Id("c").Dot(svc.toCamel("%sWith%sAsync", workflow, signal)).CallFunc(func(args *g.Group) {
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

func (svc *Manifest) genClientImplUpdateGetMethod(f *g.File, update protoreflect.FullName) {
	methodName := svc.toCamel("Get%s", svc.methods[update].GoName)
	clientType := svc.toLowerCamel("%sClient", svc.Service.GoName)

	f.Commentf("%s retrieves a handle to an existing %s update", methodName, svc.fqnForUpdate(update))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(methodName).
		Params(
			g.Id("ctx").Qual("context", "Context"),
			g.Id("req").Qual(clientPkg, "GetWorkflowUpdateHandleOptions"),
		).
		Params(
			g.Id(svc.toCamel("%sHandle", update)),
			g.Error(),
		).
		Block(
			g.Return(
				g.Op("&").Id(svc.toLowerCamel("%sHandle", update)).Custom(
					multiLineValues,
					g.Id("client").Op(":").Id("c"),
					g.Id("handle").Op(":").Id("c").Dot("client").Dot("GetWorkflowUpdateHandle").Call(
						g.Id("req"),
					),
				),
				g.Nil(),
			),
		)
}

// genClientImplUpdateMethod adds an <Update> method to a workflowClient
func (svc *Manifest) genClientImplUpdateMethod(f *g.File, update protoreflect.FullName) {
	clientType := svc.toLowerCamel("%sClient", svc.Service.GoName)
	handler := svc.methods[update]
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)

	commentWithDefaultf(f, methodSet(handler), "%s sends a(n) %s update to an existing workflow", update, svc.fqnForUpdate(update))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(svc.methods[update].GoName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
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
		BlockFunc(func(fn *g.Group) {
			if isDeprecated(handler) {
				fn.Id("c").Dot("log").Dot("WarnContext").Call(g.Id("ctx"), g.Lit("use of deprecated client method detected"), g.Lit("method"), g.Lit(svc.toCamel("%s", update)), g.Lit("update"), g.Id(svc.toCamel("%sUpdateName", update))).Line()
			}

			// initialize update options
			fn.Comment("initialize update options")
			fn.Id("o").Op(":=").Id(svc.toCamel("New%sOptions", update)).Call()
			fn.If(g.Len(g.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(g.Lit(0)).Dot("Options").Op("!=").Nil()).Block(
				g.Id("o").Op("=").Id("opts").Index(g.Lit(0)),
			)

			fn.Line()
			fn.Comment("call sync update with WorkflowUpdateStageCompleted wait policy")
			fn.List(g.Id("handle"), g.Err()).Op(":=").Id("c").Dot(svc.toCamel("%sAsync", update)).CallFunc(func(args *g.Group) {
				args.Id("ctx")
				args.Id("workflowID")
				args.Id("runID")
				if hasInput {
					args.Id("req")
				}
				args.Id("o").Dot("WithWaitPolicy").Call(g.Qual(clientPkg, "WorkflowUpdateStageCompleted"))
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.ReturnFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)

			fn.Line()
			fn.Comment("block on update completion")
			fn.Return(g.Id("handle").Dot("Get").Call(g.Id("ctx")))
		})
}

// genClientImplUpdateMethodAsync adds an <Update>Async method to a workflowClient
func (svc *Manifest) genClientImplUpdateMethodAsync(f *g.File, update protoreflect.FullName) {
	clientType := svc.toLowerCamel("%sClient", svc.Service.GoName)
	handler := svc.methods[update]
	handleName := svc.toLowerCamel("%sHandle", update)
	hasInput := !isEmpty(handler.Input)
	methodName := svc.toCamel("%sAsync", update)

	commentWithDefaultf(f, methodSet(handler), "%s sends a(n) %s update to an existing workflow", update, svc.fqnForUpdate(update))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(svc.toCamel("%sOptions", update))
		}).
		Params(
			g.Id(svc.toCamel("%sHandle", update)),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			if isDeprecated(handler) {
				fn.Id("c").Dot("log").Dot("WarnContext").Call(g.Id("ctx"), g.Lit("use of deprecated client method detected"), g.Lit("method"), g.Lit(methodName), g.Lit("update"), g.Id(svc.toCamel("%sUpdateName", update))).Line()
			}

			// initialize options
			fn.Comment("initialize update options")
			fn.Var().Id("o").Op("*").Id(svc.toCamel("%sOptions", update))
			fn.If(g.Len(g.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(g.Lit(0)).Op("!=").Nil()).Block(
				g.Id("o").Op("=").Id("opts").Index(g.Lit(0)),
			).Else().Block(
				g.Id("o").Op("=").Id(svc.toCamel("New%sOptions", update)).Call(),
			)

			fn.Line()
			fn.Comment("build UpdateWorkflowOptions")
			fn.List(g.Id("options"), g.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(args *g.Group) {
				args.Id("workflowID")
				args.Id("runID")
				if hasInput {
					args.Id("req")
				}
			})
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error initializing UpdateWorkflowWithOptions: %w"), g.Err())),
			)

			fn.Line()
			fn.Comment("update workflow")
			fn.List(g.Id("handle"), g.Err()).Op(":=").Id("c").Dot("client").Dot("UpdateWorkflow").Call(g.Id("ctx"), g.Op("*").Id("options"))
			fn.If(g.Err().Op("!=").Nil()).Block(
				g.Return(g.Nil(), g.Err()),
			)

			fn.Return(
				g.Op("&").Id(handleName).Values(
					g.Id("client").Op(":").Id("c"),
					g.Id("handle").Op(":").Id("handle"),
				),
				g.Nil(),
			)
		})
}

// genClientImplWorkflowAsyncMethod generates an <Workflow>Async client method
func (svc *Manifest) genClientImplWorkflowAsyncMethod(f *g.File, workflow protoreflect.FullName) {
	clientType := svc.toLowerCamel("%sClient", svc.Service.GoName)
	method := svc.methods[workflow]
	methodName := svc.toCamel("%sAsync", workflow)
	runImplType := svc.toLowerCamel("%sRun", workflow)
	runInterfaceType := svc.toCamel("%sRun", workflow)
	hasInput := !isEmpty(method.Input)
	deprecated := isDeprecated(method)

	commentWithDefaultf(f, methodSet(method), "%s starts a(n) %s workflow and returns a handle to the workflow run", methodName, svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
			args.Id("options").Op("...").Op("*").Id(svc.toCamel("%sOptions", workflow))
		}).
		Params(
			g.Id(runInterfaceType),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			if deprecated {
				fn.Id("c").Dot("log").Dot("WarnContext").Call(g.Id("ctx"), g.Lit("use of deprecated client method detected"), g.Lit("method"), g.Lit(methodName), g.Lit("workflow"), g.Id(svc.toCamel("%sWorkflowName", workflow))).Line()
			}

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

			// execute workflow
			fn.Id("run").Op(",").Err().Op(":=").Id("c").Dot("client").Dot("ExecuteWorkflow").CallFunc(func(args *g.Group) {
				args.Id("ctx")
				args.Id("opts")
				args.Id(svc.toCamel("%sWorkflowName", workflow))
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
func (svc *Manifest) genClientImplWorkflowCancelMethod(f *g.File) {
	clientType := svc.toLowerCamel("%sClient", svc.Service.GoName)
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
func (svc *Manifest) genClientImplWorkflowGetMethod(f *g.File, workflow protoreflect.FullName) {
	clientType := svc.toLowerCamel("%sClient", svc.Service.GoName)
	methodName := svc.toCamel("Get%s", workflow)
	runImplType := svc.toLowerCamel("%sRun", workflow)
	runInterfaceType := svc.toCamel("%sRun", workflow)

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
func (svc *Manifest) genClientImplWorkflowTerminateMethod(f *g.File) {
	clientType := svc.toLowerCamel("%sClient", svc.Service.GoName)
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
func (svc *Manifest) genClientImplWorkflowMethod(f *g.File, workflow protoreflect.FullName) {
	clientType := svc.toLowerCamel("%sClient", svc.Service.GoName)
	method := svc.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	commentWithDefaultf(f, methodSet(method), "%s executes a %s workflow and blocks until error or response received", workflow, svc.fqnForWorkflow(workflow))
	f.Func().
		Params(g.Id("c").Op("*").Id(clientType)).
		Id(svc.methods[workflow].GoName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
			}
			args.Id("options").Op("...").Op("*").Id(svc.toCamel("%sOptions", workflow))
		}).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			// execute workflow
			fn.Id("run").Op(",").Err().Op(":=").Id("c").Dot(svc.toCamel("%sAsync", workflow)).CallFunc(func(args *g.Group) {
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
func (svc *Manifest) genClientInterface(f *g.File) {
	typeName := svc.toCamel("%sClient", svc.Service.GoName)

	f.Commentf("%s describes a client for a(n) %s worker", typeName, svc.Service.Desc.FullName())
	f.Type().Id(typeName).InterfaceFunc(func(methods *g.Group) {
		for _, workflow := range svc.workflowsOrdered {
			if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
				continue
			}
			opts := svc.workflows[workflow]

			method := svc.methods[workflow]
			hasInput := !isEmpty(method.Input)
			hasOutput := !isEmpty(method.Output)
			runInterfaceType := svc.toCamel("%sRun", workflow)

			// generate <Workflow> method
			methodName := svc.toCamel("%s", workflow)
			commentWithDefaultf(methods, methodSet(method), "%s executes a(n) %s workflow and blocks until error or response received", methodName, svc.fqnForWorkflow(workflow))
			methods.Id(methodName).
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
				Line()

			// generate <Workflow>Async method
			methodName = svc.toCamel("%sAsync", workflow)
			commentf(methods, methodSet(method), "%s starts a(n) %s workflow and returns a handle to the workflow run", methodName, svc.fqnForWorkflow(workflow))
			methods.Id(methodName).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
					}
					args.Id("opts").Op("...").Op("*").Id(svc.toCamel("%sOptions", workflow))
				}).
				Params(
					g.Id(runInterfaceType),
					g.Error(),
				).
				Line()

			// generate Get<Workflow> method
			methodName = svc.toCamel("Get%s", workflow)
			commentf(methods, methodSet(method), "%s retrieves a handle to an existing %s workflow execution", methodName, svc.fqnForWorkflow(workflow))
			methods.Id(svc.toCamel("Get%s", workflow)).
				Params(
					g.Id("ctx").Qual("context", "Context"),
					g.Id("workflowID").String(),
					g.Id("runID").String(),
				).
				Params(
					g.Id(runInterfaceType),
				).
				Line()

			// add <Workflow>With<Signal> methods
			for _, signalOpts := range opts.GetSignal() {
				if !signalOpts.GetStart() {
					continue
				}
				method := svc.methods[workflow]
				signal := getFullyQualifiedRef(workflow, signalOpts.GetRef())
				handler := svc.methods[signal]
				hasWorkflowInput := !isEmpty(method.Input)
				hasWorkflowOutput := !isEmpty(method.Output)
				hasSignalInput := !isEmpty(handler.Input)

				// add synchronous flavor
				methodName := svc.toCamel("%sWith%s", workflow, signal)
				commentf(methods, methodSet(method, handler), "%s sends a(n) %s signal to a(n) %s workflow, starting it if necessary, and blocks until workflow completion", methodName, svc.fqnForSignal(signal), svc.fqnForWorkflow(workflow))
				methods.Id(methodName).
					ParamsFunc(func(args *g.Group) {
						args.Id("ctx").Qual("context", "Context")
						if hasWorkflowInput {
							args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
						}
						if hasSignalInput {
							args.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
						}
						args.Id("opts").Op("...").Op("*").Id(svc.toCamel("%sOptions", workflow))
					}).
					ParamsFunc(func(returnVals *g.Group) {
						if hasWorkflowOutput {
							returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
						}
						returnVals.Error()
					}).
					Line()

				// add async flavor
				methodName += "Async"
				commentf(methods, methodSet(method, handler), "%s sends a(n) %s signal to a(n) %s workflow, starting it if necessary, and returns a handle to the workflow execution", methodName, svc.fqnForSignal(signal), svc.fqnForWorkflow(workflow))
				methods.Id(methodName).
					ParamsFunc(func(args *g.Group) {
						args.Id("ctx").Qual("context", "Context")
						if hasWorkflowInput {
							args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
						}
						if hasSignalInput {
							args.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
						}
						args.Id("opts").Op("...").Op("*").Id(svc.toCamel("%sOptions", workflow))
					}).
					Params(
						g.Id(runInterfaceType),
						g.Error(),
					).
					Line()
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
			Params(g.Error()).
			Line()

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
			Params(g.Error()).
			Line()

		// add <Query> methods
		for _, query := range svc.queriesOrdered {
			if svc.methods[query].Desc.Parent() != svc.Service.Desc {
				continue
			}
			handler := svc.methods[query]
			hasInput := !isEmpty(handler.Input)
			commentWithDefaultf(methods, methodSet(handler), "%s executes a(n) %s query", query, svc.fqnForQuery(query))
			methods.Id(svc.toCamel("%s", query)).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("workflowID").String()
					args.Id("runID").String()
					if hasInput {
						args.Id("query").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
					}
				}).
				Params(
					g.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), svc.getMessageName(handler.Output)),
					g.Error(),
				).
				Line()
		}

		// add <Signal> methods
		for _, signal := range svc.signalsOrdered {
			if svc.methods[signal].Desc.Parent() != svc.Service.Desc {
				continue
			}
			handler := svc.methods[signal]
			hasInput := !isEmpty(handler.Input)
			commentWithDefaultf(methods, methodSet(handler), "%s sends a(n) %s signal", signal, svc.fqnForSignal(signal))
			methods.Id(svc.toCamel("%s", signal)).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("workflowID").String()
					args.Id("runID").String()
					if hasInput {
						args.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
					}
				}).
				Params(g.Error()).
				Line()
		}

		// add <Update> methods
		for _, update := range svc.updatesOrdered {
			if svc.methods[update].Desc.Parent() != svc.Service.Desc {
				continue
			}
			handler := svc.methods[update]
			hasInput := !isEmpty(handler.Input)
			hasOutput := !isEmpty(handler.Output)

			// add synchronous flavor
			methodName := svc.toCamel("%s", update)
			commentWithDefaultf(methods, methodSet(handler), "%s executes a(n) %s update and blocks until update completion", methodName, svc.fqnForUpdate(update))
			methods.Id(methodName).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("workflowID").String()
					args.Id("runID").String()
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
				Line()

			// add async flavor
			methodName = svc.toCamel("%sAsync", update)
			commentf(methods, methodSet(handler), "%s starts a(n) %s update and returns a handle to the workflow update", methodName, svc.fqnForUpdate(update))
			methods.Id(methodName).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("workflowID").String()
					args.Id("runID").String()
					if hasInput {
						args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
					}
					args.Id("opts").Op("...").Op("*").Id(svc.toCamel("%sOptions", update))
				}).
				Params(
					g.Id(svc.toCamel("%sHandle", update)),
					g.Error(),
				).
				Line()

			// add getter
			methodName = svc.toCamel("Get%s", update)
			commentf(methods, methodSet(handler), "%s retrieves a handle to an existing %s update", methodName, svc.fqnForUpdate(update))
			methods.Id(methodName).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("req").Qual(clientPkg, "GetWorkflowUpdateHandleOptions")
				}).
				Params(
					g.Id(svc.toCamel("%sHandle", update)),
					g.Error(),
				).
				Line()
		}
	})
}

func (svc *Manifest) genClientOptions(f *g.File) {
	typeName := svc.toLowerCamel("%sClientOptions", svc.Service.GoName)

	f.Commentf("%s describes optional runtime configuration for a %s", typeName, svc.toCamel("%sClient", svc.Service.GoName))
	f.Type().Id(typeName).Struct(
		g.Id("log").Op("*").Qual("log/slog", "Logger"),
	)

	constructorName := svc.toCamel("New%sClientOptions", svc.Service.GoName)
	f.Commentf("%s initializes a new %s value", constructorName, typeName)
	f.Func().Id(constructorName).Params().Op("*").Id(typeName).Block(
		g.Return(g.Op("&").Id(typeName).Values()),
	)

	f.Comment("WithLogger can be used to override the default logger")
	f.Func().
		Params(g.Id("opts").Op("*").Id(typeName)).
		Id("WithLogger").
		Params(g.Id("l").Op("*").Qual("log/slog", "Logger")).
		Op("*").Id(typeName).
		Block(
			g.If(g.Id("l").Op("!=").Nil()).Block(
				g.Id("opts").Dot("log").Op("=").Id("l"),
			),
			g.Return(g.Id("opts")),
		)

	f.Comment("getLogger returns the configured logger, or the default logger")
	f.Func().
		Params(g.Id("opts").Op("*").Id(typeName)).
		Id("getLogger").
		Params().
		Op("*").Qual("log/slog", "Logger").
		Block(
			g.If(g.Id("opts").Op("!=").Nil().Op("&&").Id("opts").Dot("log").Op("!=").Nil()).Block(
				g.Return(g.Id("opts").Dot("log")),
			),
			g.Return(g.Qual("log/slog", "Default").Call()),
		)
}

// genClientUpdateHandleImpl generates a <Update>Handle struct
func (svc *Manifest) genClientUpdateHandleImpl(f *g.File, update protoreflect.FullName) {
	clientImplType := svc.toLowerCamel("%sClient", svc.Service.GoName)
	typeName := svc.toLowerCamel("%sHandle", update)
	interfaceName := svc.toCamel("%sHandle", update)

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
func (svc *Manifest) genClientUpdateHandleImplGetMethod(f *g.File, update protoreflect.FullName) {
	typeName := svc.toLowerCamel("%sHandle", update)
	method := svc.methods[update]
	hasOutput := !isEmpty(method.Output)

	f.Comment("Get blocks until the update wait policy is met, returning the result if applicable")
	f.Func().
		Params(g.Id("h").Op("*").Id(typeName)).
		Id("Get").
		Params(g.Id("ctx").Qual("context", "Context")).
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
			fn.Var().Err().Error()
			fn.Id("doneCh").Op(":=").Make(g.Chan().Struct())
			fn.List(g.Id("gctx"), g.Id("cancel")).Op(":=").Qual("context", "WithCancel").Call(g.Qual("context", "Background").Call())
			fn.Defer().Id("cancel").Call()
			fn.Line()

			fn.Go().Func().Params().Block(
				g.For().Block(
					g.Var().Id("deadlineExceeded").Op("*").Qual(serviceerrorPkg, "DeadlineExceeded"),
					g.If(
						g.Err().Op("=").Id("h").Dot("handle").Dot("Get").CallFunc(func(args *g.Group) {
							args.Id("gctx")
							if hasOutput {
								args.Op("&").Id("resp")
							} else {
								args.Nil()
							}
						}),
						g.Err().Op("!=").Nil().Op("&&").
							Id("ctx").Dot("Err").Call().Op("==").Nil().Op("&&").
							Parens(
								g.Qual("errors", "As").Call(g.Id("err"), g.Op("&").Id("deadlineExceeded")).Op("||").
									Qual("strings", "Contains").Call(
									g.Err().Dot("Error").Call(),
									g.Qual("context", "DeadlineExceeded").Dot("Error").Call(),
								),
							),
					).Block(
						g.Continue(),
					),
					g.Break(),
				),
				g.Close(g.Id("doneCh")),
			).Call()
			fn.Line()

			fn.Select().Block(
				g.Case(g.Op("<-").Id("ctx").Dot("Done").Call()).Block(
					g.ReturnFunc(func(returnVals *g.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Id("ctx").Dot("Err").Call()
					}),
				),
				g.Case(g.Op("<-").Id("doneCh")).BlockFunc(func(bl *g.Group) {
					if hasOutput {
						bl.If(g.Err().Op("!=").Nil()).Block(
							g.Return(g.Nil(), g.Err()),
						)
						bl.Return(g.Op("&").Id("resp"), g.Nil())
					} else {
						bl.Return(g.Err())
					}
				}),
			)
		})
}

// genClientUpdateHandleImplRunIDMethod generates a <UpdateHandle>'s RunID method
func (svc *Manifest) genClientUpdateHandleImplRunIDMethod(f *g.File, update protoreflect.FullName) {
	typeName := svc.toLowerCamel("%sHandle", update)

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
func (svc *Manifest) genClientUpdateHandleImplUpdateIDMethod(f *g.File, update protoreflect.FullName) {
	typeName := svc.toLowerCamel("%sHandle", update)

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
func (svc *Manifest) genClientUpdateHandleImplWorkflowIDMethod(f *g.File, update protoreflect.FullName) {
	typeName := svc.toLowerCamel("%sHandle", update)

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
func (svc *Manifest) genClientUpdateHandleInterface(f *g.File, update protoreflect.FullName) {
	typeName := svc.toCamel("%sHandle", update)
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
					returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
				}
				returnVals.Error()
			})
	})
}

// genClientUpdateOptions generates a <Update>Options struct
func (svc *Manifest) genClientUpdateOptions(f *g.File, update protoreflect.FullName) {
	typeName := svc.toCamel("%sOptions", update)
	constructorName := "New" + typeName
	updateOpts := svc.updates[update]
	hasInput := !isEmpty(svc.methods[update].Input)

	f.Commentf("%s provides configuration for a %s update operation", typeName, svc.fqnForUpdate(update))
	f.Type().Id(typeName).Struct(
		g.Id("Options").Op("*").Qual(clientPkg, "UpdateWorkflowOptions"),
		g.Id("id").Op("*").String(),
		g.Id("waitPolicy").Qual(clientPkg, "WorkflowUpdateStage"),
	)

	f.Commentf("%s initializes a new %s value", constructorName, typeName)
	f.Func().Id(constructorName).Params().Op("*").Id(typeName).Block(
		g.Return(g.Op("&").Id(typeName).Values(
			g.Id("Options").Op(":").Op("&").Qual(clientPkg, "UpdateWorkflowOptions").Values(),
		)),
	)

	f.Comment("Build initializes a new client.UpdateWorkflowOptions with defaults and overrides applied")
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("Build").
		ParamsFunc(func(args *g.Group) {
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(svc.methods[update].Input.GoIdent.GoImportPath), svc.getMessageName(svc.methods[update].Input))
			}
		}).
		Params(
			g.Id("opts").Op("*").Qual(clientPkg, "UpdateWorkflowOptions"),
			g.Err().Error(),
		).
		BlockFunc(func(fn *g.Group) {
			fn.Comment("use user-provided UpdateWorkflowOptions if exists")
			fn.If(g.Id("o").Dot("Options").Op("!=").Nil()).Block(
				g.Id("opts").Op("=").Id("o").Dot("Options"),
			).Else().Block(
				g.Id("opts").Op("=").Op("&").Qual(clientPkg, "UpdateWorkflowOptions").Values(),
			)

			fn.Line()
			fn.Comment("set constants")
			if hasInput {
				fn.Id("opts").Dot("Args").Op("=").Index().Any().Values(g.Id("req"))
			}
			fn.Id("opts").Dot("RunID").Op("=").Id("runID")
			fn.Id("opts").Dot("UpdateName").Op("=").Id(svc.toCamel("%sUpdateName", update))
			fn.Id("opts").Dot("WorkflowID").Op("=").Id("workflowID")

			fn.Line()
			fn.Comment("set UpdateID")
			id := fn.If(g.Id("v").Op(":=").Id("o").Dot("id"), g.Id("v").Op("!=").Nil()).Block(
				g.Id("opts").Dot("UpdateID").Op("=").Op("*").Id("v"),
			)
			if idExpr := updateOpts.GetId(); idExpr != "" {
				id.Else().If(g.Id("opts").Dot("UpdateID").Op("==").Lit("")).BlockFunc(func(b *g.Group) {
					b.List(g.Id("id"), g.Err()).Op(":=").Qual(expressionPkg, "EvalExpression").CallFunc(func(args *g.Group) {
						args.Id(svc.toCamel("%sIDExpression", update))
						if hasInput {
							args.Id("req").Dot("ProtoReflect").Call()
						} else {
							args.Nil()
						}
					})
					b.If(g.Err().Op("!=").Nil()).Block(
						g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error evaluating id expression for %q update: %w"), g.Id(svc.toCamel("%sUpdateName", update)), g.Err())),
					)
					b.Id("opts").Dot("UpdateID").Op("=").Id("id")
				})
			}

			fn.Line()
			fn.Comment("set WaitPolicy")
			waitPolicy := fn.If(g.Id("v").Op(":=").Id("o").Dot("waitPolicy"), g.Id("v").Op("!=").Qual(clientPkg, "WorkflowUpdateStageUnspecified")).Block(
				g.Id("opts").Dot("WaitForStage").Op("=").Id("v"),
			)
			wp := updateOpts.GetWaitForStage()
			if wp == temporalv1.WaitPolicy_WAIT_POLICY_UNSPECIFIED && updateOpts.GetWaitPolicy() != temporalv1.WaitPolicy_WAIT_POLICY_UNSPECIFIED {
				wp = updateOpts.GetWaitPolicy()
			}
			var stage string
			switch wp {
			case temporalv1.WaitPolicy_WAIT_POLICY_ACCEPTED:
				stage = "WorkflowUpdateStageAccepted"
			case temporalv1.WaitPolicy_WAIT_POLICY_ADMITTED:
				stage = "WorkflowUpdateStageAdmitted"
			case temporalv1.WaitPolicy_WAIT_POLICY_COMPLETED, temporalv1.WaitPolicy_WAIT_POLICY_UNSPECIFIED:
				stage = "WorkflowUpdateStageCompleted"
			}
			waitPolicy.Else().If(g.Id("opts").Dot("WaitForStage").Op("==").Qual(clientPkg, "WorkflowUpdateStageUnspecified")).Block(
				g.Id("opts").Dot("WaitForStage").Op("=").Qual(clientPkg, stage),
			)

			fn.Return(g.Id("opts"), g.Nil())
		})

	f.Comment("WithUpdateID sets the UpdateID")
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("WithUpdateID").
		Params(g.Id("id").String()).
		Op("*").Id(typeName).
		Block(
			g.Id("o").Dot("id").Op("=").Op("&").Id("id"),
			g.Return(g.Id("o")),
		)

	f.Comment("WithUpdateWorkflowOptions sets the initial client.UpdateWorkflowOptions")
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("WithUpdateWorkflowOptions").
		Params(g.Id("options").Qual(clientPkg, "UpdateWorkflowOptions")).
		Op("*").Id(typeName).
		Block(
			g.Id("o").Dot("Options").Op("=").Op("&").Id("options"),
			g.Return(g.Id("o")),
		)

	f.Comment("WithWaitPolicy sets the WaitPolicy")
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("WithWaitPolicy").
		Params(g.Id("policy").Qual(clientPkg, "WorkflowUpdateStage")).
		Op("*").Id(typeName).
		Block(
			g.Id("o").Dot("waitPolicy").Op("=").Id("policy"),
			g.Return(g.Id("o")),
		)
}

// genWorkflowOptions generates a <Workflow>Options struct
func (svc *Manifest) genWorkflowOptions(f *g.File, workflow protoreflect.FullName, child bool) {
	optionsPkg := clientPkg
	optionsType := "StartWorkflowOptions"
	typeName := svc.toCamel("%sOptions", workflow)
	childQualifier := ""
	if child {
		optionsPkg = workflowPkg
		optionsType = "ChildWorkflowOptions"
		typeName = svc.toCamel("%sChildOptions", workflow)
		childQualifier = "child "
	}
	constructorName := "New" + typeName
	opts := svc.workflows[workflow]

	f.Commentf("%s provides configuration for a %s%s workflow operation", typeName, childQualifier, svc.fqnForWorkflow(workflow))
	f.Type().Id(typeName).StructFunc(func(values *g.Group) {
		values.Id("options").Qual(optionsPkg, optionsType)
		values.Id("executionTimeout").Op("*").Qual("time", "Duration")
		values.Id("id").Op("*").String()
		values.Id("idReusePolicy").Qual(enumsPkg, "WorkflowIdReusePolicy")
		values.Id("retryPolicy").Op("*").Qual(temporalPkg, "RetryPolicy")
		values.Id("runTimeout").Op("*").Qual("time", "Duration")
		values.Id("searchAttributes").Map(g.String()).Any()
		values.Id("taskQueue").Op("*").String()
		values.Id("taskTimeout").Op("*").Qual("time", "Duration")
		if child {
			values.Id("parentClosePolicy").Qual(enumsPkg, "ParentClosePolicy")
			values.Id("waitForCancellation").Op("*").Bool()
		}
	})

	f.Commentf("%s initializes a new %s value", constructorName, typeName)
	f.Func().Id(constructorName).Params().Op("*").Id(typeName).Block(
		g.Return(g.Op("&").Id(typeName).Values()),
	)

	f.Commentf("Build initializes a new %s.%s value with defaults and overrides applied", optionsPkg, optionsType)
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("Build").
		ParamsFunc(func(args *g.Group) {
			if child {
				args.Id("ctx").Qual(workflowPkg, "Context")
			}
			args.Id("req").Qual(protoreflectPkg, "Message")
		}).
		Params(
			g.Qual(optionsPkg, optionsType),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			fn.Id("opts").Op(":=").Id("o").Dot("options")

			// set ID
			idFieldName := "ID"
			if child {
				idFieldName = "WorkflowID"
			}
			id := fn.If(g.Id("v").Op(":=").Id("o").Dot("id"), g.Id("v").Op("!=").Nil()).Block(
				g.Id("opts").Dot(idFieldName).Op("=").Op("*").Id("v"),
			)
			if idExpr := opts.GetId(); idExpr != "" {
				id.Else().If(g.Id("opts").Dot(idFieldName).Op("==").Lit("")).BlockFunc(func(b *g.Group) {
					// original expression evaluation logic
					origFn := func(b *g.Group, errorReturn g.Code) {
						b.List(g.Id("id"), g.Err()).Op(":=").Qual(expressionPkg, "EvalExpression").CallFunc(func(args *g.Group) {
							args.Id(svc.toCamel("%sIDExpression", workflow))
							args.Id("req")
						})
						b.If(g.Err().Op("!=").Nil()).BlockFunc(func(returnVals *g.Group) {
							returnVals.Return(errorReturn, g.Qual("fmt", "Errorf").Call(g.Lit("error evaluating id expression for %q workflow: %w"), g.Id(svc.toCamel("%sWorkflowName", workflow)), g.Err()))
						})
					}

					if child {
						// local activity wrapper
						fixFn := func(b *g.Group) {
							b.Id("lao").Op(":=").Qual(workflowPkg, "GetLocalActivityOptions").Call(g.Id("ctx"))
							b.Id("lao").Dot("ScheduleToCloseTimeout").Op("=").Qual("time", "Second").Op("*").Lit(10)
							b.If(
								g.Err().Op(":=").Qual(workflowPkg, "ExecuteLocalActivity").Call(
									g.Qual(workflowPkg, "WithLocalActivityOptions").Call(g.Id("ctx"), g.Id("lao")),
									g.Func().Params(g.Id("ctx").Qual("context", "Context")).Params(g.String(), g.Error()).BlockFunc(func(bl *g.Group) {
										origFn(bl, g.Lit(""))
										bl.Return(g.Id("id"), g.Nil())
									}),
								).Dot("Get").Call(g.Id("ctx"), g.Op("&").Id("opts").Dot(idFieldName)),
								g.Err().Op("!=").Nil(),
							).Block(
								g.Return(g.Id("opts"), g.Qual("fmt", "Errorf").Call(g.Lit("error evaluating id expression for %q workflow: %w"), g.Id(svc.toCamel("%sWorkflowName", workflow)), g.Err())),
							)
						}

						// introduce local activity wrapper behind workflow versioning
						switch pvm := svc.patchMode(temporalv1.Patch_PV_64, workflow); pvm {
						case temporalv1.Patch_PVM_ENABLED:
							patchComment(b, temporalv1.Patch_PV_64)
							b.If(g.Add(patchVersion(temporalv1.Patch_PV_64, pvm)).Op("==").Lit(1)).BlockFunc(
								fixFn,
							).Else().BlockFunc(func(b *g.Group) {
								origFn(b, g.Id("opts"))
								b.Id("opts").Dot(idFieldName).Op("=").Id("id")
							})
						case temporalv1.Patch_PVM_MARKER:
							b.Add(patchVersion(temporalv1.Patch_PV_64, pvm))
							fixFn(b)
						case temporalv1.Patch_PVM_REMOVED:
							fixFn(b)
						default:
							origFn(b, g.Id("opts"))
							b.Id("opts").Dot(idFieldName).Op("=").Id("id")
						}
					} else {
						origFn(b, g.Id("opts"))
						b.Id("opts").Dot(idFieldName).Op("=").Id("id")
					}
				})
			}

			// set WorkflowIDReusePolicy
			idReusePolicy := fn.If(g.Id("v").Op(":=").Id("o").Dot("idReusePolicy"), g.Id("v").Op("!=").Qual(enumsPkg, "WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED")).Block(
				g.Id("opts").Dot("WorkflowIDReusePolicy").Op("=").Id("v"),
			)
			if opts.GetIdReusePolicy() != temporalv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
				idReusePolicy.Else().If(g.Id("opts").Dot("WorkflowIDReusePolicy").Op("==").Qual(enumsPkg, "WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED")).BlockFunc(func(bl *g.Group) {
					var policy string
					switch opts.GetIdReusePolicy() {
					case temporalv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE:
						policy = "WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE"
					case temporalv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY:
						policy = "WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY"
					case temporalv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE:
						policy = "WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE"
					case temporalv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING:
						policy = "WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING"
					}
					bl.Id("opts").Dot("WorkflowIDReusePolicy").Op("=").Qual(enumsPkg, policy)
				})
			}

			// set TaskQueue
			{
				var defaultTaskQueue g.Code
				if tq := opts.GetTaskQueue(); tq != "" {
					defaultTaskQueue = g.Lit(tq)
				} else if tq = svc.opts.GetTaskQueue(); tq != "" {
					defaultTaskQueue = g.Id(svc.toCamel("%sTaskQueue", svc.Service.GoName))
				}
				origFn := func(b *g.Group) {
					if defaultTaskQueue != nil { // use default if defined
						b.Id("opts").Dot("TaskQueue").Op("=").Add(defaultTaskQueue)
					} else if child { // otherwise, fallback to parent if child
						b.Id("opts").Dot("TaskQueue").Op("=").Qual(workflowPkg, "GetInfo").Call(g.Id("ctx")).Dot("TaskQueueName")
					} else { // otherwise error
						b.Return(g.Id("opts"), g.Qual("errors", "New").Call(g.Lit("TaskQueue is required")))
					}
				}
				fixFn := func(b *g.Group) {
					if defaultTaskQueue != nil {
						b.
							IfFunc(func(b *g.Group) {
								b.Id("tq").Op(":=").Qual(patchPkg, "DefaultTaskQueue").Call(g.Id("ctx"), defaultTaskQueue)
								b.Id("tq").Op("!=").Lit("").Op("&&").Id("tq").Op("!=").Qual(workflowPkg, "GetInfo").Call(g.Id("ctx")).Dot("TaskQueueName")
							}).
							BlockFunc(func(b *g.Group) {
								b.Id("opts").Dot("TaskQueue").Op("=").Id("tq")
							})
					}
				}

				fn.
					If(g.Id("v").Op(":=").Id("o").Dot("taskQueue"), g.Id("v").Op("!=").Nil()).
					Block(
						g.Id("opts").Dot("TaskQueue").Op("=").Op("*").Id("v"),
					).Else().If(g.Id("opts").Dot("TaskQueue").Op("==").Lit("")).BlockFunc(func(b *g.Group) {
					switch pvm := svc.patchMode(temporalv1.Patch_PV_77, workflow); true {
					case !child || pvm == temporalv1.Patch_PVM_DISABLED:
						origFn(b)
					case pvm == temporalv1.Patch_PVM_ENABLED:
						patchComment(b, temporalv1.Patch_PV_77)
						b.
							If(g.Add(patchVersion(temporalv1.Patch_PV_77, pvm)).Op("==").Lit(1)).
							BlockFunc(fixFn).
							Else().
							BlockFunc(origFn)
					case pvm == temporalv1.Patch_PVM_MARKER:
						patchComment(b, temporalv1.Patch_PV_77)
						patchVersion(temporalv1.Patch_PV_77, pvm)
						fixFn(b)
					case pvm == temporalv1.Patch_PVM_REMOVED:
						patchComment(b, temporalv1.Patch_PV_77)
						fixFn(b)
					}
				})
			}

			// set RetryPolicy
			retryPolicy := fn.If(g.Id("v").Op(":=").Id("o").Dot("retryPolicy"), g.Id("v").Op("!=").Nil()).Block(
				g.Id("opts").Dot("RetryPolicy").Op("=").Id("v"),
			)
			if policy := opts.GetRetryPolicy(); policy != nil {
				retryPolicy.Else().If(g.Id("opts").Dot("RetryPolicy").Op("==").Nil()).Block(
					g.Id("opts").Dot("RetryPolicy").Op("=").Op("&").Qual(temporalPkg, "RetryPolicy").CustomFunc(multiLineValues, func(fields *g.Group) {
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
							fields.Id("NonRetryableErrorTypes").Op(":").Index().String().CustomFunc(multiLineValues, func(vals *g.Group) {
								for _, err := range errs {
									vals.Lit(err)
								}
							})
						}
					}),
				)
			}

			// set SearchAttributes
			searchAttributes := fn.If(g.Id("v").Op(":=").Id("o").Dot("searchAttributes"), g.Id("v").Op("!=").Nil()).Block(
				g.Id("opts").Dot("SearchAttributes").Op("=").Id("o").Dot("searchAttributes"),
			)
			if mapping := opts.GetSearchAttributes(); mapping != "" {
				searchAttributes.Else().If(g.Id("opts").Dot("SearchAttributes").Op("==").Nil()).
					BlockFunc(func(b *g.Group) {
						// original expression evaluation logic
						origFn := func(b *g.Group, errorResult g.Code) {
							// convert input to generic mapping input
							b.List(g.Id("structured"), g.Err()).Op(":=").Qual(expressionPkg, "ToStructured").Call(g.Id("req"))
							b.If(g.Err().Op("!=").Nil()).Block(
								g.Return(errorResult, g.Qual("fmt", "Errorf").Call(g.Lit(fmt.Sprintf("error serializing input for %q search attribute mapping: %%v", svc.methods[workflow].GoName)), g.Err())),
							)

							// execute mapping
							b.List(g.Id("result"), g.Err()).Op(":=").Id(svc.toCamel("%sSearchAttributesMapping", workflow)).Dot("Query").Call(g.Id("structured"))
							b.If(g.Err().Op("!=").Nil()).Block(
								g.Return(errorResult, g.Qual("fmt", "Errorf").Call(g.Lit(fmt.Sprintf("error executing %q search attribute mapping: %%v", svc.methods[workflow].GoName)), g.Err())),
							)

							// coerce mapping result to map[string]any
							b.List(g.Id("searchAttributes"), g.Id("ok")).Op(":=").Id("result").Op(".").Parens(g.Map(g.String()).Any())
							b.If(g.Op("!").Id("ok")).Block(
								g.Return(errorResult, g.Qual("fmt", "Errorf").Call(g.Lit(fmt.Sprintf("expected %q search attribute mapping to return map[string]any, got: %%T", svc.methods[workflow].GoName)), g.Id("result"))),
							)
						}

						if child {
							// local activity wrapper
							fixFn := func(b *g.Group) {
								b.Id("lao").Op(":=").Qual(workflowPkg, "GetLocalActivityOptions").Call(g.Id("ctx"))
								b.Id("lao").Dot("ScheduleToCloseTimeout").Op("=").Qual("time", "Second").Op("*").Lit(10)
								b.If(
									g.Err().Op(":=").Qual(workflowPkg, "ExecuteLocalActivity").Call(
										g.Qual(workflowPkg, "WithLocalActivityOptions").Call(g.Id("ctx"), g.Id("lao")),
										g.Func().Params(g.Id("ctx").Qual("context", "Context")).Params(g.Map(g.String()).Any(), g.Error()).BlockFunc(func(b *g.Group) {
											origFn(b, g.Nil())
											b.Return(g.Id("searchAttributes"), g.Nil())
										}),
									).Dot("Get").Call(g.Id("ctx"), g.Op("&").Id("opts").Dot("SearchAttributes")),
									g.Err().Op("!=").Nil(),
								).Block(
									g.Return(g.Id("opts"), g.Qual("fmt", "Errorf").Call(g.Lit("error evaluating search attributes for %q workflow: %w"), g.Id(svc.toCamel("%sWorkflowName", workflow)), g.Err())),
								)
							}

							// introduce local activity wrapper behind workflow versioning
							switch pvm := svc.patchMode(temporalv1.Patch_PV_64, workflow); pvm {
							case temporalv1.Patch_PVM_ENABLED:
								patchComment(b, temporalv1.Patch_PV_64)
								b.If(g.Add(patchVersion(temporalv1.Patch_PV_64, pvm)).Op("==").Lit(1)).BlockFunc(
									fixFn,
								).Else().BlockFunc(func(b *g.Group) {
									origFn(b, g.Id("opts"))
									b.Id("opts").Dot("SearchAttributes").Op("=").Id("searchAttributes")
								})
							case temporalv1.Patch_PVM_MARKER:
								b.Add(patchVersion(temporalv1.Patch_PV_64, pvm))
								fixFn(b)
							case temporalv1.Patch_PVM_REMOVED:
								fixFn(b)
							default:
								origFn(b, g.Id("opts"))
								b.Id("opts").Dot("SearchAttributes").Op("=").Id("searchAttributes")
							}
						} else {
							origFn(b, g.Id("opts"))
							b.Id("opts").Dot("SearchAttributes").Op("=").Id("searchAttributes")
						}
					})
			}

			// set WorkflowExecutionTimeout
			executionTimeout := fn.If(g.Id("v").Op(":=").Id("o").Dot("executionTimeout"), g.Id("v").Op("!=").Nil()).Block(
				g.Id("opts").Dot("WorkflowExecutionTimeout").Op("=").Op("*").Id("v"),
			)
			if d := opts.GetExecutionTimeout().AsDuration(); d > 0 {
				executionTimeout.Else().If(g.Id("opts").Dot("WorkflowExecutionTimeout").Op("==").Lit(0)).Block(
					g.Id("opts").Dot("WorkflowExecutionTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
				)
			}

			// set WorkflowRunTimeout
			runTimeout := fn.If(g.Id("v").Op(":=").Id("o").Dot("runTimeout"), g.Id("v").Op("!=").Nil()).Block(
				g.Id("opts").Dot("WorkflowRunTimeout").Op("=").Op("*").Id("v"),
			)
			if d := opts.GetRunTimeout().AsDuration(); d > 0 {
				runTimeout.Else().If(g.Id("opts").Dot("WorkflowRunTimeout").Op("==").Lit(0)).Block(
					g.Id("opts").Dot("WorkflowRunTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
				)
			}

			// set WorkflowTaskTimeout
			taskTimeout := fn.If(g.Id("v").Op(":=").Id("o").Dot("taskTimeout"), g.Id("v").Op("!=").Nil()).Block(
				g.Id("opts").Dot("WorkflowTaskTimeout").Op("=").Op("*").Id("v"),
			)
			if d := opts.GetTaskTimeout().AsDuration(); d > 0 {
				taskTimeout.Else().If(g.Id("opts").Dot("WorkflowTaskTimeout").Op("==").Lit(0)).Block(
					g.Id("opts").Dot("WorkflowTaskTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
				)
			}

			// set ParentClosePolicy
			if child {
				parentClosePolicy := fn.If(g.Id("v").Op(":=").Id("o").Dot("parentClosePolicy"), g.Id("v").Op("!=").Qual(enumsPkg, "PARENT_CLOSE_POLICY_UNSPECIFIED")).Block(
					g.Id("opts").Dot("ParentClosePolicy").Op("=").Id("v"),
				)
				if opts.GetParentClosePolicy() != temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_UNSPECIFIED {
					parentClosePolicy.Else().If(g.Id("opts").Dot("ParentClosePolicy").Op("==").Qual(enumsPkg, "PARENT_CLOSE_POLICY_UNSPECIFIED")).BlockFunc(func(bl *g.Group) {
						var policy string
						switch opts.GetParentClosePolicy() {
						case temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_ABANDON:
							policy = "PARENT_CLOSE_POLICY_ABANDON"
						case temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL:
							policy = "PARENT_CLOSE_POLICY_REQUEST_CANCEL"
						case temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE:
							policy = "PARENT_CLOSE_POLICY_TERMINATE"
						}
						bl.Id("opts").Dot("ParentClosePolicy").Op("=").Qual(enumsPkg, policy)
					})
				}
			}

			// set WaitForCancellation
			if child {
				waitForCancellation := fn.If(g.Id("v").Op(":=").Id("o").Dot("waitForCancellation"), g.Id("v").Op("!=").Nil()).Block(
					g.Id("opts").Dot("WaitForCancellation").Op("=").Op("*").Id("v"),
				)
				if opts.GetWaitForCancellation() {
					waitForCancellation.Else().Block(
						g.Id("opts").Dot("WaitForCancellation").Op("=").Lit(true),
					)
				}
			}

			fn.Return(g.Id("opts"), g.Nil())
		})

	baseName := svc.toCamel("With%s", optionsType)
	f.Commentf("%s sets the initial %s.%s", baseName, optionsPkg, optionsType)
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id(baseName).
		Params(g.Id("options").Qual(optionsPkg, optionsType)).
		Op("*").Id(typeName).
		Block(
			g.Id("o").Dot("options").Op("=").Id("options"),
			g.Return(g.Id("o")),
		)

	f.Comment("WithExecutionTimeout sets the WorkflowExecutionTimeout value")
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("WithExecutionTimeout").
		Params(g.Id("d").Qual("time", "Duration")).
		Op("*").Id(typeName).
		Block(
			g.Id("o").Dot("executionTimeout").Op("=").Op("&").Id("d"),
			g.Return(g.Id("o")),
		)

	if child {
		f.Comment("WithID sets the WorkflowID value")
	} else {
		f.Comment("WithID sets the ID value")
	}
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("WithID").
		Params(g.Id("id").String()).
		Op("*").Id(typeName).
		Block(
			g.Id("o").Dot("id").Op("=").Op("&").Id("id"),
			g.Return(g.Id("o")),
		)

	f.Comment("WithIDReusePolicy sets the WorkflowIDReusePolicy value")
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("WithIDReusePolicy").
		Params(g.Id("policy").Qual(enumsPkg, "WorkflowIdReusePolicy")).
		Op("*").Id(typeName).
		Block(
			g.Id("o").Dot("idReusePolicy").Op("=").Id("policy"),
			g.Return(g.Id("o")),
		)

	if child {
		f.Comment("WithParentClosePolicy sets the WorkflowIDReusePolicy value")
		f.Func().
			Params(g.Id("o").Op("*").Id(typeName)).
			Id("WithParentClosePolicy").
			Params(g.Id("policy").Qual(enumsPkg, "ParentClosePolicy")).
			Op("*").Id(typeName).
			Block(
				g.Id("o").Dot("parentClosePolicy").Op("=").Id("policy"),
				g.Return(g.Id("o")),
			)
	}

	f.Comment("WithRetryPolicy sets the RetryPolicy value")
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("WithRetryPolicy").
		Params(g.Id("policy").Op("*").Qual(temporalPkg, "RetryPolicy")).
		Op("*").Id(typeName).
		Block(
			g.Id("o").Dot("retryPolicy").Op("=").Id("policy"),
			g.Return(g.Id("o")),
		)

	f.Comment("WithRunTimeout sets the WorkflowRunTimeout value")
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("WithRunTimeout").
		Params(g.Id("d").Qual("time", "Duration")).
		Op("*").Id(typeName).
		Block(
			g.Id("o").Dot("runTimeout").Op("=").Op("&").Id("d"),
			g.Return(g.Id("o")),
		)

	f.Comment("WithSearchAttributes sets the SearchAttributes value")
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("WithSearchAttributes").
		Params(g.Id("sa").Map(g.String()).Any()).
		Op("*").Id(typeName).
		Block(
			g.Id("o").Dot("searchAttributes").Op("=").Id("sa"),
			g.Return(g.Id("o")),
		)

	f.Comment("WithTaskTimeout sets the WorkflowTaskTimeout value")
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("WithTaskTimeout").
		Params(g.Id("d").Qual("time", "Duration")).
		Op("*").Id(typeName).
		Block(
			g.Id("o").Dot("taskTimeout").Op("=").Op("&").Id("d"),
			g.Return(g.Id("o")),
		)

	f.Comment("WithTaskQueue sets the TaskQueue value")
	f.Func().
		Params(g.Id("o").Op("*").Id(typeName)).
		Id("WithTaskQueue").
		Params(g.Id("tq").String()).
		Op("*").Id(typeName).
		Block(
			g.Id("o").Dot("taskQueue").Op("=").Op("&").Id("tq"),
			g.Return(g.Id("o")),
		)

	if child {
		f.Comment("WithWaitForCancellation sets the WaitForCancellation value")
		f.Func().
			Params(g.Id("o").Op("*").Id(typeName)).
			Id("WithWaitForCancellation").
			Params(g.Id("wait").Bool()).
			Op("*").Id(typeName).
			Block(
				g.Id("o").Dot("waitForCancellation").Op("=").Op("&").Id("wait"),
				g.Return(g.Id("o")),
			)
	}
}

// genClientWorkflowRunImpl generates a <Workflow>Run struct
func (svc *Manifest) genClientWorkflowRunImpl(f *g.File, workflow protoreflect.FullName) {
	clientType := svc.toLowerCamel("%sClient", svc.Service.GoName)
	typeName := svc.toLowerCamel("%sRun", workflow)
	interfaceName := svc.toCamel("%sRun", workflow)

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
func (svc *Manifest) genClientWorkflowRunImplCancelMethod(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toLowerCamel("%sRun", workflow)

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
func (svc *Manifest) genClientWorkflowRunImplGetMethod(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toLowerCamel("%sRun", workflow)
	method := svc.methods[workflow]
	hasOutput := !isEmpty(method.Output)

	f.Comment("Get blocks until the workflow is complete, returning the result if applicable")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Get").
		Params(g.Id("ctx").Qual("context", "Context")).
		ParamsFunc(func(returnVals *g.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *g.Group) {
			if hasOutput {
				fn.Var().Id("resp").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
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
func (svc *Manifest) genClientWorkflowRunImplIDMethod(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toLowerCamel("%sRun", workflow)

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
func (svc *Manifest) genClientWorkflowRunImplQueryMethod(f *g.File, workflow protoreflect.FullName, query protoreflect.FullName) {
	typeName := svc.toLowerCamel("%sRun", workflow)
	handler := svc.methods[query]
	hasInput := !isEmpty(handler.Input)

	commentWithDefaultf(f, methodSet(handler), "%s executes a(n) %s query", query, svc.fqnForQuery(query))
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id(svc.methods[query].GoName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
			}
		}).
		Params(
			g.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), svc.getMessageName(handler.Output)),
			g.Error(),
		).
		Block(
			g.Return(
				g.Id("r").Dot("client").Dot(svc.methods[query].GoName).CallFunc(func(args *g.Group) {
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

// genClientWorkflowRunImplRunMethod generates a <Workflow>Run's Run method
func (svc *Manifest) genClientWorkflowRunImplRunMethod(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toLowerCamel("%sRun", workflow)

	f.Comment("Run returns the inner client.WorkflowRun")
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id("Run").
		Params().
		Qual(clientPkg, "WorkflowRun").
		Block(
			g.Return(g.Id("r").Dot("run")),
		)
}

// genClientWorkflowRunImplRunIDMethod generates a <Workflow>Run's RunID method
func (svc *Manifest) genClientWorkflowRunImplRunIDMethod(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toLowerCamel("%sRun", workflow)

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
func (svc *Manifest) genClientWorkflowRunImplSignalMethod(f *g.File, workflow protoreflect.FullName, signal protoreflect.FullName) {
	typeName := svc.toLowerCamel("%sRun", workflow)
	handler := svc.methods[signal]
	hasInput := !isEmpty(handler.Input)

	// generate get method
	commentWithDefaultf(f, methodSet(handler), "%s sends a(n) %s signal", signal, svc.fqnForSignal(signal))
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id(svc.methods[signal].GoName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
			}
		}).
		Params(g.Error()).
		Block(
			g.ReturnFunc(func(returnVals *g.Group) {
				if svc.methodsFromSameService(signal, workflow) {
					returnVals.Id("r").Dot("client").Dot(svc.methods[signal].GoName).CallFunc(func(args *g.Group) {
						args.Id("ctx")
						args.Id("r").Dot("ID").Call()
						args.Lit("")
						if hasInput {
							args.Id("req")
						}
					})
				} else {
					returnVals.Id("r").Dot("client").Dot("client").Dot("SignalWorkflow").CallFunc(func(args *g.Group) {
						args.Id("ctx")
						args.Id("r").Dot("ID").Call()
						args.Id("r").Dot("RunID").Call()
						args.Add(svc.Qual(signal, svc.toCamel("%sSignalName", signal)))
						if hasInput {
							args.Id("req")
						} else {
							args.Nil()
						}
					})
				}
			}),
		)
}

// genClientWorkflowRunImplTerminateMethod generates a <Workflow>Run's Terminate method
func (svc *Manifest) genClientWorkflowRunImplTerminateMethod(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toLowerCamel("%sRun", workflow)

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
func (svc *Manifest) genClientWorkflowRunImplUpdateAsyncMethod(f *g.File, workflow protoreflect.FullName, update protoreflect.FullName) {
	typeName := svc.toLowerCamel("%sRun", workflow)
	methodName := svc.toCamel("%sAsync", update)
	handler := svc.methods[update]
	hasInput := !isEmpty(handler.Input)

	commentWithDefaultf(f, methodSet(handler), "%s start a(n) %s workflow update and returns a handle to the update", methodName, svc.fqnForUpdate(update))
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id(methodName).
		ParamsFunc(func(args *g.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(svc.toCamel("%sOptions", update))
		}).
		Params(
			g.Add(svc.Qual(update, svc.toCamel("%sHandle", update))),
			g.Error(),
		).
		BlockFunc(func(fn *g.Group) {
			if svc.methodsFromSamePackage(workflow, update) {
				fn.Return(
					g.Id("r").Dot("client").Dot(svc.toCamel("%sAsync", update)).CallFunc(func(args *g.Group) {
						args.Id("ctx")
						args.Id("r").Dot("ID").Call()
						args.Id("r").Dot("RunID").Call()
						if hasInput {
							args.Id("req")
						}
						args.Id("opts").Op("...")
					}),
				)
			} else {
				// initialize options
				fn.Var().Id("o").Op("*").Id(svc.toCamel("%sOptions", update))
				fn.If(g.Len(g.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(g.Lit(0)).Op("!=").Nil()).Block(
					g.Id("o").Op("=").Id("opts").Index(g.Lit(0)),
				).Else().Block(
					g.Id("o").Op("=").Id(svc.toCamel("New%sOptions", update)).Call(),
				)

				// build UpdateWorkflowWithOptions
				fn.List(g.Id("options"), g.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(args *g.Group) {
					args.Id("r").Dot("ID").Call()
					args.Id("r").Dot("RunID").Call()
					if hasInput {
						args.Id("req")
					}
				})
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Nil(), g.Qual("fmt", "Errorf").Call(g.Lit("error initializing UpdateWorkflowWithOptions: %w"), g.Err())),
				)

				// update workflow
				fn.List(g.Id("handle"), g.Err()).Op(":=").Id("c").Dot("client").Dot("UpdateWorkflowWithOptions").Call(g.Id("ctx"), g.Id("options"))
				fn.If(g.Err().Op("!=").Nil()).Block(
					g.Return(g.Nil(), g.Err()),
				)

				fn.Return(
					g.Op("&").Add(svc.Qual(update, svc.toCamel("%sHandle", update))).Values(
						g.Id("client").Op(":").Id("c"),
						g.Id("handle").Op(":").Id("handle"),
					),
					g.Nil(),
				)
			}
		})
}

// genClientWorkflowRunImplUpdateMethod generates a <Workflow>Run's <Update> method
func (svc *Manifest) genClientWorkflowRunImplUpdateMethod(f *g.File, workflow protoreflect.FullName, update protoreflect.FullName) {
	typeName := svc.toLowerCamel("%sRun", workflow)
	handler := svc.methods[update]
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)

	commentWithDefaultf(f, methodSet(handler), "%s executes a(n) %s workflow update", update, svc.fqnForUpdate(update))
	f.Func().
		Params(g.Id("r").Op("*").Id(typeName)).
		Id(svc.methods[update].GoName).
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

// genClientWorkflowRunInterface generates a <Workflow>Run interface
func (svc *Manifest) genClientWorkflowRunInterface(f *g.File, workflow protoreflect.FullName) {
	typeName := svc.toCamel("%sRun", workflow)
	opts := svc.workflows[workflow]
	method := svc.methods[workflow]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s describes a(n) %s workflow run", typeName, svc.fqnForWorkflow(workflow))
	f.Type().Id(typeName).InterfaceFunc(func(methods *g.Group) {
		methods.Comment("ID returns the workflow ID")
		methods.Id("ID").Params().String().Line()

		methods.Comment("RunID returns the workflow instance ID")
		methods.Id("RunID").Params().String().Line()

		methods.Comment("Run returns the inner client.WorkflowRun")
		methods.Id("Run").Params().Qual(clientPkg, "WorkflowRun").Line()

		methods.Comment("Get blocks until the workflow is complete and returns the result")
		methods.Id("Get").
			Params(g.Id("ctx").Qual("context", "Context")).
			ParamsFunc(func(returnVals *g.Group) {
				if hasOutput {
					returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
				}
				returnVals.Error()
			}).Line()

		methods.Comment("Cancel requests cancellation of a workflow in execution, returning an error if applicable")
		methods.Id("Cancel").
			Params(g.Id("ctx").Qual("context", "Context")).
			Params(g.Error()).Line()

		methods.Comment("Terminate terminates a workflow in execution, returning an error if applicable")
		methods.Id("Terminate").
			Params(
				g.Id("ctx").Qual("context", "Context"),
				g.Id("reason").String(),
				g.Id("details").Op("...").Interface(),
			).
			Params(g.Error()).Line()

		for _, queryOpts := range opts.GetQuery() {
			query := getFullyQualifiedRef(workflow, queryOpts.GetRef())
			handler := svc.methods[query]
			hasInput := !isEmpty(handler.Input)

			commentWithDefaultf(methods, methodSet(handler), "%s executes a(n) %s query", query, svc.fqnForQuery(query))
			methods.Id(svc.toCamel("%s", query)).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
					}
				}).
				Params(
					g.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), svc.getMessageName(handler.Output)),
					g.Error(),
				).Line()
		}

		for _, signalOpts := range opts.GetSignal() {
			signal := getFullyQualifiedRef(workflow, signalOpts.GetRef())
			handler := svc.methods[signal]
			hasInput := !isEmpty(handler.Input)

			commentWithDefaultf(methods, methodSet(handler), "%s sends a(n) %s signal", signal, svc.fqnForSignal(signal))
			methods.Id(svc.toCamel("%s", signal)).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
					}
				}).
				Params(g.Error()).Line()
		}

		for _, updateOpts := range opts.GetUpdate() {
			update := getFullyQualifiedRef(workflow, updateOpts.GetRef())
			handler := svc.methods[update]
			hasInput := !isEmpty(handler.Input)
			hasOutput := !isEmpty(handler.Output)

			// add synchronous flavor
			commentWithDefaultf(methods, methodSet(handler), "%s executes a(n) %s update", update, svc.fqnForUpdate(update))
			methods.Id(svc.toCamel("%s", update)).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
					}
					args.Id("opts").Op("...").Op("*").Add(svc.Qual(update, svc.toCamel("%sOptions", update)))
				}).
				ParamsFunc(func(returnVals *g.Group) {
					if hasOutput {
						returnVals.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), svc.getMessageName(handler.Output))
					}
					returnVals.Error()
				}).Line()

			// add async flavor
			commentWithDefaultf(methods, methodSet(handler), "%sAsync sends a(n) %s update to the workflow", update, svc.fqnForUpdate(update))
			methods.Id(svc.toCamel("%sAsync", update)).
				ParamsFunc(func(args *g.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), svc.getMessageName(handler.Input))
					}
					args.Id("opts").Op("...").Op("*").Add(svc.Qual(update, svc.toCamel("%sOptions", update)))
				}).
				Params(
					g.Add(svc.Qual(update, svc.toCamel("%sHandle", update))),
					g.Error(),
				).Line()
		}
	})
}
