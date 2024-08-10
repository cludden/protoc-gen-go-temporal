package main

import (
	"fmt"
	"log"
	"os"
	"time"

	schedulev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/schedule/v1"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type (
	Workflows struct{}

	ScheduleWorkflow struct {
		*schedulev1.ScheduleWorkflowInput
	}
)

func main() {
	app, err := schedulev1.NewExampleCli(
		schedulev1.NewExampleCliOptions().
			WithClient(withClient).
			WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
				w := worker.New(c, schedulev1.ExampleTaskQueue, worker.Options{})
				schedulev1.RegisterExampleWorkflows(w, &Workflows{})
				return w, nil
			}),
	)
	if err != nil {
		log.Fatal(err)
	}

	app.Commands = append(app.Commands, &cli.Command{
		Name: "create-schedule",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "id",
				Usage: "schedule id",
				Value: "schedule_id",
			},
		},
		Action: func(cmd *cli.Context) error {
			c, err := withClient(cmd)
			if err != nil {
				return err
			}
			defer c.Close()

			h, err := c.ScheduleClient().Create(cmd.Context, client.ScheduleOptions{
				ID: cmd.String("id"),
				Spec: client.ScheduleSpec{
					Intervals: []client.ScheduleIntervalSpec{{
						Every: time.Minute,
					}},
				},
				Action: &client.ScheduleWorkflowAction{
					ID:        fmt.Sprintf("%s/", schedulev1.ScheduleWorkflowName),
					Workflow:  schedulev1.ScheduleWorkflowName,
					TaskQueue: schedulev1.ExampleTaskQueue,
					Args: []any{
						&schedulev1.ScheduleInput{},
					},
				},
			})
			if err != nil {
				return err
			}
			fmt.Printf("schedule created: %s\n", h.GetID())
			return nil
		},
	})

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func (w *Workflows) Schedule(ctx workflow.Context, input *schedulev1.ScheduleWorkflowInput) (schedulev1.ScheduleWorkflow, error) {
	return &ScheduleWorkflow{input}, nil
}

func (w *ScheduleWorkflow) Execute(ctx workflow.Context) (*schedulev1.ScheduleOutput, error) {
	return &schedulev1.ScheduleOutput{
		StartedAt: timestamppb.New(workflow.Now(ctx)),
	}, nil
}

func withClient(cmd *cli.Context) (client.Client, error) {
	c, err := client.DialContext(cmd.Context, client.Options{})
	if err != nil {
		return nil, err
	}
	return c, nil
}
