package plugin

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Names struct {
	*Manifest
}

func (m *Manifest) Names() *Names {
	return &Names{m}
}

func (n *Names) clientImpl() string {
	return n.toCamel("%sClient", n.Service)
}

func (n *Names) clientUpdateHandleIface(update protoreflect.FullName) string {
	return n.toCamel("%sHandle", update)
}

func (n *Names) clientUpdateHandleImpl(update protoreflect.FullName) string {
	return n.toLowerCamel("%sHandle", update)
}

func (n *Names) clientUpdateWithStartMethod(workflow, update protoreflect.FullName) string {
	return n.toCamel("%sWith%s", workflow, update)
}

func (n *Names) clientUpdateWithStartMethodAsync(workflow, update protoreflect.FullName) string {
	return n.toCamel("%sWith%sAsync", workflow, update)
}

func (n *Names) clientUpdateWithStartOptionsCtor(workflow, update protoreflect.FullName) string {
	return n.toCamel("New%sWith%sOptions", workflow, update)
}

func (n *Names) clientUpdateWithStartOptionsImpl(workflow, update protoreflect.FullName) string {
	return n.toCamel("%sWith%sOptions", workflow, update)
}

func (n *Names) clientWorkflowRun(workflow protoreflect.FullName) string {
	return n.toCamel("%sRun", workflow)
}
