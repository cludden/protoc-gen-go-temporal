package shoppingcart

import (
	shoppingcartv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/shoppingcart/v1"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/proto"
)

type (
	Workflows struct{}

	ShoppingCartWorkflow struct {
		*Workflows
		shoppingcartv1.ShoppingCartWorkflowInput
		cart *shoppingcartv1.CartState
	}
)

func (w *Workflows) ShoppingCart(
	ctx workflow.Context,
	input *shoppingcartv1.ShoppingCartWorkflowInput,
) (shoppingcartv1.ShoppingCartWorkflow, error) {
	cart := input.Req.GetCart()
	if cart.GetItems() == nil {
		cart = &shoppingcartv1.CartState{
			Items: make(map[string]int32),
		}
	}
	return &ShoppingCartWorkflow{
		Workflows:                 w,
		ShoppingCartWorkflowInput: *input,
		cart:                      cart,
	}, nil
}

func (w *ShoppingCartWorkflow) Execute(
	ctx workflow.Context,
) (*shoppingcartv1.ShoppingCartOutput, error) {
	if err := workflow.Await(ctx, func() bool {
		if workflow.GetInfo(ctx).GetContinueAsNewSuggested() {
			return true
		}
		if checkout := w.Checkout.ReceiveAsync(); checkout != nil {
			return true
		}
		return false
	}); err != nil {
		return nil, err
	}

	if workflow.GetInfo(ctx).GetContinueAsNewSuggested() {
		if err := workflow.Await(ctx, func() bool {
			return workflow.AllHandlersFinished(ctx)
		}); err != nil {
			return nil, err
		}
		return nil, workflow.NewContinueAsNewError(
			ctx,
			shoppingcartv1.ShoppingCartWorkflowName,
			&shoppingcartv1.ShoppingCartInput{
				Cart: w.cart,
			},
		)
	}
	return &shoppingcartv1.ShoppingCartOutput{
		Cart: w.cart,
	}, nil
}

func (w *ShoppingCartWorkflow) Describe(
	*shoppingcartv1.DescribeInput,
) (*shoppingcartv1.DescribeOutput, error) {
	return &shoppingcartv1.DescribeOutput{
		Cart: w.cart,
	}, nil
}

func (w *ShoppingCartWorkflow) UpdateCart(
	ctx workflow.Context,
	input *shoppingcartv1.UpdateCartInput,
) (*shoppingcartv1.UpdateCartOutput, error) {
	switch input.GetAction() {
	case shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_ADD:
		w.cart.Items[input.GetItemId()] += 1
	case shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_REMOVE:
		w.cart.Items[input.GetItemId()] -= 1
		if w.cart.Items[input.GetItemId()] <= 0 {
			delete(w.cart.Items, input.GetItemId())
		}
	default:
		return nil, temporal.NewNonRetryableApplicationError(
			"Invalid update cart action",
			"Unimplemented",
			nil,
		)
	}
	return &shoppingcartv1.UpdateCartOutput{
		Cart: proto.Clone(w.cart).(*shoppingcartv1.CartState),
	}, nil
}

func (w *ShoppingCartWorkflow) ValidateUpdateCart(
	ctx workflow.Context,
	input *shoppingcartv1.UpdateCartInput,
) error {
	if input.GetAction() == shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_UNSPECIFIED ||
		input.GetItemId() == "" {
		return temporal.NewNonRetryableApplicationError(
			"Invalid update cart request",
			"InvalidArgument",
			nil,
		)
	}
	if input.GetAction() == shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_REMOVE &&
		w.cart.GetItems()[input.GetItemId()] == 0 {
		return temporal.NewNonRetryableApplicationError("Item not found in cart", "NotFound", nil)
	}
	return nil
}
