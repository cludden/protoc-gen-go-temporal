package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"testing"
	"time"

	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	commonv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/simple/common/v1"
	simplepb "github.com/cludden/protoc-gen-go-temporal/gen/test/simple/v1"
	simplemocks "github.com/cludden/protoc-gen-go-temporal/mocks/test/simple/v1"
	"github.com/cludden/protoc-gen-go-temporal/pkg/helpers"
	"github.com/cludden/protoc-gen-go-temporal/pkg/patch"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
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

func TestSomeWorkflow1Alias(t *testing.T) {
	ActivityEvents = nil
	require := require.New(t)
	suite := &testsuite.WorkflowTestSuite{}
	env := suite.NewTestWorkflowEnvironment()
	simplepb.RegisterSimpleWorkflows(env, &Workflows{})
	simplepb.RegisterSimpleActivities(env, &Activities{})

	// send signals
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(simplepb.SomeSignal1SignalName, nil)
	}, time.Minute*10)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(simplepb.SomeSignal2SignalName, &simplepb.SomeSignal2Request{RequestVal: "foo"})
	}, time.Minute*30)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(simplepb.SomeSignal2SignalName, &simplepb.SomeSignal2Request{RequestVal: "bar"})
	}, time.Minute*60)

	env.ExecuteWorkflow("mycompany.SomeWorkflow1", &simplepb.SomeWorkflow1Request{Id: "bar", RequestVal: "blah"})
	require.True(env.IsWorkflowCompleted())
	require.NoError(env.GetWorkflowError())
	var out simplepb.SomeWorkflow1Response
	require.NoError(env.GetWorkflowResult(&out))
}

func TestSomeWorkflow1Child(t *testing.T) {
	require := require.New(t)
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	simplepb.RegisterSimpleActivities(env, &Activities{})
	simplepb.RegisterSimpleWorkflows(env, &Workflows{})

	env.RegisterWorkflowWithOptions(func(ctx workflow.Context) error {
		_, err := simplepb.SomeWorkflow1Child(ctx, &simplepb.SomeWorkflow1Request{
			RequestVal: "foo",
			Id:         "test",
		})
		return err
	}, workflow.RegisterOptions{Name: "test"})

	env.OnWorkflow(simplepb.SomeWorkflow1WorkflowName, mock.Anything, mock.Anything).
		Return(&simplepb.SomeWorkflow1Response{}, nil)

	var id string
	env.SetOnLocalActivityCompletedListener(func(activityInfo *activity.Info, result converter.EncodedValue, err error) {
		require.NoError(result.Get(&id))
	})

	env.ExecuteWorkflow("test")

	require.True(env.IsWorkflowCompleted())
	require.NoError(env.GetWorkflowError())
	require.Regexp(regexp.MustCompile(`some-workflow-1/test/[a-f0-9-]{32}`), id)
}

func TestSomeWorkflow1Child_DefaultVersion(t *testing.T) {
	require := require.New(t)
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	simplepb.RegisterSimpleActivities(env, &Activities{})
	simplepb.RegisterSimpleWorkflows(env, &Workflows{})

	env.RegisterWorkflowWithOptions(func(ctx workflow.Context) error {
		_, err := simplepb.SomeWorkflow1Child(ctx, &simplepb.SomeWorkflow1Request{
			RequestVal: "foo",
			Id:         "test",
		})
		return err
	}, workflow.RegisterOptions{Name: "test"})

	env.OnGetVersion(patch.PV_64_ExpressionEvaluationLocalActivity, workflow.DefaultVersion, 1).Return(workflow.DefaultVersion)

	env.OnWorkflow(simplepb.SomeWorkflow1WorkflowName, mock.Anything, mock.Anything).
		Return(&simplepb.SomeWorkflow1Response{}, nil)

	var called bool
	env.SetOnLocalActivityCompletedListener(func(activityInfo *activity.Info, result converter.EncodedValue, err error) {
		called = true
	})

	env.ExecuteWorkflow("test")

	require.True(env.IsWorkflowCompleted())
	require.NoError(env.GetWorkflowError())
	require.False(called)
}

