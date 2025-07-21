package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/cludden/protoc-gen-go-temporal/examples/nexus/echo"
	"github.com/cludden/protoc-gen-go-temporal/examples/nexus/greeting"
	nexusv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/nexus/v1"
	"github.com/cludden/protoc-gen-go-temporal/gen/example/nexus/v1/nexusv1nexustemporal"
	"github.com/urfave/cli/v3"
	"go.temporal.io/api/nexus/v1"
	"go.temporal.io/api/operatorservice/v1"
	"go.temporal.io/sdk/client"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
)

func main() {
	greetingCmd, err := nexusv1.NewGreetingServiceCli(
		nexusv1.NewGreetingServiceCliOptions().
			WithClient(newClient).
			WithWorker(func(ctx context.Context, cmd *cli.Command, c client.Client) (worker.Worker, error) {
				w := worker.New(c, nexusv1.GreetingServiceTaskQueue, worker.Options{})
				if err := greeting.Register(w); err != nil {
					return nil, err
				}
				return w, nil
			}),
	)
	if err != nil {
		log.Fatal(err)
	}

	echoCmd, err := nexusv1.NewEchoServiceCli(
		nexusv1.NewEchoServiceCliOptions().
			WithClient(newClient).
			WithWorker(func(ctx context.Context, cmd *cli.Command, c client.Client) (worker.Worker, error) {
				w := worker.New(c, nexusv1.EchoServiceTaskQueue, worker.Options{})
				endpoint := cmd.String("endpoint")
				greeting := nexusv1nexustemporal.NewGreetingServiceNexusClient(endpoint)
				if err := echo.Register(w, greeting); err != nil {
					return nil, err
				}
				return w, nil
			}),
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := &cli.Command{
		Name:  "nexus",
		Usage: "Run the Nexus example",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "endpoint",
				Usage: "Nexus endpoint to connect to",
				Value: "greeting",
			},
		},
		Commands: []*cli.Command{
			greetingCmd,
			echoCmd,
			{
				Name: "register-endpoint",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					c, err := newClient(ctx, cmd)
					if err != nil {
						return err
					}
					defer c.Close()

					_, err = c.OperatorService().
						CreateNexusEndpoint(ctx, &operatorservice.CreateNexusEndpointRequest{
							Spec: &nexus.EndpointSpec{
								Name: cmd.String("endpoint"),
								Target: &nexus.EndpointTarget{
									Variant: &nexus.EndpointTarget_Worker_{
										Worker: &nexus.EndpointTarget_Worker{
											Namespace: "default",
											TaskQueue: nexusv1.GreetingServiceTaskQueue,
										},
									},
								},
							},
						})
					return err
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func newClient(ctx context.Context, cmd *cli.Command) (client.Client, error) {
	return client.DialContext(ctx, client.Options{
		Namespace: "default",
		Logger: tlog.NewStructuredLogger(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))),
	})
}
