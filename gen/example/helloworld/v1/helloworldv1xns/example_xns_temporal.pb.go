// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 0.0.1-next (c119af1d2fcff0c4d02fcb42bcb8179a0945b8ad)
//	go go1.24.0
//	protoc (unknown)
//
// source: example/helloworld/v1/example.proto
package helloworldv1xns

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/cludden/protoc-gen-go-temporal/gen/example/helloworld/v1"
	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	xnsv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/xns/v1"
	expression "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
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

// ExampleOptions is used to configure example.helloworld.v1.Example xns activity registration
type ExampleOptions struct {
	// errorConverter is used to customize error
	errorConverter func(error) error
	// filter is used to filter xns activity registrations. It receives as
	// input the original activity name, and should return one of the following:
	// 1. the original activity name, for no changes
	// 2. a modified activity name, to override the original activity name
	// 3. an empty string, to skip registration
	filter func(string) string
}

// NewExampleOptions initializes a new ExampleOptions value
func NewExampleOptions() *ExampleOptions {
	return &ExampleOptions{}
}

// WithErrorConverter overrides the default error converter applied to xns activity errors
func (opts *ExampleOptions) WithErrorConverter(errorConverter func(error) error) *ExampleOptions {
	opts.errorConverter = errorConverter
	return opts
}

// Filter is used to filter registered xns activities or customize their name
func (opts *ExampleOptions) WithFilter(filter func(string) string) *ExampleOptions {
	opts.filter = filter
	return opts
}

// convertError is applied to all xns activity errors
func (opts *ExampleOptions) convertError(err error) error {
	if err == nil {
		return nil
	}
	if opts != nil && opts.errorConverter != nil {
		return opts.errorConverter(err)
	}
	return xns.ErrorToApplicationError(err)
}

// filterActivity is used to filter xns activity registrations
func (opts *ExampleOptions) filterActivity(name string) string {
	if opts == nil || opts.filter == nil {
		return name
	}
	return opts.filter(name)
}

// exampleOptions is a reference to the ExampleOptions initialized at registration
var exampleOptions *ExampleOptions

// RegisterExampleActivities registers example.helloworld.v1.Example cross-namespace activities
func RegisterExampleActivities(r worker.ActivityRegistry, c v1.ExampleClient, options ...*ExampleOptions) {
	if exampleOptions == nil && len(options) > 0 && options[0] != nil {
		exampleOptions = options[0]
	}
	a := &exampleActivities{c}
	if name := exampleOptions.filterActivity("example.helloworld.v1.Example.CancelWorkflow"); name != "" {
		r.RegisterActivityWithOptions(a.CancelWorkflow, activity.RegisterOptions{Name: name})
	}
	if name := exampleOptions.filterActivity(v1.HelloWorkflowName); name != "" {
		r.RegisterActivityWithOptions(a.Hello, activity.RegisterOptions{Name: name})
	}
	if name := exampleOptions.filterActivity(v1.GoodbyeSignalName); name != "" {
		r.RegisterActivityWithOptions(a.Goodbye, activity.RegisterOptions{Name: name})
	}
}

// HelloWorkflowOptions are used to configure a(n) example.v1.Hello workflow execution
type HelloWorkflowOptions struct {
	ActivityOptions      *workflow.ActivityOptions
	Detached             bool
	HeartbeatInterval    time.Duration
	ParentClosePolicy    enumsv1.ParentClosePolicy
	StartWorkflowOptions *client.StartWorkflowOptions
}

// NewHelloWorkflowOptions initializes a new HelloWorkflowOptions value
func NewHelloWorkflowOptions() *HelloWorkflowOptions {
	return &HelloWorkflowOptions{}
}

// WithActivityOptions can be used to customize the activity options
func (opts *HelloWorkflowOptions) WithActivityOptions(ao workflow.ActivityOptions) *HelloWorkflowOptions {
	opts.ActivityOptions = &ao
	return opts
}

