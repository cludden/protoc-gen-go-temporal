package main

import (
	"context"
	"testing"
	"time"

	updatabletimerv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/updatabletimer/v1"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestUpdatableTimerWorkflow(t *testing.T) {
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	updatabletimerv1.RegisterUpdatableTimerWorkflow(env, NewUpdatableTimerWorkflow)
	client := updatabletimerv1.NewTestExampleClient(env, nil, nil)
	ctx, require := context.Background(), require.New(t)

	start := env.Now()
	initial := start.Add(time.Hour).UTC()

	timer, err := client.UpdatableTimerAsync(ctx, &updatabletimerv1.UpdatableTimerInput{
		InitialWakeUpTime: timestamppb.New(initial),
	})
	require.NoError(err)

	env.RegisterDelayedCallback(func() {
		t, err := timer.GetWakeUpTime(ctx)
		require.NoError(err)
		require.Equal(initial, t.GetWakeUpTime().AsTime())
	}, time.Minute)

	env.RegisterDelayedCallback(func() {
		require.NoError(
			timer.UpdateWakeUpTime(ctx, &updatabletimerv1.UpdateWakeUpTimeInput{
				WakeUpTime: timestamppb.New(initial.Add(time.Hour)),
			}),
		)
	}, time.Minute*30)

	env.RegisterDelayedCallback(func() {
		t, err := timer.GetWakeUpTime(ctx)
		require.NoError(err)
		require.Equal(initial.Add(time.Hour), t.GetWakeUpTime().AsTime())
	}, time.Minute*119)

	require.NoError(timer.Get(ctx))
	require.Equal(time.Hour*2, env.Now().Sub(start))
}
