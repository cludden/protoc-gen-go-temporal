package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mutexv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/mutex/v1"
	"github.com/cludden/protoc-gen-go-temporal/gen/example/mutex/v1/mutexv1xns"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Workflows struct{}

type MutexWorkflow struct {
	*mutexv1.MutexWorkflowInput
	log tlog.Logger
}

func (w *Workflows) Mutex(ctx workflow.Context, input *mutexv1.MutexWorkflowInput) (mutexv1.MutexWorkflow, error) {
	return &MutexWorkflow{input, workflow.GetLogger(ctx)}, nil
}

func (w *MutexWorkflow) Execute(ctx workflow.Context) error {
	for {
		req := w.AcquireLock.ReceiveAsync()
		if req == nil {
			w.log.Info("no more signals")
			return nil
		}

		var leaseID string
		if err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
			return uuid.NewString()
		}).Get(&leaseID); err != nil {
			return fmt.Errorf("error generating lease id: %w", err)
		}

		if err := mutexv1.LockAcquiredExternal(ctx, req.GetWorkflowId(), "", &mutexv1.LockAcquiredInput{
			LeaseId: leaseID,
		}); err != nil {
			return fmt.Errorf("error signaling lock acquired: %w", err)
		}

		timerCtx, cancelTimer := workflow.WithCancel(ctx)
		for done := false; !done; {
			workflow.NewSelector(ctx).
				AddFuture(workflow.NewTimer(timerCtx, req.GetTimeout().AsDuration()), func(workflow.Future) {
					w.log.Info("lease expired")
					done = true
				}).
				AddReceive(w.ReleaseLock.Channel, func(workflow.ReceiveChannel, bool) {
					if release := w.ReleaseLock.ReceiveAsync(); release.GetLeaseId() == leaseID {
						cancelTimer()
						done = true
					}
				}).
				Select(ctx)
		}
	}
}

type SampleWorkflowWithMutexWorkflow struct {
	*mutexv1.SampleWorkflowWithMutexWorkflowInput
	log tlog.Logger
}

func (w *Workflows) SampleWorkflowWithMutex(ctx workflow.Context, input *mutexv1.SampleWorkflowWithMutexWorkflowInput) (mutexv1.SampleWorkflowWithMutexWorkflow, error) {
	return &SampleWorkflowWithMutexWorkflow{input, workflow.GetLogger(ctx)}, nil
}

func (w *SampleWorkflowWithMutexWorkflow) Execute(ctx workflow.Context) error {
	w.log.Info("started", "resourceID", w.Req.GetResourceId())
	mutex, err := mutexv1xns.MutexWithAcquireLockAsync(
		ctx,
		&mutexv1.MutexInput{ResourceId: w.Req.GetResourceId()},
		&mutexv1.AcquireLockInput{
			WorkflowId: workflow.GetInfo(ctx).WorkflowExecution.ID,
			Timeout:    durationpb.New(time.Minute * 10),
		},
		mutexv1xns.NewMutexWorkflowOptions().WithDetached(true),
	)
	if err != nil {
		return err
	}

	lease, _ := w.LockAcquired.Receive(ctx)
	defer func() {
		if err := mutex.ReleaseLock(ctx, &mutexv1.ReleaseLockInput{LeaseId: lease.GetLeaseId()}); err != nil {
			w.log.Error("failed to release lock", "error", err)
		}
		w.log.Info("finished")
	}()
	w.log.Info("resource lock acquired", "leaseID", lease.GetLeaseId())

	w.log.Info("critical operation started")
	d := w.Req.GetSleep().AsDuration()
	if d == 0 {
		d = time.Second * 10
	}
	err = workflow.Sleep(ctx, w.Req.GetSleep().AsDuration())
	w.log.Info("critical operation finished")
	return err
}

func main() {
	app, err := mutexv1.NewExampleCli(
		mutexv1.NewExampleCliOptions().WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
			w := worker.New(c, mutexv1.ExampleTaskQueue, worker.Options{})
			mutexv1.RegisterExampleWorkflows(w, &Workflows{})
			mutexv1xns.RegisterExampleActivities(w, mutexv1.NewExampleClient(c))
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
