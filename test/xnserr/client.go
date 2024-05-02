package xnserr

import (
	xnserrv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/xnserr/v1"
	"github.com/cludden/protoc-gen-go-temporal/gen/test/xnserr/v1/xnserrv1xns"
	"github.com/cludden/protoc-gen-go-temporal/pkg/xns"
	"go.temporal.io/sdk/workflow"
)

type (
	ClientWorkflows struct{}

	CallSleepWorkflow struct {
		*xnserrv1.CallSleepWorkflowInput
	}
)

func (w *ClientWorkflows) CallSleep(ctx workflow.Context, input *xnserrv1.CallSleepWorkflowInput) (xnserrv1.CallSleepWorkflow, error) {
	return &CallSleepWorkflow{input}, nil
}

func (w *CallSleepWorkflow) Execute(ctx workflow.Context) error {
	return xnserrv1xns.Sleep(ctx, &xnserrv1.SleepRequest{
		Sleep:   w.Req.GetSleep(),
		Failure: w.Req.GetFailure(),
	}, xnserrv1xns.NewSleepWorkflowOptions().
		WithStartWorkflow(
			xns.UnmarshalStartWorkflowOptions(w.Req.GetStartWorkflowOptions()),
		).
		WithActivityOptions(workflow.ActivityOptions{
			RetryPolicy:         xns.UnmarshalRetryPolicy(w.Req.GetRetryPolicy()),
			WaitForCancellation: true,
		}),
	)
}
