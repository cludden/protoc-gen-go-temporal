package cliv3

import (
	"github.com/cludden/protoc-gen-go-temporal/gen/test/cliv3"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

type (
	Workflows struct{}

	CreateFooWorkflow struct {
		*Workflows
		*cliv3.CreateFooWorkflowInput
	}
)

func Register(r worker.Registry) {
	cliv3.RegisterExampleServiceWorkflows(r, &Workflows{})
}

func (w *Workflows) CreateFoo(ctx workflow.Context, input *cliv3.CreateFooWorkflowInput) (cliv3.CreateFooWorkflow, error) {
	return &CreateFooWorkflow{w, input}, nil
}

func (w *CreateFooWorkflow) Execute(ctx workflow.Context) (*cliv3.CreateFooOutput, error) {
	return &cliv3.CreateFooOutput{}, nil
}

func (w *CreateFooWorkflow) GetFoo(input *cliv3.GetFooInput) (*cliv3.GetFooOutput, error) {
	return &cliv3.GetFooOutput{}, nil
}

func (w *CreateFooWorkflow) UpdateFoo(ctx workflow.Context, input *cliv3.UpdateFooInput) (*cliv3.UpdateFooOutput, error) {
	return &cliv3.UpdateFooOutput{}, nil
}
