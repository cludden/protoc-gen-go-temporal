package external

import (
	"context"
	"sync"
	"testing"
	"time"

	examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	"github.com/cludden/protoc-gen-go-temporal/gen/example/v1/examplev1xns"
	mexamplev1 "github.com/cludden/protoc-gen-go-temporal/mocks/github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestProvisionFoo(t *testing.T) {
	require, ctx := require.New(t), context.Background()

	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()

	example := mexamplev1.NewMockExampleClient(t)
	examplev1xns.RegisterExampleActivities(env, example)
	external := examplev1.NewTestExternalClient(env, &Workflows{}, &Activities{})

	run, err := external.ProvisionFooAsync(ctx, &examplev1.ProvisionFooInput{RequestName: "test"})
	require.NoError(err)

	m := &sync.RWMutex{}
	var progress *examplev1.SetFooProgressRequest

	example.On("CreateFooAsync", mock.Anything, mock.Anything, mock.Anything).
		Return(func(ctx context.Context, req *examplev1.CreateFooInput, opts ...*examplev1.CreateFooOptions) (examplev1.CreateFooRun, error) {
			child := mexamplev1.NewMockCreateFooRun(t)
			child.On("Get", mock.Anything).
				Return(func(ctx context.Context) (*examplev1.CreateFooResponse, error) {
					for {
						m.RLock()
						p := progress.GetProgress()
						m.RUnlock()
						if p >= 100 {
							break
						}
						time.Sleep(time.Millisecond * 10)
					}
					return &examplev1.CreateFooResponse{
						Foo: &examplev1.Foo{
							Name:   req.GetRequestName(),
							Status: examplev1.Foo_FOO_STATUS_READY,
						},
					}, nil
				})
			return child, nil
		})

	example.On("SetFooProgress", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(func(ctx context.Context, workflowID, runID string, req *examplev1.SetFooProgressRequest) error {
			m.Lock()
			defer m.Unlock()
			progress = req
			return nil
		})

	example.On("GetFooProgress", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(func(ctx context.Context, workflowID, runID string) (*examplev1.GetFooProgressResponse, error) {
			m.RLock()
			defer m.RUnlock()
			p := progress.GetProgress()
			status := examplev1.Foo_FOO_STATUS_CREATING
			if p >= 100 {
				status = examplev1.Foo_FOO_STATUS_READY
			}
			return &examplev1.GetFooProgressResponse{
				Progress: p,
				Status:   status,
			}, nil
		})

	example.On("UpdateFooProgressAsync", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(func(ctx context.Context, workflowID, runID string, req *examplev1.SetFooProgressRequest, opts ...*examplev1.UpdateFooProgressOptions) (examplev1.UpdateFooProgressHandle, error) {
			m.Lock()
			defer m.Unlock()
			progress = req
			handle := mexamplev1.NewMockUpdateFooProgressHandle(t)
			handle.On("Get", mock.Anything).
				Return(func(ctx context.Context) (*examplev1.GetFooProgressResponse, error) {
					m.RLock()
					defer m.RUnlock()
					p := progress.GetProgress()
					status := examplev1.Foo_FOO_STATUS_CREATING
					if p >= 100 {
						status = examplev1.Foo_FOO_STATUS_READY
					}
					return &examplev1.GetFooProgressResponse{
						Progress: p,
						Status:   status,
					}, nil
				})
			handle.On("UpdateID").Return("foo")
			return handle, nil
		})

	resp, err := run.Get(ctx)
	require.NoError(err)
	require.Equal("test", resp.GetFoo().GetName())
	require.Equal(examplev1.Foo_FOO_STATUS_READY.String(), resp.GetFoo().GetStatus().String())
}
