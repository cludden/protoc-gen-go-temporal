package shoppingcart

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"

	shoppingcartv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/shoppingcart/v1"
	"github.com/cludden/protoc-gen-go-temporal/gen/example/shoppingcart/v1/shoppingcartv1xns"
	"github.com/google/uuid"
	"github.com/hairyhenderson/go-which"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/types/known/durationpb"
)

type ShoppingCartSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	env    *testsuite.TestWorkflowEnvironment
	client shoppingcartv1.ShoppingCartClient
	ctx    context.Context
}

func TestShoppingCartSuite(t *testing.T) {
	suite.Run(t, new(ShoppingCartSuite))
}

func (s *ShoppingCartSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
	s.client = shoppingcartv1.NewTestShoppingCartClient(s.env, &Workflows{}, nil)
	s.ctx = context.Background()
}

func (s *ShoppingCartSuite) TestShoppingCartWithUpdateCart() {
	handle, run, err := s.client.ShoppingCartWithUpdateCartAsync(
		s.ctx,
		&shoppingcartv1.ShoppingCartInput{},
		&shoppingcartv1.UpdateCartInput{
			Action: shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_ADD,
			ItemId: "foo",
		},
	)
	s.Require().NoError(err)

	var addFoo shoppingcartv1.UpdateCartHandle
	s.env.RegisterDelayedCallback(func() {
		var err error
		addFoo, err = run.UpdateCartAsync(s.ctx, &shoppingcartv1.UpdateCartInput{
			Action: shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_ADD,
			ItemId: "foo",
		})
		s.Require().NoError(err)
	}, time.Second)

	var removeFoo shoppingcartv1.UpdateCartHandle
	s.env.RegisterDelayedCallback(func() {
		var err error
		removeFoo, err = run.UpdateCartAsync(s.ctx, &shoppingcartv1.UpdateCartInput{
			Action: shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_REMOVE,
			ItemId: "foo",
		})
		s.Require().NoError(err)
	}, time.Second*2)

	s.env.RegisterDelayedCallback(func() {
		s.Require().NoError(
			run.Checkout(s.ctx, &shoppingcartv1.CheckoutInput{}),
		)
	}, time.Second*3)

	out, err := run.Get(s.ctx)
	s.Require().NoError(err)
	s.Require().NotNil(out)
	s.Require().Contains(out.GetCart().GetItems(), "foo")
	s.Require().Equal(int32(1), out.GetCart().GetItems()["foo"])

	resp, err := handle.Get(s.ctx)
	s.Require().NoError(err)
	s.Require().NotNil(resp)
	s.Require().Contains(resp.GetCart().GetItems(), "foo")
	s.Require().Equal(int32(1), resp.GetCart().GetItems()["foo"])

	resp, err = addFoo.Get(s.ctx)
	s.Require().NoError(err)
	s.Require().NotNil(resp)
	s.Require().Contains(resp.GetCart().GetItems(), "foo")
	s.Require().Equal(int32(2), resp.GetCart().GetItems()["foo"])

	resp, err = removeFoo.Get(s.ctx)
	s.Require().NoError(err)
	s.Require().NotNil(resp)
	s.Require().Contains(resp.GetCart().GetItems(), "foo")
	s.Require().Equal(int32(1), resp.GetCart().GetItems()["foo"])
}

