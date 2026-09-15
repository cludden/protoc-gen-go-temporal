package xns

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	xnsv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/xns/v1"
	"github.com/cludden/protoc-gen-go-temporal/pkg/convert"
	commonpb "go.temporal.io/api/common/v1"
	deploymentpb "go.temporal.io/api/deployment/v1"
	"go.temporal.io/api/enums/v1"
	workflowpb "go.temporal.io/api/workflow/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
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

	if errors.Is(err, context.Canceled) {
		return temporal.NewCanceledError(err.Error())
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
		CronSchedule:             o.CronSchedule,
		EnableEagerStart:         o.EnableEagerStart,
		ErrorWhenAlreadyStarted:  o.WorkflowExecutionErrorWhenAlreadyStarted,
		ExecutionTimeout:         durationpb.New(o.WorkflowExecutionTimeout),
		Id:                       o.ID,
		RunTimeout:               durationpb.New(o.WorkflowRunTimeout),
		StartDelay:               durationpb.New(o.StartDelay),
		StaticDetails:            o.StaticDetails,
		StaticSummary:            o.StaticSummary,
		TaskQueue:                o.TaskQueue,
		TaskTimeout:              durationpb.New(o.WorkflowTaskTimeout),
		WorkflowIdConflictPolicy: o.WorkflowIDConflictPolicy,
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
	// priority
	if o.Priority != (temporal.Priority{}) {
		priorityKey, err := convert.SafeCast[int, int32](o.Priority.PriorityKey)
		if err != nil {
			return nil, fmt.Errorf("error marshalling priority key: %w", err)
		}
		opts.Priority = &commonpb.Priority{
			PriorityKey:    priorityKey,
			FairnessKey:    o.Priority.FairnessKey,
			FairnessWeight: o.Priority.FairnessWeight,
		}
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
	// typed search attributes
	if o.TypedSearchAttributes.Size() > 0 {
		tsa, err := marshalTypedSearchAttributes(o.TypedSearchAttributes)
		if err != nil {
			return nil, err
		}
		opts.TypedSearchAttributes = tsa
	}
	// versioning override
	vo, err := marshalVersioningOverride(o.VersioningOverride)
	if err != nil {
		return nil, err
	}
	opts.VersioningOverride = vo
	return opts, nil
}

func UnmarshalStartWorkflowOptions(o *xnsv1.StartWorkflowOptions) client.StartWorkflowOptions {
	opts := client.StartWorkflowOptions{}
	if v := o.GetCronSchedule(); v != "" {
		opts.CronSchedule = v
	}
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
	if v := o.GetPriority(); v != nil {
		opts.Priority = temporal.Priority{
			PriorityKey:    int(v.GetPriorityKey()),
			FairnessKey:    v.GetFairnessKey(),
			FairnessWeight: v.GetFairnessWeight(),
		}
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
	if v := o.GetStaticDetails(); v != "" {
		opts.StaticDetails = v
	}
	if v := o.GetStaticSummary(); v != "" {
		opts.StaticSummary = v
	}
	if v := o.GetTaskQueue(); v != "" {
		opts.TaskQueue = v
	}
	if v := o.GetTaskTimeout(); v.IsValid() {
		opts.WorkflowTaskTimeout = v.AsDuration()
	}
	if v := o.GetTypedSearchAttributes(); len(v.GetIndexedFields()) > 0 {
		opts.TypedSearchAttributes = unmarshalTypedSearchAttributes(v)
	}
	if v := unmarshalVersioningOverride(o.GetVersioningOverride()); v != nil {
		opts.VersioningOverride = v
	}
	if v := o.GetWorkflowIdConflictPolicy(); v != enums.WORKFLOW_ID_CONFLICT_POLICY_UNSPECIFIED {
		opts.WorkflowIDConflictPolicy = v
	}
	return opts
}

// marshalVersioningOverride encodes a client.VersioningOverride as its canonical
// protobuf representation. Only the override oneof is populated: the deprecated
// behavior, deployment, and pinned_version fields exist for servers that predate
// it, and nothing on this wire is read by a server.
func marshalVersioningOverride(vo client.VersioningOverride) (*workflowpb.VersioningOverride, error) {
	switch v := vo.(type) {
	case nil:
		return nil, nil
	case *client.PinnedVersioningOverride:
		// a non-nil interface can still hold a nil pointer
		if v == nil {
			return nil, nil
		}
		return &workflowpb.VersioningOverride{
			Override: &workflowpb.VersioningOverride_Pinned{
				Pinned: &workflowpb.VersioningOverride_PinnedOverride{
					Behavior: workflowpb.VersioningOverride_PINNED_OVERRIDE_BEHAVIOR_PINNED,
					Version: &deploymentpb.WorkerDeploymentVersion{
						BuildId:        v.Version.BuildID,
						DeploymentName: v.Version.DeploymentName,
					},
				},
			},
		}, nil
	case *client.AutoUpgradeVersioningOverride:
		if v == nil {
			return nil, nil
		}
		return &workflowpb.VersioningOverride{
			Override: &workflowpb.VersioningOverride_AutoUpgrade{AutoUpgrade: true},
		}, nil
	default:
		return nil, fmt.Errorf("unsupported versioning override type: %T", vo)
	}
}

// unmarshalVersioningOverride decodes the override written by
// marshalVersioningOverride, falling back to the deprecated fields for values
// written by other tooling.
func unmarshalVersioningOverride(vo *workflowpb.VersioningOverride) client.VersioningOverride {
	if vo == nil {
		return nil
	}
	switch override := vo.GetOverride().(type) {
	case *workflowpb.VersioningOverride_AutoUpgrade:
		if override.AutoUpgrade {
			return &client.AutoUpgradeVersioningOverride{}
		}
	case *workflowpb.VersioningOverride_Pinned:
		return &client.PinnedVersioningOverride{
			Version: worker.WorkerDeploymentVersion{
				BuildID:        override.Pinned.GetVersion().GetBuildId(),
				DeploymentName: override.Pinned.GetVersion().GetDeploymentName(),
			},
		}
	}
	// the deprecated fields are still read, for values written by other tooling
	switch vo.GetBehavior() { //nolint:staticcheck
	case enums.VERSIONING_BEHAVIOR_AUTO_UPGRADE:
		return &client.AutoUpgradeVersioningOverride{}
	case enums.VERSIONING_BEHAVIOR_PINNED:
		if d := vo.GetDeployment(); d != nil { //nolint:staticcheck
			return &client.PinnedVersioningOverride{
				Version: worker.WorkerDeploymentVersion{
					BuildID:        d.GetBuildId(),
					DeploymentName: d.GetSeriesName(),
				},
			}
		}
		// canonical pinned version strings are "<deployment_name>.<build_id>"
		if name, buildID, ok := strings.Cut(vo.GetPinnedVersion(), "."); ok { //nolint:staticcheck
			return &client.PinnedVersioningOverride{
				Version: worker.WorkerDeploymentVersion{
					BuildID:        buildID,
					DeploymentName: name,
				},
			}
		}
	}
	return nil
}

// marshalTypedSearchAttributes encodes typed search attributes as the canonical
// payload map, preserving each attribute's indexed value type.
func marshalTypedSearchAttributes(sa temporal.SearchAttributes) (*commonpb.SearchAttributes, error) {
	fields := make(map[string]*commonpb.Payload, sa.Size())
	for key, value := range sa.GetUntypedValues() {
		payload, err := converter.GetDefaultDataConverter().ToPayload(value)
		if err != nil {
			return nil, fmt.Errorf("error marshalling typed search attribute %q: %w", key.GetName(), err)
		}
		// the server drops an attribute that arrives without a type
		if payload.GetData() != nil {
			if payload.Metadata == nil {
				payload.Metadata = make(map[string][]byte, 1)
			}
			payload.Metadata["type"] = []byte(key.GetValueType().String())
		}
		fields[key.GetName()] = payload
	}
	return &commonpb.SearchAttributes{IndexedFields: fields}, nil
}

// unmarshalTypedSearchAttributes decodes the payload map written by
// marshalTypedSearchAttributes. An attribute carrying an unrecognized type, or
// one whose payload fails to decode, is skipped rather than surfaced: this is
// unreachable for values written by MarshalStartWorkflowOptions, which refuses
// to encode an attribute it cannot represent.
func unmarshalTypedSearchAttributes(sa *commonpb.SearchAttributes) temporal.SearchAttributes {
	dc := converter.GetDefaultDataConverter()
	updates := make([]temporal.SearchAttributeUpdate, 0, len(sa.GetIndexedFields()))
	for name, payload := range sa.GetIndexedFields() {
		if payload.GetData() == nil {
			continue
		}
		valueType, err := enums.IndexedValueTypeFromString(string(payload.GetMetadata()["type"]))
		if err != nil {
			continue
		}
		switch valueType {
		case enums.INDEXED_VALUE_TYPE_BOOL:
			var value bool
			if err := dc.FromPayload(payload, &value); err == nil {
				updates = append(updates, temporal.NewSearchAttributeKeyBool(name).ValueSet(value))
			}
		case enums.INDEXED_VALUE_TYPE_DATETIME:
			var value time.Time
			if err := dc.FromPayload(payload, &value); err == nil {
				updates = append(updates, temporal.NewSearchAttributeKeyTime(name).ValueSet(value))
			}
		case enums.INDEXED_VALUE_TYPE_DOUBLE:
			var value float64
			if err := dc.FromPayload(payload, &value); err == nil {
				updates = append(updates, temporal.NewSearchAttributeKeyFloat64(name).ValueSet(value))
			}
		case enums.INDEXED_VALUE_TYPE_INT:
			var value int64
			if err := dc.FromPayload(payload, &value); err == nil {
				updates = append(updates, temporal.NewSearchAttributeKeyInt64(name).ValueSet(value))
			}
		case enums.INDEXED_VALUE_TYPE_KEYWORD:
			var value string
			if err := dc.FromPayload(payload, &value); err == nil {
				updates = append(updates, temporal.NewSearchAttributeKeyKeyword(name).ValueSet(value))
			}
		case enums.INDEXED_VALUE_TYPE_KEYWORD_LIST:
			var value []string
			if err := dc.FromPayload(payload, &value); err == nil {
				updates = append(updates, temporal.NewSearchAttributeKeyKeywordList(name).ValueSet(value))
			}
		case enums.INDEXED_VALUE_TYPE_TEXT:
			var value string
			if err := dc.FromPayload(payload, &value); err == nil {
				updates = append(updates, temporal.NewSearchAttributeKeyString(name).ValueSet(value))
			}
		}
	}
	return temporal.NewSearchAttributes(updates...)
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

func MarshalUpdateWorkflowOptions(o client.UpdateWorkflowOptions) (*xnsv1.UpdateWorkflowWithOptionsRequest, error) {
	opts := &xnsv1.UpdateWorkflowWithOptionsRequest{
		UpdateId:            o.UpdateID,
		WorkflowId:          o.WorkflowID,
		RunId:               o.RunID,
		FirstExecutionRunId: o.FirstExecutionRunID,
	}
	switch o.WaitForStage {
	case client.WorkflowUpdateStageAccepted:
		opts.WaitPolicy = xnsv1.WaitPolicy_WAIT_POLICY_ACCEPTED
	case client.WorkflowUpdateStageAdmitted:
		opts.WaitPolicy = xnsv1.WaitPolicy_WAIT_POLICY_ADMITTED
	case client.WorkflowUpdateStageCompleted:
		opts.WaitPolicy = xnsv1.WaitPolicy_WAIT_POLICY_COMPLETED
	}
	return opts, nil
}

func UnmarshalUpdateWorkflowOptions(o *xnsv1.UpdateWorkflowWithOptionsRequest) client.UpdateWorkflowOptions {
	opts := client.UpdateWorkflowOptions{
		UpdateID:            o.GetUpdateId(),
		WorkflowID:          o.GetWorkflowId(),
		RunID:               o.GetRunId(),
		FirstExecutionRunID: o.GetFirstExecutionRunId(),
	}

	wp := o.GetWaitForStage()
	if wp == xnsv1.WaitPolicy_WAIT_POLICY_UNSPECIFIED && o.GetWaitPolicy() != xnsv1.WaitPolicy_WAIT_POLICY_UNSPECIFIED {
		wp = o.GetWaitPolicy()
	}

	switch wp {
	case xnsv1.WaitPolicy_WAIT_POLICY_ACCEPTED:
		opts.WaitForStage = client.WorkflowUpdateStageAccepted
	case xnsv1.WaitPolicy_WAIT_POLICY_ADMITTED:
		opts.WaitForStage = client.WorkflowUpdateStageAdmitted
	case xnsv1.WaitPolicy_WAIT_POLICY_COMPLETED:
		opts.WaitForStage = client.WorkflowUpdateStageCompleted
	}
	return opts
}
