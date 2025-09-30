package main

import (
	"context"
	"testing"
	"time"

	issue_125v1 "github.com/cludden/protoc-gen-go-temporal/gen/test/issue-125/v1"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestIssue125(t *testing.T) {
	t.Setenv("TEMPORAL_DEBUG", "true")
	var s testsuite.WorkflowTestSuite
	ctx := context.Background()

	env := s.NewTestWorkflowEnvironment()
	client := issue_125v1.NewTestIssue125ServiceClient(env, &Workflows{}, nil)

	run, err := client.FooAsync(ctx, &issue_125v1.FooInput{Id: "1"})
	require.NoError(t, err)
	require.NotNil(t, run)

	var (
		barHandle issue_125v1.BarHandle
		bazHandle issue_125v1.BazHandle
	)

	env.RegisterDelayedCallback(func() {
		var err error
		barHandle, err = run.BarAsync(ctx, &issue_125v1.BarInput{Id: "2"})
		require.NoError(t, err)
		require.NotNil(t, barHandle)
	}, time.Second*10)

	env.RegisterDelayedCallback(func() {
		var err error
		bazHandle, err = run.BazAsync(ctx, &issue_125v1.BazInput{Id: "3"})
		require.NoError(t, err)
		require.NotNil(t, bazHandle)
	}, time.Hour*2)

	result, err := run.Get(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, "foo:1", result.GetResult())

	bar, err := barHandle.Get(ctx)
	require.NoError(t, err)
	require.NotNil(t, bar)
	require.Equal(t, "bar:2", bar.GetResult())

	baz, err := bazHandle.Get(ctx)
	require.NoError(t, err)
	require.NotNil(t, baz)
	require.Equal(t, "baz:3", baz.GetResult())
}
