package main

import (
	"testing"

	simplepb "github.com/cludden/protoc-gen-go-temporal/gen/test/simple/v1"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/workflow"
)

func TestOnlyActivites(t *testing.T) {
	require := require.New(t)

	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	simplepb.RegisterOnlyActivitiesActivities(env, &OnlyActivites{})

	env.ExecuteWorkflow(func(ctx workflow.Context) error {
		out, err := simplepb.LonelyActivity1(ctx, &simplepb.LonelyActivity1Request{})
		require.NoError(err)
		require.NotNil(out)
		return nil
	})
	require.True(env.IsWorkflowCompleted())
	require.NoError(env.GetWorkflowError())
}
