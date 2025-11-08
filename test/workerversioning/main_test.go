package main

import (
	"testing"

	workerversioningv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/workerversioning/v1"
	workermocks "github.com/cludden/protoc-gen-go-temporal/mocks/go.temporal.io/sdk/worker"
	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/workflow"
)

func TestWorkerVersioning(t *testing.T) {
	r := workermocks.NewMockRegistry(t)
	r.EXPECT().RegisterWorkflowWithOptions(mock.Anything, workflow.RegisterOptions{
		Name: workerversioningv1.FooWorkflowName,
	})
	r.EXPECT().RegisterWorkflowWithOptions(mock.Anything, workflow.RegisterOptions{
		Name: workerversioningv1.BarWorkflowName,
	})
	r.EXPECT().RegisterWorkflowWithOptions(mock.Anything, workflow.RegisterOptions{
		Name:               workerversioningv1.BazWorkflowName,
		VersioningBehavior: workflow.VersioningBehaviorPinned,
	})
	r.EXPECT().RegisterWorkflowWithOptions(mock.Anything, workflow.RegisterOptions{
		Name:               workerversioningv1.QuxWorkflowName,
		VersioningBehavior: workflow.VersioningBehaviorAutoUpgrade,
	})
	Register(r)
}