func TestSomeWorkflow2WithTestClient(t *testing.T) {
	require, ctx := require.New(t), context.Background()
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	client := simplepb.NewTestSimpleClient(env, &Workflows{simplepb.NewSimpleWorkflowFunctions(), nil, ""}, &Activities{})

	run, err := client.SomeWorkflow2Async(ctx)
	require.NoError(err)
	require.NotNil(run)

	env.RegisterDelayedCallback(func() {
		require.NoError(client.SomeSignal2(ctx, run.ID(), "", &simplepb.SomeSignal2Request{RequestVal: "bar"}))
	}, time.Second)

	var handle simplepb.SomeUpdate1Handle
	env.RegisterDelayedCallback(func() {
		handle, err = run.SomeUpdate1Async(ctx, &simplepb.SomeUpdate1Request{RequestVal: "test"})
		require.NoError(err)
	}, time.Second*3)

	require.NoError(run.Get(ctx))
	update, err := handle.Get(ctx)
	require.NoError(err)
	require.Equal("TEST", update.GetResponseVal())
}

func TestSomeWorkflow2WithMock(t *testing.T) {
	require, ctx := require.New(t), context.Background()
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	fns := simplemocks.NewMockSimpleWorkflowFunctions(t)
	fns.EXPECT().SomeWorkflow3(mock.Anything, mock.Anything).Return(nil)
	client := simplepb.NewTestSimpleClient(env, &Workflows{fns, nil, ""}, &Activities{})

	run, err := client.SomeWorkflow2Async(ctx)
	require.NoError(err)
	require.NotNil(run)

	var handle simplepb.SomeUpdate1Handle
	env.RegisterDelayedCallback(func() {
		handle, err = run.SomeUpdate1Async(ctx, &simplepb.SomeUpdate1Request{RequestVal: "test"})
		require.NoError(err)
	}, time.Second*3)

	require.NoError(run.Get(ctx))
	update, err := handle.Get(ctx)
	require.NoError(err)
	require.Equal("TEST", update.GetResponseVal())
}

func TestSomeWorkflow2Child(t *testing.T) {
	require := require.New(t)
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	simplepb.RegisterSimpleActivities(env, &Activities{})
	simplepb.RegisterSimpleWorkflows(env, &Workflows{})

	env.RegisterWorkflowWithOptions(func(ctx workflow.Context) error {
		return simplepb.SomeWorkflow2Child(ctx)
	}, workflow.RegisterOptions{Name: "test"})

	env.OnWorkflow(simplepb.SomeWorkflow2WorkflowName, mock.Anything, mock.Anything).
		Return(nil)

	var id string
	env.SetOnLocalActivityCompletedListener(func(activityInfo *activity.Info, result converter.EncodedValue, err error) {
		require.NoError(result.Get(&id))
	})

	env.ExecuteWorkflow("test")

	require.True(env.IsWorkflowCompleted())
	require.NoError(env.GetWorkflowError())
	require.Regexp(regexp.MustCompile(`some-workflow-2/[a-f0-9-]{32}`), id)
}

func TestSomeWorkflow3Child(t *testing.T) {
	require := require.New(t)
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	simplepb.RegisterSimpleActivities(env, &Activities{})
	simplepb.RegisterSimpleWorkflows(env, &Workflows{})

	env.RegisterWorkflowWithOptions(func(ctx workflow.Context) error {
		return simplepb.SomeWorkflow3Child(ctx, &simplepb.SomeWorkflow3Request{Id: "test", RequestVal: "foo"})
	}, workflow.RegisterOptions{Name: "test"})

	env.OnWorkflow(simplepb.SomeWorkflow3WorkflowName, mock.Anything, mock.Anything).
		Return(nil)

	var getVersionCalled bool
	env.OnGetVersion(patch.PV_64_ExpressionEvaluationLocalActivity, workflow.DefaultVersion, workflow.Version(1)).Run(func(args mock.Arguments) {
		getVersionCalled = true
	}).Return(workflow.Version(1))
	env.OnGetVersion(patch.PV_64_ExpressionEvaluationLocalActivity, workflow.Version(1), workflow.Version(1)).Run(func(args mock.Arguments) {
		getVersionCalled = true
	}).Return(workflow.Version(1))

	var localActivityCalled bool
	env.SetOnLocalActivityCompletedListener(func(activityInfo *activity.Info, result converter.EncodedValue, err error) {
		localActivityCalled = true
	})

	env.ExecuteWorkflow("test")

	require.True(env.IsWorkflowCompleted())
	require.NoError(env.GetWorkflowError())
	require.False(getVersionCalled)
	require.False(localActivityCalled)
}

