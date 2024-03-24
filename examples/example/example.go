package example

import (
	"context"
	"fmt"

	examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/workflow"
)

type (
	// Workflows manages shared state for workflow constructors and is used to
	// register workflows with a worker
	Workflows struct{}

	// Activities manages shared state for activities and is used to register
	// activities with a worker
	Activities struct{}
)

// CreateFooWorkflow manages workflow state for a CreateFoo workflow
type CreateFooWorkflow struct {
	// it embeds the generated workflow Input type that contains the workflow
	// input and signal helpers
	*examplev1.CreateFooWorkflowInput

	log      log.Logger
	progress float32
	status   examplev1.Foo_Status
}

// CreateFoo implements a CreateFoo workflow constructor on the shared Workflows struct
// that initializes a new CreateFooWorkflow for each execution
func (w *Workflows) CreateFoo(ctx workflow.Context, input *examplev1.CreateFooWorkflowInput) (examplev1.CreateFooWorkflow, error) {
	return &CreateFooWorkflow{
		CreateFooWorkflowInput: input,
		log:                    workflow.GetLogger(ctx),
		status:                 examplev1.Foo_FOO_STATUS_CREATING,
	}, nil
}

// Execute defines the entrypoint to a CreateFooWorkflow
func (wf *CreateFooWorkflow) Execute(ctx workflow.Context) (*examplev1.CreateFooResponse, error) {
	// listen for signals
	workflow.Go(ctx, func(ctx workflow.Context) {
		for {
			signal, _ := wf.SetFooProgress.Receive(ctx)
			wf.UpdateFooProgress(ctx, signal)
		}
	})

	// execute Notify activity using generated helper
	if err := examplev1.Notify(ctx, &examplev1.NotifyRequest{
		Message: fmt.Sprintf("creating foo resource (%s)", wf.Req.GetName()),
	}); err != nil {
		return nil, fmt.Errorf("error sending notification: %w", err)
	}

	// block until progress has reached 100 via signals and/or updates
	if err := workflow.Await(ctx, func() bool {
		return wf.status == examplev1.Foo_FOO_STATUS_READY
	}); err != nil {
		return nil, fmt.Errorf("error awaiting ready status: %w", err)
	}

	return &examplev1.CreateFooResponse{
		Foo: &examplev1.Foo{
			Name:   wf.Req.GetName(),
			Status: wf.status,
		},
	}, nil
}

// GetFooProgress defines the handler for a GetFooProgress query
func (wf *CreateFooWorkflow) GetFooProgress() (*examplev1.GetFooProgressResponse, error) {
	return &examplev1.GetFooProgressResponse{Progress: wf.progress, Status: wf.status}, nil
}

// UpdateFooProgress defines the handler for an UpdateFooProgress update
func (wf *CreateFooWorkflow) UpdateFooProgress(ctx workflow.Context, req *examplev1.SetFooProgressRequest) (*examplev1.GetFooProgressResponse, error) {
	wf.progress = req.GetProgress()
	switch {
	case wf.progress < 0:
		wf.progress, wf.status = 0, examplev1.Foo_FOO_STATUS_CREATING
	case wf.progress < 100:
		wf.status = examplev1.Foo_FOO_STATUS_CREATING
	case wf.progress >= 100:
		wf.progress, wf.status = 100, examplev1.Foo_FOO_STATUS_READY
	}
	return &examplev1.GetFooProgressResponse{Progress: wf.progress, Status: wf.status}, nil
}

// Notify defines the implementation for a Notify activity
func (a *Activities) Notify(ctx context.Context, req *examplev1.NotifyRequest) error {
	activity.GetLogger(ctx).Info("notification", "message", req.GetMessage())
	return nil
}
