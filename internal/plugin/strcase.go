package plugin

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func (m *Manifest) toCamel(format string, args ...any) string {
	for i, arg := range args {
		if fn, ok := arg.(protoreflect.FullName); ok {
			args[i] = m.methods[fn].GoName
		}
	}

	s := fmt.Sprintf(format, args...)
	return m.caser.ToCamel(s)
}

func (m *Manifest) toLowerCamel(format string, args ...any) string {
	for i, arg := range args {
		if fn, ok := arg.(protoreflect.FullName); ok {
			args[i] = m.methods[fn].GoName
		}
	}

	s := fmt.Sprintf(format, args...)
	return m.caser.ToLowerCamel(s)
}
