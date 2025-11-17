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

func (n *names) clientSignalWithStart(workflow, signal protoreflect.FullName) string {
	return n.toCamel("%sWith%s", workflow, signal)
}

func (n *names) clientSignalWithStartAsync(workflow, signal protoreflect.FullName) string {
	return n.toCamel("%sWith%sAsync", workflow, signal)
}

func (n *names) clientUpdate(update protoreflect.FullName) string {
	return n.toCamel("%s", update)
}

func (n *names) clientUpdateAsync(update protoreflect.FullName) string {
	return n.toCamel("%sAsync", update)
}

func (n *names) clientUpdateGet(update protoreflect.FullName) string {
	return n.toCamel("Get%s", update)
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
		StructFunc(func(g *j.Group) {
			g.Id("client").Qual(clientPkg, "Client")
			g.Id("log").Op("*").Qual("log/slog", "Logger")
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
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			g.Id("workflowID").String()
			g.Id("runID").String()
			if hasInput {
				g.Id("query").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
		}).
		Params(
			j.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output)),
			j.Error(),
		).
		BlockFunc(func(g *j.Group) {
			if isDeprecated(method) {
				g.Id("c").Dot("log").Dot("WarnContext").Call(j.Id("ctx"), j.Lit("use of deprecated client method detected"), j.Lit("method"), j.Lit(m.toCamel("%s", query)), j.Lit("query"), j.Id(m.toCamel("%sQueryName", query))).Line()
			}
			g.Var().Id("resp").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			g.If(
				j.List(j.Id("val"), j.Err()).Op(":=").Id("c").Dot("client").Dot("QueryWorkflow").CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("workflowID")
					g.Id("runID")
					g.Id(m.toCamel("%sQueryName", query))
					if hasInput {
						g.Id("query")
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
			g.Return(
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
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			g.Id("workflowID").String()
			g.Id("runID").String()
			if hasInput {
				g.Id("signal").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
		}).
		Params(j.Error()).
		BlockFunc(func(g *j.Group) {
			if isDeprecated(method) {
				g.Id("c").Dot("log").Dot("WarnContext").Call(j.Id("ctx"), j.Lit("use of deprecated client method detected"), j.Lit("method"), j.Lit(m.toCamel("%s", signal)), j.Lit("signal"), j.Id(m.toCamel("%sSignalName", signal)))
				g.Line()
			}
			g.Return(
				j.Id("c").Dot("client").Dot("SignalWorkflow").CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("workflowID")
					g.Id("runID")
					g.Id(m.toCamel("%sSignalName", signal))
					if hasInput {
						g.Id("signal")
					} else {
						g.Nil()
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
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasWorkflowInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if hasSignalInput {
				g.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			g.Id("options").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
		}).
		ParamsFunc(func(g *j.Group) {
			if hasWorkflowOutput {
				g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			if workflowDeprecated, signalDeprecated := isDeprecated(method), isDeprecated(handler); workflowDeprecated || signalDeprecated {
				if workflowDeprecated {
					g.Id("c").Dot("log").Dot("WarnContext").Call(j.Id("ctx"), j.Lit("use of deprecated client method detected"), j.Lit("method"), j.Lit(name), j.Lit("workflow"), j.Id(m.toCamel("%sWorkflowName", workflow)))
				}
				if signalDeprecated {
					g.Id("c").Dot("log").Dot("WarnContext").Call(j.Id("ctx"), j.Lit("use of deprecated client method detected"), j.Lit("method"), j.Lit(name), j.Lit("signal"), j.Id(m.toCamel("%sSignalName", signal)))
				}
				g.Line()
			}
			// signal with start workflow
			g.Id("run").Op(",").Err().Op(":=").Id("c").Dot(m.toCamel("%sWith%sAsync", workflow, signal)).CallFunc(func(g *j.Group) {
				g.Id("ctx")
				if hasWorkflowInput {
					g.Id("req")
				}
				if hasSignalInput {
					g.Id("signal")
				}
				g.Id("options").Op("...")
			})
			g.If(j.Err().Op("!=").Nil()).Block(
				j.ReturnFunc(func(g *j.Group) {
					if hasWorkflowOutput {
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
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasWorkflowInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			if hasSignalInput {
				g.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			g.Id("options").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
		}).
		Params(
			j.Id(m.toCamel("%sRun", workflow)),
			j.Error(),
		).
		BlockFunc(func(g *j.Group) {
			if workflowDeprecated, signalDeprecated := isDeprecated(method), isDeprecated(handler); workflowDeprecated || signalDeprecated {
				if workflowDeprecated {
					g.Id("c").Dot("log").Dot("WarnContext").Call(j.Id("ctx"), j.Lit("use of deprecated client method detected"), j.Lit("method"), j.Lit(name), j.Lit("workflow"), j.Id(m.toCamel("%sWorkflowName", workflow)))
				}
				if signalDeprecated {
					g.Id("c").Dot("log").Dot("WarnContext").Call(j.Id("ctx"), j.Lit("use of deprecated client method detected"), j.Lit("method"), j.Lit(name), j.Lit("signal"), j.Id(m.toCamel("%sSignalName", signal)))
				}
				g.Line()
			}

			// initialize options
			g.Var().Id("o").Op("*").Id(m.toCamel("%sOptions", workflow))
			g.If(j.Len(j.Id("options")).Op(">").Lit(0).Op("&&").Id("options").Index(j.Lit(0)).Op("!=").Nil()).Block(
				j.Id("o").Op("=").Id("options").Index(j.Lit(0)),
			).Else().Block(
				j.Id("o").Op("=").Id(m.toCamel("New%sOptions", workflow)).Call(),
			)

			// initialize client.StartWorkfowOptions
			g.List(j.Id("opts"), j.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(g *j.Group) {
				if hasWorkflowInput {
					g.Id("req").Dot("ProtoReflect").Call()
				} else {
					g.Nil()
				}
			})
			g.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing client.StartWorkflowOptions: %w"), j.Err())),
			)

			// signal with start workflow
			g.Id("run").Op(",").Err().Op(":=").Id("c").Dot("client").Dot("SignalWithStartWorkflow").CallFunc(func(g *j.Group) {
				g.Id("ctx")
				g.Id("opts").Dot("ID")
				g.Qual(m.goImportPathForMethod(signal), m.toCamel("%sSignalName", signal))
				if hasSignalInput {
					g.Id("signal")
				} else {
					g.Nil()
				}
				g.Id("opts")
				g.Id(m.toCamel("%sWorkflowName", workflow))
				if hasWorkflowInput {
					g.Id("req")
				}
			})
			g.If(j.Id("run").Op("==").Nil().Op("||").Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Err()),
			)
			g.Return(
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
				j.Id("c").Dot("client").Dot(methodName).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("workflowID")
					g.Id("runID")
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
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasInput {
				g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
			}
			g.Id("options").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
		}).
		Params(
			j.Id(runInterfaceType),
			j.Error(),
		).
		BlockFunc(func(g *j.Group) {
			if deprecated {
				g.Id("c").Dot("log").Dot("WarnContext").Call(j.Id("ctx"), j.Lit("use of deprecated client method detected"), j.Lit("method"), j.Lit(methodName), j.Lit("workflow"), j.Id(m.toCamel("%sWorkflowName", workflow))).Line()
			}

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

			// execute workflow
			g.Id("run").Op(",").Err().Op(":=").Id("c").Dot("client").Dot("ExecuteWorkflow").CallFunc(func(g *j.Group) {
				g.Id("ctx")
				g.Id("opts")
				g.Id(m.toCamel("%sWorkflowName", workflow))
				if hasInput {
					g.Id("req")
				}
			})
			g.If(j.Err().Op("!=").Nil()).Block(
				j.Return(j.Nil(), j.Err()),
			)
			g.If(j.Id("run").Op("==").Nil()).Block(
				j.Return(j.Nil(), j.Qual("errors", "New").Call(j.Lit("execute workflow returned nil run"))),
			)
			g.Return(
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
				j.Id("c").Dot("client").Dot(methodName).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("workflowID")
					g.Id("runID")
					g.Id("reason")
					g.Id("details").Op("...")
				}),
			),
		)
}

// genClientInterface generates a Client interface for a given service
func (m *Manifest) genClientInterface(f *j.File) {
	typeName := m.toCamel("%sClient", m.Service.GoName)

	f.Commentf("%s describes a client for a(n) %s worker", typeName, m.Service.Desc.FullName())
	f.Type().Id(typeName).InterfaceFunc(func(g *j.Group) {
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
			methodName := m.Names().clientWorkflow(workflow)
			commentWithDefaultf(g, methodSet(method), "%s executes a(n) %s workflow and blocks until error or response received", methodName, m.fqnForWorkflow(workflow))
			g.Id(methodName).
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
				Line()

			// generate <Workflow>Async method
			methodName = m.Names().clientWorkflowAsync(workflow)
			commentf(g, methodSet(method), "%s starts a(n) %s workflow and returns a handle to the workflow run", methodName, m.fqnForWorkflow(workflow))
			g.Id(methodName).
				ParamsFunc(func(g *j.Group) {
					g.Id("ctx").Qual("context", "Context")
					if hasInput {
						g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
					}
					g.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
				}).
				Params(
					j.Id(runInterfaceType),
					j.Error(),
				).
				Line()

			// generate Get<Workflow> method
			methodName = m.Names().clientWorkflowGet(workflow)
			commentf(g, methodSet(method), "%s retrieves a handle to an existing %s workflow execution", methodName, m.fqnForWorkflow(workflow))
			g.Id(m.toCamel("Get%s", workflow)).
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
				methodName := m.Names().clientSignalWithStart(workflow, signal)
				commentf(g, methodSet(method, handler), "%s sends a(n) %s signal to a(n) %s workflow, starting it if necessary, and blocks until workflow completion", methodName, m.fqnForSignal(signal), m.fqnForWorkflow(workflow))
				g.Id(methodName).
					ParamsFunc(func(g *j.Group) {
						g.Id("ctx").Qual("context", "Context")
						if hasWorkflowInput {
							g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
						}
						if hasSignalInput {
							g.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
						}
						g.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
					}).
					ParamsFunc(func(g *j.Group) {
						if hasWorkflowOutput {
							g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
						}
						g.Error()
					}).
					Line()

				// add async flavor
				methodName = m.Names().clientSignalWithStartAsync(workflow, signal)
				commentf(g, methodSet(method, handler), "%s sends a(n) %s signal to a(n) %s workflow, starting it if necessary, and returns a handle to the workflow execution", methodName, m.fqnForSignal(signal), m.fqnForWorkflow(workflow))
				g.Id(methodName).
					ParamsFunc(func(g *j.Group) {
						g.Id("ctx").Qual("context", "Context")
						if hasWorkflowInput {
							g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
						}
						if hasSignalInput {
							g.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
						}
						g.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", workflow))
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
				commentf(g, methodSet(method, handler), "%s executes a(n) %s update on a(n) %s workflow, starting it if necessary, and blocks until update completion", methodName, m.fqnForUpdate(update), m.fqnForWorkflow(workflow))
				g.Id(methodName).
					ParamsFunc(func(g *j.Group) {
						g.Id("ctx").Qual("context", "Context")
						if hasWorkflowInput {
							g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
						}
						if hasUpdateInput {
							g.Id("update").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
						}
						g.Id("opts").Op("...").Op("*").Id(optionsName)
					}).
					ParamsFunc(func(g *j.Group) {
						if hasUpdateOutput {
							g.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
						}
						g.Id(runName)
						g.Error()
					})

				asyncName := m.Names().clientUpdateWithStartAsync(workflow, update)
				commentf(g, methodSet(method, handler), "%s starts a(n) %s update on a(n) %s workflow, starting it if necessary, and returns a handle to the update execution", asyncName, m.fqnForUpdate(update), m.fqnForWorkflow(workflow))
				g.Id(asyncName).
					ParamsFunc(func(g *j.Group) {
						g.Id("ctx").Qual("context", "Context")
						if hasWorkflowInput {
							g.Id("req").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), m.getMessageName(method.Input))
						}
						if hasUpdateInput {
							g.Id("update").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
						}
						g.Id("opts").Op("...").Op("*").Id(optionsName)
					}).
					Params(j.Id(handleName), j.Id(runName), j.Error())
			}
		}

		// generate CancelWorkflow method
		methodName := "CancelWorkflow"
		g.Commentf("%s requests cancellation of an existing workflow execution", methodName)
		g.Id(methodName).
			Params(
				j.Id("ctx").Qual("context", "Context"),
				j.Id("workflowID").String(),
				j.Id("runID").String(),
			).
			Params(j.Error()).
			Line()

		// generate TerminateWorkflow method
		methodName = "TerminateWorkflow"
		g.Commentf("%s an existing workflow execution", methodName)
		g.Id(methodName).
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
			commentWithDefaultf(g, methodSet(handler), "%s executes a(n) %s query", query, m.fqnForQuery(query))
			g.Id(m.toCamel("%s", query)).
				ParamsFunc(func(g *j.Group) {
					g.Id("ctx").Qual("context", "Context")
					g.Id("workflowID").String()
					g.Id("runID").String()
					if hasInput {
						g.Id("query").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
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
			commentWithDefaultf(g, methodSet(handler), "%s sends a(n) %s signal", signal, m.fqnForSignal(signal))
			g.Id(m.toCamel("%s", signal)).
				ParamsFunc(func(g *j.Group) {
					g.Id("ctx").Qual("context", "Context")
					g.Id("workflowID").String()
					g.Id("runID").String()
					if hasInput {
						g.Id("signal").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
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
			methodName := m.Names().clientUpdate(update)
			commentWithDefaultf(g, methodSet(handler), "%s executes a(n) %s update and blocks until update completion", methodName, m.fqnForUpdate(update))
			g.Id(methodName).
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
				Line()

			// add async flavor
			methodName = m.Names().clientUpdateAsync(update)
			commentf(g, methodSet(handler), "%s starts a(n) %s update and returns a handle to the workflow update", methodName, m.fqnForUpdate(update))
			g.Id(methodName).
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
				Line()

			// add getter
			methodName = m.Names().clientUpdateGet(update)
			commentf(g, methodSet(handler), "%s retrieves a handle to an existing %s update", methodName, m.fqnForUpdate(update))
			g.Id(methodName).
				ParamsFunc(func(g *j.Group) {
					g.Id("ctx").Qual("context", "Context")
					g.Id("req").Qual(clientPkg, "GetWorkflowUpdateHandleOptions")
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
			g.Var().Err().Error()
			g.Id("doneCh").Op(":=").Make(j.Chan().Struct())
			g.List(j.Id("gctx"), j.Id("cancel")).Op(":=").Qual("context", "WithCancel").Call(j.Qual("context", "Background").Call())
			g.Defer().Id("cancel").Call()
			g.Line()

			g.Go().Func().Params().Block(
				j.For().Block(
					j.Var().Id("deadlineExceeded").Op("*").Qual(serviceerrorPkg, "DeadlineExceeded"),
					j.If(
						j.Err().Op("=").Id("h").Dot("handle").Dot("Get").CallFunc(func(g *j.Group) {
							g.Id("gctx")
							if hasOutput {
								g.Op("&").Id("resp")
							} else {
								g.Nil()
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
			g.Line()

			g.Select().Block(
				j.Case(j.Op("<-").Id("ctx").Dot("Done").Call()).Block(
					j.ReturnFunc(func(g *j.Group) {
						if hasOutput {
							g.Nil()
						}
						g.Id("ctx").Dot("Err").Call()
					}),
				),
				j.Case(j.Op("<-").Id("doneCh")).BlockFunc(func(g *j.Group) {
					if hasOutput {
						g.If(j.Err().Op("!=").Nil()).Block(
							j.Return(j.Nil(), j.Err()),
						)
						g.Return(j.Op("&").Id("resp"), j.Nil())
					} else {
						g.Return(j.Err())
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
	f.Type().Id(typeName).InterfaceFunc(func(g *j.Group) {
		g.Comment("WorkflowID returns the workflow ID")
		g.Id("WorkflowID").Params().String()

		g.Comment("RunID returns the workflow instance ID")
		g.Id("RunID").Params().String()

		g.Comment("UpdateID returns the update ID")
		g.Id("UpdateID").Params().String()

		g.Comment("Get blocks until the workflow is complete and returns the result")
		g.Id("Get").
			Params(j.Id("ctx").Qual("context", "Context")).
			ParamsFunc(func(g *j.Group) {
				if hasOutput {
					g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
				}
				g.Error()
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
		ParamsFunc(func(g *j.Group) {
			g.Id("workflowID").String()
			g.Id("runID").String()
			if hasInput {
				g.Id("req").Op("*").Qual(string(m.methods[update].Input.GoIdent.GoImportPath), m.getMessageName(m.methods[update].Input))
			}
		}).
		Params(
			j.Id("opts").Op("*").Qual(clientPkg, "UpdateWorkflowOptions"),
			j.Err().Error(),
		).
		BlockFunc(func(g *j.Group) {
			g.Comment("use user-provided UpdateWorkflowOptions if exists")
			g.If(j.Id("o").Dot("Options").Op("!=").Nil()).Block(
				j.Id("opts").Op("=").Id("o").Dot("Options"),
			).Else().Block(
				j.Id("opts").Op("=").Op("&").Qual(clientPkg, "UpdateWorkflowOptions").Values(),
			)

			g.Line()
			g.Comment("set constants")
			if hasInput {
				g.Id("opts").Dot("Args").Op("=").Index().Any().Values(j.Id("req"))
			}
			g.Id("opts").Dot("RunID").Op("=").Id("runID")
			g.Id("opts").Dot("UpdateName").Op("=").Id(m.toCamel("%sUpdateName", update))
			g.Id("opts").Dot("WorkflowID").Op("=").Id("workflowID")

			g.Line()
			g.Comment("set UpdateID")
			id := g.If(j.Id("v").Op(":=").Id("o").Dot("id"), j.Id("v").Op("!=").Nil()).Block(
				j.Id("opts").Dot("UpdateID").Op("=").Op("*").Id("v"),
			)
			if idExpr := updateOpts.GetId(); idExpr != "" {
				id.Else().If(j.Id("opts").Dot("UpdateID").Op("==").Lit("")).BlockFunc(func(g *j.Group) {
					g.List(j.Id("id"), j.Err()).Op(":=").Qual(expressionPkg, "EvalExpression").CallFunc(func(g *j.Group) {
						g.Id(m.toCamel("%sIDExpression", update))
						if hasInput {
							g.Id("req").Dot("ProtoReflect").Call()
						} else {
							g.Nil()
						}
					})
					g.If(j.Err().Op("!=").Nil()).Block(
						j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error evaluating id expression for %q update: %w"), j.Id(m.toCamel("%sUpdateName", update)), j.Err())),
					)
					g.Id("opts").Dot("UpdateID").Op("=").Id("id")
				})
			}

			g.Line()
			g.Comment("set WaitPolicy")
			waitPolicy := g.If(j.Id("v").Op(":=").Id("o").Dot("waitPolicy"), j.Id("v").Op("!=").Qual(clientPkg, "WorkflowUpdateStageUnspecified")).Block(
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
			case temporalv1.WaitPolicy_WAIT_POLICY_COMPLETED:
				stage = "WorkflowUpdateStageCompleted"
			default:
				stage = "WorkflowUpdateStageAccepted"
			}
			waitPolicy.Else().If(j.Id("opts").Dot("WaitForStage").Op("==").Qual(clientPkg, "WorkflowUpdateStageUnspecified")).Block(
				j.Id("opts").Dot("WaitForStage").Op("=").Qual(clientPkg, stage),
			)

			g.Return(j.Id("opts"), j.Nil())
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
				j.Id("r").Dot("client").Dot("CancelWorkflow").CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("r").Dot("ID").Call()
					g.Id("r").Dot("RunID").Call()
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
		ParamsFunc(func(g *j.Group) {
			if hasOutput {
				g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
			}
			g.Error()
		}).
		BlockFunc(func(g *j.Group) {
			if hasOutput {
				g.Var().Id("resp").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
				g.If(
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
				g.Return(
					j.Op("&").Id("resp"), j.Nil(),
				)
			} else {
				g.Return(
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
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasInput {
				g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
		}).
		Params(
			j.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output)),
			j.Error(),
		).
		Block(
			j.Return(
				j.Id("r").Dot("client").Dot(m.methods[query].GoName).CallFunc(func(g *j.Group) {
					g.Id("ctx")
					g.Id("r").Dot("ID").Call()
					g.Lit("")
					if hasInput {
						g.Id("req")
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
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasInput {
				g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
		}).
		Params(j.Error()).
		Block(
			j.ReturnFunc(func(g *j.Group) {
				if m.methodsFromSameService(signal, workflow) {
					g.Id("r").Dot("client").Dot(m.methods[signal].GoName).CallFunc(func(g *j.Group) {
						g.Id("ctx")
						g.Id("r").Dot("ID").Call()
						g.Lit("")
						if hasInput {
							g.Id("req")
						}
					})
				} else {
					g.Id("r").Dot("client").Dot("client").Dot("SignalWorkflow").CallFunc(func(g *j.Group) {
						g.Id("ctx")
						g.Id("r").Dot("ID").Call()
						g.Id("r").Dot("RunID").Call()
						g.Add(m.Qual(signal, m.toCamel("%sSignalName", signal)))
						if hasInput {
							g.Id("req")
						} else {
							g.Nil()
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
		ParamsFunc(func(g *j.Group) {
			g.Id("ctx").Qual("context", "Context")
			if hasInput {
				g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
			}
			g.Id("opts").Op("...").Op("*").Id(m.toCamel("%sOptions", update))
		}).
		Params(
			j.Add(m.Qual(update, m.toCamel("%sHandle", update))),
			j.Error(),
		).
		BlockFunc(func(g *j.Group) {
			if m.methodsFromSamePackage(workflow, update) {
				g.Return(
					j.Id("r").Dot("client").Dot(m.toCamel("%sAsync", update)).CallFunc(func(g *j.Group) {
						g.Id("ctx")
						g.Id("r").Dot("ID").Call()
						g.Id("r").Dot("RunID").Call()
						if hasInput {
							g.Id("req")
						}
						g.Id("opts").Op("...")
					}),
				)
			} else {
				// initialize options
				g.Var().Id("o").Op("*").Id(m.toCamel("%sOptions", update))
				g.If(j.Len(j.Id("opts")).Op(">").Lit(0).Op("&&").Id("opts").Index(j.Lit(0)).Op("!=").Nil()).Block(
					j.Id("o").Op("=").Id("opts").Index(j.Lit(0)),
				).Else().Block(
					j.Id("o").Op("=").Id(m.toCamel("New%sOptions", update)).Call(),
				)

				// build UpdateWorkflowWithOptions
				g.List(j.Id("options"), j.Err()).Op(":=").Id("o").Dot("Build").CallFunc(func(g *j.Group) {
					g.Id("r").Dot("ID").Call()
					g.Id("r").Dot("RunID").Call()
					if hasInput {
						g.Id("req")
					}
				})
				g.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Nil(), j.Qual("fmt", "Errorf").Call(j.Lit("error initializing UpdateWorkflowWithOptions: %w"), j.Err())),
				)

				// update workflow
				g.List(j.Id("handle"), j.Err()).Op(":=").Id("c").Dot("client").Dot("UpdateWorkflowWithOptions").Call(j.Id("ctx"), j.Id("options"))
				g.If(j.Err().Op("!=").Nil()).Block(
					j.Return(j.Nil(), j.Err()),
				)

				g.Return(
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

// genClientWorkflowRunInterface generates a <Workflow>Run interface
func (m *Manifest) genClientWorkflowRunInterface(f *j.File, workflow protoreflect.FullName) {
	typeName := m.toCamel("%sRun", workflow)
	opts := m.workflows[workflow]
	method := m.methods[workflow]
	hasOutput := !isEmpty(method.Output)

	f.Commentf("%s describes a(n) %s workflow run", typeName, m.fqnForWorkflow(workflow))
	f.Type().Id(typeName).InterfaceFunc(func(g *j.Group) {
		g.Comment("ID returns the workflow ID")
		g.Id("ID").Params().String().Line()

		g.Comment("RunID returns the workflow instance ID")
		g.Id("RunID").Params().String().Line()

		g.Comment("Run returns the inner client.WorkflowRun")
		g.Id("Run").Params().Qual(clientPkg, "WorkflowRun").Line()

		g.Comment("Get blocks until the workflow is complete and returns the result")
		g.Id("Get").
			Params(j.Id("ctx").Qual("context", "Context")).
			ParamsFunc(func(g *j.Group) {
				if hasOutput {
					g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), m.getMessageName(method.Output))
				}
				g.Error()
			}).Line()

		g.Comment("Cancel requests cancellation of a workflow in execution, returning an error if applicable")
		g.Id("Cancel").
			Params(j.Id("ctx").Qual("context", "Context")).
			Params(j.Error()).Line()

		g.Comment("Terminate terminates a workflow in execution, returning an error if applicable")
		g.Id("Terminate").
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

			commentWithDefaultf(g, methodSet(handler), "%s executes a(n) %s query", query, m.fqnForQuery(query))
			g.Id(m.toCamel("%s", query)).
				ParamsFunc(func(g *j.Group) {
					g.Id("ctx").Qual("context", "Context")
					if hasInput {
						g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
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

			commentWithDefaultf(g, methodSet(handler), "%s sends a(n) %s signal", signal, m.fqnForSignal(signal))
			g.Id(m.toCamel("%s", signal)).
				ParamsFunc(func(g *j.Group) {
					g.Id("ctx").Qual("context", "Context")
					if hasInput {
						g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
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
			commentWithDefaultf(g, methodSet(handler), "%s executes a(n) %s update", update, m.fqnForUpdate(update))
			g.Id(m.toCamel("%s", update)).
				ParamsFunc(func(g *j.Group) {
					g.Id("ctx").Qual("context", "Context")
					if hasInput {
						g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
					g.Id("opts").Op("...").Op("*").Add(m.Qual(update, m.toCamel("%sOptions", update)))
				}).
				ParamsFunc(func(g *j.Group) {
					if hasOutput {
						g.Op("*").Qual(string(handler.Output.GoIdent.GoImportPath), m.getMessageName(handler.Output))
					}
					g.Error()
				}).Line()

			// add async flavor
			commentWithDefaultf(g, methodSet(handler), "%sAsync sends a(n) %s update to the workflow", update, m.fqnForUpdate(update))
			g.Id(m.toCamel("%sAsync", update)).
				ParamsFunc(func(g *j.Group) {
					g.Id("ctx").Qual("context", "Context")
					if hasInput {
						g.Id("req").Op("*").Qual(string(handler.Input.GoIdent.GoImportPath), m.getMessageName(handler.Input))
					}
					g.Id("opts").Op("...").Op("*").Add(m.Qual(update, m.toCamel("%sOptions", update)))
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
	hasInput := !isEmpty(m.methods[workflow].Input)

	f.Commentf("%s provides configuration for a %s%s workflow operation", typeName, childQualifier, m.fqnForWorkflow(workflow))
	f.Type().Id(typeName).StructFunc(func(g *j.Group) {
		g.Id("options").Qual(optionsPkg, optionsType)
		g.Id("executionTimeout").Op("*").Qual("time", "Duration")
		g.Id("id").Op("*").String()
		g.Id("idReusePolicy").Qual(enumsPkg, "WorkflowIdReusePolicy")
		g.Id("retryPolicy").Op("*").Qual(temporalPkg, "RetryPolicy")
		g.Id("runTimeout").Op("*").Qual("time", "Duration")
		g.Id("searchAttributes").Map(j.String()).Any()
		g.Id("taskQueue").Op("*").String()
		g.Id("taskTimeout").Op("*").Qual("time", "Duration")
		g.Id("typedSearchAttributes").Op("*").Qual(temporalPkg, "SearchAttributes")
		g.Id("workflowIdConflictPolicy").Qual(enumsPkg, "WorkflowIdConflictPolicy")
		if child {
			g.Id("dc").Qual(converterPkg, "DataConverter")
			g.Id("parentClosePolicy").Qual(enumsPkg, "ParentClosePolicy")
			g.Id("waitForCancellation").Op("*").Bool()
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
		ParamsFunc(func(g *j.Group) {
			if child {
				g.Id("ctx").Qual(workflowPkg, "Context")
			}
			g.Id("req").Qual(protoreflectPkg, "Message")
		}).
		Params(
			j.Qual(optionsPkg, optionsType),
			j.Error(),
		).
		BlockFunc(func(g *j.Group) {
			g.Id("opts").Op(":=").Id("o").Dot("options")

			// set ID
			idFieldName := "ID"
			if child {
				idFieldName = "WorkflowID"
			}
			id := g.If(j.Id("v").Op(":=").Id("o").Dot("id"), j.Id("v").Op("!=").Nil()).Block(
				j.Id("opts").Dot(idFieldName).Op("=").Op("*").Id("v"),
			)
			if idExpr := opts.GetId(); idExpr != "" {
				id.Else().If(j.Id("opts").Dot(idFieldName).Op("==").Lit("")).BlockFunc(func(g *j.Group) {
					// original expression evaluation logic
					origFn := func(g *j.Group, errorReturn j.Code) {
						g.List(j.Id("id"), j.Err()).Op(":=").Qual(expressionPkg, "EvalExpression").CallFunc(func(g *j.Group) {
							g.Id(m.toCamel("%sIDExpression", workflow))
							g.Id("req")
						})
						g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
							g.Return(errorReturn, j.Qual("fmt", "Errorf").Call(j.Lit("error evaluating id expression for %q workflow: %w"), j.Id(m.toCamel("%sWorkflowName", workflow)), j.Err()))
						})
					}

					if child {
						// local activity wrapper
						fixFn := func(g *j.Group) {
							g.Id("lao").Op(":=").Qual(workflowPkg, "GetLocalActivityOptions").Call(j.Id("ctx"))
							g.Id("lao").Dot("ScheduleToCloseTimeout").Op("=").Qual("time", "Second").Op("*").Lit(10)
							g.If(
								j.Err().Op(":=").Qual(workflowPkg, "ExecuteLocalActivity").Call(
									j.Qual(workflowPkg, "WithLocalActivityOptions").Call(j.Id("ctx"), j.Id("lao")),
									j.Func().Params(j.Id("ctx").Qual("context", "Context")).Params(j.String(), j.Error()).BlockFunc(func(g *j.Group) {
										origFn(g, j.Lit(""))
										g.Return(j.Id("id"), j.Nil())
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
							patchComment(g, temporalv1.Patch_PV_64)
							g.If(j.Add(patchVersion(temporalv1.Patch_PV_64, pvm)).Op("==").Lit(1)).BlockFunc(
								fixFn,
							).Else().BlockFunc(func(g *j.Group) {
								origFn(g, j.Id("opts"))
								g.Id("opts").Dot(idFieldName).Op("=").Id("id")
							})
						case temporalv1.Patch_PVM_MARKER:
							g.Add(patchVersion(temporalv1.Patch_PV_64, pvm))
							fixFn(g)
						case temporalv1.Patch_PVM_REMOVED:
							fixFn(g)
						default:
							origFn(g, j.Id("opts"))
							g.Id("opts").Dot(idFieldName).Op("=").Id("id")
						}
					} else {
						origFn(g, j.Id("opts"))
						g.Id("opts").Dot(idFieldName).Op("=").Id("id")
					}
				})
			}

			// set WorkflowIDReusePolicy
			idReusePolicy := g.If(j.Id("v").Op(":=").Id("o").Dot("idReusePolicy"), j.Id("v").Op("!=").Qual(enumsPkg, "WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED")).Block(
				j.Id("opts").Dot("WorkflowIDReusePolicy").Op("=").Id("v"),
			)
			if opts.GetIdReusePolicy() != temporalv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
				idReusePolicy.Else().If(j.Id("opts").Dot("WorkflowIDReusePolicy").Op("==").Qual(enumsPkg, "WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED")).BlockFunc(func(g *j.Group) {
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
					g.Id("opts").Dot("WorkflowIDReusePolicy").Op("=").Qual(enumsPkg, policy)
				})
			}

			// set WorkflowIDConflictPolicy
			if !child {
				idConflictPolicy := g.If(j.Id("v").Op(":=").Id("o").Dot("workflowIdConflictPolicy"), j.Id("v").Op("!=").Qual(enumsPkg, "WORKFLOW_ID_CONFLICT_POLICY_UNSPECIFIED")).Block(
					j.Id("opts").Dot("WorkflowIDConflictPolicy").Op("=").Id("v"),
				)
				if p := opts.GetWorkflowIdConflictPolicy(); p != enums.WORKFLOW_ID_CONFLICT_POLICY_UNSPECIFIED {
					idConflictPolicy.Else().If(j.Id("opts").Dot("WorkflowIDConflictPolicy").Op("==").Qual(enumsPkg, "WORKFLOW_ID_CONFLICT_POLICY_UNSPECIFIED")).BlockFunc(func(g *j.Group) {
						var policy string
						switch p {
						case enums.WORKFLOW_ID_CONFLICT_POLICY_FAIL:
							policy = "WORKFLOW_ID_CONFLICT_POLICY_FAIL"
						case enums.WORKFLOW_ID_CONFLICT_POLICY_TERMINATE_EXISTING:
							policy = "WORKFLOW_ID_CONFLICT_POLICY_TERMINATE_EXISTING"
						case enums.WORKFLOW_ID_CONFLICT_POLICY_USE_EXISTING:
							policy = "WORKFLOW_ID_CONFLICT_POLICY_USE_EXISTING"
						}
						g.Id("opts").Dot("WorkflowIDConflictPolicy").Op("=").Qual(enumsPkg, policy)
					})
				}
			}

			// set TaskQueue
			g.If(j.Id("v").Op(":=").Id("o").Dot("taskQueue"), j.Id("v").Op("!=").Nil()).Block(
				j.Id("opts").Dot("TaskQueue").Op("=").Op("*").Id("v"),
			).Else().If(j.Id("opts").Dot("TaskQueue").Op("==").Lit("")).BlockFunc(func(g *j.Group) {
				var taskQueueVar j.Code
				if tq := opts.GetTaskQueue(); tq != "" {
					taskQueueVar = j.Lit(tq)
				} else if tq = m.opts.GetTaskQueue(); tq != "" {
					taskQueueVar = j.Id(m.toCamel("%sTaskQueue", m.Service.GoName))
				}
				if taskQueueVar != nil {
					g.Id("opts").Dot("TaskQueue").Op("=").Add(taskQueueVar)
				} else if child {
					g.Id("opts").Dot("TaskQueue").Op("=").Qual(workflowPkg, "GetInfo").Call(j.Id("ctx")).Dot("TaskQueueName")
				} else {
					g.Return(j.Id("opts"), j.Qual("errors", "New").Call(j.Lit("TaskQueue is required")))
				}
			})

			// set RetryPolicy
			retryPolicy := g.If(j.Id("v").Op(":=").Id("o").Dot("retryPolicy"), j.Id("v").Op("!=").Nil()).Block(
				j.Id("opts").Dot("RetryPolicy").Op("=").Id("v"),
			)
			if policy := opts.GetRetryPolicy(); policy != nil {
				retryPolicy.Else().If(j.Id("opts").Dot("RetryPolicy").Op("==").Nil()).Block(
					j.Id("opts").Dot("RetryPolicy").Op("=").Op("&").Qual(temporalPkg, "RetryPolicy").CustomFunc(multiLineValues, func(g *j.Group) {
						if d := policy.GetInitialInterval(); d.IsValid() {
							g.Id("InitialInterval").Op(":").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10))
						}
						if d := policy.GetMaxInterval(); d.IsValid() {
							g.Id("MaximumInterval").Op(":").Id(strconv.FormatInt(d.AsDuration().Nanoseconds(), 10))
						}
						if n := policy.GetBackoffCoefficient(); n != 0 {
							g.Id("BackoffCoefficient").Op(":").Lit(n)
						}
						if n := policy.GetMaxAttempts(); n != 0 {
							g.Id("MaximumAttempts").Op(":").Lit(n)
						}
						if errs := policy.GetNonRetryableErrorTypes(); len(errs) > 0 {
							g.Id("NonRetryableErrorTypes").Op(":").Index().String().CustomFunc(multiLineValues, func(g *j.Group) {
								for _, err := range errs {
									g.Lit(err)
								}
							})
						}
					}),
				)
			}

			// set SearchAttributes
			searchAttributes := g.If(j.Id("v").Op(":=").Id("o").Dot("searchAttributes"), j.Id("v").Op("!=").Nil()).Block(
				j.Id("opts").Dot("SearchAttributes").Op("=").Id("o").Dot("searchAttributes"),
			)
			if mapping := opts.GetSearchAttributes(); mapping != "" {
				searchAttributes.Else().If(j.Id("opts").Dot("SearchAttributes").Op("==").Nil()).
					BlockFunc(func(g *j.Group) {
						// original expression evaluation logic
						origFn := func(g *j.Group, errorResult j.Code) {
							// convert input to generic mapping input
							g.List(j.Id("structured"), j.Err()).Op(":=").Qual(expressionPkg, "ToStructured").Call(j.Id("req"))
							g.If(j.Err().Op("!=").Nil()).Block(
								j.Return(errorResult, j.Qual("fmt", "Errorf").Call(j.Lit(fmt.Sprintf("error serializing input for %q search attribute mapping: %%v", m.methods[workflow].GoName)), j.Err())),
							)

							// execute mapping
							g.List(j.Id("result"), j.Err()).Op(":=").Id(m.toCamel("%sSearchAttributesMapping", workflow)).Dot("Query").Call(j.Id("structured"))
							g.If(j.Err().Op("!=").Nil()).Block(
								j.Return(errorResult, j.Qual("fmt", "Errorf").Call(j.Lit(fmt.Sprintf("error executing %q search attribute mapping: %%v", m.methods[workflow].GoName)), j.Err())),
							)

							// coerce mapping result to map[string]any
							g.List(j.Id("searchAttributes"), j.Id("ok")).Op(":=").Id("result").Op(".").Parens(j.Map(j.String()).Any())
							g.If(j.Op("!").Id("ok")).Block(
								j.Return(errorResult, j.Qual("fmt", "Errorf").Call(j.Lit(fmt.Sprintf("expected %q search attribute mapping to return map[string]any, got: %%T", m.methods[workflow].GoName)), j.Id("result"))),
							)
						}

						if child {
							// local activity wrapper
							fixFn := func(g *j.Group) {
								g.Id("lao").Op(":=").Qual(workflowPkg, "GetLocalActivityOptions").Call(j.Id("ctx"))
								g.Id("lao").Dot("ScheduleToCloseTimeout").Op("=").Qual("time", "Second").Op("*").Lit(10)
								g.If(
									j.Err().Op(":=").Qual(workflowPkg, "ExecuteLocalActivity").Call(
										j.Qual(workflowPkg, "WithLocalActivityOptions").Call(j.Id("ctx"), j.Id("lao")),
										j.Func().Params(j.Id("ctx").Qual("context", "Context")).Params(j.Map(j.String()).Any(), j.Error()).BlockFunc(func(g *j.Group) {
											origFn(g, j.Nil())
											g.Return(j.Id("searchAttributes"), j.Nil())
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
								patchComment(g, temporalv1.Patch_PV_64)
								g.If(j.Add(patchVersion(temporalv1.Patch_PV_64, pvm)).Op("==").Lit(1)).BlockFunc(
									fixFn,
								).Else().BlockFunc(func(g *j.Group) {
									origFn(g, j.Id("opts"))
									g.Id("opts").Dot("SearchAttributes").Op("=").Id("searchAttributes")
								})
							case temporalv1.Patch_PVM_MARKER:
								g.Add(patchVersion(temporalv1.Patch_PV_64, pvm))
								fixFn(g)
							case temporalv1.Patch_PVM_REMOVED:
								fixFn(g)
							default:
								origFn(g, j.Id("opts"))
								g.Id("opts").Dot("SearchAttributes").Op("=").Id("searchAttributes")
							}
						} else {
							origFn(g, j.Id("opts"))
							g.Id("opts").Dot("SearchAttributes").Op("=").Id("searchAttributes")
						}
					})
			}

			// set TypedSearchAttributes
			typedSearchAttributes := g.IfFunc(func(g *j.Group) {
				g.Id("v").Op(":=").Id("o").Dot("typedSearchAttributes")
				g.Id("v").Op("!=").Nil()
			}).BlockFunc(func(g *j.Group) {
				g.Id("opts").Dot("TypedSearchAttributes").Op("=").Op("*").Id("v")
			})
			if mapping := opts.GetTypedSearchAttributes(); mapping != "" {
				typedSearchAttributes.Else().IfFunc(func(g *j.Group) {
					g.Id("opts").Dot("TypedSearchAttributes").Dot("Size").Call().Op("==").Lit(0)
				}).BlockFunc(func(g *j.Group) {
					saFn := func(g *j.Group, errorResult j.Code) {
						// initialize structured mapping input
						if hasInput {
							g.List(j.Id("structured"), j.Err()).Op(":=").
								Qual(expressionPkg, "ToStructured").Call(j.Id("req"))
							g.IfFunc(func(g *j.Group) {
								g.Err().Op("!=").Nil()
							}).BlockFunc(func(g *j.Group) {
								g.ReturnFunc(func(g *j.Group) {
									g.Add(errorResult)
									g.Qual("fmt", "Errorf").CallFunc(func(g *j.Group) {
										g.Lit(fmt.Sprintf(
											"error serializing input for %q typed search attribute "+
												"mapping: %%v",
											m.methods[workflow].GoName,
										))
										g.Err()
									})
								})
							})
						} else {
							g.Id("structured").Op(":=").Make(j.Map(j.String()).Any())
						}

						// execute mapping
						g.List(j.Id("result"), j.Err()).Op(":=").
							Id(m.Names().workflowTypedSearchAttributesMapping(workflow)).
							Dot("Query").Call(j.Id("structured"))
						g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
							g.ReturnFunc(func(g *j.Group) {
								g.Add(errorResult)
								g.Qual("fmt", "Errorf").CallFunc(func(g *j.Group) {
									g.Lit(fmt.Sprintf(
										"error executing %q typed search attribute mapping: %%v",
										m.methods[workflow].GoName,
									))
									g.Err()
								})
							})
						})

						// coerce mapping result to map[string]any
						g.List(j.Id("sa"), j.Id("ok")).Op(":=").
							Id("result").Op(".").Parens(j.Map(j.String()).Any())
						g.If(j.Op("!").Id("ok")).BlockFunc(func(g *j.Group) {
							g.ReturnFunc(func(g *j.Group) {
								g.Add(errorResult)
								g.Qual("fmt", "Errorf").CallFunc(func(g *j.Group) {
									g.Lit(fmt.Sprintf(
										"expected %q typed search attribute mapping to return "+
											"map[string]any, got: %%T",
										m.methods[workflow].GoName,
									))
									g.Id("result")
								})
							})
						})
					}
					tsaFn := func(g *j.Group, errorResult j.Code) {
						// marshal mapping result to SearchAttributes
						g.List(j.Id("tsa"), j.Err()).Op(":=").
							Qual(convertPkg, "MarshalTypedSearchAttributes").Call(j.Id("sa"))
						g.IfFunc(func(g *j.Group) {
							g.Err().Op("!=").Nil()
						}).BlockFunc(func(g *j.Group) {
							g.ReturnFunc(func(g *j.Group) {
								g.Add(errorResult)
								g.Qual("fmt", "Errorf").CallFunc(func(g *j.Group) {
									g.Lit(fmt.Sprintf(
										"error marshaling %q typed search attribute mapping: %%v",
										m.methods[workflow].GoName,
									))
									g.Err()
								})
							})
						})
					}
					if child {
						g.Id("sa").Op(":=").Make(j.Map(j.String()).Any())
						g.IfFunc(func(g *j.Group) {
							g.Err().Op(":=").Qual(workflowPkg, "ExecuteLocalActivity").
								CallFunc(func(g *j.Group) {
									g.Qual(workflowPkg, "WithLocalActivityOptions").CallFunc(func(g *j.Group) {
										g.Id("ctx")
										g.Qual(workflowPkg, "LocalActivityOptions").ValuesFunc(func(g *j.Group) {
											g.Id("StartToCloseTimeout").Op(":").Qual("time", "Second").Op("*").Lit(10)
										})
									})
									g.Func().Params(j.Id("ctx").Qual("context", "Context")).Params(j.Map(j.String()).Any(), j.Error()).BlockFunc(func(g *j.Group) {
										saFn(g, j.Nil())
										g.Return(j.Id("sa"), j.Nil())
									})
								}).
								Dot("Get").Call(j.Id("ctx"), j.Op("&").Id("sa"))
							g.Err().Op("!=").Nil()
						}).BlockFunc(func(g *j.Group) {
							g.ReturnFunc(func(g *j.Group) {
								g.Add(j.Id("opts"))
								g.Qual("fmt", "Errorf").CallFunc(func(g *j.Group) {
									g.Lit(fmt.Sprintf(
										"error evaluating typed search attributes for %q workflow: %%w",
										m.toCamel("%sWorkflowName", workflow),
									))
									g.Err()
								})
							})
						})
					} else {
						saFn(g, j.Id("opts"))
					}
					tsaFn(g, j.Id("opts"))
					g.Id("opts").Dot("TypedSearchAttributes").Op("=").Id("tsa")
				})
			}

			// set WorkflowExecutionTimeout
			executionTimeout := g.If(j.Id("v").Op(":=").Id("o").Dot("executionTimeout"), j.Id("v").Op("!=").Nil()).Block(
				j.Id("opts").Dot("WorkflowExecutionTimeout").Op("=").Op("*").Id("v"),
			)
			if d := opts.GetExecutionTimeout().AsDuration(); d > 0 {
				executionTimeout.Else().If(j.Id("opts").Dot("WorkflowExecutionTimeout").Op("==").Lit(0)).Block(
					j.Id("opts").Dot("WorkflowExecutionTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
				)
			}

			// set WorkflowRunTimeout
			runTimeout := g.If(j.Id("v").Op(":=").Id("o").Dot("runTimeout"), j.Id("v").Op("!=").Nil()).Block(
				j.Id("opts").Dot("WorkflowRunTimeout").Op("=").Op("*").Id("v"),
			)
			if d := opts.GetRunTimeout().AsDuration(); d > 0 {
				runTimeout.Else().If(j.Id("opts").Dot("WorkflowRunTimeout").Op("==").Lit(0)).Block(
					j.Id("opts").Dot("WorkflowRunTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
				)
			}

			// set WorkflowTaskTimeout
			taskTimeout := g.If(j.Id("v").Op(":=").Id("o").Dot("taskTimeout"), j.Id("v").Op("!=").Nil()).Block(
				j.Id("opts").Dot("WorkflowTaskTimeout").Op("=").Op("*").Id("v"),
			)
			if d := opts.GetTaskTimeout().AsDuration(); d > 0 {
				taskTimeout.Else().If(j.Id("opts").Dot("WorkflowTaskTimeout").Op("==").Lit(0)).Block(
					j.Id("opts").Dot("WorkflowTaskTimeout").Op("=").Id(strconv.FormatInt(d.Nanoseconds(), 10)).Comment(durafmt.Parse(d).String()),
				)
			}

			// set ParentClosePolicy
			if child {
				parentClosePolicy := g.If(j.Id("v").Op(":=").Id("o").Dot("parentClosePolicy"), j.Id("v").Op("!=").Qual(enumsPkg, "PARENT_CLOSE_POLICY_UNSPECIFIED")).Block(
					j.Id("opts").Dot("ParentClosePolicy").Op("=").Id("v"),
				)
				if opts.GetParentClosePolicy() != temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_UNSPECIFIED {
					parentClosePolicy.Else().If(j.Id("opts").Dot("ParentClosePolicy").Op("==").Qual(enumsPkg, "PARENT_CLOSE_POLICY_UNSPECIFIED")).BlockFunc(func(g *j.Group) {
						var policy string
						switch opts.GetParentClosePolicy() {
						case temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_ABANDON:
							policy = "PARENT_CLOSE_POLICY_ABANDON"
						case temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL:
							policy = "PARENT_CLOSE_POLICY_REQUEST_CANCEL"
						case temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE:
							policy = "PARENT_CLOSE_POLICY_TERMINATE"
						}
						g.Id("opts").Dot("ParentClosePolicy").Op("=").Qual(enumsPkg, policy)
					})
				}
			}

			// set WaitForCancellation
			if child {
				waitForCancellation := g.If(j.Id("v").Op(":=").Id("o").Dot("waitForCancellation"), j.Id("v").Op("!=").Nil()).Block(
					j.Id("opts").Dot("WaitForCancellation").Op("=").Op("*").Id("v"),
				)
				if opts.GetWaitForCancellation() {
					waitForCancellation.Else().Block(
						j.Id("opts").Dot("WaitForCancellation").Op("=").Lit(true),
					)
				}
			}

			g.Return(j.Id("opts"), j.Nil())
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

	if child {
		f.Comment("WithDataConverter registers a DataConverter for the child workflow")
		f.Func().
			Params(j.Id("o").Op("*").Id(typeName)).
			Id("WithDataConverter").
			Params(j.Id("dc").Qual(converterPkg, "DataConverter")).
			Op("*").Id(typeName).
			Block(
				j.Id("o").Dot("dc").Op("=").Id("dc"),
				j.Return(j.Id("o")),
			)
	}

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

	f.Comment("WithTypedSearchAttributes sets the TypedSearchAttributes value")
	f.Func().
		Params(j.Id("o").Op("*").Id(typeName)).
		Id("WithTypedSearchAttributes").
		Params(j.Id("tsa").Qual(temporalPkg, "SearchAttributes")).
		Op("*").Id(typeName).
		Block(
			j.Id("o").Dot("typedSearchAttributes").Op("=").Op("&").Id("tsa"),
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
