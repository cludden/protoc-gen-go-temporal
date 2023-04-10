package mutex

import (
	"context"
	"fmt"
	"time"

	"github.com/cludden/protoc-gen-go-temporal/example/mutexv1"
	"github.com/google/uuid"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Workflows manages shared state for workflow constructors
type Workflows struct{}

// MutexWorkflow provides a mutex over a shared resource
type MutexWorkflow struct {
	*mutexv1.MutexInput
	log log.Logger
}

// LockAccount initializes a new MutexWorkflow value
func (w *Workflows) Mutex(ctx workflow.Context, input *mutexv1.MutexInput) (mutexv1.MutexWorkflow, error) {
	return &MutexWorkflow{input, log.With(workflow.GetLogger(ctx), "resource", input.Req.GetResource())}, nil
}

// Execute defines the entrypoint to a MutexWorkflow
func (wf *MutexWorkflow) Execute(ctx workflow.Context) error {
	for {
		wf.log.Info("dequeuing lease request")
		lease := wf.AcquireLease.ReceiveAsync()
		if lease == nil {
			return nil
		}

		wf.log.Info("generating lease id")
		var leaseID string
		if err := workflow.SideEffect(ctx, func(_ workflow.Context) any {
			return uuid.New().String()
		}).Get(&leaseID); err != nil {
			return fmt.Errorf("error generating lease id: %w", err)
		}
		log := log.With(wf.log, "lease", leaseID)

		log.Info("notifying lease holder")
		info := workflow.GetInfo(ctx)
		if err := mutexv1.LeaseAcquiredExternal(ctx, lease.GetWorkflowId(), "", &mutexv1.LeaseAcquiredSignal{
			WorkflowId: info.WorkflowExecution.ID,
			RunId:      info.WorkflowExecution.RunID,
			LeaseId:    leaseID,
		}).Get(ctx, nil); err != nil {
			return fmt.Errorf("error signalling lock acquired: %w", err)
		}

		log.Info("initializing lease timer")
		timerCtx, timerCancel := workflow.WithCancel(ctx)
		timer := workflow.NewTimer(timerCtx, lease.GetTimeout().AsDuration())

		for done := false; !done; {
			workflow.NewSelector(ctx).
				AddFuture(timer, func(f workflow.Future) {
					log.Info("lease expired")
					done = true
				}).
				AddReceive(wf.RenewLease.Channel, func(workflow.ReceiveChannel, bool) {
					s := wf.RenewLease.ReceiveAsync()
					if s.GetLeaseId() != leaseID {
						return
					}
					log.Info("extending lease")
					timerCancel()
					timerCtx, timerCancel = workflow.WithCancel(ctx)
					timer = workflow.NewTimer(timerCtx, s.GetTimeout().AsDuration())
				}).
				AddReceive(wf.RevokeLease.Channel, func(workflow.ReceiveChannel, bool) {
					s := wf.RevokeLease.ReceiveAsync()
					if s.GetLeaseId() != leaseID {
						return
					}
					log.Info("revoking lease")
					timerCancel()
					done = true
				}).
				Select(ctx)
		}
	}
}

// SampleWorkflowWithMutexWorkflow simulates a long running workflow requiring exclusive access to a shared resource
type SampleWorkflowWithMutexWorkflow struct {
	*mutexv1.SampleWorkflowWithMutexInput
	log log.Logger
}

// SampleWorkflowWithMutex initializes a new SampleWorkflowWithMutexWorkflow value
func (w *Workflows) SampleWorkflowWithMutex(ctx workflow.Context, input *mutexv1.SampleWorkflowWithMutexInput) (mutexv1.SampleWorkflowWithMutexWorkflow, error) {
	return &SampleWorkflowWithMutexWorkflow{input, log.With(
		workflow.GetLogger(ctx), "resource", input.Req.GetResource(), "workflow", workflow.GetInfo(ctx).WorkflowExecution.ID,
	)}, nil
}

// Execute defines the entrypoint to a TransferWorkflow
func (wf *SampleWorkflowWithMutexWorkflow) Execute(ctx workflow.Context) (resp *mutexv1.SampleWorkflowWithMutexResponse, err error) {
	wf.log.Info("started")

	wf.log.Info("requesting lease")
	if err := mutexv1.Mutex(ctx, nil, &mutexv1.MutexRequest{Resource: wf.Req.GetResource()}).Get(ctx); err != nil {
		return nil, fmt.Errorf("error requesting lease: %w", err)
	}

	wf.log.Info("waiting until lease acquired")
	lease, _ := wf.LeaseAcquired.Receive(ctx)
	wf.log.Info("lease acquired", "lease", lease.GetLeaseId())
	defer func() {
		wf.log.Info("revoking lease", "lease", lease.GetLeaseId())
		cancelCtx, _ := workflow.NewDisconnectedContext(ctx)
		if mutexv1.RevokeLeaseExternal(cancelCtx, lease.GetWorkflowId(), lease.GetRunId(), &mutexv1.RevokeLeaseSignal{
			LeaseId: lease.GetLeaseId(),
		}).Get(ctx, nil); err != nil {
			wf.log.Error("error revoking lease", "error", err, "lease", lease.GetLeaseId())
		}
	}()

	// emulate long running process
	wf.log.Info("critical operation started")
	_ = workflow.Sleep(ctx, 10*time.Second)
	wf.log.Info("critical operation finished")

	return &mutexv1.SampleWorkflowWithMutexResponse{Result: lease.GetLeaseId()}, nil
}

// Activities manages shared state for activities
type Activites struct {
	Client mutexv1.Client
}

// Mutex locks a shared resource and can be called from a parent workflow
func (a *Activites) Mutex(ctx context.Context, req *mutexv1.MutexRequest) error {
	_, err := a.Client.StartMutexWithAcquireLease(ctx, nil, req, &mutexv1.AcquireLeaseSignal{
		WorkflowId: activity.GetInfo(ctx).WorkflowExecution.ID,
		Timeout:    durationpb.New(time.Minute * 2),
	})
	return err
}
