package main

import (
	"context"
	"fmt"

	xnsv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/xns/v1"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

type ExampleWorkflows struct{}

type CreateFooWorkflow struct {
	*xnsv1.CreateFooWorkflowInput
	progress float32
	status   xnsv1.Foo_Status
}

func (w *ExampleWorkflows) CreateFoo(ctx workflow.Context, input *xnsv1.CreateFooWorkflowInput) (xnsv1.CreateFooWorkflow, error) {
	return &CreateFooWorkflow{input, 0, xnsv1.Foo_FOO_STATUS_CREATING}, nil
}

func (wf *CreateFooWorkflow) Execute(ctx workflow.Context) (*xnsv1.CreateFooResponse, error) {
	workflow.Go(ctx, func(ctx workflow.Context) {
		for {
			signal, _ := wf.SetFooProgress.Receive(ctx)
			wf.UpdateFooProgress(ctx, signal)
		}
	})

	err := xnsv1.Notify(ctx, &xnsv1.NotifyRequest{
		Message: fmt.Sprintf("creating foo resource (%s)", wf.Req.GetName()),
	})
	if err != nil {
		return nil, fmt.Errorf("error sending notification: %w", err)
	}

	workflow.Await(ctx, func() bool {
		return wf.status == xnsv1.Foo_FOO_STATUS_READY
	})

	return &xnsv1.CreateFooResponse{
		Foo: &xnsv1.Foo{
			Name:   wf.Req.GetName(),
			Status: wf.status,
		},
	}, nil
}

func (wf *CreateFooWorkflow) GetFooProgress() (*xnsv1.GetFooProgressResponse, error) {
	return &xnsv1.GetFooProgressResponse{Progress: wf.progress, Status: wf.status}, nil
}

func (wf *CreateFooWorkflow) UpdateFooProgress(ctx workflow.Context, req *xnsv1.SetFooProgressRequest) (*xnsv1.GetFooProgressResponse, error) {
	wf.progress = req.GetProgress()
	switch {
	case wf.progress < 0:
		wf.progress, wf.status = 0, xnsv1.Foo_FOO_STATUS_CREATING
	case wf.progress < 100:
		wf.status = xnsv1.Foo_FOO_STATUS_CREATING
	case wf.progress >= 100:
		wf.progress, wf.status = 100, xnsv1.Foo_FOO_STATUS_READY
	}
	return &xnsv1.GetFooProgressResponse{Progress: wf.progress, Status: wf.status}, nil
}

// ExampleActivities manages shared state for activities and is used to register
// activities with a worker
type ExampleActivities struct{}

// Notify defines the implementation for a Notify activity
func (a *ExampleActivities) Notify(ctx context.Context, req *xnsv1.NotifyRequest) error {
	activity.GetLogger(ctx).Info("notification", "message", req.GetMessage())
	return nil
}
