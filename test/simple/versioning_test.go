package main

import (
	"context"
	"testing"

	simplepb "github.com/cludden/protoc-gen-go-temporal/gen/test/simple/v1"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

// TestUpdateWithStartVersioningOverride covers the update-with-start path, which
// reaches the versioning override only by delegating to the workflow options'
// own Build.
func TestUpdateWithStartVersioningOverride(t *testing.T) {
	pinned := &client.PinnedVersioningOverride{
		Version: worker.WorkerDeploymentVersion{DeploymentName: "example", BuildID: "abc123"},
	}

	// capture the start workflow options handed to the start operation
	var captured client.StartWorkflowOptions
	capture := func(swo client.StartWorkflowOptions) client.WithStartWorkflowOperation {
		captured = swo
		return nil
	}

	_, err := simplepb.NewSomeWorkflow1WithSomeUpdate2Options().
		WithSomeWorkflow1Options(
			simplepb.NewSomeWorkflow1Options().WithVersioningOverride(pinned),
		).
		Build(
			context.Background(),
			capture,
			&simplepb.SomeWorkflow1Request{Id: "some-id"},
			&simplepb.SomeUpdate2Request{},
		)
	require.NoError(t, err)
	require.Equal(t, pinned, captured.VersioningOverride)
}
