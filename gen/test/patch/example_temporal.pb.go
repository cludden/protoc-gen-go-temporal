// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 0.0.1-next (c119af1d2fcff0c4d02fcb42bcb8179a0945b8ad)
//	go go1.24.0
//	protoc (unknown)
//
// source: test/patch/example.proto
package patch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	helpers "github.com/cludden/protoc-gen-go-temporal/pkg/helpers"
	scheme "github.com/cludden/protoc-gen-go-temporal/pkg/scheme"
	gohomedir "github.com/mitchellh/go-homedir"
	v2 "github.com/urfave/cli/v2"
	enumsv1 "go.temporal.io/api/enums/v1"
	activity "go.temporal.io/sdk/activity"
	client "go.temporal.io/sdk/client"
	temporal "go.temporal.io/sdk/temporal"
	testsuite "go.temporal.io/sdk/testsuite"
	worker "go.temporal.io/sdk/worker"
	workflow "go.temporal.io/sdk/workflow"
	protojson "google.golang.org/protobuf/encoding/protojson"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"log/slog"
	"os"
	"sort"
	"time"
)

// FooServiceTaskQueue is the default task-queue for a test.patch.FooService worker
const FooServiceTaskQueue = "foo-queue"

// test.patch.FooService workflow names
const (
	FooWorkflowName = "test.patch.FooService.Foo"
)

// test.patch.FooService activity names
const (
	FooActivityName = "test.patch.FooService.Foo"
)

// FooServiceClient describes a client for a(n) test.patch.FooService worker
type FooServiceClient interface {
	// Foo executes a(n) test.patch.FooService.Foo workflow and blocks until error or response received
	Foo(ctx context.Context, req *FooInput, opts ...*FooOptions) (*FooOutput, error)

	// FooAsync starts a(n) test.patch.FooService.Foo workflow and returns a handle to the workflow run
	FooAsync(ctx context.Context, req *FooInput, opts ...*FooOptions) (FooRun, error)

	// GetFoo retrieves a handle to an existing test.patch.FooService.Foo workflow execution
	GetFoo(ctx context.Context, workflowID string, runID string) FooRun

	// CancelWorkflow requests cancellation of an existing workflow execution
	CancelWorkflow(ctx context.Context, workflowID string, runID string) error

	// TerminateWorkflow an existing workflow execution
	TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error
}

// fooServiceClient implements a temporal client for a test.patch.FooService service
type fooServiceClient struct {
	client client.Client
	log    *slog.Logger
}

// NewFooServiceClient initializes a new test.patch.FooService client
func NewFooServiceClient(c client.Client, options ...*fooServiceClientOptions) FooServiceClient {
	var cfg *fooServiceClientOptions
	if len(options) > 0 {
		cfg = options[0]
	} else {
		cfg = NewFooServiceClientOptions()
	}
	return &fooServiceClient{
		client: c,
		log:    cfg.getLogger(),
	}
}

// NewFooServiceClientWithOptions initializes a new FooService client with the given options
func NewFooServiceClientWithOptions(c client.Client, opts client.Options, options ...*fooServiceClientOptions) (FooServiceClient, error) {
	var err error
	c, err = client.NewClientFromExisting(c, opts)
	if err != nil {
		return nil, fmt.Errorf("error initializing client with options: %w", err)
	}
	var cfg *fooServiceClientOptions
	if len(options) > 0 {
		cfg = options[0]
	} else {
		cfg = NewFooServiceClientOptions()
	}
	return &fooServiceClient{
		client: c,
		log:    cfg.getLogger(),
	}, nil
}

// fooServiceClientOptions describes optional runtime configuration for a FooServiceClient
type fooServiceClientOptions struct {
	log *slog.Logger
}

// NewFooServiceClientOptions initializes a new fooServiceClientOptions value
func NewFooServiceClientOptions() *fooServiceClientOptions {
	return &fooServiceClientOptions{}
}

// WithLogger can be used to override the default logger
func (opts *fooServiceClientOptions) WithLogger(l *slog.Logger) *fooServiceClientOptions {
	if l != nil {
		opts.log = l
	}
	return opts
}

// getLogger returns the configured logger, or the default logger
func (opts *fooServiceClientOptions) getLogger() *slog.Logger {
	if opts != nil && opts.log != nil {
		return opts.log
	}
	return slog.Default()
}

// test.patch.FooService.Foo executes a test.patch.FooService.Foo workflow and blocks until error or response received
func (c *fooServiceClient) Foo(ctx context.Context, req *FooInput, options ...*FooOptions) (*FooOutput, error) {
	run, err := c.FooAsync(ctx, req, options...)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// FooAsync starts a(n) test.patch.FooService.Foo workflow and returns a handle to the workflow run
func (c *fooServiceClient) FooAsync(ctx context.Context, req *FooInput, options ...*FooOptions) (FooRun, error) {
	var o *FooOptions
	if len(options) > 0 && options[0] != nil {
		o = options[0]
	} else {
		o = NewFooOptions()
	}
	opts, err := o.Build(req.ProtoReflect())
	if err != nil {
		return nil, fmt.Errorf("error initializing client.StartWorkflowOptions: %w", err)
	}
	run, err := c.client.ExecuteWorkflow(ctx, opts, FooWorkflowName, req)
	if err != nil {
		return nil, err
	}
	if run == nil {
		return nil, errors.New("execute workflow returned nil run")
	}
	return &fooRun{
		client: c,
		run:    run,
	}, nil
}

// GetFoo fetches an existing test.patch.FooService.Foo execution
func (c *fooServiceClient) GetFoo(ctx context.Context, workflowID string, runID string) FooRun {
	return &fooRun{
		client: c,
		run:    c.client.GetWorkflow(ctx, workflowID, runID),
	}
}

// CancelWorkflow requests cancellation of an existing workflow execution
func (c *fooServiceClient) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	return c.client.CancelWorkflow(ctx, workflowID, runID)
}

