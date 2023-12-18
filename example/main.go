package main

import (
	"context"
	"fmt"
	"log"
	"os"

	examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	logger "go.temporal.io/server/common/log"
)

// Workflows manages shared state for workflow constructors and is used to
// register workflows with a worker
type Workflows struct{}

// ============================================================================

// CreateFooWorkflow manages workflow state for a CreateFoo workflow
type CreateFooWorkflow struct {
	// it embeds the generated workflow Input type that contains the workflow
	// input and signal helpers
	*examplev1.CreateFooWorkflowInput

	progress float32
	status   examplev1.Foo_Status
}

// CreateFoo implements a CreateFoo workflow constructor on the shared Workflows struct
// that initializes a new CreateFooWorkflow for each execution
func (w *Workflows) CreateFoo(ctx workflow.Context, input *examplev1.CreateFooWorkflowInput) (examplev1.CreateFooWorkflow, error) {
	return &CreateFooWorkflow{input, 0, examplev1.Foo_FOO_STATUS_CREATING}, nil
}

// Execute defines the entrypoint to a CreateFooWorkflow
func (wf *CreateFooWorkflow) Execute(ctx workflow.Context) (*examplev1.CreateFooResponse, error) {
	// listen for signals
	workflow.Go(ctx, func(ctx workflow.Context) {
		for {
			signal, _ := wf.SetFooProgress.Receive(ctx)
			wf.UpdateFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: signal.GetProgress()})
		}
	})

	// execute Notify activity using generated helper
	err := examplev1.Notify(ctx, &examplev1.NotifyRequest{
		Message: fmt.Sprintf("creating foo resource (%s)", wf.Req.GetRequestName()),
	})
	if err != nil {
		return nil, fmt.Errorf("error sending notification: %w", err)
	}

	// block until progress has reached 100 via signals and/or updates
	workflow.Await(ctx, func() bool {
		return wf.status == examplev1.Foo_FOO_STATUS_READY
	})

	return &examplev1.CreateFooResponse{
		Foo: &examplev1.Foo{
			Name:   wf.Req.GetRequestName(),
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

// ============================================================================

// Activities manages shared state for activities and is used to register
// activities with a worker
type Activities struct{}

// Notify defines the implementation for a Notify activity
func (a *Activities) Notify(ctx context.Context, req *examplev1.NotifyRequest) error {
	activity.GetLogger(ctx).Info("notification", "message", req.GetMessage())
	return nil
}

// ============================================================================

func main() {
	// initialize the generated cli application
	app, err := examplev1.NewExampleCli(
		examplev1.NewExampleCliOptions().
			WithClient(func(cmd *cli.Context) (client.Client, error) {
				return client.Dial(client.Options{
					Logger: logger.NewSdkLogger(logger.NewCLILogger()),
				})
			}).
			WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
				w := worker.New(c, examplev1.ExampleTaskQueue, worker.Options{})
				// register activities and workflows using generated helpers
				examplev1.RegisterExampleActivities(w, &Activities{})
				examplev1.RegisterExampleWorkflows(w, &Workflows{})
				return w, nil
			}),
	)
	if err != nil {
		log.Fatalf("error initializing commands: %v", err)
	}
	app.Name = "example"
	app.Usage = "an example temporal cli"

	// run cli
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
