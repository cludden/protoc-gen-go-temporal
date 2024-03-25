package main

import (
	"context"
	"log"
	"os"

	helloworldv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/helloworld/v1"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	sdklog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

type (
	// HelloWorldWorkflow provides a helloworldv1.HelloWorldWorkflow implementation
	HelloWorldWorkflow struct {
		*helloworldv1.HelloWorldWorkflowInput
		log sdklog.Logger
	}

	// Activities provides a helloworldv1.HelloWorldActivities implementation
	Activities struct{}
)

// NewHelloWorldWorkflow initializes a new helloworldv1.HelloWorldWorkflow value
func NewHelloWorldWorkflow(ctx workflow.Context, input *helloworldv1.HelloWorldWorkflowInput) (helloworldv1.HelloWorldWorkflow, error) {
	return &HelloWorldWorkflow{input, workflow.GetLogger(ctx)}, nil
}

// Execute defines the entrypoint to a HelloWorld workflow
func (w *HelloWorldWorkflow) Execute(ctx workflow.Context) (*helloworldv1.HelloWorldOutput, error) {
	result, err := helloworldv1.HelloWorld(ctx, w.Req)
	if err != nil {
		w.log.Error("Activity failed.", "error", err)
		return nil, err
	}

	w.log.Info("HelloWorld workflow completed.", "result", result.GetResult())
	return result, nil
}

// HelloWorld defines the entrypoint to a HelloWorld activity
func (a *Activities) HelloWorld(ctx context.Context, input *helloworldv1.HelloWorldInput) (*helloworldv1.HelloWorldOutput, error) {
	activity.GetLogger(ctx).Info("Activity", "name", input.GetName())
	return &helloworldv1.HelloWorldOutput{
		Result: "Hello " + input.GetName() + "!",
	}, nil
}

func main() {
	app, err := helloworldv1.NewHelloWorldCli(
		helloworldv1.NewHelloWorldCliOptions().
			WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
				w := worker.New(c, helloworldv1.HelloWorldTaskQueue, worker.Options{})
				helloworldv1.RegisterHelloWorldWorkflow(w, NewHelloWorldWorkflow)
				helloworldv1.RegisterHelloWorldActivities(w, &Activities{})
				return w, nil
			}),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
