package main

import (
	"context"
	"testing"

	simplepb "github.com/cludden/protoc-gen-go-temporal/gen/test/simple/v1"
	"github.com/cludden/protoc-gen-go-temporal/internal/testutil"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/worker"
)

func TestCustomTaskQueue(t *testing.T) {
	// override the default task queue for the test
	defaultTaskQueue := simplepb.SimpleTaskQueue
	simplepb.SimpleTaskQueue = "custom-task-queue"
	defer func() { simplepb.SimpleTaskQueue = defaultTaskQueue }()

	c, ctx := testutil.NewIntegrationEnv(t, nil), context.Background()

	// start a worker on a custom task queue
	w := worker.New(c, "custom-task-queue", worker.Options{})
	simplepb.RegisterSimpleWorkflows(w, &Workflows{taskQueue: "custom-task-queue"})
	simplepb.RegisterSimpleActivities(w, &Activities{})
	require.NoError(t, w.Start())
	t.Cleanup(w.Stop)

	client := simplepb.NewSimpleClient(c)
	run, err := client.SomeWorkflow1Async(ctx, &simplepb.SomeWorkflow1Request{
		Id:         "test-id",
		RequestVal: "test request",
	})
	require.NoError(t, err)

	require.NoError(t, run.SomeSignal1(ctx))
	require.NoError(t, run.SomeSignal2(ctx, &simplepb.SomeSignal2Request{RequestVal: "foo"}))
	require.NoError(t, run.SomeSignal2(ctx, &simplepb.SomeSignal2Request{RequestVal: "bar"}))

	out, err := run.Get(ctx)
	require.NoError(t, err)
	require.Equal(t, "started with param test request\nsome activity 3 with response some response\nsome local activity 3 with response some response\nsome signal 1\nsome signal 2 with param foo\nsome signal 2 with param bar", out.GetResponseVal())
}
