
export let buf_gen_yaml = `version: v2
managed:
  enabled: true
  override:
    - file_option: go_package_prefix
      value: example/gen
plugins:
  - local: protoc-gen-go
    out: gen
    opt:
      - paths=source_relative
  - local: protoc-gen-go_temporal
    out: gen
    strategy: all
    opt:
      - cli-categories=true
      - cli-enabled=true
      - cli-v3=true
      - docs-out=./proto/README.md
`

export let buf_yaml = `version: v2
modules:
  - path: proto
deps:
  - buf.build/cludden/protoc-gen-go-temporal
lint:
  use:
    - BASIC
`

export let cli_fragments = [
  {
    language: 'sh',
    title: 'build cli binary',
    output: '',
    content: 'go build -o example cmd/example/main.go',
  },
  {
    language: 'sh',
    title: 'print cli usage details',
    output: 'img/cli-usage.png',
    content: `example -h`,
  },
  {
    language: 'sh',
    title: 'start a workflow',
    output: 'img/cli-start-workflow.png',
    content: `example create-foo --name test -d`,
  },
  {
    language: 'sh',
    title: `send a signal`,
    output: '',
    content: `example set-foo-progress -w create-foo/test --progress 5.7`,
  },
  {
    language: 'sh',
    title: 'query workflow',
    output: 'img/cli-query.png',
    content: 'example get-foo-progress -w create-foo/test',
  },
  {
    language: 'sh',
    title: 'update workflow',
    output: 'img/cli-update.png',
    content: 'example update-foo-progress -w create-foo/test --progress 100',
  },
];

export let xns_example = `package main

import (
	"fmt"
	"log"

	examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	"github.com/cludden/protoc-gen-go-temporal/gen/example/v1/examplev1xns"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func main() {
	// initialize temporal client for current namespace
	c, _ := client.Dial(client.Options{
		Namespace: "default",
	})
	defer c.Close()

	// initialize temporal worker in current namespace
	w := worker.New(c, "my-task-queue", worker.Options{})
	w.RegisterWorkflow(SomeWorkflow)

	// initialize temporal client for proto service namespace
	xnsc, _ := client.NewClientFromExisting(c, client.Options{
		Namespace: "example",
	})

	// register generated cross-namespace activities using the appropriate
	// temporal client
	examplev1xns.RegisterExampleActivities(w, examplev1.NewExampleClient(xnsc))
	
	// start worker
	_ = w.Run(w.InterruptCh())
}

func SomeWorkflow(ctx workflow.Context) error {
	log := workflow.GetLogger(ctx)

	// start workflow in target namespace
	run, err := examplev1xns.CreateFooAsync(ctx, &examplev1.CreateFooRequest{Name: w.Req.GetName()})
	if err != nil {
		return fmt.Errorf("error initializing CreateFoo workflow: %w", err)
	}

	// send signal
	if err := run.SetFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 5.7}); err != nil {
		return fmt.Errorf("error signaling SetFooProgress: %w", err)
	}
	log.Info("SetFooProgress", "progress", 5.7)

	// execute query
	progress, err := run.GetFooProgress(ctx)
	if err != nil {
		return fmt.Errorf("error querying GetFooProgress: %w", err)
	}
	log.Info("GetFooProgress", "status", progress.GetStatus().String(), "progress", progress.GetProgress())

	// execute update
	update, err := run.UpdateFooProgressAsync(ctx, &examplev1.SetFooProgressRequest{Progress: 100})
	if err != nil {
		return fmt.Errorf("error initializing UpdateFooProgress: %w", err)
	}
	progress, err = update.Get(ctx)
	if err != nil {
		return fmt.Errorf("error updating UpdateFooProgress: %w", err)
	}
	log.Info("UpdateFooProgress", "status", progress.GetStatus().String(), "progress", progress.GetProgress())

	// await workflow completion
	resp, err := run.Get(ctx)
	if err != nil {
		return err
	}
	return nil
}
`