// WithDetached can be used to start a workflow execution and exit immediately
func (opts *HelloWorkflowOptions) WithDetached(d bool) *HelloWorkflowOptions {
	opts.Detached = d
	return opts
}

// WithHeartbeatInterval can be used to customize the activity heartbeat interval
func (opts *HelloWorkflowOptions) WithHeartbeatInterval(d time.Duration) *HelloWorkflowOptions {
	opts.HeartbeatInterval = d
	return opts
}

// WithParentClosePolicy can be used to customize the cancellation propagation behavior
func (opts *HelloWorkflowOptions) WithParentClosePolicy(policy enumsv1.ParentClosePolicy) *HelloWorkflowOptions {
	opts.ParentClosePolicy = policy
	return opts
}

// WithStartWorkflowOptions can be used to customize the start workflow options
func (opts *HelloWorkflowOptions) WithStartWorkflow(swo client.StartWorkflowOptions) *HelloWorkflowOptions {
	opts.StartWorkflowOptions = &swo
	return opts
}

// HelloRun provides a handle to a example.v1.Hello workflow execution
type HelloRun interface {
	// Cancel cancels the workflow
	Cancel(workflow.Context) error

	// Future returns the inner workflow.Future
	Future() workflow.Future

	// Get returns the inner workflow.Future
	Get(workflow.Context) (*v1.HelloResponse, error)

	// ID returns the workflow id
	ID() string

	// Goodbye signals a running workflow to exit
	Goodbye(workflow.Context, *v1.GoodbyeRequest, ...*GoodbyeSignalOptions) error

	// Goodbye signals a running workflow to exit
	GoodbyeAsync(workflow.Context, *v1.GoodbyeRequest, ...*GoodbyeSignalOptions) (GoodbyeSignalHandle, error)
}

// helloRun provides a(n) HelloRun implementation
type helloRun struct {
	cancel func()
	future workflow.Future
	id     string
}

// Cancel the underlying workflow execution
func (r *helloRun) Cancel(ctx workflow.Context) error {
	if r.cancel != nil {
		r.cancel()
		if _, err := r.Get(ctx); err != nil && !errors.Is(err, workflow.ErrCanceled) {
			return err
		}
		return nil
	}
	return CancelExampleWorkflow(ctx, r.id, "")
}

// Future returns the underlying activity future
func (r *helloRun) Future() workflow.Future {
	return r.future
}

