package main

import (
	"log"
	"os"

	example "github.com/cludden/protoc-gen-go-temporal/example"
	examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	app, err := examplev1.NewExampleCli(
		examplev1.NewExampleCliOptions().
			WithClient(func(cmd *cli.Context) (client.Client, error) {
				return client.Dial(client.Options{})
			}).
			WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
				w := worker.New(c, examplev1.ExampleTaskQueue, worker.Options{})
				examplev1.RegisterExampleWorkflows(w, &example.Workflows{})
				examplev1.RegisterExampleActivities(w, &example.Activities{})
				return w, nil
			}),
	)
	if err != nil {
		log.Fatalf("error initializing cli: %v", err)
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
