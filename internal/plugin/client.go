package plugin

import (
	"cmp"
	"fmt"
	"strconv"

	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	j "github.com/dave/jennifer/jen"
	"github.com/hako/durafmt"
	"go.temporal.io/api/enums/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (n *names) clientCtor() string {
	return n.toCamel("New%sClient", n.Service.GoName)
}

func (n *names) clientImpl() string {
	return n.toLowerCamel("%sClient", n.Service.GoName)
}

func (n *names) clientUpdate(update protoreflect.FullName) string {
	return n.toCamel("%s", update)
}

func (n *names) clientUpdateAsync(update protoreflect.FullName) string {
	return n.toCamel("%sAsync", update)
}

func (n *names) clientUpdateHandleIface(update protoreflect.FullName) string {
	return n.toCamel("%sHandle", update)
}

func (n *names) clientUpdateHandleImpl(update protoreflect.FullName) string {
	return n.toLowerCamel("%sHandle", update)
}

func (n *names) clientUpdateOptions(update protoreflect.FullName) string {
	return n.toCamel("%sOptions", update)
}

func (n *names) clientUpdateOptionsCtor(update protoreflect.FullName) string {
	return n.toCamel("New%sOptions", update)
}

func (n *names) clientUpdateWithStart(workflow, update protoreflect.FullName) string {
	return n.toCamel("%sWith%s", workflow, update)
}

func (n *names) clientUpdateWithStartAsync(workflow, update protoreflect.FullName) string {
	return n.toCamel("%sWith%sAsync", workflow, update)
}

func (n *names) clientUpdateWithStartOptions(workflow, update protoreflect.FullName) string {
	return n.toCamel("%sWith%sOptions", workflow, update)
}

func (n *names) clientUpdateWithStartOptionsCtor(workflow, update protoreflect.FullName) string {
	return n.toCamel("New%sWith%sOptions", workflow, update)
}

func (n *names) clientWorkflow(workflow protoreflect.FullName) string {
	return n.toCamel("%s", workflow)
}

func (n *names) clientWorkflowAsync(workflow protoreflect.FullName) string {
	return n.toCamel("%sAsync", workflow)
}

func (n *names) clientWorkflowRun(workflow protoreflect.FullName) string {
	return n.toCamel("%sRun", workflow)
}

func (n *names) clientWorkflowGet(workflow protoreflect.FullName) string {
	return n.toCamel("Get%s", workflow)
}

func (n *names) clientWorkflowOptions(workflow protoreflect.FullName) string {
	return n.toCamel("%sOptions", workflow)
}

func (n *names) clientWorkflowOptionsCtor(workflow protoreflect.FullName) string {
	return n.toCamel("New%sOptions", workflow)
}

// genClientImpl generates a <service>Client implementation
func (m *Manifest) genClientImpl(f *j.File) {
	typeName := m.toLowerCamel("%sClient", m.Service.GoName)

	f.Commentf("%s implements a temporal client for a %s service", typeName, m.Service.Desc.FullName())
	f.Type().
		Id(typeName).
		StructFunc(func(fields *j.Group) {
			fields.Id("client").Qual(clientPkg, "Client")
			fields.Id("log").Op("*").Qual("log/slog", "Logger")
		})
}

// genClientImplConstructor generates a New<Service>Client function
func (m *Manifest) genClientImplConstructor(f *j.File) {
	methodName := m.toCamel("New%sClient", m.Service.GoName)
	implName := m.toLowerCamel("%sClient", m.Service.GoName)
	interfaceName := m.toCamel("%sClient", m.Service.GoName)
	optionsName := m.toLowerCamel("%sClientOptions", m.Service.GoName)

	f.Commentf("%s initializes a new %s client", methodName, m.Service.Desc.FullName())
	f.Func().
		Id(methodName).
		Params(
			j.Id("c").Qual(clientPkg, "Client"),
			j.Id("options").Op("...").Op("*").Id(optionsName),
		).
		Params(
			j.Id(interfaceName),
		).
		Block(
			j.Var().Id("cfg").Op("*").Id(optionsName),
			j.If(j.Len(j.Id("options")).Op(">").Lit(0)).
				Block(
					j.Id("cfg").Op("=").Id("options").Index(j.Lit(0)),
				).
				Else().
				Block(
					j.Id("cfg").Op("=").Id(m.toCamel("New%sClientOptions", m.Service.GoName)).Call(),
				),
			j.Return(
				j.Op("&").Id(implName).Custom(
					multiLineValues,
					j.Id("client").Op(":").Id("c"),
					j.Id("log").Op(":").Id("cfg").Dot("getLogger").Call(),
				),
			),
		)

	methodName += "WithOptions"
	f.Commentf("%s initializes a new %s client with the given options", methodName, m.Service.GoName)
	f.Func().
		Id(methodName).
		Params(
			j.Id("c").Qual(clientPkg, "Client"),
			j.Id("opts").Qual(clientPkg, "Options"),
			j.Id("options").Op("...").Op("*").Id(optionsName),
		).
		Params(
			j.Id(interfaceName),
			j.Error(),
		).
		Block(
			j.Var().Err().Error(),
			j.List(j.Id("c"), j.Err()).Op("=").Qual(clientPkg, "NewClientFromExisting").Call(j.Id("c"), j.Id("opts")),
			j.If().Err().Op("!=").Nil().Block(
				j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing client with options: %w"), j.Err())),
			),
			j.Var().Id("cfg").Op("*").Id(optionsName),
			j.If(j.Len(j.Id("options")).Op(">").Lit(0)).
				Block(
					j.Id("cfg").Op("=").Id("options").Index(j.Lit(0)),
				).
				Else().
				Block(
					j.Id("cfg").Op("=").Id(m.toCamel("New%sClientOptions", m.Service.GoName)).Call(),
				),
			j.Return(
				j.Op("&").Id(implName).Custom(
					multiLineValues,
					j.Id("client").Op(":").Id("c"),
					j.Id("log").Op(":").Id("cfg").Dot("getLogger").Call(),
				),
				j.Nil(),
			),
		)
}

// genClientImplQueryMethod adds a <Query> method to a workflowClient
func (m *Manifest) genClientImplQueryMethod(f *j.File, query protoreflect.FullName) {
	clientType := m.toLowerCamel("%sClient", m.Service.GoName)
	method := m.methods[query]
	hasInput := !isEmpty(method.Input)

	commentWithDefaultf(f, methodSet(method), "%s sends a(n) %s query to an existing workflow", query, m.fqnForQuery(query))
	f.Func().
		Params(j.Id("c").Op("*").Id(clientType)).
		Id(m.methods[query].GoName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("query").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
		}).
		Params(
			j.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output)),
			j.Error(),
		).
		BlockFunc(func(fn *j.Group) {
			if isDeprecated(method) {
				fn.Id("c").Dot("log").Dot("WarnContext").Call(j.Id("ctx"), j.Lit("use of deprecated client method detected"), j.Lit("method"), j.Lit(m.toCamel("%s", query)), j.Lit("query"), j.Id(m.toCamel("%sQueryName", query))).Line()
			}
			fn.Var().Id("resp").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			fn.If(
				j.List(j.Id("val"), j.Err()).Op(":=").Id("c").Dot("client").Dot("QueryWorkflow").CallFunc(func(args *j.Group) {
					args.Id("ctx")
					args.Id("workflowID")
					args.Id("runID")
					args.Id(m.toCamel("%sQueryName", query))
					if hasInput {
						args.Id("query")
					}
				}),
				j.Err().Op("!=").Nil(),
			).Block(
				j.Return(j.Nil(), j.Err()),
			).Else().If(
				j.Err().Op("=").Id("val").Dot("Get").Call(
					j.Op("&").Id("resp"),
				),
				j.Err().Op("!=").Nil(),
			).Block(
				j.Return(j.Nil(), j.Err()),
			)
			fn.Return(
				j.Op("&").Id("resp"), j.Nil(),
			)
		})
}

// genClientImplSignalMethod adds a <Signal> method to a workflowClient
func (m *Manifest) genClientImplSignalMethod(f *j.File, signal protoreflect.FullName) {
	clientType := m.toLowerCamel("%sClient", m.Service.GoName)
	method := m.methods[signal]
	hasInput := !isEmpty(method.Input)

	commentWithDefaultf(f, methodSet(method), "%s sends a(n) %s signal to an existing workflow", signal, m.fqnForSignal(signal))
	f.Func().
		Params(j.Id("c").Op("*").Id(clientType)).
		Id(m.methods[signal].GoName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual("context", "Context")
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("signal").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
		}).
		Params(j.Error()).
		BlockFunc(func(fn *j.Group) {
			if isDeprecated(method) {
				fn.Id("c").Dot("log").Dot("WarnContext").Call(j.Id("ctx"), j.Lit("use of deprecated client method detected"), j.Lit("method"), j.Lit(m.toCamel("%s", signal)), j.Lit("signal"), j.Id(m.toCamel("%sSignalName", signal)))
				fn.Line()
			}
			fn.Return(
				j.Id("c").Dot("client").Dot("SignalWorkflow").CallFunc(func(args *j.Group) {
					args.Id("ctx")
					args.Id("workflowID")
					args.Id("runID")
					args.Id(m.toCamel("%sSignalName", signal))
					if hasInput {
						args.Id("signal")
					} else {
						args.Nil()
					}
				}),
			)
		})
}

// genClientImplSignalWithStartMethod adds a Start<Workflow>With<Signal> client method
func (m *Manifest) genClientImplSignalWithStartMethod(f *j.File, workflow, signal protoreflect.FullName) {
	clientType := m.toLowerCamel("%sClient", m.Service.GoName)
	method := m.methods[workflow]
	handler := m.methods[signal]
	name := m.toCamel("%sWith%s", workflow, signal)
	hasWorkflowInput := !isEmpty(method.Input)
	hasWorkflowOutput := !isEmpty(method.Output)
	hasSignalInput := !isEmpty(handler.Input)

	commentWithDefaultf(f, methodSet(method, handler), "%s starts a(n) %s workflow and sends a(n) %s signal in a transaction", name, m.fqnForWorkflow(workflow), m.fqnForSignal(signal))
	f.Func().
		Params(j.Id("c").Op("*").Id(clientType)).
		Id(name).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasWorkflowInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if hasSignalInput {
				args.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			args.Id("options").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
		}).
		ParamsFunc(func(returnVals *j.Group) {
			if hasWorkflowOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *j.Group) {
			if workflowDeprecated, signalDeprecated := isDeprecated(method), isDeprecated(handler); workflowDeprecated || signalDeprecated {
				if workflowDeprecated {
					fn.Id("c").Dot("log").Dot("WarnContext").Call(j.Id("ctx"), j.Lit("use of deprecated client method detected"), j.Lit("method"), j.Lit(name), j.Lit("workflow"), j.Id(m.toCamel("%sWorkflowName", workflow)))
				}
				if signalDeprecated {
					fn.Id("c").Dot("log").Dot("WarnContext").Call(j.Id("ctx"), j.Lit("use of deprecated client method detected"), j.Lit("method"), j.Lit(name), j.Lit("signal"), j.Id(m.toCamel("%sSignalName", signal)))
				}
				fn.Line()
			}
			// signal with start workflow
			fn.Id("run").Op(",").Err().Op(":=").Id("c").Dot(m.toCamel("%sWith%sAsync", workflow, signal)).CallFunc(func(args *j.Group) {
				args.Id("ctx")
				if hasWorkflowInput {
					args.Id("req")
				}
				if hasSignalInput {
					args.Id("signal")
				}
				args.Id("options").Op("...")
			})
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.ReturnFunc(func(returnVals *j.Group) {
					if hasWorkflowOutput {
						returnVals.Nil()
					}
					returnVals.Err()
				}),
			)
			fn.Return(
				j.Id("run").Dot("Get").Call(j.Id("ctx")),
			)
		})
}

