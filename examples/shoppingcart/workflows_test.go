package shoppingcart

import (
	"context"
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
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/proto"
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

	s.env.RegisterDelayedCallback(func() {
		s.Require().NoError(
			run.Checkout(s.ctx, &shoppingcartv1.CheckoutInput{}),
		)
	}, time.Second)

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
}

func (s *ShoppingCartSuite) TestShoppingCartWithUpdateCartAndContinueAsNew() {
	s.env.SetContinueAsNewSuggested(true)

	var updatesCompleted int
	s.env.RegisterDelayedCallback(func() {
		s.env.UpdateWorkflow(shoppingcartv1.UpdateCartUpdateName, uuid.NewString(), &testsuite.TestUpdateCallback{
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
		}, &shoppingcartv1.UpdateCartInput{
			Action: shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_ADD,
			ItemId: "foo",
		})
	}, 0)

	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(shoppingcartv1.CheckoutSignalName, &shoppingcartv1.CheckoutInput{})
	}, time.Second)

	s.env.ExecuteWorkflow(shoppingcartv1.ShoppingCartWorkflowName, &shoppingcartv1.ShoppingCartInput{})

	s.Require().True(s.env.IsWorkflowCompleted())
	err := s.env.GetWorkflowError()
	s.Require().True(workflow.IsContinueAsNewError(err), "Expected workflow to continue as new, got: %v", err)
}

func TestShoppingCartWithUpdateCartE2E(t *testing.T) {
	existingPath := which.Which("temporal")
	if existingPath == "" {
		t.Skip("temporal CLI not found in PATH, skipping E2E test")
	}
	ctx, r := context.Background(), require.New(t)

	logger := log.NewStructuredLogger(slog.New(slog.NewTextHandler(os.Stderr, nil)))

	srv, err := testsuite.StartDevServer(ctx, testsuite.DevServerOptions{
		ExistingPath: existingPath,
		EnableUI:     true,
		ClientOptions: &client.Options{
			HostPort: "localhost:7233",
			Logger:   logger,
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

	dw.RegisterWorkflowWithOptions(func(ctx workflow.Context, input *shoppingcartv1.ShoppingCartInput) (*shoppingcartv1.ShoppingCartOutput, error) {
		update, run, err := shoppingcartv1xns.ShoppingCartWithUpdateCart(
			ctx,
			input,
			&shoppingcartv1.UpdateCartInput{
				Action: shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_ADD,
				ItemId: "foo",
			},
		)
		if err != nil {
			return nil, err
		}
		out, err := run.Get(ctx)
		if err != nil {
			return nil, err
		}
		if !proto.Equal(update.GetCart(), out.GetCart()) {
			return nil, temporal.NewNonRetryableApplicationError(
				"Cart state mismatch",
				"InvalidState",
				nil,
			)
		}
		return out, nil
	}, workflow.RegisterOptions{
		Name: "test",
	})

	r.NoError(dw.Start())
	t.Cleanup(dw.Stop)

	id := uuid.NewString()

	run, err := dc.ExecuteWorkflow(
		ctx,
		client.StartWorkflowOptions{
			ID:        id,
			TaskQueue: "test",
		},
		"test",
		&shoppingcartv1.ShoppingCartInput{},
	)
	r.NoError(err)

	var out shoppingcartv1.ShoppingCartOutput
	r.NoError(run.Get(ctx, &out))
	r.Contains(out.GetCart().GetItems(), "foo")
	r.Equal(int32(1), out.GetCart().GetItems()["foo"])
}
