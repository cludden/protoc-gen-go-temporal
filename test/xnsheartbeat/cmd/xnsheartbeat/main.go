package main

import (
	"log"
	"os"

	xnsheartbeatv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/xnsheartbeat/v1"
	"github.com/cludden/protoc-gen-go-temporal/gen/test/xnsheartbeat/v1/xnsheartbeatv1xns"
	"github.com/cludden/protoc-gen-go-temporal/test/xnsheartbeat"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	cmd, err := xnsheartbeatv1.NewXnsHeartbeatServiceCli(
		xnsheartbeatv1.NewXnsHeartbeatServiceCliOptions().
			WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
				w := worker.New(c, xnsheartbeatv1.XnsHeartbeatServiceTaskQueue, worker.Options{
					DefaultHeartbeatThrottleInterval: 0,
					MaxHeartbeatThrottleInterval:     0,
				})
				xnsheartbeatv1.RegisterXnsHeartbeatServiceWorkflows(w, &xnsheartbeat.Workflows{})
				return w, nil
			}),
	)
	if err != nil {
		log.Fatal(err)
	}

	callerCmd, err := xnsheartbeatv1.NewXnsHeartbeatCallerServiceCliCommand(
		xnsheartbeatv1.NewXnsHeartbeatCallerServiceCliOptions().
			WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
				w := worker.New(c, xnsheartbeatv1.XnsHeartbeatCallerServiceTaskQueue, worker.Options{
					DefaultHeartbeatThrottleInterval: 0,
					MaxHeartbeatThrottleInterval:     0,
				})
				xnsheartbeatv1.RegisterXnsHeartbeatCallerServiceWorkflows(w, &xnsheartbeat.CallerWorkflows{})
				cc := xnsheartbeatv1.NewXnsHeartbeatServiceClient(c)
				xnsheartbeatv1xns.RegisterXnsHeartbeatServiceActivities(w, cc)
				return w, nil
			}),
	)
	if err != nil {
		log.Fatal(err)
	}
	cmd.Commands = append(cmd.Commands, callerCmd)

	if err := cmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