// genClientImplSignalWithStartMethodAsync adds a <Workflow>With<Signal>Async client method
func (m *Manifest) genClientImplSignalWithStartMethodAsync(f *j.File, workflow, signal protoreflect.FullName) {
	clientType := m.toLowerCamel("%sClient", m.Service.GoName)
	method := m.methods[workflow]
	handler := m.methods[signal]
	name := m.toCamel("%sWith%sAsync", workflow, signal)
	runName := m.toLowerCamel("%sRun", workflow)
	hasWorkflowInput := !isEmpty(method.Input)
	hasSignalInput := !isEmpty(handler.Input)

	commentWithDefaultf(f, methodSet(method, handler), "%s starts a(n) %s workflow and sends a(n) %s signal in a transaction", name, m.fqnForWorkflow(workflow), m.fqnForSignal(signal))
	f.Func().
		Params(j.Id("c").Op("*").Id(clientType)).
		Id(name).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasWorkflowInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if hasSignalInput {
				args.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			args.Id("options").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
		}).
		Params(
			j.Id(m.toCamel("%sRun", workflow)),
			j.Error(),
		).
		BlockFunc(func(fn *j.Group) {
			if workflowDeprecated, signalDeprecated := isDeprecated(method), isDeprecated(handler); workflowDeprecated || signalDeprecated {
				if workflowDeprecated {
					fn.Id("c").Dot("log").Dot("WarnContext").Call(j.Id("ctx"), j.Lit("use of deprecated client method detected"), j.Lit("method"), j.Lit(name), j.Lit("workflow"), j.Id(m.toCamel("%sWorkflowName", workflow)))
				}
				if signalDeprecated {
					fn.Id("c").Dot("log").Dot("WarnContext").Call(j.Id("ctx"), j.Lit("use of deprecated client method detected"), j.Lit("method"), j.Lit(name), j.Lit("signal"), j.Id(m.toCamel("%sSignalName", signal)))
				}
				fn.Line()
			}

			// initialize options
			fn.Var().Id("o").Op("*").Id(m.toCamel("%sOptions", workflow))
			fn.If(j.Len(j.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(j.Lit(0)).Op("!=").Nil()).Block(
				j.Id("o").Op("=").Id("options").Index(j.Lit(0)),
			).Else().Block(
				j.Id("o").Op("=").Id(m.toCamel("New%sOptions", workflow)).Call(),
			)

			// initialize client.StartWorkfowOptions
			fn.List(j.Id("opts"), j.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(args *j.Group) {
				if hasWorkflowInput {
					args.Id("req").Dot("ProtoReflect").Call()
				} else {
					args.Nil()
				}
			})
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing client.StartWorkflowOptions: %w"), j.Err())),
			)

			// signal with start workflow
			fn.Id("run").Op(",").Err().Op(":=").Id("c").Dot("client").Dot("SignalWithStartWorkflow").CallFunc(func(args *j.Group) {
				args.Id("ctx")
				args.Id("opts").Dot("ID")
				args.Qual(m.goImportPathForMethod(signal), m.toCamel("%sSignalName", signal))
				if hasSignalInput {
					args.Id("signal")
				} else {
					args.Nil()
				}
				args.Id("opts")
				args.Id(m.toCamel("%sWorkflowName", workflow))
				if hasWorkflowInput {
					args.Id("req")
				}
			})
			fn.If(j.Id("run").Op("==").Nil().Op("||").Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Err()),
			)
			fn.Return(
				j.Op("&").Id(runName).Block(
					j.Id("client").Op(":").Id("c").Op(","),
					j.Id("run").Op(":").Id("run").Op(","),
				),
				j.Nil(),
			)
		})
}

func (m *Manifest) genClientImplUpdateGetMethod(f *j.File, update protoreflect.FullName) {
	methodName := m.toCamel("Get%s", m.methods[update].GoName)
	clientType := m.toLowerCamel("%sClient", m.Service.GoName)

	f.Commentf("%s retrieves a handle to an existing %s update", methodName, m.fqnForUpdate(update))
	f.Func().
		Params(j.Id("c").Op("*").Id(clientType)).
		Id(methodName).
		Params(
			j.Id("ctx").Qual("context", "Context"),
			j.Id("req").Qual(clientPkg, "GetWorkflowUpdateHandleOptions"),
		).
		Params(
			j.Id(m.toCamel("%sHandle", update)),
			j.Error(),
		).
		Block(
			j.Return(
				j.Op("&").Id(m.toLowerCamel("%sHandle", update)).Custom(
					multiLineValues,
					j.Id("client").Op(":").Id("c"),
					j.Id("handle").Op(":").Id("c").Dot("client").Dot("GetWorkflowUpdateHandle").Call(
						j.Id("req"),
					),
				),
				j.Nil(),
			),
		)
}

// genClientImplUpdateMethod adds an <Update> method to a workflowClient
func (m *Manifest) genClientImplUpdateMethod(f *j.File, update protoreflect.FullName) {
	clientType := m.toLowerCamel("%sClient", m.Service.GoName)
	handler := m.methods[update]
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)

	methodName := m.Names().clientUpdate(update)
	asyncName := m.Names().clientUpdateAsync(update)

	commentWithDefaultf(f, methodSet(handler), "%s sends a(n) %s update to an existing workflow", update, m.fqnForUpdate(update))
	f.Func().
		Params(j.Id("c").Op("*").Id(clientType)).
		Id(methodName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			g.Id("workflowID").String()
			g.Id("runID").String()
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
		BlockFunc(func(g *j.Group) {
			if isDeprecated(handler) {
				g.Id("c").Dot("log").Dot("WarnContext").Call(j.Id("ctx"), j.Lit("use of deprecated client method detected"), j.Lit("method"), j.Lit(m.toCamel("%s", update)), j.Lit("update"), j.Id(m.toCamel("%sUpdateName", update))).Line()
			}

			// initialize update options
			g.Comment("initialize update options")
			g.Id("o").Op(":=").Id(m.toCamel("New%sOptions", update)).Call()
			g.If(j.Len(j.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(j.Lit(0)).Dot("Options").Op("!=").Nil()).Block(
				j.Id("o").Op("=").Id("opts").Index(j.Lit(0)),
			)

			g.Line()
			g.Comment("call sync update with WorkflowUpdateStageCompleted wait policy")
			g.List(j.Id("handle"), j.Err()).Op(":=").Id("c").Dot(asyncName).CallFunc(func(g *j.Group) {
				g.Id("ctx")
				g.Id("workflowID")
				g.Id("runID")
				if hasInput {
					g.Id("req")
				}
				g.Id("o").Dot("WithWaitPolicy").Call(j.Qual(clientPkg, "WorkflowUpdateStageCompleted"))
			})
			g.If(j.Err().Op("!=").Nil()).Block(
				j.ReturnFunc(func(g *j.Group) {
					if hasOutput {
						g.Nil()
					}
					g.Err()
				}),
			)

			g.Line()
			g.Comment("block on update completion")
			g.Return(j.Id("handle").Dot("Get").Call(j.Id("ctx")))
		})
}

// genClientImplUpdateMethodAsync adds an <Update>Async method to a workflowClient
func (m *Manifest) genClientImplUpdateMethodAsync(f *j.File, update protoreflect.FullName) {
	clientType := m.toLowerCamel("%sClient", m.Service.GoName)
	handler := m.methods[update]
	handleName := m.toLowerCamel("%sHandle", update)
	hasInput := !isEmpty(handler.Input)
	methodName := m.Names().clientUpdateAsync(update)

	commentWithDefaultf(f, methodSet(handler), "%s sends a(n) %s update to an existing workflow", update, m.fqnForUpdate(update))
	f.Func().
		Params(j.Id("c").Op("*").Id(clientType)).
		Id(methodName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			g.Id("workflowID").String()
			g.Id("runID").String()
			if hasInput {
				g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			g.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", update))
		}).
		Params(
			j.Id(m.toCamel("%sHandle", update)),
			j.Error(),
		).
		BlockFunc(func(g *j.Group) {
			if isDeprecated(handler) {
				g.Id("c").Dot("log").Dot("WarnContext").Call(j.Id("ctx"), j.Lit("use of deprecated client method detected"), j.Lit("method"), j.Lit(methodName), j.Lit("update"), j.Id(m.toCamel("%sUpdateName", update))).Line()
			}

			// initialize options
			g.Comment("initialize update options")
			g.Var().Id("o").Op("*").Id(m.toCamel("%sOptions", update))
			g.If(j.Len(j.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(j.Lit(0)).Op("!=").Nil()).Block(
				j.Id("o").Op("=").Id("opts").Index(j.Lit(0)),
			).Else().Block(
				j.Id("o").Op("=").Id(m.toCamel("New%sOptions", update)).Call(),
			)

			g.Line()
			g.Comment("build UpdateWorkflowOptions")
			g.List(j.Id("options"), j.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(g *j.Group) {
				g.Id("workflowID")
				g.Id("runID")
				if hasInput {
					g.Id("req")
				}
			})
			g.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing UpdateWorkflowWithOptions: %w"), j.Err())),
			)

			g.Line()
			g.Comment("update workflow")
			g.List(j.Id("handle"), j.Err()).Op(":=").Id("c").Dot("client").Dot("UpdateWorkflow").Call(j.Id("ctx"), j.Op("*").Id("options"))
			g.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Err()),
			)

			g.Return(
				j.Op("&").Id(handleName).Values(
					j.Id("client").Op(":").Id("c"),
					j.Id("handle").Op(":").Id("handle"),
				),
				j.Nil(),
			)
		})
}

