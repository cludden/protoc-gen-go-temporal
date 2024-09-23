package plugin

import (
	"fmt"

	g "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/compiler/protogen"
)

const (
	schemePkg = "github.com/cludden/protoc-gen-go-temporal/pkg/scheme"
)

func (svc *Manifest) renderCodec(f *g.File) {
	svc.genSectionHeader(f, "Codec")
	optName := svc.toCamel("With%sSchemeTypes", svc.GoName)

	f.Commentf("%s registers all %s protobuf types with the given scheme", optName, svc.GoName)
	f.Func().
		Id(optName).
		Params().
		Qual(schemePkg, "Option").
		Block(
			g.Return(
				g.Func().
					Params(g.Id("s").Op("*").Qual(schemePkg, "Scheme")).
					BlockFunc(func(fn *g.Group) {
						types := make(map[string]struct{})
						for _, a := range svc.activitiesOrdered {
							if svc.methods[a].Desc.Parent() != svc.Service.Desc {
								continue
							}
							method := svc.methods[a]
							registerType(svc.Plugin.Plugin, fn, types, svc.Service, nil, method.Input)
							registerType(svc.Plugin.Plugin, fn, types, svc.Service, nil, method.Output)
						}
						for _, q := range svc.queriesOrdered {
							if svc.methods[q].Desc.Parent() != svc.Service.Desc {
								continue
							}
							method := svc.methods[q]
							registerType(svc.Plugin.Plugin, fn, types, svc.Service, nil, method.Input)
							registerType(svc.Plugin.Plugin, fn, types, svc.Service, nil, method.Output)
						}
						for _, s := range svc.signalsOrdered {
							if svc.methods[s].Desc.Parent() != svc.Service.Desc {
								continue
							}
							method := svc.methods[s]
							registerType(svc.Plugin.Plugin, fn, types, svc.Service, nil, method.Input)
							registerType(svc.Plugin.Plugin, fn, types, svc.Service, nil, method.Output)
						}
						for _, u := range svc.updatesOrdered {
							if svc.methods[u].Desc.Parent() != svc.Service.Desc {
								continue
							}
							method := svc.methods[u]
							registerType(svc.Plugin.Plugin, fn, types, svc.Service, nil, method.Input)
							registerType(svc.Plugin.Plugin, fn, types, svc.Service, nil, method.Output)
						}
						for _, w := range svc.workflowsOrdered {
							if svc.methods[w].Desc.Parent() != svc.Service.Desc {
								continue
							}
							method := svc.methods[w]
							registerType(svc.Plugin.Plugin, fn, types, svc.Service, nil, method.Input)
							registerType(svc.Plugin.Plugin, fn, types, svc.Service, nil, method.Output)
						}
					}),
			),
		)
}

func registerType(p *protogen.Plugin, fn *g.Group, cache map[string]struct{}, svc *protogen.Service, messages g.Code, msg *protogen.Message) {
	if string(svc.Desc.ParentFile().Package()) != string(msg.Desc.ParentFile().Package()) {
		return
	}

	if _, ok := cache[string(msg.Desc.FullName())]; ok || isEmpty(msg) {
		return
	}
	f, ok := p.FilesByPath[msg.Desc.ParentFile().Path()]
	if !ok {
		p.Error(fmt.Errorf("unable to locate parent file for msg: %s", msg.Desc.ParentFile().Path()))
		return
	}
	if messages == nil {
		messages = g.Id(f.GoDescriptorIdent.GoName)
	}
	fn.Id("s").
		Dot("RegisterType").
		Call(
			g.Add(messages).
				Dot("Messages").
				Call().
				Dot("ByName").
				Call(g.Lit(string(msg.Desc.Name()))),
		)
	cache[string(msg.Desc.FullName())] = struct{}{}

	messages = g.Add(messages).
		Dot("Messages").
		Call().
		Dot("ByName").
		Call(g.Lit(string(msg.Desc.Name())))

	for _, f := range msg.Messages {
		registerType(p, fn, cache, svc, messages, f)
	}
}