func (s *ShoppingCartSuite) TestShoppingCartWithUpdateCartAndContinueAsNew() {
	s.env.SetContinueAsNewSuggested(true)

	var updatesCompleted int
	s.env.RegisterDelayedCallback(func() {
		s.env.UpdateWorkflow(
			shoppingcartv1.UpdateCartUpdateName,
			uuid.NewString(),
			&testsuite.TestUpdateCallback{
				OnAccept: func() {},
				OnReject: func(err error) { s.Require().Fail("unexpected rejection") },
				OnComplete: func(i interface{}, err error) {
					s.Require().NoError(err)
					cartState, ok := i.(*shoppingcartv1.UpdateCartOutput)
					if !ok {
						s.Require().Fail("Invalid return type")
					}
					s.Require().Equal(int32(1), cartState.GetCart().GetItems()["foo"])
					updatesCompleted++
				},
			},
			&shoppingcartv1.UpdateCartInput{
				Action: shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_ADD,
				ItemId: "foo",
			},
		)
	}, 0)

	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(shoppingcartv1.CheckoutSignalName, &shoppingcartv1.CheckoutInput{})
	}, time.Second)

	s.env.ExecuteWorkflow(
		shoppingcartv1.ShoppingCartWorkflowName,
		&shoppingcartv1.ShoppingCartInput{},
	)

	s.Require().True(s.env.IsWorkflowCompleted())
	err := s.env.GetWorkflowError()
	s.Require().
		True(workflow.IsContinueAsNewError(err), "Expected workflow to continue as new, got: %v", err)
}

func (s *ShoppingCartSuite) TestShoppingCartWithUpdateCartSynchronous() {
	s.env.RegisterDelayedCallback(func() {
		err := s.client.Checkout(s.ctx, "", "", &shoppingcartv1.CheckoutInput{})
		s.Require().NoError(err)
	}, time.Minute)

	update, run, err := s.client.ShoppingCartWithUpdateCart(
		s.ctx,
		&shoppingcartv1.ShoppingCartInput{},
		&shoppingcartv1.UpdateCartInput{
			Action: shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_ADD,
			ItemId: "foo",
		},
	)
	s.Require().NoError(err)
	s.Require().NotNil(update)
	s.Require().Contains(update.GetCart().GetItems(), "foo")
	s.Require().Equal(int32(1), update.GetCart().GetItems()["foo"])

	resp, err := run.Get(s.ctx)
	s.Require().NoError(err)
	s.Require().NotNil(resp)
	s.Require().Contains(resp.GetCart().GetItems(), "foo")
	s.Require().Equal(int32(1), resp.GetCart().GetItems()["foo"])
}

