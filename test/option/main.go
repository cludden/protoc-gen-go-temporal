package main

import (
	"context"

	optionv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/option/v1"
	"go.temporal.io/sdk/workflow"
)

type (
	TestWorkflows struct{}

	TestActivities struct{}

	WorkflowWithInputWorkflow struct {
		*optionv1.WorkflowWithInputWorkflowInput
	}
)

func (w *TestWorkflows) WorkflowWithInput(ctx workflow.Context, input *optionv1.WorkflowWithInputWorkflowInput) (optionv1.WorkflowWithInputWorkflow, error) {
	return &WorkflowWithInputWorkflow{input}, nil
}

func (w *WorkflowWithInputWorkflow) Execute(ctx workflow.Context) error {
	return nil
}

func (w *WorkflowWithInputWorkflow) UpdateWithInput(ctx workflow.Context, input *optionv1.UpdateWithInputRequest) error {
	return nil
}

func (a *TestActivities) ActivityWithInput(ctx context.Context, input *optionv1.ActivityWithInputRequest) (*optionv1.ActivityWithInputResponse, error) {
	return nil, nil
}
