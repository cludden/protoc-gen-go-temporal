package simple_test

import (
	"context"
	"testing"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/cludden/protoc-gen-go-temporal/test/simple"
	simplepb "github.com/cludden/protoc-gen-go-temporal/test/simple/gen/test/simple/v1"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func TestSomeWorkflow1(t *testing.T) {
	require := require.New(t)

	// initialize docker pool
	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Skipf("error initializing docker pool: %v", err)
	}
	if err := pool.Client.Ping(); err != nil {
		t.Skipf("error pinging docker daemon: %v", err)
	}

	// start temporalite container
	temporalite, err := pool.Run("cludden/temporalite", "0.3.0", nil)
	require.NoError(err)
	require.NoError(temporalite.Expire(120))

	// initialize temporal client
	var c client.Client
	require.NoError(retry.Do(func() (err error) {
		c, err = client.Dial(client.Options{
			HostPort: temporalite.GetHostPort("7233/tcp"),
		})
		return err
	}, retry.Delay(time.Second*30), retry.MaxDelay(time.Second*30), retry.Attempts(3)))
	defer c.Close()

	// initialize worker and register workflows, activities
	w := worker.New(c, simplepb.SimpleTaskQueue, worker.Options{})
	simple.Register(w)
	require.NoError(w.Start())
	defer w.Stop()

	// initialize simple client
	client := simplepb.NewClient(c)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start the workflow
	run, err := client.SomeWorkflow1Async(ctx, &simplepb.SomeWorkflow1Request{
		Id:         "foo",
		RequestVal: "some request",
	})
	require.NoError(err)
	require.Regexp("^some-workflow-1/foo/.{32}", run.ID())

	// send signals
	require.NoError(run.SomeSignal1(ctx))
	require.NoError(run.SomeSignal2(ctx, &simplepb.SomeSignal2Request{RequestVal: "foo"}))

	// query until we get the right events
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
