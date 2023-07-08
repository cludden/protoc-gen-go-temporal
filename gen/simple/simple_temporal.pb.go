// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 0.8.1-next (85d262dc2134d7428579c049b2e8eb0115fce7cc)
//	go go1.20.4
//	protoc (unknown)
//
// source: simple/simple.proto
package simple

import (
	"context"
	"errors"
	"fmt"
	expression "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	v1 "go.temporal.io/api/enums/v1"
	activity "go.temporal.io/sdk/activity"
	client "go.temporal.io/sdk/client"
	temporal "go.temporal.io/sdk/temporal"
	worker "go.temporal.io/sdk/worker"
	workflow "go.temporal.io/sdk/workflow"
)

// SimpleTaskQueue is the default task-queue for a Simple worker
const SimpleTaskQueue = "my-task-queue"

// Simple workflow names
const (
	SomeWorkflow1WorkflowName = "mycompany.simple.SomeWorkflow1"
	SomeWorkflow2WorkflowName = "mycompany.simple.SomeWorkflow2"
	SomeWorkflow3WorkflowName = "mycompany.simple.Simple.SomeWorkflow3Workflow"
)

// Simple id expressions
var (
	SomeWorkflow1IDExpression = expression.MustParseExpression("some-workflow-1/${!id}/${!uuid_v4()}")
	SomeWorkflow3IDExpression = expression.MustParseExpression("some-workflow-3/${!id}/${!request_val}")
)

// Simple query names
const (
	SomeQuery1QueryName = "mycompany.simple.Simple.SomeQuery1Query"
	SomeQuery2QueryName = "mycompany.simple.Simple.SomeQuery2Query"
)

// Simple signal names
const (
	SomeSignal1SignalName = "mycompany.simple.Simple.SomeSignal1Signal"
	SomeSignal2SignalName = "mycompany.simple.Simple.SomeSignal2Signal"
)

// Simple activity names
const (
	SomeActivity1ActivityName = "mycompany.simple.SomeActivity1"
	SomeActivity2ActivityName = "mycompany.simple.Simple.SomeActivity2Activity"
	SomeActivity3ActivityName = "mycompany.simple.Simple.SomeActivity3Activity"
)

// Client describes a client for a Simple worker
type Client interface {
	// SomeWorkflow1 does some workflow thing.
	SomeWorkflow1(ctx context.Context, opts *client.StartWorkflowOptions, req *SomeWorkflow1Request) (*SomeWorkflow1Response, error)
	// ExecuteSomeWorkflow1 executes a SomeWorkflow1 workflow
	ExecuteSomeWorkflow1(ctx context.Context, opts *client.StartWorkflowOptions, req *SomeWorkflow1Request) (SomeWorkflow1Run, error)
	// GetSomeWorkflow1 retrieves a SomeWorkflow1 workflow execution
	GetSomeWorkflow1(ctx context.Context, workflowID string, runID string) (SomeWorkflow1Run, error)
	// SomeWorkflow2 does some workflow thing.
	SomeWorkflow2(ctx context.Context, opts *client.StartWorkflowOptions) error
	// ExecuteSomeWorkflow2 executes a SomeWorkflow2 workflow
	ExecuteSomeWorkflow2(ctx context.Context, opts *client.StartWorkflowOptions) (SomeWorkflow2Run, error)
	// GetSomeWorkflow2 retrieves a SomeWorkflow2 workflow execution
	GetSomeWorkflow2(ctx context.Context, workflowID string, runID string) (SomeWorkflow2Run, error)
	// StartSomeWorkflow2WithSomeSignal1 sends a SomeSignal1 signal to a SomeWorkflow2 workflow, starting it if not present
	StartSomeWorkflow2WithSomeSignal1(ctx context.Context, opts *client.StartWorkflowOptions) (SomeWorkflow2Run, error)
	// SomeWorkflow3 does some workflow thing.
	SomeWorkflow3(ctx context.Context, opts *client.StartWorkflowOptions, req *SomeWorkflow3Request) error
	// ExecuteSomeWorkflow3 executes a SomeWorkflow3 workflow
	ExecuteSomeWorkflow3(ctx context.Context, opts *client.StartWorkflowOptions, req *SomeWorkflow3Request) (SomeWorkflow3Run, error)
	// GetSomeWorkflow3 retrieves a SomeWorkflow3 workflow execution
	GetSomeWorkflow3(ctx context.Context, workflowID string, runID string) (SomeWorkflow3Run, error)
	// StartSomeWorkflow3WithSomeSignal2 sends a SomeSignal2 signal to a SomeWorkflow3 workflow, starting it if not present
	StartSomeWorkflow3WithSomeSignal2(ctx context.Context, opts *client.StartWorkflowOptions, req *SomeWorkflow3Request, signal *SomeSignal2Request) (SomeWorkflow3Run, error)
	// QuerySomeQuery1 sends a SomeQuery1 query to an existing workflow
	QuerySomeQuery1(ctx context.Context, workflowID string, runID string) (*SomeQuery1Response, error)
	// QuerySomeQuery2 sends a SomeQuery2 query to an existing workflow
	QuerySomeQuery2(ctx context.Context, workflowID string, runID string, query *SomeQuery2Request) (*SomeQuery2Response, error)
	// SignalSomeSignal1 sends a SomeSignal1 signal to an existing workflow
	SignalSomeSignal1(ctx context.Context, workflowID string, runID string) error
	// SignalSomeSignal2 sends a SomeSignal2 signal to an existing workflow
	SignalSomeSignal2(ctx context.Context, workflowID string, runID string, signal *SomeSignal2Request) error
}

// Compile-time check that workflowClient satisfies Client
var _ Client = &workflowClient{}

// workflowClient implements a temporal client for a Simple service
type workflowClient struct {
	client client.Client
}

// NewClient initializes a new Simple client
func NewClient(c client.Client) Client {
	return &workflowClient{client: c}
}

