package main

import (
	"regexp"
	"testing"

	simplepb "github.com/cludden/protoc-gen-go-temporal/gen/test/simple/v1"
	"github.com/cludden/protoc-gen-go-temporal/pkg/patch"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/workflow"
)

func TestOtherWorkflowChild(t *testing.T) {
	require := require.New(t)
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	simplepb.RegisterOtherActivities(env, &OtherActivities{})
	simplepb.RegisterOtherWorkflows(env, &OtherWorkflows{})

	env.RegisterWorkflowWithOptions(func(ctx workflow.Context) error {
		_, err := simplepb.OtherWorkflowChild(ctx, &simplepb.OtherWorkflowRequest{})
		return err
	}, workflow.RegisterOptions{Name: "test"})

	env.OnWorkflow(simplepb.OtherWorkflowWorkflowName, mock.Anything, mock.Anything).
		Return(&simplepb.OtherWorkflowResponse{}, nil)

	var markerExists bool
	env.OnGetVersion(patch.PV_64_ExpressionEvaluationLocalActivity, workflow.Version(1), workflow.Version(1)).Run(func(args mock.Arguments) {
		markerExists = true
	}).Return(workflow.Version(1))

	var id string
	env.SetOnLocalActivityCompletedListener(func(activityInfo *activity.Info, result converter.EncodedValue, err error) {
		require.NoError(result.Get(&id))
	})

	env.ExecuteWorkflow("test")

	require.True(env.IsWorkflowCompleted())
	require.NoError(env.GetWorkflowError())
	require.True(markerExists)
	require.Regexp(regexp.MustCompile(`other-workflow/[a-f0-9-]{32}`), id)
}
