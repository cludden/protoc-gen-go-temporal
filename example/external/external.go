package external

import (
	"fmt"

	examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	"github.com/cludden/protoc-gen-go-temporal/gen/example/v1/examplev1xns"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/workflow"
)

type (
	Activities struct{}

	Workflows struct{}
)

type ProvisionFooWorkflow struct {
	*examplev1.ProvisionFooWorkflowInput
	log log.Logger
}

func (wfs *Workflows) ProvisionFoo(ctx workflow.Context, input *examplev1.ProvisionFooWorkflowInput) (examplev1.ProvisionFooWorkflow, error) {
	return &ProvisionFooWorkflow{input, workflow.GetLogger(ctx)}, nil
}

func (w *ProvisionFooWorkflow) Execute(ctx workflow.Context) (*examplev1.ProvisionFooResponse, error) {
	run, err := examplev1xns.CreateFooAsync(ctx, &examplev1.CreateFooInput{RequestName: w.Req.GetRequestName()})
	if err != nil {
		return nil, fmt.Errorf("error initializing CreateFoo workflow: %w", err)
	}

	if err := run.SetFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 5.7}); err != nil {
		return nil, fmt.Errorf("error signaling SetFooProgress: %w", err)
	}

	progress, err := run.GetFooProgress(ctx)
	if err != nil {
		return nil, fmt.Errorf("error querying GetFooProgress: %w", err)
	}
	w.log.Info("GetFooProgress", "status", progress.GetStatus().String(), "progress", progress.GetProgress())

	update, err := run.UpdateFooProgressAsync(ctx, &examplev1.SetFooProgressRequest{Progress: 100})
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
	return &examplev1.ProvisionFooResponse{Foo: resp.GetFoo()}, nil
}
