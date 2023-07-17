package main

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/avast/retry-go/v4"
	simplepb "github.com/cludden/protoc-gen-go-temporal/gen/simple"
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
	Register(w)
	require.NoError(w.Start())
	defer w.Stop()

	// initialize simple simple
	simple := simplepb.NewSimpleClient(c)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start the workflow
	run, err := simple.SomeWorkflow1Async(ctx, &simplepb.SomeWorkflow1Request{
		Id:         "foo",
		RequestVal: "some request",
	}, simplepb.NewSomeWorkflow1Options().WithStartWorkflowOptions(client.StartWorkflowOptions{}))
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
	}, ActivityEvents)

	resp, err := run.Get(ctx)
	require.NoError(err)
	require.NotNil(resp)
}

func TestCli(t *testing.T) {
	require := require.New(t)
	app, err := newCli()
	require.NoError(err)

	cases := []struct {
		cmd   []string
		err   string
		match []string
	}{
		{
			cmd:   []string{"-h"},
			match: []string{`COMMANDS:\s+simple\s+other\b`},
		},
	}

	for _, c := range cases {
		stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
		app.Writer, app.ErrWriter = stdout, stderr
		err := app.Run(append([]string{"test"}, c.cmd...))
		if c.err != "" {
			require.ErrorContains(err, c.err)
		} else {
			require.NoError(err)
			for _, pattern := range c.match {
				require.Regexp(pattern, stdout.String())
			}
		}
	}
}
