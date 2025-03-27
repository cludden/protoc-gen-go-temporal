package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/cludden/protoc-gen-go-temporal/examples/shoppingcart"
	shoppingcartv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/shoppingcart/v1"
	"github.com/cludden/protoc-gen-go-temporal/gen/example/shoppingcart/v1/shoppingcartv1xns"
	"github.com/urfave/cli/v2"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/types/known/durationpb"
)

func main() {
	app, err := shoppingcartv1.NewShoppingCartCli(
		shoppingcartv1.NewShoppingCartCliOptions().
			WithClient(newClient("shoppingcart")),
	)
	if err != nil {
		log.Fatalf("failed to create shopping cart CLI: %v", err)
	}

	app.Commands = append(app.Commands, &cli.Command{
		Name:   "start",
		Usage:  "Start the shopping cart workers",
		Action: start,
	})
	app.Flags = append(app.Flags, &cli.StringFlag{
		Name:    "log-level",
		Aliases: []string{"l"},
		Value:   "debug",
		Usage:   "Set the logging level (debug, info, warn, error)",
	})

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func start(cmd *cli.Context) error {
	logger := newLogger(cmd)
	c, err := newClient("default")(cmd)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}
	defer c.Close()

	// register shoppingcart namespace
	_, _ = c.WorkflowService().
		RegisterNamespace(cmd.Context, &workflowservice.RegisterNamespaceRequest{
			Namespace:                        "shoppingcart",
			WorkflowExecutionRetentionPeriod: durationpb.New(24 * time.Hour),
		})

	// create client for shoppingcart namespace
	sc, err := client.NewClientFromExistingWithContext(cmd.Context, c, client.Options{
		Namespace: "shoppingcart",
		Logger:    tlog.NewStructuredLogger(logger),
	})
	if err != nil {
		return fmt.Errorf("failed to create shopping cart client: %w", err)
	}
	defer sc.Close()

	// create and start default worker
	dw := worker.New(c, "default", worker.Options{
		MaxHeartbeatThrottleInterval:     0,
		DefaultHeartbeatThrottleInterval: 0,
	})
	shoppingcartv1xns.RegisterShoppingCartActivities(dw, shoppingcartv1.NewShoppingCartClient(sc))

	// register default workflow
	dw.RegisterWorkflowWithOptions(func(ctx workflow.Context) error {
		_, run, err := shoppingcartv1xns.ShoppingCartWithUpdateCart(
			ctx,
			&shoppingcartv1.ShoppingCartInput{},
			&shoppingcartv1.UpdateCartInput{
				Action: shoppingcartv1.UpdateCartAction_UPDATE_CART_ACTION_ADD,
				ItemId: "foo",
			},
		)
		if err != nil {
			return err
		}

		workflow.Go(ctx, func(ctx workflow.Context) {
			workflow.GetSignalChannel(ctx, shoppingcartv1.CheckoutSignalName).Receive(ctx, nil)
			if err := run.Checkout(ctx, &shoppingcartv1.CheckoutInput{}); err != nil {
				workflow.GetLogger(ctx).Error("failed to checkout", "error", err)
			}
		})

		if _, err := run.Get(ctx); err != nil {
			workflow.GetLogger(ctx).Error("failed to get shopping cart", "error", err)
		}
		return nil
	}, workflow.RegisterOptions{Name: "default"})

	if err := dw.Start(); err != nil {
		return fmt.Errorf("failed to start worker: %w", err)
	}
	defer dw.Stop()

	// create and start shoppingcart worker
	sw := worker.New(sc, shoppingcartv1.ShoppingCartTaskQueue, worker.Options{})
	shoppingcartv1.RegisterShoppingCartWorkflows(sw, &shoppingcart.Workflows{})
	if err := sw.Start(); err != nil {
		return fmt.Errorf("failed to start shopping cart worker: %w", err)
	}
	defer sw.Stop()

	<-cmd.Context.Done()
	log.Println("shutting down workers...")
	return nil
}

func newClient(namespace string) func(cmd *cli.Context) (client.Client, error) {
	return func(cmd *cli.Context) (client.Client, error) {
		return client.DialContext(cmd.Context, client.Options{
			Namespace: namespace,
			Logger:    tlog.NewStructuredLogger(newLogger(cmd)),
		})
	}
}

func newLogger(cmd *cli.Context) *slog.Logger {
	var level slog.Level
	switch cmd.String("log-level") {
	case "error":
		level = slog.LevelError
	case "warn":
		level = slog.LevelWarn
	case "info":
		level = slog.LevelInfo
	default:
		level = slog.LevelDebug
	}
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
}
