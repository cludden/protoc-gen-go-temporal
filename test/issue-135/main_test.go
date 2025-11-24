package main

import (
	"testing"

	issue_135v1 "github.com/cludden/protoc-gen-go-temporal/gen/test/issue-135/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestIssue135(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestActivityEnvironment()
	issue_135v1.RegisterFOOBarActivities(env, &Activities{})
	result, err := env.ExecuteActivity(issue_135v1.DoActivityName, &issue_135v1.DoRequest{})
	require.NoError(t, err)
	require.NotNil(t, result)
	var out issue_135v1.DoResponse
	require.NoError(t, result.Get(&out))
	require.Empty(t, cmp.Diff(&issue_135v1.DoResponse{}, &out, protocmp.Transform()))
}