func TestSomeWorkflow4C(t *testing.T) {
	require := require.New(t)
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	c := simplepb.NewTestSimpleClient(env, &Workflows{
		items: []*simplepb.Foo{
			{Foo: "foo"},
			{Foo: "bar"},
			{Foo: "baz"},
			{Foo: "qux"},
			{Foo: "quux"},
		},
	}, &Activities{})
	resp, err := c.SomeWorkflow4(context.Background(), &commonv1.PaginatedRequest{Limit: 3, Cursor: []byte(`1`)})
	require.NoError(err)
	require.Equal("4", string(resp.GetNextCursor()))
	require.Len(resp.GetItems(), 3)
	for i, foo := range []string{"bar", "baz", "qux"} {
		var f simplepb.Foo
		require.NoError(anypb.UnmarshalTo(resp.GetItems()[i], &f, proto.UnmarshalOptions{}))
		require.Equal(foo, f.GetFoo())
	}
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
			match: []string{`COMMANDS:\s+simple,\s+s\s+Simple\s+operations\s+other\s+mycompany.simple.Other operations\b`},
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
		opts   helpers.UnmarshalCliFlagsOptions
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
			opts: helpers.UnmarshalCliFlagsOptions{FromFile: "f"},
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
			c.assert(simplepb.UnmarshalCliFlagsToOtherWorkflowRequest(cmd, c.opts))
			return nil
		}
		require.NoError(app.RunContext(ctx, append([]string{
			"simple", "other-workflow",
		}, c.args...)))
	}
}

type testActivitiesTQ struct {
	Activities
}

func (a *testActivitiesTQ) SomeActivity3(ctx context.Context, req *simplepb.SomeActivity3Request) (*simplepb.SomeActivity3Response, error) {
	if tq := activity.GetInfo(ctx).TaskQueue; tq != "some-other-task-queue" {
		return nil, fmt.Errorf("expected task queue %q, got: %q", "some-other-task-queue", tq)
	}
	return &simplepb.SomeActivity3Response{ResponseVal: req.GetRequestVal()}, nil
}

func TestActivityTaskQueue(t *testing.T) {
	require := require.New(t)

	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	simplepb.RegisterSimpleActivities(env, &testActivitiesTQ{})

	env.ExecuteWorkflow(func(ctx workflow.Context) error {
		_, err := simplepb.SomeActivity3(ctx, &simplepb.SomeActivity3Request{
			RequestVal: "foo",
		})
		require.NoError(err)
		return nil
	})
	require.True(env.IsWorkflowCompleted())
	require.NoError(env.GetWorkflowError())
}

