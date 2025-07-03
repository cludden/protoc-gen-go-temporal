package plugin

import "google.golang.org/protobuf/reflect/protoreflect"

type names struct {
	*Manifest
}

func (m *Manifest) Names() *names {
	return &names{m}
}

func (n *names) updateName(update protoreflect.FullName) string {
	return n.toCamel("%sUpdateName", update)
}

func (n *names) workflowName(workflow protoreflect.FullName) string {
	return n.toCamel("%sWorkflowName", workflow)
}