// Get blocks on activity completion and returns the underlying workflow result
func (r *helloRun) Get(ctx workflow.Context) (*v1.HelloResponse, error) {
	var resp v1.HelloResponse
	if err := r.future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ID returns the underlying workflow id
func (r *helloRun) ID() string {
	return r.id
}

// Goodbye signals a running workflow to exit
func (r *helloRun) Goodbye(ctx workflow.Context, req *v1.GoodbyeRequest, opts ...*GoodbyeSignalOptions) error {
	return Goodbye(ctx, r.ID(), "", req, opts...)
}

// Goodbye signals a running workflow to exit
func (r *helloRun) GoodbyeAsync(ctx workflow.Context, req *v1.GoodbyeRequest, opts ...*GoodbyeSignalOptions) (GoodbyeSignalHandle, error) {
	return GoodbyeAsync(ctx, r.ID(), "", req, opts...)
}

// Hello prints a friendly greeting and waits for goodbye
func Hello(ctx workflow.Context, req *v1.HelloRequest, opts ...*HelloWorkflowOptions) (*v1.HelloResponse, error) {
	run, err := HelloAsync(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// Hello prints a friendly greeting and waits for goodbye
func HelloAsync(ctx workflow.Context, req *v1.HelloRequest, opts ...*HelloWorkflowOptions) (HelloRun, error) {
	activityName := exampleOptions.filterActivity(v1.HelloWorkflowName)
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", v1.HelloWorkflowName),
			"Unimplemented",
			nil,
		)
	}

	opt := &HelloWorkflowOptions{}
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
			id, err := expression.EvalExpression(v1.HelloIdexpression, req.ProtoReflect())
			if err != nil {
				workflow.GetLogger(ctx).Error("error evaluating id expression for \"example.helloworld.v1.Example.Hello\" workflow", "error", err)
				return nil
			}
			return id
		}).Get(&wo.ID); err != nil {
			return nil, err
		}
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
	return &helloRun{
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

// GoodbyeSignalOptions are used to configure a(n) example.helloworld.v1.Example.Goodbye signal execution
type GoodbyeSignalOptions struct {
	ActivityOptions   *workflow.ActivityOptions
	HeartbeatInterval time.Duration
}

// NewGoodbyeSignalOptions initializes a new GoodbyeSignalOptions value
func NewGoodbyeSignalOptions() *GoodbyeSignalOptions {
	return &GoodbyeSignalOptions{}
}

// WithActivityOptions can be used to customize the activity options
func (opts *GoodbyeSignalOptions) WithActivityOptions(ao workflow.ActivityOptions) *GoodbyeSignalOptions {
	opts.ActivityOptions = &ao
	return opts
}

// WithHeartbeatInterval can be used to customize the activity heartbeat interval
func (opts *GoodbyeSignalOptions) WithHeartbeatInterval(d time.Duration) *GoodbyeSignalOptions {
	opts.HeartbeatInterval = d
	return opts
}

// GoodbyeSignalHandle provides a handle for a example.helloworld.v1.Example.Goodbye signal activity
type GoodbyeSignalHandle interface {
	// Cancel cancels the workflow
	Cancel(workflow.Context) error
	// Future returns the inner workflow.Future
	Future() workflow.Future
	// Get returns the inner workflow.Future
	Get(workflow.Context) error
}

// goodbyeSignalHandle provides a(n) GoodbyeQueryHandle implementation
type goodbyeSignalHandle struct {
	cancel func()
	future workflow.Future
}

// Cancel the underlying signal activity
func (r *goodbyeSignalHandle) Cancel(ctx workflow.Context) error {
	r.cancel()
	if err := r.Get(ctx); err != nil && !errors.Is(err, workflow.ErrCanceled) {
		return err
	}
	return nil
}

// Future returns the underlying activity future
func (r *goodbyeSignalHandle) Future() workflow.Future {
	return r.future
}

// Get blocks on activity completion
func (r *goodbyeSignalHandle) Get(ctx workflow.Context) error {
	return r.future.Get(ctx, nil)
}

// Goodbye signals a running workflow to exit
func Goodbye(ctx workflow.Context, workflowID string, runID string, req *v1.GoodbyeRequest, opts ...*GoodbyeSignalOptions) error {
	handle, err := GoodbyeAsync(ctx, workflowID, runID, req, opts...)
	if err != nil {
		return err
	}
	return handle.Get(ctx)
}

// GoodbyeAsync executes a(n) example.helloworld.v1.Example.Goodbye signal
func GoodbyeAsync(ctx workflow.Context, workflowID string, runID string, req *v1.GoodbyeRequest, opts ...*GoodbyeSignalOptions) (GoodbyeSignalHandle, error) {
	activityName := exampleOptions.filterActivity(v1.GoodbyeSignalName)
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", v1.GoodbyeSignalName),
			"Unimplemented",
			nil,
		)
	}

	opt := &GoodbyeSignalOptions{}
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
		ao.ScheduleToCloseTimeout = 60000000000 // 1 minute
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// marshal workflow request
	wreq, err := anypb.New(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling workflow request: %w", err)
	}

	ctx, cancel := workflow.WithCancel(ctx)
	return &goodbyeSignalHandle{
		cancel: cancel,
		future: workflow.ExecuteActivity(ctx, activityName, &xnsv1.SignalRequest{
			HeartbeatInterval: durationpb.New(opt.HeartbeatInterval),
			WorkflowId:        workflowID,
			RunId:             runID,
			Request:           wreq,
		}),
	}, nil
}

