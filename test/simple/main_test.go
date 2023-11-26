package main

import (
	"bytes"
	"context"
	"os"
	"path"
	"testing"
	"time"

	simplepb "github.com/cludden/protoc-gen-go-temporal/gen/simple"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/testsuite"
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

func TestUnmarshalCLIFlagsToOtherWorkflowRequest(t *testing.T) {
	ctx, require := context.Background(), require.New(t)
	app, err := simplepb.NewOtherCli()
	require.NoError(err)
	require.NotNil(app)
	app.Setup()

	dir := t.TempDir()
	inputFile := path.Join(dir, "req.json")
	require.NoError(os.WriteFile(inputFile, []byte(`{"someVal":"foo","baz":{"baz":"test"}}`), 0777))

	var command *cli.Command
	for _, c := range app.Commands {
		if c.Name == "other-workflow" {
			command = c
			break
		}
	}
	if command == nil {
		t.FailNow()
	}

	cases := []struct {
		args   []string
		assert func(req *simplepb.OtherWorkflowRequest, err error)
	}{
		{
			args: []string{"--some-val", "foo"},
			assert: func(req *simplepb.OtherWorkflowRequest, err error) {
				require.NoError(err)
				require.NotNil(req)
				require.Equal("foo", req.GetSomeVal())
			},
		},
		{
			args: []string{"--example-duration", "3m20s"},
			assert: func(req *simplepb.OtherWorkflowRequest, err error) {
				require.NoError(err)
				require.NotNil(req)
				require.Equal(time.Second*200, req.GetExampleDuration().AsDuration())
			},
		},
		{
			args: []string{"--example-timestamp", "2023-11-26T11:21:46.715511-07:00"},
			assert: func(req *simplepb.OtherWorkflowRequest, err error) {
				require.NoError(err)
				require.NotNil(req)
				require.Equal("2023-11-26T18:21:46.715511Z", req.GetExampleTimestamp().AsTime().Format(time.RFC3339Nano))
			},
		},
		{
			args: []string{"--example-timestamp", "2023-11-26T11:21:46.715511-07:00", "--example-duration", "3m20s"},
			assert: func(req *simplepb.OtherWorkflowRequest, err error) {
				require.NoError(err)
				require.NotNil(req)
				require.Equal("2023-11-26T18:21:46.715511Z", req.GetExampleTimestamp().AsTime().Format(time.RFC3339Nano))
			},
		},
		{
			args: []string{
				"-f", inputFile,
				"--qux", `{"qux":"example"}`,
			},
			assert: func(req *simplepb.OtherWorkflowRequest, err error) {
				require.NoError(err)
				require.NotNil(req)
				require.Equal("foo", req.GetSomeVal())
				require.Equal("test", req.GetBaz().GetBaz())
				require.Equal("example", req.GetQux().GetQux())
			},
		},
	}

	for _, c := range cases {
		command.Action = func(cmd *cli.Context) error {
			c.assert(simplepb.UnmarshalCliFlagsToOtherWorkflowRequest(cmd))
			return nil
		}
		require.NoError(app.RunContext(ctx, append([]string{
			"simple", "other-workflow",
		}, c.args...)))
	}
}
