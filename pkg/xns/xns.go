package xns

import (
	"errors"
	"fmt"

	xnsv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/xns/v1"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/api/update/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
)

func Code(err error) string {
	if terr := Unwrap(err); terr != nil {
		return terr.Type()
	}
	return ""
}

func IsNonRetryable(err error) bool {
	if terr := Unwrap(err); terr != nil {
		return terr.NonRetryable()
	}
	return false
}

func Unwrap(err error) *temporal.ApplicationError {
	if err == nil {
		return nil
	}
	var x *temporal.ApplicationError
	if errors.As(err, &x) {
		return x
	}
	return Unwrap(errors.Unwrap(err))
}

// ErrorToApplicationError converts an arbitrary Go error into a temporal application error
// with the appropriate retryable configuration
func ErrorToApplicationError(err error) error {
	if err == nil {
		return nil
	}

	// extract workflow execution cause
	var workflowExecutionErr *temporal.WorkflowExecutionError
	if errors.As(err, &workflowExecutionErr) {
		if inner := workflowExecutionErr.Unwrap(); inner != nil {
			err = inner
		}
	}

	var application *temporal.ApplicationError
	if errors.As(err, &application) {
		return temporal.NewNonRetryableApplicationError(application.Message(), application.Type(), application)
	}

	var childWorkflowExecutionErr *temporal.ChildWorkflowExecutionError
	if errors.As(err, &childWorkflowExecutionErr) {
		return temporal.NewNonRetryableApplicationError(childWorkflowExecutionErr.Error(), "ChildWorkflowExecutionError", childWorkflowExecutionErr.Unwrap())
	}

	var canceledErr *temporal.CanceledError
	if errors.As(err, &canceledErr) {
		return temporal.NewNonRetryableApplicationError(canceledErr.Error(), "CanceledError", canceledErr)
	}

	var terminatedErr *temporal.TerminatedError
	if errors.As(err, &terminatedErr) {
		return temporal.NewNonRetryableApplicationError(terminatedErr.Error(), "TerminatedError", terminatedErr)
	}

	var timeoutErr *temporal.TimeoutError
	if errors.As(err, &timeoutErr) {
		return temporal.NewNonRetryableApplicationError(timeoutErr.Error(), "TimeoutError", timeoutErr)
	}

	if errors.As(err, &workflowExecutionErr) {
		return temporal.NewNonRetryableApplicationError(workflowExecutionErr.Error(), "WorkflowExecutionError", workflowExecutionErr.Unwrap())
	}

	return err
}

func MarshalStartWorkflowOptions(o client.StartWorkflowOptions) (*xnsv1.StartWorkflowOptions, error) {
	opts := &xnsv1.StartWorkflowOptions{
		EnableEagerStart:        o.EnableEagerStart,
		ErrorWhenAlreadyStarted: o.WorkflowExecutionErrorWhenAlreadyStarted,
		ExecutionTimeout:        durationpb.New(o.WorkflowExecutionTimeout),
		Id:                      o.ID,
		RunTimeout:              durationpb.New(o.WorkflowRunTimeout),
		StartDelay:              durationpb.New(o.StartDelay),
		TaskQueue:               o.TaskQueue,
		TaskTimeout:             durationpb.New(o.WorkflowTaskTimeout),
	}
	// id reuse
	switch o.WorkflowIDReusePolicy {
	case enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE:
		opts.IdReusePolicy = xnsv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
	case enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY:
		opts.IdReusePolicy = xnsv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY
	case enums.WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE:
		opts.IdReusePolicy = xnsv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE
	case enums.WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING:
		opts.IdReusePolicy = xnsv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING
	}
	// memo
	if len(o.Memo) > 0 {
		memo, err := structpb.NewStruct(o.Memo)
		if err != nil {
			return nil, fmt.Errorf("error marshalling memo: %w", err)
		}
		opts.Memo = memo
	}
	// retry policy
	if o.RetryPolicy != nil {
		opts.RetryPolicy = &xnsv1.RetryPolicy{
			BackoffCoefficient:     o.RetryPolicy.BackoffCoefficient,
			InitialInterval:        durationpb.New(o.RetryPolicy.InitialInterval),
			MaxAttempts:            o.RetryPolicy.MaximumAttempts,
			MaxInterval:            durationpb.New(o.RetryPolicy.MaximumInterval),
			NonRetryableErrorTypes: o.RetryPolicy.NonRetryableErrorTypes,
		}
	}
	// search attributes
	if len(o.SearchAttributes) > 0 {
		sa, err := structpb.NewStruct(o.SearchAttributes)
		if err != nil {
			return nil, fmt.Errorf("error marshalling search attributes: %w", err)
		}
		opts.SearchAttirbutes = sa
	}
	return opts, nil
}

