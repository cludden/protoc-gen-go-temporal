// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 1.12.1-next (96aa1f87eb6a65319940f8e96aa513502b4c3d4d)
//	go go1.22.1
//	protoc (unknown)
//
// source: example/searchattributes/v1/searchattributes.proto
package searchattributesv1

import (
	"context"
	"errors"
	"fmt"
	expression "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	helpers "github.com/cludden/protoc-gen-go-temporal/pkg/helpers"
	scheme "github.com/cludden/protoc-gen-go-temporal/pkg/scheme"
	gohomedir "github.com/mitchellh/go-homedir"
	v2 "github.com/urfave/cli/v2"
	v1 "go.temporal.io/api/enums/v1"
	client "go.temporal.io/sdk/client"
	temporal "go.temporal.io/sdk/temporal"
	testsuite "go.temporal.io/sdk/testsuite"
	worker "go.temporal.io/sdk/worker"
	workflow "go.temporal.io/sdk/workflow"
	protojson "google.golang.org/protobuf/encoding/protojson"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"os"
	"sort"
	"time"
)

// ExampleTaskQueue is the default task-queue for a example.searchattributes.v1.Example worker
const ExampleTaskQueue = "searchattributes"

// example.searchattributes.v1.Example workflow names
const (
	SearchAttributesWorkflowName = "example.searchattributes.v1.Example.SearchAttributes"
)

// example.searchattributes.v1.Example workflow id expressions
var (
	SearchAttributesIdexpression = expression.MustParseExpression("search_attributes_${! uuid_v4() }")
)

// example.searchattributes.v1.Example workflow search attribute mappings
var (
	SearchAttributesSearchAttributesMapping = expression.MustParseMapping("CustomKeywordField = customKeywordField \nCustomTextField = customTextField \nCustomIntField = customIntField.int64() \nCustomDoubleField = customDoubleField \nCustomBoolField = customBoolField \nCustomDatetimeField = customDatetimeField.ts_parse(\"2006-01-02T15:04:05Z\") \n")
)

// ExampleClient describes a client for a(n) example.searchattributes.v1.Example worker
type ExampleClient interface {
	// SearchAttributes executes a(n) example.searchattributes.v1.Example.SearchAttributes workflow and blocks until error or response received
	SearchAttributes(ctx context.Context, req *SearchAttributesInput, opts ...*SearchAttributesOptions) error

	// SearchAttributesAsync starts a(n) example.searchattributes.v1.Example.SearchAttributes workflow and returns a handle to the workflow run
	SearchAttributesAsync(ctx context.Context, req *SearchAttributesInput, opts ...*SearchAttributesOptions) (SearchAttributesRun, error)

	// GetSearchAttributes retrieves a handle to an existing example.searchattributes.v1.Example.SearchAttributes workflow execution
	GetSearchAttributes(ctx context.Context, workflowID string, runID string) SearchAttributesRun

	// CancelWorkflow requests cancellation of an existing workflow execution
	CancelWorkflow(ctx context.Context, workflowID string, runID string) error

	// TerminateWorkflow an existing workflow execution
	TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error
}

// exampleClient implements a temporal client for a example.searchattributes.v1.Example service
type exampleClient struct {
	client client.Client
	log    *slog.Logger
}

// NewExampleClient initializes a new example.searchattributes.v1.Example client
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

// example.searchattributes.v1.Example.SearchAttributes executes a example.searchattributes.v1.Example.SearchAttributes workflow and blocks until error or response received
func (c *exampleClient) SearchAttributes(ctx context.Context, req *SearchAttributesInput, options ...*SearchAttributesOptions) error {
	run, err := c.SearchAttributesAsync(ctx, req, options...)
	if err != nil {
		return err
	}
	return run.Get(ctx)
}

