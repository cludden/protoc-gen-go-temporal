// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 1.14.6-next (3d95ea2ffa441f72e0ab5441f2394e8f6115afe9)
//	go go1.23.4
//	protoc (unknown)
//
// source: test/patch/example.proto
package patchxns

import (
	"context"
	"errors"
	"fmt"
	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	xnsv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/xns/v1"
	patch "github.com/cludden/protoc-gen-go-temporal/gen/test/patch"
	xns "github.com/cludden/protoc-gen-go-temporal/pkg/xns"
	uuid "github.com/google/uuid"
	enumsv1 "go.temporal.io/api/enums/v1"
	activity "go.temporal.io/sdk/activity"
	client "go.temporal.io/sdk/client"
	temporal "go.temporal.io/sdk/temporal"
	worker "go.temporal.io/sdk/worker"
	workflow "go.temporal.io/sdk/workflow"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	"time"
)

// FooServiceOptions is used to configure test.patch.FooService xns activity registration
type FooServiceOptions struct {
	// errorConverter is used to customize error
	errorConverter func(error) error
	// filter is used to filter xns activity registrations. It receives as
	// input the original activity name, and should return one of the following:
	// 1. the original activity name, for no changes
	// 2. a modified activity name, to override the original activity name
	// 3. an empty string, to skip registration
	filter func(string) string
}

// NewFooServiceOptions initializes a new FooServiceOptions value
func NewFooServiceOptions() *FooServiceOptions {
	return &FooServiceOptions{}
}

// WithErrorConverter overrides the default error converter applied to xns activity errors
func (opts *FooServiceOptions) WithErrorConverter(errorConverter func(error) error) *FooServiceOptions {
	opts.errorConverter = errorConverter
	return opts
}

// Filter is used to filter registered xns activities or customize their name
func (opts *FooServiceOptions) WithFilter(filter func(string) string) *FooServiceOptions {
	opts.filter = filter
	return opts
}

// convertError is applied to all xns activity errors
func (opts *FooServiceOptions) convertError(err error) error {
	if err == nil {
		return nil
	}
	if opts != nil && opts.errorConverter != nil {
		return opts.errorConverter(err)
	}
	return xns.ErrorToApplicationError(err)
}

// filterActivity is used to filter xns activity registrations
func (opts *FooServiceOptions) filterActivity(name string) string {
	if opts == nil || opts.filter == nil {
		return name
	}
	return opts.filter(name)
}

// fooServiceOptions is a reference to the FooServiceOptions initialized at registration
var fooServiceOptions *FooServiceOptions

// RegisterFooServiceActivities registers test.patch.FooService cross-namespace activities
func RegisterFooServiceActivities(r worker.ActivityRegistry, c patch.FooServiceClient, options ...*FooServiceOptions) {
	if fooServiceOptions == nil && len(options) > 0 && options[0] != nil {
		fooServiceOptions = options[0]
	}
	a := &fooServiceActivities{c}
	if name := fooServiceOptions.filterActivity("test.patch.FooService.CancelWorkflow"); name != "" {
		r.RegisterActivityWithOptions(a.CancelWorkflow, activity.RegisterOptions{Name: name})
	}
	if name := fooServiceOptions.filterActivity(patch.FooWorkflowName); name != "" {
		r.RegisterActivityWithOptions(a.Foo, activity.RegisterOptions{Name: name})
	}
}

// FooWorkflowOptions are used to configure a(n) test.patch.FooService.Foo workflow execution
type FooWorkflowOptions struct {
	ActivityOptions      *workflow.ActivityOptions
	Detached             bool
	HeartbeatInterval    time.Duration
	ParentClosePolicy    enumsv1.ParentClosePolicy
	StartWorkflowOptions *client.StartWorkflowOptions
}

// NewFooWorkflowOptions initializes a new FooWorkflowOptions value
func NewFooWorkflowOptions() *FooWorkflowOptions {
	return &FooWorkflowOptions{}
}

// WithActivityOptions can be used to customize the activity options
func (opts *FooWorkflowOptions) WithActivityOptions(ao workflow.ActivityOptions) *FooWorkflowOptions {
	opts.ActivityOptions = &ao
	return opts
}

// WithDetached can be used to start a workflow execution and exit immediately
func (opts *FooWorkflowOptions) WithDetached(d bool) *FooWorkflowOptions {
	opts.Detached = d
	return opts
}

// WithHeartbeatInterval can be used to customize the activity heartbeat interval
func (opts *FooWorkflowOptions) WithHeartbeatInterval(d time.Duration) *FooWorkflowOptions {
	opts.HeartbeatInterval = d
	return opts
}

// WithParentClosePolicy can be used to customize the cancellation propagation behavior
func (opts *FooWorkflowOptions) WithParentClosePolicy(policy enumsv1.ParentClosePolicy) *FooWorkflowOptions {
	opts.ParentClosePolicy = policy
	return opts
}

