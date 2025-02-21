// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 0.0.1-next (e330e8ea8508c34826279253890ab55fac39c457)
//	go go1.24.0
//	protoc (unknown)
//
// source: test/xnserr/v1/xnserr.proto
package xnserrv1xns

import (
	"context"
	"errors"
	"fmt"
	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	xnsv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/xns/v1"
	v1 "github.com/cludden/protoc-gen-go-temporal/gen/test/xnserr/v1"
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

// ServerOptions is used to configure test.xnserr.v1.Server xns activity registration
type ServerOptions struct {
	// errorConverter is used to customize error
	errorConverter func(error) error
	// filter is used to filter xns activity registrations. It receives as
	// input the original activity name, and should return one of the following:
	// 1. the original activity name, for no changes
	// 2. a modified activity name, to override the original activity name
	// 3. an empty string, to skip registration
	filter func(string) string
}

// NewServerOptions initializes a new ServerOptions value
func NewServerOptions() *ServerOptions {
	return &ServerOptions{}
}

// WithErrorConverter overrides the default error converter applied to xns activity errors
func (opts *ServerOptions) WithErrorConverter(errorConverter func(error) error) *ServerOptions {
	opts.errorConverter = errorConverter
	return opts
}

// Filter is used to filter registered xns activities or customize their name
func (opts *ServerOptions) WithFilter(filter func(string) string) *ServerOptions {
	opts.filter = filter
	return opts
}

// convertError is applied to all xns activity errors
func (opts *ServerOptions) convertError(err error) error {
	if err == nil {
		return nil
	}
	if opts != nil && opts.errorConverter != nil {
		return opts.errorConverter(err)
	}
	return xns.ErrorToApplicationError(err)
}

// filterActivity is used to filter xns activity registrations
func (opts *ServerOptions) filterActivity(name string) string {
	if opts == nil || opts.filter == nil {
		return name
	}
	return opts.filter(name)
}

// serverOptions is a reference to the ServerOptions initialized at registration
var serverOptions *ServerOptions

// RegisterServerActivities registers test.xnserr.v1.Server cross-namespace activities
func RegisterServerActivities(r worker.ActivityRegistry, c v1.ServerClient, options ...*ServerOptions) {
	if serverOptions == nil && len(options) > 0 && options[0] != nil {
		serverOptions = options[0]
	}
	a := &serverActivities{c}
	if name := serverOptions.filterActivity("test.xnserr.v1.Server.CancelWorkflow"); name != "" {
		r.RegisterActivityWithOptions(a.CancelWorkflow, activity.RegisterOptions{Name: name})
	}
	if name := serverOptions.filterActivity(v1.SleepWorkflowName); name != "" {
		r.RegisterActivityWithOptions(a.Sleep, activity.RegisterOptions{Name: name})
	}
}

// SleepWorkflowOptions are used to configure a(n) test.xnserr.v1.Server.Sleep workflow execution
type SleepWorkflowOptions struct {
	ActivityOptions      *workflow.ActivityOptions
	Detached             bool
	HeartbeatInterval    time.Duration
	ParentClosePolicy    enumsv1.ParentClosePolicy
	StartWorkflowOptions *client.StartWorkflowOptions
}

// NewSleepWorkflowOptions initializes a new SleepWorkflowOptions value
func NewSleepWorkflowOptions() *SleepWorkflowOptions {
	return &SleepWorkflowOptions{}
}

// WithActivityOptions can be used to customize the activity options
func (opts *SleepWorkflowOptions) WithActivityOptions(ao workflow.ActivityOptions) *SleepWorkflowOptions {
	opts.ActivityOptions = &ao
	return opts
}

// WithDetached can be used to start a workflow execution and exit immediately
func (opts *SleepWorkflowOptions) WithDetached(d bool) *SleepWorkflowOptions {
	opts.Detached = d
	return opts
}

// WithHeartbeatInterval can be used to customize the activity heartbeat interval
func (opts *SleepWorkflowOptions) WithHeartbeatInterval(d time.Duration) *SleepWorkflowOptions {
	opts.HeartbeatInterval = d
	return opts
}

// WithParentClosePolicy can be used to customize the cancellation propagation behavior
func (opts *SleepWorkflowOptions) WithParentClosePolicy(policy enumsv1.ParentClosePolicy) *SleepWorkflowOptions {
	opts.ParentClosePolicy = policy
	return opts
}

// WithStartWorkflowOptions can be used to customize the start workflow options
func (opts *SleepWorkflowOptions) WithStartWorkflow(swo client.StartWorkflowOptions) *SleepWorkflowOptions {
	opts.StartWorkflowOptions = &swo
	return opts
}

// SleepRun provides a handle to a test.xnserr.v1.Server.Sleep workflow execution
type SleepRun interface {
	// Cancel cancels the workflow
	Cancel(workflow.Context) error

	// Future returns the inner workflow.Future
	Future() workflow.Future

	// Get returns the inner workflow.Future
	Get(workflow.Context) error

	// ID returns the workflow id
	ID() string
}

// sleepRun provides a(n) SleepRun implementation
type sleepRun struct {
	cancel func()
	future workflow.Future
	id     string
}

// Cancel the underlying workflow execution
func (r *sleepRun) Cancel(ctx workflow.Context) error {
	if r.cancel != nil {
		r.cancel()
		if err := r.Get(ctx); err != nil && !errors.Is(err, workflow.ErrCanceled) {
			return err
		}
		return nil
	}
	return CancelServerWorkflow(ctx, r.id, "")
}

// Future returns the underlying activity future
func (r *sleepRun) Future() workflow.Future {
	return r.future
}

// Get blocks on activity completion and returns the underlying workflow result
func (r *sleepRun) Get(ctx workflow.Context) error {
	if err := r.future.Get(ctx, nil); err != nil {
		return err
	}
	return nil
}

// ID returns the underlying workflow id
func (r *sleepRun) ID() string {
	return r.id
}

// Sleep executes a(n) test.xnserr.v1.Server.Sleep workflow and blocks until error or response is received
func Sleep(ctx workflow.Context, req *v1.SleepRequest, opts ...*SleepWorkflowOptions) error {
	run, err := SleepAsync(ctx, req, opts...)
	if err != nil {
		return err
	}
	return run.Get(ctx)
}

// SleepAsync executes a(n) test.xnserr.v1.Server.Sleep workflow and returns a handle to the underlying activity
func SleepAsync(ctx workflow.Context, req *v1.SleepRequest, opts ...*SleepWorkflowOptions) (SleepRun, error) {
	activityName := serverOptions.filterActivity(v1.SleepWorkflowName)
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", v1.SleepWorkflowName),
			"Unimplemented",
			nil,
		)
	}

	opt := &SleepWorkflowOptions{}
	if len(opts) > 0 && opts[0] != nil {
		opt = opts[0]
	}
	if opt.HeartbeatInterval == 0 {
		opt.HeartbeatInterval = 10000000000 // 10 seconds
	}

	// configure activity options
	ao := workflow.GetActivityOptions(ctx)
	if opt.ActivityOptions != nil {
		ao = *opt.ActivityOptions
	}
	if ao.HeartbeatTimeout == 0 {
		ao.HeartbeatTimeout = 30000000000 // 30 seconds
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

	parentClosePolicy := temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL
	switch opt.ParentClosePolicy {
	case enumsv1.PARENT_CLOSE_POLICY_ABANDON:
		parentClosePolicy = temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_ABANDON
	case enumsv1.PARENT_CLOSE_POLICY_REQUEST_CANCEL:
		parentClosePolicy = temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL
	case enumsv1.PARENT_CLOSE_POLICY_TERMINATE:
		parentClosePolicy = temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE
	}

	ctx, cancel := workflow.WithCancel(ctx)
	return &sleepRun{
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

// CancelServerWorkflow cancels an existing workflow
func CancelServerWorkflow(ctx workflow.Context, workflowID string, runID string) error {
	return CancelServerWorkflowAsync(ctx, workflowID, runID).Get(ctx, nil)
}

// CancelServerWorkflowAsync cancels an existing workflow
func CancelServerWorkflowAsync(ctx workflow.Context, workflowID string, runID string) workflow.Future {
	activityName := serverOptions.filterActivity("test.xnserr.v1.Server.CancelWorkflow")
	if activityName == "" {
		f, s := workflow.NewFuture(ctx)
		s.SetError(temporal.NewNonRetryableApplicationError(
			"no activity registered for test.xnserr.v1.Server.CancelWorkflow",
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

// serverActivities provides activities that can be used to interact with a(n) Server service's workflow, queries, signals, and updates across namespaces
type serverActivities struct {
	client v1.ServerClient
}

// CancelWorkflow cancels an existing workflow execution
func (a *serverActivities) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	return a.client.CancelWorkflow(ctx, workflowID, runID)
}

// Sleep executes a(n) test.xnserr.v1.Server.Sleep workflow via an activity
func (a *serverActivities) Sleep(ctx context.Context, input *xnsv1.WorkflowRequest) (err error) {
	// unmarshal workflow request
	var req v1.SleepRequest
	if err := input.Request.UnmarshalTo(&req); err != nil {
		return serverOptions.convertError(temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("error unmarshalling workflow request of type %s as github.com/cludden/protoc-gen-go-temporal/gen/test/xnserr/v1.SleepRequest", input.Request.GetTypeUrl()),
			"InvalidArgument",
			err,
		))
	}

	// initialize workflow execution
	var run v1.SleepRun
	run, err = a.client.SleepAsync(ctx, &req, v1.NewSleepOptions().WithStartWorkflowOptions(
		xns.UnmarshalStartWorkflowOptions(input.GetStartWorkflowOptions()),
	))
	if err != nil {
		return serverOptions.convertError(err)
	}

	// exit early if detached enabled
	if input.GetDetached() {
		return nil
	}

	// otherwise, wait for execution to complete in child goroutine
	doneCh := make(chan struct{})
	go func() {
		err = run.Get(ctx)
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
			return temporal.NewApplicationError("worker is stopping", "WorkerStopped")

		// catch parent activity context cancellation. in most cases, this should indicate a
		// server-sent cancellation, but there's a non-zero possibility that this cancellation
		// is received due to the worker stopping, prior to detecting the closing of the worker
		// stop channel. to give us an opportunity to detect a cancellation stemming from the
		// worker closing, we again check to see if the worker stop channel is closed before
		// propagating the cancellation
		case <-ctx.Done():
			select {
			case <-activity.GetWorkerStopChannel(ctx):
				return temporal.NewApplicationError("worker is stopping", "WorkerStopped")
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
						return serverOptions.convertError(err)
					}
				}
				return serverOptions.convertError(temporal.NewCanceledError(ctx.Err().Error()))
			}

		// handle workflow completion
		case <-doneCh:
			return serverOptions.convertError(err)
		}
	}
}

// ClientOptions is used to configure test.xnserr.v1.Client xns activity registration
type ClientOptions struct {
	// errorConverter is used to customize error
	errorConverter func(error) error
	// filter is used to filter xns activity registrations. It receives as
	// input the original activity name, and should return one of the following:
	// 1. the original activity name, for no changes
	// 2. a modified activity name, to override the original activity name
	// 3. an empty string, to skip registration
	filter func(string) string
}

// NewClientOptions initializes a new ClientOptions value
func NewClientOptions() *ClientOptions {
	return &ClientOptions{}
}

// WithErrorConverter overrides the default error converter applied to xns activity errors
func (opts *ClientOptions) WithErrorConverter(errorConverter func(error) error) *ClientOptions {
	opts.errorConverter = errorConverter
	return opts
}

// Filter is used to filter registered xns activities or customize their name
func (opts *ClientOptions) WithFilter(filter func(string) string) *ClientOptions {
	opts.filter = filter
	return opts
}

// convertError is applied to all xns activity errors
func (opts *ClientOptions) convertError(err error) error {
	if err == nil {
		return nil
	}
	if opts != nil && opts.errorConverter != nil {
		return opts.errorConverter(err)
	}
	return xns.ErrorToApplicationError(err)
}

// filterActivity is used to filter xns activity registrations
func (opts *ClientOptions) filterActivity(name string) string {
	if opts == nil || opts.filter == nil {
		return name
	}
	return opts.filter(name)
}

// clientOptions is a reference to the ClientOptions initialized at registration
var clientOptions *ClientOptions

// RegisterClientActivities registers test.xnserr.v1.Client cross-namespace activities
func RegisterClientActivities(r worker.ActivityRegistry, c v1.ClientClient, options ...*ClientOptions) {
	if clientOptions == nil && len(options) > 0 && options[0] != nil {
		clientOptions = options[0]
	}
	a := &clientActivities{c}
	if name := clientOptions.filterActivity("test.xnserr.v1.Client.CancelWorkflow"); name != "" {
		r.RegisterActivityWithOptions(a.CancelWorkflow, activity.RegisterOptions{Name: name})
	}
	if name := clientOptions.filterActivity(v1.CallSleepWorkflowName); name != "" {
		r.RegisterActivityWithOptions(a.CallSleep, activity.RegisterOptions{Name: name})
	}
}

// CallSleepWorkflowOptions are used to configure a(n) test.xnserr.v1.Client.CallSleep workflow execution
type CallSleepWorkflowOptions struct {
	ActivityOptions      *workflow.ActivityOptions
	Detached             bool
	HeartbeatInterval    time.Duration
	ParentClosePolicy    enumsv1.ParentClosePolicy
	StartWorkflowOptions *client.StartWorkflowOptions
}

// NewCallSleepWorkflowOptions initializes a new CallSleepWorkflowOptions value
func NewCallSleepWorkflowOptions() *CallSleepWorkflowOptions {
	return &CallSleepWorkflowOptions{}
}

// WithActivityOptions can be used to customize the activity options
func (opts *CallSleepWorkflowOptions) WithActivityOptions(ao workflow.ActivityOptions) *CallSleepWorkflowOptions {
	opts.ActivityOptions = &ao
	return opts
}

// WithDetached can be used to start a workflow execution and exit immediately
func (opts *CallSleepWorkflowOptions) WithDetached(d bool) *CallSleepWorkflowOptions {
	opts.Detached = d
	return opts
}

// WithHeartbeatInterval can be used to customize the activity heartbeat interval
func (opts *CallSleepWorkflowOptions) WithHeartbeatInterval(d time.Duration) *CallSleepWorkflowOptions {
	opts.HeartbeatInterval = d
	return opts
}

// WithParentClosePolicy can be used to customize the cancellation propagation behavior
func (opts *CallSleepWorkflowOptions) WithParentClosePolicy(policy enumsv1.ParentClosePolicy) *CallSleepWorkflowOptions {
	opts.ParentClosePolicy = policy
	return opts
}

// WithStartWorkflowOptions can be used to customize the start workflow options
func (opts *CallSleepWorkflowOptions) WithStartWorkflow(swo client.StartWorkflowOptions) *CallSleepWorkflowOptions {
	opts.StartWorkflowOptions = &swo
	return opts
}

// CallSleepRun provides a handle to a test.xnserr.v1.Client.CallSleep workflow execution
type CallSleepRun interface {
	// Cancel cancels the workflow
	Cancel(workflow.Context) error

	// Future returns the inner workflow.Future
	Future() workflow.Future

	// Get returns the inner workflow.Future
	Get(workflow.Context) error

	// ID returns the workflow id
	ID() string
}

// callSleepRun provides a(n) CallSleepRun implementation
type callSleepRun struct {
	cancel func()
	future workflow.Future
	id     string
}

// Cancel the underlying workflow execution
func (r *callSleepRun) Cancel(ctx workflow.Context) error {
	if r.cancel != nil {
		r.cancel()
		if err := r.Get(ctx); err != nil && !errors.Is(err, workflow.ErrCanceled) {
			return err
		}
		return nil
	}
	return CancelClientWorkflow(ctx, r.id, "")
}

// Future returns the underlying activity future
func (r *callSleepRun) Future() workflow.Future {
	return r.future
}

// Get blocks on activity completion and returns the underlying workflow result
func (r *callSleepRun) Get(ctx workflow.Context) error {
	if err := r.future.Get(ctx, nil); err != nil {
		return err
	}
	return nil
}

// ID returns the underlying workflow id
func (r *callSleepRun) ID() string {
	return r.id
}

// CallSleep executes a(n) test.xnserr.v1.Client.CallSleep workflow and blocks until error or response is received
func CallSleep(ctx workflow.Context, req *v1.CallSleepRequest, opts ...*CallSleepWorkflowOptions) error {
	run, err := CallSleepAsync(ctx, req, opts...)
	if err != nil {
		return err
	}
	return run.Get(ctx)
}

// CallSleepAsync executes a(n) test.xnserr.v1.Client.CallSleep workflow and returns a handle to the underlying activity
func CallSleepAsync(ctx workflow.Context, req *v1.CallSleepRequest, opts ...*CallSleepWorkflowOptions) (CallSleepRun, error) {
	activityName := clientOptions.filterActivity(v1.CallSleepWorkflowName)
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", v1.CallSleepWorkflowName),
			"Unimplemented",
			nil,
		)
	}

	opt := &CallSleepWorkflowOptions{}
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
	return &callSleepRun{
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

// CancelClientWorkflow cancels an existing workflow
func CancelClientWorkflow(ctx workflow.Context, workflowID string, runID string) error {
	return CancelClientWorkflowAsync(ctx, workflowID, runID).Get(ctx, nil)
}

// CancelClientWorkflowAsync cancels an existing workflow
func CancelClientWorkflowAsync(ctx workflow.Context, workflowID string, runID string) workflow.Future {
	activityName := clientOptions.filterActivity("test.xnserr.v1.Client.CancelWorkflow")
	if activityName == "" {
		f, s := workflow.NewFuture(ctx)
		s.SetError(temporal.NewNonRetryableApplicationError(
			"no activity registered for test.xnserr.v1.Client.CancelWorkflow",
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

// clientActivities provides activities that can be used to interact with a(n) Client service's workflow, queries, signals, and updates across namespaces
type clientActivities struct {
	client v1.ClientClient
}

// CancelWorkflow cancels an existing workflow execution
func (a *clientActivities) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	return a.client.CancelWorkflow(ctx, workflowID, runID)
}

// CallSleep executes a(n) test.xnserr.v1.Client.CallSleep workflow via an activity
func (a *clientActivities) CallSleep(ctx context.Context, input *xnsv1.WorkflowRequest) (err error) {
	// unmarshal workflow request
	var req v1.CallSleepRequest
	if err := input.Request.UnmarshalTo(&req); err != nil {
		return clientOptions.convertError(temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("error unmarshalling workflow request of type %s as github.com/cludden/protoc-gen-go-temporal/gen/test/xnserr/v1.CallSleepRequest", input.Request.GetTypeUrl()),
			"InvalidArgument",
			err,
		))
	}

	// initialize workflow execution
	var run v1.CallSleepRun
	run, err = a.client.CallSleepAsync(ctx, &req, v1.NewCallSleepOptions().WithStartWorkflowOptions(
		xns.UnmarshalStartWorkflowOptions(input.GetStartWorkflowOptions()),
	))
	if err != nil {
		return clientOptions.convertError(err)
	}

	// exit early if detached enabled
	if input.GetDetached() {
		return nil
	}

	// otherwise, wait for execution to complete in child goroutine
	doneCh := make(chan struct{})
	go func() {
		err = run.Get(ctx)
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
			return temporal.NewApplicationError("worker is stopping", "WorkerStopped")

		// catch parent activity context cancellation. in most cases, this should indicate a
		// server-sent cancellation, but there's a non-zero possibility that this cancellation
		// is received due to the worker stopping, prior to detecting the closing of the worker
		// stop channel. to give us an opportunity to detect a cancellation stemming from the
		// worker closing, we again check to see if the worker stop channel is closed before
		// propagating the cancellation
		case <-ctx.Done():
			select {
			case <-activity.GetWorkerStopChannel(ctx):
				return temporal.NewApplicationError("worker is stopping", "WorkerStopped")
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
						return clientOptions.convertError(err)
					}
				}
				return clientOptions.convertError(temporal.NewCanceledError(ctx.Err().Error()))
			}

		// handle workflow completion
		case <-doneCh:
			return clientOptions.convertError(err)
		}
	}
}