func (m *Manifest) genClientImplUpdateWithStartMethod(f *j.File, workflow, update protoreflect.FullName) {
	method := m.methods[workflow]
	hasWorkflowInput := !isEmpty(method.Input)

	handler := m.methods[update]
	hasUpdateInput := !isEmpty(handler.Input)
	hasUpdateOutput := !isEmpty(handler.Output)

	methodName := m.Names().clientUpdateWithStart(workflow, update)
	asyncName := m.Names().clientUpdateWithStartAsync(workflow, update)
	clientImplName := m.Names().clientImpl()
	optionsName := m.Names().clientUpdateWithStartOptions(workflow, update)
	runName := m.Names().clientWorkflowRun(workflow)

	commentf(f, methodSet(method, handler), "%s starts a(n) %s workflow and executes a(n) %s update in a transaction", methodName, m.fqnForWorkflow(workflow), m.fqnForUpdate(update))
	f.Func().
		Params(j.Id("c").Op("*").Id(clientImplName)).
		Id(methodName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasWorkflowInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if hasUpdateInput {
				g.Id("update").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			g.Id("options").Op("...").Op("*").Id(optionsName)
		}).
		ParamsFunc(func(g *j.Group) {
			if hasUpdateOutput {
				g.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
			}
			g.Id(runName)
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			g.List(j.Id("updateHandle"), j.Id("run"), j.Err()).Op(":=").Id("c").Dot(asyncName).CallFunc(func(g *j.Group) {
				g.Id("ctx")
				if hasWorkflowInput {
					g.Id("req")
				}
				if hasUpdateInput {
					g.Id("update")
				}
				g.Id("options").Op("...")
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

			if hasUpdateOutput {
				g.List(j.Id("out"), j.Err()).Op(":=").Id("updateHandle").Dot("Get").Call(j.Id("ctx"))
				g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
					g.Return(j.Nil(), j.Id("run"), j.Err())
				})
			} else {
				g.If(j.Err().Op(":=").Id("updateHandle").Dot("Get").Call(j.Id("ctx")), j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
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

func (m *Manifest) genClientImplUpdateWithStartMethodAsync(f *j.File, workflow, update protoreflect.FullName) {
	method := m.methods[workflow]
	hasWorkflowInput := !isEmpty(method.Input)

	handler := m.methods[update]
	hasUpdateInput := !isEmpty(handler.Input)

	methodName := m.Names().clientUpdateWithStartAsync(workflow, update)
	workflowName := m.Names().workflowName(workflow)
	clientImplName := m.Names().clientImpl()
	optionsName := m.Names().clientUpdateWithStartOptions(workflow, update)
	optionsCtorName := m.Names().clientUpdateWithStartOptionsCtor(workflow, update)
	runName := m.Names().clientWorkflowRun(workflow)
	handleIfaceName := m.Names().clientUpdateHandleIface(update)
	handleImplName := m.Names().clientUpdateHandleImpl(update)
	getWorkflow := m.Names().clientWorkflowGet(workflow)

	commentf(f, methodSet(method, handler), "%s starts a(n) %s workflow and executes a(n) %s update in a transaction", methodName, m.fqnForWorkflow(workflow), m.fqnForUpdate(update))
	f.Func().
		Params(j.Id("c").Op("*").Id(clientImplName)).
		Id(methodName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasWorkflowInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if hasUpdateInput {
				g.Id("update").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			g.Id("options").Op("...").Op("*").Id(optionsName)
		}).
		Params(j.Id(handleIfaceName), j.Id(runName), j.Error()).
		BlockFunc(func(g *j.Group) {
			// initialize method options
			g.Var().Id("o").Op("*").Id(optionsName)
			g.If(j.Len(j.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(j.Lit(0)).Op("!=").Nil()).BlockFunc(func(g *j.Group) {
				g.Id("o").Op("=").Id("options").Index(j.Lit(0))
			}).Else().BlockFunc(func(g *j.Group) {
				g.Id("o").Op("=").Id(optionsCtorName).Call()
			})

			// initialize client.UpdateWorkflowWithOptions
			g.List(j.Id("opts"), j.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(g *j.Group) {
				g.Id("ctx")
				g.Func().Params(j.Id("swo").Qual(clientPkg, "StartWorkflowOptions")).Qual(clientPkg, "WithStartWorkflowOperation").BlockFunc(func(g *j.Group) {
					g.ReturnFunc(func(g *j.Group) {
						g.Id("c").Dot("client").Dot("NewWithStartWorkflowOperation").CallFunc(func(g *j.Group) {
							g.Id("swo")
							g.Id(workflowName)
							if hasWorkflowInput {
								g.Id("req")
							}
						})
					})
				})
				if hasWorkflowInput {
					g.Id("req")
				}
				if hasUpdateInput {
					g.Id("update")
				}
			})
			g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
				g.Return(j.Nil(), j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing UpdateWorkflowWithOptions: %w"), j.Err()))
			})

			g.List(j.Id("handle"), j.Err()).Op(":=").Id("c").Dot("client").Dot("UpdateWithStartWorkflow").Call(j.Id("ctx"), j.Id("opts"))
			g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
				g.Return(j.Nil(), j.Nil(), j.Err())
			})
			g.ReturnFunc(func(g *j.Group) {
				g.Op("&").Id(handleImplName).Custom(multiLineValues,
					j.Id("client").Op(":").Id("c"),
					j.Id("handle").Op(":").Id("handle"),
				)
				g.Id("c").Dot(getWorkflow).Call(j.Id("ctx"), j.Id("handle").Dot("WorkflowID").Call(), j.Id("handle").Dot("RunID").Call())
				g.Nil()
			})
		})
}

func (m *Manifest) genClientUpdateWithStartOptions(f *j.File, workflow, update protoreflect.FullName) {
	method := m.methods[workflow]
	opts := m.workflows[workflow]
	hasWorkflowInput := !isEmpty(method.Input)

	handler := m.methods[update]
	hasUpdateInput := !isEmpty(handler.Input)

	workflowOptions := m.Names().clientWorkflowOptions(workflow)
	workflowOptionsCtor := m.Names().clientWorkflowOptionsCtor(workflow)
	updateOptions := m.Names().clientUpdateOptions(update)
	updateOptionsCtor := m.Names().clientUpdateOptionsCtor(update)

	optionsName := m.Names().clientUpdateWithStartOptions(workflow, update)
	commentf(f, methodSet(m.methods[workflow], m.methods[update]), "%s is the options for a %s workflow with a %s update", optionsName, m.fqnForWorkflow(workflow), m.fqnForUpdate(update))
	f.Type().
		Id(optionsName).
		StructFunc(func(g *j.Group) {
			g.Id("options").Qual(clientPkg, "UpdateWithStartWorkflowOptions")
			g.Id("workflowOptions").Op("*").Id(workflowOptions)
			g.Id("updateOptions").Op("*").Id(updateOptions)
		})

	ctorName := m.Names().clientUpdateWithStartOptionsCtor(workflow, update)
	f.Commentf("%s initializes a new %s value", ctorName, optionsName)
	f.Func().
		Id(ctorName).
		Params().
		Op("*").Id(optionsName).
		BlockFunc(func(g *j.Group) {
			g.ReturnFunc(func(g *j.Group) {
				g.Op("&").Id(optionsName).Values()
			})
		})

	f.Commentf("Build transforms %s into valid client.UpdateWithStartWorkflowOptions", optionsName)
	f.Func().
		Params(j.Id("o").Op("*").Id(optionsName)).
		Id("Build").
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			g.Id("op").Func().Params(j.Qual(clientPkg, "StartWorkflowOptions")).Qual(clientPkg, "WithStartWorkflowOperation")
			if hasWorkflowInput {
				g.Id("input").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if hasUpdateInput {
				g.Id("update").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
		}).
		ParamsFunc(func(g *j.Group) {
			g.Id("options").Qual(clientPkg, "UpdateWithStartWorkflowOptions")
			g.Err().Error()
		}).
		BlockFunc(func(g *j.Group) {
			// initialize result to empty or user-provided UpdateWithStartWorkflowOptions
			g.Id("options").Op("=").Id("o").Dot("options")

			// initialize workflow options
			g.If(j.Id("o").Dot("workflowOptions").Op("==").Nil()).Block(
				j.Id("o").Dot("workflowOptions").Op("=").Id(workflowOptionsCtor).Call(),
			)

			// initialize start workflow options
			g.List(j.Id("swo"), j.Err()).Op(":=").Id("o").Dot("workflowOptions").Dot("Build").CallFunc(func(g *j.Group) {
				if hasWorkflowInput {
					g.Id("input").Dot("ProtoReflect").Call()
				} else {
					g.Nil()
				}
			})
			g.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Id("options"), j.Err()),
			)

			// set default WorkflowIDConflictPolicy if not specified
			defaultWorkflowIDConflictPolicy := enums.WORKFLOW_ID_CONFLICT_POLICY_FAIL
			for _, u := range opts.GetUpdate() {
				if getFullyQualifiedRef(workflow, u.GetRef()) == update {
					defaultWorkflowIDConflictPolicy = cmp.Or(u.GetWorkflowIdConflictPolicy(), defaultWorkflowIDConflictPolicy)
				}
			}
			if defaultWorkflowIDConflictPolicy != enums.WORKFLOW_ID_CONFLICT_POLICY_UNSPECIFIED {
				g.If(j.Id("swo").Dot("WorkflowIDConflictPolicy").Op("==").Qual(enumsPkg, "WORKFLOW_ID_CONFLICT_POLICY_UNSPECIFIED")).BlockFunc(func(g *j.Group) {
					var enumv string
					switch defaultWorkflowIDConflictPolicy {
					case enums.WORKFLOW_ID_CONFLICT_POLICY_FAIL:
						enumv = "WORKFLOW_ID_CONFLICT_POLICY_FAIL"
					case enums.WORKFLOW_ID_CONFLICT_POLICY_TERMINATE_EXISTING:
						enumv = "WORKFLOW_ID_CONFLICT_POLICY_TERMINATE_EXISTING"
					case enums.WORKFLOW_ID_CONFLICT_POLICY_USE_EXISTING:
						enumv = "WORKFLOW_ID_CONFLICT_POLICY_USE_EXISTING"
					}
					g.Id("swo").Dot("WorkflowIDConflictPolicy").Op("=").Qual(enumsPkg, enumv)
				})
			}

			// initialize start workflow operation
			g.Id("options").Dot("StartWorkflowOperation").Op("=").Id("op").Call(j.Id("swo"))

			// initialize update options
			g.If(j.Id("o").Dot("updateOptions").Op("==").Nil()).Block(
				j.Id("o").Dot("updateOptions").Op("=").Id(updateOptionsCtor).Call(),
			)

			// initialize workflow update options
			g.List(j.Id("uo"), j.Err()).Op(":=").Id("o").Dot("updateOptions").Dot("Build").CallFunc(func(g *j.Group) {
				g.Id("swo").Dot("ID")
				g.Lit("")
				if hasUpdateInput {
					g.Id("update")
				}
			})
			g.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Id("options"), j.Err()),
			)
			g.Id("options").Dot("UpdateOptions").Op("=").Op("*").Id("uo")

			g.ReturnFunc(func(g *j.Group) {
				g.Id("options")
				g.Nil()
			})
		})

	f.Comment("WithUpdateWithStartWorkflowOptions sets the UpdateWithStartWorkflowOptions")
	f.Func().
		Params(j.Id("o").Op("*").Id(optionsName)).
		Id("WithUpdateWithStartWorkflowOptions").
		Params(j.Id("options").Qual(clientPkg, "UpdateWithStartWorkflowOptions")).
		Params(j.Op("*").Id(optionsName)).
		BlockFunc(func(g *j.Group) {
			g.Id("o").Dot("options").Op("=").Id("options")
			g.Return(j.Id("o"))
		})

	withWorkflowOptions := m.toCamel("With%s", workflowOptions)
	f.Commentf("%s sets the %s", withWorkflowOptions, withWorkflowOptions)
	f.Func().
		Params(j.Id("o").Op("*").Id(optionsName)).
		Id(withWorkflowOptions).
		Params(j.Id("options").Op("*").Id(workflowOptions)).
		Params(j.Op("*").Id(optionsName)).
		BlockFunc(func(g *j.Group) {
			g.Id("o").Dot("workflowOptions").Op("=").Id("options")
			g.Return(j.Id("o"))
		})

	withUpdateOptions := m.toCamel("With%s", updateOptions)
	f.Commentf("%s sets the %s", withUpdateOptions, updateOptions)
	f.Func().
		Params(j.Id("o").Op("*").Id(optionsName)).
		Id(withUpdateOptions).
		Params(j.Id("options").Op("*").Id(updateOptions)).
		Params(j.Op("*").Id(optionsName)).
		BlockFunc(func(g *j.Group) {
			g.Id("o").Dot("updateOptions").Op("=").Id("options")
			g.Return(j.Id("o"))
		})
}

// genClientImplWorkflowCancelMethod generates a Cancel<Workflow> client method
func (m *Manifest) genClientImplWorkflowCancelMethod(f *j.File) {
	clientType := m.toLowerCamel("%sClient", m.Service.GoName)
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
			j.Return(
				j.Id("c").Dot("client").Dot(methodName).CallFunc(func(args *j.Group) {
					args.Id("ctx")
					args.Id("workflowID")
					args.Id("runID")
				}),
			),
		)
}

// genClientImplWorkflowGetMethod generates a Get<Workflow> client method
func (m *Manifest) genClientImplWorkflowGetMethod(f *j.File, workflow protoreflect.FullName) {
	clientType := m.toLowerCamel("%sClient", m.Service.GoName)
	methodName := m.toCamel("Get%s", workflow)
	runImplType := m.toLowerCamel("%sRun", workflow)
	runInterfaceType := m.toCamel("%sRun", workflow)

	f.Commentf("%s fetches an existing %s execution", methodName, m.fqnForWorkflow(workflow))
	f.Func().
		Params(j.Id("c").Op("*").Id(clientType)).
		Id(methodName).
		Params(
			j.Id("ctx").Qual("context", "Context"),
			j.Id("workflowID").String(),
			j.Id("runID").String(),
		).
		Params(
			j.Id(runInterfaceType),
		).
		Block(
			j.Return(
				j.Op("&").Id(runImplType).Block(
					j.Id("client").Op(":").Id("c").Op(","),
					j.Id("run").Op(":").Id("c").Dot("client").Dot("GetWorkflow").Call(
						j.Id("ctx"), j.Id("workflowID"), j.Id("runID"),
					).Op(","),
				),
			),
		)
}