func TestShoppingCartWithUpdateCartXnsE2E(t *testing.T) {
	existingPath := which.Which("temporal")
	if existingPath == "" {
		t.Skip("temporal CLI not found in PATH, skipping E2E test")
	}
	ctx, r := context.Background(), require.New(t)

	logger := log.NewStructuredLogger(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	srv, err := testsuite.StartDevServer(ctx, testsuite.DevServerOptions{
		ExistingPath: existingPath,
		EnableUI:     true,
		ClientOptions: &client.Options{
			Logger: logger,
		},
	})
	r.NoError(err)
	t.Cleanup(func() { r.NoError(srv.Stop()) })

	dc := srv.Client()
	t.Cleanup(dc.Close)

	_, _ = dc.WorkflowService().RegisterNamespace(ctx, &workflowservice.RegisterNamespaceRequest{
		Namespace:                        "shoppingcart",
		WorkflowExecutionRetentionPeriod: durationpb.New(time.Hour * 24),
	})

	sc, err := client.NewClientFromExisting(dc, client.Options{
		Namespace: "shoppingcart",
		Logger:    logger,
	})
	r.NoError(err)
	t.Cleanup(sc.Close)

	sw := worker.New(sc, shoppingcartv1.ShoppingCartTaskQueue, worker.Options{})
	shoppingcartv1.RegisterShoppingCartWorkflows(sw, &Workflows{})
	r.NoError(sw.Start())
	t.Cleanup(sw.Stop)

	dw := worker.New(dc, "test", worker.Options{})
	shoppingcartv1xns.RegisterShoppingCartActivities(dw, shoppingcartv1.NewShoppingCartClient(sc))

	dw.RegisterWorkflowWithOptions(
		func(ctx workflow.Context, input *shoppingcartv1.ShoppingCartInput) (*shoppingcartv1.ShoppingCartOutput, error) {
			uwsopts := shoppingcartv1xns.NewShoppingCartWithUpdateCartOptions().
				WithWorkflowOptions(
					shoppingcartv1xns.NewShoppingCartWorkflowOptions().
						WithStartWorkflow(client.StartWorkflowOptions{
							ID: "shoppingcart-test",
						}),
				)

			update, run, err := shoppingcartv1xns.ShoppingCartWithUpdateCart(
				ctx,
				input,
				&shoppingcartv1.UpdateCartInput{
					Action: shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_ADD,
					ItemId: "foo",
				},
				uwsopts,
			)
			if err != nil {
				return nil, err
			} else if c := update.GetCart().GetItems()["foo"]; c != 1 {
				return nil, fmt.Errorf("expected cart to have 1 foo item, got %v", c)
			}

			update2, _, err := shoppingcartv1xns.ShoppingCartWithUpdateCart(
				ctx,
				input,
				&shoppingcartv1.UpdateCartInput{
					Action: shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_ADD,
					ItemId: "foo",
				},
				uwsopts,
			)
			if err != nil {
				return nil, err
			} else if c := update2.GetCart().GetItems()["foo"]; c != 2 {
				return nil, fmt.Errorf("expected cart to have 2 foo items, got %v", c)
			}

			update3, _, err := shoppingcartv1xns.ShoppingCartWithUpdateCart(
				ctx,
				input,
				&shoppingcartv1.UpdateCartInput{
					Action: shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_REMOVE,
					ItemId: "foo",
				},
				uwsopts,
			)
			if err != nil {
				return nil, err
			} else if c := update3.GetCart().GetItems()["foo"]; c != 1 {
				return nil, fmt.Errorf("expected cart to have 1 foo item, got %v", c)
			}

			workflow.GetSignalChannel(ctx, "checkout").Receive(ctx, nil)
			if err := run.Checkout(ctx, &shoppingcartv1.CheckoutInput{}); err != nil {
				return nil, err
			}

			out, err := run.Get(ctx)
			if err != nil {
				return nil, err
			} else if c := out.GetCart().GetItems()["foo"]; c != 1 {
				return nil, fmt.Errorf("expected cart to have 1 foo item after checkout, got %v", c)
			}
			return out, nil
		},
		workflow.RegisterOptions{
			Name: "test",
		},
	)

	r.NoError(dw.Start())
	t.Cleanup(dw.Stop)

	run, err := dc.ExecuteWorkflow(
		ctx,
		client.StartWorkflowOptions{
			ID:        uuid.NewString(),
			TaskQueue: "test",
		},
		"test",
		&shoppingcartv1.ShoppingCartInput{},
	)
	r.NoError(err)

	time.Sleep(time.Second)
	r.NoError(dc.SignalWorkflow(ctx, run.GetID(), run.GetRunID(), "checkout", nil))

	var out shoppingcartv1.ShoppingCartOutput
	r.NoError(run.Get(ctx, &out))
	r.Contains(out.GetCart().GetItems(), "foo")
	r.Equal(int32(1), out.GetCart().GetItems()["foo"])
}

func TestShoppingCartWithUpdateCartXnsWithCancellationE2E(t *testing.T) {
	existingPath := which.Which("temporal")
	if existingPath == "" {
		t.Skip("temporal CLI not found in PATH, skipping E2E test")
	}
	ctx, r := context.Background(), require.New(t)

	logger := log.NewStructuredLogger(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	srv, err := testsuite.StartDevServer(ctx, testsuite.DevServerOptions{
		ExistingPath: existingPath,
		EnableUI:     true,
		ClientOptions: &client.Options{
			Logger: logger,
		},
	})
	r.NoError(err)
	t.Cleanup(func() { r.NoError(srv.Stop()) })

	dc := srv.Client()
	t.Cleanup(dc.Close)

	_, _ = dc.WorkflowService().RegisterNamespace(ctx, &workflowservice.RegisterNamespaceRequest{
		Namespace:                        "shoppingcart",
		WorkflowExecutionRetentionPeriod: durationpb.New(time.Hour * 24),
	})

	sc, err := client.NewClientFromExisting(dc, client.Options{
		Namespace: "shoppingcart",
		Logger:    logger,
	})
	r.NoError(err)

	sw := worker.New(sc, shoppingcartv1.ShoppingCartTaskQueue, worker.Options{})
	shoppingcartv1.RegisterShoppingCartWorkflows(sw, &Workflows{})
	r.NoError(sw.Start())
	t.Cleanup(sw.Stop)

	dw := worker.New(dc, "test", worker.Options{
		MaxHeartbeatThrottleInterval:     0,
		DefaultHeartbeatThrottleInterval: 0,
	})
	shoppingcartv1xns.RegisterShoppingCartActivities(dw, shoppingcartv1.NewShoppingCartClient(sc))

	dw.RegisterWorkflowWithOptions(
		func(ctx workflow.Context, input *shoppingcartv1.ShoppingCartInput) (out *shoppingcartv1.ShoppingCartOutput, err error) {
			uwsopts := shoppingcartv1xns.NewShoppingCartWithUpdateCartOptions().
				WithWorkflowOptions(
					shoppingcartv1xns.NewShoppingCartWorkflowOptions().
						WithStartWorkflow(client.StartWorkflowOptions{
							ID: "shoppingcart-test",
						}),
				)

			update, run, err := shoppingcartv1xns.ShoppingCartWithUpdateCart(
				ctx,
				input,
				&shoppingcartv1.UpdateCartInput{
					Action: shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_ADD,
					ItemId: "foo",
				},
				uwsopts,
			)
			if err != nil {
				return nil, err
			} else if c := update.GetCart().GetItems()["foo"]; c != 1 {
				return nil, fmt.Errorf("expected cart to have 1 foo item, got %v", c)
			}

			update2, _, err := shoppingcartv1xns.ShoppingCartWithUpdateCart(
				ctx,
				input,
				&shoppingcartv1.UpdateCartInput{
					Action: shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_ADD,
					ItemId: "foo",
				},
				uwsopts,
			)
			if err != nil {
				return nil, err
			} else if c := update2.GetCart().GetItems()["foo"]; c != 2 {
				return nil, fmt.Errorf("expected cart to have 2 foo items, got %v", c)
			}

			update3, _, err := shoppingcartv1xns.ShoppingCartWithUpdateCart(
				ctx,
				input,
				&shoppingcartv1.UpdateCartInput{
					Action: shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_REMOVE,
					ItemId: "foo",
				},
				uwsopts,
			)
			if err != nil {
				return nil, err
			} else if c := update3.GetCart().GetItems()["foo"]; c != 1 {
				return nil, fmt.Errorf("expected cart to have 1 foo item, got %v", c)
			}

			workflow.Go(ctx, func(ctx workflow.Context) {
				workflow.GetSignalChannel(ctx, "checkout").Receive(ctx, nil)
				if err := run.Checkout(ctx, &shoppingcartv1.CheckoutInput{}); err != nil {
					workflow.GetLogger(ctx).Error("checkout failed", "error", err)
				}
			})

			return run.Get(ctx)
		},
		workflow.RegisterOptions{
			Name: "test",
		},
	)

	r.NoError(dw.Start())
	t.Cleanup(dw.Stop)

	run, err := dc.ExecuteWorkflow(
		ctx,
		client.StartWorkflowOptions{
			ID:        uuid.NewString(),
			TaskQueue: "test",
		},
		"test",
		&shoppingcartv1.ShoppingCartInput{},
	)
	r.NoError(err)

	time.Sleep(time.Second * 2)
	r.NoError(dc.CancelWorkflow(ctx, run.GetID(), run.GetRunID()))

	r.True(temporal.IsCanceledError(run.Get(ctx, nil)))

	r.Eventually(func() bool {
		resp, err := sc.DescribeWorkflowExecution(ctx, "shoppingcart-test", "")
		if err != nil {
			return false
		}
		return resp.WorkflowExecutionInfo.Status == enums.WORKFLOW_EXECUTION_STATUS_CANCELED
	}, time.Second*10, time.Millisecond*100, "Expected xns workflow to be canceled")
}