// SearchAttributesAsync starts a(n) example.searchattributes.v1.Example.SearchAttributes workflow and returns a handle to the workflow run
func (c *exampleClient) SearchAttributesAsync(ctx context.Context, req *SearchAttributesInput, options ...*SearchAttributesOptions) (SearchAttributesRun, error) {
	var o *SearchAttributesOptions
	if len(options) > 0 && options[0] != nil {
		o = options[0]
	} else {
		o = NewSearchAttributesOptions()
	}
	opts, err := o.Build(req.ProtoReflect())
	if err != nil {
		return nil, fmt.Errorf("error initializing client.StartWorkflowOptions: %w", err)
	}
	run, err := c.client.ExecuteWorkflow(ctx, opts, SearchAttributesWorkflowName, req)
	if err != nil {
		return nil, err
	}
	if run == nil {
		return nil, errors.New("execute workflow returned nil run")
	}
	return &searchAttributesRun{
		client: c,
		run:    run,
	}, nil
}

// GetSearchAttributes fetches an existing example.searchattributes.v1.Example.SearchAttributes execution
func (c *exampleClient) GetSearchAttributes(ctx context.Context, workflowID string, runID string) SearchAttributesRun {
	return &searchAttributesRun{
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

// SearchAttributesOptions provides configuration for a example.searchattributes.v1.Example.SearchAttributes workflow operation
type SearchAttributesOptions struct {
	options          client.StartWorkflowOptions
	executionTimeout *time.Duration
	id               *string
	idReusePolicy    v1.WorkflowIdReusePolicy
	retryPolicy      *temporal.RetryPolicy
	runTimeout       *time.Duration
	searchAttributes map[string]any
	taskQueue        *string
	taskTimeout      *time.Duration
}

// NewSearchAttributesOptions initializes a new SearchAttributesOptions value
func NewSearchAttributesOptions() *SearchAttributesOptions {
	return &SearchAttributesOptions{}
}

// Build initializes a new go.temporal.io/sdk/client.StartWorkflowOptions value with defaults and overrides applied
func (o *SearchAttributesOptions) Build(req protoreflect.Message) (client.StartWorkflowOptions, error) {
	opts := o.options
	if v := o.id; v != nil {
		opts.ID = *v
	} else if opts.ID == "" {
		id, err := expression.EvalExpression(SearchAttributesIdexpression, req)
		if err != nil {
			return opts, fmt.Errorf("error evaluating id expression for %q workflow: %w", SearchAttributesWorkflowName, err)
		}
		opts.ID = id
	}
	if v := o.idReusePolicy; v != v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
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
	} else if opts.SearchAttributes == nil {
		structured, err := expression.ToStructured(req)
		if err != nil {
			return opts, fmt.Errorf("error serializing input for \"SearchAttributes\" search attribute mapping: %v", err)
		}
		result, err := SearchAttributesSearchAttributesMapping.Query(structured)
		if err != nil {
			return opts, fmt.Errorf("error executing \"SearchAttributes\" search attribute mapping: %v", err)
		}
		searchAttributes, ok := result.(map[string]any)
		if !ok {
			return opts, fmt.Errorf("expected \"SearchAttributes\" search attribute mapping to return map[string]any, got: %T", result)
		}
		opts.SearchAttributes = searchAttributes
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
func (o *SearchAttributesOptions) WithStartWorkflowOptions(options client.StartWorkflowOptions) *SearchAttributesOptions {
	o.options = options
	return o
}

// WithExecutionTimeout sets the WorkflowExecutionTimeout value
func (o *SearchAttributesOptions) WithExecutionTimeout(d time.Duration) *SearchAttributesOptions {
	o.executionTimeout = &d
	return o
}

// WithID sets the ID value
func (o *SearchAttributesOptions) WithID(id string) *SearchAttributesOptions {
	o.id = &id
	return o
}

// WithIDReusePolicy sets the WorkflowIDReusePolicy value
func (o *SearchAttributesOptions) WithIDReusePolicy(policy v1.WorkflowIdReusePolicy) *SearchAttributesOptions {
	o.idReusePolicy = policy
	return o
}

// WithRetryPolicy sets the RetryPolicy value
func (o *SearchAttributesOptions) WithRetryPolicy(policy *temporal.RetryPolicy) *SearchAttributesOptions {
	o.retryPolicy = policy
	return o
}

// WithRunTimeout sets the WorkflowRunTimeout value
func (o *SearchAttributesOptions) WithRunTimeout(d time.Duration) *SearchAttributesOptions {
	o.runTimeout = &d
	return o
}

// WithSearchAttributes sets the SearchAttributes value
func (o *SearchAttributesOptions) WithSearchAttributes(sa map[string]any) *SearchAttributesOptions {
	o.searchAttributes = sa
	return o
}

// WithTaskTimeout sets the WorkflowTaskTimeout value
func (o *SearchAttributesOptions) WithTaskTimeout(d time.Duration) *SearchAttributesOptions {
	o.taskTimeout = &d
	return o
}

// WithTaskQueue sets the TaskQueue value
func (o *SearchAttributesOptions) WithTaskQueue(tq string) *SearchAttributesOptions {
	o.taskQueue = &tq
	return o
}

// SearchAttributesRun describes a(n) example.searchattributes.v1.Example.SearchAttributes workflow run
type SearchAttributesRun interface {
	// ID returns the workflow ID
	ID() string

	// RunID returns the workflow instance ID
	RunID() string

	// Run returns the inner client.WorkflowRun
	Run() client.WorkflowRun

	// Get blocks until the workflow is complete and returns the result
	Get(ctx context.Context) error

	// Cancel requests cancellation of a workflow in execution, returning an error if applicable
	Cancel(ctx context.Context) error

	// Terminate terminates a workflow in execution, returning an error if applicable
	Terminate(ctx context.Context, reason string, details ...interface{}) error
}

// searchAttributesRun provides an internal implementation of a(n) SearchAttributesRunRun
type searchAttributesRun struct {
	client *exampleClient
	run    client.WorkflowRun
}

// ID returns the workflow ID
func (r *searchAttributesRun) ID() string {
	return r.run.GetID()
}

// Run returns the inner client.WorkflowRun
func (r *searchAttributesRun) Run() client.WorkflowRun {
	return r.run
}

// RunID returns the execution ID
func (r *searchAttributesRun) RunID() string {
	return r.run.GetRunID()
}

// Cancel requests cancellation of a workflow in execution, returning an error if applicable
func (r *searchAttributesRun) Cancel(ctx context.Context) error {
	return r.client.CancelWorkflow(ctx, r.ID(), r.RunID())
}

// Get blocks until the workflow is complete, returning the result if applicable
func (r *searchAttributesRun) Get(ctx context.Context) error {
	return r.run.Get(ctx, nil)
}

// Terminate terminates a workflow in execution, returning an error if applicable
func (r *searchAttributesRun) Terminate(ctx context.Context, reason string, details ...interface{}) error {
	return r.client.TerminateWorkflow(ctx, r.ID(), r.RunID(), reason, details...)
}

// Reference to generated workflow functions
var (
	// SearchAttributesFunction implements a "example.searchattributes.v1.Example.SearchAttributes" workflow
	SearchAttributesFunction func(workflow.Context, *SearchAttributesInput) error
)

// ExampleWorkflowFunctions describes a mockable dependency for inlining workflows within other workflows
type (
	// ExampleWorkflowFunctions describes a mockable dependency for inlining workflows within other workflows
	ExampleWorkflowFunctions interface {
		// SearchAttributes executes a "example.searchattributes.v1.Example.SearchAttributes" workflow inline
		SearchAttributes(workflow.Context, *SearchAttributesInput) error
	}
	// exampleWorkflowFunctions provides an internal ExampleWorkflowFunctions implementation
	exampleWorkflowFunctions struct{}
)

func NewExampleWorkflowFunctions() ExampleWorkflowFunctions {
	return &exampleWorkflowFunctions{}
}

// SearchAttributes executes a "example.searchattributes.v1.Example.SearchAttributes" workflow inline
func (f *exampleWorkflowFunctions) SearchAttributes(ctx workflow.Context, req *SearchAttributesInput) error {
	if SearchAttributesFunction == nil {
		return errors.New("SearchAttributes requires workflow registration via RegisterExampleWorkflows or RegisterSearchAttributesWorkflow")
	}
	return SearchAttributesFunction(ctx, req)
}

// ExampleWorkflows provides methods for initializing new example.searchattributes.v1.Example workflow values
type ExampleWorkflows interface {
	// SearchAttributes initializes a new a(n) SearchAttributesWorkflow implementation
	SearchAttributes(ctx workflow.Context, input *SearchAttributesWorkflowInput) (SearchAttributesWorkflow, error)
}

// RegisterExampleWorkflows registers example.searchattributes.v1.Example workflows with the given worker
func RegisterExampleWorkflows(r worker.WorkflowRegistry, workflows ExampleWorkflows) {
	RegisterSearchAttributesWorkflow(r, workflows.SearchAttributes)
}

// RegisterSearchAttributesWorkflow registers a example.searchattributes.v1.Example.SearchAttributes workflow with the given worker
func RegisterSearchAttributesWorkflow(r worker.WorkflowRegistry, wf func(workflow.Context, *SearchAttributesWorkflowInput) (SearchAttributesWorkflow, error)) {
	SearchAttributesFunction = buildSearchAttributes(wf)
	r.RegisterWorkflowWithOptions(SearchAttributesFunction, workflow.RegisterOptions{Name: SearchAttributesWorkflowName})
}

// buildSearchAttributes converts a SearchAttributes workflow struct into a valid workflow function
func buildSearchAttributes(ctor func(workflow.Context, *SearchAttributesWorkflowInput) (SearchAttributesWorkflow, error)) func(workflow.Context, *SearchAttributesInput) error {
	return func(ctx workflow.Context, req *SearchAttributesInput) error {
		input := &SearchAttributesWorkflowInput{
			Req: req,
		}
		wf, err := ctor(ctx, input)
		if err != nil {
			return err
		}
		if initializable, ok := wf.(helpers.Initializable); ok {
			if err := initializable.Initialize(ctx); err != nil {
				return err
			}
		}
		return wf.Execute(ctx)
	}
}

// SearchAttributesWorkflowInput describes the input to a(n) example.searchattributes.v1.Example.SearchAttributes workflow constructor
type SearchAttributesWorkflowInput struct {
	Req *SearchAttributesInput
}

// SearchAttributesWorkflow describes a(n) example.searchattributes.v1.Example.SearchAttributes workflow implementation
//
// workflow details: (id: "search_attributes_${! uuid_v4() }")
type SearchAttributesWorkflow interface {
	// Execute defines the entrypoint to a(n) example.searchattributes.v1.Example.SearchAttributes workflow
	Execute(ctx workflow.Context) error
}

// SearchAttributesChild executes a child example.searchattributes.v1.Example.SearchAttributes workflow and blocks until error or response received
func SearchAttributesChild(ctx workflow.Context, req *SearchAttributesInput, options ...*SearchAttributesChildOptions) error {
	childRun, err := SearchAttributesChildAsync(ctx, req, options...)
	if err != nil {
		return err
	}
	return childRun.Get(ctx)
}

// SearchAttributesChildAsync starts a child example.searchattributes.v1.Example.SearchAttributes workflow and returns a handle to the child workflow run
func SearchAttributesChildAsync(ctx workflow.Context, req *SearchAttributesInput, options ...*SearchAttributesChildOptions) (*SearchAttributesChildRun, error) {
	var o *SearchAttributesChildOptions
	if len(options) > 0 && options[0] != nil {
		o = options[0]
	} else {
		o = NewSearchAttributesChildOptions()
	}
	opts, err := o.Build(ctx, req.ProtoReflect())
	if err != nil {
		return nil, fmt.Errorf("error initializing workflow.ChildWorkflowOptions: %w", err)
	}
	ctx = workflow.WithChildOptions(ctx, opts)
	return &SearchAttributesChildRun{Future: workflow.ExecuteChildWorkflow(ctx, SearchAttributesWorkflowName, req)}, nil
}

// SearchAttributesChildOptions provides configuration for a child example.searchattributes.v1.Example.SearchAttributes workflow operation
type SearchAttributesChildOptions struct {
	options             workflow.ChildWorkflowOptions
	executionTimeout    *time.Duration
	id                  *string
	idReusePolicy       v1.WorkflowIdReusePolicy
	retryPolicy         *temporal.RetryPolicy
	runTimeout          *time.Duration
	searchAttributes    map[string]any
	taskQueue           *string
	taskTimeout         *time.Duration
	parentClosePolicy   v1.ParentClosePolicy
	waitForCancellation *bool
}

// NewSearchAttributesChildOptions initializes a new SearchAttributesChildOptions value
func NewSearchAttributesChildOptions() *SearchAttributesChildOptions {
	return &SearchAttributesChildOptions{}
}

// Build initializes a new go.temporal.io/sdk/workflow.ChildWorkflowOptions value with defaults and overrides applied
func (o *SearchAttributesChildOptions) Build(ctx workflow.Context, req protoreflect.Message) (workflow.ChildWorkflowOptions, error) {
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
				id, err := expression.EvalExpression(SearchAttributesIdexpression, req)
				if err != nil {
					return "", fmt.Errorf("error evaluating id expression for %q workflow: %w", SearchAttributesWorkflowName, err)
				}
				return id, nil
			}).Get(ctx, &opts.WorkflowID); err != nil {
				return opts, fmt.Errorf("error evaluating id expression for %q workflow: %w", SearchAttributesWorkflowName, err)
			}
		} else {
			id, err := expression.EvalExpression(SearchAttributesIdexpression, req)
			if err != nil {
				return opts, fmt.Errorf("error evaluating id expression for %q workflow: %w", SearchAttributesWorkflowName, err)
			}
			opts.WorkflowID = id
		}
	}
	if v := o.idReusePolicy; v != v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
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
	} else if opts.SearchAttributes == nil {
		// wrap expression evaluation in local activity
		// more info: https://cludden.github.io/protoc-gen-go-temporal/docs/guides/patches#pv_64-expression-evaluation-local-activity
		if workflow.GetVersion(ctx, "cludden_protoc-gen-go-temporal_64_expression-evaluation-local-activity", workflow.DefaultVersion, 1) == 1 {
			lao := workflow.GetLocalActivityOptions(ctx)
			lao.ScheduleToCloseTimeout = time.Second * 10
			if err := workflow.ExecuteLocalActivity(workflow.WithLocalActivityOptions(ctx, lao), func(ctx context.Context) (map[string]any, error) {
				structured, err := expression.ToStructured(req)
				if err != nil {
					return nil, fmt.Errorf("error serializing input for \"SearchAttributes\" search attribute mapping: %v", err)
				}
				result, err := SearchAttributesSearchAttributesMapping.Query(structured)
				if err != nil {
					return nil, fmt.Errorf("error executing \"SearchAttributes\" search attribute mapping: %v", err)
				}
				searchAttributes, ok := result.(map[string]any)
				if !ok {
					return nil, fmt.Errorf("expected \"SearchAttributes\" search attribute mapping to return map[string]any, got: %T", result)
				}
				return searchAttributes, nil
			}).Get(ctx, &opts.SearchAttributes); err != nil {
				return opts, fmt.Errorf("error evaluating search attributes for %q workflow: %w", SearchAttributesWorkflowName, err)
			}
		} else {
			structured, err := expression.ToStructured(req)
			if err != nil {
				return opts, fmt.Errorf("error serializing input for \"SearchAttributes\" search attribute mapping: %v", err)
			}
			result, err := SearchAttributesSearchAttributesMapping.Query(structured)
			if err != nil {
				return opts, fmt.Errorf("error executing \"SearchAttributes\" search attribute mapping: %v", err)
			}
			searchAttributes, ok := result.(map[string]any)
			if !ok {
				return opts, fmt.Errorf("expected \"SearchAttributes\" search attribute mapping to return map[string]any, got: %T", result)
			}
			opts.SearchAttributes = searchAttributes
		}
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
	if v := o.parentClosePolicy; v != v1.PARENT_CLOSE_POLICY_UNSPECIFIED {
		opts.ParentClosePolicy = v
	}
	if v := o.waitForCancellation; v != nil {
		opts.WaitForCancellation = *v
	}
	return opts, nil
}