func UnmarshalStartWorkflowOptions(o *xnsv1.StartWorkflowOptions) client.StartWorkflowOptions {
	opts := client.StartWorkflowOptions{}
	if v := o.GetEnableEagerStart(); v {
		opts.EnableEagerStart = v
	}
	if v := o.GetErrorWhenAlreadyStarted(); v {
		opts.WorkflowExecutionErrorWhenAlreadyStarted = v
	}
	if v := o.GetExecutionTimeout(); v.IsValid() {
		opts.WorkflowExecutionTimeout = v.AsDuration()
	}
	if v := o.GetId(); v != "" {
		opts.ID = v
	}
	if v := o.GetIdReusePolicy(); v != xnsv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		switch v {
		case xnsv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE:
			opts.WorkflowIDReusePolicy = enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
		case xnsv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY:
			opts.WorkflowIDReusePolicy = enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY
		case xnsv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE:
			opts.WorkflowIDReusePolicy = enums.WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE
		case xnsv1.IDReusePolicy_WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING:
			opts.WorkflowIDReusePolicy = enums.WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING
		}
	}
	if v := o.GetMemo(); len(v.GetFields()) > 0 {
		opts.Memo = v.AsMap()
	}
	if v := o.GetRetryPolicy(); v != nil {
		opts.RetryPolicy = UnmarshalRetryPolicy(v)
	}
	if v := o.GetRunTimeout(); v.IsValid() {
		opts.WorkflowRunTimeout = v.AsDuration()
	}
	if v := o.GetSearchAttirbutes(); len(v.GetFields()) > 0 {
		opts.SearchAttributes = v.AsMap()
	}
	if v := o.GetStartDelay(); v.IsValid() {
		opts.StartDelay = v.AsDuration()
	}
	if v := o.GetTaskQueue(); v != "" {
		opts.TaskQueue = v
	}
	if v := o.GetTaskTimeout(); v.IsValid() {
		opts.WorkflowTaskTimeout = v.AsDuration()
	}
	return opts
}

func UnmarshalRetryPolicy(rp *xnsv1.RetryPolicy) *temporal.RetryPolicy {
	if rp == nil {
		return nil
	}

	result := &temporal.RetryPolicy{}
	empty := true
	if x := rp.GetBackoffCoefficient(); x != 0 {
		result.BackoffCoefficient, empty = x, false
	}
	if x := rp.GetInitialInterval(); x.IsValid() {
		result.InitialInterval, empty = x.AsDuration(), false
	}
	if x := rp.GetMaxAttempts(); x > 0 {
		result.MaximumAttempts, empty = x, false
	}
	if x := rp.GetMaxInterval(); x.IsValid() {
		result.MaximumInterval, empty = x.AsDuration(), false
	}
	if x := rp.GetNonRetryableErrorTypes(); len(x) > 0 {
		result.NonRetryableErrorTypes, empty = x, false
	}

	if empty {
		return nil
	}
	return result
}

func MarshalUpdateWorkflowOptions(o client.UpdateWorkflowWithOptionsRequest) (*xnsv1.UpdateWorkflowWithOptionsRequest, error) {
	opts := &xnsv1.UpdateWorkflowWithOptionsRequest{
		UpdateId:            o.UpdateID,
		WorkflowId:          o.WorkflowID,
		RunId:               o.RunID,
		FirstExecutionRunId: o.FirstExecutionRunID,
	}
	if o.WaitPolicy != nil {
		switch o.WaitPolicy.LifecycleStage {
		case enums.UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ACCEPTED:
			opts.WaitPolicy = xnsv1.WaitPolicy_WAIT_POLICY_ACCEPTED
		case enums.UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ADMITTED:
			opts.WaitPolicy = xnsv1.WaitPolicy_WAIT_POLICY_ADMITTED
		case enums.UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_COMPLETED:
			opts.WaitPolicy = xnsv1.WaitPolicy_WAIT_POLICY_COMPLETED
		}
	}
	return opts, nil
}

func UnmarshalUpdateWorkflowOptions(o *xnsv1.UpdateWorkflowWithOptionsRequest) client.UpdateWorkflowWithOptionsRequest {
	opts := client.UpdateWorkflowWithOptionsRequest{
		UpdateID:            o.GetUpdateId(),
		WorkflowID:          o.GetWorkflowId(),
		RunID:               o.GetRunId(),
		FirstExecutionRunID: o.GetFirstExecutionRunId(),
	}
	switch o.WaitPolicy {
	case xnsv1.WaitPolicy_WAIT_POLICY_ACCEPTED:
		opts.WaitPolicy = &update.WaitPolicy{LifecycleStage: enums.UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ACCEPTED}
	case xnsv1.WaitPolicy_WAIT_POLICY_ADMITTED:
		opts.WaitPolicy = &update.WaitPolicy{LifecycleStage: enums.UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ADMITTED}
	case xnsv1.WaitPolicy_WAIT_POLICY_COMPLETED:
		opts.WaitPolicy = &update.WaitPolicy{LifecycleStage: enums.UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_COMPLETED}
	}
	return opts
}
