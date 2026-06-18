package testutil

import (
	"testing"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
)

var _ worker.ActivityRegistry = &activityTestEnvironment{}

type activityTestEnvironment struct {
	*testsuite.TestActivityEnvironment
}

func NewActivityTestEnvironment(t *testing.T) *activityTestEnvironment {
	var s testsuite.WorkflowTestSuite
	return &activityTestEnvironment{
		TestActivityEnvironment: s.NewTestActivityEnvironment(),
	}
}

func (s *activityTestEnvironment) RegisterDynamicActivity(a any, options activity.DynamicRegisterOptions) {
	panic("not implemented")
}