// WithStartWorkflowOptions can be used to customize the start workflow options
func (opts *FooWorkflowOptions) WithStartWorkflow(swo client.StartWorkflowOptions) *FooWorkflowOptions {
	opts.StartWorkflowOptions = &swo
	return opts
}

// FooRun provides a handle to a test.patch.FooService.Foo workflow execution
type FooRun interface {
	// Cancel cancels the workflow
	Cancel(workflow.Context) error

	// Future returns the inner workflow.Future
	Future() workflow.Future

	// Get returns the inner workflow.Future
	Get(workflow.Context) (*patch.FooOutput, error)

	// ID returns the workflow id
	ID() string
}

// fooRun provides a(n) FooRun implementation
type fooRun struct {
	cancel func()
	future workflow.Future
	id     string
}

// Cancel the underlying workflow execution
func (r *fooRun) Cancel(ctx workflow.Context) error {
	if r.cancel != nil {
		r.cancel()
		if _, err := r.Get(ctx); err != nil && !errors.Is(err, workflow.ErrCanceled) {
			return err
		}
		return nil
	}
	return CancelFooServiceWorkflow(ctx, r.id, "")
}

// Future returns the underlying activity future
func (r *fooRun) Future() workflow.Future {
	return r.future
}

// Get blocks on activity completion and returns the underlying workflow result
func (r *fooRun) Get(ctx workflow.Context) (*patch.FooOutput, error) {
	var resp patch.FooOutput
	if err := r.future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ID returns the underlying workflow id
func (r *fooRun) ID() string {
	return r.id
}

// Foo executes a(n) test.patch.FooService.Foo workflow and blocks until error or response is received
func Foo(ctx workflow.Context, req *patch.FooInput, opts ...*FooWorkflowOptions) (*patch.FooOutput, error) {
	run, err := FooAsync(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// FooAsync executes a(n) test.patch.FooService.Foo workflow and returns a handle to the underlying activity
func FooAsync(ctx workflow.Context, req *patch.FooInput, opts ...*FooWorkflowOptions) (FooRun, error) {
	activityName := fooServiceOptions.filterActivity(patch.FooWorkflowName)
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", patch.FooWorkflowName),
			"Unimplemented",
			nil,
		)
	}

	opt := &FooWorkflowOptions{}
	if len(opts) > 0 && opts[0] != nil {
		opt = opts[0]
	}
	if opt.HeartbeatInterval == 0 {
		opt.HeartbeatInterval = time.Second * 30
	}

	// configure activity options
	ao := workflow.GetActivityOptions(ctx)
	if opt.ActivityOptions != nil {
		ao = *opt.ActivityOptions
	}
	if ao.HeartbeatTimeout == 0 {
		ao.HeartbeatTimeout = opt.HeartbeatInterval * 2
	}
	// WaitForCancellation must be set otherwise the underlying workflow is not guaranteed to be canceled
	ao.WaitForCancellation = true

	if ao.StartToCloseTimeout == 0 && ao.ScheduleToCloseTimeout == 0 {
		ao.ScheduleToCloseTimeout = 86400000000000 // 1 day
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// configure start workflow options
	wo := client.StartWorkflowOptions{}
	if opt.StartWorkflowOptions != nil {
		wo = *opt.StartWorkflowOptions
	}
	if wo.ID == "" {
		if err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
			id, err := uuid.NewRandom()
			if err != nil {
				workflow.GetLogger(ctx).Error("error generating workflow id", "error", err)
				return nil
			}
			return id
		}).Get(&wo.ID); err != nil {
			return nil, err
		}
	}
	if wo.ID == "" {
		return nil, temporal.NewNonRetryableApplicationError("workflow id is required", "InvalidArgument", nil)
	}

	// marshal start workflow options protobuf message
	swo, err := xns.MarshalStartWorkflowOptions(wo)
	if err != nil {
		return nil, fmt.Errorf("error marshalling start workflow options: %w", err)
	}

	// marshal workflow request protobuf message
	wreq, err := anypb.New(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling workflow request: %w", err)
	}

	var parentClosePolicy temporalv1.ParentClosePolicy
	switch opt.ParentClosePolicy {
	case enumsv1.PARENT_CLOSE_POLICY_ABANDON:
		parentClosePolicy = temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_ABANDON
	case enumsv1.PARENT_CLOSE_POLICY_REQUEST_CANCEL:
		parentClosePolicy = temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL
	case enumsv1.PARENT_CLOSE_POLICY_TERMINATE:
		parentClosePolicy = temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE
	}

	ctx, cancel := workflow.WithCancel(ctx)
	return &fooRun{
		cancel: cancel,
		id:     wo.ID,
		future: workflow.ExecuteActivity(ctx, activityName, &xnsv1.WorkflowRequest{
			Detached:             opt.Detached,
			HeartbeatInterval:    durationpb.New(opt.HeartbeatInterval),
			ParentClosePolicy:    parentClosePolicy,
			Request:              wreq,
			StartWorkflowOptions: swo,
		}),
	}, nil
}

