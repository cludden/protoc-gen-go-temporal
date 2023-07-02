package mutex

import (
	"context"
	"testing"

	mutexv1 "github.com/cludden/protoc-gen-go-temporal/example/gen/example/mutex"
	"github.com/cludden/protoc-gen-go-temporal/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/client"
)

func TestSampleWorkflowWithMutexWorkflow(t *testing.T) {
	require := require.New(t)
	c := mocks.NewClient(t)
	var opts client.StartWorkflowOptions
	c.On("ExecuteWorkflow", mock.Anything, mock.Anything, mutexv1.SampleWorkflowWithMutexWorkflowName, mock.AnythingOfType("*mutex.SampleWorkflowWithMutexRequest")).
		Return(func(ctx context.Context, options client.StartWorkflowOptions, wf interface{}, req ...interface{}) (client.WorkflowRun, error) {
			opts = options
			return &sampleWorkflowWithMutexRun{}, nil
		}, nil)

	mutex := mutexv1.NewClient(c)
	run, err := mutex.SampleWorkflowWithMutexAsync(context.Background(), &mutexv1.SampleWorkflowWithMutexRequest{Resource: "foo"})
	require.NoError(err)
	require.NotNil(run)

	require.Regexp(`sample-workflow-with-mutex/foo/.{32}`, opts.ID)
	require.Equal(mutexv1.MutexTaskQueue, opts.TaskQueue)
	require.NotNil(opts.SearchAttributes)
	require.Len(opts.SearchAttributes, 2)
	require.Equal("foo", opts.SearchAttributes["resource"])
	require.Equal("bar", opts.SearchAttributes["foo"])
}

// ============================================================================

var _ client.WorkflowRun = &sampleWorkflowWithMutexRun{}

type sampleWorkflowWithMutexRun struct{}

func (r *sampleWorkflowWithMutexRun) Get(context.Context, any) error {
	return nil
}

func (r *sampleWorkflowWithMutexRun) GetID() string {
	return ""
}

func (r *sampleWorkflowWithMutexRun) GetRunID() string {
	return ""
}

func (r *sampleWorkflowWithMutexRun) GetWithOptions(context.Context, any, client.WorkflowRunGetOptions) error {
	return nil
}
