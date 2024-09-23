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

func (svc *Manifest) getNexusWorkflowOperationFutureIfaceName(workflow protoreflect.FullName) string {
	return svc.toCamel("%sWorkflowOperationFuture", workflow)
}

func (svc *Manifest) getNexusWorkflowOperationFutureImplName(workflow protoreflect.FullName) string {
	return svc.toLowerCamel("%sWorkflowOperationFuture", workflow)
}

func (svc *Manifest) getNexusWorkflowOperationName(workflow protoreflect.FullName) string {
	return svc.toCamel("%sWorkflowOperation", workflow)
}

func (svc *Manifest) getNexusWorkflowOperationNameConstName(workflow protoreflect.FullName) string {
	return svc.toCamel("%sWorkflowOperatioName", workflow)
}
