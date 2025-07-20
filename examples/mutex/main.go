package main

import (
	"cmp"
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

type (
	Workflows struct{}

	MutexWorkflow struct {
		*mutexv1.MutexWorkflowInput
		log      tlog.Logger
		requests []*leaseRequest
	}

	leaseRequest struct {
		input *mutexv1.AcquireLockInput
		f     workflow.Settable
	}
)

func (w *Workflows) Mutex(
	ctx workflow.Context,
	input *mutexv1.MutexWorkflowInput,
) (mutexv1.MutexWorkflow, error) {
	return &MutexWorkflow{
		input,
		workflow.GetLogger(ctx),
		make([]*leaseRequest, 0),
	}, nil
}

func (w *MutexWorkflow) Execute(ctx workflow.Context) error {
	var req *leaseRequest
	for len(w.requests) > 0 {
		if req, w.requests = w.requests[0], w.requests[1:]; req == nil {
			w.log.Info("no more lock requests")
			return nil
		}

		var leaseID string
		if err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
			return uuid.NewString()
		}).Get(&leaseID); err != nil {
			req.f.SetError(fmt.Errorf("error generating lease id: %w", err))
			continue
		}
		req.f.SetValue(leaseID)

		timerCtx, cancelTimer := workflow.WithCancel(ctx)
		for done := false; !done; {
			timer := workflow.NewTimer(timerCtx, req.input.GetTimeout().AsDuration())
			workflow.NewSelector(ctx).
				AddFuture(timer, func(workflow.Future) {
					w.log.Info("lease expired")
					done = true
				}).
				AddReceive(w.ReleaseLock.Channel, func(workflow.ReceiveChannel, bool) {
					if r := w.ReleaseLock.ReceiveAsync(); r.GetLeaseId() == leaseID {
						cancelTimer()
						done = true
					}
				}).
				Select(ctx)
		}
	}
	return nil
}

func (w *MutexWorkflow) AcquireLock(
	ctx workflow.Context,
	input *mutexv1.AcquireLockInput,
) (*mutexv1.AcquireLockOutput, error) {
	f, set := workflow.NewFuture(ctx)
	w.requests = append(w.requests, &leaseRequest{
		input: input,
		f:     set,
	})

	var leaseID string
	if err := f.Get(ctx, &leaseID); err != nil {
		return nil, fmt.Errorf("error waiting for lease id: %w", err)
	}
	return &mutexv1.AcquireLockOutput{LeaseId: leaseID}, nil
}

type SampleWorkflowWithMutexWorkflow struct {
	*mutexv1.SampleWorkflowWithMutexWorkflowInput
	log tlog.Logger
}

func (w *Workflows) SampleWorkflowWithMutex(
	ctx workflow.Context,
	input *mutexv1.SampleWorkflowWithMutexWorkflowInput,
) (mutexv1.SampleWorkflowWithMutexWorkflow, error) {
	return &SampleWorkflowWithMutexWorkflow{input, workflow.GetLogger(ctx)}, nil
}

func (w *SampleWorkflowWithMutexWorkflow) Execute(ctx workflow.Context) error {
	w.log.Info("started", "resourceID", w.Req.GetResourceId())
	lease, mutex, err := mutexv1xns.MutexWithAcquireLock(
		ctx,
		&mutexv1.MutexInput{ResourceId: w.Req.GetResourceId()},
		&mutexv1.AcquireLockInput{
			Timeout: durationpb.New(time.Minute * 10),
		},
	)
	if err != nil {
		return err
	}
	defer func() {
		if err := mutex.ReleaseLock(ctx, &mutexv1.ReleaseLockInput{
			LeaseId: lease.GetLeaseId(),
		}); err != nil {
			w.log.Error("failed to release lock", "error", err)
		}
	}()
	w.log.Info("resource lock acquired", "leaseID", lease.GetLeaseId())

	w.log.Info("critical operation started")
	err = workflow.Sleep(ctx, cmp.Or(w.Req.GetSleep().AsDuration(), time.Second*10))
	w.log.Info("critical operation finished")
	return err
}

func main() {
	app, err := mutexv1.NewExampleCli(
		mutexv1.NewExampleCliOptions().
			WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
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