// genClientImplWorkflowMethod generates an <Workflow> client method
func (m *Manifest) genClientImplWorkflowMethod(f *j.File, workflow protoreflect.FullName) {
	clientType := m.toLowerCamel("%sClient", m.Service.GoName)
	method := m.methods[workflow]
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	methodName := m.Names().clientWorkflow(workflow)
	asyncName := m.Names().clientWorkflowAsync(workflow)

	commentWithDefaultf(f, methodSet(method), "%s executes a %s workflow and blocks until error or response received", workflow, m.fqnForWorkflow(workflow))
	f.Func().
		Params(j.Id("c").Op("*").Id(clientType)).
		Id(methodName).
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			g.Id("options").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
		}).
		ParamsFunc(func(g *j.Group) {
			if hasOutput {
				g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			// execute workflow
			g.Id("run").Op(",").Err().Op(":=").Id("c").Dot(asyncName).CallFunc(func(g *j.Group) {
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
				j.Id("run").Dot("Get").Call(j.Id("ctx")),
			)
		})
}

// genClientImplWorkflowMethodAsync generates an <Workflow>Async client method
func (m *Manifest) genClientImplWorkflowMethodAsync(f *j.File, workflow protoreflect.FullName) {
	clientType := m.toLowerCamel("%sClient", m.Service.GoName)
	method := m.methods[workflow]
	methodName := m.toCamel("%sAsync", workflow)
	runImplType := m.toLowerCamel("%sRun", workflow)
	runInterfaceType := m.toCamel("%sRun", workflow)
	hasInput := !isEmpty(method.Input)
	deprecated := isDeprecated(method)

	commentWithDefaultf(f, methodSet(method), "%s starts a(n) %s workflow and returns a handle to the workflow run", methodName, m.fqnForWorkflow(workflow))
	f.Func().
		Params(j.Id("c").Op("*").Id(clientType)).
		Id(methodName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			args.Id("options").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
		}).
		Params(
			j.Id(runInterfaceType),
			j.Error(),
		).
		BlockFunc(func(fn *j.Group) {
			if deprecated {
				fn.Id("c").Dot("log").Dot("WarnContext").Call(j.Id("ctx"), j.Lit("use of deprecated client method detected"), j.Lit("method"), j.Lit(methodName), j.Lit("workflow"), j.Id(m.toCamel("%sWorkflowName", workflow))).Line()
			}

			// initialize options
			fn.Var().Id("o").Op("*").Id(m.toCamel("%sOptions", workflow))
			fn.If(j.Len(j.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(j.Lit(0)).Op("!=").Nil()).Block(
				j.Id("o").Op("=").Id("options").Index(j.Lit(0)),
			).Else().Block(
				j.Id("o").Op("=").Id(m.toCamel("New%sOptions", workflow)).Call(),
			)

			// initialize client.StartWorkfowOptions
			fn.List(j.Id("opts"), j.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(args *j.Group) {
				if hasInput {
					args.Id("req").Dot("ProtoReflect").Call()
				} else {
					args.Nil()
				}
			})
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing client.StartWorkflowOptions: %w"), j.Err())),
			)

			// execute workflow
			fn.Id("run").Op(",").Err().Op(":=").Id("c").Dot("client").Dot("ExecuteWorkflow").CallFunc(func(args *j.Group) {
				args.Id("ctx")
				args.Id("opts")
				args.Id(m.toCamel("%sWorkflowName", workflow))
				if hasInput {
					args.Id("req")
				}
			})
			fn.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Err()),
			)
			fn.If(j.Id("run").Op("==").Nil()).Block(
				j.Return(j.Nil(), j.Qual("errors", "New").Call(j.Lit("execute workflow returned nil run"))),
			)
			fn.Return(
				j.Op("&").Id(runImplType).Block(
					j.Id("client").Op(":").Id("c").Op(","),
					j.Id("run").Op(":").Id("run").Op(","),
				),
				j.Nil(),
			)
		})
}

// genClientImplWorkflowTerminateMethod generates a TerminateWorkflow client method
func (m *Manifest) genClientImplWorkflowTerminateMethod(f *j.File) {
	clientType := m.toLowerCamel("%sClient", m.Service.GoName)
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
				j.Id("c").Dot("client").Dot(methodName).CallFunc(func(args *j.Group) {
					args.Id("ctx")
					args.Id("workflowID")
					args.Id("runID")
					args.Id("reason")
					args.Id("details").Op("...")
				}),
			),
		)
}

// genClientInterface generates a Client interface for a given service
func (m *Manifest) genClientInterface(f *j.File) {
	typeName := m.toCamel("%sClient", m.Service.GoName)

	f.Commentf("%s describes a client for a(n) %s worker", typeName, m.Service.Desc.FullName())
	f.Type().Id(typeName).InterfaceFunc(func(methods *j.Group) {
		for _, workflow := range m.workflowsOrdered {
			if m.methods[workflow].Desc.Parent() != m.Service.Desc {
				continue
			}
			opts := m.workflows[workflow]

			method := m.methods[workflow]
			hasInput := !isEmpty(method.Input)
			hasOutput := !isEmpty(method.Output)
			runInterfaceType := m.toCamel("%sRun", workflow)

			// generate <Workflow> method
			methodName := m.toCamel("%s", workflow)
			commentWithDefaultf(methods, methodSet(method), "%s executes a(n) %s workflow and blocks until error or response received", methodName, m.fqnForWorkflow(workflow))
			methods.Id(methodName).
				ParamsFunc(func(args *j.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
					}
					args.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
				}).
				ParamsFunc(func(returnVals *j.Group) {
					if hasOutput {
						returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
					}
					returnVals.Error()
				}).
				Line()

			// generate <Workflow>Async method
			methodName = m.toCamel("%sAsync", workflow)
			commentf(methods, methodSet(method), "%s starts a(n) %s workflow and returns a handle to the workflow run", methodName, m.fqnForWorkflow(workflow))
			methods.Id(methodName).
				ParamsFunc(func(args *j.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
					}
					args.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
				}).
				Params(
					j.Id(runInterfaceType),
					j.Error(),
				).
				Line()

			// generate Get<Workflow> method
			methodName = m.toCamel("Get%s", workflow)
			commentf(methods, methodSet(method), "%s retrieves a handle to an existing %s workflow execution", methodName, m.fqnForWorkflow(workflow))
			methods.Id(m.toCamel("Get%s", workflow)).
				Params(
					j.Id("ctx").Qual("context", "Context"),
					j.Id("workflowID").String(),
					j.Id("runID").String(),
				).
				Params(
					j.Id(runInterfaceType),
				).
				Line()

			// add <Workflow>With<Signal> methods
			for _, signalOpts := range opts.GetSignal() {
				if !signalOpts.GetStart() {
					continue
				}
				method := m.methods[workflow]
				signal := getFullyQualifiedRef(workflow, signalOpts.GetRef())
				handler := m.methods[signal]
				hasWorkflowInput := !isEmpty(method.Input)
				hasWorkflowOutput := !isEmpty(method.Output)
				hasSignalInput := !isEmpty(handler.Input)

				// add synchronous flavor
				methodName := m.toCamel("%sWith%s", workflow, signal)
				commentf(methods, methodSet(method, handler), "%s sends a(n) %s signal to a(n) %s workflow, starting it if necessary, and blocks until workflow completion", methodName, m.fqnForSignal(signal), m.fqnForWorkflow(workflow))
				methods.Id(methodName).
					ParamsFunc(func(args *j.Group) {
						args.Id("ctx").Qual("context", "Context")
						if hasWorkflowInput {
							args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
						}
						if hasSignalInput {
							args.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
						}
						args.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
					}).
					ParamsFunc(func(returnVals *j.Group) {
						if hasWorkflowOutput {
							returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
						}
						returnVals.Error()
					}).
					Line()

				// add async flavor
				methodName += "Async"
				commentf(methods, methodSet(method, handler), "%s sends a(n) %s signal to a(n) %s workflow, starting it if necessary, and returns a handle to the workflow execution", methodName, m.fqnForSignal(signal), m.fqnForWorkflow(workflow))
				methods.Id(methodName).
					ParamsFunc(func(args *j.Group) {
						args.Id("ctx").Qual("context", "Context")
						if hasWorkflowInput {
							args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
						}
						if hasSignalInput {
							args.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
						}
						args.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
					}).
					Params(
						j.Id(runInterfaceType),
						j.Error(),
					).
					Line()
			}

			// add <Workflow>With<Update> methods
			for _, updateOpts := range opts.GetUpdate() {
				if !updateOpts.GetStart() {
					continue
				}
				method := m.methods[workflow]
				update := getFullyQualifiedRef(workflow, updateOpts.GetRef())
				handler := m.methods[update]
				hasWorkflowInput := !isEmpty(method.Input)
				hasUpdateInput := !isEmpty(handler.Input)
				hasUpdateOutput := !isEmpty(handler.Output)
				optionsName := m.Names().clientUpdateWithStartOptions(workflow, update)
				runName := m.Names().clientWorkflowRun(workflow)
				handleName := m.Names().clientUpdateHandleIface(update)

				methodName := m.Names().clientUpdateWithStart(workflow, update)
				commentf(methods, methodSet(method, handler), "%s executes a(n) %s update on a(n) %s workflow, starting it if necessary, and blocks until update completion", methodName, m.fqnForUpdate(update), m.fqnForWorkflow(workflow))
				methods.Id(methodName).
					ParamsFunc(func(args *j.Group) {
						args.Id("ctx").Qual("context", "Context")
						if hasWorkflowInput {
							args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
						}
						if hasUpdateInput {
							args.Id("update").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
						}
						args.Id("opts").Op("...").Op("*").Id(optionsName)
					}).
					ParamsFunc(func(returnVals *j.Group) {
						if hasUpdateOutput {
							returnVals.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
						}
						returnVals.Id(runName)
						returnVals.Error()
					})

				asyncName := m.Names().clientUpdateWithStartAsync(workflow, update)
				commentf(methods, methodSet(method, handler), "%s starts a(n) %s update on a(n) %s workflow, starting it if necessary, and returns a handle to the update execution", asyncName, m.fqnForUpdate(update), m.fqnForWorkflow(workflow))
				methods.Id(asyncName).
					ParamsFunc(func(args *j.Group) {
						args.Id("ctx").Qual("context", "Context")
						if hasWorkflowInput {
							args.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
						}
						if hasUpdateInput {
							args.Id("update").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
						}
						args.Id("opts").Op("...").Op("*").Id(optionsName)
					}).
					Params(j.Id(handleName), j.Id(runName), j.Error())
			}
		}

		// generate CancelWorkflow method
		methodName := "CancelWorkflow"
		methods.Commentf("%s requests cancellation of an existing workflow execution", methodName)
		methods.Id(methodName).
			Params(
				j.Id("ctx").Qual("context", "Context"),
				j.Id("workflowID").String(),
				j.Id("runID").String(),
			).
			Params(j.Error()).
			Line()

		// generate TerminateWorkflow method
		methodName = "TerminateWorkflow"
		methods.Commentf("%s an existing workflow execution", methodName)
		methods.Id(methodName).
			Params(
				j.Id("ctx").Qual("context", "Context"),
				j.Id("workflowID").String(),
				j.Id("runID").String(),
				j.Id("reason").String(),
				j.Id("details").Op("...").Interface(),
			).
			Params(j.Error()).
			Line()

		// add <Query> methods
		for _, query := range m.queriesOrdered {
			if m.methods[query].Desc.Parent() != m.Service.Desc {
				continue
			}
			handler := m.methods[query]
			hasInput := !isEmpty(handler.Input)
			commentWithDefaultf(methods, methodSet(handler), "%s executes a(n) %s query", query, m.fqnForQuery(query))
			methods.Id(m.toCamel("%s", query)).
				ParamsFunc(func(args *j.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("workflowID").String()
					args.Id("runID").String()
					if hasInput {
						args.Id("query").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
				}).
				Params(
					j.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output)),
					j.Error(),
				).
				Line()
		}

		// add <Signal> methods
		for _, signal := range m.signalsOrdered {
			if m.methods[signal].Desc.Parent() != m.Service.Desc {
				continue
			}
			handler := m.methods[signal]
			hasInput := !isEmpty(handler.Input)
			commentWithDefaultf(methods, methodSet(handler), "%s sends a(n) %s signal", signal, m.fqnForSignal(signal))
			methods.Id(m.toCamel("%s", signal)).
				ParamsFunc(func(args *j.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("workflowID").String()
					args.Id("runID").String()
					if hasInput {
						args.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
				}).
				Params(j.Error()).
				Line()
		}

		// add <Update> methods
		for _, update := range m.updatesOrdered {
			if m.methods[update].Desc.Parent() != m.Service.Desc {
				continue
			}
			handler := m.methods[update]
			hasInput := !isEmpty(handler.Input)
			hasOutput := !isEmpty(handler.Output)

			// add synchronous flavor
			methodName := m.toCamel("%s", update)
			commentWithDefaultf(methods, methodSet(handler), "%s executes a(n) %s update and blocks until update completion", methodName, m.fqnForUpdate(update))
			methods.Id(methodName).
				ParamsFunc(func(args *j.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("workflowID").String()
					args.Id("runID").String()
					if hasInput {
						args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
					args.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", update))
				}).
				ParamsFunc(func(returnVals *j.Group) {
					if hasOutput {
						returnVals.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
					}
					returnVals.Error()
				}).
				Line()

			// add async flavor
			methodName = m.toCamel("%sAsync", update)
			commentf(methods, methodSet(handler), "%s starts a(n) %s update and returns a handle to the workflow update", methodName, m.fqnForUpdate(update))
			methods.Id(methodName).
				ParamsFunc(func(args *j.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("workflowID").String()
					args.Id("runID").String()
					if hasInput {
						args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
					args.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", update))
				}).
				Params(
					j.Id(m.toCamel("%sHandle", update)),
					j.Error(),
				).
				Line()

			// add getter
			methodName = m.toCamel("Get%s", update)
			commentf(methods, methodSet(handler), "%s retrieves a handle to an existing %s update", methodName, m.fqnForUpdate(update))
			methods.Id(methodName).
				ParamsFunc(func(args *j.Group) {
					args.Id("ctx").Qual("context", "Context")
					args.Id("req").Qual(clientPkg, "GetWorkflowUpdateHandleOptions")
				}).
				Params(
					j.Id(m.toCamel("%sHandle", update)),
					j.Error(),
				).
				Line()
		}
	})
}

