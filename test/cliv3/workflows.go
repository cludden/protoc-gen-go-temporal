package cliv3

import (
	"github.com/cludden/protoc-gen-go-temporal/gen/test/cliv3"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/proto"
)

type (
	Workflows struct{}

	CreateFooWorkflow struct {
		*Workflows
		*cliv3.CreateFooWorkflowInput
		progress float64
	}
)

func Register(r worker.Registry) {
	cliv3.RegisterExampleServiceWorkflows(r, &Workflows{})
}

func (w *Workflows) CreateFoo(
	ctx workflow.Context,
	input *cliv3.CreateFooWorkflowInput,
) (cliv3.CreateFooWorkflow, error) {
	return &CreateFooWorkflow{w, input, 0}, nil
}

func (w *CreateFooWorkflow) Execute(ctx workflow.Context) (*cliv3.CreateFooOutput, error) {
	workflow.Go(ctx, func(ctx workflow.Context) {
		for {
			signal, _ := w.SignalFoo.Receive(ctx)
			if p := signal.GetProgress(); p > w.progress && p <= 1 {
				w.progress = p
			}
		}
	})

	ready, s := workflow.NewFuture(ctx)
	workflow.Go(ctx, func(ctx workflow.Context) {
		s.Set(nil, workflow.Await(ctx, func() bool {
			return w.progress == 1
		}))
	})

	ectx, ecancel := workflow.WithCancel(ctx)
	expired := workflow.NewTimer(ectx, w.Req.GetExpiresAt().AsTime().Sub(workflow.Now(ctx)))

	var isReady bool
	workflow.NewSelector(ctx).
		AddFuture(expired, func(workflow.Future) {}).
		AddFuture(ready, func(workflow.Future) {
			ecancel()
			isReady = true
		}).
		Select(ctx)
	if !isReady {
		return nil, temporal.NewNonRetryableApplicationError(
			"create foo workflow expired",
			"CreateFooWorkflowExpired",
			nil,
		)
	}
	return cliv3.CreateFooOutput_builder{
		Id: proto.String(workflow.GetInfo(ctx).WorkflowExecution.RunID),
	}.Build(), nil
}

func (w *CreateFooWorkflow) GetFoo(input *cliv3.GetFooInput) (*cliv3.GetFooOutput, error) {
	return cliv3.GetFooOutput_builder{
		Name:        proto.String(w.Req.GetName()),
		Description: proto.String(w.Req.GetDescription()),
		Progress:    proto.Float64(w.progress),
	}.Build(), nil
}

func (w *CreateFooWorkflow) UpdateFoo(
	ctx workflow.Context,
	input *cliv3.UpdateFooInput,
) (*cliv3.UpdateFooOutput, error) {
	w.progress = input.GetProgress()
	return cliv3.UpdateFooOutput_builder{
		Id: proto.String(workflow.GetCurrentUpdateInfo(ctx).ID),
	}.Build(), nil
}

func (w *CreateFooWorkflow) ValidateUpdateFoo(ctx workflow.Context, input *cliv3.UpdateFooInput) error {
	if p := input.GetProgress(); p < w.progress || p > 1 {
		return temporal.NewNonRetryableApplicationError(
			"update foo workflow invalid progress",
			"UpdateFooWorkflowInvalidProgress",
			nil,
		)
	}
	return nil
}
