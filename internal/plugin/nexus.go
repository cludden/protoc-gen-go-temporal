package plugin

import (
	"fmt"
	"path"

	"github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (svc *Manifest) renderProtocGenGoNexus(f *jen.File) {
	svc.genProtocGenGoNexusHandler(f)
}

func (svc *Manifest) genProtocGenGoNexusHandler(f *jen.File) {
	typeName := svc.toCamel("%sNexusHandler", svc.Service.GoName)
	nexusGoPackageName := fmt.Sprintf("%snexus", svc.File.GoPackageName)
	nexusGoImportPath := path.Join(string(svc.File.GoImportPath), nexusGoPackageName)

	f.Commentf("%s is an implementation of the protoc-gen-go-nexus handler", typeName)
	f.Type().Id(typeName).StructFunc(func(g *jen.Group) {
		g.Qual(nexusGoImportPath, svc.toCamel("Unimplemented%sNexusHandler", svc.Service.GoName))
	})

	svc.genProtocGenGoNexusRegisterService(f, typeName)

	for _, workflow := range svc.workflowsOrdered {
		if svc.methods[workflow].Desc.Parent() != svc.Service.Desc {
			continue
		}
		svc.genProtocGenGoNexusHandlerWorkflow(f, typeName, workflow)
	}
}

func (svc *Manifest) genProtocGenGoNexusHandlerWorkflow(f *jen.File, typeName string, w protoreflect.FullName) {
	method := svc.methods[w]
	methodName := svc.toCamel("%s", method.GoName)
	hasInput := !isEmpty(method.Input)
	hasOutput := !isEmpty(method.Output)

	commentWithDefaultf(f, methodSet(method), "returns a nexus operation for executing a %s workflow", svc.fqnForWorkflow(w))
	f.Func().
		Params(jen.Id("h").Op("*").Id(typeName)).
		Id(methodName).
		Params(jen.Id("name").String()).
		Params(
			jen.Qual("github.com/nexus-rpc/sdk-go/nexus", "Operation").
				TypesFunc(func(g *jen.Group) {
					for _, msg := range []*protogen.Message{method.Input, method.Output} {
						if isEmpty(msg) {
							g.Qual("github.com/nexus-rpc/sdk-go/nexus", "NoValue")
						} else {
							g.Op("*").Qual(string(msg.GoIdent.GoImportPath), msg.GoIdent.GoName)
						}
					}
				}),
		).
		BlockFunc(func(g *jen.Group) {
			g.ReturnFunc(func(g *jen.Group) {
				g.Qual("go.temporal.io/sdk/temporalnexus", "MustNewWorkflowRunOperationWithOptions").CallFunc(func(g *jen.Group) {
					g.Qual("go.temporal.io/sdk/temporalnexus", "WorkflowRunOperationOptions").
						TypesFunc(func(g *jen.Group) {
							for _, msg := range []*protogen.Message{method.Input, method.Output} {
								if isEmpty(msg) {
									g.Qual("github.com/nexus-rpc/sdk-go/nexus", "NoValue")
								} else {
									g.Op("*").Qual(string(msg.GoIdent.GoImportPath), msg.GoIdent.GoName)
								}
							}
						}).
						Values(jen.Dict{
							jen.Id("Name"): jen.Id("name"),
							jen.Id("Handler"): jen.Func().
								ParamsFunc(func(g *jen.Group) {
									g.Id("ctx").Qual("context", "Context")
									if hasInput {
										g.Id("input").Op("*").Qual(string(method.Input.GoIdent.GoImportPath), svc.getMessageName(method.Input))
									} else {
										g.Id("_").Qual("github.com/nexus-rpc/sdk-go/nexus", "NoValue")
									}
									g.Id("opts").Qual("github.com/nexus-rpc/sdk-go/nexus", "StartOperationOptions")
								}).
								ParamsFunc(func(g *jen.Group) {
									g.Qual("go.temporal.io/sdk/temporalnexus", "WorkflowHandle").TypesFunc(func(g *jen.Group) {
										if hasOutput {
											g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
										} else {
											g.Qual("github.com/nexus-rpc/sdk-go/nexus", "NoValue")
										}
									})
									g.Error()
								}).
								BlockFunc(func(g *jen.Group) {
									g.List(jen.Id("o"), jen.Err()).Op(":=").Qual(string(svc.File.GoImportPath), svc.toCamel("New%sOptions", w)).Call().Dot("Build").CallFunc(func(g *jen.Group) {
										if hasInput {
											g.Id("input").Dot("ProtoReflect").Call()
										} else {
											g.Nil()
										}
									})
									g.If(jen.Err().Op("!=").Nil()).BlockFunc(func(g *jen.Group) {
										g.Return(jen.Nil(), jen.Err())
									})
									g.Return(
										jen.Qual("go.temporal.io/sdk/temporalnexus", "ExecuteUntypedWorkflow").TypesFunc(func(g *jen.Group) {
											if hasOutput {
												g.Op("*").Qual(string(method.Output.GoIdent.GoImportPath), svc.getMessageName(method.Output))
											} else {
												g.Qual("github.com/nexus-rpc/sdk-go/nexus", "NoValue")
											}
										}).CallFunc(func(g *jen.Group) {
											g.Id("ctx")
											g.Id("opts")
											g.Id("o")
											g.Qual(string(svc.File.GoImportPath), svc.toCamel("%sWorkflowName", w))
											if hasInput {
												g.Id("input")
											}
										}),
									)
								}),
						})
				})
			})
		})
}

func (svc *Manifest) genProtocGenGoNexusRegisterService(f *jen.File, handlerTypeName string) {
	fnName := fmt.Sprintf("Register%sNexusService", svc.Service.GoName)
	nexusGoPackageName := fmt.Sprintf("%snexus", svc.File.GoPackageName)
	nexusGoImportPath := path.Join(string(svc.File.GoImportPath), nexusGoPackageName)

	f.Commentf("%s initializes a new %s nexus service and registers it with the provided registry", fnName, svc.Service.GoName)
	f.Func().
		Id(fnName).
		ParamsFunc(func(g *jen.Group) {
			g.Id("r").Qual("go.temporal.io/sdk/worker", "NexusServiceRegistry")
		}).
		Error().
		BlockFunc(func(g *jen.Group) {
			g.List(jen.Id("svc"), jen.Err()).Op(":=").
				Qual(nexusGoImportPath, svc.toCamel("New%sNexusService", svc.Service.GoName)).
				Call(jen.Op("&").Id(handlerTypeName).Values())
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Err()),
			)
			g.Id("r").Dot("RegisterNexusService").Call(jen.Id("svc"))
			g.Return(jen.Nil())
		})
}