// WithChildWorkflowOptions sets the initial go.temporal.io/sdk/workflow.ChildWorkflowOptions
func (o *SearchAttributesChildOptions) WithChildWorkflowOptions(options workflow.ChildWorkflowOptions) *SearchAttributesChildOptions {
	o.options = options
	return o
}

// WithExecutionTimeout sets the WorkflowExecutionTimeout value
func (o *SearchAttributesChildOptions) WithExecutionTimeout(d time.Duration) *SearchAttributesChildOptions {
	o.executionTimeout = &d
	return o
}

// WithID sets the WorkflowID value
func (o *SearchAttributesChildOptions) WithID(id string) *SearchAttributesChildOptions {
	o.id = &id
	return o
}

// WithIDReusePolicy sets the WorkflowIDReusePolicy value
func (o *SearchAttributesChildOptions) WithIDReusePolicy(policy v1.WorkflowIdReusePolicy) *SearchAttributesChildOptions {
	o.idReusePolicy = policy
	return o
}

// WithParentClosePolicy sets the WorkflowIDReusePolicy value
func (o *SearchAttributesChildOptions) WithParentClosePolicy(policy v1.ParentClosePolicy) *SearchAttributesChildOptions {
	o.parentClosePolicy = policy
	return o
}

// WithRetryPolicy sets the RetryPolicy value
func (o *SearchAttributesChildOptions) WithRetryPolicy(policy *temporal.RetryPolicy) *SearchAttributesChildOptions {
	o.retryPolicy = policy
	return o
}