func TestActivityOptions(t *testing.T) {
	var suite testsuite.WorkflowTestSuite

	cases := []struct {
		desc     string
		opts     []*simplepb.SomeActivity4ActivityOptions
		expected workflow.ActivityOptions
	}{
		{
			desc: "defaults",
			expected: workflow.ActivityOptions{
				HeartbeatTimeout: time.Second * 30,
				RetryPolicy: &temporal.RetryPolicy{
					MaximumAttempts: 5,
				},
				ScheduleToCloseTimeout: time.Second * 300,
				ScheduleToStartTimeout: time.Second * 5,
				StartToCloseTimeout:    time.Second * 60,
				TaskQueue:              "some-other-task-queue",
				WaitForCancellation:    true,
			},
			opts: []*simplepb.SomeActivity4ActivityOptions{
				simplepb.NewSomeActivity4ActivityOptions(),
				simplepb.NewSomeActivity4ActivityOptions().
					WithActivityOptions(workflow.ActivityOptions{}),
			},
		},
		{
			desc: "base",
			expected: workflow.ActivityOptions{
				HeartbeatTimeout: time.Second * 20,
				RetryPolicy: &temporal.RetryPolicy{
					MaximumAttempts: 3,
				},
				ScheduleToCloseTimeout: time.Second * 299,
				ScheduleToStartTimeout: time.Second * 4,
				StartToCloseTimeout:    time.Second * 59,
				TaskQueue:              "bar",
				WaitForCancellation:    true,
			},
			opts: []*simplepb.SomeActivity4ActivityOptions{
				simplepb.NewSomeActivity4ActivityOptions().
					WithActivityOptions(workflow.ActivityOptions{
						HeartbeatTimeout: time.Second * 20,
						RetryPolicy: &temporal.RetryPolicy{
							MaximumAttempts: 3,
						},
						ScheduleToCloseTimeout: time.Second * 299,
						ScheduleToStartTimeout: time.Second * 4,
						StartToCloseTimeout:    time.Second * 59,
						TaskQueue:              "bar",
						WaitForCancellation:    true,
					}),
			},
		},
		{
			desc: "base partial",
			expected: workflow.ActivityOptions{
				HeartbeatTimeout: time.Second * 20,
				RetryPolicy: &temporal.RetryPolicy{
					MaximumAttempts: 5,
				},
				ScheduleToCloseTimeout: time.Second * 299,
				ScheduleToStartTimeout: time.Second * 5,
				StartToCloseTimeout:    time.Second * 59,
				TaskQueue:              "some-other-task-queue",
				WaitForCancellation:    true,
			},
			opts: []*simplepb.SomeActivity4ActivityOptions{
				simplepb.NewSomeActivity4ActivityOptions().
					WithActivityOptions(workflow.ActivityOptions{
						HeartbeatTimeout:       time.Second * 20,
						ScheduleToCloseTimeout: time.Second * 299,
						StartToCloseTimeout:    time.Second * 59,
						WaitForCancellation:    false,
					}),
			},
		},
		{
			desc: "overrides",
			expected: workflow.ActivityOptions{
				HeartbeatTimeout:       time.Second * 31,
				RetryPolicy:            &temporal.RetryPolicy{},
				ScheduleToCloseTimeout: time.Second * 301,
				ScheduleToStartTimeout: time.Second * 6,
				StartToCloseTimeout:    time.Second * 61,
				TaskQueue:              "bar",
				WaitForCancellation:    false,
			},
			opts: []*simplepb.SomeActivity4ActivityOptions{
				simplepb.NewSomeActivity4ActivityOptions().
					WithHeartbeatTimeout(time.Second * 31).
					WithRetryPolicy(&temporal.RetryPolicy{}).
					WithScheduleToCloseTimeout(time.Second * 301).
					WithScheduleToStartTimeout(time.Second * 6).
					WithStartToCloseTimeout(time.Second * 61).
					WithWaitForCancellation(false).
					WithActivityOptions(workflow.ActivityOptions{
						HeartbeatTimeout: time.Second * 20,
						RetryPolicy: &temporal.RetryPolicy{
							MaximumAttempts: 3,
						},
						ScheduleToCloseTimeout: time.Second * 299,
						ScheduleToStartTimeout: time.Second * 4,
						StartToCloseTimeout:    time.Second * 59,
						TaskQueue:              "bar",
						WaitForCancellation:    true,
					}),
			},
		},
		{
			desc: "overrides partial",
			expected: workflow.ActivityOptions{
				HeartbeatTimeout: time.Second * 20,
				RetryPolicy: &temporal.RetryPolicy{
					MaximumAttempts: 5,
				},
				ScheduleToCloseTimeout: time.Second * 300,
				ScheduleToStartTimeout: time.Second * 5,
				StartToCloseTimeout:    time.Second * 60,
				TaskQueue:              "some-other-task-queue",
				WaitForCancellation:    false,
			},
			opts: []*simplepb.SomeActivity4ActivityOptions{
				simplepb.NewSomeActivity4ActivityOptions().
					WithWaitForCancellation(false).
					WithActivityOptions(workflow.ActivityOptions{
						HeartbeatTimeout: time.Second * 20,
					}),
			},
		},
	}

	for _, c := range cases {
		for i, opts := range c.opts {
			t.Run(fmt.Sprintf("%s-%d", c.desc, i), func(t *testing.T) {
				env := suite.NewTestWorkflowEnvironment()
				var ao workflow.ActivityOptions
				env.ExecuteWorkflow(func(ctx workflow.Context) (err error) {
					ctx, err = opts.Build(ctx)
					if err != nil {
						return err
					}
					ao = workflow.GetActivityOptions(ctx)
					return nil
				})
				require.True(t, env.IsWorkflowCompleted())
				require.NoError(t, env.GetWorkflowError())
				require.Equal(t, c.expected, ao)
			})
		}

	}
}

