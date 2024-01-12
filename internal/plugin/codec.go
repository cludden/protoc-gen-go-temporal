package plugin

import (
	"fmt"

	g "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/compiler/protogen"
)

const (
	schemePkg = "github.com/cludden/protoc-gen-go-temporal/pkg/scheme"
)

func (svc *Service) renderCodec(f *g.File) {
	optName := toCamel("With%sSchemeTypes", svc.GoName)

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
							method := svc.methods[a]
							registerType(svc.Plugin, fn, types, method.Input)
							registerType(svc.Plugin, fn, types, method.Output)
						}
						for _, q := range svc.queriesOrdered {
							method := svc.methods[q]
							registerType(svc.Plugin, fn, types, method.Input)
							registerType(svc.Plugin, fn, types, method.Output)
						}
						for _, s := range svc.signalsOrdered {
							method := svc.methods[s]
							registerType(svc.Plugin, fn, types, method.Input)
							registerType(svc.Plugin, fn, types, method.Output)
						}
						for _, u := range svc.updatesOrdered {
							method := svc.methods[u]
							registerType(svc.Plugin, fn, types, method.Input)
							registerType(svc.Plugin, fn, types, method.Output)
						}
						for _, w := range svc.workflowsOrdered {
							method := svc.methods[w]
							registerType(svc.Plugin, fn, types, method.Input)
							registerType(svc.Plugin, fn, types, method.Output)
						}
					}),
			),
		)
}

func registerType(p *protogen.Plugin, fn *g.Group, cache map[string]struct{}, msg *protogen.Message) {
	if _, ok := cache[string(msg.Desc.FullName())]; ok || isEmpty(msg) {
		return
	}
	f, ok := p.FilesByPath[msg.Desc.ParentFile().Path()]
	if !ok {
		p.Error(fmt.Errorf("unable to locate parent file for msg: %s", msg.Desc.ParentFile().Path()))
		return
	}
	fn.Id("s").
		Dot("RegisterType").
		Call(
			g.Id(f.GoDescriptorIdent.GoName).
				Dot("Messages").
				Call().
				Dot("ByName").
				Call(g.Lit(string(msg.Desc.FullName().Name()))),
		)
	cache[string(msg.Desc.FullName())] = struct{}{}
}
