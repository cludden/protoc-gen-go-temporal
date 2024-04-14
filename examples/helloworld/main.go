package main

import (
	"log"
	"os"

	example "github.com/cludden/protoc-gen-go-temporal/examples/helloworld/internal"
	helloworldv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/helloworld/v1"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	app, err := helloworldv1.NewExampleCli(
		helloworldv1.NewExampleCliOptions().
			WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
				w := worker.New(c, helloworldv1.ExampleTaskQueue, worker.Options{})
				helloworldv1.RegisterExampleWorkflows(w, &example.Workflows{})
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