func TestLocalActivityOptions(t *testing.T) {
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	env.RegisterActivityWithOptions(func(ctx context.Context) error {
		return nil
	}, activity.RegisterOptions{Name: simplepb.SomeActivity4ActivityName})
	env.ExecuteWorkflow(func(ctx workflow.Context) error {
		return simplepb.SomeActivity4Local(ctx, simplepb.NewSomeActivity4LocalActivityOptions().
			Local(func(ctx context.Context) error {
				return errors.New("uh oh")
			}),
		)
	})
	require.True(t, env.IsWorkflowCompleted())
	require.ErrorContains(t, env.GetWorkflowError(), "uh oh")
}

func TestContinueAsNew(t *testing.T) {
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	simple := simplepb.NewTestSimpleClient(env, &Workflows{}, nil)
	out, err := simple.ExampleContinueAsNew(context.Background(), &simplepb.ExampleContinueAsNewRequest{
		Remaining: 3,
	})
	require.Nil(t, out)
	require.Error(t, err)
	require.True(t, workflow.IsContinueAsNewError(err))
	var canerr *workflow.ContinueAsNewError
	require.ErrorAs(t, err, &canerr)
	require.NotNil(t, canerr)
	var next simplepb.ExampleContinueAsNewRequest
	require.NoError(t, converter.GetDefaultDataConverter().FromPayloads(canerr.Input, &next))
	require.Equal(t, int32(2), next.GetRemaining())
	require.Nil(t, next.GetRetryPolicy())
}

func TestContinueAsNewWithOptions(t *testing.T) {
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	simple := simplepb.NewTestSimpleClient(env, &Workflows{}, nil)
	out, err := simple.ExampleContinueAsNew(context.Background(), &simplepb.ExampleContinueAsNewRequest{
		Remaining: 3,
		RetryPolicy: &temporalv1.RetryPolicy{
			InitialInterval:    durationpb.New(time.Second * 10),
			BackoffCoefficient: 2.0,
			MaxInterval:        durationpb.New(time.Second * 30),
			MaxAttempts:        5,
		},
	})
	require.Nil(t, out)
	require.Error(t, err)
	require.True(t, workflow.IsContinueAsNewError(err))
	var canerr *workflow.ContinueAsNewError
	require.ErrorAs(t, err, &canerr)
	require.NotNil(t, canerr)
	var next simplepb.ExampleContinueAsNewRequest
	require.NoError(t, converter.GetDefaultDataConverter().FromPayloads(canerr.Input, &next))
	require.Equal(t, int32(2), next.GetRemaining())
	require.NotNil(t, next.GetRetryPolicy())
	require.Equal(t, time.Second*10, next.GetRetryPolicy().GetInitialInterval().AsDuration())
	require.Equal(t, 2.0, next.GetRetryPolicy().GetBackoffCoefficient())
	require.Equal(t, time.Second*30, next.GetRetryPolicy().GetMaxInterval().AsDuration())
	require.Equal(t, int32(5), next.GetRetryPolicy().GetMaxAttempts())
	require.NotNil(t, canerr.RetryPolicy)
	require.Equal(t, time.Second*10, canerr.RetryPolicy.InitialInterval)
	require.Equal(t, 2.0, canerr.RetryPolicy.BackoffCoefficient)
	require.Equal(t, time.Second*30, canerr.RetryPolicy.MaximumInterval)
	require.Equal(t, int32(5), canerr.RetryPolicy.MaximumAttempts)
}