// WithRunTimeout sets the WorkflowRunTimeout value
func (o *SearchAttributesChildOptions) WithRunTimeout(d time.Duration) *SearchAttributesChildOptions {
	o.runTimeout = &d
	return o
}

// WithSearchAttributes sets the SearchAttributes value
func (o *SearchAttributesChildOptions) WithSearchAttributes(sa map[string]any) *SearchAttributesChildOptions {
	o.searchAttributes = sa
	return o
}

// WithTaskTimeout sets the WorkflowTaskTimeout value
func (o *SearchAttributesChildOptions) WithTaskTimeout(d time.Duration) *SearchAttributesChildOptions {
	o.taskTimeout = &d
	return o
}

// WithTaskQueue sets the TaskQueue value
func (o *SearchAttributesChildOptions) WithTaskQueue(tq string) *SearchAttributesChildOptions {
	o.taskQueue = &tq
	return o
}

// WithWaitForCancellation sets the WaitForCancellation value
func (o *SearchAttributesChildOptions) WithWaitForCancellation(wait bool) *SearchAttributesChildOptions {
	o.waitForCancellation = &wait
	return o
}

// SearchAttributesChildRun describes a child SearchAttributes workflow run
type SearchAttributesChildRun struct {
	Future workflow.ChildWorkflowFuture
}

