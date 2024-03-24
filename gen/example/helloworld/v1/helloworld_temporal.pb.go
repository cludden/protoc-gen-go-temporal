// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 1.10.2-next (cf7bb8abe8c16d8977887195eae0e5ba5b97300c)
//	go go1.21.5
//	protoc (unknown)
//
// source: example/helloworld/v1/helloworld.proto
package helloworldv1

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	expression "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	helpers "github.com/cludden/protoc-gen-go-temporal/pkg/helpers"
	scheme "github.com/cludden/protoc-gen-go-temporal/pkg/scheme"
	gohomedir "github.com/mitchellh/go-homedir"
	v2 "github.com/urfave/cli/v2"
	activity "go.temporal.io/sdk/activity"
	client "go.temporal.io/sdk/client"
	testsuite "go.temporal.io/sdk/testsuite"
	worker "go.temporal.io/sdk/worker"
	workflow "go.temporal.io/sdk/workflow"
	protojson "google.golang.org/protobuf/encoding/protojson"
	"log/slog"
	"os"
	"sort"
)

// ExampleTaskQueue is the default task-queue for a example.helloworld.v1.Example worker
const ExampleTaskQueue = "hello-world"

// example.helloworld.v1.Example workflow names
const (
	HelloWorldWorkflowName = "example.helloworld.v1.Example.HelloWorld"
)

// example.helloworld.v1.Example workflow id expressions
var (
	HelloWorldIdexpression = expression.MustParseExpression("hello_world/${! uuid_v4() }")
)

// example.helloworld.v1.Example activity names
const (
	HelloWorldActivityName = "example.helloworld.v1.Example.HelloWorld"
)

// ExampleClient describes a client for a(n) example.helloworld.v1.Example worker
type ExampleClient interface {
	// HelloWorld executes a(n) example.helloworld.v1.Example.HelloWorld workflow and blocks until error or response received
	HelloWorld(ctx context.Context, req *HelloWorldInput, opts ...*HelloWorldOptions) (*HelloWorldOutput, error)

	// HelloWorldAsync starts a(n) example.helloworld.v1.Example.HelloWorld workflow and returns a handle to the workflow run
	HelloWorldAsync(ctx context.Context, req *HelloWorldInput, opts ...*HelloWorldOptions) (HelloWorldRun, error)

	// GetHelloWorld retrieves a handle to an existing example.helloworld.v1.Example.HelloWorld workflow execution
	GetHelloWorld(ctx context.Context, workflowID string, runID string) HelloWorldRun

	// CancelWorkflow requests cancellation of an existing workflow execution
	CancelWorkflow(ctx context.Context, workflowID string, runID string) error

	// TerminateWorkflow an existing workflow execution
	TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error
}

// exampleClient implements a temporal client for a example.helloworld.v1.Example service
type exampleClient struct {
	client client.Client
	log    *slog.Logger
}

// NewExampleClient initializes a new example.helloworld.v1.Example client
func NewExampleClient(c client.Client, options ...*exampleClientOptions) ExampleClient {
	var cfg *exampleClientOptions
	if len(options) > 0 {
		cfg = options[0]
	} else {
		cfg = NewExampleClientOptions()
	}
	return &exampleClient{
		client: c,
		log:    cfg.getLogger(),
	}
}

// NewExampleClientWithOptions initializes a new Example client with the given options
func NewExampleClientWithOptions(c client.Client, opts client.Options, options ...*exampleClientOptions) (ExampleClient, error) {
	var err error
	c, err = client.NewClientFromExisting(c, opts)
	if err != nil {
		return nil, fmt.Errorf("error initializing client with options: %w", err)
	}
	var cfg *exampleClientOptions
	if len(options) > 0 {
		cfg = options[0]
	} else {
		cfg = NewExampleClientOptions()
	}
	return &exampleClient{
		client: c,
		log:    cfg.getLogger(),
	}, nil
}

// exampleClientOptions describes optional runtime configuration for a ExampleClient
type exampleClientOptions struct {
	log *slog.Logger
}