func (m *Manifest) genClientOptions(f *j.File) {
	typeName := m.toLowerCamel("%sClientOptions", m.Service.GoName)

	f.Commentf("%s describes optional runtime configuration for a %s", typeName, m.toCamel("%sClient", m.Service.GoName))
	f.Type().Id(typeName).Struct(
		j.Id("log").Op("*").Qual("log/slog", "Logger"),
	)

	constructorName := m.toCamel("New%sClientOptions", m.Service.GoName)
	f.Commentf("%s initializes a new %s value", constructorName, typeName)
	f.Func().Id(constructorName).Params().Op("*").Id(typeName).Block(
		j.Return(j.Op("&").Id(typeName).Values()),
	)

	f.Comment("WithLogger can be used to override the default logger")
	f.Func().
		Params(j.Id("opts").Op("*").Id(typeName)).
		Id("WithLogger").
		Params(j.Id("l").Op("*").Qual("log/slog", "Logger")).
		Op("*").Id(typeName).
		Block(
			j.If(j.Id("l").Op("!=").Nil()).Block(
				j.Id("opts").Dot("log").Op("=").Id("l"),
			),
			j.Return(j.Id("opts")),
		)

	f.Comment("getLogger returns the configured logger, or the default logger")
	f.Func().
		Params(j.Id("opts").Op("*").Id(typeName)).
		Id("getLogger").
		Params().
		Op("*").Qual("log/slog", "Logger").
		Block(
			j.If(j.Id("opts").Op("!=").Nil().Op("&&").Id("opts").Dot("log").Op("!=").Nil()).Block(
				j.Return(j.Id("opts").Dot("log")),
			),
			j.Return(j.Qual("log/slog", "Default").Call()),
		)
}

// genClientUpdateHandleImpl generates a <Update>Handle struct
func (m *Manifest) genClientUpdateHandleImpl(f *j.File, update protoreflect.FullName) {
	clientImplType := m.toLowerCamel("%sClient", m.Service.GoName)
	typeName := m.toLowerCamel("%sHandle", update)
	interfaceName := m.toCamel("%sHandle", update)

	// generate struct
	f.Commentf("%s provides an internal implementation of a(n) %s", typeName, interfaceName)
	f.Type().
		Id(typeName).
		Struct(
			j.Id("client").Op("*").Id(clientImplType),
			j.Id("handle").Qual(clientPkg, "WorkflowUpdateHandle"),
		)
}

// genClientUpdateHandleImplGetMethod generates a <UpdateHandle>'s Get method
func (m *Manifest) genClientUpdateHandleImplGetMethod(f *j.File, update protoreflect.FullName) {
	typeName := m.toLowerCamel("%sHandle", update)
	method := m.methods[update]
	hasOutput := !isEmpty(method.Output)

	f.Comment("Get blocks until the update wait policy is met, returning the result if applicable")
	f.Func().
		Params(j.Id("h").Op("*").Id(typeName)).
		Id("Get").
		Params(j.Id("ctx").Qual("context", "Context")).
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
			fn.Var().Err().Error()
			fn.Id("doneCh").Op(":=").Make(j.Chan().Struct())
			fn.List(j.Id("gctx"), j.Id("cancel")).Op(":=").Qual("context", "WithCancel").Call(j.Qual("context", "Background").Call())
			fn.Defer().Id("cancel").Call()
			fn.Line()

			fn.Go().Func().Params().Block(
				j.For().Block(
					j.Var().Id("deadlineExceeded").Op("*").Qual(serviceerrorPkg, "DeadlineExceeded"),
					j.If(
						j.Err().Op("=").Id("h").Dot("handle").Dot("Get").CallFunc(func(args *j.Group) {
							args.Id("gctx")
							if hasOutput {
								args.Op("&").Id("resp")
							} else {
								args.Nil()
							}
						}),
						j.Err().Op("!=").Nil().Op("&&").
							Id("ctx").Dot("Err").Call().Op("==").Nil().Op("&&").
							Parens(
								j.Qual("errors", "As").Call(j.Id("err"), j.Op("&").Id("deadlineExceeded")).Op("||").
									Qual("strings", "Contains").Call(
									j.Err().Dot("Error").Call(),
									j.Qual("context", "DeadlineExceeded").Dot("Error").Call(),
								),
							),
					).Block(
						j.Continue(),
					),
					j.Break(),
				),
				j.Close(j.Id("doneCh")),
			).Call()
			fn.Line()

			fn.Select().Block(
				j.Case(j.Op("<-").Id("ctx").Dot("Done").Call()).Block(
					j.ReturnFunc(func(returnVals *j.Group) {
						if hasOutput {
							returnVals.Nil()
						}
						returnVals.Id("ctx").Dot("Err").Call()
					}),
				),
				j.Case(j.Op("<-").Id("doneCh")).BlockFunc(func(bl *j.Group) {
					if hasOutput {
						bl.If(j.Err().Op("!=").Nil()).Block(
							j.Return(j.Nil(), j.Err()),
						)
						bl.Return(j.Op("&").Id("resp"), j.Nil())
					} else {
						bl.Return(j.Err())
					}
				}),
			)
		})
}

// genClientUpdateHandleImplRunIDMethod generates a <UpdateHandle>'s RunID method
func (m *Manifest) genClientUpdateHandleImplRunIDMethod(f *j.File, update protoreflect.FullName) {
	typeName := m.toLowerCamel("%sHandle", update)

	f.Comment("RunID returns the execution ID")
	f.Func().
		Params(j.Id("h").Op("*").Id(typeName)).
		Id("RunID").
		Params().
		String().
		Block(
			j.Return(j.Id("h").Dot("handle").Dot("RunID").Call()),
		)
}

// genClientUpdateHandleImplUpdateIDMethod generates a <UpdateHandle>'s UpdateID method
func (m *Manifest) genClientUpdateHandleImplUpdateIDMethod(f *j.File, update protoreflect.FullName) {
	typeName := m.toLowerCamel("%sHandle", update)

	f.Comment("UpdateID returns the update ID")
	f.Func().
		Params(j.Id("h").Op("*").Id(typeName)).
		Id("UpdateID").
		Params().
		String().
		Block(
			j.Return(j.Id("h").Dot("handle").Dot("UpdateID").Call()),
		)
}

// genClientUpdateHandleImplWorkflowIDMethod generates a <UpdateHandle>'s WorkflowID method
func (m *Manifest) genClientUpdateHandleImplWorkflowIDMethod(f *j.File, update protoreflect.FullName) {
	typeName := m.toLowerCamel("%sHandle", update)

	f.Comment("WorkflowID returns the workflow ID")
	f.Func().
		Params(j.Id("h").Op("*").Id(typeName)).
		Id("WorkflowID").
		Params().
		String().
		Block(
			j.Return(j.Id("h").Dot("handle").Dot("WorkflowID").Call()),
		)
}

// genClientUpdateHandleInterface generates a <Workflow>Run interface
func (m *Manifest) genClientUpdateHandleInterface(f *j.File, update protoreflect.FullName) {
	typeName := m.toCamel("%sHandle", update)
	method := m.methods[update]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s describes a(n) %s update handle", typeName, m.fqnForUpdate(update))
	f.Type().Id(typeName).InterfaceFunc(func(methods *j.Group) {
		methods.Comment("WorkflowID returns the workflow ID")
		methods.Id("WorkflowID").Params().String()

		methods.Comment("RunID returns the workflow instance ID")
		methods.Id("RunID").Params().String()

		methods.Comment("UpdateID returns the update ID")
		methods.Id("UpdateID").Params().String()

		methods.Comment("Get blocks until the workflow is complete and returns the result")
		methods.Id("Get").
			Params(j.Id("ctx").Qual("context", "Context")).
			ParamsFunc(func(returnVals *j.Group) {
				if hasOutput {
					returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
				}
				returnVals.Error()
			})
	})
}

// genClientUpdateOptions generates a <Update>Options struct
func (m *Manifest) genClientUpdateOptions(f *j.File, update protoreflect.FullName) {
	typeName := m.toCamel("%sOptions", update)
	constructorName := "New" + typeName
	updateOpts := m.updates[update]
	hasInput := !isEmpty(m.methods[update].Input)

	f.Commentf("%s provides configuration for a %s update operation", typeName, m.fqnForUpdate(update))
	f.Type().Id(typeName).Struct(
		j.Id("Options").Op("*").Qual(clientPkg, "UpdateWorkflowOptions"),
		j.Id("id").Op("*").String(),
		j.Id("waitPolicy").Qual(clientPkg, "WorkflowUpdateStage"),
	)

	f.Commentf("%s initializes a new %s value", constructorName, typeName)
	f.Func().Id(constructorName).Params().Op("*").Id(typeName).Block(
		j.Return(j.Op("&").Id(typeName).Values(
			j.Id("Options").Op(":").Op("&").Qual(clientPkg, "UpdateWorkflowOptions").Values(),
		)),
	)

	f.Comment("Build initializes a new client.UpdateWorkflowOptions with defaults and overrides applied")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("Build").
		ParamsFunc(func(args *j.Group) {
			args.Id("workflowID").String()
			args.Id("runID").String()
			if hasInput {
				args.Id("req").Op("*").Qual(string(m.methods[update].Input.GoIdent.GoImportPath), m.getMessageName(m.methods[update].Input))
			}
		}).
		Params(
			j.Id("opts").Op("*").Qual(clientPkg, "UpdateWorkflowOptions"),
			j.Err().Error(),
		).
		BlockFunc(func(fn *j.Group) {
			fn.Comment("use user-provided UpdateWorkflowOptions if exists")
			fn.If(j.Id("o").Dot("Options").Op("!=").Nil()).Block(
				j.Id("opts").Op("=").Id("o").Dot("Options"),
			).Else().Block(
				j.Id("opts").Op("=").Op("&").Qual(clientPkg, "UpdateWorkflowOptions").Values(),
			)

			fn.Line()
			fn.Comment("set constants")
			if hasInput {
				fn.Id("opts").Dot("Args").Op("=").Index().Any().Values(j.Id("req"))
			}
			fn.Id("opts").Dot("RunID").Op("=").Id("runID")
			fn.Id("opts").Dot("UpdateName").Op("=").Id(m.toCamel("%sUpdateName", update))
			fn.Id("opts").Dot("WorkflowID").Op("=").Id("workflowID")

			fn.Line()
			fn.Comment("set UpdateID")
			id := fn.If(j.Id("v").Op(":=").Id("o").Dot("id"), j.Id("v").Op("!=").Nil()).Block(
				j.Id("opts").Dot("UpdateID").Op("=").Op("*").Id("v"),
			)
			if idExpr := updateOpts.GetId(); idExpr != "" {
				id.Else().If(j.Id("opts").Dot("UpdateID").Op("==").Lit("")).BlockFunc(func(b *j.Group) {
					b.List(j.Id("id"), j.Err()).Op(":=").Qual(expressionPkg, "EvalExpression").CallFunc(func(args *j.Group) {
						args.Id(m.toCamel("%sIDExpression", update))
						if hasInput {
							args.Id("req").Dot("ProtoReflect").Call()
						} else {
							args.Nil()
						}
					})
					b.If(j.Err().Op("!=").Nil()).Block(
						j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error evaluating id expression for %q update: %w"), j.Id(m.toCamel("%sUpdateName", update)), j.Err())),
					)
					b.Id("opts").Dot("UpdateID").Op("=").Id("id")
				})
			}

			fn.Line()
			fn.Comment("set WaitPolicy")
			waitPolicy := fn.If(j.Id("v").Op(":=").Id("o").Dot("waitPolicy"), j.Id("v").Op("!=").Qual(clientPkg, "WorkflowUpdateStageUnspecified")).Block(
				j.Id("opts").Dot("WaitForStage").Op("=").Id("v"),
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
			waitPolicy.Else().If(j.Id("opts").Dot("WaitForStage").Op("==").Qual(clientPkg, "WorkflowUpdateStageUnspecified")).Block(
				j.Id("opts").Dot("WaitForStage").Op("=").Qual(clientPkg, stage),
			)

			fn.Return(j.Id("opts"), j.Nil())
		})

	f.Comment("WithUpdateID sets the UpdateID")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithUpdateID").
		Params(j.Id("id").String()).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("id").Op("=").Op("&").Id("id"),
			j.Return(j.Id("o")),
		)

	f.Comment("WithUpdateWorkflowOptions sets the initial client.UpdateWorkflowOptions")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithUpdateWorkflowOptions").
		Params(j.Id("options").Qual(clientPkg, "UpdateWorkflowOptions")).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("Options").Op("=").Op("&").Id("options"),
			j.Return(j.Id("o")),
		)

	f.Comment("WithWaitPolicy sets the WaitPolicy")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithWaitPolicy").
		Params(j.Id("policy").Qual(clientPkg, "WorkflowUpdateStage")).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("waitPolicy").Op("=").Id("policy"),
			j.Return(j.Id("o")),
		)
}

