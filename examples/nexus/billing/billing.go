package billing

import (
	"time"

	nexusv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/nexus/v1"
	"go.temporal.io/sdk/workflow"
)

type (
	Workflows struct{}

	ChargeWorkflow struct {
		*nexusv1.ChargeWorkflowInput
	}
)

func (w *Workflows) Charge(ctx workflow.Context, input *nexusv1.ChargeWorkflowInput) (nexusv1.ChargeWorkflow, error) {
	return &ChargeWorkflow{input}, nil
}

func (w *ChargeWorkflow) Execute(ctx workflow.Context) (*nexusv1.ChargeOutput, error) {
	if err := workflow.Sleep(ctx, time.Second*5); err != nil {
		return nil, err
	}
	order := w.Req.GetOrder()
	order.Status = nexusv1.OrderStatus_ORDER_STATUS_IN_TRANSIT
	return &nexusv1.ChargeOutput{Order: order}, nil
}
