package echo

import (
	nexusv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/nexus/v1"
	nexusv1temporal "github.com/cludden/protoc-gen-go-temporal/gen/example/nexus/v1/nexusv1temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

type (
	workflows struct {
		greeting nexusv1temporal.GreetingServiceNexusClient
	}

	echoWorkflow struct {
		*workflows
		*nexusv1.EchoWorkflowInput
	}
)

func Register(r worker.WorkflowRegistry, greeting nexusv1temporal.GreetingServiceNexusClient) error {
	w := &workflows{greeting: greeting}
	nexusv1.RegisterEchoServiceWorkflows(r, w)
	return nil
}

func (w *workflows) Echo(ctx workflow.Context, input *nexusv1.EchoWorkflowInput) (nexusv1.EchoWorkflow, error) {
	return &echoWorkflow{w, input}, nil
}

func (w *echoWorkflow) Execute(ctx workflow.Context) (*nexusv1.EchoOutput, error) {
	hello, err := w.greeting.Hello(ctx, &nexusv1.HelloInput{
		Name:     w.Req.GetName(),
		Language: w.Req.GetLanguage(),
	})
	if err != nil {
		return nil, err
	}
	return &nexusv1.EchoOutput{
		Message: hello.GetMessage(),
	}, nil
}
