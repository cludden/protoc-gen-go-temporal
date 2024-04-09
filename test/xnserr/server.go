package xnserr

import (
	xnserrv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/xnserr/v1"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type (
	ServerWorkflows struct{}

	SleepWorkflow struct {
		*xnserrv1.SleepWorkflowInput
	}
)

func (w *ServerWorkflows) Sleep(ctx workflow.Context, input *xnserrv1.SleepWorkflowInput) (xnserrv1.SleepWorkflow, error) {
	return &SleepWorkflow{input}, nil
}

func (w *SleepWorkflow) Execute(ctx workflow.Context) error {
	if d := w.Req.GetSleep(); d.IsValid() {
		if err := workflow.Sleep(ctx, d.AsDuration()); err != nil {
			return err
		}
	}
	f := w.Req.GetFailure()
	switch f.GetInfo() {
	case xnserrv1.FailureInfo_FAILURE_INFO_ACTIVITY:
		return &temporal.ActivityError{}
	case xnserrv1.FailureInfo_FAILURE_INFO_APPLICATION_ERROR:
		errType := f.GetApplicationErrorType()
		if errType == "" {
			errType = "SleepError"
		}
		if f.GetNonRetryable() {
			return temporal.NewNonRetryableApplicationError(f.GetMessage(), errType, nil)
		}
		return temporal.NewApplicationError(f.GetMessage(), errType)
	case xnserrv1.FailureInfo_FAILURE_INFO_CANCELED:
		return &temporal.CanceledError{}
	case xnserrv1.FailureInfo_FAILURE_INFO_CHILD_WORKFLOW_EXECUTION:
		return &temporal.ChildWorkflowExecutionError{}
	case xnserrv1.FailureInfo_FAILURE_INFO_TERMINATED:
		return &temporal.TerminatedError{}
	case xnserrv1.FailureInfo_FAILURE_INFO_TIMEOUT:
		return &temporal.TimeoutError{}
	default:
		return nil
	}
}
