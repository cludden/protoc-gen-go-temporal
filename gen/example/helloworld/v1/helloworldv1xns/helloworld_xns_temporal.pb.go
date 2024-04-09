// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 1.10.4-next (7304ecfb0827e07546ab6bc6830e717a08451b63)
//	go go1.22.1
//	protoc (unknown)
//
// source: example/helloworld/v1/helloworld.proto
package helloworldv1xns

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/cludden/protoc-gen-go-temporal/gen/example/helloworld/v1"
	v11 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/xns/v1"
	expression "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	xns "github.com/cludden/protoc-gen-go-temporal/pkg/xns"
	uuid "github.com/google/uuid"
	activity "go.temporal.io/sdk/activity"
	client "go.temporal.io/sdk/client"
	temporal "go.temporal.io/sdk/temporal"
	worker "go.temporal.io/sdk/worker"
	workflow "go.temporal.io/sdk/workflow"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	"time"
)

// HelloWorldOptions is used to configure example.helloworld.v1.HelloWorld xns activity registration
type HelloWorldOptions struct {
	// errorConverter is used to customize error
	errorConverter func(error) error
	// filter is used to filter xns activity registrations. It receives as
	// input the original activity name, and should return one of the following:
	// 1. the original activity name, for no changes
	// 2. a modified activity name, to override the original activity name
	// 3. an empty string, to skip registration
	filter func(string) string
}

// NewHelloWorldOptions initializes a new HelloWorldOptions value
func NewHelloWorldOptions() *HelloWorldOptions {
	return &HelloWorldOptions{}
}

// WithErrorConverter overrides the default error converter applied to xns activity errors
func (opts *HelloWorldOptions) WithErrorConverter(errorConverter func(error) error) *HelloWorldOptions {
	opts.errorConverter = errorConverter
	return opts
}

// Filter is used to filter registered xns activities or customize their name
func (opts *HelloWorldOptions) WithFilter(filter func(string) string) *HelloWorldOptions {
	opts.filter = filter
	return opts
}

// convertError is applied to all xns activity errors
func (opts *HelloWorldOptions) convertError(err error) error {
	if err == nil {
		return nil
	}
	if opts != nil && opts.errorConverter != nil {
		return opts.errorConverter(err)
	}
	return xns.ErrorToApplicationError(err)
}

// filterActivity is used to filter xns activity registrations
func (opts *HelloWorldOptions) filterActivity(name string) string {
	if opts == nil || opts.filter == nil {
		return name
	}
	return opts.filter(name)
}

// helloWorldOptions is a reference to the HelloWorldOptions initialized at registration
var helloWorldOptions *HelloWorldOptions

// RegisterHelloWorldActivities registers example.helloworld.v1.HelloWorld cross-namespace activities
func RegisterHelloWorldActivities(r worker.ActivityRegistry, c v1.HelloWorldClient, options ...*HelloWorldOptions) {
	if helloWorldOptions == nil && len(options) > 0 && options[0] != nil {
		helloWorldOptions = options[0]
	}
	a := &helloWorldActivities{c}
	if name := helloWorldOptions.filterActivity("example.helloworld.v1.HelloWorld.CancelWorkflow"); name != "" {
		r.RegisterActivityWithOptions(a.CancelWorkflow, activity.RegisterOptions{Name: name})
	}
	if name := helloWorldOptions.filterActivity(v1.HelloWorldWorkflowName); name != "" {
		r.RegisterActivityWithOptions(a.HelloWorld, activity.RegisterOptions{Name: name})
	}
}

// HelloWorldWorkflowOptions are used to configure a(n) HelloWorld workflow execution
type HelloWorldWorkflowOptions struct {
	ActivityOptions      *workflow.ActivityOptions
	Detached             bool
	HeartbeatInterval    time.Duration
	StartWorkflowOptions *client.StartWorkflowOptions
}

// NewHelloWorldWorkflowOptions initializes a new HelloWorldWorkflowOptions value
func NewHelloWorldWorkflowOptions() *HelloWorldWorkflowOptions {
	return &HelloWorldWorkflowOptions{}
}

// WithActivityOptions can be used to customize the activity options
func (opts *HelloWorldWorkflowOptions) WithActivityOptions(ao workflow.ActivityOptions) *HelloWorldWorkflowOptions {
	opts.ActivityOptions = &ao
	return opts
}

// WithDetached can be used to start a workflow execution and exit immediately
func (opts *HelloWorldWorkflowOptions) WithDetached(d bool) *HelloWorldWorkflowOptions {
	opts.Detached = d
	return opts
}

// WithHeartbeatInterval can be used to customize the activity heartbeat interval
func (opts *HelloWorldWorkflowOptions) WithHeartbeatInterval(d time.Duration) *HelloWorldWorkflowOptions {
	opts.HeartbeatInterval = d
	return opts
}

// WithStartWorkflowOptions can be used to customize the start workflow options
func (opts *HelloWorldWorkflowOptions) WithStartWorkflow(swo client.StartWorkflowOptions) *HelloWorldWorkflowOptions {
	opts.StartWorkflowOptions = &swo
	return opts
}

// HelloWorldRun provides a handle to a HelloWorld workflow execution
type HelloWorldRun interface {
	// Cancel cancels the workflow
	Cancel(workflow.Context) error

	// Future returns the inner workflow.Future
	Future() workflow.Future

	// Get returns the inner workflow.Future
	Get(workflow.Context) (*v1.HelloWorldOutput, error)

	// ID returns the workflow id
	ID() string
}

