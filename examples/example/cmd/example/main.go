package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/cludden/protoc-gen-go-temporal/examples/example"
	examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	"github.com/urfave/cli/v3"
	"go.temporal.io/sdk/client"
	logsdk "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
)

func main() {
	// initialize the generated cli application
	app, err := examplev1.NewExampleCli(
		examplev1.NewExampleCliOptions().
			WithClient(func(ctx context.Context, cmd *cli.Command) (client.Client, error) {
				return client.DialContext(ctx, client.Options{
					Logger: logsdk.NewStructuredLogger(slog.New(slog.NewTextHandler(os.Stdout, nil))),
				})
			}).
			WithWorker(func(ctx context.Context, cmd *cli.Command, c client.Client) (worker.Worker, error) {
				// register activities and workflows using generated helpers
				w := worker.New(c, examplev1.ExampleTaskQueue, worker.Options{})
				examplev1.RegisterExampleActivities(w, &example.Activities{})
				examplev1.RegisterExampleWorkflows(w, &example.Workflows{})
				return w, nil
			}),
	)
	if err != nil {
		log.Fatalf("error initializing example cli: %v", err)
	}

	// run cli
	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
