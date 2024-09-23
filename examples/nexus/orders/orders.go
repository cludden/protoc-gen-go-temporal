package orders

import (
	nexusv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/nexus/v1"
	"github.com/google/uuid"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type (
	Workflows struct {
		BillingEndpoint  string
		ShippingEndpoint string
	}

	CreateOrderWorkflow struct {
		*Workflows
		*nexusv1.CreateOrderWorkflowInput
		billing  nexusv1.BillingNexusClient
		shipping nexusv1.ShippingNexusClient
	}
)

func (w *Workflows) CreateOrder(ctx workflow.Context, input *nexusv1.CreateOrderWorkflowInput) (nexusv1.CreateOrderWorkflow, error) {
	return &CreateOrderWorkflow{w, input, nexusv1.NewBillingNexusClient(w.BillingEndpoint), nexusv1.NewShippingNexusClient(w.ShippingEndpoint)}, nil
}

func (w *CreateOrderWorkflow) Execute(ctx workflow.Context) (*nexusv1.CreateOrderOutput, error) {
	var id string
	if err := workflow.SideEffect(ctx, func(ctx workflow.Context) any {
		return uuid.NewString()
	}).Get(&id); err != nil {
		return nil, err
	}

	order := &nexusv1.Order{
		Id:         id,
		CustomerId: w.Req.GetCustomerId(),
		Items:      w.Req.GetItems(),
		ReceivedAt: timestamppb.New(workflow.Now(ctx)),
	}

	charge, err := w.billing.Charge(ctx, &nexusv1.ChargeInput{Order: order})
	if err != nil {
		return nil, err
	}
	order = charge.GetOrder()

	if err := w.shipping.Shipment(ctx, &nexusv1.ShipmentInput{Order: order}); err != nil {
		return nil, err
	}
	order.Status = nexusv1.OrderStatus_ORDER_STATUS_COMPLETED

	return &nexusv1.CreateOrderOutput{Order: order}, nil
}
