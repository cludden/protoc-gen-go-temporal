# protoc-gen-go-temporal

a protoc plugin for generating temporal clients and workers in Go from protobuf schemas

inspired by [github.com/cretz/temporal-sdk-go-advanced](https://github.com/cretz/temporal-sdk-go-advanced)

**Table of Contents**

- [protoc-gen-go-temporal](#protoc-gen-go-temporal)
	- [Features](#features)
	- [Getting Started](#getting-started)
	- [Options](#options)
		- [Service Options](#service-options)
		- [Method Options](#method-options)
		- [ID Expressions](#id-expressions)
	- [CLI](#cli)
	- [License](#license)

## Features

- typed client with:
  - methods for executing workflows, queries, and signals
  - methods for cancelling or terminating workflows
  - default `client.StartWorkflowOptions`
  - dynamic workflow and update ids via [Bloblang expressions](#id-expressions)
  - default timeouts, id reuse policies, retry policies
- typed worker helpers with:
  - functions for calling activities and local activities from workflows
  - functions for executing child workflows and signalling external workflows
  - default `workflow.ActivityOptions`, `workflow.ChildWorkflowOptions`
  - default timeouts, parent cose policies, retry policies
- configurable CLI with:
  - commands for executing workflows, synchronously or asynchronously
  - commands for starting workflows with signals, synchronously or asynchronously
  - commands for querying existing workflwos
  - commands for sending signals to existing workflows
  - typed flags for conventiently specifying workflow, query, and signal inputs

## Getting Started
1. Install [buf](https://docs.buf.build/installation)

2. Install Go protoc plugin
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

3. Install this plugin

	3a. Grab a binary for your OS from [the releases page](https://github.com/cludden/protoc-gen-go-temporal/releases) and place in your $PATH

	3b. `go install github.com/cludden/protoc-gen-go-temporal/cmd/protoc-gen-go_temporal@<version>`

4. Initialize buf repository
```shell
mkdir proto && cd proto && buf init
```

5. Add dependency to `buf.yaml`
```yaml
version: v1
deps:
  - buf.build/cludden/protoc-gen-go-temporal@<version>
```

6. Add plugin to `buf.gen.yaml` and exclude it from managed mode go prefix
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

7. Define your service  
<small><b><i>note:</i></b> see [example](./example/) and [test](./test/) for more details on generated code and usage</small>

```protobuf
syntax="proto3";

package example.v1;

import "google/protobuf/empty.proto";
import "temporal/v1/temporal.proto";

service Example {
  option (temporal.v1.service) = {
    task_queue: "example-v1"
    features: { cli: CLI_FEATURE_ENABLED }
  };

  // CreateFoo creates a new foo operation
  rpc CreateFoo(CreateFooRequest) returns (CreateFooResponse) {
    option (temporal.v1.workflow) = {
      default_options {
        id: 'create-foo/${!name.slug()}'
        id_reuse_policy: WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
        execution_timeout: { seconds: 3600 }
      }
      query: { ref: 'GetFooProgress' }
      signal: { ref: 'SetFooProgress', start: true }
    };
  }

  // GetFooProgress returns the status of a CreateFoo operation
  rpc GetFooProgress(google.protobuf.Empty) returns (GetFooProgressResponse) {
    option (temporal.v1.query) = {};
  }

  // Notify sends a notification
  rpc Notify(NotifyRequest) returns (google.protobuf.Empty) {
    option (temporal.v1.activity) = {
      default_options {
        start_to_close_timeout: { seconds: 30 }
        retry_policy: {
          max_attempts: 3
        }
      }
    };
  }

  // SetFooProgress sets the current status of a CreateFoo operation
  rpc SetFooProgress(SetFooProgressRequest) returns (google.protobuf.Empty) {
    option (temporal.v1.signal) = {};
  }
}

// CreateFooRequest describes the input to a CreateFoo workflow
message CreateFooRequest {
  // unique foo name
  string name = 1;
}

// SampleWorkflowWithMutexResponse describes the output from a CreateFoo workflow
message CreateFooResponse {
  Foo foo = 1; 
}

// Foo describes an illustrative foo resource
message Foo {
  string name = 1;
  Status status = 2;

  enum Status {
    FOO_STATUS_UNSPECIFIED = 0;
    FOO_STATUS_READY = 1;
    FOO_STATUS_CREATING = 2;
  }
}

// GetFooProgressResponse describes the output from a GetFooProgress query
message GetFooProgressResponse {
  float progress = 1;
  Foo.Status status = 2;
}

// NotifyRequest describes the input to a Notify activity
message NotifyRequest {
  string message = 1;
}

// SetFooProgressRequest describes the input to a SetFooProgress signal
message SetFooProgressRequest {
  // value of current workflow progress
  float progress = 1;
}
```

8. Generate temporal worker, client, and cli types, methods, interfaces, and functions

```shell
buf generate
```

9. Implement your activities, workflows

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"

	examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	logger "go.temporal.io/server/common/log"
)

// Workflows manages shared state for workflow constructors
type Workflows struct{}

// ============================================================================

// CreateFooWorkflow creates a new Foo resource
type CreateFooWorkflow struct {
	*examplev1.CreateFooInput
	progress float32
	status   examplev1.Foo_Status
}

// CreateFoo initializes a new CreateFooWorkflow
func (w *Workflows) CreateFoo(ctx workflow.Context, input *examplev1.CreateFooInput) (examplev1.CreateFooWorkflow, error) {
	return &CreateFooWorkflow{input, 0, examplev1.Foo_FOO_STATUS_UNSPECIFIED}, nil
}

// Execute defines the entrypoint to a CreateFooWorkflow
func (wf *CreateFooWorkflow) Execute(ctx workflow.Context) (*examplev1.CreateFooResponse, error) {
	// execute Notify activity using generated helper
	if err := examplev1.Notify(ctx, nil, &examplev1.NotifyRequest{Message: fmt.Sprintf("creating foo resource (%s)", wf.Req.GetName())}).Get(ctx); err != nil {
		return nil, fmt.Errorf("error sending notification: %w", err)
	}

	// wait until signalled progress reaches 100
	for progress := float32(0); progress < 100; {
		signal, _ := wf.SetFooProgress.Receive(ctx)
		progress = signal.GetProgress()
		switch {
		case progress < 0:
			progress, wf.status = 0, examplev1.Foo_FOO_STATUS_UNSPECIFIED
		case progress < 100:
			wf.status = examplev1.Foo_FOO_STATUS_CREATING
		case progress >= 100:
			progress, wf.status = 100, examplev1.Foo_FOO_STATUS_READY
		}
		wf.progress = progress
	}

	return &examplev1.CreateFooResponse{
		Foo: &examplev1.Foo{
			Name:   wf.Req.GetName(),
			Status: wf.status,
		},
	}, nil
}

// GetFooProgress defines the handler for a GetFooProgress query
func (wf *CreateFooWorkflow) GetFooProgress() (*examplev1.GetFooProgressResponse, error) {
	return &examplev1.GetFooProgressResponse{Progress: wf.progress, Status: wf.status}, nil
}

// ============================================================================

// Activities manages shared state for activities
type Activities struct{}

// Notify sends a notification
func (a *Activities) Notify(ctx context.Context, req *examplev1.NotifyRequest) error {
	activity.GetLogger(ctx).Info("notification", "message", req.GetMessage())
	return nil
}

// ============================================================================

func main() {
	// initialize cli app
	app := &cli.App{
		Name:  "Example",
		Usage: "an example temporal cli",
		// add cleanup logic to parent app
		After: func(cmd *cli.Context) error {
			if c, ok := cmd.App.Metadata["client"]; ok {
				c.(client.Client).Close()
			}
			return nil
		},
	}

	// initialize client commands using generated constructor
	var err error
	if app.Commands, err = examplev1.NewCommands(
		// provide a client initializer for use by commands
		examplev1.WithClientForCommand(func(cmd *cli.Context) (client.Client, error) {
			c, err := client.Dial(client.Options{
				Logger: logger.NewSdkLogger(logger.NewCLILogger()),
			})
			if err != nil {
				return nil, fmt.Errorf("error initializing client: %w", err)
			}
			// set a reference to the client in app metadata for use by app cleanup
			cmd.App.Metadata["client"] = c
			return c, nil
		}),
	); err != nil {
		log.Fatalf("error initializing commands: %v", err)
	}

	// add worker command
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "worker",
		Usage: "run service worker",
		Action: func(cmd *cli.Context) error {
			// initialize temporal client
			c, err := client.Dial(client.Options{})
			if err != nil {
				return fmt.Errorf("error initializing client: %w", err)
			}
			defer c.Close()

			// register workflows & activities using generated registration helpers, start worker
			w := worker.New(c, examplev1.ExampleTaskQueue, worker.Options{})
			examplev1.RegisterActivities(w, &Activities{})
			examplev1.RegisterWorkflows(w, &Workflows{})
			if err := w.Start(); err != nil {
				return fmt.Errorf("error starting worker: %w", err)
			}
			defer w.Stop()

			<-cmd.Context.Done()
			return nil
		},
	})
	sort.Slice(app.Commands, func(i, j int) bool {
		return app.Commands[i].Name < app.Commands[j].Name
	})

	// run cli
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
```

10. Run your worker

*start temporal*
```shell
temporal server start-dev
```

*start worker*
```shell
go run example/main.go worker
```

11.  Execute workflows, signals, queries

*with generated client*
```go
package main

import (
	"context"
	"log"

	examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	"go.temporal.io/sdk/client"
)

func main() {
	c, _ := client.Dial(client.Options{})
	client, ctx := examplev1.NewClient(c), context.Background()

	run, _ := client.ExecuteCreateFoo(ctx, nil, &examplev1.CreateFooRequest{Name: "test"})
	log.Printf("started workflow: workflow_id=%s, run_id=%s\n", run.ID(), run.RunID())

	_ = run.SetFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 5.7})
	log.Println("signalling progress")

	progress, _ := run.GetFooProgress(ctx)
	log.Printf("querying progress: %s\n", progress.String())

	_ = run.SetFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 100})
	log.Println("signalling progress")

	resp, _ := run.Get(ctx)
	log.Println("workflow completed: %s\n", resp.String())
}
```

*with generated cli*
```shell
$ go run example/main.go create-foo -d --name test
success
workflow_id: create-foo/test
execution_id: 44cacae1-6a13-4b4a-8db7-d29eaafd1499

$ go run example/main.go set-foo-progress -w create-foo/test --progress 5.7
success

$ go run example/main.go get-foo-progress -w create-foo/test
{
  "progress": 5.7,
  "status": 2
}

$ go run example/main.go set-foo-progress -w create-foo/test --progress 100
success

$ go run example/main.go get-foo-progress -w create-foo/test
{
  "progress": 100,
  "status": 1
}
```

## Options

See [reference documentation](./docs/api/temporal/v1/api.md) for all Service and Method options supported by this plugin.

### Service Options

| field | type | description |
| :--- | :---: | :--- |
| features | [Features](./docs/api/temporal/v1/api.md#serviceoptionsfeatures) | specifies settings for optional features |
| namespace | `string` | default namespace for child workflows, activities |
| task_queue | `string` | default task queue for all workflows, activities |

*Example*
```protobuf
syntax="proto3";

import "temporal/v1/temporal.proto";

service Example {
	option(temporal.v1.service) = {
		features: { cli: CLI_FEATURE_ENABLED }
		task_queue: 'example-v1'
	};
}
```

### Method Options

| field | type | description |
| :--- | :---: | :--- |
| activity | [ActivityOptions](./docs/api/temporal/v1/api.md#activityoptions) | default settings for Temporal activities |
| query | [QueryOptions](./docs/api/temporal/v1/api.md#queryoptions) | default settings for Temporal queries |
| signal | [SignalOptions](./docs/api/temporal/v1/api.md#signaloptions) | default settings for Temporal signals |
| workflow | [WorkflowOptions](./docs/api/temporal/v1/api.md#workflowoptions) | default settings for Temporal workflows |

*Example*
```protobuf
syntax="proto3";

import "google/protobuf/empty.proto";
import "temporal/v1/temporal.proto";

service Example {
	rpc MyWorkflow(MyWorkflowRequest) returns (MyWorkflowResponse) {
		option (temporal.v1.workflow) = {
			default_options: {
				id: 'my-workflow/${! uuid_v4() }'
				execution_timeout: { seconds: 3600 }
			}
			query: { ref: 'MyQuery' }
			signal: { ref: 'MySignal', start: true }
		};
	}

	rpc MyActivity(MyActivityRequest) returns (MyActivityResponse) {
		option (temporal.v1.activity) = {
			default_options: {
				start_to_close_timeout: { seconds: 30 }
				retry_policy: {
					max_attempts: 3
				}
			}
		};
	}

	rpc MyQuery(MyQueryRequest) returns (MyQueryResponse) {
		option (temporal.v1.query) = {};
	}

	rpc MySignal(MySignalRequest) returns (google.protobuf.Empty) {
		option (temporal.v1.signal) = {};
	}
}
```

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

## CLI

This plugin can optionally generate a configurable CLI using [github.com/urfave/cli/v2](https://github.com/urfave/cli/v2). To enable this functionality, use the corresponding [service option](#service-options). When enabled, this plugin will generate a CLI command for each workflow, start-workflow-with-signal, query, and signal. Each command provides typed flags for configuring the corresponding inputs and options.

```go
package main

import (
	"fmt"
	"log"
	"os"

	examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
)

func main() {
	// initialize commands using generated constructor
	commands, err := examplev1.NewCommands(
		// provide a client factory for use by commands
		examplev1.WithClientForCommand(func(cmd *cli.Context) (client.Client, error) {
			c, err := client.Dial(client.Options{})
			if err != nil {
				return nil, fmt.Errorf("error initializing client: %w", err)
			}
			// set a reference to the client in app metadata for use by app cleanup
			cmd.App.Metadata["client"] = c
			return c, nil
		}),
	)
	if err != nil {
		log.Fatalf("error initializing commands: %v", err)
	}

	// run cli
	if err := (&cli.App{
		Name:     "Example",
		Usage:    "an example temporal cli",
		Commands: commands,
		// add cleanup logic to global "After" hook
		After: func(cmd *cli.Context) error {
			if c, ok := cmd.App.Metadata["client"]; ok {
				c.(client.Client).Close()
			}
			return nil
		},
	}).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
```

## License
Licensed under the [MIT License](LICENSE.md)  
Copyright for portions of project cludden/protoc-gen-go-temporal are held by Chad Retz, 2021 as part of project cretz/temporal-sdk-go-advanced. All other copyright for project cludden/protoc-gen-go-temporal are held by Chris Ludden, 2023.