// TerminateWorkflow terminates an existing workflow execution
func (c *fooServiceClient) TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error {
	return c.client.TerminateWorkflow(ctx, workflowID, runID, reason, details...)
}

// FooOptions provides configuration for a test.patch.FooService.Foo workflow operation
type FooOptions struct {
	options          client.StartWorkflowOptions
	executionTimeout *time.Duration
	id               *string
	idReusePolicy    enumsv1.WorkflowIdReusePolicy
	retryPolicy      *temporal.RetryPolicy
	runTimeout       *time.Duration
	searchAttributes map[string]any
	taskQueue        *string
	taskTimeout      *time.Duration
}

// NewFooOptions initializes a new FooOptions value
func NewFooOptions() *FooOptions {
	return &FooOptions{}
}

// Build initializes a new go.temporal.io/sdk/client.StartWorkflowOptions value with defaults and overrides applied
func (o *FooOptions) Build(req protoreflect.Message) (client.StartWorkflowOptions, error) {
	opts := o.options
	if v := o.id; v != nil {
		opts.ID = *v
	}
	if v := o.idReusePolicy; v != enumsv1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v
	}
	if v := o.taskQueue; v != nil {
		opts.TaskQueue = *v
	} else if opts.TaskQueue == "" {
		opts.TaskQueue = FooServiceTaskQueue
	}
	if v := o.retryPolicy; v != nil {
		opts.RetryPolicy = v
	}
	if v := o.searchAttributes; v != nil {
		opts.SearchAttributes = o.searchAttributes
	}
	if v := o.executionTimeout; v != nil {
		opts.WorkflowExecutionTimeout = *v
	}
	if v := o.runTimeout; v != nil {
		opts.WorkflowRunTimeout = *v
	}
	if v := o.taskTimeout; v != nil {
		opts.WorkflowTaskTimeout = *v
	}
	return opts, nil
}

// WithStartWorkflowOptions sets the initial go.temporal.io/sdk/client.StartWorkflowOptions
func (o *FooOptions) WithStartWorkflowOptions(options client.StartWorkflowOptions) *FooOptions {
	o.options = options
	return o
}

// WithExecutionTimeout sets the WorkflowExecutionTimeout value
func (o *FooOptions) WithExecutionTimeout(d time.Duration) *FooOptions {
	o.executionTimeout = &d
	return o
}

// WithID sets the ID value
func (o *FooOptions) WithID(id string) *FooOptions {
	o.id = &id
	return o
}

// WithIDReusePolicy sets the WorkflowIDReusePolicy value
func (o *FooOptions) WithIDReusePolicy(policy enumsv1.WorkflowIdReusePolicy) *FooOptions {
	o.idReusePolicy = policy
	return o
}

// WithRetryPolicy sets the RetryPolicy value
func (o *FooOptions) WithRetryPolicy(policy *temporal.RetryPolicy) *FooOptions {
	o.retryPolicy = policy
	return o
}

// WithRunTimeout sets the WorkflowRunTimeout value
func (o *FooOptions) WithRunTimeout(d time.Duration) *FooOptions {
	o.runTimeout = &d
	return o
}

// WithSearchAttributes sets the SearchAttributes value
func (o *FooOptions) WithSearchAttributes(sa map[string]any) *FooOptions {
	o.searchAttributes = sa
	return o
}

// WithTaskTimeout sets the WorkflowTaskTimeout value
func (o *FooOptions) WithTaskTimeout(d time.Duration) *FooOptions {
	o.taskTimeout = &d
	return o
}

// WithTaskQueue sets the TaskQueue value
func (o *FooOptions) WithTaskQueue(tq string) *FooOptions {
	o.taskQueue = &tq
	return o
}

// FooRun describes a(n) test.patch.FooService.Foo workflow run
type FooRun interface {
	// ID returns the workflow ID
	ID() string

	// RunID returns the workflow instance ID
	RunID() string

	// Run returns the inner client.WorkflowRun
	Run() client.WorkflowRun

	// Get blocks until the workflow is complete and returns the result
	Get(ctx context.Context) (*FooOutput, error)

	// Cancel requests cancellation of a workflow in execution, returning an error if applicable
	Cancel(ctx context.Context) error

	// Terminate terminates a workflow in execution, returning an error if applicable
	Terminate(ctx context.Context, reason string, details ...interface{}) error
}

// fooRun provides an internal implementation of a(n) FooRunRun
type fooRun struct {
	client *fooServiceClient
	run    client.WorkflowRun
}

// ID returns the workflow ID
func (r *fooRun) ID() string {
	return r.run.GetID()
}

// Run returns the inner client.WorkflowRun
func (r *fooRun) Run() client.WorkflowRun {
	return r.run
}

// RunID returns the execution ID
func (r *fooRun) RunID() string {
	return r.run.GetRunID()
}