// CancelExampleWorkflow cancels an existing workflow
func CancelExampleWorkflow(ctx workflow.Context, workflowID string, runID string) error {
	return CancelExampleWorkflowAsync(ctx, workflowID, runID).Get(ctx, nil)
}

// CancelExampleWorkflowAsync cancels an existing workflow
func CancelExampleWorkflowAsync(ctx workflow.Context, workflowID string, runID string) workflow.Future {
	activityName := exampleOptions.filterActivity("example.helloworld.v1.Example.CancelWorkflow")
	if activityName == "" {
		f, s := workflow.NewFuture(ctx)
		s.SetError(temporal.NewNonRetryableApplicationError(
			"no activity registered for example.helloworld.v1.Example.CancelWorkflow",
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

// exampleActivities provides activities that can be used to interact with a(n) Example service's workflow, queries, signals, and updates across namespaces
type exampleActivities struct {
	client v1.ExampleClient
}

// CancelWorkflow cancels an existing workflow execution
func (a *exampleActivities) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	return a.client.CancelWorkflow(ctx, workflowID, runID)
}

// Hello executes a(n) example.v1.Hello workflow via an activity
func (a *exampleActivities) Hello(ctx context.Context, input *xnsv1.WorkflowRequest) (resp *v1.HelloResponse, err error) {
	// unmarshal workflow request
	var req v1.HelloRequest
	if err := input.Request.UnmarshalTo(&req); err != nil {
		return nil, exampleOptions.convertError(temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("error unmarshalling workflow request of type %s as github.com/cludden/protoc-gen-go-temporal/gen/example/helloworld/v1.HelloRequest", input.Request.GetTypeUrl()),
			"InvalidArgument",
			err,
		))
	}

	// initialize workflow execution
	var run v1.HelloRun
	run, err = a.client.HelloAsync(ctx, &req, v1.NewHelloOptions().WithStartWorkflowOptions(
		xns.UnmarshalStartWorkflowOptions(input.GetStartWorkflowOptions()),
	))
	if err != nil {
		return nil, exampleOptions.convertError(err)
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
						return nil, exampleOptions.convertError(err)
					}
				}
				return nil, exampleOptions.convertError(temporal.NewCanceledError(ctx.Err().Error()))
			}

		// handle workflow completion
		case <-doneCh:
			return resp, exampleOptions.convertError(err)
		}
	}
}

// Goodbye executes a(n) example.helloworld.v1.Example.Goodbye signal via an activity
func (a *exampleActivities) Goodbye(ctx context.Context, input *xnsv1.SignalRequest) (err error) {
	// unmarshal signal request
	var req v1.GoodbyeRequest
	if err := input.Request.UnmarshalTo(&req); err != nil {
		return exampleOptions.convertError(temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("error unmarshalling signal request of type %s as github.com/cludden/protoc-gen-go-temporal/gen/example/helloworld/v1.GoodbyeRequest", input.Request.GetTypeUrl()),
			"InvalidArgument",
			err,
		))
	}
	// execute signal in child goroutine
	doneCh := make(chan struct{})
	go func() {
		err = a.client.Goodbye(ctx, input.GetWorkflowId(), input.GetRunId(), &req)
		close(doneCh)
	}()

	heartbeatInterval := input.GetHeartbeatInterval().AsDuration()
	if heartbeatInterval == 0 {
		heartbeatInterval = time.Second * 10
	}

	// heartbeat activity while waiting for signal to complete
	for {
		select {
		case <-time.After(heartbeatInterval):
			activity.RecordHeartbeat(ctx)
		case <-ctx.Done():
			exampleOptions.convertError(ctx.Err())
		case <-doneCh:
			return exampleOptions.convertError(err)
		}
	}
}
