package main

import (
	"context"
	"log"
	"os"

	acronymv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/acronym/v1"
	"go.temporal.io/sdk/workflow"
)

func main() {
	app, err := acronymv1.NewAWSCli()
	if err != nil {
		log.Fatal(err)
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type (
	Workflows struct{}

	Activities struct{}
)

// =============================================================================

type ManageAWSWorkflow struct {
	*acronymv1.ManageAWSWorkflowInput
}

func (w *Workflows) ManageAWS(ctx workflow.Context, input *acronymv1.ManageAWSWorkflowInput) (acronymv1.ManageAWSWorkflow, error) {
	return &ManageAWSWorkflow{input}, nil
}

func (w *ManageAWSWorkflow) Execute(ctx workflow.Context) (*acronymv1.ManageAWSResponse, error) {
	return &acronymv1.ManageAWSResponse{}, nil
}

// =============================================================================

type ManageAWSResourceWorkflow struct {
	*acronymv1.ManageAWSResourceWorkflowInput
}

func (w *Workflows) ManageAWSResource(ctx workflow.Context, input *acronymv1.ManageAWSResourceWorkflowInput) (acronymv1.ManageAWSResourceWorkflow, error) {
	return &ManageAWSResourceWorkflow{input}, nil
}

func (w *ManageAWSResourceWorkflow) Execute(ctx workflow.Context) (*acronymv1.ManageAWSResourceResponse, error) {
	return acronymv1.ManageAWSResource(ctx, w.Req)
}

// =============================================================================

type SomethingV1FooBarWorkflow struct {
	*acronymv1.SomethingV1FooBarWorkflowInput
}

func (w *Workflows) SomethingV1FooBar(ctx workflow.Context, input *acronymv1.SomethingV1FooBarWorkflowInput) (acronymv1.SomethingV1FooBarWorkflow, error) {
	return &SomethingV1FooBarWorkflow{input}, nil
}

func (w *SomethingV1FooBarWorkflow) Execute(ctx workflow.Context) (*acronymv1.SomethingV1FooBarResponse, error) {
	return &acronymv1.SomethingV1FooBarResponse{}, nil
}

// =============================================================================

type SomethingV2FooBarWorkflow struct {
	*acronymv1.SomethingV2FooBarWorkflowInput
}

func (w *Workflows) SomethingV2FooBar(ctx workflow.Context, input *acronymv1.SomethingV2FooBarWorkflowInput) (acronymv1.SomethingV2FooBarWorkflow, error) {
	return &SomethingV2FooBarWorkflow{input}, nil
}

func (w *SomethingV2FooBarWorkflow) Execute(ctx workflow.Context) (*acronymv1.SomethingV2FooBarResponse, error) {
	return &acronymv1.SomethingV2FooBarResponse{}, nil
}

// =============================================================================

func (a *Activities) ManageAWSResource(ctx context.Context, input *acronymv1.ManageAWSResourceRequest) (*acronymv1.ManageAWSResourceResponse, error) {
	return &acronymv1.ManageAWSResourceResponse{}, nil
}

func (a *Activities) ManageAWSResourceURN(ctx context.Context, input *acronymv1.ManageAWSResourceURNRequest) (*acronymv1.ManageAWSResourceURNResponse, error) {
	return &acronymv1.ManageAWSResourceURNResponse{}, nil
}