// helloWorldRun provides a(n) HelloWorldRun implementation
type helloWorldRun struct {
	cancel func()
	future workflow.Future
	id     string
}

// Cancel the underlying workflow execution
func (r *helloWorldRun) Cancel(ctx workflow.Context) error {
	if r.cancel != nil {
		r.cancel()
		if _, err := r.Get(ctx); err != nil && !errors.Is(err, workflow.ErrCanceled) {
			return err
		}
		return nil
	}
	return CancelHelloWorldWorkflow(ctx, r.id, "")
}

// Future returns the underlying activity future
func (r *helloWorldRun) Future() workflow.Future {
	return r.future
}

// Get blocks on activity completion and returns the underlying workflow result
func (r *helloWorldRun) Get(ctx workflow.Context) (*v1.HelloWorldOutput, error) {
	var resp v1.HelloWorldOutput
	if err := r.future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ID returns the underlying workflow id
func (r *helloWorldRun) ID() string {
	return r.id
}

// HelloWorld describes a Temporal workflow and activity with the same name
// and signature
func HelloWorld(ctx workflow.Context, req *v1.HelloWorldInput, opts ...*HelloWorldWorkflowOptions) (*v1.HelloWorldOutput, error) {
	run, err := HelloWorldAsync(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// HelloWorld describes a Temporal workflow and activity with the same name
// and signature
func HelloWorldAsync(ctx workflow.Context, req *v1.HelloWorldInput, opts ...*HelloWorldWorkflowOptions) (HelloWorldRun, error) {
	activityName := helloWorldOptions.filterActivity(v1.HelloWorldWorkflowName)
	if activityName == "" {
		return nil, temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("no activity registered for %s", v1.HelloWorldWorkflowName),
			"Unimplemented",
			nil,
		)
	}

	opt := &HelloWorldWorkflowOptions{}
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
			id, err := expression.EvalExpression(v1.HelloWorldIdexpression, req.ProtoReflect())
			if err != nil {
				workflow.GetLogger(ctx).Error("error evaluating id expression for \"example.helloworld.v1.HelloWorld.HelloWorld\" workflow", "error", err)
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

	ctx, cancel := workflow.WithCancel(ctx)
	return &helloWorldRun{
		cancel: cancel,
		id:     wo.ID,
		future: workflow.ExecuteActivity(ctx, activityName, &v11.WorkflowRequest{
			Detached:             opt.Detached,
			HeartbeatInterval:    durationpb.New(opt.HeartbeatInterval),
			Request:              wreq,
			StartWorkflowOptions: swo,
		}),
	}, nil
}

// CancelHelloWorldWorkflow cancels an existing workflow
func CancelHelloWorldWorkflow(ctx workflow.Context, workflowID string, runID string) error {
	return CancelHelloWorldWorkflowAsync(ctx, workflowID, runID).Get(ctx, nil)
}

// CancelHelloWorldWorkflowAsync cancels an existing workflow
func CancelHelloWorldWorkflowAsync(ctx workflow.Context, workflowID string, runID string) workflow.Future {
	activityName := helloWorldOptions.filterActivity("example.helloworld.v1.HelloWorld.CancelWorkflow")
	if activityName == "" {
		f, s := workflow.NewFuture(ctx)
		s.SetError(temporal.NewNonRetryableApplicationError(
			"no activity registered for example.helloworld.v1.HelloWorld.CancelWorkflow",
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

// helloWorldActivities provides activities that can be used to interact with a(n) HelloWorld service's workflow, queries, signals, and updates across namespaces
type helloWorldActivities struct {
	client v1.HelloWorldClient
}

// CancelWorkflow cancels an existing workflow execution
func (a *helloWorldActivities) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	return a.client.CancelWorkflow(ctx, workflowID, runID)
}

// HelloWorld executes a(n) HelloWorld workflow via an activity
func (a *helloWorldActivities) HelloWorld(ctx context.Context, input *v11.WorkflowRequest) (resp *v1.HelloWorldOutput, err error) {
	// unmarshal workflow request
	var req v1.HelloWorldInput
	if err := input.Request.UnmarshalTo(&req); err != nil {
		return nil, helloWorldOptions.convertError(temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("error unmarshalling workflow request of type %s as github.com/cludden/protoc-gen-go-temporal/gen/example/helloworld/v1.HelloWorldInput", input.Request.GetTypeUrl()),
			"InvalidArgument",
			err,
		))
	}

	// initialize workflow execution
	var run v1.HelloWorldRun
	run, err = a.client.HelloWorldAsync(ctx, &req, v1.NewHelloWorldOptions().WithStartWorkflowOptions(
		xns.UnmarshalStartWorkflowOptions(input.GetStartWorkflowOptions()),
	))
	if err != nil {
		return nil, helloWorldOptions.convertError(err)
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
		heartbeatInterval = time.Minute
	}

	// heartbeat activity while waiting for workflow execution to complete
	for {
		select {
		case <-time.After(heartbeatInterval):
			activity.RecordHeartbeat(ctx, run.ID())
		case <-ctx.Done():
			if err := run.Cancel(ctx); err != nil {
				return nil, helloWorldOptions.convertError(err)
			}
			return nil, helloWorldOptions.convertError(workflow.ErrCanceled)
		case <-doneCh:
			return resp, helloWorldOptions.convertError(err)
		}
	}
}