// Cancel requests cancellation of a workflow in execution, returning an error if applicable
func (r *fooRun) Cancel(ctx context.Context) error {
	return r.client.CancelWorkflow(ctx, r.ID(), r.RunID())
}

// Get blocks until the workflow is complete, returning the result if applicable
func (r *fooRun) Get(ctx context.Context) (*FooOutput, error) {
	var resp FooOutput
	if err := r.run.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Terminate terminates a workflow in execution, returning an error if applicable
func (r *fooRun) Terminate(ctx context.Context, reason string, details ...interface{}) error {
	return r.client.TerminateWorkflow(ctx, r.ID(), r.RunID(), reason, details...)
}

// Reference to generated workflow functions
var (
	// FooFunction implements a "test.patch.FooService.Foo" workflow
	FooFunction func(workflow.Context, *FooInput) (*FooOutput, error)
)

// FooServiceWorkflowFunctions describes a mockable dependency for inlining workflows within other workflows
type (
	// FooServiceWorkflowFunctions describes a mockable dependency for inlining workflows within other workflows
	FooServiceWorkflowFunctions interface {
		// Foo executes a "test.patch.FooService.Foo" workflow inline
		Foo(workflow.Context, *FooInput) (*FooOutput, error)
	}
	// fooServiceWorkflowFunctions provides an internal FooServiceWorkflowFunctions implementation
	fooServiceWorkflowFunctions struct{}
)

func NewFooServiceWorkflowFunctions() FooServiceWorkflowFunctions {
	return &fooServiceWorkflowFunctions{}
}

// Foo executes a "test.patch.FooService.Foo" workflow inline
func (f *fooServiceWorkflowFunctions) Foo(ctx workflow.Context, req *FooInput) (*FooOutput, error) {
	if FooFunction == nil {
		return nil, errors.New("Foo requires workflow registration via RegisterFooServiceWorkflows or RegisterFooWorkflow")
	}
	return FooFunction(ctx, req)
}

// FooServiceWorkflows provides methods for initializing new test.patch.FooService workflow values
type FooServiceWorkflows interface {
	// Foo initializes a new a(n) FooWorkflow implementation
	Foo(ctx workflow.Context, input *FooWorkflowInput) (FooWorkflow, error)
}

// RegisterFooServiceWorkflows registers test.patch.FooService workflows with the given worker
func RegisterFooServiceWorkflows(r worker.WorkflowRegistry, workflows FooServiceWorkflows) {
	RegisterFooWorkflow(r, workflows.Foo)
}

// RegisterFooWorkflow registers a test.patch.FooService.Foo workflow with the given worker
func RegisterFooWorkflow(r worker.WorkflowRegistry, wf func(workflow.Context, *FooWorkflowInput) (FooWorkflow, error)) {
	FooFunction = buildFoo(wf)
	r.RegisterWorkflowWithOptions(FooFunction, workflow.RegisterOptions{Name: FooWorkflowName})
}

// buildFoo converts a Foo workflow struct into a valid workflow function
func buildFoo(ctor func(workflow.Context, *FooWorkflowInput) (FooWorkflow, error)) func(workflow.Context, *FooInput) (*FooOutput, error) {
	return func(ctx workflow.Context, req *FooInput) (*FooOutput, error) {
		input := &FooWorkflowInput{
			Req: req,
		}
		wf, err := ctor(ctx, input)
		if err != nil {
			return nil, err
		}
		if initializable, ok := wf.(helpers.Initializable); ok {
			if err := initializable.Initialize(ctx); err != nil {
				return nil, err
			}
		}
		return wf.Execute(ctx)
	}
}

// FooWorkflowInput describes the input to a(n) test.patch.FooService.Foo workflow constructor
type FooWorkflowInput struct {
	Req *FooInput
}

// FooWorkflow describes a(n) test.patch.FooService.Foo workflow implementation
type FooWorkflow interface {
	// Execute defines the entrypoint to a(n) test.patch.FooService.Foo workflow
	Execute(ctx workflow.Context) (*FooOutput, error)
}

// FooChild executes a child test.patch.FooService.Foo workflow and blocks until error or response received
func FooChild(ctx workflow.Context, req *FooInput, options ...*FooChildOptions) (*FooOutput, error) {
	childRun, err := FooChildAsync(ctx, req, options...)
	if err != nil {
		return nil, err
	}
	return childRun.Get(ctx)
}

// FooChildAsync starts a child test.patch.FooService.Foo workflow and returns a handle to the child workflow run
func FooChildAsync(ctx workflow.Context, req *FooInput, options ...*FooChildOptions) (*FooChildRun, error) {
	var o *FooChildOptions
	if len(options) > 0 && options[0] != nil {
		o = options[0]
	} else {
		o = NewFooChildOptions()
	}
	opts, err := o.Build(ctx, req.ProtoReflect())
	if err != nil {
		return nil, fmt.Errorf("error initializing workflow.ChildWorkflowOptions: %w", err)
	}
	ctx = workflow.WithChildOptions(ctx, opts)
	return &FooChildRun{Future: workflow.ExecuteChildWorkflow(ctx, FooWorkflowName, req)}, nil
}

// FooChildOptions provides configuration for a child test.patch.FooService.Foo workflow operation
type FooChildOptions struct {
	options             workflow.ChildWorkflowOptions
	executionTimeout    *time.Duration
	id                  *string
	idReusePolicy       enumsv1.WorkflowIdReusePolicy
	retryPolicy         *temporal.RetryPolicy
	runTimeout          *time.Duration
	searchAttributes    map[string]any
	taskQueue           *string
	taskTimeout         *time.Duration
	parentClosePolicy   enumsv1.ParentClosePolicy
	waitForCancellation *bool
}

// NewFooChildOptions initializes a new FooChildOptions value
func NewFooChildOptions() *FooChildOptions {
	return &FooChildOptions{}
}

// Build initializes a new go.temporal.io/sdk/workflow.ChildWorkflowOptions value with defaults and overrides applied
func (o *FooChildOptions) Build(ctx workflow.Context, req protoreflect.Message) (workflow.ChildWorkflowOptions, error) {
	opts := o.options
	if v := o.id; v != nil {
		opts.WorkflowID = *v
	}
	if v := o.idReusePolicy; v != enumsv1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v
	}
	if v := o.taskQueue; v != nil {
		opts.TaskQueue = *v
	} else if opts.TaskQueue == "" {
		opts.TaskQueue = FooServiceTaskQueue
	}
	if v := o.retryPolicy; v != nil {
		opts.RetryPolicy = v
	}
	if v := o.searchAttributes; v != nil {
		opts.SearchAttributes = o.searchAttributes
	}
	if v := o.executionTimeout; v != nil {
		opts.WorkflowExecutionTimeout = *v
	}
	if v := o.runTimeout; v != nil {
		opts.WorkflowRunTimeout = *v
	}
	if v := o.taskTimeout; v != nil {
		opts.WorkflowTaskTimeout = *v
	}
	if v := o.parentClosePolicy; v != enumsv1.PARENT_CLOSE_POLICY_UNSPECIFIED {
		opts.ParentClosePolicy = v
	}
	if v := o.waitForCancellation; v != nil {
		opts.WaitForCancellation = *v
	}
	return opts, nil
}

