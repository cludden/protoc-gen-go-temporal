package main

import (
	"testing"

	activityv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/activity/v1"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/proto"
)

func TestActivity(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	activityv1.RegisterExampleActivities(env, &Activities{})
	env.OnActivity(activityv1.FooActivityName, mock.Anything, mock.MatchedBy(func(input *activityv1.FooInput) bool {
		return input.GetInput() == "input"
	})).Return(&activityv1.FooOutput{Output: "output"}, nil)
	env.OnActivity(activityv1.BarActivityName, mock.Anything, mock.MatchedBy(func(input *activityv1.BarInput) bool {
		return input.GetInput() == "output"
	})).Return(nil)
	env.OnActivity(activityv1.BazActivityName, mock.Anything).Return(&activityv1.BazOutput{Output: "test"}, nil)
	env.OnActivity(activityv1.QuxActivityName, mock.Anything).Return(nil)
	env.ExecuteWorkflow(func(ctx workflow.Context) (*activityv1.BazOutput, error) {
		foo, err := activityv1.Foo(ctx, &activityv1.FooInput{Input: "input"})
		if err != nil {
			return nil, err
		}
		if err = activityv1.Bar(ctx, &activityv1.BarInput{Input: foo.GetOutput()}); err != nil {
			return nil, err
		}
		baz, err := activityv1.Baz(ctx)
		if err != nil {
			return nil, err
		}
		return baz, activityv1.Qux(ctx)
	})
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	var out activityv1.BazOutput
	require.NoError(t, env.GetWorkflowResult(&out))
	require.True(t, proto.Equal(&activityv1.BazOutput{Output: "test"}, &out))
}
