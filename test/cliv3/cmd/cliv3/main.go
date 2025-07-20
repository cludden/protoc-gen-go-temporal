package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/cludden/protoc-gen-go-temporal/gen/test/cliv3"
	example "github.com/cludden/protoc-gen-go-temporal/test/cliv3"
	"github.com/urfave/cli/v3"
	"go.temporal.io/sdk/client"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
)

func main() {
	cmd, err := cliv3.NewExampleServiceCli(
		cliv3.NewExampleServiceCliOptions().
			WithClient(func(ctx context.Context, c *cli.Command) (client.Client, error) {
				return client.DialContext(ctx, client.Options{
					Logger: tlog.NewStructuredLogger(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
						Level: slog.LevelDebug,
					}))),
				})
			}).
			WithWorker(func(ctx context.Context, cmd *cli.Command, c client.Client) (worker.Worker, error) {
				w := worker.New(c, cliv3.ExampleServiceTaskQueue, worker.Options{})
				example.Register(w)
				return w, nil
			}),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
