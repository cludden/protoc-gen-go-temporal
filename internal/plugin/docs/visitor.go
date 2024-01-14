package docs

import (
	"sort"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type visitor struct {
	seen map[string]struct{}
	refs map[string][]string
}

func newVisitor() *visitor {
	return &visitor{
		seen: make(map[string]struct{}),
		refs: make(map[string][]string),
	}
}

func (v *visitor) walk(p *protogen.Plugin) map[string][]string {
	for _, f := range p.Files {
		for _, svc := range f.Services {
			for _, m := range svc.Methods {
				for _, msg := range []*protogen.Message{m.Input, m.Output} {
					v.walkMessage(msg)
				}
			}
		}
	}
	for pkg, refs := range v.refs {
		sort.Strings(refs)
		v.refs[pkg] = refs
	}
	return v.refs
}

func (v *visitor) walkMessage(msg *protogen.Message) {
	if !notEmpty(msg) {
		return
	}

	if _, ok := v.seen[string(msg.Desc.FullName())]; ok {
		return
	}

	pkgName := string(msg.Desc.ParentFile().Package())
	if _, ok := v.refs[pkgName]; !ok {
		v.refs[pkgName] = []string{}
	}
	v.refs[pkgName] = append(v.refs[pkgName], string(msg.Desc.FullName()))
	v.seen[string(msg.Desc.FullName())] = struct{}{}

	for _, field := range msg.Fields {
		switch field.Desc.Kind() {
		case protoreflect.EnumKind:
			if _, ok := v.seen[string(field.Enum.Desc.FullName())]; !ok {
				v.seen[string(field.Enum.Desc.FullName())] = struct{}{}
				enumPkgName := field.Enum.Desc.ParentFile().Package()
				if _, ok := v.refs[string(enumPkgName)]; !ok {
					v.refs[string(enumPkgName)] = []string{}
				}
				v.refs[string(enumPkgName)] = append(v.refs[string(enumPkgName)], string(field.Enum.Desc.FullName()))
			}
		case protoreflect.MessageKind:
			v.walkMessage(field.Message)
		}
	}
}
