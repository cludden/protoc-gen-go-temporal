// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 1.14.4-next (d2f9bfc0a16171313109a7b80bc10fabc41196e0)
//	go go1.22.6
//	protoc (unknown)
//
// source: example/helloworld/v1/example.proto
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
	enumsv1 "go.temporal.io/api/enums/v1"
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

// =============================================================================
// Constants (example.helloworld.v1.Example)
// =============================================================================

// ExampleTaskQueue is the default task-queue for a example.helloworld.v1.Example worker
const ExampleTaskQueue = "hello-world"

// example.helloworld.v1.Example workflow names
const (
	HelloWorkflowName = "example.v1.Hello"
)

// example.helloworld.v1.Example workflow id expressions
var (
	HelloIdexpression = expression.MustParseExpression("hello/${! name.or(\"World\") }")
)

// example.helloworld.v1.Example signal names
const (
	GoodbyeSignalName = "example.helloworld.v1.Example.Goodbye"
)

// =============================================================================
// Client (example.helloworld.v1.Example)
// =============================================================================

// ExampleClient describes a client for a(n) example.helloworld.v1.Example worker
type ExampleClient interface {
	// Hello prints a friendly greeting and waits for goodbye
	Hello(ctx context.Context, req *HelloRequest, opts ...*HelloOptions) (*HelloResponse, error)

	// HelloAsync starts a(n) example.v1.Hello workflow and returns a handle to the workflow run
	HelloAsync(ctx context.Context, req *HelloRequest, opts ...*HelloOptions) (HelloRun, error)

	// GetHello retrieves a handle to an existing example.v1.Hello workflow execution
	GetHello(ctx context.Context, workflowID string, runID string) HelloRun

	// CancelWorkflow requests cancellation of an existing workflow execution
	CancelWorkflow(ctx context.Context, workflowID string, runID string) error

	// TerminateWorkflow an existing workflow execution
	TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error

	// Goodbye signals a running workflow to exit
	Goodbye(ctx context.Context, workflowID string, runID string, signal *GoodbyeRequest) error
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

// Hello prints a friendly greeting and waits for goodbye
func (c *exampleClient) Hello(ctx context.Context, req *HelloRequest, options ...*HelloOptions) (*HelloResponse, error) {
	run, err := c.HelloAsync(ctx, req, options...)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// Hello prints a friendly greeting and waits for goodbye
func (c *exampleClient) HelloAsync(ctx context.Context, req *HelloRequest, options ...*HelloOptions) (HelloRun, error) {
	var o *HelloOptions
	if len(options) > 0 && options[0] != nil {
		o = options[0]
	} else {
		o = NewHelloOptions()
	}
	opts, err := o.Build(req.ProtoReflect())
	if err != nil {
		return nil, fmt.Errorf("error initializing client.StartWorkflowOptions: %w", err)
	}
	run, err := c.client.ExecuteWorkflow(ctx, opts, HelloWorkflowName, req)
	if err != nil {
		return nil, err
	}
	if run == nil {
		return nil, errors.New("execute workflow returned nil run")
	}
	return &helloRun{
		client: c,
		run:    run,
	}, nil
}

// GetHello fetches an existing example.v1.Hello execution
func (c *exampleClient) GetHello(ctx context.Context, workflowID string, runID string) HelloRun {
	return &helloRun{
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

// Goodbye signals a running workflow to exit
func (c *exampleClient) Goodbye(ctx context.Context, workflowID string, runID string, signal *GoodbyeRequest) error {
	return c.client.SignalWorkflow(ctx, workflowID, runID, GoodbyeSignalName, signal)
}

// HelloOptions provides configuration for a example.v1.Hello workflow operation
type HelloOptions struct {
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

// NewHelloOptions initializes a new HelloOptions value
func NewHelloOptions() *HelloOptions {
	return &HelloOptions{}
}

// Build initializes a new go.temporal.io/sdk/client.StartWorkflowOptions value with defaults and overrides applied
func (o *HelloOptions) Build(req protoreflect.Message) (client.StartWorkflowOptions, error) {
	opts := o.options
	if v := o.id; v != nil {
		opts.ID = *v
	} else if opts.ID == "" {
		id, err := expression.EvalExpression(HelloIdexpression, req)
		if err != nil {
			return opts, fmt.Errorf("error evaluating id expression for %q workflow: %w", HelloWorkflowName, err)
		}
		opts.ID = id
	}
	if v := o.idReusePolicy; v != enumsv1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v
	}
	if v := o.taskQueue; v != nil {
		opts.TaskQueue = *v
	} else if opts.TaskQueue == "" {
		opts.TaskQueue = ExampleTaskQueue
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
func (o *HelloOptions) WithStartWorkflowOptions(options client.StartWorkflowOptions) *HelloOptions {
	o.options = options
	return o
}

// WithExecutionTimeout sets the WorkflowExecutionTimeout value
func (o *HelloOptions) WithExecutionTimeout(d time.Duration) *HelloOptions {
	o.executionTimeout = &d
	return o
}

// WithID sets the ID value
func (o *HelloOptions) WithID(id string) *HelloOptions {
	o.id = &id
	return o
}

// WithIDReusePolicy sets the WorkflowIDReusePolicy value
func (o *HelloOptions) WithIDReusePolicy(policy enumsv1.WorkflowIdReusePolicy) *HelloOptions {
	o.idReusePolicy = policy
	return o
}

// WithRetryPolicy sets the RetryPolicy value
func (o *HelloOptions) WithRetryPolicy(policy *temporal.RetryPolicy) *HelloOptions {
	o.retryPolicy = policy
	return o
}

// WithRunTimeout sets the WorkflowRunTimeout value
func (o *HelloOptions) WithRunTimeout(d time.Duration) *HelloOptions {
	o.runTimeout = &d
	return o
}

// WithSearchAttributes sets the SearchAttributes value
func (o *HelloOptions) WithSearchAttributes(sa map[string]any) *HelloOptions {
	o.searchAttributes = sa
	return o
}

// WithTaskTimeout sets the WorkflowTaskTimeout value
func (o *HelloOptions) WithTaskTimeout(d time.Duration) *HelloOptions {
	o.taskTimeout = &d
	return o
}

// WithTaskQueue sets the TaskQueue value
func (o *HelloOptions) WithTaskQueue(tq string) *HelloOptions {
	o.taskQueue = &tq
	return o
}

// HelloRun describes a(n) example.v1.Hello workflow run
type HelloRun interface {
	// ID returns the workflow ID
	ID() string

	// RunID returns the workflow instance ID
	RunID() string

	// Run returns the inner client.WorkflowRun
	Run() client.WorkflowRun

	// Get blocks until the workflow is complete and returns the result
	Get(ctx context.Context) (*HelloResponse, error)

	// Cancel requests cancellation of a workflow in execution, returning an error if applicable
	Cancel(ctx context.Context) error

	// Terminate terminates a workflow in execution, returning an error if applicable
	Terminate(ctx context.Context, reason string, details ...interface{}) error

	// Goodbye signals a running workflow to exit
	Goodbye(ctx context.Context, req *GoodbyeRequest) error
}

// helloRun provides an internal implementation of a(n) HelloRunRun
type helloRun struct {
	client *exampleClient
	run    client.WorkflowRun
}

// ID returns the workflow ID
func (r *helloRun) ID() string {
	return r.run.GetID()
}

// Run returns the inner client.WorkflowRun
func (r *helloRun) Run() client.WorkflowRun {
	return r.run
}

// RunID returns the execution ID
func (r *helloRun) RunID() string {
	return r.run.GetRunID()
}

// Cancel requests cancellation of a workflow in execution, returning an error if applicable
func (r *helloRun) Cancel(ctx context.Context) error {
	return r.client.CancelWorkflow(ctx, r.ID(), r.RunID())
}

// Get blocks until the workflow is complete, returning the result if applicable
func (r *helloRun) Get(ctx context.Context) (*HelloResponse, error) {
	var resp HelloResponse
	if err := r.run.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Terminate terminates a workflow in execution, returning an error if applicable
func (r *helloRun) Terminate(ctx context.Context, reason string, details ...interface{}) error {
	return r.client.TerminateWorkflow(ctx, r.ID(), r.RunID(), reason, details...)
}

// Goodbye signals a running workflow to exit
func (r *helloRun) Goodbye(ctx context.Context, req *GoodbyeRequest) error {
	return r.client.Goodbye(ctx, r.ID(), "", req)
}

// Reference to generated workflow functions
var (
	// Hello prints a friendly greeting and waits for goodbye
	HelloFunction func(workflow.Context, *HelloRequest) (*HelloResponse, error)
)

// ExampleWorkflowFunctions describes a mockable dependency for inlining workflows within other workflows
type (
	// ExampleWorkflowFunctions describes a mockable dependency for inlining workflows within other workflows
	ExampleWorkflowFunctions interface {
		// Hello prints a friendly greeting and waits for goodbye
		Hello(workflow.Context, *HelloRequest) (*HelloResponse, error)
	}
	// exampleWorkflowFunctions provides an internal ExampleWorkflowFunctions implementation
	exampleWorkflowFunctions struct{}
)

func NewExampleWorkflowFunctions() ExampleWorkflowFunctions {
	return &exampleWorkflowFunctions{}
}

// Hello prints a friendly greeting and waits for goodbye
func (f *exampleWorkflowFunctions) Hello(ctx workflow.Context, req *HelloRequest) (*HelloResponse, error) {
	if HelloFunction == nil {
		return nil, errors.New("Hello requires workflow registration via RegisterExampleWorkflows or RegisterHelloWorkflow")
	}
	return HelloFunction(ctx, req)
}

// ExampleWorkflows provides methods for initializing new example.helloworld.v1.Example workflow values
type ExampleWorkflows interface {
	// Hello prints a friendly greeting and waits for goodbye
	Hello(ctx workflow.Context, input *HelloWorkflowInput) (HelloWorkflow, error)
}

// RegisterExampleWorkflows registers example.helloworld.v1.Example workflows with the given worker
func RegisterExampleWorkflows(r worker.WorkflowRegistry, workflows ExampleWorkflows) {
	RegisterHelloWorkflow(r, workflows.Hello)
}

// RegisterHelloWorkflow registers a example.helloworld.v1.Example.Hello workflow with the given worker
func RegisterHelloWorkflow(r worker.WorkflowRegistry, wf func(workflow.Context, *HelloWorkflowInput) (HelloWorkflow, error)) {
	HelloFunction = buildHello(wf)
	r.RegisterWorkflowWithOptions(HelloFunction, workflow.RegisterOptions{Name: HelloWorkflowName})
}

// buildHello converts a Hello workflow struct into a valid workflow function
func buildHello(ctor func(workflow.Context, *HelloWorkflowInput) (HelloWorkflow, error)) func(workflow.Context, *HelloRequest) (*HelloResponse, error) {
	return func(ctx workflow.Context, req *HelloRequest) (*HelloResponse, error) {
		input := &HelloWorkflowInput{
			Req: req,
			Goodbye: &GoodbyeSignal{
				Channel: workflow.GetSignalChannel(ctx, GoodbyeSignalName),
			},
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

// HelloWorkflowInput describes the input to a(n) example.v1.Hello workflow constructor
type HelloWorkflowInput struct {
	Req     *HelloRequest
	Goodbye *GoodbyeSignal
}

// Hello prints a friendly greeting and waits for goodbye
//
// workflow details: (name: "example.v1.Hello", id: "hello/${! name.or("World") }")
type HelloWorkflow interface {
	// Execute defines the entrypoint to a(n) example.v1.Hello workflow
	Execute(ctx workflow.Context) (*HelloResponse, error)
}

// Hello prints a friendly greeting and waits for goodbye
func HelloChild(ctx workflow.Context, req *HelloRequest, options ...*HelloChildOptions) (*HelloResponse, error) {
	childRun, err := HelloChildAsync(ctx, req, options...)
	if err != nil {
		return nil, err
	}
	return childRun.Get(ctx)
}

// Hello prints a friendly greeting and waits for goodbye
func HelloChildAsync(ctx workflow.Context, req *HelloRequest, options ...*HelloChildOptions) (*HelloChildRun, error) {
	var o *HelloChildOptions
	if len(options) > 0 && options[0] != nil {
		o = options[0]
	} else {
		o = NewHelloChildOptions()
	}
	opts, err := o.Build(ctx, req.ProtoReflect())
	if err != nil {
		return nil, fmt.Errorf("error initializing workflow.ChildWorkflowOptions: %w", err)
	}
	ctx = workflow.WithChildOptions(ctx, opts)
	return &HelloChildRun{Future: workflow.ExecuteChildWorkflow(ctx, HelloWorkflowName, req)}, nil
}

// HelloChildOptions provides configuration for a child example.v1.Hello workflow operation
type HelloChildOptions struct {
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

// NewHelloChildOptions initializes a new HelloChildOptions value
func NewHelloChildOptions() *HelloChildOptions {
	return &HelloChildOptions{}
}

// Build initializes a new go.temporal.io/sdk/workflow.ChildWorkflowOptions value with defaults and overrides applied
func (o *HelloChildOptions) Build(ctx workflow.Context, req protoreflect.Message) (workflow.ChildWorkflowOptions, error) {
	opts := o.options
	if v := o.id; v != nil {
		opts.WorkflowID = *v
	} else if opts.WorkflowID == "" {
		// wrap expression evaluation in local activity
		// more info: https://cludden.github.io/protoc-gen-go-temporal/docs/guides/patches#pv_64-expression-evaluation-local-activity
		if workflow.GetVersion(ctx, "cludden_protoc-gen-go-temporal_64_expression-evaluation-local-activity", workflow.DefaultVersion, 1) == 1 {
			lao := workflow.GetLocalActivityOptions(ctx)
			lao.ScheduleToCloseTimeout = time.Second * 10
			if err := workflow.ExecuteLocalActivity(workflow.WithLocalActivityOptions(ctx, lao), func(ctx context.Context) (string, error) {
				id, err := expression.EvalExpression(HelloIdexpression, req)
				if err != nil {
					return "", fmt.Errorf("error evaluating id expression for %q workflow: %w", HelloWorkflowName, err)
				}
				return id, nil
			}).Get(ctx, &opts.WorkflowID); err != nil {
				return opts, fmt.Errorf("error evaluating id expression for %q workflow: %w", HelloWorkflowName, err)
			}
		} else {
			id, err := expression.EvalExpression(HelloIdexpression, req)
			if err != nil {
				return opts, fmt.Errorf("error evaluating id expression for %q workflow: %w", HelloWorkflowName, err)
			}
			opts.WorkflowID = id
		}
	}
	if v := o.idReusePolicy; v != enumsv1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v
	}
	if v := o.taskQueue; v != nil {
		opts.TaskQueue = *v
	} else if opts.TaskQueue == "" {
		opts.TaskQueue = ExampleTaskQueue
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
func (o *HelloChildOptions) WithChildWorkflowOptions(options workflow.ChildWorkflowOptions) *HelloChildOptions {
	o.options = options
	return o
}

// WithExecutionTimeout sets the WorkflowExecutionTimeout value
func (o *HelloChildOptions) WithExecutionTimeout(d time.Duration) *HelloChildOptions {
	o.executionTimeout = &d
	return o
}

// WithID sets the WorkflowID value
func (o *HelloChildOptions) WithID(id string) *HelloChildOptions {
	o.id = &id
	return o
}

// WithIDReusePolicy sets the WorkflowIDReusePolicy value
func (o *HelloChildOptions) WithIDReusePolicy(policy enumsv1.WorkflowIdReusePolicy) *HelloChildOptions {
	o.idReusePolicy = policy
	return o
}

// WithParentClosePolicy sets the WorkflowIDReusePolicy value
func (o *HelloChildOptions) WithParentClosePolicy(policy enumsv1.ParentClosePolicy) *HelloChildOptions {
	o.parentClosePolicy = policy
	return o
}

// WithRetryPolicy sets the RetryPolicy value
func (o *HelloChildOptions) WithRetryPolicy(policy *temporal.RetryPolicy) *HelloChildOptions {
	o.retryPolicy = policy
	return o
}

// WithRunTimeout sets the WorkflowRunTimeout value
func (o *HelloChildOptions) WithRunTimeout(d time.Duration) *HelloChildOptions {
	o.runTimeout = &d
	return o
}

// WithSearchAttributes sets the SearchAttributes value
func (o *HelloChildOptions) WithSearchAttributes(sa map[string]any) *HelloChildOptions {
	o.searchAttributes = sa
	return o
}

// WithTaskTimeout sets the WorkflowTaskTimeout value
func (o *HelloChildOptions) WithTaskTimeout(d time.Duration) *HelloChildOptions {
	o.taskTimeout = &d
	return o
}

// WithTaskQueue sets the TaskQueue value
func (o *HelloChildOptions) WithTaskQueue(tq string) *HelloChildOptions {
	o.taskQueue = &tq
	return o
}

// WithWaitForCancellation sets the WaitForCancellation value
func (o *HelloChildOptions) WithWaitForCancellation(wait bool) *HelloChildOptions {
	o.waitForCancellation = &wait
	return o
}

// HelloChildRun describes a child Hello workflow run
type HelloChildRun struct {
	Future workflow.ChildWorkflowFuture
}

// Get blocks until the workflow is completed, returning the response value
func (r *HelloChildRun) Get(ctx workflow.Context) (*HelloResponse, error) {
	var resp HelloResponse
	if err := r.Future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Select adds this completion to the selector. Callback can be nil.
func (r *HelloChildRun) Select(sel workflow.Selector, fn func(*HelloChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future, func(workflow.Future) {
		if fn != nil {
			fn(r)
		}
	})
}

// SelectStart adds waiting for start to the selector. Callback can be nil.
func (r *HelloChildRun) SelectStart(sel workflow.Selector, fn func(*HelloChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future.GetChildWorkflowExecution(), func(workflow.Future) {
		if fn != nil {
			fn(r)
		}
	})
}

// WaitStart waits for the child workflow to start
func (r *HelloChildRun) WaitStart(ctx workflow.Context) (*workflow.Execution, error) {
	var exec workflow.Execution
	if err := r.Future.GetChildWorkflowExecution().Get(ctx, &exec); err != nil {
		return nil, err
	}
	return &exec, nil
}

// Goodbye sends a(n) "example.helloworld.v1.Example.Goodbye" signal request to the child workflow
func (r *HelloChildRun) Goodbye(ctx workflow.Context, input *GoodbyeRequest) error {
	return r.GoodbyeAsync(ctx, input).Get(ctx, nil)
}

// GoodbyeAsync sends a(n) "example.helloworld.v1.Example.Goodbye" signal request to the child workflow
func (r *HelloChildRun) GoodbyeAsync(ctx workflow.Context, input *GoodbyeRequest) workflow.Future {
	return r.Future.SignalChildWorkflow(ctx, GoodbyeSignalName, input)
}

// GoodbyeSignal describes a(n) example.helloworld.v1.Example.Goodbye signal
type GoodbyeSignal struct {
	Channel workflow.ReceiveChannel
}

// NewGoodbyeSignal initializes a new example.helloworld.v1.Example.Goodbye signal wrapper
func NewGoodbyeSignal(ctx workflow.Context) *GoodbyeSignal {
	return &GoodbyeSignal{Channel: workflow.GetSignalChannel(ctx, GoodbyeSignalName)}
}

// Receive blocks until a(n) example.helloworld.v1.Example.Goodbye signal is received
func (s *GoodbyeSignal) Receive(ctx workflow.Context) (*GoodbyeRequest, bool) {
	var resp GoodbyeRequest
	more := s.Channel.Receive(ctx, &resp)
	return &resp, more
}

// ReceiveAsync checks for a example.helloworld.v1.Example.Goodbye signal without blocking
func (s *GoodbyeSignal) ReceiveAsync() *GoodbyeRequest {
	var resp GoodbyeRequest
	if ok := s.Channel.ReceiveAsync(&resp); !ok {
		return nil
	}
	return &resp
}

// ReceiveWithTimeout blocks until a(n) example.helloworld.v1.Example.Goodbye signal is received or timeout expires.
// Returns more value of false when Channel is closed.
// Returns ok value of false when no value was found in the channel for the duration of timeout or the ctx was canceled.
// resp will be nil if ok is false.
func (s *GoodbyeSignal) ReceiveWithTimeout(ctx workflow.Context, timeout time.Duration) (resp *GoodbyeRequest, ok bool, more bool) {
	resp = &GoodbyeRequest{}
	if ok, more = s.Channel.ReceiveWithTimeout(ctx, timeout, &resp); !ok {
		return nil, false, more
	}
	return
}

// Select checks for a(n) example.helloworld.v1.Example.Goodbye signal without blocking
func (s *GoodbyeSignal) Select(sel workflow.Selector, fn func(*GoodbyeRequest)) workflow.Selector {
	return sel.AddReceive(s.Channel, func(workflow.ReceiveChannel, bool) {
		req := s.ReceiveAsync()
		if fn != nil {
			fn(req)
		}
	})
}

// Goodbye signals a running workflow to exit
func GoodbyeExternal(ctx workflow.Context, workflowID string, runID string, req *GoodbyeRequest) error {
	return GoodbyeExternalAsync(ctx, workflowID, runID, req).Get(ctx, nil)
}

// Goodbye signals a running workflow to exit
func GoodbyeExternalAsync(ctx workflow.Context, workflowID string, runID string, req *GoodbyeRequest) workflow.Future {
	return workflow.SignalExternalWorkflow(ctx, workflowID, runID, GoodbyeSignalName, req)
}

// ExampleActivities describes available worker activities
type ExampleActivities interface{}

// RegisterExampleActivities registers activities with a worker
func RegisterExampleActivities(r worker.ActivityRegistry, activities ExampleActivities) {}

// =============================================================================
// Test Client (example.helloworld.v1.Example)
// =============================================================================

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

// Hello executes a(n) example.v1.Hello workflow in the test environment
func (c *TestExampleClient) Hello(ctx context.Context, req *HelloRequest, opts ...*HelloOptions) (*HelloResponse, error) {
	run, err := c.HelloAsync(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// HelloAsync executes a(n) example.v1.Hello workflow in the test environment
func (c *TestExampleClient) HelloAsync(ctx context.Context, req *HelloRequest, options ...*HelloOptions) (HelloRun, error) {
	var o *HelloOptions
	if len(options) > 0 && options[0] != nil {
		o = options[0]
	} else {
		o = NewHelloOptions()
	}
	opts, err := o.Build(req.ProtoReflect())
	if err != nil {
		return nil, fmt.Errorf("error initializing client.StartWorkflowOptions: %w", err)
	}
	return &testHelloRun{client: c, env: c.env, opts: &opts, req: req, workflows: c.workflows}, nil
}

// GetHello is a noop
func (c *TestExampleClient) GetHello(ctx context.Context, workflowID string, runID string) HelloRun {
	return &testHelloRun{env: c.env, workflows: c.workflows}
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

// Goodbye executes a example.helloworld.v1.Example.Goodbye signal
func (c *TestExampleClient) Goodbye(ctx context.Context, workflowID string, runID string, req *GoodbyeRequest) error {
	c.env.SignalWorkflow(GoodbyeSignalName, req)
	return nil
}

var _ HelloRun = &testHelloRun{}

// testHelloRun provides convenience methods for interacting with a(n) example.v1.Hello workflow in the test environment
type testHelloRun struct {
	client    *TestExampleClient
	env       *testsuite.TestWorkflowEnvironment
	opts      *client.StartWorkflowOptions
	req       *HelloRequest
	workflows ExampleWorkflows
}

// Cancel requests cancellation of a workflow in execution, returning an error if applicable
func (r *testHelloRun) Cancel(ctx context.Context) error {
	return r.client.CancelWorkflow(ctx, r.ID(), r.RunID())
}

// Get retrieves a test example.v1.Hello workflow result
func (r *testHelloRun) Get(context.Context) (*HelloResponse, error) {
	r.env.ExecuteWorkflow(HelloWorkflowName, r.req)
	if !r.env.IsWorkflowCompleted() {
		return nil, errors.New("workflow in progress")
	}
	if err := r.env.GetWorkflowError(); err != nil {
		return nil, err
	}
	var result HelloResponse
	if err := r.env.GetWorkflowResult(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ID returns a test example.v1.Hello workflow run's workflow ID
func (r *testHelloRun) ID() string {
	if r.opts != nil {
		return r.opts.ID
	}
	return ""
}

// Run noop implementation
func (r *testHelloRun) Run() client.WorkflowRun {
	return nil
}

// RunID noop implementation
func (r *testHelloRun) RunID() string {
	return ""
}

// Terminate terminates a workflow in execution, returning an error if applicable
func (r *testHelloRun) Terminate(ctx context.Context, reason string, details ...interface{}) error {
	return r.client.TerminateWorkflow(ctx, r.ID(), r.RunID(), reason, details...)
}

// Goodbye executes a example.helloworld.v1.Example.Goodbye signal against a test example.v1.Hello workflow
func (r *testHelloRun) Goodbye(ctx context.Context, req *GoodbyeRequest) error {
	return r.client.Goodbye(ctx, r.ID(), r.RunID(), req)
}

// =============================================================================
// CLI (example.helloworld.v1.Example)
// =============================================================================

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
		{
			Name:                   "goodbye",
			Usage:                  "Goodbye signals a running workflow to exit",
			Category:               "SIGNALS",
			UseShortOptionHandling: true,
			Before:                 opts.before,
			After:                  opts.after,
			Flags: []v2.Flag{
				&v2.StringFlag{
					Name:     "workflow-id",
					Usage:    "workflow id",
					Required: true,
					Aliases:  []string{"w"},
				},
				&v2.StringFlag{
					Name:    "run-id",
					Usage:   "run id",
					Aliases: []string{"r"},
				},
				&v2.StringFlag{
					Name:    "input-file",
					Usage:   "path to json-formatted input file",
					Aliases: []string{"f"},
				},
				&v2.StringFlag{
					Name:     "message",
					Usage:    "set the value of the operation's \"Message\" parameter",
					Category: "INPUT",
				},
			},
			Action: func(cmd *v2.Context) error {
				c, err := opts.clientForCommand(cmd)
				if err != nil {
					return fmt.Errorf("error initializing client for command: %w", err)
				}
				defer c.Close()
				client := NewExampleClient(c)
				req, err := UnmarshalCliFlagsToGoodbyeRequest(cmd)
				if err != nil {
					return fmt.Errorf("error unmarshalling request: %w", err)
				}
				if err := client.Goodbye(cmd.Context, cmd.String("workflow-id"), cmd.String("run-id"), req); err != nil {
					return fmt.Errorf("error sending %q signal: %w", GoodbyeSignalName, err)
				}
				fmt.Println("success")
				return nil
			},
		},
		{
			Name:                   "hello",
			Usage:                  "Hello prints a friendly greeting and waits for goodbye",
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
				req, err := UnmarshalCliFlagsToHelloRequest(cmd)
				if err != nil {
					return fmt.Errorf("error unmarshalling request: %w", err)
				}
				opts := client.StartWorkflowOptions{}
				if tq := cmd.String("task-queue"); tq != "" {
					opts.TaskQueue = tq
				}
				run, err := c.HelloAsync(cmd.Context, req, NewHelloOptions().WithStartWorkflowOptions(opts))
				if err != nil {
					return fmt.Errorf("error starting %s workflow: %w", HelloWorkflowName, err)
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

// UnmarshalCliFlagsToGoodbyeRequest unmarshals a GoodbyeRequest from command line flags
func UnmarshalCliFlagsToGoodbyeRequest(cmd *v2.Context) (*GoodbyeRequest, error) {
	var result GoodbyeRequest
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
	if cmd.IsSet("message") {
		hasValues = true
		result.Message = cmd.String("message")
	}
	if !hasValues {
		return nil, nil
	}
	return &result, nil
}

// UnmarshalCliFlagsToHelloRequest unmarshals a HelloRequest from command line flags
func UnmarshalCliFlagsToHelloRequest(cmd *v2.Context) (*HelloRequest, error) {
	var result HelloRequest
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

// =============================================================================
// Codec (example.helloworld.v1.Example)
// =============================================================================

// WithExampleSchemeTypes registers all Example protobuf types with the given scheme
func WithExampleSchemeTypes() scheme.Option {
	return func(s *scheme.Scheme) {
		s.RegisterType(File_example_helloworld_v1_example_proto.Messages().ByName("GoodbyeRequest"))
		s.RegisterType(File_example_helloworld_v1_example_proto.Messages().ByName("HelloRequest"))
		s.RegisterType(File_example_helloworld_v1_example_proto.Messages().ByName("HelloResponse"))
	}
}
