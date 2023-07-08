// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 0.8.1-next (85d262dc2134d7428579c049b2e8eb0115fce7cc)
//	go go1.20.4
//	protoc (unknown)
//
// source: example/v1/example.proto
package examplev1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	expression "github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	v2 "github.com/urfave/cli/v2"
	v1 "go.temporal.io/api/enums/v1"
	activity "go.temporal.io/sdk/activity"
	client "go.temporal.io/sdk/client"
	temporal "go.temporal.io/sdk/temporal"
	worker "go.temporal.io/sdk/worker"
	workflow "go.temporal.io/sdk/workflow"
)

// ExampleTaskQueue is the default task-queue for a Example worker
const ExampleTaskQueue = "example-v1"

// Example workflow names
const (
	CreateFooWorkflowName = "example.v1.Example.CreateFooWorkflow"
)

// Example id expressions
var (
	CreateFooIDExpression = expression.MustParseExpression("create-foo/${!name.slug()}")
)

// Example query names
const (
	GetFooProgressQueryName = "example.v1.Example.GetFooProgressQuery"
)

// Example signal names
const (
	SetFooProgressSignalName = "example.v1.Example.SetFooProgressSignal"
)

// Example activity names
const (
	NotifyActivityName = "example.v1.Example.NotifyActivity"
)

// Client describes a client for a Example worker
type Client interface {
	// CreateFoo creates a new foo operation
	CreateFoo(ctx context.Context, opts *client.StartWorkflowOptions, req *CreateFooRequest) (*CreateFooResponse, error)
	// ExecuteCreateFoo executes a CreateFoo workflow
	ExecuteCreateFoo(ctx context.Context, opts *client.StartWorkflowOptions, req *CreateFooRequest) (CreateFooRun, error)
	// GetCreateFoo retrieves a CreateFoo workflow execution
	GetCreateFoo(ctx context.Context, workflowID string, runID string) (CreateFooRun, error)
	// StartCreateFooWithSetFooProgress sends a SetFooProgress signal to a CreateFoo workflow, starting it if not present
	StartCreateFooWithSetFooProgress(ctx context.Context, opts *client.StartWorkflowOptions, req *CreateFooRequest, signal *SetFooProgressRequest) (CreateFooRun, error)
	// QueryGetFooProgress sends a GetFooProgress query to an existing workflow
	QueryGetFooProgress(ctx context.Context, workflowID string, runID string) (*GetFooProgressResponse, error)
	// SignalSetFooProgress sends a SetFooProgress signal to an existing workflow
	SignalSetFooProgress(ctx context.Context, workflowID string, runID string, signal *SetFooProgressRequest) error
}

// Compile-time check that workflowClient satisfies Client
var _ Client = &workflowClient{}

// workflowClient implements a temporal client for a Example service
type workflowClient struct {
	client client.Client
}

// NewClient initializes a new Example client
func NewClient(c client.Client) Client {
	return &workflowClient{client: c}
}

// NewClientWithOptions initializes a new Example client with the given options
func NewClientWithOptions(c client.Client, opts client.Options) (Client, error) {
	var err error
	c, err = client.NewClientFromExisting(c, opts)
	if err != nil {
		return nil, fmt.Errorf("error initializing client with options: %w", err)
	}
	return &workflowClient{client: c}, nil
}

// CreateFoo creates a new foo operation
func (c *workflowClient) CreateFoo(ctx context.Context, opts *client.StartWorkflowOptions, req *CreateFooRequest) (*CreateFooResponse, error) {
	run, err := c.ExecuteCreateFoo(ctx, opts, req)
	if err != nil {
		return nil, err
	}
	return run.Get(ctx)
}

// ExecuteCreateFoo starts a CreateFoo workflow
func (c *workflowClient) ExecuteCreateFoo(ctx context.Context, opts *client.StartWorkflowOptions, req *CreateFooRequest) (CreateFooRun, error) {
	if opts == nil {
		opts = &client.StartWorkflowOptions{}
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "example-v1"
	}
	if opts.ID == "" {
		id, err := expression.EvalExpression(CreateFooIDExpression, req.ProtoReflect())
		if err != nil {
			return nil, err
		}
		opts.ID = id
	}
	if opts.WorkflowIDReusePolicy == v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v1.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
	}
	if opts.WorkflowExecutionTimeout == 0 {
		opts.WorkflowRunTimeout = 3600000000000 // 1h0m0s
	}
	run, err := c.client.ExecuteWorkflow(ctx, *opts, CreateFooWorkflowName, req)
	if err != nil {
		return nil, err
	}
	if run == nil {
		return nil, errors.New("execute workflow returned nil run")
	}
	return &createFooRun{
		client: c,
		run:    run,
	}, nil
}