// NewExampleClientOptions initializes a new exampleClientOptions value
func NewExampleClientOptions() *exampleClientOptions {
	return &exampleClientOptions{}
}

// WithLogger can be used to override the default logger
func (opts *exampleClientOptions) WithLogger(l *slog.Logger) *exampleClientOptions {
	if l != nil {
		opts.log = l
	}
	return opts
}

// getLogger returns the configured logger, or the default logger
func (opts *exampleClientOptions) getLogger() *slog.Logger {
	if opts != nil && opts.log != nil {
		return opts.log
	}
	return slog.Default()
}

// example.helloworld.v1.Example.HelloWorld executes a example.helloworld.v1.Example.HelloWorld workflow and blocks until error or response received
func (c *exampleClient) HelloWorld(ctx context.Context, req *HelloWorldInput, options ...*HelloWorldOptions) (*HelloWorldOutput, error) {
	run, err := c.HelloWorldAsync(ctx, req, options...)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// HelloWorldAsync starts a(n) example.helloworld.v1.Example.HelloWorld workflow and returns a handle to the workflow run
func (c *exampleClient) HelloWorldAsync(ctx context.Context, req *HelloWorldInput, options ...*HelloWorldOptions) (HelloWorldRun, error) {
	opts := &client.StartWorkflowOptions{}
	if len(options) > 0 && options[0].opts != nil {
		opts = options[0].opts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = ExampleTaskQueue
	}
	if opts.ID == "" {
		id, err := expression.EvalExpression(HelloWorldIdexpression, req.ProtoReflect())
		if err != nil {
			return nil, fmt.Errorf("error evaluating id expression for \"HelloWorld\" workflow: %w", err)
		}
		opts.ID = id
	}
	run, err := c.client.ExecuteWorkflow(ctx, *opts, HelloWorldWorkflowName, req)
	if err != nil {
		return nil, err
	}
	if run == nil {
		return nil, errors.New("execute workflow returned nil run")
	}
	return &helloWorldRun{
		client: c,
		run:    run,
	}, nil
}

// GetHelloWorld fetches an existing example.helloworld.v1.Example.HelloWorld execution
func (c *exampleClient) GetHelloWorld(ctx context.Context, workflowID string, runID string) HelloWorldRun {
	return &helloWorldRun{
		client: c,
		run:    c.client.GetWorkflow(ctx, workflowID, runID),
	}
}

// CancelWorkflow requests cancellation of an existing workflow execution
func (c *exampleClient) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	return c.client.CancelWorkflow(ctx, workflowID, runID)
}

// TerminateWorkflow terminates an existing workflow execution
func (c *exampleClient) TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error {
	return c.client.TerminateWorkflow(ctx, workflowID, runID, reason, details...)
}

// HelloWorldOptions provides configuration for a example.helloworld.v1.Example.HelloWorld workflow operation
type HelloWorldOptions struct {
	opts *client.StartWorkflowOptions
}

// NewHelloWorldOptions initializes a new HelloWorldOptions value
func NewHelloWorldOptions() *HelloWorldOptions {
	return &HelloWorldOptions{}
}

// WithStartWorkflowOptions sets the initial client.StartWorkflowOptions
func (opts *HelloWorldOptions) WithStartWorkflowOptions(options client.StartWorkflowOptions) *HelloWorldOptions {
	opts.opts = &options
	return opts
}

// HelloWorldRun describes a(n) example.helloworld.v1.Example.HelloWorld workflow run
type HelloWorldRun interface {
	// ID returns the workflow ID
	ID() string

	// RunID returns the workflow instance ID
	RunID() string

	// Run returns the inner client.WorkflowRun
	Run() client.WorkflowRun

	// Get blocks until the workflow is complete and returns the result
	Get(ctx context.Context) (*HelloWorldOutput, error)

	// Cancel requests cancellation of a workflow in execution, returning an error if applicable
	Cancel(ctx context.Context) error

	// Terminate terminates a workflow in execution, returning an error if applicable
	Terminate(ctx context.Context, reason string, details ...interface{}) error
}

