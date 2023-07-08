package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"

	examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	logger "go.temporal.io/server/common/log"
)

// Workflows manages shared state for workflow constructors
type Workflows struct{}

// ============================================================================

// CreateFooWorkflow creates a new Foo resource
type CreateFooWorkflow struct {
	*examplev1.CreateFooInput
	progress float32
	status   examplev1.Foo_Status
}

// CreateFoo initializes a new CreateFooWorkflow
func (w *Workflows) CreateFoo(ctx workflow.Context, input *examplev1.CreateFooInput) (examplev1.CreateFooWorkflow, error) {
	return &CreateFooWorkflow{input, 0, examplev1.Foo_FOO_STATUS_UNSPECIFIED}, nil
}

// Execute defines the entrypoint to a CreateFooWorkflow
func (wf *CreateFooWorkflow) Execute(ctx workflow.Context) (*examplev1.CreateFooResponse, error) {
	// execute Notify activity using generated helper
	if err := examplev1.Notify(ctx, nil, &examplev1.NotifyRequest{Message: fmt.Sprintf("creating foo resource (%s)", wf.Req.GetName())}).Get(ctx); err != nil {
		return nil, fmt.Errorf("error sending notification: %w", err)
	}

	// wait until signalled progress reaches 100
	for progress := float32(0); progress < 100; {
		signal, _ := wf.SetFooProgress.Receive(ctx)
		progress = signal.GetProgress()
		switch {
		case progress < 0:
			progress, wf.status = 0, examplev1.Foo_FOO_STATUS_UNSPECIFIED
		case progress < 100:
			wf.status = examplev1.Foo_FOO_STATUS_CREATING
		case progress >= 100:
			progress, wf.status = 100, examplev1.Foo_FOO_STATUS_READY
		}
		wf.progress = progress
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
	// initialize cli app
	app := &cli.App{
		Name:  "Example",
		Usage: "an example temporal cli",
		// add cleanup logic to parent app
		After: func(cmd *cli.Context) error {
			if c, ok := cmd.App.Metadata["client"]; ok {
				c.(client.Client).Close()
			}
			return nil
		},
	}

	// initialize client commands using generated constructor
	var err error
	if app.Commands, err = examplev1.NewCommands(
		// provide a client initializer for use by commands
		examplev1.WithClientForCommand(func(cmd *cli.Context) (client.Client, error) {
			c, err := client.Dial(client.Options{
				Logger: logger.NewSdkLogger(logger.NewCLILogger()),
			})
			if err != nil {
				return nil, fmt.Errorf("error initializing client: %w", err)
			}
			// set a reference to the client in app metadata for use by app cleanup
			cmd.App.Metadata["client"] = c
			return c, nil
		}),
	); err != nil {
		log.Fatalf("error initializing commands: %v", err)
	}

	// add worker command
	app.Commands = append(app.Commands, &cli.Command{
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
	sort.Slice(app.Commands, func(i, j int) bool {
		return app.Commands[i].Name < app.Commands[j].Name
	})

	// run cli
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