// WithChildWorkflowOptions sets the initial go.temporal.io/sdk/workflow.ChildWorkflowOptions
func (o *FooChildOptions) WithChildWorkflowOptions(options workflow.ChildWorkflowOptions) *FooChildOptions {
	o.options = options
	return o
}

// WithExecutionTimeout sets the WorkflowExecutionTimeout value
func (o *FooChildOptions) WithExecutionTimeout(d time.Duration) *FooChildOptions {
	o.executionTimeout = &d
	return o
}

// WithID sets the WorkflowID value
func (o *FooChildOptions) WithID(id string) *FooChildOptions {
	o.id = &id
	return o
}

// WithIDReusePolicy sets the WorkflowIDReusePolicy value
func (o *FooChildOptions) WithIDReusePolicy(policy enumsv1.WorkflowIdReusePolicy) *FooChildOptions {
	o.idReusePolicy = policy
	return o
}

// WithParentClosePolicy sets the WorkflowIDReusePolicy value
func (o *FooChildOptions) WithParentClosePolicy(policy enumsv1.ParentClosePolicy) *FooChildOptions {
	o.parentClosePolicy = policy
	return o
}

// WithRetryPolicy sets the RetryPolicy value
func (o *FooChildOptions) WithRetryPolicy(policy *temporal.RetryPolicy) *FooChildOptions {
	o.retryPolicy = policy
	return o
}

// WithRunTimeout sets the WorkflowRunTimeout value
func (o *FooChildOptions) WithRunTimeout(d time.Duration) *FooChildOptions {
	o.runTimeout = &d
	return o
}

// WithSearchAttributes sets the SearchAttributes value
func (o *FooChildOptions) WithSearchAttributes(sa map[string]any) *FooChildOptions {
	o.searchAttributes = sa
	return o
}

// WithTaskTimeout sets the WorkflowTaskTimeout value
func (o *FooChildOptions) WithTaskTimeout(d time.Duration) *FooChildOptions {
	o.taskTimeout = &d
	return o
}

// WithTaskQueue sets the TaskQueue value
func (o *FooChildOptions) WithTaskQueue(tq string) *FooChildOptions {
	o.taskQueue = &tq
	return o
}

// WithWaitForCancellation sets the WaitForCancellation value
func (o *FooChildOptions) WithWaitForCancellation(wait bool) *FooChildOptions {
	o.waitForCancellation = &wait
	return o
}

// FooChildRun describes a child Foo workflow run
type FooChildRun struct {
	Future workflow.ChildWorkflowFuture
}

