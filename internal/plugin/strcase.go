package plugin

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (svc *Manifest) toCamel(format string, args ...any) string {
	for i, arg := range args {
		if fn, ok := arg.(protoreflect.FullName); ok {
			args[i] = svc.methods[fn].GoName
		}
	}

	s := fmt.Sprintf(format, args...)

	for _, initialism := range svc.cfg.InitialismsStrCase {
		if strings.Contains(s, initialism) {
			strcase.ConfigureAcronym(s, s)
		}
	}

	return strcase.ToCamel(s)
}

func (svc *Manifest) toLowerCamel(format string, args ...any) string {
	for i, arg := range args {
		if fn, ok := arg.(protoreflect.FullName); ok {
			args[i] = svc.methods[fn].GoName
		}
	}

	s := fmt.Sprintf(format, args...)

	for _, initialism := range svc.cfg.InitialismsStrCase {
		if strings.Contains(s, initialism) {
			strcase.ConfigureAcronym(s, s)
		}
	}

	return strcase.ToLowerCamel(fmt.Sprintf(format, args...))
}
