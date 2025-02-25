package plugin

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func (svc *Manifest) toCamel(format string, args ...any) string {
	for i, arg := range args {
		if fn, ok := arg.(protoreflect.FullName); ok {
			args[i] = svc.methods[fn].GoName
		}
	}

	s := fmt.Sprintf(format, args...)
	return svc.caser.ToCamel(s)
}

func (svc *Manifest) toLowerCamel(format string, args ...any) string {
	for i, arg := range args {
		if fn, ok := arg.(protoreflect.FullName); ok {
			args[i] = svc.methods[fn].GoName
		}
	}

	s := fmt.Sprintf(format, args...)
	return svc.caser.ToLowerCamel(s)
}
