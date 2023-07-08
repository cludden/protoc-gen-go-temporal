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
)

// Workflows manages shared state for workflow constructors
type Workflows struct{}

// ============================================================================

// CreateFooWorkflow creates a new Foo resource
type CreateFooWorkflow struct {
	*examplev1.CreateFooInput
	progress float32
	status   examplev1.Foo_Status
	update   workflow.Settable
	updated  workflow.Future
}

// CreateFoo initializes a new CreateFooWorkflow
func (w *Workflows) CreateFoo(ctx workflow.Context, input *examplev1.CreateFooInput) (examplev1.CreateFooWorkflow, error) {
	return &CreateFooWorkflow{input, 0, examplev1.Foo_FOO_STATUS_UNSPECIFIED, nil, nil}, nil
}

// Execute defines the entrypoint to a CreateFooWorkflow
func (wf *CreateFooWorkflow) Execute(ctx workflow.Context) (*examplev1.CreateFooResponse, error) {
	// execute Notify activity using generated helper
	if err := examplev1.Notify(ctx, &examplev1.NotifyRequest{Message: fmt.Sprintf("creating foo resource (%s)", wf.Req.GetName())}); err != nil {
		return nil, fmt.Errorf("error sending notification: %w", err)
	}

	// wait until signalled progress reaches 100
	for wf.progress = float32(0); wf.progress < 100; {
		wf.updated, wf.update = workflow.NewFuture(ctx)
		workflow.NewSelector(ctx).
			AddReceive(wf.SetFooProgress.Channel, func(workflow.ReceiveChannel, bool) {
				wf.UpdateFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: wf.SetFooProgress.ReceiveAsync().GetProgress()})
			}).
			AddFuture(wf.updated, func(workflow.Future) {}).
			Select(ctx)
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

// UpdateFooProgress defines the handler for a UpdateFooProgress update
func (wf *CreateFooWorkflow) UpdateFooProgress(ctx workflow.Context, req *examplev1.SetFooProgressRequest) (*examplev1.GetFooProgressResponse, error) {
	progress := req.GetProgress()
	switch {
	case progress < 0:
		progress, wf.status = 0, examplev1.Foo_FOO_STATUS_UNSPECIFIED
	case progress < 100:
		wf.status = examplev1.Foo_FOO_STATUS_CREATING
	case progress >= 100:
		progress, wf.status = 100, examplev1.Foo_FOO_STATUS_READY
	}
	wf.progress = progress
	wf.update.SetValue(progress)
	return &examplev1.GetFooProgressResponse{Progress: wf.progress, Status: wf.status}, nil
}

// ============================================================================

// Activities manages shared state for activities
type Activities struct{}

// Notify sends a notification
func (a *Activities) Notify(ctx context.Context, req *examplev1.NotifyRequest) error {
	activity.GetLogger(ctx).Info("notification", "message", req.GetMessage())
	return nil
}

// ============================================================================

func main() {
	// initialize client commands using generated constructor
	commands, err := examplev1.NewCommands(
		// provide a client initializer for use by commands
		examplev1.WithClientForCommand(func(cmd *cli.Context) (client.Client, error) {
			return client.Dial(client.Options{})
		}),
	)
	if err != nil {
		log.Fatalf("error initializing commands: %v", err)
	}

	// add worker command
	commands = append(commands, &cli.Command{
		Name:  "worker",
		Usage: "run service worker",
		Action: func(cmd *cli.Context) error {
			// initialize temporal client
			c, err := client.Dial(client.Options{})
			if err != nil {
				return fmt.Errorf("error initializing client: %w", err)
			}
			defer c.Close()

			// register workflows & activities using generated registration helpers, start worker
			w := worker.New(c, examplev1.ExampleTaskQueue, worker.Options{})
			examplev1.RegisterActivities(w, &Activities{})
			examplev1.RegisterWorkflows(w, &Workflows{})
			if err := w.Start(); err != nil {
				return fmt.Errorf("error starting worker: %w", err)
			}
			defer w.Stop()

			<-cmd.Context.Done()
			return nil
		},
	})

	// run cli
	if err := (&cli.App{
		Name:     "Example",
		Usage:    "an example temporal cli",
		Commands: commands,
	}).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
