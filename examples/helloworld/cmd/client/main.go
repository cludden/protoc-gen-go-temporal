package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	helloworldv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/helloworld/v1"
	"go.temporal.io/sdk/client"
	sdklog "go.temporal.io/sdk/log"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	client, err := client.Dial(client.Options{
		Logger: sdklog.NewStructuredLogger(logger),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	example := helloworldv1.NewExampleClient(client)
	run, err := example.HelloAsync(ctx, &helloworldv1.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatal(err)
	}
	logger = logger.With("workflow_id", run.ID())
	logger.Info("workflow started")

	_, ctx = <-ctx.Done(), context.Background()
	logger.Info("received shutdown signal, sending Goodbye signal to workflow")
	if err := run.Goodbye(ctx, &helloworldv1.GoodbyeRequest{}); err != nil {
		log.Fatal(err)
	}

	out, err := run.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("workflow completed", "result", out.String())
}
