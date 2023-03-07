package simple_test

import (
	"context"
	"testing"
	"time"

	simplepb "github.com/cludden/protoc-gen-go-temporal/gen/simple"
	"github.com/cludden/protoc-gen-go-temporal/test/simple"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func TestSomeWorkflow1(t *testing.T) {
	require := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	temporalClient, err := client.Dial(client.Options{})
	require.NoError(err)
	defer temporalClient.Close()

	taskQueue := uuid.NewString()
	w := worker.New(temporalClient, taskQueue, worker.Options{})
	simple.Register(w)
	require.NoError(w.Start())
	defer w.Stop()

	// Start server and worker w/ workflow registered

	// Create the client
	c := simplepb.NewClient(temporalClient)

	// Start the workflow
	run, err := c.ExecuteSomeWorkflow1(
		ctx,
		&client.StartWorkflowOptions{TaskQueue: taskQueue},
		&simplepb.SomeWorkflow1Request{RequestVal: "some request"},
	)
	require.NoError(err)

	err = run.SomeSignal1(ctx)
	require.NoError(err)

	err = run.SomeSignal2(ctx, &simplepb.SomeSignal2Request{RequestVal: "foo"})
	require.NoError(err)

	// Query until we get the right events
	require.Eventually(func() bool {
		resp, err := run.SomeQuery1(ctx)
		require.NoError(err)
		for _, item := range []string{
			"started with param some request",
			"some activity 3 with response some response",
			"some local activity 3 with response some response",
			"some query 1",
		} {
			require.Contains(resp.GetResponseVal(), item)
		}
		return true
	}, 2*time.Second, 200*time.Millisecond)

	// Check the activity events
	require.Equal([]string{
		"some activity 3 with param some activity param",
		"some activity 3 with param some local activity param",
	}, simple.ActivityEvents)

	resp, err := run.Get(ctx)
	require.NoError(err)
	require.NotNil(resp)
}