// CancelFooServiceWorkflow cancels an existing workflow
func CancelFooServiceWorkflow(ctx workflow.Context, workflowID string, runID string) error {
	return CancelFooServiceWorkflowAsync(ctx, workflowID, runID).Get(ctx, nil)
}

// CancelFooServiceWorkflowAsync cancels an existing workflow
func CancelFooServiceWorkflowAsync(ctx workflow.Context, workflowID string, runID string) workflow.Future {
	activityName := fooServiceOptions.filterActivity("test.patch.FooService.CancelWorkflow")
	if activityName == "" {
		f, s := workflow.NewFuture(ctx)
		s.SetError(temporal.NewNonRetryableApplicationError(
			"no activity registered for test.patch.FooService.CancelWorkflow",
			"Unimplemented",
			nil,
		))
		return f
	}
	ao := workflow.GetActivityOptions(ctx)
	if ao.StartToCloseTimeout == 0 && ao.ScheduleToCloseTimeout == 0 {
		ao.StartToCloseTimeout = time.Minute
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	return workflow.ExecuteActivity(ctx, activityName, workflowID, runID)
}

// fooServiceActivities provides activities that can be used to interact with a(n) FooService service's workflow, queries, signals, and updates across namespaces
type fooServiceActivities struct {
	client patch.FooServiceClient
}

// CancelWorkflow cancels an existing workflow execution
func (a *fooServiceActivities) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	return a.client.CancelWorkflow(ctx, workflowID, runID)
}

// Foo executes a(n) test.patch.FooService.Foo workflow via an activity
func (a *fooServiceActivities) Foo(ctx context.Context, input *xnsv1.WorkflowRequest) (resp *patch.FooOutput, err error) {
	// unmarshal workflow request
	var req patch.FooInput
	if err := input.Request.UnmarshalTo(&req); err != nil {
		return nil, fooServiceOptions.convertError(temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("error unmarshalling workflow request of type %s as github.com/cludden/protoc-gen-go-temporal/gen/test/patch.FooInput", input.Request.GetTypeUrl()),
			"InvalidArgument",
			err,
		))
	}

	// initialize workflow execution
	var run patch.FooRun
	run, err = a.client.FooAsync(ctx, &req, patch.NewFooOptions().WithStartWorkflowOptions(
		xns.UnmarshalStartWorkflowOptions(input.GetStartWorkflowOptions()),
	))
	if err != nil {
		return nil, fooServiceOptions.convertError(err)
	}

	// exit early if detached enabled
	if input.GetDetached() {
		return nil, nil
	}

	// otherwise, wait for execution to complete in child goroutine
	doneCh := make(chan struct{})
	go func() {
		resp, err = run.Get(ctx)
		close(doneCh)
	}()

	heartbeatInterval := input.GetHeartbeatInterval().AsDuration()
	if heartbeatInterval == 0 {
		heartbeatInterval = time.Second * 30
	}

	// heartbeat activity while waiting for workflow execution to complete
	for {
		select {
		// send heartbeats periodically
		case <-time.After(heartbeatInterval):
			activity.RecordHeartbeat(ctx, run.ID())

		// return retryable error on worker close
		case <-activity.GetWorkerStopChannel(ctx):
			return nil, temporal.NewApplicationError("worker is stopping", "WorkerStopped")

		// catch parent activity context cancellation. in most cases, this should indicate a
		// server-sent cancellation, but there's a non-zero possibility that this cancellation
		// is received due to the worker stopping, prior to detecting the closing of the worker
		// stop channel. to give us an opportunity to detect a cancellation stemming from the
		// worker closing, we again check to see if the worker stop channel is closed before
		// propagating the cancellation
		case <-ctx.Done():
			select {
			case <-activity.GetWorkerStopChannel(ctx):
				return nil, temporal.NewApplicationError("worker is stopping", "WorkerStopped")
			default:
				parentClosePolicy := input.GetParentClosePolicy()
				if parentClosePolicy == temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL || parentClosePolicy == temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE {
					disconnectedCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
					defer cancel()
					if parentClosePolicy == temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL {
						err = run.Cancel(disconnectedCtx)
					} else {
						err = run.Terminate(disconnectedCtx, "xns activity cancellation received", "error", ctx.Err())
					}
					if err != nil {
						return nil, fooServiceOptions.convertError(err)
					}
				}
				return nil, fooServiceOptions.convertError(temporal.NewCanceledError(ctx.Err().Error()))
			}

		// handle workflow completion
		case <-doneCh:
			return resp, fooServiceOptions.convertError(err)
		}
	}
}