// Get blocks until the workflow is completed, returning the response value
func (r *SearchAttributesChildRun) Get(ctx workflow.Context) error {
	if err := r.Future.Get(ctx, nil); err != nil {
		return err
	}
	return nil
}

// Select adds this completion to the selector. Callback can be nil.
func (r *SearchAttributesChildRun) Select(sel workflow.Selector, fn func(*SearchAttributesChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future, func(workflow.Future) {
		if fn != nil {
			fn(r)
		}
	})
}

// SelectStart adds waiting for start to the selector. Callback can be nil.
func (r *SearchAttributesChildRun) SelectStart(sel workflow.Selector, fn func(*SearchAttributesChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future.GetChildWorkflowExecution(), func(workflow.Future) {
		if fn != nil {
			fn(r)
		}
	})
}

// WaitStart waits for the child workflow to start
func (r *SearchAttributesChildRun) WaitStart(ctx workflow.Context) (*workflow.Execution, error) {
	var exec workflow.Execution
	if err := r.Future.GetChildWorkflowExecution().Get(ctx, &exec); err != nil {
		return nil, err
	}
	return &exec, nil
}

// ExampleActivities describes available worker activities
type ExampleActivities interface{}

// RegisterExampleActivities registers activities with a worker
func RegisterExampleActivities(r worker.ActivityRegistry, activities ExampleActivities) {}

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

