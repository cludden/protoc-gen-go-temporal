package test

import (
	"context"
	"testing"
	"time"

	"github.com/cludden/protoc-gen-go-temporal/examples/nexus/billing"
	"github.com/cludden/protoc-gen-go-temporal/examples/nexus/orders"
	"github.com/cludden/protoc-gen-go-temporal/examples/nexus/shipping"
	nexusv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/nexus/v1"
	"github.com/stretchr/testify/require"
	"go.temporal.io/api/nexus/v1"
	"go.temporal.io/api/operatorservice/v1"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
	"google.golang.org/protobuf/types/known/durationpb"
)

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	r, ctx := require.New(t), context.Background()

	// start dev server
	srv, err := testsuite.StartDevServer(ctx, testsuite.DevServerOptions{
		ClientOptions: &client.Options{
			HostPort: "0.0.0.0:7233",
		},
		EnableUI: true,
		ExtraArgs: []string{
			"--http-port", "7243",
			"--dynamic-config-value", "system.enableNexus=true",
		},
	})
	r.NoError(err)
	t.Cleanup(func() { r.NoError(srv.Stop()) })

	// defer client cleanup
	c := srv.Client()
	t.Cleanup(c.Close)

	// create namespaces
	for _, ns := range []string{"orders", "billing", "shipping"} {
		_, err := c.WorkflowService().RegisterNamespace(ctx, &workflowservice.RegisterNamespaceRequest{
			Namespace:                        ns,
			WorkflowExecutionRetentionPeriod: durationpb.New(time.Hour * 24),
		})
		r.NoError(err)
	}

	// create nexus endpoints
	_, err = c.OperatorService().CreateNexusEndpoint(ctx, &operatorservice.CreateNexusEndpointRequest{
		Spec: &nexus.EndpointSpec{
			Name: "billing-test",
			Target: &nexus.EndpointTarget{
				Variant: &nexus.EndpointTarget_Worker_{
					Worker: &nexus.EndpointTarget_Worker{
						Namespace: "billing",
						TaskQueue: "billing-v1",
					},
				},
			},
		},
	})
	r.NoError(err)

	_, err = c.OperatorService().CreateNexusEndpoint(ctx, &operatorservice.CreateNexusEndpointRequest{
		Spec: &nexus.EndpointSpec{
			Name: "shipping-test",
			Target: &nexus.EndpointTarget{
				Variant: &nexus.EndpointTarget_Worker_{
					Worker: &nexus.EndpointTarget_Worker{
						Namespace: "shipping",
						TaskQueue: "shipping-v1",
					},
				},
			},
		},
	})
	r.NoError(err)

	// initialize orders worker
	ordersClient, err := client.NewClientFromExistingWithContext(ctx, c, client.Options{Namespace: "orders"})
	r.NoError(err)

	ordersWorker := worker.New(ordersClient, nexusv1.OrdersTaskQueue, worker.Options{})
	nexusv1.RegisterOrdersWorkflows(ordersWorker, &orders.Workflows{
		BillingEndpoint:  "billing-test",
		ShippingEndpoint: "shipping-test",
	})
	r.NoError(ordersWorker.Start())
	t.Cleanup(ordersWorker.Stop)

	// initialize billing worker
	billingClient, err := client.NewClientFromExistingWithContext(ctx, c, client.Options{Namespace: "billing"})
	r.NoError(err)

	billingWorker := worker.New(billingClient, nexusv1.BillingTaskQueue, worker.Options{})
	nexusv1.RegisterBillingWorkflows(billingWorker, &billing.Workflows{})
	r.NoError(nexusv1.RegisterBillingService(billingWorker))
	r.NoError(billingWorker.Start())
	t.Cleanup(billingWorker.Stop)

	// initialize shipping worker
	shippingClient, err := client.NewClientFromExistingWithContext(ctx, c, client.Options{Namespace: "shipping"})
	r.NoError(err)

	shippingWorker := worker.New(shippingClient, nexusv1.ShippingTaskQueue, worker.Options{})
	nexusv1.RegisterShippingWorkflows(shippingWorker, &shipping.Workflows{})
	r.NoError(nexusv1.RegisterShippingService(shippingWorker))
	r.NoError(shippingWorker.Start())
	t.Cleanup(shippingWorker.Stop)

	// create order
	order, err := nexusv1.NewOrdersClient(ordersClient).CreateOrder(ctx, &nexusv1.CreateOrderInput{
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
	})
	r.NoError(err)
	r.NotNil(order)
}
