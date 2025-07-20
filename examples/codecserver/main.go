package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/cludden/protoc-gen-go-temporal/examples/example"
	examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	"github.com/cludden/protoc-gen-go-temporal/pkg/codec"
	"github.com/cludden/protoc-gen-go-temporal/pkg/scheme"
	"github.com/urfave/cli/v3"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
)

func main() {
	app, err := examplev1.NewExampleCli(
		examplev1.NewExampleCliOptions().
			WithClient(func(ctx context.Context, cmd *cli.Command) (client.Client, error) {
				return client.DialContext(ctx, client.Options{
					DataConverter: converter.NewCompositeDataConverter(
						converter.NewNilPayloadConverter(),
						converter.NewByteSlicePayloadConverter(),
						converter.NewProtoPayloadConverter(),
					),
					Logger: tlog.NewStructuredLogger(slog.Default()),
				})
			}).
			WithWorker(func(ctx context.Context, cmd *cli.Command, c client.Client) (worker.Worker, error) {
				w := worker.New(c, examplev1.ExampleTaskQueue, worker.Options{})
				examplev1.RegisterExampleActivities(w, &example.Activities{})
				examplev1.RegisterExampleWorkflows(w, &example.Workflows{})
				return w, nil
			}),
	)
	if err != nil {
		log.Fatalf("error initializing example cli: %v", err)
	}

	app.Commands = append(app.Commands, &cli.Command{
		Name:  "codec",
		Usage: "run remote codec server",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			handler := converter.NewPayloadCodecHTTPHandler(
				codec.NewProtoJSONCodec(
					scheme.New(
						examplev1.WithExampleSchemeTypes(),
					),
				),
			)

			srv := &http.Server{
				Addr:    "0.0.0.0:8080",
				Handler: handler,
			}

			go func() {
				sigChan := make(chan os.Signal, 1)
				signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
				<-sigChan

				if err := srv.Shutdown(context.Background()); err != nil {
					log.Fatalf("error shutting down server: %v", err)
				}
			}()

			if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("server error: %v", err)
			}
			return nil
		},
	})

	// run cli
	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