// SearchAttributes executes a(n) example.searchattributes.v1.Example.SearchAttributes workflow in the test environment
func (c *TestExampleClient) SearchAttributes(ctx context.Context, req *SearchAttributesInput, opts ...*SearchAttributesOptions) error {
	run, err := c.SearchAttributesAsync(ctx, req, opts...)
	if err != nil {
		return err
	}
	return run.Get(ctx)
}

// SearchAttributesAsync executes a(n) example.searchattributes.v1.Example.SearchAttributes workflow in the test environment
func (c *TestExampleClient) SearchAttributesAsync(ctx context.Context, req *SearchAttributesInput, options ...*SearchAttributesOptions) (SearchAttributesRun, error) {
	var o *SearchAttributesOptions
	if len(options) > 0 && options[0] != nil {
		o = options[0]
	} else {
		o = NewSearchAttributesOptions()
	}
	opts, err := o.Build(req.ProtoReflect())
	if err != nil {
		return nil, fmt.Errorf("error initializing client.StartWorkflowOptions: %w", err)
	}
	return &testSearchAttributesRun{client: c, env: c.env, opts: &opts, req: req, workflows: c.workflows}, nil
}

// GetSearchAttributes is a noop
func (c *TestExampleClient) GetSearchAttributes(ctx context.Context, workflowID string, runID string) SearchAttributesRun {
	return &testSearchAttributesRun{env: c.env, workflows: c.workflows}
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

var _ SearchAttributesRun = &testSearchAttributesRun{}

// testSearchAttributesRun provides convenience methods for interacting with a(n) example.searchattributes.v1.Example.SearchAttributes workflow in the test environment
type testSearchAttributesRun struct {
	client    *TestExampleClient
	env       *testsuite.TestWorkflowEnvironment
	opts      *client.StartWorkflowOptions
	req       *SearchAttributesInput
	workflows ExampleWorkflows
}

// Cancel requests cancellation of a workflow in execution, returning an error if applicable
func (r *testSearchAttributesRun) Cancel(ctx context.Context) error {
	return r.client.CancelWorkflow(ctx, r.ID(), r.RunID())
}

// Get retrieves a test example.searchattributes.v1.Example.SearchAttributes workflow result
func (r *testSearchAttributesRun) Get(context.Context) error {
	r.env.ExecuteWorkflow(SearchAttributesWorkflowName, r.req)
	if !r.env.IsWorkflowCompleted() {
		return errors.New("workflow in progress")
	}
	if err := r.env.GetWorkflowError(); err != nil {
		return err
	}
	return nil
}

// ID returns a test example.searchattributes.v1.Example.SearchAttributes workflow run's workflow ID
func (r *testSearchAttributesRun) ID() string {
	if r.opts != nil {
		return r.opts.ID
	}
	return ""
}

// Run noop implementation
func (r *testSearchAttributesRun) Run() client.WorkflowRun {
	return nil
}

// RunID noop implementation
func (r *testSearchAttributesRun) RunID() string {
	return ""
}

// Terminate terminates a workflow in execution, returning an error if applicable
func (r *testSearchAttributesRun) Terminate(ctx context.Context, reason string, details ...interface{}) error {
	return r.client.TerminateWorkflow(ctx, r.ID(), r.RunID(), reason, details...)
}

// ExampleCliOptions describes runtime configuration for example.searchattributes.v1.Example cli
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

// NewExampleCli initializes a cli for a(n) example.searchattributes.v1.Example service
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

// NewExampleCliCommand initializes a cli command for a example.searchattributes.v1.Example service with subcommands for each query, signal, update, and workflow
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

// newExampleCommands initializes (sub)commands for a example.searchattributes.v1.Example cli or command
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
			Name:                   "search-attributes",
			Usage:                  "executes a(n) example.searchattributes.v1.Example.SearchAttributes workflow",
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
					Value:   "searchattributes",
				},
				&v2.StringFlag{
					Name:    "input-file",
					Usage:   "path to json-formatted input file",
					Aliases: []string{"f"},
				},
				&v2.StringFlag{
					Name:     "custom-keyword-field",
					Usage:    "set the value of the operation's \"CustomKeywordField\" parameter",
					Category: "INPUT",
				},
				&v2.StringFlag{
					Name:     "custom-text-field",
					Usage:    "set the value of the operation's \"CustomTextField\" parameter",
					Category: "INPUT",
				},
				&v2.Int64Flag{
					Name:     "custom-int-field",
					Usage:    "set the value of the operation's \"CustomIntField\" parameter",
					Category: "INPUT",
				},
				&v2.Float64Flag{
					Name:     "custom-double-field",
					Usage:    "set the value of the operation's \"CustomDoubleField\" parameter",
					Category: "INPUT",
				},
				&v2.BoolFlag{
					Name:     "custom-bool-field",
					Usage:    "set the value of the operation's \"CustomBoolField\" parameter",
					Category: "INPUT",
				},
				&v2.StringFlag{
					Name:     "custom-datetime-field",
					Usage:    "set the value of the operation's \"CustomDatetimeField\" parameter (e.g. \"2017-01-15T01:30:15.01Z\")",
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
				req, err := UnmarshalCliFlagsToSearchAttributesInput(cmd)
				if err != nil {
					return fmt.Errorf("error unmarshalling request: %w", err)
				}
				opts := client.StartWorkflowOptions{}
				if tq := cmd.String("task-queue"); tq != "" {
					opts.TaskQueue = tq
				}
				run, err := c.SearchAttributesAsync(cmd.Context, req, NewSearchAttributesOptions().WithStartWorkflowOptions(opts))
				if err != nil {
					return fmt.Errorf("error starting %s workflow: %w", SearchAttributesWorkflowName, err)
				}
				if cmd.Bool("detach") {
					fmt.Println("success")
					fmt.Printf("workflow id: %s\n", run.ID())
					fmt.Printf("run id: %s\n", run.RunID())
					return nil
				}
				if err := run.Get(cmd.Context); err != nil {
					return err
				} else {
					return nil
				}
			},
		},
	}
	if opts.worker != nil {
		commands = append(commands, []*v2.Command{
			{
				Name:                   "worker",
				Usage:                  "runs a example.searchattributes.v1.Example worker process",
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

// UnmarshalCliFlagsToSearchAttributesInput unmarshals a SearchAttributesInput from command line flags
func UnmarshalCliFlagsToSearchAttributesInput(cmd *v2.Context) (*SearchAttributesInput, error) {
	var result SearchAttributesInput
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
	if cmd.IsSet("custom-keyword-field") {
		hasValues = true
		result.CustomKeywordField = cmd.String("custom-keyword-field")
	}
	if cmd.IsSet("custom-text-field") {
		hasValues = true
		result.CustomTextField = cmd.String("custom-text-field")
	}
	if cmd.IsSet("custom-int-field") {
		hasValues = true
		result.CustomIntField = cmd.Int64("custom-int-field")
	}
	if cmd.IsSet("custom-double-field") {
		hasValues = true
		result.CustomDoubleField = cmd.Float64("custom-double-field")
	}
	if cmd.IsSet("custom-bool-field") {
		hasValues = true
		result.CustomBoolField = cmd.Bool("custom-bool-field")
	}
	if cmd.IsSet("custom-datetime-field") {
		hasValues = true
		v, err := time.Parse(time.RFC3339Nano, cmd.String("custom-datetime-field"))
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling \"custom-datetime-field\" timestamp flag: %w", err)
		}
		result.CustomDatetimeField = timestamppb.New(v)
	}
	if !hasValues {
		return nil, nil
	}
	return &result, nil
}

// WithExampleSchemeTypes registers all Example protobuf types with the given scheme
func WithExampleSchemeTypes() scheme.Option {
	return func(s *scheme.Scheme) {
		s.RegisterType(File_example_searchattributes_v1_searchattributes_proto.Messages().ByName("SearchAttributesInput"))
	}
}