// helloWorldRun provides an internal implementation of a(n) HelloWorldRunRun
type helloWorldRun struct {
	client *exampleClient
	run    client.WorkflowRun
}

// ID returns the workflow ID
func (r *helloWorldRun) ID() string {
	return r.run.GetID()
}

// Run returns the inner client.WorkflowRun
func (r *helloWorldRun) Run() client.WorkflowRun {
	return r.run
}

// RunID returns the execution ID
func (r *helloWorldRun) RunID() string {
	return r.run.GetRunID()
}

// Cancel requests cancellation of a workflow in execution, returning an error if applicable
func (r *helloWorldRun) Cancel(ctx context.Context) error {
	return r.client.CancelWorkflow(ctx, r.ID(), r.RunID())
}

// Get blocks until the workflow is complete, returning the result if applicable
func (r *helloWorldRun) Get(ctx context.Context) (*HelloWorldOutput, error) {
	var resp HelloWorldOutput
	if err := r.run.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Terminate terminates a workflow in execution, returning an error if applicable
func (r *helloWorldRun) Terminate(ctx context.Context, reason string, details ...interface{}) error {
	return r.client.TerminateWorkflow(ctx, r.ID(), r.RunID(), reason, details...)
}

// Reference to generated workflow functions
var (
	// HelloWorldFunction implements a "HelloWorldWorkflow" workflow
	HelloWorldFunction func(workflow.Context, *HelloWorldInput) (*HelloWorldOutput, error)
)

// ExampleWorkflows provides methods for initializing new example.helloworld.v1.Example workflow values
type ExampleWorkflows interface {
	// example.helloworld.v1.Example.HelloWorld initializes a new a(n) HelloWorldWorkflow implementation
	HelloWorld(ctx workflow.Context, input *HelloWorldWorkflowInput) (HelloWorldWorkflow, error)
}

// RegisterExampleWorkflows registers example.helloworld.v1.Example workflows with the given worker
func RegisterExampleWorkflows(r worker.WorkflowRegistry, workflows ExampleWorkflows) {
	RegisterHelloWorldWorkflow(r, workflows.HelloWorld)
}

// RegisterHelloWorldWorkflow registers a example.helloworld.v1.Example.HelloWorld workflow with the given worker
func RegisterHelloWorldWorkflow(r worker.WorkflowRegistry, wf func(workflow.Context, *HelloWorldWorkflowInput) (HelloWorldWorkflow, error)) {
	HelloWorldFunction = buildHelloWorld(wf)
	r.RegisterWorkflowWithOptions(HelloWorldFunction, workflow.RegisterOptions{Name: HelloWorldWorkflowName})
}

// buildHelloWorld converts a HelloWorld workflow struct into a valid workflow function
func buildHelloWorld(ctor func(workflow.Context, *HelloWorldWorkflowInput) (HelloWorldWorkflow, error)) func(workflow.Context, *HelloWorldInput) (*HelloWorldOutput, error) {
	return func(ctx workflow.Context, req *HelloWorldInput) (*HelloWorldOutput, error) {
		input := &HelloWorldWorkflowInput{
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

// HelloWorldWorkflowInput describes the input to a(n) example.helloworld.v1.Example.HelloWorld workflow constructor
type HelloWorldWorkflowInput struct {
	Req *HelloWorldInput
}

// HelloWorldWorkflow describes a(n) example.helloworld.v1.Example.HelloWorld workflow implementation
//
// workflow details: (id: "hello_world/${! uuid_v4() }")
type HelloWorldWorkflow interface {
	// Execute defines the entrypoint to a(n) example.helloworld.v1.Example.HelloWorld workflow
	Execute(ctx workflow.Context) (*HelloWorldOutput, error)
}

// HelloWorldChild executes a child example.helloworld.v1.Example.HelloWorld workflow and blocks until error or response received
func HelloWorldChild(ctx workflow.Context, req *HelloWorldInput, options ...*HelloWorldChildOptions) (*HelloWorldOutput, error) {
	childRun, err := HelloWorldChildAsync(ctx, req, options...)
	if err != nil {
		return nil, err
	}
	return childRun.Get(ctx)
}

// HelloWorldChildAsync starts a child example.helloworld.v1.Example.HelloWorld workflow and returns a handle to the child workflow run
func HelloWorldChildAsync(ctx workflow.Context, req *HelloWorldInput, options ...*HelloWorldChildOptions) (*HelloWorldChildRun, error) {
	var opts *workflow.ChildWorkflowOptions
	if len(options) > 0 && options[0].opts != nil {
		opts = options[0].opts
	} else {
		childOpts := workflow.GetChildWorkflowOptions(ctx)
		opts = &childOpts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = ExampleTaskQueue
	}
	if opts.WorkflowID == "" {
		id, err := expression.EvalExpression(HelloWorldIdexpression, req.ProtoReflect())
		if err != nil {
			panic(err)
		}
		opts.WorkflowID = id
	}
	ctx = workflow.WithChildOptions(ctx, *opts)
	return &HelloWorldChildRun{Future: workflow.ExecuteChildWorkflow(ctx, HelloWorldWorkflowName, req)}, nil
}

// HelloWorldChildOptions provides configuration for a example.helloworld.v1.Example.HelloWorld workflow operation
type HelloWorldChildOptions struct {
	opts *workflow.ChildWorkflowOptions
}

// NewHelloWorldChildOptions initializes a new HelloWorldChildOptions value
func NewHelloWorldChildOptions() *HelloWorldChildOptions {
	return &HelloWorldChildOptions{}
}

// WithChildWorkflowOptions sets the initial client.StartWorkflowOptions
func (opts *HelloWorldChildOptions) WithChildWorkflowOptions(options workflow.ChildWorkflowOptions) *HelloWorldChildOptions {
	opts.opts = &options
	return opts
}

// HelloWorldChildRun describes a child HelloWorld workflow run
type HelloWorldChildRun struct {
	Future workflow.ChildWorkflowFuture
}

// Get blocks until the workflow is completed, returning the response value
func (r *HelloWorldChildRun) Get(ctx workflow.Context) (*HelloWorldOutput, error) {
	var resp HelloWorldOutput
	if err := r.Future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Select adds this completion to the selector. Callback can be nil.
func (r *HelloWorldChildRun) Select(sel workflow.Selector, fn func(*HelloWorldChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future, func(workflow.Future) {
		if fn != nil {
			fn(r)
		}
	})
}

// SelectStart adds waiting for start to the selector. Callback can be nil.
func (r *HelloWorldChildRun) SelectStart(sel workflow.Selector, fn func(*HelloWorldChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future.GetChildWorkflowExecution(), func(workflow.Future) {
		if fn != nil {
			fn(r)
		}
	})
}

// WaitStart waits for the child workflow to start
func (r *HelloWorldChildRun) WaitStart(ctx workflow.Context) (*workflow.Execution, error) {
	var exec workflow.Execution
	if err := r.Future.GetChildWorkflowExecution().Get(ctx, &exec); err != nil {
		return nil, err
	}
	return &exec, nil
}

// ExampleActivities describes available worker activities
type ExampleActivities interface {
	// example.helloworld.v1.Example.HelloWorld implements a(n) example.helloworld.v1.Example.HelloWorld activity definition
	HelloWorld(ctx context.Context, req *HelloWorldInput) (*HelloWorldOutput, error)
}

// RegisterExampleActivities registers activities with a worker
func RegisterExampleActivities(r worker.ActivityRegistry, activities ExampleActivities) {
	RegisterHelloWorldActivity(r, activities.HelloWorld)
}

// RegisterHelloWorldActivity registers a example.helloworld.v1.Example.HelloWorld activity
func RegisterHelloWorldActivity(r worker.ActivityRegistry, fn func(context.Context, *HelloWorldInput) (*HelloWorldOutput, error)) {
	r.RegisterActivityWithOptions(fn, activity.RegisterOptions{
		Name: HelloWorldActivityName,
	})
}

// HelloWorldFuture describes a(n) example.helloworld.v1.Example.HelloWorld activity execution
type HelloWorldFuture struct {
	Future workflow.Future
}

// Get blocks on the activity's completion, returning the response
func (f *HelloWorldFuture) Get(ctx workflow.Context) (*HelloWorldOutput, error) {
	var resp HelloWorldOutput
	if err := f.Future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Select adds the activity's completion to the selector, callback can be nil
func (f *HelloWorldFuture) Select(sel workflow.Selector, fn func(*HelloWorldFuture)) workflow.Selector {
	return sel.AddFuture(f.Future, func(workflow.Future) {
		if fn != nil {
			fn(f)
		}
	})
}

// HelloWorld executes a(n) example.helloworld.v1.Example.HelloWorld activity
func HelloWorld(ctx workflow.Context, req *HelloWorldInput, options ...*HelloWorldActivityOptions) (*HelloWorldOutput, error) {
	var opts *HelloWorldActivityOptions
	if len(options) > 0 && options[0] != nil {
		opts = options[0]
	} else {
		opts = NewHelloWorldActivityOptions()
	}
	if opts.opts == nil {
		activityOpts := workflow.GetActivityOptions(ctx)
		opts.opts = &activityOpts
	}
	if opts.opts.StartToCloseTimeout == 0 {
		opts.opts.StartToCloseTimeout = 10000000000 // 10s
	}
	ctx = workflow.WithActivityOptions(ctx, *opts.opts)
	var activity any
	activity = HelloWorldActivityName
	future := &HelloWorldFuture{Future: workflow.ExecuteActivity(ctx, activity, req)}
	return future.Get(ctx)
}

// HelloWorldAsync executes a(n) example.helloworld.v1.Example.HelloWorld activity (asynchronously)
func HelloWorldAsync(ctx workflow.Context, req *HelloWorldInput, options ...*HelloWorldActivityOptions) *HelloWorldFuture {
	var opts *HelloWorldActivityOptions
	if len(options) > 0 && options[0] != nil {
		opts = options[0]
	} else {
		opts = NewHelloWorldActivityOptions()
	}
	if opts.opts == nil {
		activityOpts := workflow.GetActivityOptions(ctx)
		opts.opts = &activityOpts
	}
	if opts.opts.StartToCloseTimeout == 0 {
		opts.opts.StartToCloseTimeout = 10000000000 // 10s
	}
	ctx = workflow.WithActivityOptions(ctx, *opts.opts)
	var activity any
	activity = HelloWorldActivityName
	future := &HelloWorldFuture{Future: workflow.ExecuteActivity(ctx, activity, req)}
	return future
}

// HelloWorldLocal executes a(n) example.helloworld.v1.Example.HelloWorld activity (locally)
func HelloWorldLocal(ctx workflow.Context, req *HelloWorldInput, options ...*HelloWorldLocalActivityOptions) (*HelloWorldOutput, error) {
	var opts *HelloWorldLocalActivityOptions
	if len(options) > 0 && options[0] != nil {
		opts = options[0]
	} else {
		opts = NewHelloWorldLocalActivityOptions()
	}
	if opts.opts == nil {
		activityOpts := workflow.GetLocalActivityOptions(ctx)
		opts.opts = &activityOpts
	}
	if opts.opts.StartToCloseTimeout == 0 {
		opts.opts.StartToCloseTimeout = 10000000000 // 10s
	}
	ctx = workflow.WithLocalActivityOptions(ctx, *opts.opts)
	var activity any
	if opts.fn != nil {
		activity = opts.fn
	} else {
		activity = HelloWorldActivityName
	}
	future := &HelloWorldFuture{Future: workflow.ExecuteLocalActivity(ctx, activity, req)}
	return future.Get(ctx)
}

// HelloWorldLocalAsync executes a(n) example.helloworld.v1.Example.HelloWorld activity (asynchronously, locally)
func HelloWorldLocalAsync(ctx workflow.Context, req *HelloWorldInput, options ...*HelloWorldLocalActivityOptions) *HelloWorldFuture {
	var opts *HelloWorldLocalActivityOptions
	if len(options) > 0 && options[0] != nil {
		opts = options[0]
	} else {
		opts = NewHelloWorldLocalActivityOptions()
	}
	if opts.opts == nil {
		activityOpts := workflow.GetLocalActivityOptions(ctx)
		opts.opts = &activityOpts
	}
	if opts.opts.StartToCloseTimeout == 0 {
		opts.opts.StartToCloseTimeout = 10000000000 // 10s
	}
	ctx = workflow.WithLocalActivityOptions(ctx, *opts.opts)
	var activity any
	if opts.fn != nil {
		activity = opts.fn
	} else {
		activity = HelloWorldActivityName
	}
	future := &HelloWorldFuture{Future: workflow.ExecuteLocalActivity(ctx, activity, req)}
	return future
}

// HelloWorldLocalActivityOptions provides configuration for a local example.helloworld.v1.Example.HelloWorld activity
type HelloWorldLocalActivityOptions struct {
	fn   func(context.Context, *HelloWorldInput) (*HelloWorldOutput, error)
	opts *workflow.LocalActivityOptions
}

// NewHelloWorldLocalActivityOptions sets default LocalActivityOptions
func NewHelloWorldLocalActivityOptions() *HelloWorldLocalActivityOptions {
	return &HelloWorldLocalActivityOptions{}
}

// Local provides a local example.helloworld.v1.Example.HelloWorld activity implementation
func (opts *HelloWorldLocalActivityOptions) Local(fn func(context.Context, *HelloWorldInput) (*HelloWorldOutput, error)) *HelloWorldLocalActivityOptions {
	opts.fn = fn
	return opts
}

// WithLocalActivityOptions sets default LocalActivityOptions
func (opts *HelloWorldLocalActivityOptions) WithLocalActivityOptions(options workflow.LocalActivityOptions) *HelloWorldLocalActivityOptions {
	opts.opts = &options
	return opts
}

// HelloWorldActivityOptions provides configuration for a(n) example.helloworld.v1.Example.HelloWorld activity
type HelloWorldActivityOptions struct {
	opts *workflow.ActivityOptions
}

// NewHelloWorldActivityOptions sets default ActivityOptions
func NewHelloWorldActivityOptions() *HelloWorldActivityOptions {
	return &HelloWorldActivityOptions{}
}

// WithActivityOptions sets default ActivityOptions
func (opts *HelloWorldActivityOptions) WithActivityOptions(options workflow.ActivityOptions) *HelloWorldActivityOptions {
	opts.opts = &options
	return opts
}

// TestClient provides a testsuite-compatible Client
type TestExampleClient struct {
	env       *testsuite.TestWorkflowEnvironment
	workflows ExampleWorkflows
}

var _ ExampleClient = &TestExampleClient{}

// NewTestExampleClient initializes a new TestExampleClient value
func NewTestExampleClient(env *testsuite.TestWorkflowEnvironment, workflows ExampleWorkflows, activities ExampleActivities) *TestExampleClient {
	if workflows != nil {
		RegisterExampleWorkflows(env, workflows)
	}
	if activities != nil {
		RegisterExampleActivities(env, activities)
	}
	return &TestExampleClient{env, workflows}
}

// HelloWorld executes a(n) example.helloworld.v1.Example.HelloWorld workflow in the test environment
func (c *TestExampleClient) HelloWorld(ctx context.Context, req *HelloWorldInput, opts ...*HelloWorldOptions) (*HelloWorldOutput, error) {
	run, err := c.HelloWorldAsync(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// HelloWorldAsync executes a(n) example.helloworld.v1.Example.HelloWorld workflow in the test environment
func (c *TestExampleClient) HelloWorldAsync(ctx context.Context, req *HelloWorldInput, options ...*HelloWorldOptions) (HelloWorldRun, error) {
	opts := &client.StartWorkflowOptions{}
	if len(options) > 0 && options[0].opts != nil {
		opts = options[0].opts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = ExampleTaskQueue
	}
	if opts.ID == "" {
		id, err := expression.EvalExpression(HelloWorldIdexpression, req.ProtoReflect())
		if err != nil {
			return nil, fmt.Errorf("error evaluating id expression for \"HelloWorld\" workflow: %w", err)
		}
		opts.ID = id
	}
	return &testHelloWorldRun{client: c, env: c.env, opts: opts, req: req, workflows: c.workflows}, nil
}

// GetHelloWorld is a noop
func (c *TestExampleClient) GetHelloWorld(ctx context.Context, workflowID string, runID string) HelloWorldRun {
	return &testHelloWorldRun{env: c.env, workflows: c.workflows}
}

// CancelWorkflow requests cancellation of an existing workflow execution
func (c *TestExampleClient) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	c.env.CancelWorkflow()
	return nil
}

// TerminateWorkflow terminates an existing workflow execution
func (c *TestExampleClient) TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error {
	return c.CancelWorkflow(ctx, workflowID, runID)
}

var _ HelloWorldRun = &testHelloWorldRun{}

// testHelloWorldRun provides convenience methods for interacting with a(n) example.helloworld.v1.Example.HelloWorld workflow in the test environment
type testHelloWorldRun struct {
	client    *TestExampleClient
	env       *testsuite.TestWorkflowEnvironment
	opts      *client.StartWorkflowOptions
	req       *HelloWorldInput
	workflows ExampleWorkflows
}

// Cancel requests cancellation of a workflow in execution, returning an error if applicable
func (r *testHelloWorldRun) Cancel(ctx context.Context) error {
	return r.client.CancelWorkflow(ctx, r.ID(), r.RunID())
}

// Get retrieves a test example.helloworld.v1.Example.HelloWorld workflow result
func (r *testHelloWorldRun) Get(context.Context) (*HelloWorldOutput, error) {
	r.env.ExecuteWorkflow(HelloWorldWorkflowName, r.req)
	if !r.env.IsWorkflowCompleted() {
		return nil, errors.New("workflow in progress")
	}
	if err := r.env.GetWorkflowError(); err != nil {
		return nil, err
	}
	var result HelloWorldOutput
	if err := r.env.GetWorkflowResult(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ID returns a test example.helloworld.v1.Example.HelloWorld workflow run's workflow ID
func (r *testHelloWorldRun) ID() string {
	if r.opts != nil {
		return r.opts.ID
	}
	return ""
}

// Run noop implementation
func (r *testHelloWorldRun) Run() client.WorkflowRun {
	return nil
}

// RunID noop implementation
func (r *testHelloWorldRun) RunID() string {
	return ""
}

// Terminate terminates a workflow in execution, returning an error if applicable
func (r *testHelloWorldRun) Terminate(ctx context.Context, reason string, details ...interface{}) error {
	return r.client.TerminateWorkflow(ctx, r.ID(), r.RunID(), reason, details...)
}

// ExampleCliOptions describes runtime configuration for example.helloworld.v1.Example cli
type ExampleCliOptions struct {
	after            func(*v2.Context) error
	before           func(*v2.Context) error
	clientForCommand func(*v2.Context) (client.Client, error)
	worker           func(*v2.Context, client.Client) (worker.Worker, error)
}

// NewExampleCliOptions initializes a new ExampleCliOptions value
func NewExampleCliOptions() *ExampleCliOptions {
	return &ExampleCliOptions{}
}

// WithAfter injects a custom After hook to be run after any command invocation
func (opts *ExampleCliOptions) WithAfter(fn func(*v2.Context) error) *ExampleCliOptions {
	opts.after = fn
	return opts
}

// WithBefore injects a custom Before hook to be run prior to any command invocation
func (opts *ExampleCliOptions) WithBefore(fn func(*v2.Context) error) *ExampleCliOptions {
	opts.before = fn
	return opts
}

// WithClient provides a Temporal client factory for use by commands
func (opts *ExampleCliOptions) WithClient(fn func(*v2.Context) (client.Client, error)) *ExampleCliOptions {
	opts.clientForCommand = fn
	return opts
}

// WithWorker provides an method for initializing a worker
func (opts *ExampleCliOptions) WithWorker(fn func(*v2.Context, client.Client) (worker.Worker, error)) *ExampleCliOptions {
	opts.worker = fn
	return opts
}

// NewExampleCli initializes a cli for a(n) example.helloworld.v1.Example service
func NewExampleCli(options ...*ExampleCliOptions) (*v2.App, error) {
	commands, err := newExampleCommands(options...)
	if err != nil {
		return nil, fmt.Errorf("error initializing subcommands: %w", err)
	}
	return &v2.App{
		Name:     "example",
		Commands: commands,
	}, nil
}

// NewExampleCliCommand initializes a cli command for a example.helloworld.v1.Example service with subcommands for each query, signal, update, and workflow
func NewExampleCliCommand(options ...*ExampleCliOptions) (*v2.Command, error) {
	subcommands, err := newExampleCommands(options...)
	if err != nil {
		return nil, fmt.Errorf("error initializing subcommands: %w", err)
	}
	return &v2.Command{
		Name:        "example",
		Subcommands: subcommands,
	}, nil
}

// newExampleCommands initializes (sub)commands for a example.helloworld.v1.Example cli or command
func newExampleCommands(options ...*ExampleCliOptions) ([]*v2.Command, error) {
	opts := &ExampleCliOptions{}
	if len(options) > 0 {
		opts = options[0]
	}
	if opts.clientForCommand == nil {
		opts.clientForCommand = func(*v2.Context) (client.Client, error) {
			return client.Dial(client.Options{})
		}
	}
	commands := []*v2.Command{
		// example.helloworld.v1.Example.HelloWorld executes a(n) example.helloworld.v1.Example.HelloWorld workflow,
		{
			Name:                   "hello-world",
			Usage:                  "example.helloworld.v1.Example.HelloWorld executes a(n) example.helloworld.v1.Example.HelloWorld workflow",
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
					Value:   "hello-world",
				},
				&v2.StringFlag{
					Name:    "input-file",
					Usage:   "path to json-formatted input file",
					Aliases: []string{"f"},
				},
				&v2.StringFlag{
					Name:     "name",
					Usage:    "set the value of the operation's \"Name\" parameter",
					Category: "INPUT",
				},
			},
			Action: func(cmd *v2.Context) error {
				tc, err := opts.clientForCommand(cmd)
				if err != nil {
					return fmt.Errorf("error initializing client for command: %w", err)
				}
				defer tc.Close()
				c := NewExampleClient(tc)
				req, err := UnmarshalCliFlagsToHelloWorldInput(cmd)
				if err != nil {
					return fmt.Errorf("error unmarshalling request: %w", err)
				}
				opts := client.StartWorkflowOptions{}
				if tq := cmd.String("task-queue"); tq != "" {
					opts.TaskQueue = tq
				}
				run, err := c.HelloWorldAsync(cmd.Context, req, NewHelloWorldOptions().WithStartWorkflowOptions(opts))
				if err != nil {
					return fmt.Errorf("error starting %s workflow: %w", HelloWorldWorkflowName, err)
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
				Usage:                  "runs a example.helloworld.v1.Example worker process",
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

// UnmarshalCliFlagsToHelloWorldInput unmarshals a HelloWorldInput from command line flags
func UnmarshalCliFlagsToHelloWorldInput(cmd *v2.Context) (*HelloWorldInput, error) {
	var result HelloWorldInput
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
	if cmd.IsSet("name") {
		hasValues = true
		result.Name = cmd.String("name")
	}
	if !hasValues {
		return nil, nil
	}
	return &result, nil
}

// WithExampleSchemeTypes registers all Example protobuf types with the given scheme
func WithExampleSchemeTypes() scheme.Option {
	return func(s *scheme.Scheme) {
		s.RegisterType(File_example_helloworld_v1_helloworld_proto.Messages().ByName("HelloWorldInput"))
		s.RegisterType(File_example_helloworld_v1_helloworld_proto.Messages().ByName("HelloWorldOutput"))
	}
}
