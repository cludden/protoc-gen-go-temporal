package example

import (
	helloworldv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/helloworld/v1"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/workflow"
)

type (
	Workflows struct{}

	// HelloWorkflow provides a helloworldv1.HelloWorkflow implementation
	HelloWorkflow struct {
		*helloworldv1.HelloWorkflowInput
		log log.Logger
	}
)

// NewHelloWorkflow initializes a new helloworldv1.HelloWorkflow value
func (w *Workflows) Hello(ctx workflow.Context, input *helloworldv1.HelloWorkflowInput) (helloworldv1.HelloWorkflow, error) {
	return &HelloWorkflow{input, workflow.GetLogger(ctx)}, nil
}

// Execute defines the entrypoint to a Hello workflow
func (w *HelloWorkflow) Execute(ctx workflow.Context) (*helloworldv1.HelloResponse, error) {
	w.log.Info("Hello workflow started", "request", w.Req)

	goodbye, _ := w.Goodbye.Receive(ctx)
	w.log.Info("Goodbye received", "signal", goodbye)

	return &helloworldv1.HelloResponse{}, nil
}
