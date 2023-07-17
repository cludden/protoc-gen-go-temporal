package plugin

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func toCamel(format string, a ...any) string {
	return strcase.ToCamel(fmt.Sprintf(format, a...))
}

func toLowerCamel(format string, a ...any) string {
	return strcase.ToLowerCamel(fmt.Sprintf(format, a...))
}
