package main

import (
	"context"
	"log"

	activityv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/activity/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.DialContext(context.Background(), client.Options{})
	if err != nil {
		log.Fatalf("unable to create client: %v", err)
	}
	defer c.Close()

	w := worker.New(c, activityv1.ExampleTaskQueue, worker.Options{})
	activityv1.RegisterExampleActivities(w, &Activities{})

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalf("unable to start worker: %v", err)
	}
}

type Activities struct {
	activityv1.ExampleActivities
}

func (a *Activities) Foo(ctx context.Context, input *activityv1.FooInput) (*activityv1.FooOutput, error) {
	return &activityv1.FooOutput{Output: input.GetInput()}, nil
}

func (a *Activities) Bar(ctx context.Context, input *activityv1.BarInput) error {
	return nil
}

func (a *Activities) Baz(ctx context.Context) (*activityv1.BazOutput, error) {
	return &activityv1.BazOutput{}, nil
}

func (a *Activities) Qux(ctx context.Context) error {
	return nil
}