// GetCreateFoo fetches an existing CreateFoo execution
func (c *workflowClient) GetCreateFoo(ctx context.Context, workflowID string, runID string) (CreateFooRun, error) {
	return &createFooRun{
		client: c,
		run:    c.client.GetWorkflow(ctx, workflowID, runID),
	}, nil
}

// StartCreateFooWithSetFooProgress starts a CreateFoo workflow and sends a SetFooProgress signal in a transaction
func (c *workflowClient) StartCreateFooWithSetFooProgress(ctx context.Context, opts *client.StartWorkflowOptions, req *CreateFooRequest, signal *SetFooProgressRequest) (CreateFooRun, error) {
	if opts == nil {
		opts = &client.StartWorkflowOptions{}
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "example-v1"
	}
	if opts.ID == "" {
		id, err := expression.EvalExpression(CreateFooIDExpression, req.ProtoReflect())
		if err != nil {
			return nil, err
		}
		opts.ID = id
	}
	if opts.WorkflowIDReusePolicy == v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v1.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
	}
	if opts.WorkflowExecutionTimeout == 0 {
		opts.WorkflowRunTimeout = 3600000000000 // 1h0m0s
	}
	run, err := c.client.SignalWithStartWorkflow(ctx, opts.ID, SetFooProgressSignalName, signal, *opts, CreateFooWorkflowName, req)
	if run == nil || err != nil {
		return nil, err
	}
	return &createFooRun{
		client: c,
		run:    run,
	}, nil
}

