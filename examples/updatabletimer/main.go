package main

import (
	"log"
	"os"

	updatabletimerv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/updatabletimer/v1"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UpdatableTimerWorkflow struct {
	*updatabletimerv1.UpdatableTimerWorkflowInput
	log        tlog.Logger
	wakeUpTime *timestamppb.Timestamp
}

func NewUpdatableTimerWorkflow(ctx workflow.Context, input *updatabletimerv1.UpdatableTimerWorkflowInput) (updatabletimerv1.UpdatableTimerWorkflow, error) {
	return &UpdatableTimerWorkflow{input, workflow.GetLogger(ctx), input.Req.GetInitialWakeUpTime()}, nil
}

func (w *UpdatableTimerWorkflow) Execute(ctx workflow.Context) error {
	var timerFired bool
	for !timerFired && ctx.Err() == nil {
		timerCtx, cancelTimer := workflow.WithCancel(ctx)
		timer := workflow.NewTimer(timerCtx, w.wakeUpTime.AsTime().Sub(workflow.Now(ctx)))
		w.log.Info("SleepUntil", "WakeUpTime", w.wakeUpTime)

		workflow.NewSelector(ctx).
			AddFuture(timer, func(f workflow.Future) {
				if err := f.Get(timerCtx, nil); err != nil {
					w.log.Info("Timer canceled")
				} else {
					w.log.Info("Timer fired")
					timerFired = true
				}
			}).
			AddReceive(w.UpdateWakeUpTime.Channel, func(workflow.ReceiveChannel, bool) {
				defer cancelTimer()
				w.wakeUpTime = w.UpdateWakeUpTime.ReceiveAsync().GetWakeUpTime()
				w.log.Info("WakeUpTime updated", "WakeUpTime", w.wakeUpTime)
			}).
			Select(ctx)
	}
	return ctx.Err()
}

func (w *UpdatableTimerWorkflow) GetWakeUpTime() (*updatabletimerv1.GetWakeUpTimeOutput, error) {
	return &updatabletimerv1.GetWakeUpTimeOutput{WakeUpTime: w.wakeUpTime}, nil
}

func main() {
	app, err := updatabletimerv1.NewExampleCli(
		updatabletimerv1.NewExampleCliOptions().WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
			w := worker.New(c, updatabletimerv1.ExampleTaskQueue, worker.Options{})
			updatabletimerv1.RegisterUpdatableTimerWorkflow(w, NewUpdatableTimerWorkflow)
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
