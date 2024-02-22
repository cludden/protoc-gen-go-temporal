package plugin

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (svc *Manifest) toCamel(format string, args ...any) string {
	for i, arg := range args {
		if fn, ok := arg.(protoreflect.FullName); ok {
			args[i] = svc.methods[fn].GoName
		}
	}
	return strcase.ToCamel(fmt.Sprintf(format, args...))
}

func (svc *Manifest) toLowerCamel(format string, args ...any) string {
	for i, arg := range args {
		if fn, ok := arg.(protoreflect.FullName); ok {
			args[i] = svc.methods[fn].GoName
		}
	}
	return strcase.ToLowerCamel(fmt.Sprintf(format, args...))
}