// QueryGetFooProgress sends a GetFooProgress query to an existing workflow
func (c *workflowClient) QueryGetFooProgress(ctx context.Context, workflowID string, runID string) (*GetFooProgressResponse, error) {
	var resp GetFooProgressResponse
	if val, err := c.client.QueryWorkflow(ctx, workflowID, runID, GetFooProgressQueryName); err != nil {
		return nil, err
	} else if err = val.Get(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SignalSetFooProgress sends a SetFooProgress signal to an existing workflow
func (c *workflowClient) SignalSetFooProgress(ctx context.Context, workflowID string, runID string, signal *SetFooProgressRequest) error {
	return c.client.SignalWorkflow(ctx, workflowID, runID, SetFooProgressSignalName, signal)
}

// CreateFooRun describes a CreateFoo workflow run
type CreateFooRun interface {
	// ID returns the workflow ID
	ID() string
	// RunID returns the workflow instance ID
	RunID() string
	// Get blocks until the workflow is complete and returns the result
	Get(ctx context.Context) (*CreateFooResponse, error)
	// GetFooProgress runs the GetFooProgress query against the workflow
	GetFooProgress(ctx context.Context) (*GetFooProgressResponse, error)
	// SetFooProgress sends a SetFooProgress signal to the workflow
	SetFooProgress(ctx context.Context, req *SetFooProgressRequest) error
}

// createFooRun provides an internal implementation of a CreateFooRun
type createFooRun struct {
	client *workflowClient
	run    client.WorkflowRun
}

// ID returns the workflow ID
func (r *createFooRun) ID() string {
	return r.run.GetID()
}

// RunID returns the execution ID
func (r *createFooRun) RunID() string {
	return r.run.GetRunID()
}

// Get blocks until the workflow is complete, returning the result if applicable
func (r *createFooRun) Get(ctx context.Context) (*CreateFooResponse, error) {
	var resp CreateFooResponse
	if err := r.run.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFooProgress executes a GetFooProgress query against the workflow
func (r *createFooRun) GetFooProgress(ctx context.Context) (*GetFooProgressResponse, error) {
	return r.client.QueryGetFooProgress(ctx, r.ID(), "")
}

// SetFooProgress sends a SetFooProgress signal to the workflow
func (r *createFooRun) SetFooProgress(ctx context.Context, req *SetFooProgressRequest) error {
	return r.client.SignalSetFooProgress(ctx, r.ID(), "", req)
}

// Workflows provides methods for initializing new Example workflow values
type Workflows interface {
	// CreateFoo initializes a new CreateFooWorkflow value
	CreateFoo(ctx workflow.Context, input *CreateFooInput) (CreateFooWorkflow, error)
}

// RegisterWorkflows registers Example workflows with the given worker
func RegisterWorkflows(r worker.Registry, workflows Workflows) {
	RegisterCreateFooWorkflow(r, workflows.CreateFoo)
}

// RegisterCreateFooWorkflow registers a CreateFoo workflow with the given worker
func RegisterCreateFooWorkflow(r worker.Registry, wf func(workflow.Context, *CreateFooInput) (CreateFooWorkflow, error)) {
	r.RegisterWorkflowWithOptions(buildCreateFoo(wf), workflow.RegisterOptions{Name: CreateFooWorkflowName})
}

// buildCreateFoo converts a CreateFoo workflow struct into a valid workflow function
func buildCreateFoo(wf func(workflow.Context, *CreateFooInput) (CreateFooWorkflow, error)) func(workflow.Context, *CreateFooRequest) (*CreateFooResponse, error) {
	return (&createFoo{wf}).CreateFoo
}

// createFoo provides an CreateFoo method for calling the user's implementation
type createFoo struct {
	ctor func(workflow.Context, *CreateFooInput) (CreateFooWorkflow, error)
}

// CreateFoo constructs a new CreateFoo value and executes it
func (w *createFoo) CreateFoo(ctx workflow.Context, req *CreateFooRequest) (*CreateFooResponse, error) {
	input := &CreateFooInput{
		Req: req,
		SetFooProgress: &SetFooProgressSignal{
			Channel: workflow.GetSignalChannel(ctx, SetFooProgressSignalName),
		},
	}
	wf, err := w.ctor(ctx, input)
	if err != nil {
		return nil, err
	}
	if err := workflow.SetQueryHandler(ctx, GetFooProgressQueryName, wf.GetFooProgress); err != nil {
		return nil, err
	}
	return wf.Execute(ctx)
}

// CreateFooInput describes the input to a CreateFoo workflow constructor
type CreateFooInput struct {
	Req            *CreateFooRequest
	SetFooProgress *SetFooProgressSignal
}

// CreateFoo creates a new foo operation
type CreateFooWorkflow interface {
	// Execute a CreateFoo workflow
	Execute(ctx workflow.Context) (*CreateFooResponse, error)
	// GetFooProgress query handler
	GetFooProgress() (*GetFooProgressResponse, error)
}

// CreateFooChild executes a child CreateFoo workflow
func CreateFooChild(ctx workflow.Context, opts *workflow.ChildWorkflowOptions, req *CreateFooRequest) *CreateFooChildRun {
	if opts == nil {
		childOpts := workflow.GetChildWorkflowOptions(ctx)
		opts = &childOpts
	}
	if opts.TaskQueue == "" {
		opts.TaskQueue = "example-v1"
	}
	if opts.WorkflowID == "" {
		id, err := expression.EvalExpression(CreateFooIDExpression, req.ProtoReflect())
		if err != nil {
			panic(err)
		}
		opts.WorkflowID = id
	}
	if opts.WorkflowIDReusePolicy == v1.WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED {
		opts.WorkflowIDReusePolicy = v1.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
	}
	if opts.WorkflowExecutionTimeout == 0 {
		opts.WorkflowRunTimeout = 3600000000000 // 1h0m0s
	}
	ctx = workflow.WithChildOptions(ctx, *opts)
	return &CreateFooChildRun{Future: workflow.ExecuteChildWorkflow(ctx, CreateFooWorkflowName, req)}
}

// CreateFooChildRun describes a child CreateFoo workflow run
type CreateFooChildRun struct {
	Future workflow.ChildWorkflowFuture
}

// Get blocks until the workflow is completed, returning the response value
func (r *CreateFooChildRun) Get(ctx workflow.Context) (*CreateFooResponse, error) {
	var resp CreateFooResponse
	if err := r.Future.Get(ctx, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Select adds this completion to the selector. Callback can be nil.
func (r *CreateFooChildRun) Select(sel workflow.Selector, fn func(CreateFooChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future, func(workflow.Future) {
		if fn != nil {
			fn(*r)
		}
	})
}

// SelectStart adds waiting for start to the selector. Callback can be nil.
func (r *CreateFooChildRun) SelectStart(sel workflow.Selector, fn func(CreateFooChildRun)) workflow.Selector {
	return sel.AddFuture(r.Future.GetChildWorkflowExecution(), func(workflow.Future) {
		if fn != nil {
			fn(*r)
		}
	})
}

// WaitStart waits for the child workflow to start
func (r *CreateFooChildRun) WaitStart(ctx workflow.Context) (*workflow.Execution, error) {
	var exec workflow.Execution
	if err := r.Future.GetChildWorkflowExecution().Get(ctx, &exec); err != nil {
		return nil, err
	}
	return &exec, nil
}

// SetFooProgress sends the corresponding signal request to the child workflow
func (r *CreateFooChildRun) SetFooProgress(ctx workflow.Context, input *SetFooProgressRequest) workflow.Future {
	return r.Future.SignalChildWorkflow(ctx, SetFooProgressSignalName, input)
}

// SetFooProgressSignal describes a SetFooProgress signal
type SetFooProgressSignal struct {
	Channel workflow.ReceiveChannel
}

// Receive blocks until a SetFooProgress signal is received
func (s *SetFooProgressSignal) Receive(ctx workflow.Context) (*SetFooProgressRequest, bool) {
	var resp SetFooProgressRequest
	more := s.Channel.Receive(ctx, &resp)
	return &resp, more
}

// ReceiveAsync checks for a SetFooProgress signal without blocking
func (s *SetFooProgressSignal) ReceiveAsync() *SetFooProgressRequest {
	var resp SetFooProgressRequest
	if ok := s.Channel.ReceiveAsync(&resp); !ok {
		return nil
	}
	return &resp
}

// Select checks for a SetFooProgress signal without blocking
func (s *SetFooProgressSignal) Select(sel workflow.Selector, fn func(*SetFooProgressRequest)) workflow.Selector {
	return sel.AddReceive(s.Channel, func(workflow.ReceiveChannel, bool) {
		req := s.ReceiveAsync()
		if fn != nil {
			fn(req)
		}
	})
}

// SetFooProgressExternal sends a SetFooProgress signal to an existing workflow
func SetFooProgressExternal(ctx workflow.Context, workflowID string, runID string, req *SetFooProgressRequest) workflow.Future {
	return workflow.SignalExternalWorkflow(ctx, workflowID, runID, SetFooProgressSignalName, req)
}

// Activities describes available worker activites
type Activities interface {
	// Notify sends a notification
	Notify(ctx context.Context, req *NotifyRequest) error
}

// RegisterActivities registers activities with a worker
func RegisterActivities(r worker.Registry, activities Activities) {
	RegisterNotifyActivity(r, activities.Notify)
}

// RegisterNotifyActivity registers a Notify activity
func RegisterNotifyActivity(r worker.Registry, fn func(context.Context, *NotifyRequest) error) {
	r.RegisterActivityWithOptions(fn, activity.RegisterOptions{
		Name: NotifyActivityName,
	})
}

// NotifyFuture describes a Notify activity execution
type NotifyFuture struct {
	Future workflow.Future
}

// Get blocks on a Notify execution, returning the response
func (f *NotifyFuture) Get(ctx workflow.Context) error {
	return f.Future.Get(ctx, nil)
}

// Select adds the Notify completion to the selector, callback can be nil
func (f *NotifyFuture) Select(sel workflow.Selector, fn func(*NotifyFuture)) workflow.Selector {
	return sel.AddFuture(f.Future, func(workflow.Future) {
		if fn != nil {
			fn(f)
		}
	})
}

// Notify sends a notification
func Notify(ctx workflow.Context, opts *workflow.ActivityOptions, req *NotifyRequest) *NotifyFuture {
	if opts == nil {
		activityOpts := workflow.GetActivityOptions(ctx)
		opts = &activityOpts
	}
	if opts.RetryPolicy == nil {
		opts.RetryPolicy = &temporal.RetryPolicy{MaximumAttempts: int32(3)}
	}
	if opts.StartToCloseTimeout == 0 {
		opts.StartToCloseTimeout = 30000000000 // 30s
	}
	ctx = workflow.WithActivityOptions(ctx, *opts)
	return &NotifyFuture{Future: workflow.ExecuteActivity(ctx, NotifyActivityName, req)}
}

// Notify sends a notification
func NotifyLocal(ctx workflow.Context, opts *workflow.LocalActivityOptions, fn func(context.Context, *NotifyRequest) error, req *NotifyRequest) *NotifyFuture {
	if opts == nil {
		activityOpts := workflow.GetLocalActivityOptions(ctx)
		opts = &activityOpts
	}
	if opts.RetryPolicy == nil {
		opts.RetryPolicy = &temporal.RetryPolicy{MaximumAttempts: int32(3)}
	}
	if opts.StartToCloseTimeout == 0 {
		opts.StartToCloseTimeout = 30000000000 // 30s
	}
	ctx = workflow.WithLocalActivityOptions(ctx, *opts)
	var activity any
	if fn == nil {
		activity = NotifyActivityName
	} else {
		activity = fn
	}
	return &NotifyFuture{Future: workflow.ExecuteLocalActivity(ctx, activity, req)}
}

// Define CLI options type
type CLIOptions struct {
	clientForCommand func(*v2.Context) (client.Client, error)
}

// Option describes a CLI option
type CLIOption func(*CLIOptions) error

// WithClientForCommand injects a client factory for use by individual commands
func WithClientForCommand(ctor func(*v2.Context) (client.Client, error)) CLIOption {
	return func(opts *CLIOptions) error {
		if ctor == nil {
			return errors.New("ClientForCommand cannot be nil")
		}
		opts.clientForCommand = ctor
		return nil
	}
}

// NewCommands contains cli commands for a Example service
func NewCommands(options ...CLIOption) ([]*v2.Command, error) {
	opts := &CLIOptions{}
	for _, opt := range options {
		if err := opt(opts); err != nil {
			return nil, err
		}
	}
	if opts.clientForCommand == nil {
		return nil, fmt.Errorf("missing required ClientForCommand")
	}
	commands := []*v2.Command{
		// GetFooProgress returns the status of a CreateFoo operation,
		{
			Name:                   "get-foo-progress",
			Usage:                  "GetFooProgress returns the status of a CreateFoo operation",
			Category:               "QUERIES",
			UseShortOptionHandling: true,
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
			},
			Action: func(cmd *v2.Context) error {
				c, err := opts.clientForCommand(cmd)
				if err != nil {
					return fmt.Errorf("error initializing client for command: %w", err)
				}
				defer c.Close()
				client := NewClient(c)
				resp, err := client.QueryGetFooProgress(cmd.Context, cmd.String("workflow-id"), cmd.String("run-id"))
				if err != nil {
					return err
				}
				b, err := json.MarshalIndent(resp, "", "  ")
				if err != nil {
					return fmt.Errorf("error formatting response for display: %w", err)
				}
				fmt.Println(string(b))
				return nil
			},
		},
		// SetFooProgress sets the current status of a CreateFoo operation,
		{
			Name:                   "set-foo-progress",
			Usage:                  "SetFooProgress sets the current status of a CreateFoo operation",
			Category:               "SIGNALS",
			UseShortOptionHandling: true,
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
				&v2.Float64Flag{
					Name:  "progress",
					Usage: "value of current workflow progress",
				},
			},
			Action: func(cmd *v2.Context) error {
				c, err := opts.clientForCommand(cmd)
				if err != nil {
					return fmt.Errorf("error initializing client for command: %w", err)
				}
				defer c.Close()
				client := NewClient(c)
				req, err := unmarshalCliFlagsToSetFooProgressRequest(cmd)
				if err != nil {
					return fmt.Errorf("error unmarshalling request: %w", err)
				}
				if err := client.SignalSetFooProgress(cmd.Context, cmd.String("workflow-id"), cmd.String("run-id"), req); err != nil {
					return fmt.Errorf("error sending \"example.v1.Example.SetFooProgress\" signal: %w", err)
				}
				fmt.Println("success")
				return nil
			},
		},
		// CreateFoo creates a new foo operation,
		{
			Name:                   "create-foo",
			Usage:                  "CreateFoo creates a new foo operation",
			Category:               "WORKFLOWS",
			UseShortOptionHandling: true,
			Flags: []v2.Flag{
				&v2.BoolFlag{
					Name:    "detach",
					Usage:   "run workflow in the background and print workflow and execution id",
					Aliases: []string{"d"},
				},
				&v2.StringFlag{
					Name:  "name",
					Usage: "unique foo name",
				},
			},
			Action: func(cmd *v2.Context) error {
				c, err := opts.clientForCommand(cmd)
				if err != nil {
					return fmt.Errorf("error initializing client for command: %w", err)
				}
				defer c.Close()
				client := NewClient(c)
				req, err := unmarshalCliFlagsToCreateFooRequest(cmd)
				if err != nil {
					return fmt.Errorf("error unmarshalling request: %w", err)
				}
				run, err := client.ExecuteCreateFoo(cmd.Context, nil, req)
				if err != nil {
					return fmt.Errorf("error starting CreateFoo workflow: %w", err)
				}
				if cmd.Bool("detach") {
					fmt.Println("success")
					fmt.Printf("workflow_id: %s\n", run.ID())
					fmt.Printf("execution_id: %s\n", run.RunID())
					return nil
				}
				if resp, err := run.Get(cmd.Context); err != nil {
					return err
				} else {
					b, err := json.MarshalIndent(resp, "", "  ")
					if err != nil {
						return fmt.Errorf("error formatting response for display: %w", err)
					}
					fmt.Println(string(b))
					return nil
				}
			},
		},
		// sends a SetFooProgress signal to a CreateFoo worfklow, starting it if it doesn't exist,
		{
			Name:                   "create-foo-with-set-foo-progress",
			Usage:                  "sends a SetFooProgress signal to a CreateFoo worfklow, starting it if it doesn't exist",
			Category:               "WORKFLOWS",
			UseShortOptionHandling: true,
			Flags: []v2.Flag{
				&v2.BoolFlag{
					Name:    "detach",
					Usage:   "run workflow in the background and print workflow and execution id",
					Aliases: []string{"d"},
				},
				&v2.StringFlag{
					Name:     "name",
					Usage:    "unique foo name",
					Category: "request",
				},
				&v2.Float64Flag{
					Name:     "progress",
					Usage:    "value of current workflow progress",
					Category: "signal",
				},
			},
			Action: func(cmd *v2.Context) error {
				c, err := opts.clientForCommand(cmd)
				if err != nil {
					return fmt.Errorf("error initializing client for command: %w", err)
				}
				defer c.Close()
				client := NewClient(c)
				req, err := unmarshalCliFlagsToCreateFooRequest(cmd)
				if err != nil {
					return fmt.Errorf("error unmarshalling request: %w", err)
				}
				signal, err := unmarshalCliFlagsToSetFooProgressRequest(cmd)
				if err != nil {
					return fmt.Errorf("error unmarshalling signal: %w", err)
				}
				run, err := client.StartCreateFooWithSetFooProgress(cmd.Context, nil, req, signal)
				if err != nil {
					return fmt.Errorf("error signaling CreateFoo workflow: %w", err)
				}
				if cmd.Bool("detach") {
					fmt.Println("success")
					fmt.Printf("workflow_id: %s\n", run.ID())
					fmt.Printf("execution_id: %s\n", run.RunID())
					return nil
				}
				if resp, err := run.Get(cmd.Context); err != nil {
					return err
				} else {
					b, err := json.MarshalIndent(resp, "", "  ")
					if err != nil {
						return fmt.Errorf("error formatting response for display: %w", err)
					}
					fmt.Println(string(b))
					return nil
				}
			},
		},
	}
	return commands, nil
}

// unmarshalCliFlagsToSetFooProgressRequest unmarshals a SetFooProgressRequest from command line flags
func unmarshalCliFlagsToSetFooProgressRequest(cmd *v2.Context) (*SetFooProgressRequest, error) {
	var result SetFooProgressRequest
	var hasValues bool
	if cmd.IsSet("progress") {
		hasValues = true
		result.Progress = float32(cmd.Float64("progress"))
	}
	if !hasValues {
		return nil, nil
	}
	return &result, nil
}

// unmarshalCliFlagsToCreateFooRequest unmarshals a CreateFooRequest from command line flags
func unmarshalCliFlagsToCreateFooRequest(cmd *v2.Context) (*CreateFooRequest, error) {
	var result CreateFooRequest
	var hasValues bool
	if cmd.IsSet("name") {
		hasValues = true
		result.Name = cmd.String("name")
	}
	if !hasValues {
		return nil, nil
	}
	return &result, nil
}
