package main

import (
	"context"
	"io"
	"log/slog"
	"math"
	"sync"
	"testing"
	"time"

	mutexv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/mutex/v1"
	"github.com/cludden/protoc-gen-go-temporal/gen/example/mutex/v1/mutexv1xns"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
	"google.golang.org/protobuf/types/known/durationpb"
)

func TestSampleWorkflowWithMutexWorkflow(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	ctx, require := context.Background(), require.New(t)
	srv, err := testsuite.StartDevServer(ctx, testsuite.DevServerOptions{
		ClientOptions: &client.Options{
			Logger: log.NewStructuredLogger(slog.New(slog.NewJSONHandler(io.Discard, nil))),
		},
	})
	require.NoError(err)
	defer srv.Stop()

	c := srv.Client()
	defer c.Close()
	client := mutexv1.NewExampleClient(c)

	w := worker.New(c, mutexv1.ExampleTaskQueue, worker.Options{})
	mutexv1.RegisterExampleWorkflows(w, &Workflows{})
	mutexv1xns.RegisterExampleActivities(w, client)
	require.NoError(w.Start())
	defer w.Stop()

	var a, b time.Time
	var g sync.WaitGroup
	g.Add(2)
	go func() {
		defer g.Done()
		require.NoError(client.SampleWorkflowWithMutex(ctx, &mutexv1.SampleWorkflowWithMutexInput{
			ResourceId: "test",
			Sleep:      durationpb.New(time.Second * 3),
		}))
		a = time.Now()
	}()
	go func() {
		defer g.Done()
		require.NoError(client.SampleWorkflowWithMutex(ctx, &mutexv1.SampleWorkflowWithMutexInput{
			ResourceId: "test",
			Sleep:      durationpb.New(time.Second * 3),
		}))
		b = time.Now()
	}()
	g.Wait()
	require.GreaterOrEqual(math.Abs(a.Sub(b).Seconds()), float64(3), "one workflow should finish at least 3s after the other")
}