// NewClientWithOptions initializes a new Simple client with the given options
func NewClientWithOptions(c client.Client, opts client.Options) (Client, error) {
	var err error
	c, err = client.NewClientFromExisting(c, opts)
	if err != nil {
		return nil, fmt.Errorf("error initializing client with options: %w", err)
	}
	return &workflowClient{client: c}, nil
}

// SomeWorkflow1 does some workflow thing.
func (c *workflowClient) SomeWorkflow1(ctx context.Context, opts *client.StartWorkflowOptions, req *SomeWorkflow1Request) (*SomeWorkflow1Response, error) {
	run, err := c.ExecuteSomeWorkflow1(ctx, opts, req)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// ExecuteSomeWorkflow1 starts a SomeWorkflow1 workflow
func (c *workflowClient) ExecuteSomeWorkflow1(ctx context.Context, opts *client.StartWorkflowOptions, req *SomeWorkflow1Request) (SomeWorkflow1Run, error) {
	if opts == nil {
		opts = &client.StartWorkflowOptions{}
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "my-task-queue"
	}
	if opts.ID == "" {
		id, err := expression.EvalExpression(SomeWorkflow1IDExpression, req.ProtoReflect())
		if err != nil {
			return nil, err
		}
		opts.ID = id
	}
	run, err := c.client.ExecuteWorkflow(ctx, *opts, SomeWorkflow1WorkflowName, req)
	if err != nil {
		return nil, err
	}
	if run == nil {
		return nil, errors.New("execute workflow returned nil run")
	}
	return &someWorkflow1Run{
		client: c,
		run:    run,
	}, nil
}

// GetSomeWorkflow1 fetches an existing SomeWorkflow1 execution
func (c *workflowClient) GetSomeWorkflow1(ctx context.Context, workflowID string, runID string) (SomeWorkflow1Run, error) {
	return &someWorkflow1Run{
		client: c,
		run:    c.client.GetWorkflow(ctx, workflowID, runID),
	}, nil
}

// SomeWorkflow2 does some workflow thing.
func (c *workflowClient) SomeWorkflow2(ctx context.Context, opts *client.StartWorkflowOptions) error {
	run, err := c.ExecuteSomeWorkflow2(ctx, opts)
	if err != nil {
		return err
	}
	return run.Get(ctx)
}

// ExecuteSomeWorkflow2 starts a SomeWorkflow2 workflow
func (c *workflowClient) ExecuteSomeWorkflow2(ctx context.Context, opts *client.StartWorkflowOptions) (SomeWorkflow2Run, error) {
	if opts == nil {
		opts = &client.StartWorkflowOptions{}
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "my-task-queue"
	}
	run, err := c.client.ExecuteWorkflow(ctx, *opts, SomeWorkflow2WorkflowName)
	if err != nil {
		return nil, err
	}
	if run == nil {
		return nil, errors.New("execute workflow returned nil run")
	}
	return &someWorkflow2Run{
		client: c,
		run:    run,
	}, nil
}

// GetSomeWorkflow2 fetches an existing SomeWorkflow2 execution
func (c *workflowClient) GetSomeWorkflow2(ctx context.Context, workflowID string, runID string) (SomeWorkflow2Run, error) {
	return &someWorkflow2Run{
		client: c,
		run:    c.client.GetWorkflow(ctx, workflowID, runID),
	}, nil
}

// StartSomeWorkflow2WithSomeSignal1 starts a SomeWorkflow2 workflow and sends a SomeSignal1 signal in a transaction
func (c *workflowClient) StartSomeWorkflow2WithSomeSignal1(ctx context.Context, opts *client.StartWorkflowOptions) (SomeWorkflow2Run, error) {
	if opts == nil {
		opts = &client.StartWorkflowOptions{}
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "my-task-queue"
	}
	run, err := c.client.SignalWithStartWorkflow(ctx, opts.ID, SomeSignal1SignalName, nil, *opts, SomeWorkflow2WorkflowName)
	if run == nil || err != nil {
		return nil, err
	}
	return &someWorkflow2Run{
		client: c,
		run:    run,
	}, nil
}

// SomeWorkflow3 does some workflow thing.
func (c *workflowClient) SomeWorkflow3(ctx context.Context, opts *client.StartWorkflowOptions, req *SomeWorkflow3Request) error {
	run, err := c.ExecuteSomeWorkflow3(ctx, opts, req)
	if err != nil {
		return err
	}
	return run.Get(ctx)
}

// ExecuteSomeWorkflow3 starts a SomeWorkflow3 workflow
func (c *workflowClient) ExecuteSomeWorkflow3(ctx context.Context, opts *client.StartWorkflowOptions, req *SomeWorkflow3Request) (SomeWorkflow3Run, error) {
	if opts == nil {
		opts = &client.StartWorkflowOptions{}
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "my-task-queue-2"
	}
	if opts.ID == "" {
		id, err := expression.EvalExpression(SomeWorkflow3IDExpression, req.ProtoReflect())
		if err != nil {
			return nil, err
		}
		opts.ID = id
	}
	if opts.WorkflowIDReusePolicy == v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v1.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
	}
	if opts.RetryPolicy == nil {
		opts.RetryPolicy = &temporal.RetryPolicy{MaximumAttempts: int32(2)}
	}
	if opts.WorkflowExecutionTimeout == 0 {
		opts.WorkflowRunTimeout = 3600000000000 // 1h0m0s
	}
	run, err := c.client.ExecuteWorkflow(ctx, *opts, SomeWorkflow3WorkflowName, req)
	if err != nil {
		return nil, err
	}
	if run == nil {
		return nil, errors.New("execute workflow returned nil run")
	}
	return &someWorkflow3Run{
		client: c,
		run:    run,
	}, nil
}

