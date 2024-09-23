package orders

import (
	"context"
	"testing"

	"github.com/cludden/protoc-gen-go-temporal/examples/nexus/billing"
	"github.com/cludden/protoc-gen-go-temporal/examples/nexus/shipping"
	nexusv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/nexus/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/testing/protocmp"
)

type OrdersSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	env    *testsuite.TestWorkflowEnvironment
	orders nexusv1.OrdersClient
}

func TestOrders(t *testing.T) {
	suite.Run(t, new(OrdersSuite))
}

func (s *OrdersSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()

	nexusv1.RegisterBillingWorkflows(s.env, &billing.Workflows{})
	s.Require().NoError(nexusv1.RegisterBillingService(s.env))
	nexusv1.RegisterShippingWorkflows(s.env, &shipping.Workflows{})
	s.Require().NoError(nexusv1.RegisterShippingService(s.env))

	s.orders = nexusv1.NewTestOrdersClient(s.env, &Workflows{
		BillingEndpoint:  "billing-test",
		ShippingEndpoint: "shipping-test",
	}, nil)
}

func (s *OrdersSuite) SetupSubTest() {
	s.SetupTest()
}

func genCreateOrderInput(fns ...func(*nexusv1.CreateOrderInput)) *nexusv1.CreateOrderInput {
	v := &nexusv1.CreateOrderInput{
		CustomerId: "abc",
		Items: []*nexusv1.Item{
			{
				Sku:      "0123",
				Quantity: 1,
			},
			{
				Sku:      "4567",
				Quantity: 2,
			},
		},
	}
	for _, fn := range fns {
		fn(v)
	}
	return v
}

func (s *OrdersSuite) TestCreateOrder() {
	cases := []struct {
		desc          string
		input         *nexusv1.CreateOrderInput
		errors        []string
		chargeError   error
		shipmentError error
	}{
		{
			desc: "success",
		},
	}

	for _, c := range cases {
		s.Run(c.desc, func() {
			r := s.Require()

			input := c.input
			if input == nil {
				input = genCreateOrderInput()
			}

			setupMocks := func() {
				s.env.
					OnWorkflow(nexusv1.ChargeWorkflowName, mock.Anything, mock.MatchedBy(func(req *nexusv1.ChargeInput) bool {
						c1 := s.Equal(input.GetCustomerId(), req.GetOrder().GetCustomerId())
						c2 := s.Empty(cmp.Diff(input.GetItems(), req.GetOrder().GetItems(), protocmp.Transform()))
						c3 := s.Equal(nexusv1.OrderStatus_ORDER_STATUS_PENDING.String(), req.GetOrder().GetStatus().String())
						return c1 && c2 && c3
					})).
					Return(func(ctx workflow.Context, req *nexusv1.ChargeInput) (*nexusv1.ChargeOutput, error) {
						if c.chargeError != nil {
							return nil, c.chargeError
						}
						order := req.GetOrder()
						order.Status = nexusv1.OrderStatus_ORDER_STATUS_IN_TRANSIT
						return &nexusv1.ChargeOutput{Order: order}, nil
					})
				if c.chargeError != nil {
					return
				}

				s.env.
					OnWorkflow(nexusv1.ShipmentWorkflowName, mock.Anything, mock.MatchedBy(func(req *nexusv1.ShipmentInput) bool {
						return s.Equal(input.GetCustomerId(), req.GetOrder().GetCustomerId()) &&
							s.Empty(cmp.Diff(input.GetItems(), req.GetOrder().GetItems(), protocmp.Transform())) &&
							s.Equal(nexusv1.OrderStatus_ORDER_STATUS_IN_TRANSIT.String(), req.GetOrder().GetStatus().String())
					})).
					Return(func(ctx workflow.Context, req *nexusv1.ShipmentInput) error {
						if c.shipmentError != nil {
							return c.shipmentError
						}
						return nil
					})
			}
			setupMocks()

			out, err := s.orders.CreateOrder(context.Background(), input)
			if len(c.errors) > 0 {
				r.Error(err)
				for _, msg := range c.errors {
					r.ErrorContains(err, msg)
				}
			} else {
				r.NoError(err)
				r.NotNil(out)
				r.Equal(nexusv1.OrderStatus_ORDER_STATUS_COMPLETED, out.GetOrder().GetStatus())
			}
		})
	}
}