// genClientWorkflowRunImpl generates a <Workflow>Run struct
func (m *Manifest) genClientWorkflowRunImpl(f *j.File, workflow protoreflect.FullName) {
	clientType := m.toLowerCamel("%sClient", m.Service.GoName)
	typeName := m.toLowerCamel("%sRun", workflow)
	interfaceName := m.toCamel("%sRun", workflow)

	// generate struct
	f.Commentf("%s provides an internal implementation of a(n) %sRun", typeName, interfaceName)
	f.Type().
		Id(typeName).
		Struct(
			j.Id("client").Op("*").Id(clientType),
			j.Id("run").Qual(clientPkg, "WorkflowRun"),
		)
}

// genClientWorkflowRunImplCancelMethod generates a <Workflow>Run's Cancel method
func (m *Manifest) genClientWorkflowRunImplCancelMethod(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toLowerCamel("%sRun", workflow)

	f.Comment("Cancel requests cancellation of a workflow in execution, returning an error if applicable")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("Cancel").
		Params(j.Id("ctx").Qual("context", "Context")).
		Params(j.Error()).
		Block(
			j.Return(
				j.Id("r").Dot("client").Dot("CancelWorkflow").CallFunc(func(args *j.Group) {
					args.Id("ctx")
					args.Id("r").Dot("ID").Call()
					args.Id("r").Dot("RunID").Call()
				}),
			),
		)
}

// genClientWorkflowRunImplGetMethod generates a <Workflow>Run's Get method
func (m *Manifest) genClientWorkflowRunImplGetMethod(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toLowerCamel("%sRun", workflow)
	method := m.methods[workflow]
	hasOutput := !isEmpty(method.Output)

	f.Comment("Get blocks until the workflow is complete, returning the result if applicable")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("Get").
		Params(j.Id("ctx").Qual("context", "Context")).
		ParamsFunc(func(returnVals *j.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			returnVals.Error()
		}).
		BlockFunc(func(fn *j.Group) {
			if hasOutput {
				fn.Var().Id("resp").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
				fn.If(
					j.Err().Op(":=").Id("r").Dot("run").Dot("Get").Call(
						j.Id("ctx"),
						j.Op("&").Id("resp"),
					),
					j.Err().Op("!=").Nil(),
				).Block(
					j.Return(
						j.Nil(), j.Err(),
					),
				)
				fn.Return(
					j.Op("&").Id("resp"), j.Nil(),
				)
			} else {
				fn.Return(
					j.Id("r").Dot("run").Dot("Get").Call(
						j.Id("ctx"),
						j.Nil(),
					),
				)
			}
		})
}

// genClientWorkflowRunImplIDMethod generates a <Workflow>Run's ID method
func (m *Manifest) genClientWorkflowRunImplIDMethod(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toLowerCamel("%sRun", workflow)

	f.Comment("ID returns the workflow ID")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("ID").
		Params().
		String().
		Block(
			j.Return(j.Id("r").Dot("run").Dot("GetID").Call()),
		)
}

