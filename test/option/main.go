package main

import (
	"context"

	optionv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/option/v1"
	"go.temporal.io/sdk/workflow"
)

type (
	Workflows struct{}

	Activities struct{}

	WorkflowWithInputWorkflow struct {
		*optionv1.WorkflowWithInputWorkflowInput
	}
)

func (w *Workflows) WorkflowWithInput(ctx workflow.Context, input *optionv1.WorkflowWithInputWorkflowInput) (optionv1.WorkflowWithInputWorkflow, error) {
	return &WorkflowWithInputWorkflow{input}, nil
}

func (w *WorkflowWithInputWorkflow) Execute(ctx workflow.Context) error {
	return nil
}

func (w *WorkflowWithInputWorkflow) UpdateWithInput(ctx workflow.Context, input *optionv1.UpdateWithInputRequest) error {
	return nil
}

func (a *Activities) ActivityWithInput(ctx context.Context, input *optionv1.ActivityWithInputRequest) (*optionv1.ActivityWithInputResponse, error) {
	return nil, nil
}