// GetSomeWorkflow3 fetches an existing SomeWorkflow3 execution
func (c *workflowClient) GetSomeWorkflow3(ctx context.Context, workflowID string, runID string) (SomeWorkflow3Run, error) {
	return &someWorkflow3Run{
		client: c,
		run:    c.client.GetWorkflow(ctx, workflowID, runID),
	}, nil
}

// StartSomeWorkflow3WithSomeSignal2 starts a SomeWorkflow3 workflow and sends a SomeSignal2 signal in a transaction
func (c *workflowClient) StartSomeWorkflow3WithSomeSignal2(ctx context.Context, opts *client.StartWorkflowOptions, req *SomeWorkflow3Request, signal *SomeSignal2Request) (SomeWorkflow3Run, error) {
	if opts == nil {
		opts = &client.StartWorkflowOptions{}
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "my-task-queue-2"
	}
	if opts.ID == "" {
		id, err := expression.EvalExpression(SomeWorkflow3IDExpression, req.ProtoReflect())
		if err != nil {
			return nil, err
		}
		opts.ID = id
	}
	if opts.WorkflowIDReusePolicy == v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v1.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
	}
	if opts.RetryPolicy == nil {
		opts.RetryPolicy = &temporal.RetryPolicy{MaximumAttempts: int32(2)}
	}
	if opts.WorkflowExecutionTimeout == 0 {
		opts.WorkflowRunTimeout = 3600000000000 // 1h0m0s
	}
	run, err := c.client.SignalWithStartWorkflow(ctx, opts.ID, SomeSignal2SignalName, signal, *opts, SomeWorkflow3WorkflowName, req)
	if run == nil || err != nil {
		return nil, err
	}
	return &someWorkflow3Run{
		client: c,
		run:    run,
	}, nil
}