// genClientWorkflowRunImplQueryMethod generates a <WOrkflow>Run's <Query> method
func (m *Manifest) genClientWorkflowRunImplQueryMethod(f *j.File, workflow protoreflect.FullName, query protoreflect.FullName) {
	typeName := m.toLowerCamel("%sRun", workflow)
	handler := m.methods[query]
	hasInput := !isEmpty(handler.Input)

	commentWithDefaultf(f, methodSet(handler), "%s executes a(n) %s query", query, m.fqnForQuery(query))
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id(m.methods[query].GoName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
		}).
		Params(
			j.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output)),
			j.Error(),
		).
		Block(
			j.Return(
				j.Id("r").Dot("client").Dot(m.methods[query].GoName).CallFunc(func(args *j.Group) {
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
func (m *Manifest) genClientWorkflowRunImplRunMethod(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toLowerCamel("%sRun", workflow)

	f.Comment("Run returns the inner client.WorkflowRun")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("Run").
		Params().
		Qual(clientPkg, "WorkflowRun").
		Block(
			j.Return(j.Id("r").Dot("run")),
		)
}

// genClientWorkflowRunImplRunIDMethod generates a <Workflow>Run's RunID method
func (m *Manifest) genClientWorkflowRunImplRunIDMethod(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toLowerCamel("%sRun", workflow)

	f.Comment("RunID returns the execution ID")
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id("RunID").
		Params().
		String().
		Block(
			j.Return(j.Id("r").Dot("run").Dot("GetRunID").Call()),
		)
}

// genClientWorkflowRunImplSignalMethod generates a <Workflow>Run's <Signal> method
func (m *Manifest) genClientWorkflowRunImplSignalMethod(f *j.File, workflow protoreflect.FullName, signal protoreflect.FullName) {
	typeName := m.toLowerCamel("%sRun", workflow)
	handler := m.methods[signal]
	hasInput := !isEmpty(handler.Input)

	// generate get method
	commentWithDefaultf(f, methodSet(handler), "%s sends a(n) %s signal", signal, m.fqnForSignal(signal))
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id(m.methods[signal].GoName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
		}).
		Params(j.Error()).
		Block(
			j.ReturnFunc(func(returnVals *j.Group) {
				if m.methodsFromSameService(signal, workflow) {
					returnVals.Id("r").Dot("client").Dot(m.methods[signal].GoName).CallFunc(func(args *j.Group) {
						args.Id("ctx")
						args.Id("r").Dot("ID").Call()
						args.Lit("")
						if hasInput {
							args.Id("req")
						}
					})
				} else {
					returnVals.Id("r").Dot("client").Dot("client").Dot("SignalWorkflow").CallFunc(func(args *j.Group) {
						args.Id("ctx")
						args.Id("r").Dot("ID").Call()
						args.Id("r").Dot("RunID").Call()
						args.Add(m.Qual(signal, m.toCamel("%sSignalName", signal)))
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
func (m *Manifest) genClientWorkflowRunImplTerminateMethod(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toLowerCamel("%sRun", workflow)

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
				j.Id("r").Dot("client").Dot("TerminateWorkflow").CallFunc(func(args *j.Group) {
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
func (m *Manifest) genClientWorkflowRunImplUpdateAsyncMethod(f *j.File, workflow protoreflect.FullName, update protoreflect.FullName) {
	typeName := m.toLowerCamel("%sRun", workflow)
	methodName := m.toCamel("%sAsync", update)
	handler := m.methods[update]
	hasInput := !isEmpty(handler.Input)

	commentWithDefaultf(f, methodSet(handler), "%s start a(n) %s workflow update and returns a handle to the update", methodName, m.fqnForUpdate(update))
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id(methodName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", update))
		}).
		Params(
			j.Add(m.Qual(update, m.toCamel("%sHandle", update))),
			j.Error(),
		).
		BlockFunc(func(fn *j.Group) {
			if m.methodsFromSamePackage(workflow, update) {
				fn.Return(
					j.Id("r").Dot("client").Dot(m.toCamel("%sAsync", update)).CallFunc(func(args *j.Group) {
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
				fn.Var().Id("o").Op("*").Id(m.toCamel("%sOptions", update))
				fn.If(j.Len(j.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(j.Lit(0)).Op("!=").Nil()).Block(
					j.Id("o").Op("=").Id("opts").Index(j.Lit(0)),
				).Else().Block(
					j.Id("o").Op("=").Id(m.toCamel("New%sOptions", update)).Call(),
				)

				// build UpdateWorkflowWithOptions
				fn.List(j.Id("options"), j.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(args *j.Group) {
					args.Id("r").Dot("ID").Call()
					args.Id("r").Dot("RunID").Call()
					if hasInput {
						args.Id("req")
					}
				})
				fn.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing UpdateWorkflowWithOptions: %w"), j.Err())),
				)

				// update workflow
				fn.List(j.Id("handle"), j.Err()).Op(":=").Id("c").Dot("client").Dot("UpdateWorkflowWithOptions").Call(j.Id("ctx"), j.Id("options"))
				fn.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Nil(), j.Err()),
				)

				fn.Return(
					j.Op("&").Add(m.Qual(update, m.toCamel("%sHandle", update))).Values(
						j.Id("client").Op(":").Id("c"),
						j.Id("handle").Op(":").Id("handle"),
					),
					j.Nil(),
				)
			}
		})
}

// genClientWorkflowRunImplUpdateMethod generates a <Workflow>Run's <Update> method
func (m *Manifest) genClientWorkflowRunImplUpdateMethod(f *j.File, workflow protoreflect.FullName, update protoreflect.FullName) {
	typeName := m.toLowerCamel("%sRun", workflow)
	handler := m.methods[update]
	hasInput := !isEmpty(handler.Input)
	hasOutput := !isEmpty(handler.Output)

	commentWithDefaultf(f, methodSet(handler), "%s executes a(n) %s workflow update", update, m.fqnForUpdate(update))
	f.Func().
		Params(j.Id("r").Op("*").Id(typeName)).
		Id(m.methods[update].GoName).
		ParamsFunc(func(args *j.Group) {
			args.Id("ctx").Qual("context", "Context")
			if hasInput {
				args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			args.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", update))
		}).
		ParamsFunc(func(returnVals *j.Group) {
			if hasOutput {
				returnVals.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
			}
			returnVals.Error()
		}).
		Block(
			j.Return(
				j.Id("r").Dot("client").Dot(m.methods[update].GoName).CallFunc(func(args *j.Group) {
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
func (m *Manifest) genClientWorkflowRunInterface(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toCamel("%sRun", workflow)
	opts := m.workflows[workflow]
	method := m.methods[workflow]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s describes a(n) %s workflow run", typeName, m.fqnForWorkflow(workflow))
	f.Type().Id(typeName).InterfaceFunc(func(methods *j.Group) {
		methods.Comment("ID returns the workflow ID")
		methods.Id("ID").Params().String().Line()

		methods.Comment("RunID returns the workflow instance ID")
		methods.Id("RunID").Params().String().Line()

		methods.Comment("Run returns the inner client.WorkflowRun")
		methods.Id("Run").Params().Qual(clientPkg, "WorkflowRun").Line()

		methods.Comment("Get blocks until the workflow is complete and returns the result")
		methods.Id("Get").
			Params(j.Id("ctx").Qual("context", "Context")).
			ParamsFunc(func(returnVals *j.Group) {
				if hasOutput {
					returnVals.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
				}
				returnVals.Error()
			}).Line()

		methods.Comment("Cancel requests cancellation of a workflow in execution, returning an error if applicable")
		methods.Id("Cancel").
			Params(j.Id("ctx").Qual("context", "Context")).
			Params(j.Error()).Line()

		methods.Comment("Terminate terminates a workflow in execution, returning an error if applicable")
		methods.Id("Terminate").
			Params(
				j.Id("ctx").Qual("context", "Context"),
				j.Id("reason").String(),
				j.Id("details").Op("...").Interface(),
			).
			Params(j.Error()).Line()

		for _, queryOpts := range opts.GetQuery() {
			query := getFullyQualifiedRef(workflow, queryOpts.GetRef())
			handler := m.methods[query]
			hasInput := !isEmpty(handler.Input)

			commentWithDefaultf(methods, methodSet(handler), "%s executes a(n) %s query", query, m.fqnForQuery(query))
			methods.Id(m.toCamel("%s", query)).
				ParamsFunc(func(args *j.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
				}).
				Params(
					j.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output)),
					j.Error(),
				).Line()
		}

		for _, signalOpts := range opts.GetSignal() {
			signal := getFullyQualifiedRef(workflow, signalOpts.GetRef())
			handler := m.methods[signal]
			hasInput := !isEmpty(handler.Input)

			commentWithDefaultf(methods, methodSet(handler), "%s sends a(n) %s signal", signal, m.fqnForSignal(signal))
			methods.Id(m.toCamel("%s", signal)).
				ParamsFunc(func(args *j.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
				}).
				Params(j.Error()).Line()
		}

		for _, updateOpts := range opts.GetUpdate() {
			update := getFullyQualifiedRef(workflow, updateOpts.GetRef())
			handler := m.methods[update]
			hasInput := !isEmpty(handler.Input)
			hasOutput := !isEmpty(handler.Output)

			// add synchronous flavor
			commentWithDefaultf(methods, methodSet(handler), "%s executes a(n) %s update", update, m.fqnForUpdate(update))
			methods.Id(m.toCamel("%s", update)).
				ParamsFunc(func(args *j.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
					args.Id("opts").Op("...").Op("*").Add(m.Qual(update, m.toCamel("%sOptions", update)))
				}).
				ParamsFunc(func(returnVals *j.Group) {
					if hasOutput {
						returnVals.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
					}
					returnVals.Error()
				}).Line()

			// add async flavor
			commentWithDefaultf(methods, methodSet(handler), "%sAsync sends a(n) %s update to the workflow", update, m.fqnForUpdate(update))
			methods.Id(m.toCamel("%sAsync", update)).
				ParamsFunc(func(args *j.Group) {
					args.Id("ctx").Qual("context", "Context")
					if hasInput {
						args.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
					args.Id("opts").Op("...").Op("*").Add(m.Qual(update, m.toCamel("%sOptions", update)))
				}).
				Params(
					j.Add(m.Qual(update, m.toCamel("%sHandle", update))),
					j.Error(),
				).Line()
		}
	})
}

// genWorkflowOptions generates a <Workflow>Options struct
func (m *Manifest) genWorkflowOptions(f *j.File, workflow protoreflect.FullName, child bool) {
	optionsPkg := clientPkg
	optionsType := "StartWorkflowOptions"
	typeName := m.toCamel("%sOptions", workflow)
	childQualifier := ""
	if child {
		optionsPkg = workflowPkg
		optionsType = "ChildWorkflowOptions"
		typeName = m.toCamel("%sChildOptions", workflow)
		childQualifier = "child "
	}
	constructorName := "New" + typeName
	opts := m.workflows[workflow]

	f.Commentf("%s provides configuration for a %s%s workflow operation", typeName, childQualifier, m.fqnForWorkflow(workflow))
	f.Type().Id(typeName).StructFunc(func(values *j.Group) {
		values.Id("options").Qual(optionsPkg, optionsType)
		values.Id("executionTimeout").Op("*").Qual("time", "Duration")
		values.Id("id").Op("*").String()
		values.Id("idReusePolicy").Qual(enumsPkg, "WorkflowIdReusePolicy")
		values.Id("retryPolicy").Op("*").Qual(temporalPkg, "RetryPolicy")
		values.Id("runTimeout").Op("*").Qual("time", "Duration")
		values.Id("searchAttributes").Map(j.String()).Any()
		values.Id("taskQueue").Op("*").String()
		values.Id("taskTimeout").Op("*").Qual("time", "Duration")
		values.Id("workflowIdConflictPolicy").Qual(enumsPkg, "WorkflowIdConflictPolicy")
		if child {
			values.Id("parentClosePolicy").Qual(enumsPkg, "ParentClosePolicy")
			values.Id("waitForCancellation").Op("*").Bool()
		}
	})

	f.Commentf("%s initializes a new %s value", constructorName, typeName)
	f.Func().Id(constructorName).Params().Op("*").Id(typeName).Block(
		j.Return(j.Op("&").Id(typeName).Values()),
	)

	f.Commentf("Build initializes a new %s.%s value with defaults and overrides applied", optionsPkg, optionsType)
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("Build").
		ParamsFunc(func(args *j.Group) {
			if child {
				args.Id("ctx").Qual(workflowPkg, "Context")
			}
			args.Id("req").Qual(protoreflectPkg, "Message")
		}).
		Params(
			j.Qual(optionsPkg, optionsType),
			j.Error(),
		).
		BlockFunc(func(fn *j.Group) {
			fn.Id("opts").Op(":=").Id("o").Dot("options")

			// set ID
			idFieldName := "ID"
			if child {
				idFieldName = "WorkflowID"
			}
			id := fn.If(j.Id("v").Op(":=").Id("o").Dot("id"), j.Id("v").Op("!=").Nil()).Block(
				j.Id("opts").Dot(idFieldName).Op("=").Op("*").Id("v"),
			)
			if idExpr := opts.GetId(); idExpr != "" {
				id.Else().If(j.Id("opts").Dot(idFieldName).Op("==").Lit("")).BlockFunc(func(b *j.Group) {
					// original expression evaluation logic
					origFn := func(b *j.Group, errorReturn j.Code) {
						b.List(j.Id("id"), j.Err()).Op(":=").Qual(expressionPkg, "EvalExpression").CallFunc(func(args *j.Group) {
							args.Id(m.toCamel("%sIDExpression", workflow))
							args.Id("req")
						})
						b.If(j.Err().Op("!=").Nil()).BlockFunc(func(returnVals *j.Group) {
							returnVals.Return(errorReturn, j.Qual("fmt", "Errorf").Call(j.Lit("error evaluating id expression for %q workflow: %w"), j.Id(m.toCamel("%sWorkflowName", workflow)), j.Err()))
						})
					}

					if child {
						// local activity wrapper
						fixFn := func(b *j.Group) {
							b.Id("lao").Op(":=").Qual(workflowPkg, "GetLocalActivityOptions").Call(j.Id("ctx"))
							b.Id("lao").Dot("ScheduleToCloseTimeout").Op("=").Qual("time", "Second").Op("*").Lit(10)
							b.If(
								j.Err().Op(":=").Qual(workflowPkg, "ExecuteLocalActivity").Call(
									j.Qual(workflowPkg, "WithLocalActivityOptions").Call(j.Id("ctx"), j.Id("lao")),
									j.Func().Params(j.Id("ctx").Qual("context", "Context")).Params(j.String(), j.Error()).BlockFunc(func(bl *j.Group) {
										origFn(bl, j.Lit(""))
										bl.Return(j.Id("id"), j.Nil())
									}),
								).Dot("Get").Call(j.Id("ctx"), j.Op("&").Id("opts").Dot(idFieldName)),
								j.Err().Op("!=").Nil(),
							).Block(
								j.Return(j.Id("opts"), j.Qual("fmt", "Errorf").Call(j.Lit("error evaluating id expression for %q workflow: %w"), j.Id(m.toCamel("%sWorkflowName", workflow)), j.Err())),
							)
						}

						// introduce local activity wrapper behind workflow versioning
						switch pvm := m.patchMode(temporalv1.Patch_PV_64, workflow); pvm {
						case temporalv1.Patch_PVM_ENABLED:
							patchComment(b, temporalv1.Patch_PV_64)
							b.If(j.Add(patchVersion(temporalv1.Patch_PV_64, pvm)).Op("==").Lit(1)).BlockFunc(
								fixFn,
							).Else().BlockFunc(func(b *j.Group) {
								origFn(b, j.Id("opts"))
								b.Id("opts").Dot(idFieldName).Op("=").Id("id")
							})
						case temporalv1.Patch_PVM_MARKER:
							b.Add(patchVersion(temporalv1.Patch_PV_64, pvm))
							fixFn(b)
						case temporalv1.Patch_PVM_REMOVED:
							fixFn(b)
						default:
							origFn(b, j.Id("opts"))
							b.Id("opts").Dot(idFieldName).Op("=").Id("id")
						}
					} else {
						origFn(b, j.Id("opts"))
						b.Id("opts").Dot(idFieldName).Op("=").Id("id")
					}
				})
			}

			// set WorkflowIDReusePolicy
			idReusePolicy := fn.If(j.Id("v").Op(":=").Id("o").Dot("idReusePolicy"), j.Id("v").Op("!=").Qual(enumsPkg, "WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED")).Block(
				j.Id("opts").Dot("WorkflowIDReusePolicy").Op("=").Id("v"),
			)
			if opts.GetIdReusePolicy() != temporalv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
				idReusePolicy.Else().If(j.Id("opts").Dot("WorkflowIDReusePolicy").Op("==").Qual(enumsPkg, "WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED")).BlockFunc(func(bl *j.Group) {
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
			fn.If(j.Id("v").Op(":=").Id("o").Dot("taskQueue"), j.Id("v").Op("!=").Nil()).Block(
				j.Id("opts").Dot("TaskQueue").Op("=").Op("*").Id("v"),
			).Else().If(j.Id("opts").Dot("TaskQueue").Op("==").Lit("")).BlockFunc(func(bl *j.Group) {
				var taskQueueVar j.Code
				if tq := opts.GetTaskQueue(); tq != "" {
					taskQueueVar = j.Lit(tq)
				} else if tq = m.opts.GetTaskQueue(); tq != "" {
					taskQueueVar = j.Id(m.toCamel("%sTaskQueue", m.Service.GoName))
				}
				if taskQueueVar != nil {
					bl.Id("opts").Dot("TaskQueue").Op("=").Add(taskQueueVar)
				} else if child {
					bl.Id("opts").Dot("TaskQueue").Op("=").Qual(workflowPkg, "GetInfo").Call(j.Id("ctx")).Dot("TaskQueueName")
				} else {
					bl.Return(j.Id("opts"), j.Qual("errors", "New").Call(j.Lit("TaskQueue is required")))
				}
			})

			// set RetryPolicy
			retryPolicy := fn.If(j.Id("v").Op(":=").Id("o").Dot("retryPolicy"), j.Id("v").Op("!=").Nil()).Block(
				j.Id("opts").Dot("RetryPolicy").Op("=").Id("v"),
			)
			if policy := opts.GetRetryPolicy(); policy != nil {
				retryPolicy.Else().If(j.Id("opts").Dot("RetryPolicy").Op("==").Nil()).Block(
					j.Id("opts").Dot("RetryPolicy").Op("=").Op("&").Qual(temporalPkg, "RetryPolicy").CustomFunc(multiLineValues, func(fields *j.Group) {
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
							fields.Id("NonRetryableErrorTypes").Op(":").Index().String().CustomFunc(multiLineValues, func(vals *j.Group) {
								for _, err := range errs {
									vals.Lit(err)
								}
							})
						}
					}),
				)
			}

			// set SearchAttributes
			searchAttributes := fn.If(j.Id("v").Op(":=").Id("o").Dot("searchAttributes"), j.Id("v").Op("!=").Nil()).Block(
				j.Id("opts").Dot("SearchAttributes").Op("=").Id("o").Dot("searchAttributes"),
			)
			if mapping := opts.GetSearchAttributes(); mapping != "" {
				searchAttributes.Else().If(j.Id("opts").Dot("SearchAttributes").Op("==").Nil()).
					BlockFunc(func(b *j.Group) {
						// original expression evaluation logic
						origFn := func(b *j.Group, errorResult j.Code) {
							// convert input to generic mapping input
							b.List(j.Id("structured"), j.Err()).Op(":=").Qual(expressionPkg, "ToStructured").Call(j.Id("req"))
							b.If(j.Err().Op("!=").Nil()).Block(
								j.Return(errorResult, j.Qual("fmt", "Errorf").Call(j.Lit(fmt.Sprintf("error serializing input for %q search attribute mapping: %%v", m.methods[workflow].GoName)), j.Err())),
							)

							// execute mapping
							b.List(j.Id("result"), j.Err()).Op(":=").Id(m.toCamel("%sSearchAttributesMapping", workflow)).Dot("Query").Call(j.Id("structured"))
							b.If(j.Err().Op("!=").Nil()).Block(
								j.Return(errorResult, j.Qual("fmt", "Errorf").Call(j.Lit(fmt.Sprintf("error executing %q search attribute mapping: %%v", m.methods[workflow].GoName)), j.Err())),
							)

							// coerce mapping result to map[string]any
							b.List(j.Id("searchAttributes"), j.Id("ok")).Op(":=").Id("result").Op(".").Parens(j.Map(j.String()).Any())
							b.If(j.Op("!").Id("ok")).Block(
								j.Return(errorResult, j.Qual("fmt", "Errorf").Call(j.Lit(fmt.Sprintf("expected %q search attribute mapping to return map[string]any, got: %%T", m.methods[workflow].GoName)), j.Id("result"))),
							)
						}

						if child {
							// local activity wrapper
							fixFn := func(b *j.Group) {
								b.Id("lao").Op(":=").Qual(workflowPkg, "GetLocalActivityOptions").Call(j.Id("ctx"))
								b.Id("lao").Dot("ScheduleToCloseTimeout").Op("=").Qual("time", "Second").Op("*").Lit(10)
								b.If(
									j.Err().Op(":=").Qual(workflowPkg, "ExecuteLocalActivity").Call(
										j.Qual(workflowPkg, "WithLocalActivityOptions").Call(j.Id("ctx"), j.Id("lao")),
										j.Func().Params(j.Id("ctx").Qual("context", "Context")).Params(j.Map(j.String()).Any(), j.Error()).BlockFunc(func(b *j.Group) {
											origFn(b, j.Nil())
											b.Return(j.Id("searchAttributes"), j.Nil())
										}),
									).Dot("Get").Call(j.Id("ctx"), j.Op("&").Id("opts").Dot("SearchAttributes")),
									j.Err().Op("!=").Nil(),
								).Block(
									j.Return(j.Id("opts"), j.Qual("fmt", "Errorf").Call(j.Lit("error evaluating search attributes for %q workflow: %w"), j.Id(m.toCamel("%sWorkflowName", workflow)), j.Err())),
								)
							}

							// introduce local activity wrapper behind workflow versioning
							switch pvm := m.patchMode(temporalv1.Patch_PV_64, workflow); pvm {
							case temporalv1.Patch_PVM_ENABLED:
								patchComment(b, temporalv1.Patch_PV_64)
								b.If(j.Add(patchVersion(temporalv1.Patch_PV_64, pvm)).Op("==").Lit(1)).BlockFunc(
									fixFn,
								).Else().BlockFunc(func(b *j.Group) {
									origFn(b, j.Id("opts"))
									b.Id("opts").Dot("SearchAttributes").Op("=").Id("searchAttributes")
								})
							case temporalv1.Patch_PVM_MARKER:
								b.Add(patchVersion(temporalv1.Patch_PV_64, pvm))
								fixFn(b)
							case temporalv1.Patch_PVM_REMOVED:
								fixFn(b)
							default:
								origFn(b, j.Id("opts"))
								b.Id("opts").Dot("SearchAttributes").Op("=").Id("searchAttributes")
							}
						} else {
							origFn(b, j.Id("opts"))
							b.Id("opts").Dot("SearchAttributes").Op("=").Id("searchAttributes")
						}
					})
			}

			// set WorkflowExecutionTimeout
			executionTimeout := fn.If(j.Id("v").Op(":=").Id("o").Dot("executionTimeout"), j.Id("v").Op("!=").Nil()).Block(
				j.Id("opts").Dot("WorkflowExecutionTimeout").Op("=").Op("*").Id("v"),
			)
			if d := opts.GetExecutionTimeout().AsDuration(); d > 0 {
				executionTimeout.Else().If(j.Id("opts").Dot("WorkflowExecutionTimeout").Op("==").Lit(0)).Block(
					j.Id("opts").Dot("WorkflowExecutionTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
				)
			}

			// set WorkflowRunTimeout
			runTimeout := fn.If(j.Id("v").Op(":=").Id("o").Dot("runTimeout"), j.Id("v").Op("!=").Nil()).Block(
				j.Id("opts").Dot("WorkflowRunTimeout").Op("=").Op("*").Id("v"),
			)
			if d := opts.GetRunTimeout().AsDuration(); d > 0 {
				runTimeout.Else().If(j.Id("opts").Dot("WorkflowRunTimeout").Op("==").Lit(0)).Block(
					j.Id("opts").Dot("WorkflowRunTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
				)
			}

			// set WorkflowTaskTimeout
			taskTimeout := fn.If(j.Id("v").Op(":=").Id("o").Dot("taskTimeout"), j.Id("v").Op("!=").Nil()).Block(
				j.Id("opts").Dot("WorkflowTaskTimeout").Op("=").Op("*").Id("v"),
			)
			if d := opts.GetTaskTimeout().AsDuration(); d > 0 {
				taskTimeout.Else().If(j.Id("opts").Dot("WorkflowTaskTimeout").Op("==").Lit(0)).Block(
					j.Id("opts").Dot("WorkflowTaskTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
				)
			}

			// set ParentClosePolicy
			if child {
				parentClosePolicy := fn.If(j.Id("v").Op(":=").Id("o").Dot("parentClosePolicy"), j.Id("v").Op("!=").Qual(enumsPkg, "PARENT_CLOSE_POLICY_UNSPECIFIED")).Block(
					j.Id("opts").Dot("ParentClosePolicy").Op("=").Id("v"),
				)
				if opts.GetParentClosePolicy() != temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_UNSPECIFIED {
					parentClosePolicy.Else().If(j.Id("opts").Dot("ParentClosePolicy").Op("==").Qual(enumsPkg, "PARENT_CLOSE_POLICY_UNSPECIFIED")).BlockFunc(func(bl *j.Group) {
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
				waitForCancellation := fn.If(j.Id("v").Op(":=").Id("o").Dot("waitForCancellation"), j.Id("v").Op("!=").Nil()).Block(
					j.Id("opts").Dot("WaitForCancellation").Op("=").Op("*").Id("v"),
				)
				if opts.GetWaitForCancellation() {
					waitForCancellation.Else().Block(
						j.Id("opts").Dot("WaitForCancellation").Op("=").Lit(true),
					)
				}
			}

			if p := opts.GetWorkflowIdConflictPolicy(); p != enums.WORKFLOW_ID_CONFLICT_POLICY_UNSPECIFIED {
				fn.If(j.Id("opts").Dot("workflowIdConflictPolicy").Op("!=").Qual(enumsPkg, "WORKFLOW_ID_CONFLICT_POLICY_UNSPECIFIED")).BlockFunc(func(bl *j.Group) {
					var policy string
					switch p {
					case enums.WORKFLOW_ID_CONFLICT_POLICY_FAIL:
						policy = "WORKFLOW_ID_CONFLICT_POLICY_FAIL"
					case enums.WORKFLOW_ID_CONFLICT_POLICY_TERMINATE_EXISTING:
						policy = "WORKFLOW_ID_CONFLICT_POLICY_TERMINATE_EXISTING"
					case enums.WORKFLOW_ID_CONFLICT_POLICY_USE_EXISTING:
						policy = "WORKFLOW_ID_CONFLICT_POLICY_USE_EXISTING"
					}
					bl.Id("opts").Dot("WorkflowIdConflictPolicy").Op("=").Qual(enumsPkg, policy)
				})
			}

			fn.Return(j.Id("opts"), j.Nil())
		})

	baseName := m.toCamel("With%s", optionsType)
	f.Commentf("%s sets the initial %s.%s", baseName, optionsPkg, optionsType)
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id(baseName).
		Params(j.Id("options").Qual(optionsPkg, optionsType)).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("options").Op("=").Id("options"),
			j.Return(j.Id("o")),
		)

	f.Comment("WithExecutionTimeout sets the WorkflowExecutionTimeout value")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithExecutionTimeout").
		Params(j.Id("d").Qual("time", "Duration")).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("executionTimeout").Op("=").Op("&").Id("d"),
			j.Return(j.Id("o")),
		)

	if child {
		f.Comment("WithID sets the WorkflowID value")
	} else {
		f.Comment("WithID sets the ID value")
	}
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithID").
		Params(j.Id("id").String()).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("id").Op("=").Op("&").Id("id"),
			j.Return(j.Id("o")),
		)

	f.Comment("WithIDReusePolicy sets the WorkflowIDReusePolicy value")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithIDReusePolicy").
		Params(j.Id("policy").Qual(enumsPkg, "WorkflowIdReusePolicy")).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("idReusePolicy").Op("=").Id("policy"),
			j.Return(j.Id("o")),
		)

	if child {
		f.Comment("WithParentClosePolicy sets the WorkflowIDReusePolicy value")
		f.Func().
			Params(j.Id("o").Op("*").Id(typeName)).
			Id("WithParentClosePolicy").
			Params(j.Id("policy").Qual(enumsPkg, "ParentClosePolicy")).
			Op("*").Id(typeName).
			Block(
				j.Id("o").Dot("parentClosePolicy").Op("=").Id("policy"),
				j.Return(j.Id("o")),
			)
	}

	f.Comment("WithRetryPolicy sets the RetryPolicy value")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithRetryPolicy").
		Params(j.Id("policy").Op("*").Qual(temporalPkg, "RetryPolicy")).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("retryPolicy").Op("=").Id("policy"),
			j.Return(j.Id("o")),
		)

	f.Comment("WithRunTimeout sets the WorkflowRunTimeout value")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithRunTimeout").
		Params(j.Id("d").Qual("time", "Duration")).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("runTimeout").Op("=").Op("&").Id("d"),
			j.Return(j.Id("o")),
		)

	f.Comment("WithSearchAttributes sets the SearchAttributes value")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithSearchAttributes").
		Params(j.Id("sa").Map(j.String()).Any()).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("searchAttributes").Op("=").Id("sa"),
			j.Return(j.Id("o")),
		)

	f.Comment("WithTaskTimeout sets the WorkflowTaskTimeout value")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithTaskTimeout").
		Params(j.Id("d").Qual("time", "Duration")).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("taskTimeout").Op("=").Op("&").Id("d"),
			j.Return(j.Id("o")),
		)

	f.Comment("WithTaskQueue sets the TaskQueue value")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithTaskQueue").
		Params(j.Id("tq").String()).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("taskQueue").Op("=").Op("&").Id("tq"),
			j.Return(j.Id("o")),
		)

	if child {
		f.Comment("WithWaitForCancellation sets the WaitForCancellation value")
		f.Func().
			Params(j.Id("o").Op("*").Id(typeName)).
			Id("WithWaitForCancellation").
			Params(j.Id("wait").Bool()).
			Op("*").Id(typeName).
			Block(
				j.Id("o").Dot("waitForCancellation").Op("=").Op("&").Id("wait"),
				j.Return(j.Id("o")),
			)
	}

	f.Comment("WithWorkflowIdConflictPolicy sets the WorkflowIdConflictPolicy value")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithWorkflowIdConflictPolicy").
		Params(j.Id("policy").Qual(enumsPkg, "WorkflowIdConflictPolicy")).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("workflowIdConflictPolicy").Op("=").Id("policy"),
			j.Return(j.Id("o")),
		)
}
