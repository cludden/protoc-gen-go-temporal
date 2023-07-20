package main

import (
	"bytes"
	"context"
	"testing"
	"time"

	simplepb "github.com/cludden/protoc-gen-go-temporal/gen/simple"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
)

func TestSomeWorkflow1WithTestClient(t *testing.T) {
	ActivityEvents = nil
	require := require.New(t)
	suite := &testsuite.WorkflowTestSuite{}
	env := suite.NewTestWorkflowEnvironment()
	client := simplepb.NewTestSimpleClient(env, &Workflows{}, &Activities{})

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
	env.RegisterDelayedCallback(func() {
		require.NoError(run.SomeSignal1(ctx))
	}, time.Minute*10)
	env.RegisterDelayedCallback(func() {
		require.NoError(run.SomeSignal2(ctx, &simplepb.SomeSignal2Request{RequestVal: "foo"}))
	}, time.Minute*30)

	// query until we get the right events
	env.RegisterDelayedCallback(func() {
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
	}, time.Minute*41)

	env.RegisterDelayedCallback(func() {
		require.NoError(run.SomeSignal2(ctx, &simplepb.SomeSignal2Request{RequestVal: "foo"}))
	}, time.Minute*50)

	resp, err := run.Get(ctx)
	require.NoError(err)
	require.NotNil(resp)

	// Check the activity events
	require.Equal([]string{
		"some activity 3 with param some activity param",
		"some activity 3 with param some local activity param",
	}, ActivityEvents)
}

func TestSomeWorkflow1WithClient(t *testing.T) {
	ActivityEvents = nil
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
	require.NoError(pool.Retry(func() (err error) {
		c, err = client.Dial(client.Options{
			HostPort: temporalite.GetHostPort("7233/tcp"),
		})
		return err
	}))
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
	require.NoError(run.SomeSignal2(ctx, &simplepb.SomeSignal2Request{RequestVal: "bar"}))

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

func TestSomeWorkflow2WithTestClient(t *testing.T) {
	require, ctx := require.New(t), context.Background()
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	client := simplepb.NewTestSimpleClient(env, &Workflows{}, &Activities{})

	run, err := client.SomeWorkflow2Async(ctx)
	require.NoError(err)
	require.NotNil(run)

	var handle simplepb.SomeUpdate1Handle
	env.RegisterDelayedCallback(func() {
		handle, err = run.SomeUpdate1Async(ctx, &simplepb.SomeUpdate1Request{RequestVal: "test"})
		require.NoError(err)
	}, time.Second)

	require.NoError(run.Get(ctx))
	update, err := handle.Get(ctx)
	require.NoError(err)
	require.Equal("TEST", update.GetResponseVal())
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
