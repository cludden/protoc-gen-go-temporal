
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

export let xns_example = `package xns

import (
    "fmt"

    examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
    "github.com/cludden/protoc-gen-go-temporal/gen/example/v1/examplev1xns"
    "go.temporal.io/sdk/workflow"
)

func SomeWorkflow(ctx workflow.Context) error {
    log := workflow.GetLogger(ctx)

    run, err := examplev1xns.CreateFooAsync(ctx, &examplev1.CreateFooRequest{Name: w.Req.GetName()})
    if err != nil {
        return fmt.Errorf("error initializing CreateFoo workflow: %w", err)
    }

    if err := run.SetFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 5.7}); err != nil {
        return fmt.Errorf("error signaling SetFooProgress: %w", err)
    }
    log.Info("SetFooProgress", "progress", 5.7)

    progress, err := run.GetFooProgress(ctx)
    if err != nil {
        return fmt.Errorf("error querying GetFooProgress: %w", err)
    }
    log.Info("GetFooProgress", "status", progress.GetStatus().String(), "progress", progress.GetProgress())

    update, err := run.UpdateFooProgressAsync(ctx, &examplev1.SetFooProgressRequest{Progress: 100})
    if err != nil {
        return fmt.Errorf("error initializing UpdateFooProgress: %w", err)
    }
    progress, err = update.Get(ctx)
    if err != nil {
        return fmt.Errorf("error updating UpdateFooProgress: %w", err)
    }
    log.Info("UpdateFooProgress", "status", progress.GetStatus().String(), "progress", progress.GetProgress())

    resp, err := run.Get(ctx)
    if err != nil {
        return err
    }
    return nil
}
`
