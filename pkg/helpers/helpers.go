package helpers

import "go.temporal.io/sdk/workflow"

// Initializable describes a Workflow that supports initialization prior to registering signal, query, or update handlers
type Initializable interface {
	Initialize(ctx workflow.Context) error
}
