# protoc-gen-go-temporal

A protoc plugin for generating typed Temporal clients and workers in Go from protobuf schemas. This plugin allows Workflow authors to configure sensible defaults and guardrails, simplifies the implementation and testing of Temporal workers, and streamlines integration by providing typed client SDKs and a generated CLI application. 

<small>inspired by [Chad Retz's](https://github.com/cretz/) awesome [github.com/cretz/temporal-sdk-go-advanced](https://github.com/cretz/temporal-sdk-go-advanced) and [Jacob LeGrone's](https://github.com/jlegrone/) excellent Replay talk on [Temporal @ Datadog](https://youtu.be/LxgkAoTSI8Q)</small>

**Table of Contents**

- [protoc-gen-go-temporal](#protoc-gen-go-temporal)
  - [How it works](#how-it-works)
  - [Features](#features)
  - [Getting Started](#getting-started)
  - [Options](#options)
    - [Plugin Options](#plugin-options)
    - [Service Options](#service-options)
    - [Method Options](#method-options)
    - [Bloblang Expressions](#bloblang-expressions)
  - [CLI](#cli)
  - [Test Client](#test-client)
  - [Cross-Namespace (XNS)](#cross-namespace-xns)
  - [Documentation](#documentation)
  - [License](#license)

## How it works

1. Annotate your protobuf services and methods with Temporal options provided by this plugin
2. Generate Go code that includes types, methods, and functions for implementing Temporal clients, workers, and cli applications
3. Define implementations for the required Workflow and Activity interfaces
4. Run your Temporal worker using the generated helpers and interact with it using the generated client and/or cli

## Features

Generated **Client** with:
  - methods for executing workflows, queries, signals, and updates
  - methods for cancelling or terminating workflows
  - default `client.StartWorkflowOptions` and `client.UpdateWorkflowWithOptionsRequest`
  - dynamic workflow ids, update ids, and search attributes via [Bloblang expressions](#bloblang-expressions)
  - default timeouts, id reuse policies, retry policies, wait policies
  - experimental [cross-namespace (xns)](#cross-namespace-xns) support


Generated **Worker** resources with:
  - functions for calling activities and local activities from workflows
  - functions for executing child workflows and signalling external workflows
  - default `workflow.ActivityOptions`, `workflow.ChildWorkflowOptions`
  - default timeouts, parent cose policies, retry policies


Optional **CLI** with:
  - commands for executing workflows, synchronously or asynchronously
  - commands for starting workflows with signals, synchronously or asynchronously
  - commands for querying existing workflows
  - commands for sending signals to existing workflows
  - typed flags for conventiently specifying workflow, query, and signal inputs


## Getting Started
1. Install [buf](https://docs.buf.build/installation)
   
2. Install this plugin
   1. via `homebrew`
    ```shell
    brew install cludden/formula/protoc-gen-go_temporal
    ```
   2. manually by grabbing a binary for your OS from [the releases page](https://github.com/cludden/protoc-gen-go-temporal/releases) and placing it in your $PATH
   3. via `go`:
    ```shell
    go install github.com/cludden/protoc-gen-go-temporal/cmd/protoc-gen-go_temporal@<version>
    ```

3. Install Go protoc plugin
  ```shell
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  ```

4. Initialize buf repository
  ```shell
  mkdir proto && cd proto && buf mod init
  ```

5. Add dependency to `buf.yaml`
  ```yaml
  version: v1
  deps:
    - buf.build/cludden/protoc-gen-go-temporal:<version>
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
      opt: paths=source_relative,cli-enabled=true,cli-categories=true,workflow-update-enabled=true
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
    };

    // CreateFoo creates a new foo operation
    rpc CreateFoo(CreateFooRequest) returns (CreateFooResponse) {
      option (temporal.v1.workflow) = {
        execution_timeout: { seconds: 3600 } // foos can take awhile to create
        id_reuse_policy: WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE
        id: 'create-foo/${!name.slug()}'
        query: { ref: 'GetFooProgress' }
        signal: { ref: 'SetFooProgress', start: true }
        update: { ref: 'UpdateFooProgress' }
      };
    }

    // GetFooProgress returns the status of a CreateFoo operation
    rpc GetFooProgress(google.protobuf.Empty) returns (GetFooProgressResponse) {
      option (temporal.v1.query) = {};
    }

    // Notify sends a notification
    rpc Notify(NotifyRequest) returns (google.protobuf.Empty) {
      option (temporal.v1.activity) = {
        start_to_close_timeout: { seconds: 30 }
        retry_policy: {
          max_attempts: 3
        }
      };
    }

    // SetFooProgress sets the current status of a CreateFoo operation
    rpc SetFooProgress(SetFooProgressRequest) returns (google.protobuf.Empty) {
      option (temporal.v1.signal) = {};
    }

    // UpdateFooProgress sets the current status of a CreateFoo operation
    rpc UpdateFooProgress(SetFooProgressRequest) returns (GetFooProgressResponse) {
      option (temporal.v1.update) = {
        id: 'update-progress/${! progress.string() }',
      };
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
  buf mod update && buf generate
  ```

9. Implement the required Workflow and Activity interfaces
  ```go
  package main

  import (
    "context"
    "fmt"
    "log"
    "os"

    examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
    "github.com/urfave/cli/v2"
    "go.temporal.io/sdk/activity"
    "go.temporal.io/sdk/client"
    "go.temporal.io/sdk/worker"
    "go.temporal.io/sdk/workflow"
    logger "go.temporal.io/server/common/log"
  )

  // Workflows manages shared state for workflow constructors and is used to
  // register workflows with a worker
  type Workflows struct{}

  // ============================================================================

  // CreateFooWorkflow manages workflow state for a CreateFoo workflow
  type CreateFooWorkflow struct {
    // it embeds the generated workflow Input type that contains the workflow
    // input and signal helpers
    *examplev1.CreateFooWorkflowInput

    progress float32
    status   examplev1.Foo_Status
  }

  // CreateFoo implements a CreateFoo workflow constructor on the shared Workflows struct
  // that initializes a new CreateFooWorkflow for each execution
  func (w *Workflows) CreateFoo(ctx workflow.Context, input *examplev1.CreateFooInput) (examplev1.CreateFooWorkflow, error) {
    return &CreateFooWorkflow{input, 0, examplev1.Foo_FOO_STATUS_CREATING}, nil
  }

  // Execute defines the entrypoint to a CreateFooWorkflow
  func (wf *CreateFooWorkflow) Execute(ctx workflow.Context) (*examplev1.CreateFooResponse, error) {
    // listen for signals
    workflow.Go(ctx, func(ctx workflow.Context) {
      for {
        signal, _ := wf.SetFooProgress.Receive(ctx)
        wf.UpdateFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: signal.GetProgress()})
      }
    })

    // execute Notify activity using generated helper
    err := examplev1.Notify(ctx, &examplev1.NotifyRequest{
      Message: fmt.Sprintf("creating foo resource (%s)", wf.Req.GetName()),
    })
    if err != nil {
      return nil, fmt.Errorf("error sending notification: %w", err)
    }

    // block until progress has reached 100 via signals and/or updates
    workflow.Await(ctx, func() bool {
      return wf.status == examplev1.Foo_FOO_STATUS_READY
    })

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

  // UpdateFooProgress defines the handler for an UpdateFooProgress update
  func (wf *CreateFooWorkflow) UpdateFooProgress(ctx workflow.Context, req *examplev1.SetFooProgressRequest) (*examplev1.GetFooProgressResponse, error) {
    wf.progress = req.GetProgress()
    switch {
    case wf.progress < 0:
      wf.progress, wf.status = 0, examplev1.Foo_FOO_STATUS_CREATING
    case wf.progress < 100:
      wf.status = examplev1.Foo_FOO_STATUS_CREATING
    case wf.progress >= 100:
      wf.progress, wf.status = 100, examplev1.Foo_FOO_STATUS_READY
    }
    return &examplev1.GetFooProgressResponse{Progress: wf.progress, Status: wf.status}, nil
  }

  // ============================================================================

  // Activities manages shared state for activities and is used to register
  // activities with a worker
  type Activities struct{}

  // Notify defines the implementation for a Notify activity
  func (a *Activities) Notify(ctx context.Context, req *examplev1.NotifyRequest) error {
    activity.GetLogger(ctx).Info("notification", "message", req.GetMessage())
    return nil
  }

  // ============================================================================

  func main() {
    // initialize the generated cli application
    app, err := examplev1.NewExampleCli(
      examplev1.NewExampleCliOptions().
        WithClient(func(cmd *cli.Context) (client.Client, error) {
          return client.Dial(client.Options{
            Logger: logger.NewSdkLogger(logger.NewCLILogger()),
          })
        }).
        WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
          w := worker.New(c, examplev1.ExampleTaskQueue, worker.Options{})
          // register activities and workflows using generated helpers
          examplev1.RegisterExampleActivities(w, &Activities{})
          examplev1.RegisterExampleWorkflows(w, &Workflows{})
          return w, nil
        }),
    )
    if err != nil {
      log.Fatalf("error initializing commands: %v", err)
    }
    app.Name = "example"
    app.Usage = "an example temporal cli"

    // run cli
    if err := app.Run(os.Args); err != nil {
      log.Fatal(err)
    }
  }
  ```

10. Run your worker
    
  *start temporal*
  ```shell
  temporal server start-dev --dynamic-config-value "frontend.enableUpdateWorkflowExecution=true"
  ```

  *start worker*
  ```shell
  go get -u github.com/cludden/protoc-gen-go-temporal@<release> && go mod tidy
  go run main.go worker
  ```

11.  Execute workflows, queries, signals, and updates
  
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

    run, _ := client.CreateFooAsync(ctx, &examplev1.CreateFooRequest{Name: "test"})
    log.Printf("started workflow: workflow_id=%s, run_id=%s\n", run.ID(), run.RunID())

    log.Println("signalling progress")
    _ = run.SetFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 5.7})

    progress, _ := run.GetFooProgress(ctx)
    log.Printf("queried progress: %s\n", progress.String())

    update, _ := run.UpdateFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 100})
    log.Printf("updated progress: %s\n", update.String())

    resp, _ := run.Get(ctx)
    log.Printf("workflow completed: %s\n", resp.String())
  }
  ```

  *with generated cli*
  ```shell
  $ go run main.go -h
  NAME:
  Example - an example temporal cli

  USAGE:
  Example [global options] command [command options] [arguments...]

  COMMANDS:
  worker   run service worker
  help, h  Shows a list of commands or help for one command
  QUERIES:
    get-foo-progress  GetFooProgress returns the status of a CreateFoo operation
  SIGNALS:
    set-foo-progress  SetFooProgress sets the current status of a CreateFoo operation
  UPDATES:
    update-foo-progress  UpdateFooProgress sets the current status of a CreateFoo operation
  WORKFLOWS:
    create-foo                        CreateFoo creates a new foo operation
    create-foo-with-set-foo-progress  sends a SetFooProgress signal to a CreateFoo workflow, starting it if necessary

  GLOBAL OPTIONS:
  --help, -h  show help (default: false)

  $ go run main.go create-foo -d --name test
  success
  workflow id: create-foo/test
  run id: 44cacae1-6a13-4b4a-8db7-d29eaafd1499

  $ go run main.go set-foo-progress -w create-foo/test --progress 5.7
  success

  $ go run main.go get-foo-progress -w create-foo/test
  {
    "progress": 5.7,
    "status": "FOO_STATUS_CREATING"
  }

  $ go run main.go update-foo-progress -w create-foo/test --progress 100
  {
    "progress": 100,
    "status": "FOO_STATUS_READY"
  }

  $ go run main.go get-foo-progress -w create-foo/test
  {
    "progress": 100,
    "status": "FOO_STATUS_READY"
  }
  ```

## Options

See [reference documentation](./docs/api/temporal/v1/api.md) for all Service and Method options supported by this plugin.

### Plugin Options

Plugin options are used to globally configure this plugin's behavior at runtime. They can be specified
via:
- protoc commaond line flags

  ```shell
  --go_temporal_opt=<flag>=<value>,<flag>=<value>
  ```

- buf generate options

  ```yaml
  # buf.gen.yaml
  plugins:
  # ..
    - plugin: go_temporal
      out: gen
      opt: <flag>=<value>,<flag>=<value>
      strategy: all
  ```

| flag | type | description | default |
| :--- | :---: | :--- | :---: |
| cli-categories | `bool` | enables cli categories | `true` |
| cli-enabled | `bool` | enables cli generation | `false` |
| disable-workflow-input-rename | `bool` | disables renamed workflow input suffix | `false` |
| enable-patch-support | `bool` | enables experimental support for [protoc-gen-go-patch](https://github.com/alta/protopatch) | `false` |
| enable-xns | `bool` | enables [experimental cross-namespace support](#cross-namespace-xns) | `false` |
| workflow-update-enabled | `bool` | enables experimental workflow update | `false` |

### Service Options

| field | type | description |
| :--- | :---: | :--- |
| task_queue | `string` | default task queue for all workflows, activities |

*Example*
```protobuf
syntax="proto3";

import "temporal/v1/temporal.proto";

service Example {
  option (temporal.v1.service) = {
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
| update | [UpdateOptions](./docs/api/temporal/v1/api.md#updateoptions) | default settings for Temporal updates. workflow updates are considered experimental and must be enabled via [service-level features](./docs/api/temporal/v1/api.md#serviceoptionsfeatures)|
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
      update: { ref: 'MyUpdate' }
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

  rpc MyUpdate(MyUpdateRequest) returns (MyUpdateResponse) {
    option (temporal.v1.update) = {
      id: 'my-update/${!uuid_v4()}'
      validate: true
    };
  }
}
```

### Bloblang Expressions
Default workflow IDs, update IDs, and search attributes can be defined using [Bloblang](https://www.benthos.dev/docs/guides/bloblang/about) expressions via the `${!<expression>}` interpolation syntax. The expression is evaluated against the protojson serialized input, allowing it to leverage fields from the input parameter, as well as Bloblang's native [functions](https://www.benthos.dev/docs/guides/bloblang/functions) and [methods](https://www.benthos.dev/docs/guides/bloblang/methods). 

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
      id: 'say-greeting/${! greeting.or("hello").capitalize() }/${! subject.or("world").capitalize() }/${! uuid_v4() }'
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

run, _ := example.ExecuteSayGreeting(context.Background(), &examplev1.SayGreetingRequest{})
require.Regexp(`^say-greeting/Hello/World/[a-f0-9-]{32}$`, run.ID())

run, _ := example.ExecuteSayGreeting(context.Background(), &examplev1.SayGreetingRequest{
  Greeting: "howdy",
  Subject: "stranger",
})
require.Regexp(`^say-greeting/Howdy/Stranger/[a-f0-9-]{32}$`, run.ID())
```

## CLI

This plugin can optionally generate a configurable CLI using [github.com/urfave/cli/v2](https://github.com/urfave/cli/v2). To enable this functionality, use the corresponding [plugin option](#plugin-options). When enabled, this plugin will generate a CLI command for each workflow, start-workflow-with-signal, query, and signal. Each command provides typed flags for configuring the corresponding inputs and options.

```go
package main

import (
  "log"
  "os"

  example "github.com/cludden/protoc-gen-go-temporal/example"
  examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
  "github.com/urfave/cli/v2"
  "go.temporal.io/sdk/client"
  "go.temporal.io/sdk/worker"
)

func main() {
  app, err := examplev1.NewExampleCli(
    examplev1.NewExampleCliOptions().
      WithClient(func(cmd *cli.Context) (client.Client, error) {
        return client.Dial(client.Options{})
      }).
      WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
        w := worker.New(c, examplev1.ExampleTaskQueue, worker.Options{})
        examplev1.RegisterExampleWorkflows(w, &example.Workflows{})
        examplev1.RegisterExampleActivities(w, &example.Activities{})
        return w, nil
      }),
  )
  if err != nil {
    log.Fatalf("error initializing cli: %v", err)
  }
  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}
```

## Test Client

The generated code includes resources that are compatible with the Temporal Go SDK's [testsuite](https://pkg.go.dev/go.temporal.io/sdk@v1.23.1/testsuite) module. See [tests](./test/simple/main_test.go) for example usage.

**_Note:_** that all queries, signals, and udpates must be called via the test environment's `RegisterDelayedCallback` method prior to invoking the test client's synchronous `<Workflow>` method or an asynchronous workflow run's `Get` method.

## Cross-Namespace (XNS)

*__Experimental__*

This plugin provides experimental support for cross-namespace and/or cross-cluster integration by enabling the `enable-xns` plugin option. When enabled, the plugin will generate an additional `path/to/generated/code/<package>xns` go package containing types, methods, and helpers for calling workflows, queries, signals, and updates from other Temporal workflows via activities. The activities use [heartbeating](https://docs.temporal.io/activities#activity-heartbeat) to maintain liveness for long-running workflows or updates, and their associated timeouts can be configured using the generated options helpers. For an example of xns integration, see the [example/external](./example/external/external.go) package.

## Documentation

- [Generated code reference](./docs/generated.md)

## License
Licensed under the [MIT License](LICENSE.md)  
Copyright for portions of project cludden/protoc-gen-go-temporal are held by Chad Retz, 2021 as part of project cretz/temporal-sdk-go-advanced. All other copyright for project cludden/protoc-gen-go-temporal are held by Chris Ludden, 2023.
