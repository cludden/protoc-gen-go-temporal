package plugin

import (
	"fmt"

	j "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/compiler/protogen"
)

const (
	schemePkg = "github.com/cludden/protoc-gen-go-temporal/pkg/scheme"
)

func (m *Manifest) renderCodec(f *j.File) {
	optName := m.toCamel("With%sSchemeTypes", m.GoName)

	f.Commentf("%s registers all %s protobuf types with the given scheme", optName, m.GoName)
	f.Func().
		Id(optName).
		Params().
		Qual(schemePkg, "Option").
		Block(
			j.Return(
				j.Func().
					Params(j.Id("s").Op("*").Qual(schemePkg, "Scheme")).
					BlockFunc(func(g *j.Group) {
						types := make(map[string]struct{})
						for _, a := range m.activitiesOrdered {
							if m.methods[a].Desc.Parent() != m.Service.Desc {
								continue
							}
							method := m.methods[a]
							registerType(m.Plugin.Plugin, g, types, m.Service, nil, method.Input)
							registerType(m.Plugin.Plugin, g, types, m.Service, nil, method.Output)
						}
						for _, q := range m.queriesOrdered {
							if m.methods[q].Desc.Parent() != m.Service.Desc {
								continue
							}
							method := m.methods[q]
							registerType(m.Plugin.Plugin, g, types, m.Service, nil, method.Input)
							registerType(m.Plugin.Plugin, g, types, m.Service, nil, method.Output)
						}
						for _, s := range m.signalsOrdered {
							if m.methods[s].Desc.Parent() != m.Service.Desc {
								continue
							}
							method := m.methods[s]
							registerType(m.Plugin.Plugin, g, types, m.Service, nil, method.Input)
							registerType(m.Plugin.Plugin, g, types, m.Service, nil, method.Output)
						}
						for _, u := range m.updatesOrdered {
							if m.methods[u].Desc.Parent() != m.Service.Desc {
								continue
							}
							method := m.methods[u]
							registerType(m.Plugin.Plugin, g, types, m.Service, nil, method.Input)
							registerType(m.Plugin.Plugin, g, types, m.Service, nil, method.Output)
						}
						for _, w := range m.workflowsOrdered {
							if m.methods[w].Desc.Parent() != m.Service.Desc {
								continue
							}
							method := m.methods[w]
							registerType(m.Plugin.Plugin, g, types, m.Service, nil, method.Input)
							registerType(m.Plugin.Plugin, g, types, m.Service, nil, method.Output)
						}
					}),
			),
		)
}

func registerType(p *protogen.Plugin, g *j.Group, cache map[string]struct{}, svc *protogen.Service, messages j.Code, msg *protogen.Message) {
	if _, ok := cache[string(msg.Desc.FullName())]; ok || isEmpty(msg) {
		return
	}
	f, ok := p.FilesByPath[msg.Desc.ParentFile().Path()]
	if !ok {
		p.Error(fmt.Errorf("unable to locate parent file for msg: %s", msg.Desc.ParentFile().Path()))
		return
	}

	if messages == nil {
		if string(msg.Desc.ParentFile().Package()) != string(svc.Desc.ParentFile().Package()) {
			messages = j.Qual(string(f.GoImportPath), f.GoDescriptorIdent.GoName)
		} else {
			messages = j.Id(f.GoDescriptorIdent.GoName)
		}
	}
	g.Id("s").
		Dot("RegisterType").
		Call(
			j.Add(messages).
				Dot("Messages").
				Call().
				Dot("ByName").
				Call(j.Lit(string(msg.Desc.Name()))),
		)
	cache[string(msg.Desc.FullName())] = struct{}{}

	messages = j.Add(messages).
		Dot("Messages").
		Call().
		Dot("ByName").
		Call(j.Lit(string(msg.Desc.Name())))

	for _, f := range msg.Messages {
		registerType(p, g, cache, svc, messages, f)
	}
}
