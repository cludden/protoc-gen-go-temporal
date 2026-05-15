package plugin

import (
	"slices"
	"strings"

	nexusv1 "github.com/bergundy/nexus-proto-annotations/go/nexus/v1"
	j "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func (m *Manifest) renderNexus(f *j.File, file *protogen.File, svc *protogen.Service) bool {
	if !m.nexusGetShouldIncludeService(svc) {
		return false
	} else if !slices.ContainsFunc(svc.Methods, m.nexusGetShouldIncludeOperation) {
		return false
	}
	m.nexusGenHandlerImpl(f, file, svc)
	return true
}

func (m *Manifest) nexusGenHandlerImpl(f *j.File, file *protogen.File, svc *protogen.Service) {
	handlerImpl := m.Names().nexusHandler(svc)

	f.Commentf("Nexus handler for %s", svc.Desc.FullName())
	f.Comment("")
	f.Comment("Deprecated: Use native nexus implementation instead")
	f.Type().Id(handlerImpl).Struct()

	for _, workflow := range m.workflowsOrdered {
		if m.methods[workflow].Desc.Parent() != m.Service.Desc {
			continue
		}
		method := m.methods[workflow]
		if !m.nexusGetShouldIncludeOperation(method) {
			continue
		}
		input, hasInput, output, _ := m.nexusGetMethodIO(method)

		operation := method.GoName
		commentWithDefaultf(f, methodSet(method), "Nexus operation for %s workflow", m.fqnForWorkflow(workflow))
		f.Func().
			ParamsFunc(func(g *j.Group) {
				g.Id("h").Op("*").Id(handlerImpl)
			}).
			Id(operation).
			ParamsFunc(func(g *j.Group) {
				g.Id("name").String()
			}).
			Qual(nexusPkg, "Operation").
			TypesFunc(func(g *j.Group) {
				g.Add(input)
				g.Add(output)
			}).
			BlockFunc(func(g *j.Group) {
				g.ReturnFunc(func(g *j.Group) {
					g.Qual(temporalnexusPkg, "MustNewWorkflowRunOperationWithOptions").CallFunc(func(g *j.Group) {
						g.Qual(temporalnexusPkg, "WorkflowRunOperationOptions").
							TypesFunc(func(g *j.Group) {
								g.Add(input)
								g.Add(output)
							}).
							Values(j.Dict{
								j.Id("Name"): j.Id("name"),
								j.Id("Handler"): j.Func().
									ParamsFunc(func(g *j.Group) {
										g.Id("ctx").Qual("context", "Context")
										g.Id("input").Add(input)
										g.Id("opts").Qual(nexusPkg, "StartOperationOptions")
									}).
									ParamsFunc(func(g *j.Group) {
										g.Qual(temporalnexusPkg, "WorkflowHandle").TypesFunc(func(g *j.Group) {
											g.Add(output)
										})
										g.Error()
									}).
									BlockFunc(func(g *j.Group) {
										g.List(j.Id("o"), j.Err()).Op(":=").Qual(string(file.GoImportPath), m.toCamel("New%sOptions", workflow)).Call().Dot("Build").CallFunc(func(g *j.Group) {
											if hasInput {
												g.Id("input").Dot("ProtoReflect").Call()
											} else {
												g.Nil()
											}
										})
										g.If(j.Err().Op("!=").Nil()).BlockFunc(func(g *j.Group) {
											g.Return(
												j.Nil(),
												j.Qual(nexusPkg, "HandlerErrorf").Call(
													j.Qual(nexusPkg, "HandlerErrorTypeBadRequest"),
													j.Lit("failed to build workflow options: %w"),
													j.Err(),
												),
											)
										})
										g.Return(
											j.Qual(temporalnexusPkg, "ExecuteUntypedWorkflow").TypesFunc(func(g *j.Group) {
												g.Add(output)
											}).CallFunc(func(g *j.Group) {
												g.Id("ctx")
												g.Id("opts")
												g.Id("o")
												g.Qual(string(file.GoImportPath), m.Names().workflowName(workflow))
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
}

func (p *Manifest) nexusGetMethodIO(m *protogen.Method) (input *j.Statement, hasInput bool, output *j.Statement, hasOutput bool) {
	in := m.Input
	if !isEmpty(in) {
		hasInput = true
		input = j.Op("*").Qual(string(in.GoIdent.GoImportPath), in.GoIdent.GoName)
	} else {
		input = j.Qual(nexusPkg, "NoValue")
	}
	out := m.Output
	if !isEmpty(out) {
		hasOutput = true
		output = j.Op("*").Qual(string(out.GoIdent.GoImportPath), out.GoIdent.GoName)
	} else {
		output = j.Qual(nexusPkg, "NoValue")
	}
	return input, hasInput, output, hasOutput
}

func (p *Manifest) nexusGetShouldIncludeOperation(m *protogen.Method) bool {
	if _, isWorkflow := p.workflows[m.Desc.FullName()]; !isWorkflow {
		return false
	}
	tags := p.nexusGetOperationOptions(m).GetTags()
	if len(p.includeOperationTags) > 0 && !slices.ContainsFunc(tags, func(t string) bool {
		_, ok := p.includeOperationTags[t]
		return ok
	}) {
		return false
	}
	return !slices.ContainsFunc(tags, func(t string) bool {
		_, ok := p.excludeOperationTags[t]
		return ok
	})
}

func (m *Manifest) nexusGetShouldIncludeService(svc *protogen.Service) bool {
	tags := m.nexusGetServiceOptions(svc).GetTags()
	if len(m.includeServiceTags) > 0 && !slices.ContainsFunc(tags, func(t string) bool {
		_, ok := m.includeServiceTags[t]
		return ok
	}) {
		return false
	}
	return !slices.ContainsFunc(tags, func(t string) bool {
		_, ok := m.excludeServiceTags[t]
		return ok
	})
}

func (m *Plugin) nexusGetTags(tags string) map[string]struct{} {
	index := make(map[string]struct{})
	for _, tag := range strings.Split(tags, ";") {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}
		index[tag] = struct{}{}
	}
	return index
}

func (p *Plugin) nexusGetOperationOptions(m *protogen.Method) *nexusv1.OperationOptions {
	opts, _ := proto.GetExtension(m.Desc.Options(), nexusv1.E_Operation).(*nexusv1.OperationOptions)
	return opts
}

func (p *Plugin) nexusGetServiceOptions(svc *protogen.Service) *nexusv1.ServiceOptions {
	opts, _ := proto.GetExtension(svc.Desc.Options(), nexusv1.E_Service).(*nexusv1.ServiceOptions)
	return opts
}

func (n *names) nexusHandler(service *protogen.Service) string {
	return n.toCamel("%sNexusHandler", service.GoName)
}

func (n *names) nexusRegisterService(service *protogen.Service) string {
	return n.toCamel("Register%sNexusService", service.GoName)
}
