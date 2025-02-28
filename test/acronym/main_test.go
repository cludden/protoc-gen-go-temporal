package main

import (
	"context"
	"testing"

	acronymv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/acronym/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestExample(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	client := acronymv1.NewTestAWSClient(env, &Workflows{}, &Activities{})

	input := &acronymv1.ManageAWSRequest{Urn: "test"}

	env.
		OnActivity(acronymv1.ManageAWSResourceActivityName, mock.Anything, mock.MatchedBy(func(req *acronymv1.ManageAWSResourceRequest) bool {
			return assert.Empty(t, cmp.Diff(&acronymv1.ManageAWSResourceRequest{
				Urn: input.GetUrn(),
			}, req, protocmp.Transform()))
		})).
		Return(&acronymv1.ManageAWSResourceResponse{}, nil)

	resp, err := client.ManageAWS(context.Background(), input)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