// Get blocks until the workflow is completed, returning the response value
func (r *FooChildRun) Get(ctx workflow.Context) (*FooOutput, error) {
	var resp FooOutput
	if err := r.Future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Select adds this completion to the selector. Callback can be nil.
func (r *FooChildRun) Select(sel workflow.Selector, fn func(*FooChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future, func(workflow.Future) {
		if fn != nil {
			fn(r)
		}
	})
}

// SelectStart adds waiting for start to the selector. Callback can be nil.
func (r *FooChildRun) SelectStart(sel workflow.Selector, fn func(*FooChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future.GetChildWorkflowExecution(), func(workflow.Future) {
		if fn != nil {
			fn(r)
		}
	})
}

// WaitStart waits for the child workflow to start
func (r *FooChildRun) WaitStart(ctx workflow.Context) (*workflow.Execution, error) {
	var exec workflow.Execution
	if err := r.Future.GetChildWorkflowExecution().Get(ctx, &exec); err != nil {
		return nil, err
	}
	return &exec, nil
}

// FooServiceActivities describes available worker activities
type FooServiceActivities interface {
	// test.patch.FooService.Foo implements a(n) test.patch.FooService.Foo activity definition
	Foo(ctx context.Context, req *FooInput) (*FooOutput, error)
}

// RegisterFooServiceActivities registers activities with a worker
func RegisterFooServiceActivities(r worker.ActivityRegistry, activities FooServiceActivities) {
	RegisterFooActivity(r, activities.Foo)
}

// RegisterFooActivity registers a test.patch.FooService.Foo activity
func RegisterFooActivity(r worker.ActivityRegistry, fn func(context.Context, *FooInput) (*FooOutput, error)) {
	r.RegisterActivityWithOptions(fn, activity.RegisterOptions{
		Name: FooActivityName,
	})
}

// FooFuture describes a(n) test.patch.FooService.Foo activity execution
type FooFuture struct {
	Future workflow.Future
}

// Get blocks on the activity's completion, returning the response
func (f *FooFuture) Get(ctx workflow.Context) (*FooOutput, error) {
	var resp FooOutput
	if err := f.Future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Select adds the activity's completion to the selector, callback can be nil
func (f *FooFuture) Select(sel workflow.Selector, fn func(*FooFuture)) workflow.Selector {
	return sel.AddFuture(f.Future, func(workflow.Future) {
		if fn != nil {
			fn(f)
		}
	})
}

// Foo executes a(n) test.patch.FooService.Foo activity
func Foo(ctx workflow.Context, req *FooInput, options ...*FooActivityOptions) (*FooOutput, error) {
	return FooAsync(ctx, req, options...).Get(ctx)
}

// FooAsync executes a(n) test.patch.FooService.Foo activity (asynchronously)
func FooAsync(ctx workflow.Context, req *FooInput, options ...*FooActivityOptions) *FooFuture {
	var o *FooActivityOptions
	if len(options) > 0 && options[0] != nil {
		o = options[0]
	} else {
		o = NewFooActivityOptions()
	}
	var err error
	if ctx, err = o.Build(ctx); err != nil {
		errF, errS := workflow.NewFuture(ctx)
		errS.SetError(err)
		return &FooFuture{Future: errF}
	}
	activity := FooActivityName
	future := &FooFuture{Future: workflow.ExecuteActivity(ctx, activity, req)}
	return future
}

// FooLocal executes a(n) test.patch.FooService.Foo activity (locally)
func FooLocal(ctx workflow.Context, req *FooInput, options ...*FooLocalActivityOptions) (*FooOutput, error) {
	return FooLocalAsync(ctx, req, options...).Get(ctx)
}

// FooLocalAsync executes a(n) test.patch.FooService.Foo activity (asynchronously, locally)
func FooLocalAsync(ctx workflow.Context, req *FooInput, options ...*FooLocalActivityOptions) *FooFuture {
	var o *FooLocalActivityOptions
	if len(options) > 0 && options[0] != nil {
		o = options[0]
	} else {
		o = NewFooLocalActivityOptions()
	}
	var err error
	if ctx, err = o.Build(ctx); err != nil {
		errF, errS := workflow.NewFuture(ctx)
		errS.SetError(err)
		return &FooFuture{Future: errF}
	}
	var activity any
	if o.fn != nil {
		activity = o.fn
	} else {
		activity = FooActivityName
	}
	future := &FooFuture{Future: workflow.ExecuteLocalActivity(ctx, activity, req)}
	return future
}

// FooActivityOptions provides configuration for a(n) test.patch.FooService.Foo activity
type FooActivityOptions struct {
	options                workflow.ActivityOptions
	retryPolicy            *temporal.RetryPolicy
	scheduleToCloseTimeout *time.Duration
	startToCloseTimeout    *time.Duration
	heartbeatTimeout       *time.Duration
	scheduleToStartTimeout *time.Duration
	taskQueue              *string
	waitForCancellation    *bool
}

// NewFooActivityOptions initializes a new FooActivityOptions value
func NewFooActivityOptions() *FooActivityOptions {
	return &FooActivityOptions{}
}

// Build initializes a workflow.Context with appropriate ActivityOptions values derived from schema defaults and any user-defined overrides
func (o *FooActivityOptions) Build(ctx workflow.Context) (workflow.Context, error) {
	opts := o.options
	if v := o.heartbeatTimeout; v != nil {
		opts.HeartbeatTimeout = *v
	}
	if v := o.retryPolicy; v != nil {
		opts.RetryPolicy = v
	}
	if v := o.scheduleToCloseTimeout; v != nil {
		opts.ScheduleToCloseTimeout = *v
	}
	if v := o.scheduleToStartTimeout; v != nil {
		opts.ScheduleToStartTimeout = *v
	}
	if v := o.startToCloseTimeout; v != nil {
		opts.StartToCloseTimeout = *v
	} else if opts.StartToCloseTimeout == 0 {
		opts.StartToCloseTimeout = 2000000000 // 2 seconds
	}
	if v := o.taskQueue; v != nil {
		opts.TaskQueue = *v
	} else if opts.TaskQueue == "" {
		opts.TaskQueue = FooServiceTaskQueue
	}
	if v := o.waitForCancellation; v != nil {
		opts.WaitForCancellation = *v
	}
	return workflow.WithActivityOptions(ctx, opts), nil
}

// WithActivityOptions specifies an initial ActivityOptions value to which defaults will be applied
func (o *FooActivityOptions) WithActivityOptions(options workflow.ActivityOptions) *FooActivityOptions {
	o.options = options
	return o
}

// WithHeartbeatTimeout sets the HeartbeatTimeout value
func (o *FooActivityOptions) WithHeartbeatTimeout(d time.Duration) *FooActivityOptions {
	o.heartbeatTimeout = &d
	return o
}

// WithRetryPolicy sets the RetryPolicy value
func (o *FooActivityOptions) WithRetryPolicy(policy *temporal.RetryPolicy) *FooActivityOptions {
	o.retryPolicy = policy
	return o
}

// WithScheduleToCloseTimeout sets the ScheduleToCloseTimeout value
func (o *FooActivityOptions) WithScheduleToCloseTimeout(d time.Duration) *FooActivityOptions {
	o.scheduleToCloseTimeout = &d
	return o
}

// WithScheduleToStartTimeout sets the ScheduleToStartTimeout value
func (o *FooActivityOptions) WithScheduleToStartTimeout(d time.Duration) *FooActivityOptions {
	o.scheduleToStartTimeout = &d
	return o
}

// WithStartToCloseTimeout sets the StartToCloseTimeout value
func (o *FooActivityOptions) WithStartToCloseTimeout(d time.Duration) *FooActivityOptions {
	o.startToCloseTimeout = &d
	return o
}

// WithTaskQueue sets the TaskQueue value
func (o *FooActivityOptions) WithTaskQueue(tq string) *FooActivityOptions {
	o.taskQueue = &tq
	return o
}

// WithWaitForCancellation sets the WaitForCancellation value
func (o *FooActivityOptions) WithWaitForCancellation(wait bool) *FooActivityOptions {
	o.waitForCancellation = &wait
	return o
}

// FooLocalActivityOptions provides configuration for a(n) test.patch.FooService.Foo activity
type FooLocalActivityOptions struct {
	options                workflow.LocalActivityOptions
	retryPolicy            *temporal.RetryPolicy
	scheduleToCloseTimeout *time.Duration
	startToCloseTimeout    *time.Duration
	fn                     func(context.Context, *FooInput) (*FooOutput, error)
}

// NewFooLocalActivityOptions initializes a new FooLocalActivityOptions value
func NewFooLocalActivityOptions() *FooLocalActivityOptions {
	return &FooLocalActivityOptions{}
}

// Build initializes a workflow.Context with appropriate LocalActivityOptions values derived from schema defaults and any user-defined overrides
func (o *FooLocalActivityOptions) Build(ctx workflow.Context) (workflow.Context, error) {
	opts := o.options
	if v := o.retryPolicy; v != nil {
		opts.RetryPolicy = v
	}
	if v := o.scheduleToCloseTimeout; v != nil {
		opts.ScheduleToCloseTimeout = *v
	}
	if v := o.startToCloseTimeout; v != nil {
		opts.StartToCloseTimeout = *v
	} else if opts.StartToCloseTimeout == 0 {
		opts.StartToCloseTimeout = 2000000000 // 2 seconds
	}
	return workflow.WithLocalActivityOptions(ctx, opts), nil
}

// Local specifies a custom test.patch.FooService.Foo implementation
func (o *FooLocalActivityOptions) Local(fn func(context.Context, *FooInput) (*FooOutput, error)) *FooLocalActivityOptions {
	o.fn = fn
	return o
}

// WithLocalActivityOptions specifies an initial LocalActivityOptions value to which defaults will be applied
func (o *FooLocalActivityOptions) WithLocalActivityOptions(options workflow.LocalActivityOptions) *FooLocalActivityOptions {
	o.options = options
	return o
}

// WithRetryPolicy sets the RetryPolicy value
func (o *FooLocalActivityOptions) WithRetryPolicy(policy *temporal.RetryPolicy) *FooLocalActivityOptions {
	o.retryPolicy = policy
	return o
}

// WithScheduleToCloseTimeout sets the ScheduleToCloseTimeout value
func (o *FooLocalActivityOptions) WithScheduleToCloseTimeout(d time.Duration) *FooLocalActivityOptions {
	o.scheduleToCloseTimeout = &d
	return o
}

// WithStartToCloseTimeout sets the StartToCloseTimeout value
func (o *FooLocalActivityOptions) WithStartToCloseTimeout(d time.Duration) *FooLocalActivityOptions {
	o.startToCloseTimeout = &d
	return o
}

// TestClient provides a testsuite-compatible Client
type TestFooServiceClient struct {
	env       *testsuite.TestWorkflowEnvironment
	workflows FooServiceWorkflows
}

var _ FooServiceClient = &TestFooServiceClient{}

// NewTestFooServiceClient initializes a new TestFooServiceClient value
func NewTestFooServiceClient(env *testsuite.TestWorkflowEnvironment, workflows FooServiceWorkflows, activities FooServiceActivities) *TestFooServiceClient {
	if workflows != nil {
		RegisterFooServiceWorkflows(env, workflows)
	}
	if activities != nil {
		RegisterFooServiceActivities(env, activities)
	}
	return &TestFooServiceClient{env, workflows}
}

// Foo executes a(n) test.patch.FooService.Foo workflow in the test environment
func (c *TestFooServiceClient) Foo(ctx context.Context, req *FooInput, opts ...*FooOptions) (*FooOutput, error) {
	run, err := c.FooAsync(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// FooAsync executes a(n) test.patch.FooService.Foo workflow in the test environment
func (c *TestFooServiceClient) FooAsync(ctx context.Context, req *FooInput, options ...*FooOptions) (FooRun, error) {
	var o *FooOptions
	if len(options) > 0 && options[0] != nil {
		o = options[0]
	} else {
		o = NewFooOptions()
	}
	opts, err := o.Build(req.ProtoReflect())
	if err != nil {
		return nil, fmt.Errorf("error initializing client.StartWorkflowOptions: %w", err)
	}
	return &testFooRun{client: c, env: c.env, opts: &opts, req: req, workflows: c.workflows}, nil
}

// GetFoo is a noop
func (c *TestFooServiceClient) GetFoo(ctx context.Context, workflowID string, runID string) FooRun {
	return &testFooRun{env: c.env, workflows: c.workflows}
}

// CancelWorkflow requests cancellation of an existing workflow execution
func (c *TestFooServiceClient) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	c.env.CancelWorkflow()
	return nil
}

// TerminateWorkflow terminates an existing workflow execution
func (c *TestFooServiceClient) TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error {
	return c.CancelWorkflow(ctx, workflowID, runID)
}

var _ FooRun = &testFooRun{}

// testFooRun provides convenience methods for interacting with a(n) test.patch.FooService.Foo workflow in the test environment
type testFooRun struct {
	client    *TestFooServiceClient
	env       *testsuite.TestWorkflowEnvironment
	opts      *client.StartWorkflowOptions
	req       *FooInput
	workflows FooServiceWorkflows
}

// Cancel requests cancellation of a workflow in execution, returning an error if applicable
func (r *testFooRun) Cancel(ctx context.Context) error {
	return r.client.CancelWorkflow(ctx, r.ID(), r.RunID())
}

// Get retrieves a test test.patch.FooService.Foo workflow result
func (r *testFooRun) Get(context.Context) (*FooOutput, error) {
	r.env.ExecuteWorkflow(FooWorkflowName, r.req)
	if !r.env.IsWorkflowCompleted() {
		return nil, errors.New("workflow in progress")
	}
	if err := r.env.GetWorkflowError(); err != nil {
		return nil, err
	}
	var result FooOutput
	if err := r.env.GetWorkflowResult(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ID returns a test test.patch.FooService.Foo workflow run's workflow ID
func (r *testFooRun) ID() string {
	if r.opts != nil {
		return r.opts.ID
	}
	return ""
}

// Run noop implementation
func (r *testFooRun) Run() client.WorkflowRun {
	return nil
}

// RunID noop implementation
func (r *testFooRun) RunID() string {
	return ""
}

// Terminate terminates a workflow in execution, returning an error if applicable
func (r *testFooRun) Terminate(ctx context.Context, reason string, details ...interface{}) error {
	return r.client.TerminateWorkflow(ctx, r.ID(), r.RunID(), reason, details...)
}

// FooServiceCliOptions describes runtime configuration for test.patch.FooService cli
type FooServiceCliOptions struct {
	after            func(*v2.Context) error
	before           func(*v2.Context) error
	clientForCommand func(*v2.Context) (client.Client, error)
	worker           func(*v2.Context, client.Client) (worker.Worker, error)
}

// NewFooServiceCliOptions initializes a new FooServiceCliOptions value
func NewFooServiceCliOptions() *FooServiceCliOptions {
	return &FooServiceCliOptions{}
}

// WithAfter injects a custom After hook to be run after any command invocation
func (opts *FooServiceCliOptions) WithAfter(fn func(*v2.Context) error) *FooServiceCliOptions {
	opts.after = fn
	return opts
}

// WithBefore injects a custom Before hook to be run prior to any command invocation
func (opts *FooServiceCliOptions) WithBefore(fn func(*v2.Context) error) *FooServiceCliOptions {
	opts.before = fn
	return opts
}

// WithClient provides a Temporal client factory for use by commands
func (opts *FooServiceCliOptions) WithClient(fn func(*v2.Context) (client.Client, error)) *FooServiceCliOptions {
	opts.clientForCommand = fn
	return opts
}

// WithWorker provides an method for initializing a worker
func (opts *FooServiceCliOptions) WithWorker(fn func(*v2.Context, client.Client) (worker.Worker, error)) *FooServiceCliOptions {
	opts.worker = fn
	return opts
}

// NewFooServiceCli initializes a cli for a(n) test.patch.FooService service
func NewFooServiceCli(options ...*FooServiceCliOptions) (*v2.App, error) {
	commands, err := newFooServiceCommands(options...)
	if err != nil {
		return nil, fmt.Errorf("error initializing subcommands: %w", err)
	}
	return &v2.App{
		Name:     "foo-service",
		Commands: commands,
	}, nil
}

// NewFooServiceCliCommand initializes a cli command for a test.patch.FooService service with subcommands for each query, signal, update, and workflow
func NewFooServiceCliCommand(options ...*FooServiceCliOptions) (*v2.Command, error) {
	subcommands, err := newFooServiceCommands(options...)
	if err != nil {
		return nil, fmt.Errorf("error initializing subcommands: %w", err)
	}
	return &v2.Command{
		Name:        "foo-service",
		Subcommands: subcommands,
	}, nil
}

// newFooServiceCommands initializes (sub)commands for a test.patch.FooService cli or command
func newFooServiceCommands(options ...*FooServiceCliOptions) ([]*v2.Command, error) {
	opts := &FooServiceCliOptions{}
	if len(options) > 0 {
		opts = options[0]
	}
	if opts.clientForCommand == nil {
		opts.clientForCommand = func(*v2.Context) (client.Client, error) {
			return client.Dial(client.Options{})
		}
	}
	commands := []*v2.Command{
		{
			Name:                   "foo",
			Usage:                  "executes a(n) test.patch.FooService.Foo workflow",
			Category:               "WORKFLOWS",
			UseShortOptionHandling: true,
			Before:                 opts.before,
			After:                  opts.after,
			Flags: []v2.Flag{
				&v2.BoolFlag{
					Name:    "detach",
					Usage:   "run workflow in the background and print workflow and execution id",
					Aliases: []string{"d"},
				},
				&v2.StringFlag{
					Name:    "task-queue",
					Usage:   "task queue name",
					Aliases: []string{"t"},
					EnvVars: []string{"TEMPORAL_TASK_QUEUE_NAME", "TEMPORAL_TASK_QUEUE", "TASK_QUEUE_NAME", "TASK_QUEUE"},
					Value:   "foo-queue",
				},
				&v2.StringFlag{
					Name:    "input-file",
					Usage:   "path to json-formatted input file",
					Aliases: []string{"f"},
				},
				&v2.StringFlag{
					Name:     "foo-id",
					Usage:    "set the value of the operation's \"FooID\" parameter",
					Category: "INPUT",
				},
			},
			Action: func(cmd *v2.Context) error {
				tc, err := opts.clientForCommand(cmd)
				if err != nil {
					return fmt.Errorf("error initializing client for command: %w", err)
				}
				defer tc.Close()
				c := NewFooServiceClient(tc)
				req, err := UnmarshalCliFlagsToFooInput(cmd)
				if err != nil {
					return fmt.Errorf("error unmarshalling request: %w", err)
				}
				opts := client.StartWorkflowOptions{}
				if tq := cmd.String("task-queue"); tq != "" {
					opts.TaskQueue = tq
				}
				run, err := c.FooAsync(cmd.Context, req, NewFooOptions().WithStartWorkflowOptions(opts))
				if err != nil {
					return fmt.Errorf("error starting %s workflow: %w", FooWorkflowName, err)
				}
				if cmd.Bool("detach") {
					fmt.Println("success")
					fmt.Printf("workflow id: %s\n", run.ID())
					fmt.Printf("run id: %s\n", run.RunID())
					return nil
				}
				if resp, err := run.Get(cmd.Context); err != nil {
					return err
				} else {
					b, err := protojson.Marshal(resp)
					if err != nil {
						return fmt.Errorf("error serializing response json: %w", err)
					}
					var out bytes.Buffer
					if err := json.Indent(&out, b, "", "  "); err != nil {
						return fmt.Errorf("error formatting json: %w", err)
					}
					fmt.Println(out.String())
					return nil
				}
			},
		},
	}
	if opts.worker != nil {
		commands = append(commands, []*v2.Command{
			{
				Name:                   "worker",
				Usage:                  "runs a test.patch.FooService worker process",
				UseShortOptionHandling: true,
				Before:                 opts.before,
				After:                  opts.after,
				Action: func(cmd *v2.Context) error {
					c, err := opts.clientForCommand(cmd)
					if err != nil {
						return fmt.Errorf("error initializing client for command: %w", err)
					}
					defer c.Close()
					w, err := opts.worker(cmd, c)
					if opts.worker != nil {
						if err != nil {
							return fmt.Errorf("error initializing worker: %w", err)
						}
					}
					if err := w.Start(); err != nil {
						return fmt.Errorf("error starting worker: %w", err)
					}
					defer w.Stop()
					<-cmd.Context.Done()
					return nil
				},
			},
		}...)
	}
	sort.Slice(commands, func(i, j int) bool {
		return commands[i].Name < commands[j].Name
	})
	return commands, nil
}

// UnmarshalCliFlagsToFooInput unmarshals a FooInput from command line flags
func UnmarshalCliFlagsToFooInput(cmd *v2.Context) (*FooInput, error) {
	var result FooInput
	var hasValues bool
	if cmd.IsSet("input-file") {
		inputFile, err := gohomedir.Expand(cmd.String("input-file"))
		if err != nil {
			inputFile = cmd.String("input-file")
		}
		b, err := os.ReadFile(inputFile)
		if err != nil {
			return nil, fmt.Errorf("error reading input-file: %w", err)
		}
		if err := protojson.Unmarshal(b, &result); err != nil {
			return nil, fmt.Errorf("error parsing input-file json: %w", err)
		}
		hasValues = true
	}
	if cmd.IsSet("foo-id") {
		hasValues = true
		result.FooID = cmd.String("foo-id")
	}
	if !hasValues {
		return nil, nil
	}
	return &result, nil
}

// WithFooServiceSchemeTypes registers all FooService protobuf types with the given scheme
func WithFooServiceSchemeTypes() scheme.Option {
	return func(s *scheme.Scheme) {
		s.RegisterType(File_test_patch_example_proto.Messages().ByName("FooInput"))
		s.RegisterType(File_test_patch_example_proto.Messages().ByName("FooOutput"))
	}
}