// QuerySomeQuery1 sends a SomeQuery1 query to an existing workflow
func (c *workflowClient) QuerySomeQuery1(ctx context.Context, workflowID string, runID string) (*SomeQuery1Response, error) {
	var resp SomeQuery1Response
	if val, err := c.client.QueryWorkflow(ctx, workflowID, runID, SomeQuery1QueryName); err != nil {
		return nil, err
	} else if err = val.Get(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QuerySomeQuery2 sends a SomeQuery2 query to an existing workflow
func (c *workflowClient) QuerySomeQuery2(ctx context.Context, workflowID string, runID string, query *SomeQuery2Request) (*SomeQuery2Response, error) {
	var resp SomeQuery2Response
	if val, err := c.client.QueryWorkflow(ctx, workflowID, runID, SomeQuery2QueryName, query); err != nil {
		return nil, err
	} else if err = val.Get(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SignalSomeSignal1 sends a SomeSignal1 signal to an existing workflow
func (c *workflowClient) SignalSomeSignal1(ctx context.Context, workflowID string, runID string) error {
	return c.client.SignalWorkflow(ctx, workflowID, runID, SomeSignal1SignalName, nil)
}

// SignalSomeSignal2 sends a SomeSignal2 signal to an existing workflow
func (c *workflowClient) SignalSomeSignal2(ctx context.Context, workflowID string, runID string, signal *SomeSignal2Request) error {
	return c.client.SignalWorkflow(ctx, workflowID, runID, SomeSignal2SignalName, signal)
}

// SomeWorkflow1Run describes a SomeWorkflow1 workflow run
type SomeWorkflow1Run interface {
	// ID returns the workflow ID
	ID() string
	// RunID returns the workflow instance ID
	RunID() string
	// Get blocks until the workflow is complete and returns the result
	Get(ctx context.Context) (*SomeWorkflow1Response, error)
	// SomeQuery1 runs the SomeQuery1 query against the workflow
	SomeQuery1(ctx context.Context) (*SomeQuery1Response, error)
	// SomeQuery2 runs the SomeQuery2 query against the workflow
	SomeQuery2(ctx context.Context, req *SomeQuery2Request) (*SomeQuery2Response, error)
	// SomeSignal1 sends a SomeSignal1 signal to the workflow
	SomeSignal1(ctx context.Context) error
	// SomeSignal2 sends a SomeSignal2 signal to the workflow
	SomeSignal2(ctx context.Context, req *SomeSignal2Request) error
}

// someWorkflow1Run provides an internal implementation of a SomeWorkflow1Run
type someWorkflow1Run struct {
	client *workflowClient
	run    client.WorkflowRun
}

// ID returns the workflow ID
func (r *someWorkflow1Run) ID() string {
	return r.run.GetID()
}

// RunID returns the execution ID
func (r *someWorkflow1Run) RunID() string {
	return r.run.GetRunID()
}

// Get blocks until the workflow is complete, returning the result if applicable
func (r *someWorkflow1Run) Get(ctx context.Context) (*SomeWorkflow1Response, error) {
	var resp SomeWorkflow1Response
	if err := r.run.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SomeQuery1 executes a SomeQuery1 query against the workflow
func (r *someWorkflow1Run) SomeQuery1(ctx context.Context) (*SomeQuery1Response, error) {
	return r.client.QuerySomeQuery1(ctx, r.ID(), "")
}

// SomeQuery2 executes a SomeQuery2 query against the workflow
func (r *someWorkflow1Run) SomeQuery2(ctx context.Context, req *SomeQuery2Request) (*SomeQuery2Response, error) {
	return r.client.QuerySomeQuery2(ctx, r.ID(), "", req)
}

// SomeSignal1 sends a SomeSignal1 signal to the workflow
func (r *someWorkflow1Run) SomeSignal1(ctx context.Context) error {
	return r.client.SignalSomeSignal1(ctx, r.ID(), "")
}

// SomeSignal2 sends a SomeSignal2 signal to the workflow
func (r *someWorkflow1Run) SomeSignal2(ctx context.Context, req *SomeSignal2Request) error {
	return r.client.SignalSomeSignal2(ctx, r.ID(), "", req)
}

// SomeWorkflow2Run describes a SomeWorkflow2 workflow run
type SomeWorkflow2Run interface {
	// ID returns the workflow ID
	ID() string
	// RunID returns the workflow instance ID
	RunID() string
	// Get blocks until the workflow is complete and returns the result
	Get(ctx context.Context) error
	// SomeSignal1 sends a SomeSignal1 signal to the workflow
	SomeSignal1(ctx context.Context) error
}

// someWorkflow2Run provides an internal implementation of a SomeWorkflow2Run
type someWorkflow2Run struct {
	client *workflowClient
	run    client.WorkflowRun
}

// ID returns the workflow ID
func (r *someWorkflow2Run) ID() string {
	return r.run.GetID()
}

// RunID returns the execution ID
func (r *someWorkflow2Run) RunID() string {
	return r.run.GetRunID()
}

// Get blocks until the workflow is complete, returning the result if applicable
func (r *someWorkflow2Run) Get(ctx context.Context) error {
	return r.run.Get(ctx, nil)
}

// SomeSignal1 sends a SomeSignal1 signal to the workflow
func (r *someWorkflow2Run) SomeSignal1(ctx context.Context) error {
	return r.client.SignalSomeSignal1(ctx, r.ID(), "")
}

// SomeWorkflow3Run describes a SomeWorkflow3 workflow run
type SomeWorkflow3Run interface {
	// ID returns the workflow ID
	ID() string
	// RunID returns the workflow instance ID
	RunID() string
	// Get blocks until the workflow is complete and returns the result
	Get(ctx context.Context) error
	// SomeSignal2 sends a SomeSignal2 signal to the workflow
	SomeSignal2(ctx context.Context, req *SomeSignal2Request) error
}

// someWorkflow3Run provides an internal implementation of a SomeWorkflow3Run
type someWorkflow3Run struct {
	client *workflowClient
	run    client.WorkflowRun
}

// ID returns the workflow ID
func (r *someWorkflow3Run) ID() string {
	return r.run.GetID()
}

// RunID returns the execution ID
func (r *someWorkflow3Run) RunID() string {
	return r.run.GetRunID()
}

// Get blocks until the workflow is complete, returning the result if applicable
func (r *someWorkflow3Run) Get(ctx context.Context) error {
	return r.run.Get(ctx, nil)
}

// SomeSignal2 sends a SomeSignal2 signal to the workflow
func (r *someWorkflow3Run) SomeSignal2(ctx context.Context, req *SomeSignal2Request) error {
	return r.client.SignalSomeSignal2(ctx, r.ID(), "", req)
}

// Workflows provides methods for initializing new Simple workflow values
type Workflows interface {
	// SomeWorkflow1 initializes a new SomeWorkflow1Workflow value
	SomeWorkflow1(ctx workflow.Context, input *SomeWorkflow1Input) (SomeWorkflow1Workflow, error)
	// SomeWorkflow2 initializes a new SomeWorkflow2Workflow value
	SomeWorkflow2(ctx workflow.Context, input *SomeWorkflow2Input) (SomeWorkflow2Workflow, error)
	// SomeWorkflow3 initializes a new SomeWorkflow3Workflow value
	SomeWorkflow3(ctx workflow.Context, input *SomeWorkflow3Input) (SomeWorkflow3Workflow, error)
}

// RegisterWorkflows registers Simple workflows with the given worker
func RegisterWorkflows(r worker.Registry, workflows Workflows) {
	RegisterSomeWorkflow1Workflow(r, workflows.SomeWorkflow1)
	RegisterSomeWorkflow2Workflow(r, workflows.SomeWorkflow2)
	RegisterSomeWorkflow3Workflow(r, workflows.SomeWorkflow3)
}

// RegisterSomeWorkflow1Workflow registers a SomeWorkflow1 workflow with the given worker
func RegisterSomeWorkflow1Workflow(r worker.Registry, wf func(workflow.Context, *SomeWorkflow1Input) (SomeWorkflow1Workflow, error)) {
	r.RegisterWorkflowWithOptions(buildSomeWorkflow1(wf), workflow.RegisterOptions{Name: SomeWorkflow1WorkflowName})
}

// buildSomeWorkflow1 converts a SomeWorkflow1 workflow struct into a valid workflow function
func buildSomeWorkflow1(wf func(workflow.Context, *SomeWorkflow1Input) (SomeWorkflow1Workflow, error)) func(workflow.Context, *SomeWorkflow1Request) (*SomeWorkflow1Response, error) {
	return (&someWorkflow1{wf}).SomeWorkflow1
}

// someWorkflow1 provides an SomeWorkflow1 method for calling the user's implementation
type someWorkflow1 struct {
	ctor func(workflow.Context, *SomeWorkflow1Input) (SomeWorkflow1Workflow, error)
}

// SomeWorkflow1 constructs a new SomeWorkflow1 value and executes it
func (w *someWorkflow1) SomeWorkflow1(ctx workflow.Context, req *SomeWorkflow1Request) (*SomeWorkflow1Response, error) {
	input := &SomeWorkflow1Input{
		Req: req,
		SomeSignal1: &SomeSignal1Signal{
			Channel: workflow.GetSignalChannel(ctx, SomeSignal1SignalName),
		},
		SomeSignal2: &SomeSignal2Signal{
			Channel: workflow.GetSignalChannel(ctx, SomeSignal2SignalName),
		},
	}
	wf, err := w.ctor(ctx, input)
	if err != nil {
		return nil, err
	}
	if err := workflow.SetQueryHandler(ctx, SomeQuery1QueryName, wf.SomeQuery1); err != nil {
		return nil, err
	}
	if err := workflow.SetQueryHandler(ctx, SomeQuery2QueryName, wf.SomeQuery2); err != nil {
		return nil, err
	}
	return wf.Execute(ctx)
}

// SomeWorkflow1Input describes the input to a SomeWorkflow1 workflow constructor
type SomeWorkflow1Input struct {
	Req         *SomeWorkflow1Request
	SomeSignal1 *SomeSignal1Signal
	SomeSignal2 *SomeSignal2Signal
}

// SomeWorkflow1 does some workflow thing.
type SomeWorkflow1Workflow interface {
	// Execute a SomeWorkflow1 workflow
	Execute(ctx workflow.Context) (*SomeWorkflow1Response, error)
	// SomeQuery1 query handler
	SomeQuery1() (*SomeQuery1Response, error)
	// SomeQuery2 query handler
	SomeQuery2(*SomeQuery2Request) (*SomeQuery2Response, error)
}

// SomeWorkflow1Child executes a child SomeWorkflow1 workflow
func SomeWorkflow1Child(ctx workflow.Context, opts *workflow.ChildWorkflowOptions, req *SomeWorkflow1Request) *SomeWorkflow1ChildRun {
	if opts == nil {
		childOpts := workflow.GetChildWorkflowOptions(ctx)
		opts = &childOpts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "my-task-queue"
	}
	if opts.WorkflowID == "" {
		id, err := expression.EvalExpression(SomeWorkflow1IDExpression, req.ProtoReflect())
		if err != nil {
			panic(err)
		}
		opts.WorkflowID = id
	}
	ctx = workflow.WithChildOptions(ctx, *opts)
	return &SomeWorkflow1ChildRun{Future: workflow.ExecuteChildWorkflow(ctx, SomeWorkflow1WorkflowName, req)}
}

// SomeWorkflow1ChildRun describes a child SomeWorkflow1 workflow run
type SomeWorkflow1ChildRun struct {
	Future workflow.ChildWorkflowFuture
}

// Get blocks until the workflow is completed, returning the response value
func (r *SomeWorkflow1ChildRun) Get(ctx workflow.Context) (*SomeWorkflow1Response, error) {
	var resp SomeWorkflow1Response
	if err := r.Future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Select adds this completion to the selector. Callback can be nil.
func (r *SomeWorkflow1ChildRun) Select(sel workflow.Selector, fn func(SomeWorkflow1ChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future, func(workflow.Future) {
		if fn != nil {
			fn(*r)
		}
	})
}

// SelectStart adds waiting for start to the selector. Callback can be nil.
func (r *SomeWorkflow1ChildRun) SelectStart(sel workflow.Selector, fn func(SomeWorkflow1ChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future.GetChildWorkflowExecution(), func(workflow.Future) {
		if fn != nil {
			fn(*r)
		}
	})
}

// WaitStart waits for the child workflow to start
func (r *SomeWorkflow1ChildRun) WaitStart(ctx workflow.Context) (*workflow.Execution, error) {
	var exec workflow.Execution
	if err := r.Future.GetChildWorkflowExecution().Get(ctx, &exec); err != nil {
		return nil, err
	}
	return &exec, nil
}

// SomeSignal1 sends the corresponding signal request to the child workflow
func (r *SomeWorkflow1ChildRun) SomeSignal1(ctx workflow.Context) workflow.Future {
	return r.Future.SignalChildWorkflow(ctx, SomeSignal1SignalName, nil)
}

// SomeSignal2 sends the corresponding signal request to the child workflow
func (r *SomeWorkflow1ChildRun) SomeSignal2(ctx workflow.Context, input *SomeSignal2Request) workflow.Future {
	return r.Future.SignalChildWorkflow(ctx, SomeSignal2SignalName, input)
}

// RegisterSomeWorkflow2Workflow registers a SomeWorkflow2 workflow with the given worker
func RegisterSomeWorkflow2Workflow(r worker.Registry, wf func(workflow.Context, *SomeWorkflow2Input) (SomeWorkflow2Workflow, error)) {
	r.RegisterWorkflowWithOptions(buildSomeWorkflow2(wf), workflow.RegisterOptions{Name: SomeWorkflow2WorkflowName})
}

// buildSomeWorkflow2 converts a SomeWorkflow2 workflow struct into a valid workflow function
func buildSomeWorkflow2(wf func(workflow.Context, *SomeWorkflow2Input) (SomeWorkflow2Workflow, error)) func(workflow.Context) error {
	return (&someWorkflow2{wf}).SomeWorkflow2
}

// someWorkflow2 provides an SomeWorkflow2 method for calling the user's implementation
type someWorkflow2 struct {
	ctor func(workflow.Context, *SomeWorkflow2Input) (SomeWorkflow2Workflow, error)
}

// SomeWorkflow2 constructs a new SomeWorkflow2 value and executes it
func (w *someWorkflow2) SomeWorkflow2(ctx workflow.Context) error {
	input := &SomeWorkflow2Input{
		SomeSignal1: &SomeSignal1Signal{
			Channel: workflow.GetSignalChannel(ctx, SomeSignal1SignalName),
		},
	}
	wf, err := w.ctor(ctx, input)
	if err != nil {
		return err
	}
	return wf.Execute(ctx)
}

// SomeWorkflow2Input describes the input to a SomeWorkflow2 workflow constructor
type SomeWorkflow2Input struct {
	SomeSignal1 *SomeSignal1Signal
}

// SomeWorkflow2 does some workflow thing.
type SomeWorkflow2Workflow interface {
	// Execute a SomeWorkflow2 workflow
	Execute(ctx workflow.Context) error
}

// SomeWorkflow2Child executes a child SomeWorkflow2 workflow
func SomeWorkflow2Child(ctx workflow.Context, opts *workflow.ChildWorkflowOptions) *SomeWorkflow2ChildRun {
	if opts == nil {
		childOpts := workflow.GetChildWorkflowOptions(ctx)
		opts = &childOpts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "my-task-queue"
	}
	ctx = workflow.WithChildOptions(ctx, *opts)
	return &SomeWorkflow2ChildRun{Future: workflow.ExecuteChildWorkflow(ctx, SomeWorkflow2WorkflowName, nil)}
}

// SomeWorkflow2ChildRun describes a child SomeWorkflow2 workflow run
type SomeWorkflow2ChildRun struct {
	Future workflow.ChildWorkflowFuture
}

// Get blocks until the workflow is completed, returning the response value
func (r *SomeWorkflow2ChildRun) Get(ctx workflow.Context) error {
	if err := r.Future.Get(ctx, nil); err != nil {
		return err
	}
	return nil
}

// Select adds this completion to the selector. Callback can be nil.
func (r *SomeWorkflow2ChildRun) Select(sel workflow.Selector, fn func(SomeWorkflow2ChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future, func(workflow.Future) {
		if fn != nil {
			fn(*r)
		}
	})
}

// SelectStart adds waiting for start to the selector. Callback can be nil.
func (r *SomeWorkflow2ChildRun) SelectStart(sel workflow.Selector, fn func(SomeWorkflow2ChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future.GetChildWorkflowExecution(), func(workflow.Future) {
		if fn != nil {
			fn(*r)
		}
	})
}

// WaitStart waits for the child workflow to start
func (r *SomeWorkflow2ChildRun) WaitStart(ctx workflow.Context) (*workflow.Execution, error) {
	var exec workflow.Execution
	if err := r.Future.GetChildWorkflowExecution().Get(ctx, &exec); err != nil {
		return nil, err
	}
	return &exec, nil
}

// SomeSignal1 sends the corresponding signal request to the child workflow
func (r *SomeWorkflow2ChildRun) SomeSignal1(ctx workflow.Context) workflow.Future {
	return r.Future.SignalChildWorkflow(ctx, SomeSignal1SignalName, nil)
}

// RegisterSomeWorkflow3Workflow registers a SomeWorkflow3 workflow with the given worker
func RegisterSomeWorkflow3Workflow(r worker.Registry, wf func(workflow.Context, *SomeWorkflow3Input) (SomeWorkflow3Workflow, error)) {
	r.RegisterWorkflowWithOptions(buildSomeWorkflow3(wf), workflow.RegisterOptions{Name: SomeWorkflow3WorkflowName})
}

// buildSomeWorkflow3 converts a SomeWorkflow3 workflow struct into a valid workflow function
func buildSomeWorkflow3(wf func(workflow.Context, *SomeWorkflow3Input) (SomeWorkflow3Workflow, error)) func(workflow.Context, *SomeWorkflow3Request) error {
	return (&someWorkflow3{wf}).SomeWorkflow3
}

// someWorkflow3 provides an SomeWorkflow3 method for calling the user's implementation
type someWorkflow3 struct {
	ctor func(workflow.Context, *SomeWorkflow3Input) (SomeWorkflow3Workflow, error)
}

// SomeWorkflow3 constructs a new SomeWorkflow3 value and executes it
func (w *someWorkflow3) SomeWorkflow3(ctx workflow.Context, req *SomeWorkflow3Request) error {
	input := &SomeWorkflow3Input{
		Req: req,
		SomeSignal2: &SomeSignal2Signal{
			Channel: workflow.GetSignalChannel(ctx, SomeSignal2SignalName),
		},
	}
	wf, err := w.ctor(ctx, input)
	if err != nil {
		return err
	}
	return wf.Execute(ctx)
}

// SomeWorkflow3Input describes the input to a SomeWorkflow3 workflow constructor
type SomeWorkflow3Input struct {
	Req         *SomeWorkflow3Request
	SomeSignal2 *SomeSignal2Signal
}

// SomeWorkflow3 does some workflow thing.
type SomeWorkflow3Workflow interface {
	// Execute a SomeWorkflow3 workflow
	Execute(ctx workflow.Context) error
}

// SomeWorkflow3Child executes a child SomeWorkflow3 workflow
func SomeWorkflow3Child(ctx workflow.Context, opts *workflow.ChildWorkflowOptions, req *SomeWorkflow3Request) *SomeWorkflow3ChildRun {
	if opts == nil {
		childOpts := workflow.GetChildWorkflowOptions(ctx)
		opts = &childOpts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "my-task-queue-2"
	}
	if opts.WorkflowID == "" {
		id, err := expression.EvalExpression(SomeWorkflow3IDExpression, req.ProtoReflect())
		if err != nil {
			panic(err)
		}
		opts.WorkflowID = id
	}
	if opts.WorkflowIDReusePolicy == v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v1.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
	}
	if opts.RetryPolicy == nil {
		opts.RetryPolicy = &temporal.RetryPolicy{MaximumAttempts: int32(2)}
	}
	if opts.WorkflowExecutionTimeout == 0 {
		opts.WorkflowRunTimeout = 3600000000000 // 1h0m0s
	}
	ctx = workflow.WithChildOptions(ctx, *opts)
	return &SomeWorkflow3ChildRun{Future: workflow.ExecuteChildWorkflow(ctx, SomeWorkflow3WorkflowName, req)}
}

// SomeWorkflow3ChildRun describes a child SomeWorkflow3 workflow run
type SomeWorkflow3ChildRun struct {
	Future workflow.ChildWorkflowFuture
}

// Get blocks until the workflow is completed, returning the response value
func (r *SomeWorkflow3ChildRun) Get(ctx workflow.Context) error {
	if err := r.Future.Get(ctx, nil); err != nil {
		return err
	}
	return nil
}

// Select adds this completion to the selector. Callback can be nil.
func (r *SomeWorkflow3ChildRun) Select(sel workflow.Selector, fn func(SomeWorkflow3ChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future, func(workflow.Future) {
		if fn != nil {
			fn(*r)
		}
	})
}

// SelectStart adds waiting for start to the selector. Callback can be nil.
func (r *SomeWorkflow3ChildRun) SelectStart(sel workflow.Selector, fn func(SomeWorkflow3ChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future.GetChildWorkflowExecution(), func(workflow.Future) {
		if fn != nil {
			fn(*r)
		}
	})
}

// WaitStart waits for the child workflow to start
func (r *SomeWorkflow3ChildRun) WaitStart(ctx workflow.Context) (*workflow.Execution, error) {
	var exec workflow.Execution
	if err := r.Future.GetChildWorkflowExecution().Get(ctx, &exec); err != nil {
		return nil, err
	}
	return &exec, nil
}

// SomeSignal2 sends the corresponding signal request to the child workflow
func (r *SomeWorkflow3ChildRun) SomeSignal2(ctx workflow.Context, input *SomeSignal2Request) workflow.Future {
	return r.Future.SignalChildWorkflow(ctx, SomeSignal2SignalName, input)
}

// SomeSignal1Signal describes a SomeSignal1 signal
type SomeSignal1Signal struct {
	Channel workflow.ReceiveChannel
}

// Receive blocks until a SomeSignal1 signal is received
func (s *SomeSignal1Signal) Receive(ctx workflow.Context) bool {
	more := s.Channel.Receive(ctx, nil)
	return more
}

// ReceiveAsync checks for a SomeSignal1 signal without blocking
func (s *SomeSignal1Signal) ReceiveAsync() bool {
	return s.Channel.ReceiveAsync(nil)
}

// Select checks for a SomeSignal1 signal without blocking
func (s *SomeSignal1Signal) Select(sel workflow.Selector, fn func()) workflow.Selector {
	return sel.AddReceive(s.Channel, func(workflow.ReceiveChannel, bool) {
		s.ReceiveAsync()
		if fn != nil {
			fn()
		}
	})
}

// SomeSignal1External sends a SomeSignal1 signal to an existing workflow
func SomeSignal1External(ctx workflow.Context, workflowID string, runID string) workflow.Future {
	return workflow.SignalExternalWorkflow(ctx, workflowID, runID, SomeSignal1SignalName, nil)
}

// SomeSignal2Signal describes a SomeSignal2 signal
type SomeSignal2Signal struct {
	Channel workflow.ReceiveChannel
}

// Receive blocks until a SomeSignal2 signal is received
func (s *SomeSignal2Signal) Receive(ctx workflow.Context) (*SomeSignal2Request, bool) {
	var resp SomeSignal2Request
	more := s.Channel.Receive(ctx, &resp)
	return &resp, more
}

// ReceiveAsync checks for a SomeSignal2 signal without blocking
func (s *SomeSignal2Signal) ReceiveAsync() *SomeSignal2Request {
	var resp SomeSignal2Request
	if ok := s.Channel.ReceiveAsync(&resp); !ok {
		return nil
	}
	return &resp
}

// Select checks for a SomeSignal2 signal without blocking
func (s *SomeSignal2Signal) Select(sel workflow.Selector, fn func(*SomeSignal2Request)) workflow.Selector {
	return sel.AddReceive(s.Channel, func(workflow.ReceiveChannel, bool) {
		req := s.ReceiveAsync()
		if fn != nil {
			fn(req)
		}
	})
}

// SomeSignal2External sends a SomeSignal2 signal to an existing workflow
func SomeSignal2External(ctx workflow.Context, workflowID string, runID string, req *SomeSignal2Request) workflow.Future {
	return workflow.SignalExternalWorkflow(ctx, workflowID, runID, SomeSignal2SignalName, req)
}

// Activities describes available worker activites
type Activities interface {
	// SomeActivity1 does some activity thing.
	SomeActivity1(ctx context.Context) error
	// SomeActivity2 does some activity thing.
	SomeActivity2(ctx context.Context, req *SomeActivity2Request) error
	// SomeActivity3 does some activity thing.
	SomeActivity3(ctx context.Context, req *SomeActivity3Request) (*SomeActivity3Response, error)
}

// RegisterActivities registers activities with a worker
func RegisterActivities(r worker.Registry, activities Activities) {
	RegisterSomeActivity1Activity(r, activities.SomeActivity1)
	RegisterSomeActivity2Activity(r, activities.SomeActivity2)
	RegisterSomeActivity3Activity(r, activities.SomeActivity3)
}

// RegisterSomeActivity1Activity registers a SomeActivity1 activity
func RegisterSomeActivity1Activity(r worker.Registry, fn func(context.Context) error) {
	r.RegisterActivityWithOptions(fn, activity.RegisterOptions{
		Name: SomeActivity1ActivityName,
	})
}

// SomeActivity1Future describes a SomeActivity1 activity execution
type SomeActivity1Future struct {
	Future workflow.Future
}

// Get blocks on a SomeActivity1 execution, returning the response
func (f *SomeActivity1Future) Get(ctx workflow.Context) error {
	return f.Future.Get(ctx, nil)
}

// Select adds the SomeActivity1 completion to the selector, callback can be nil
func (f *SomeActivity1Future) Select(sel workflow.Selector, fn func(*SomeActivity1Future)) workflow.Selector {
	return sel.AddFuture(f.Future, func(workflow.Future) {
		if fn != nil {
			fn(f)
		}
	})
}

// SomeActivity1 does some activity thing.
func SomeActivity1(ctx workflow.Context, opts *workflow.ActivityOptions) *SomeActivity1Future {
	if opts == nil {
		activityOpts := workflow.GetActivityOptions(ctx)
		opts = &activityOpts
	}
	ctx = workflow.WithActivityOptions(ctx, *opts)
	return &SomeActivity1Future{Future: workflow.ExecuteActivity(ctx, SomeActivity1ActivityName)}
}

// SomeActivity1 does some activity thing.
func SomeActivity1Local(ctx workflow.Context, opts *workflow.LocalActivityOptions, fn func(context.Context) error) *SomeActivity1Future {
	if opts == nil {
		activityOpts := workflow.GetLocalActivityOptions(ctx)
		opts = &activityOpts
	}
	ctx = workflow.WithLocalActivityOptions(ctx, *opts)
	var activity any
	if fn == nil {
		activity = SomeActivity1ActivityName
	} else {
		activity = fn
	}
	return &SomeActivity1Future{Future: workflow.ExecuteLocalActivity(ctx, activity)}
}

// RegisterSomeActivity2Activity registers a SomeActivity2 activity
func RegisterSomeActivity2Activity(r worker.Registry, fn func(context.Context, *SomeActivity2Request) error) {
	r.RegisterActivityWithOptions(fn, activity.RegisterOptions{
		Name: SomeActivity2ActivityName,
	})
}

// SomeActivity2Future describes a SomeActivity2 activity execution
type SomeActivity2Future struct {
	Future workflow.Future
}

// Get blocks on a SomeActivity2 execution, returning the response
func (f *SomeActivity2Future) Get(ctx workflow.Context) error {
	return f.Future.Get(ctx, nil)
}

// Select adds the SomeActivity2 completion to the selector, callback can be nil
func (f *SomeActivity2Future) Select(sel workflow.Selector, fn func(*SomeActivity2Future)) workflow.Selector {
	return sel.AddFuture(f.Future, func(workflow.Future) {
		if fn != nil {
			fn(f)
		}
	})
}

// SomeActivity2 does some activity thing.
func SomeActivity2(ctx workflow.Context, opts *workflow.ActivityOptions, req *SomeActivity2Request) *SomeActivity2Future {
	if opts == nil {
		activityOpts := workflow.GetActivityOptions(ctx)
		opts = &activityOpts
	}
	if opts.RetryPolicy == nil {
		opts.RetryPolicy = &temporal.RetryPolicy{MaximumInterval: 30000000000}
	}
	if opts.StartToCloseTimeout == 0 {
		opts.StartToCloseTimeout = 10000000000 // 10s
	}
	ctx = workflow.WithActivityOptions(ctx, *opts)
	return &SomeActivity2Future{Future: workflow.ExecuteActivity(ctx, SomeActivity2ActivityName, req)}
}

// SomeActivity2 does some activity thing.
func SomeActivity2Local(ctx workflow.Context, opts *workflow.LocalActivityOptions, fn func(context.Context, *SomeActivity2Request) error, req *SomeActivity2Request) *SomeActivity2Future {
	if opts == nil {
		activityOpts := workflow.GetLocalActivityOptions(ctx)
		opts = &activityOpts
	}
	if opts.RetryPolicy == nil {
		opts.RetryPolicy = &temporal.RetryPolicy{MaximumInterval: 30000000000}
	}
	if opts.StartToCloseTimeout == 0 {
		opts.StartToCloseTimeout = 10000000000 // 10s
	}
	ctx = workflow.WithLocalActivityOptions(ctx, *opts)
	var activity any
	if fn == nil {
		activity = SomeActivity2ActivityName
	} else {
		activity = fn
	}
	return &SomeActivity2Future{Future: workflow.ExecuteLocalActivity(ctx, activity, req)}
}

// RegisterSomeActivity3Activity registers a SomeActivity3 activity
func RegisterSomeActivity3Activity(r worker.Registry, fn func(context.Context, *SomeActivity3Request) (*SomeActivity3Response, error)) {
	r.RegisterActivityWithOptions(fn, activity.RegisterOptions{
		Name: SomeActivity3ActivityName,
	})
}

// SomeActivity3Future describes a SomeActivity3 activity execution
type SomeActivity3Future struct {
	Future workflow.Future
}

// Get blocks on a SomeActivity3 execution, returning the response
func (f *SomeActivity3Future) Get(ctx workflow.Context) (*SomeActivity3Response, error) {
	var resp SomeActivity3Response
	if err := f.Future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Select adds the SomeActivity3 completion to the selector, callback can be nil
func (f *SomeActivity3Future) Select(sel workflow.Selector, fn func(*SomeActivity3Future)) workflow.Selector {
	return sel.AddFuture(f.Future, func(workflow.Future) {
		if fn != nil {
			fn(f)
		}
	})
}

// SomeActivity3 does some activity thing.
func SomeActivity3(ctx workflow.Context, opts *workflow.ActivityOptions, req *SomeActivity3Request) *SomeActivity3Future {
	if opts == nil {
		activityOpts := workflow.GetActivityOptions(ctx)
		opts = &activityOpts
	}
	if opts.RetryPolicy == nil {
		opts.RetryPolicy = &temporal.RetryPolicy{MaximumAttempts: int32(5)}
	}
	if opts.StartToCloseTimeout == 0 {
		opts.StartToCloseTimeout = 10000000000 // 10s
	}
	ctx = workflow.WithActivityOptions(ctx, *opts)
	return &SomeActivity3Future{Future: workflow.ExecuteActivity(ctx, SomeActivity3ActivityName, req)}
}

// SomeActivity3 does some activity thing.
func SomeActivity3Local(ctx workflow.Context, opts *workflow.LocalActivityOptions, fn func(context.Context, *SomeActivity3Request) (*SomeActivity3Response, error), req *SomeActivity3Request) *SomeActivity3Future {
	if opts == nil {
		activityOpts := workflow.GetLocalActivityOptions(ctx)
		opts = &activityOpts
	}
	if opts.RetryPolicy == nil {
		opts.RetryPolicy = &temporal.RetryPolicy{MaximumAttempts: int32(5)}
	}
	if opts.StartToCloseTimeout == 0 {
		opts.StartToCloseTimeout = 10000000000 // 10s
	}
	ctx = workflow.WithLocalActivityOptions(ctx, *opts)
	var activity any
	if fn == nil {
		activity = SomeActivity3ActivityName
	} else {
		activity = fn
	}
	return &SomeActivity3Future{Future: workflow.ExecuteLocalActivity(ctx, activity, req)}
}
