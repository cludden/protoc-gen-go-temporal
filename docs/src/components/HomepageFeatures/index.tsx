import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';
import { useColorMode } from '@docusaurus/theme-common';
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import CodeBlock from "@theme/CodeBlock";
import Link from '@docusaurus/Link';
import useBaseUrl from '@docusaurus/useBaseUrl';
import About from './_about.mdx';

const examples = [
  {
    label: 'Annotate',
    content: 'Annotate your protobuf services and methods with Temporal options.',
    fragments: [
      {
        language: 'protobuf',
        title: 'proto/helloworld/v1/example.proto',
        content: `syntax="proto3";

package helloworld.v1;

import "google/protobuf/empty.proto";
import "temporal/v1/temporal.proto";

service Example {
  option (temporal.v1.service) = {
    task_queue: "hello-world"
  };

  // Hello prints a friendly greeting and waits for goodbye
  rpc Hello(HelloRequest) returns (HelloResponse) {
    option (temporal.v1.workflow) = {
      name: "helloworld.v1.Hello"
      id: 'hello/\${! name.or("World") }'
      signal: { ref: "Goodbye" }
    };
  }

  // Goodbye signals a running workflow to exit
  rpc Goodbye(GoodbyeRequest) returns (google.protobuf.Empty) {
    option (temporal.v1.signal) = {};
  }
}

// HelloRequest describes the input to a Hello workflow
message HelloRequest {
  string name = 1;
}

// HelloResponse describes the output from a Hello workflow
message HelloResponse {
  string result = 1;
}

message GoodbyeRequest {
  string message = 1;
}
        `,
      }
    ],
  },
  {
    label: 'Generate',
    content: 'Generate Go code for implementing Temporal Clients, Workers, and CLI applications.',
    fragments: [
      {
        language: 'yaml',
        title: 'buf.gen.yaml',
        content: `version: v1
managed:
  enabled: true
  go_package_prefix:
    default: example/gen
    except:
      - buf.build/cludden/protoc-gen-go-temporal
plugins:
  - plugin: go
    out: gen
    opt: paths=source_relative
  - plugin: go_temporal
    out: gen
    opt: paths=source_relative,cli-enabled=true,cli-categories=true,workflow-update-enabled=true,docs-out=./proto/README.md
    strategy: all
    `
      },
      {
        language: 'sh',
        content: `buf generate`
      }
    ],
  },
  {
    label: 'Implement',
    content: 'Implement the required Workflow and Activity interfaces',
    fragments: [
      {
        language: 'go',
        title: 'internal/example/example.go',
        content: `package example

import (
  helloworldv1 "path/to/gen/helloworld/v1"
  "go.temporal.io/sdk/log"
  "go.temporal.io/sdk/workflow"
)

type (
  Workflows struct{}

  // HelloWorkflow provides a helloworldv1.HelloWorkflow implementation
  HelloWorkflow struct {
    *helloworldv1.HelloWorkflowInput
    log log.Logger
  }
)

// NewHelloWorkflow initializes a new helloworldv1.HelloWorkflow value
func (w *Workflows) Hello(ctx workflow.Context, input *helloworldv1.HelloWorkflowInput) (helloworldv1.HelloWorkflow, error) {
  return &HelloWorkflow{input, workflow.GetLogger(ctx)}, nil
}

// Execute defines the entrypoint to a Hello workflow
func (w *HelloWorkflow) Execute(ctx workflow.Context) (*helloworldv1.HelloResponse, error) {
  w.log.Info("Hello workflow started", "request", w.Req)

  goodbye, _ := w.Goodbye.Receive(ctx)
  w.log.Info("Goodbye received", "signal", goodbye)

  return &helloworldv1.HelloResponse{}, nil
}
    `
      }
    ],
  },
  {
    label: 'Run',
    content: 'Run your Temporal Worker using the generated helpers.',
    fragments: [
      {
        language: 'go',
        title: 'main.go',
        content: `package main

import (
  "log"
  "os"

  example "internal"
  helloworldv1 "path/to/gen/helloworld/v1"
  "github.com/urfave/cli/v2"
  "go.temporal.io/sdk/client"
  "go.temporal.io/sdk/worker"
)

func main() {
  app, err := helloworldv1.NewExampleCli(
    helloworldv1.NewExampleCliOptions().
      WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
        w := worker.New(c, helloworldv1.ExampleTaskQueue, worker.Options{})
        helloworldv1.RegisterExampleWorkflows(w, &example.Workflows{})
        return w, nil
      }),
  )
  if err != nil {
    log.Fatal(err)
  }

  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}
    `
      }
    ],
  },
  {
    label: 'Client',
    content: 'Interact with your workers from any Go application using the generated Client.',
    fragments: [
      {
        language: 'go',
        title: 'cmd/client/main.go',
        content: `package main

import (
  "context"
  "log"
  "log/slog"
  "os"
  "os/signal"
  "syscall"

  helloworldv1 "path/to/gen/helloworld/v1"
  "go.temporal.io/sdk/client"
  sdklog "go.temporal.io/sdk/log"
)

func main() {
  ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
  defer cancel()
  logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

  client, err := client.Dial(client.Options{
    Logger: sdklog.NewStructuredLogger(logger),
  })
  if err != nil {
    log.Fatal(err)
  }
  defer client.Close()

  example := helloworldv1.NewExampleClient(client)
  run, err := example.HelloAsync(ctx, &helloworldv1.HelloRequest{Name: "World"})
  if err != nil {
    log.Fatal(err)
  }
  logger = logger.With("workflow_id", run.ID())
  logger.Info("workflow started")

  _, ctx = <-ctx.Done(), context.Background()
  logger.Info("received shutdown signal, sending Goodbye signal to workflow")
  if err := run.Goodbye(ctx, &helloworldv1.GoodbyeRequest{}); err != nil {
    log.Fatal(err)
  }

  out, err := run.Get(ctx)
  if err != nil {
    log.Fatal(err)
  }
  logger.Info("workflow completed", "result", out.String())
}    
    `
      }
    ],
  },
  {
    label: 'CLI',
    content: 'Or from your local machine using the generated Command Line Interface.',
    fragments: [
      {
        language: 'sh',
        title: `print cli usage`,
        content: `go run main.go -h
NAME:
    example - A new cli application

USAGE:
    example [global options] command [command options] [arguments...]

COMMANDS:
    worker   runs a example.helloworld.v1.Example worker process
    help, h  Shows a list of commands or help for one command
    SIGNALS:
      goodbye  Goodbye signals a running workflow to exit
    WORKFLOWS:
      hello  Hello prints a friendly greeting and waits for goodbye

GLOBAL OPTIONS:
    --help, -h  show help
    `
      },
      {
        language: 'sh',
        title: `start a workflow`,
        content: `go run main.go hello --name Temporal -d
success
workflow id: hello/Temporal
run id: e55c6b09-7d05-418e-ad7e-8b40b9b3b867
    `
      },
      {
        language: 'sh',
        title: `send a signal`,
        content: `go run main.go goodbye --message 👋 -d
success
    `
      }
    ],
  },
  {
    label: 'XNS',
    content: 'Or from other Temporal workflows in a different Namespace or Cluster.',
    fragments: [
      {
        language: 'go',
        title: 'main.go',
        content: `package main

import (
  "time"

  helloworldv1 "path/to/gen/helloworld/v1"
  "path/to/gen/helloworld/v1/helloworldv1xns"
  "go.temporal.io/sdk/client"
  "go.temporal.io/sdk/worker"
  "go.temporal.io/sdk/workflow"
)

func SomeOtherWorkflow(ctx workflow.Context) error {
  run, err := helloworldv1xns.HelloAsync(ctx, &helloworldv1.HelloRequest{
    Name: workflow.GetInfo(ctx).WorkflowExecution.ID,
  })
  if err != nil {
    return err
  }

  workflow.Sleep(ctx, time.Second*30)
  if err := run.Goodbye(ctx, &helloworldv1.GoodbyeRequest{}); err != nil {
    return err
  }

  _, err = run.Get(ctx)
  return err
}

func main() {
  c, _ := client.Dial(client.Options{})
  defer c.Close()

  // initialize client for a different namespace/cluster
  xnsc, _ := client.NewClientFromExisting(c, client.Options{Namespace: "helloworld"})

  w := worker.New(c, "my-task-queue", worker.Options{})
  w.RegisterWorkflow(SomeOtherWorkflow)
  helloworldv1xns.RegisterExampleActivities(w, helloworldv1.NewExampleClient(xnsc))
  w.Run(w.InterruptCh())
}
        `
      }
    ],
  }
]

export default function HomepageFeatures(): JSX.Element {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          <div className="col col--6">
            <About></About>
          </div>
          <div className="col col--6">
            <Tabs>
              {examples.map((example, idx) => (
                <TabItem value={idx.toString()} label={example.label}>
                  <p>{example.content}</p>
                  {example.fragments.map((fragment, _) => (
                    <CodeBlock language={fragment.language} title={fragment.title}>{fragment.content}</CodeBlock>
                  ))}
                </TabItem>
              ))}
            </Tabs>
          </div>
        </div>
      </div>
    </section>
  );
}
