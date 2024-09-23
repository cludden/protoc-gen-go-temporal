package shipping

import (
	"time"

	nexusv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/nexus/v1"
	"go.temporal.io/sdk/workflow"
)

type (
	Workflows struct{}

	ShipmentWorkflow struct {
		*nexusv1.ShipmentWorkflowInput
	}
)

func (w *Workflows) Shipment(ctx workflow.Context, input *nexusv1.ShipmentWorkflowInput) (nexusv1.ShipmentWorkflow, error) {
	return &ShipmentWorkflow{input}, nil
}

func (w *ShipmentWorkflow) Execute(ctx workflow.Context) error {
	return workflow.Sleep(ctx, time.Second*10)
}
