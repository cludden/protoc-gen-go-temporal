package greeting

import (
	"context"
	"fmt"

	nexusv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/nexus/v1"
	"github.com/cludden/protoc-gen-go-temporal/gen/example/nexus/v1/nexusv1temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

type (
	workflows struct{}

	activities struct{}

	helloWorkflow struct {
		*workflows
		*nexusv1.HelloWorkflowInput
	}
)

func Register(r worker.Registry) error {
	nexusv1.RegisterGreetingServiceWorkflows(r, &workflows{})
	nexusv1.RegisterGreetingServiceActivities(r, &activities{})
	return nexusv1temporal.RegisterGreetingServiceNexusService(r)
}

func (w *workflows) Hello(ctx workflow.Context, input *nexusv1.HelloWorkflowInput) (nexusv1.HelloWorkflow, error) {
	return &helloWorkflow{w, input}, nil
}

func (w *helloWorkflow) Execute(ctx workflow.Context) (*nexusv1.HelloOutput, error) {
	return nexusv1.Hello(ctx, w.Req)
}

func (a *activities) Hello(ctx context.Context, input *nexusv1.HelloInput) (*nexusv1.HelloOutput, error) {
	switch input.Language {
	case nexusv1.Language_LANGUAGE_ENGLISH:
		return &nexusv1.HelloOutput{Message: "Hello " + input.Name + " 👋"}, nil
	case nexusv1.Language_LANGUAGE_FRENCH:
		return &nexusv1.HelloOutput{Message: "Bonjour " + input.Name + " 👋"}, nil
	case nexusv1.Language_LANGUAGE_GERMAN:
		return &nexusv1.HelloOutput{Message: "Hallo " + input.Name + " 👋"}, nil
	case nexusv1.Language_LANGUAGE_SPANISH:
		return &nexusv1.HelloOutput{Message: "¡Hola! " + input.Name + " 👋"}, nil
	case nexusv1.Language_LANGUAGE_TURKISH:
		return &nexusv1.HelloOutput{Message: "Merhaba " + input.Name + " 👋"}, nil
	}
	return nil, fmt.Errorf("unsupported language %q", input.Language)
}
