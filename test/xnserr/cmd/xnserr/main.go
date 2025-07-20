package main

import (
	"log"
	"os"

	xnserrv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/xnserr/v1"
	"github.com/cludden/protoc-gen-go-temporal/gen/test/xnserr/v1/xnserrv1xns"
	"github.com/cludden/protoc-gen-go-temporal/test/xnserr"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	server, err := xnserrv1.NewServerCliCommand(
		xnserrv1.NewServerCliOptions().
			WithClient(newClientForNamespace("server")).
			WithWorker(func(ctx *cli.Context, c client.Client) (worker.Worker, error) {
				w := worker.New(c, xnserrv1.ServerTaskQueue, worker.Options{})
				xnserrv1.RegisterServerWorkflows(w, &xnserr.ServerWorkflows{})
				return w, nil
			}),
	)
	if err != nil {
		log.Fatalf("failed to create server command: %v", err)
	}

	client, err := xnserrv1.NewClientCliCommand(
		xnserrv1.NewClientCliOptions().
			WithClient(newClientForNamespace("default")).
			WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
				sc, err := client.NewClientFromExistingWithContext(cmd.Context, c, client.Options{
					Namespace: "server",
				})
				if err != nil {
					return nil, err
				}

				w := worker.New(c, xnserrv1.ClientTaskQueue, worker.Options{})
				xnserrv1.RegisterClientWorkflows(w, &xnserr.ClientWorkflows{})
				xnserrv1xns.RegisterServerActivities(w, xnserrv1.NewServerClient(sc))
				return w, nil
			}),
	)
	if err != nil {
		log.Fatalf("failed to create client command: %v", err)
	}

	app := &cli.App{
		Name: "xnserr",
		Commands: []*cli.Command{
			client,
			server,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}

func newClientForNamespace(ns string) func(cmd *cli.Context) (client.Client, error) {
	return func(cmd *cli.Context) (client.Client, error) {
		return client.DialContext(cmd.Context, client.Options{
			Namespace: ns,
		})
	}
}
