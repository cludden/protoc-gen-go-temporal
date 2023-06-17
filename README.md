# protoc-gen-go-temporal

a protoc plugin for generating temporal clients and workers in Go from protobuf schemas

inspired by [github.com/cretz/temporal-sdk-go-advanced](https://github.com/cretz/temporal-sdk-go-advanced)

**Features:**
- define default `client.StartWorkflowOptions`, `workflow.ActivityOptions`, `workflow.ChildWorkflowOptions` including:
  - default workflow ids that can leverage inputs via [Bloblang ID expressions](#id-expressions)
  - default timeouts, retry policies, id reuse policies
- generates typed client and workflow helpers
  - generates client with methods for executing workflows, queries, signals
  - generates methods for calling activities and local activities from workflows
  - generates methods for executing child workflows and signalling external workflows

## Getting Started
1. Install [buf](https://docs.buf.build/installation)

2. Install protoc plugins
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/cludden/protoc-gen-go-temporal/cmd/protoc-gen-go_temporal@latest
```

3. Initialize buf repository
```shell
mkdir proto && cd proto && buf init
```

4. Add dependency to `buf.yaml`
```yaml
version: v1
deps:
  - buf.build/cludden/protoc-gen-go-temporal
```

5. Add plugin to `buf.gen.yaml` and exclude it from managed mode go prefix
```yaml
version: v1
managed:
  enabled: true
  go_package_prefix:
    default: github.com/foo/bar/gen
    except:
      - buf.build/cludden/protoc-gen-go-temporal
plugins:
  - plugin: go
    out: gen
    opt: paths=source_relative
  - plugin: go_temporal
    out: gen
    opt: paths=source_relative
    strategy: all
```

6. Define your service  
<small><b><i>note:</i></b> see [advanced](#advanced), [example](./example/), and [test](./test/) for more details on generated code and usage</small>

```protobuf
syntax="proto3";

package example.v1;

import "google/protobuf/duration.proto";
import "google/protobuf/empty.proto";
import "temporal/v1/temporal.proto";

service Example {
  option (temporal.v1.service) = {
    task_queue: "example-v1"
  };

  // HelloWorld defines a workflow with a single activity of the same name
  rpc HelloWorld(HelloWorldRequest) returns (HelloWorldResponse) {
    option (temporal.v1.workflow) = {
      default_options {
        id: 'hello-world/${! uuid_v4() }'
        execution_timeout: { seconds: 30 }
      }
    };
    option (temporal.v1.activity) = {};
  }
}

// HelloWorldRequest describes the input to a HelloWorld workflow/activity
message HelloWorldRequest {
  // Customize greeting
  // @gotags: validate:"omitempty,oneof=hello hi hey hola"
  string greeting = 1;
}

// HelloWorldResponse descibes the output from a HelloWorld workflow/activity
message HelloWorldResponse {
  string result = 1;
}
```

7. Generate temporal worker and client types, methods, interfaces, and functions
```shell
buf generate
```

8. Implement your activities, workflows

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    examplev1 "<yourproject>/gen/v1" // import generated code
    "github.com/google/uuid"
    "go.temporal.io/sdk/activity"
    "go.temporal.io/sdk/client"
    logger "go.temporal.io/sdk/log"
    "go.temporal.io/sdk/workflow"
    "go.temporal.io/sdk/worker"
    "google.golang.org/protobuf/types/known/durationpb"
)

// Define a struct to manage workflow constructors, this simplifies registration
// with lots of workflows
type Workflows struct{}

// ============================================================================

// Define a struct per workflow for managing workflow state
type HelloWorldWorkflow struct {
  *examplev1.HelloWorldInput
  log logger.Logger
}

// Define a workflow constructor
func (w *Workflows) HelloWorld(ctx workflow.Context, input *examplev1.HelloWorldInput) (examplev1.HelloWorldWorkflow, error) {
  return &HelloWorldWorkflow(input, workflow.GetLogger(ctx))
}

// Define workflow function
func (wf *HelloWorldWorkflow) Execute(ctx workflow.Context) (*examplev1.HelloWorldResponse, error) {
  // call activity of the same name from our definition above
  return examplev1.HelloWorld(ctx, nil, wf.Req)
}

// ============================================================================

// Define a struct to manage activities
type Activities struct {}

func (a *Activities) HelloWorld(ctx context.Context, req *examplev1.HelloWorldRequest) (*examplev1.HelloWorldResponse{}, error) {
  return &examplev1.HelloWorldResponse{Result: req.GetGreeting()}, nil
}

// ============================================================================

func main() {
    // initialize temporal client
    c, _ := client.Dial(client.Options{})
    defer c.Close()

    // register workflows & activities using generated registration helpers, start worker
    w := worker.New(c, mutexv1.MutexTaskQueue, worker.Options{})
    mutexv1.RegisterActivities(w, &mutex.Activites{Client: mutexv1.NewClient(c)})
	  mutexv1.RegisterWorkflows(w, &mutex.Workflows{})
    _ := w.Start()
    defer w.Stop()

    // initialize generated client and execute workflow
    example := examplev1.NewClient(c)
    resp, err := example.HelloWorld(context.Background(), nil, &examplev1.HelloWorldRequest{})
    if err != nil {
      log.Fatalf("error executing %s workflow: %v", examplev1.HelloWorldWorkflowName, err)
    }
    log.Printf("hello world successful: %sv", resp)
}


```

### Advanced

Below is an excerpt from the more advanced [example](./example/) that showcases an implementation of the [mutex]((https://github.com/temporalio/samples-go/tree/main/mutex)) workflow from `github.com/temporalio/samples-go`

**Schema:**

```protobuf
syntax="proto3";

// buf:lint:ignore PACKAGE_DIRECTORY_MATCH
package mycompany.mutex.v1;

import "google/protobuf/duration.proto";
import "google/protobuf/empty.proto";
import "temporal/v1/temporal.proto";

service Mutex {
  option (temporal.v1.service) = {
    task_queue: "mutex-v1"
  };

  // ##########################################################################
  // Workflows
  // ##########################################################################

  // Mutex provides a mutex over a shared resource
  rpc Mutex(MutexRequest) returns (google.protobuf.Empty) {
    option (temporal.v1.workflow) = {
      default_options {
        id: 'mutex/${!resource}'
        id_reuse_policy: WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
        execution_timeout: { seconds: 3600 }
      }
      signal: { ref: 'AcquireLease', start: true }
      signal: { ref: 'RenewLease' }
      signal: { ref: 'RevokeLease' }
    };
    option (temporal.v1.activity) = {};
  }

  // SampleWorkflowWithMutex provides an example of a running workflow that uses
  // a Mutex workflow to prevent concurrent access to a shared resource
  rpc SampleWorkflowWithMutex(SampleWorkflowWithMutexRequest) returns (SampleWorkflowWithMutexResponse) {
    option (temporal.v1.workflow) = {
      default_options {
        id: 'sample-workflow-with-mutex/${!resource}/${!uuid_v4()}'
        id_reuse_policy: WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY
        execution_timeout: { seconds: 3600 }
      }
      signal: { ref: 'LeaseAcquired' }
    };
  }

  // ##########################################################################
  // Signals
  // ##########################################################################

  // AcquireLease enqueues a lease on the given resource
  rpc AcquireLease(AcquireLeaseRequest) returns (google.protobuf.Empty) {
    option (temporal.v1.signal) = {};
  }

  // LeaseAcquired notifies the calling workflow that a lease has been required
  rpc LeaseAcquired(LeaseAcquiredRequest) returns (google.protobuf.Empty) {
    option (temporal.v1.signal) = {};
  }

  // RenewLease extends the validity of an existing lease
  rpc RenewLease(RenewLeaseRequest) returns (google.protobuf.Empty) {
    option (temporal.v1.signal) = {};
  }

  // RevokeLease revokes an existing lease
  rpc RevokeLease(RevokeLeaseRequest) returns (google.protobuf.Empty) {
    option (temporal.v1.signal) = {};
  }
}

// ############################################################################
// Workflow Messages
// ############################################################################

// MutexRequest describes the input to a Mutex workflow/activity
message MutexRequest {
  string resource = 1;
}

// SampleWorkflowWithMutexRequest describes the input to a SampleWorkflowWithMutex workflow
message SampleWorkflowWithMutexRequest {
  string resource = 1;
  string dest = 2;
  double amount = 3;
}

// SampleWorkflowWithMutexResponse describes the output from a SampleWorkflowWithMutex workflow
message SampleWorkflowWithMutexResponse {
  string result = 1;
}

// ############################################################################
// Signal Messages
// ############################################################################

// AcquireLeaseRequest describes the input to a AcquireLease signal
message AcquireLeaseRequest {
  string workflow_id = 1;
  google.protobuf.Duration timeout = 2;
}

// LeaseAcquiredRequest describes the input to a LeaseAcquired signal
message LeaseAcquiredRequest {
  string workflow_id = 1;
  string run_id = 2;
  string lease_id = 3;
}

// RenewLeaseRequest describes the input to a RenewLease signal
message RenewLeaseRequest {
  string lease_id = 1;
  google.protobuf.Duration timeout = 2;
}

// RevokeLeaseRequest describes the input to a RevokeLease signal
message RevokeLeaseRequest {
  string lease_id = 1;
}
```

**Implementation:**

```go
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/cludden/protoc-gen-go-temporal/example/mutexv1"
    "github.com/google/uuid"
    "go.temporal.io/sdk/activity"
    "go.temporal.io/sdk/client"
    "go.temporal.io/sdk/log"
    "go.temporal.io/sdk/workflow"
    "go.temporal.io/sdk/worker"
    "google.golang.org/protobuf/types/known/durationpb"
)

func main() {
    c, _ := client.Dial(client.Options{})
    defer c.Close()

    w := worker.New(c, mutexv1.MutexTaskQueue, worker.Options{})
    mutexv1.RegisterActivities(w, &mutex.Activites{Client: mutexv1.NewClient(c)})
	  mutexv1.RegisterWorkflows(w, &mutex.Workflows{})
    w.Run(worker.InterruptCh())
}

// Workflows manages shared state for workflow constructors, local activities, side effects
type Workflows struct{}

// ============================================================================

// MutexWorkflow provides a mutex over a shared resource
type MutexWorkflow struct {
    *mutexv1.MutexInput
    log log.Logger
}

// LockAccount initializes a new MutexWorkflow value
func (w *Workflows) Mutex(ctx workflow.Context, input *mutexv1.MutexInput) (mutexv1.MutexWorkflow, error) {
    return &MutexWorkflow{input, log.With(workflow.GetLogger(ctx), "resource", input.Req.GetResource())}, nil
}

// Execute defines the entrypoint to a MutexWorkflow
func (wf *MutexWorkflow) Execute(ctx workflow.Context) error {
    for {
        wf.log.Info("dequeuing lease request")
        lease := wf.AcquireLease.ReceiveAsync()
        if lease == nil {
            return nil
        }

        wf.log.Info("generating lease id")
        var leaseID string
        if err := workflow.SideEffect(ctx, func(_ workflow.Context) any {
            return uuid.New().String()
        }).Get(&leaseID); err != nil {
            return fmt.Errorf("error generating lease id: %w", err)
        }
        log := log.With(wf.log, "lease", leaseID)

        log.Info("notifying lease holder")
        info := workflow.GetInfo(ctx)
        if err := mutexv1.LeaseAcquiredExternal(ctx, lease.GetWorkflowId(), "", &mutexv1.LeaseAcquiredSignal{
            WorkflowId: info.WorkflowExecution.ID,
            RunId:      info.WorkflowExecution.RunID,
            LeaseId:    leaseID,
        }).Get(ctx, nil); err != nil {
            return fmt.Errorf("error signalling lock acquired: %w", err)
        }

        log.Info("initializing lease timer")
        timerCtx, timerCancel := workflow.WithCancel(ctx)
        timer := workflow.NewTimer(timerCtx, lease.GetTimeout().AsDuration())

        for done := false; !done; {
            workflow.NewSelector(ctx).
                AddFuture(timer, func(f workflow.Future) {
                    log.Info("lease expired")
                    done = true
                }).
                AddReceive(wf.RenewLease.Channel, func(workflow.ReceiveChannel, bool) {
                    s := wf.RenewLease.ReceiveAsync()
                    if s.GetLeaseId() != leaseID {
                        return
                    }
                    log.Info("extending lease")
                    timerCancel()
                    timerCtx, timerCancel = workflow.WithCancel(ctx)
                    timer = workflow.NewTimer(timerCtx, s.GetTimeout().AsDuration())
                }).
                AddReceive(wf.RevokeLease.Channel, func(workflow.ReceiveChannel, bool) {
                    s := wf.RevokeLease.ReceiveAsync()
                    if s.GetLeaseId() != leaseID {
                        return
                    }
                    log.Info("revoking lease")
                    timerCancel()
                    done = true
                }).
                Select(ctx)
        }
    }
}

// ============================================================================

// SampleWorkflowWithMutexWorkflow simulates a long running workflow requiring exclusive access to a shared resource
type SampleWorkflowWithMutexWorkflow struct {
    *mutexv1.SampleWorkflowWithMutexInput
    log log.Logger
}

// SampleWorkflowWithMutex initializes a new SampleWorkflowWithMutexWorkflow value
func (w *Workflows) SampleWorkflowWithMutex(ctx workflow.Context, input *mutexv1.SampleWorkflowWithMutexInput) (mutexv1.SampleWorkflowWithMutexWorkflow, error) {
    return &SampleWorkflowWithMutexWorkflow{input, log.With(
        workflow.GetLogger(ctx), "resource", input.Req.GetResource(), "workflow", workflow.GetInfo(ctx).WorkflowExecution.ID,
    )}, nil
}

// Execute defines the entrypoint to a TransferWorkflow
func (wf *SampleWorkflowWithMutexWorkflow) Execute(ctx workflow.Context) (resp *mutexv1.SampleWorkflowWithMutexResponse, err error) {
    wf.log.Info("started")

    wf.log.Info("requesting lease")
    if err := mutexv1.Mutex(ctx, nil, &mutexv1.MutexRequest{Resource: wf.Req.GetResource()}).Get(ctx); err != nil {
        return nil, fmt.Errorf("error requesting lease: %w", err)
    }

    wf.log.Info("waiting until lease acquired")
    lease, _ := wf.LeaseAcquired.Receive(ctx)
    wf.log.Info("lease acquired", "lease", lease.GetLeaseId())
    defer func() {
        wf.log.Info("revoking lease", "lease", lease.GetLeaseId())
        cancelCtx, _ := workflow.NewDisconnectedContext(ctx)
        if mutexv1.RevokeLeaseExternal(cancelCtx, lease.GetWorkflowId(), lease.GetRunId(), &mutexv1.RevokeLeaseSignal{
            LeaseId: lease.GetLeaseId(),
        }).Get(ctx, nil); err != nil {
            wf.log.Error("error revoking lease", "error", err, "lease", lease.GetLeaseId())
        }
    }()

    // emulate long running process
    wf.log.Info("critical operation started")
    _ = workflow.Sleep(ctx, 10*time.Second)
    wf.log.Info("critical operation finished")

    return &mutexv1.SampleWorkflowWithMutexResponse{Result: lease.GetLeaseId()}, nil
}

// ============================================================================

// Activities manages shared state for activities
type Activites struct {
    Client mutexv1.Client
}

// Mutex locks a shared resource and can be called from a parent workflow
func (a *Activites) Mutex(ctx context.Context, req *mutexv1.MutexRequest) error {
    _, err := a.Client.StartMutexWithAcquireLease(ctx, nil, req, &mutexv1.AcquireLeaseSignal{
        WorkflowId: activity.GetInfo(ctx).WorkflowExecution.ID,
        Timeout:    durationpb.New(time.Minute * 2),
    })
    return err
}
```

## Options

See [temporal.proto](proto/temporal/v1/temporal.proto) for Service and Method options supported by this plugin.

### ID Expressions
Workflows can specify a default workflow ID that support [Bloblang](https://www.benthos.dev/docs/guides/bloblang/about) ID expressions. The expression is evaluated against a JSON-like input structure, allowing it to leverage fields from the Workflow's input parameter as well as Bloblang's native [functions](https://www.benthos.dev/docs/guides/bloblang/functions) and [methods](https://www.benthos.dev/docs/guides/bloblang/methods). 

**Example**

The following schema definition:
```protobuf
syntax="proto3"

package example.v1;

import "google/protobuf/empty.proto";
import "temporal/v1/temporal.proto";

service Example {
  rpc SayGreeting(SayGreetingRequest) returns (google.protobuf.Empty) {
    option (temporal.v1.workflow) = {
      default_options {
        id: 'say-greeting/${! greeting.or("hello").capitalize() }/${! subject.or("world").capitalize() }/${! uuid_v4() }'
      }
    };
  }
}

message SayGreetingRequest {
  string greeting = 1;
  string subject = 2;
}
```

Can be used like so:
```go
c, _ := client.Dial(client.Options{})
example := examplev1.NewClient(c)

run, _ := example.ExecuteSayGreeting(context.Background(), nil, &examplev1.SayGreetingRequest{})
require.Regexp(`^say-greeting/Hello/World/[a-f0-9-]{32}$`, run.ID())

run, _ := example.ExecuteSayGreeting(context.Background(), nil, &examplev1.SayGreetingRequest{
  Greeting: "howdy",
  Subject: "stranger",
})
require.Regexp(`^say-greeting/Howdy/Stranger/[a-f0-9-]{32}$`, run.ID())
```

## License
Licensed under the [MIT License](LICENSE.md)  
Copyright for portions of project cludden/protoc-gen-go-temporal are held by Chad Retz, 2021 as part of project cretz/temporal-sdk-go-advanced. All other copyright for project cludden/protoc-gen-go-temporal are held by Chris Ludden, 2023.