package main

import (
	"context"

	simplepb "github.com/cludden/protoc-gen-go-temporal/gen/test/simple/v1"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

type (
	DeprecatedWorkflows struct{}

	DeprecatedActivities struct{}
)

func RegisterDeprecated(r worker.Registry) {
	simplepb.RegisterDeprecatedWorkflows(r, &DeprecatedWorkflows{})
	simplepb.RegisterDeprecatedActivities(r, &DeprecatedActivities{})
}

// ============================================================================

type SomeDeprecatedWorkflow1 struct {
	*simplepb.SomeDeprecatedWorkflow1WorkflowInput
}

func (w *DeprecatedWorkflows) SomeDeprecatedWorkflow1(ctx workflow.Context, input *simplepb.SomeDeprecatedWorkflow1WorkflowInput) (simplepb.SomeDeprecatedWorkflow1Workflow, error) {
	return &SomeDeprecatedWorkflow1{input}, nil
}

func (w *SomeDeprecatedWorkflow1) Execute(ctx workflow.Context) (*simplepb.SomeDeprecatedMessage, error) {
	return nil, nil
}

func (w *SomeDeprecatedWorkflow1) SomeDeprecatedQuery1(*simplepb.SomeDeprecatedMessage) (*simplepb.SomeDeprecatedMessage, error) {
	return nil, nil
}

func (w *SomeDeprecatedWorkflow1) SomeDeprecatedUpdate1(workflow.Context, *simplepb.SomeDeprecatedMessage) (*simplepb.SomeDeprecatedMessage, error) {
	return nil, nil
}

// ============================================================================

type SomeDeprecatedWorkflow2 struct {
	*simplepb.SomeDeprecatedWorkflow2WorkflowInput
}

func (w *DeprecatedWorkflows) SomeDeprecatedWorkflow2(ctx workflow.Context, input *simplepb.SomeDeprecatedWorkflow2WorkflowInput) (simplepb.SomeDeprecatedWorkflow2Workflow, error) {
	return &SomeDeprecatedWorkflow2{input}, nil
}

func (w *SomeDeprecatedWorkflow2) Execute(ctx workflow.Context) (*simplepb.SomeDeprecatedMessage, error) {
	return nil, nil
}

func (w *SomeDeprecatedWorkflow2) SomeDeprecatedQuery2(*simplepb.SomeDeprecatedMessage) (*simplepb.SomeDeprecatedMessage, error) {
	return nil, nil
}

func (w *SomeDeprecatedWorkflow2) SomeDeprecatedUpdate2(workflow.Context, *simplepb.SomeDeprecatedMessage) (*simplepb.SomeDeprecatedMessage, error) {
	return nil, nil
}

// ============================================================================

func (a *DeprecatedActivities) SomeDeprecatedActivity1(context.Context, *simplepb.SomeDeprecatedMessage) (*simplepb.SomeDeprecatedMessage, error) {
	return nil, nil
}

func (a *DeprecatedActivities) SomeDeprecatedActivity2(context.Context, *simplepb.SomeDeprecatedMessage) (*simplepb.SomeDeprecatedMessage, error) {
	return nil, nil
}
