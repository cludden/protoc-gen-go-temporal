package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	xnsv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/xns/v1"
	"github.com/cludden/protoc-gen-go-temporal/gen/example/xns/v1/xnsv1xns"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

type (
	Workflows struct{}
)

type ProvisionFooWorkflow struct {
	*xnsv1.ProvisionFooWorkflowInput
	log tlog.Logger
}

func (wfs *Workflows) ProvisionFoo(ctx workflow.Context, input *xnsv1.ProvisionFooWorkflowInput) (xnsv1.ProvisionFooWorkflow, error) {
	return &ProvisionFooWorkflow{input, workflow.GetLogger(ctx)}, nil
}

func (w *ProvisionFooWorkflow) Execute(ctx workflow.Context) (*xnsv1.ProvisionFooResponse, error) {
	run, err := xnsv1xns.CreateFooAsync(ctx, &xnsv1.CreateFooRequest{Name: w.Req.GetName()})
	if err != nil {
		return nil, fmt.Errorf("error initializing CreateFoo workflow: %w", err)
	}

	if err := run.SetFooProgress(ctx, &xnsv1.SetFooProgressRequest{Progress: 5.7}); err != nil {
		return nil, fmt.Errorf("error signaling SetFooProgress: %w", err)
	}
	w.log.Info("SetFooProgress", "progress", 5.7)

	progress, err := run.GetFooProgress(ctx)
	if err != nil {
		return nil, fmt.Errorf("error querying GetFooProgress: %w", err)
	}
	w.log.Info("GetFooProgress", "status", progress.GetStatus().String(), "progress", progress.GetProgress())

	update, err := run.UpdateFooProgressAsync(ctx, &xnsv1.SetFooProgressRequest{Progress: 100})
	if err != nil {
		return nil, fmt.Errorf("error initializing UpdateFooProgress: %w", err)
	}
	progress, err = update.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("error updating UpdateFooProgress: %w", err)
	}
	w.log.Info("UpdateFooProgress", "status", progress.GetStatus().String(), "progress", progress.GetProgress())

	resp, err := run.Get(ctx)
	if err != nil {
		return nil, err
	}
	return &xnsv1.ProvisionFooResponse{Foo: resp.GetFoo()}, nil
}

func main() {
	app := &cli.App{}

	exampleCmd, err := xnsv1.NewExampleCliCommand()
	if err != nil {
		log.Fatal(err)
	}
	app.Commands = append(app.Commands, exampleCmd)

	xnsCmd, err := xnsv1.NewXnsCliCommand()
	if err != nil {
		log.Fatal(err)
	}
	app.Commands = append(app.Commands, xnsCmd)

	app.Commands = append(app.Commands, &cli.Command{
		Name: "worker",
		Action: func(cmd *cli.Context) error {
			c, err := client.Dial(client.Options{
				Namespace: "example",
			})
			if err != nil {
				return err
			}
			defer c.Close()

			xnsc, err := client.NewClientFromExisting(c, client.Options{
				Namespace: "default",
			})
			if err != nil {
				return err
			}

			examplew := worker.New(c, xnsv1.ExampleTaskQueue, worker.Options{})
			xnsv1.RegisterExampleWorkflows(examplew, &ExampleWorkflows{})
			xnsv1.RegisterExampleActivities(examplew, &ExampleActivities{})

			xnsw := worker.New(xnsc, xnsv1.XnsTaskQueue, worker.Options{})
			xnsv1.RegisterXnsWorkflows(xnsw, &Workflows{})
			xnsv1xns.RegisterExampleActivities(xnsw, xnsv1.NewExampleClient(c))

			var g sync.WaitGroup
			closeCh := make(chan any)
			g.Add(2)
			go func() {
				defer g.Done()
				examplew.Run(closeCh)
			}()
			go func() {
				defer g.Done()
				xnsw.Run(closeCh)
			}()

			interruptCh := make(chan os.Signal, 1)
			signal.Notify(interruptCh, syscall.SIGINT, syscall.SIGTERM)
			go func() {
				<-interruptCh
				close(closeCh)
			}()
			g.Wait()
			return nil
		},
	})

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
