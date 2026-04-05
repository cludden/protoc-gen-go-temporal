package clienttaskqueue

import (
	"context"
	"errors"
	"testing"

	clienttaskqueuepb "github.com/cludden/protoc-gen-go-temporal/gen/test/client-task-queue/v1"
	clientmocks "github.com/cludden/protoc-gen-go-temporal/mocks/go.temporal.io/sdk/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
)

func TestClientTaskQueue(t *testing.T) {
	cases := map[string]struct {
		assert func(*testing.T, *clientmocks.MockClient)
	}{
		"execute workflow with client task queue takes precedence over service, workflow task queue": {
			assert: func(t *testing.T, c *clientmocks.MockClient) {
				sc := clienttaskqueuepb.NewClientTaskQueueServiceClient(
					c, clienttaskqueuepb.NewClientTaskQueueServiceClientOptions().WithTaskQueue("baz"),
				)
				c.EXPECT().
					ExecuteWorkflow(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					RunAndReturn(func(
						ctx context.Context,
						options client.StartWorkflowOptions,
						workflow any,
						args ...any,
					) (client.WorkflowRun, error) {
						if !assert.Equal(t, "baz", options.TaskQueue) {
							return nil, errors.New("fail")
						}
						run := clientmocks.NewMockWorkflowRun(t)
						return run, nil
					})
				_, err := sc.FooAsync(context.Background(), &clienttaskqueuepb.FooInput{
					Id: "foo",
				})
				require.NoError(t, err)
			},
		},
		"signal with start workflow with client task queue takes precedence over service, workflow task queue": {
			assert: func(t *testing.T, c *clientmocks.MockClient) {
				sc := clienttaskqueuepb.NewClientTaskQueueServiceClient(
					c, clienttaskqueuepb.NewClientTaskQueueServiceClientOptions().WithTaskQueue("baz"),
				)
				c.EXPECT().
					SignalWithStartWorkflow(
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
					).
					RunAndReturn(func(
						ctx context.Context,
						workflowID string,
						signalName string,
						signalArg any,
						options client.StartWorkflowOptions,
						workflow any,
						workflowArgs ...any,
					) (client.WorkflowRun, error) {
						if !assert.Equal(t, "baz", options.TaskQueue) {
							return nil, errors.New("fail")
						}
						run := clientmocks.NewMockWorkflowRun(t)
						return run, nil
					})
				_, err := sc.FooWithBarAsync(
					context.Background(),
					&clienttaskqueuepb.FooInput{Id: "1"},
					&clienttaskqueuepb.BarInput{Id: "2"},
				)
				require.NoError(t, err)
			},
		},
		"update with start workflow with client task queue takes precedence over service, workflow task queue": {
			assert: func(t *testing.T, c *clientmocks.MockClient) {
				sc := clienttaskqueuepb.NewClientTaskQueueServiceClient(
					c, clienttaskqueuepb.NewClientTaskQueueServiceClientOptions().WithTaskQueue("baz"),
				)
				c.EXPECT().
					GetWorkflow(mock.Anything, mock.Anything, mock.Anything).
					RunAndReturn(func(ctx context.Context, workflowID, runID string) client.WorkflowRun {
						return nil
					})
				c.EXPECT().
					NewWithStartWorkflowOperation(mock.Anything, mock.Anything, mock.Anything).
					RunAndReturn(func(
						options client.StartWorkflowOptions,
						workflow any,
						args ...any,
					) client.WithStartWorkflowOperation {
						run := clientmocks.NewMockWithStartWorkflowOperation(t)
						assert.Equal(t, "baz", options.TaskQueue)
						return run
					})
				c.EXPECT().
					UpdateWithStartWorkflow(mock.Anything, mock.Anything).
					RunAndReturn(func(
						ctx context.Context,
						options client.UpdateWithStartWorkflowOptions,
					) (client.WorkflowUpdateHandle, error) {
						h := clientmocks.NewMockWorkflowUpdateHandle(t)
						h.EXPECT().WorkflowID().Return("test-workflow-id")
						h.EXPECT().RunID().Return("test-run-id")
						return h, nil
					})
				_, _, err := sc.FooWithBazAsync(
					context.Background(),
					&clienttaskqueuepb.FooInput{Id: "1"},
					&clienttaskqueuepb.BazInput{Id: "2"},
				)
				require.NoError(t, err)
			},
		},
		"execute workflow with client task queue takes precedence over workflow task queue": {
			assert: func(t *testing.T, c *clientmocks.MockClient) {
				sc := clienttaskqueuepb.NewClientTaskQueueServiceClient(
					c, clienttaskqueuepb.NewClientTaskQueueServiceClientOptions().WithTaskQueue("baz"),
				)
				c.EXPECT().
					ExecuteWorkflow(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					RunAndReturn(func(
						ctx context.Context,
						options client.StartWorkflowOptions,
						workflow any,
						args ...any,
					) (client.WorkflowRun, error) {
						if !assert.Equal(t, "baz", options.TaskQueue) {
							return nil, errors.New("fail")
						}
						run := clientmocks.NewMockWorkflowRun(t)
						return run, nil
					})
				_, err := sc.QuxAsync(context.Background(), &clienttaskqueuepb.QuxInput{
					Id: "foo",
				})
				require.NoError(t, err)
			},
		},
		"signal with start workflow with client task queue takes precedence over workflow task queue": {
			assert: func(t *testing.T, c *clientmocks.MockClient) {
				sc := clienttaskqueuepb.NewClientTaskQueueServiceClient(
					c, clienttaskqueuepb.NewClientTaskQueueServiceClientOptions().WithTaskQueue("baz"),
				)
				c.EXPECT().
					SignalWithStartWorkflow(
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
					).
					RunAndReturn(func(
						ctx context.Context,
						workflowID string,
						signalName string,
						signalArg any,
						options client.StartWorkflowOptions,
						workflow any,
						workflowArgs ...any,
					) (client.WorkflowRun, error) {
						if !assert.Equal(t, "baz", options.TaskQueue) {
							return nil, errors.New("fail")
						}
						run := clientmocks.NewMockWorkflowRun(t)
						return run, nil
					})
				_, err := sc.QuxWithBarAsync(
					context.Background(),
					&clienttaskqueuepb.QuxInput{Id: "1"},
					&clienttaskqueuepb.BarInput{Id: "2"},
				)
				require.NoError(t, err)
			},
		},
		"update with start workflow with client task queue takes precedence over workflow task queue": {
			assert: func(t *testing.T, c *clientmocks.MockClient) {
				sc := clienttaskqueuepb.NewClientTaskQueueServiceClient(
					c, clienttaskqueuepb.NewClientTaskQueueServiceClientOptions().WithTaskQueue("baz"),
				)
				c.EXPECT().
					GetWorkflow(mock.Anything, mock.Anything, mock.Anything).
					RunAndReturn(func(ctx context.Context, workflowID, runID string) client.WorkflowRun {
						return nil
					})
				c.EXPECT().
					NewWithStartWorkflowOperation(mock.Anything, mock.Anything, mock.Anything).
					RunAndReturn(func(
						options client.StartWorkflowOptions,
						workflow any,
						args ...any,
					) client.WithStartWorkflowOperation {
						run := clientmocks.NewMockWithStartWorkflowOperation(t)
						assert.Equal(t, "baz", options.TaskQueue)
						return run
					})
				c.EXPECT().
					UpdateWithStartWorkflow(mock.Anything, mock.Anything).
					RunAndReturn(func(
						ctx context.Context,
						options client.UpdateWithStartWorkflowOptions,
					) (client.WorkflowUpdateHandle, error) {
						h := clientmocks.NewMockWorkflowUpdateHandle(t)
						h.EXPECT().WorkflowID().Return("test-workflow-id")
						h.EXPECT().RunID().Return("test-run-id")
						return h, nil
					})
				_, _, err := sc.QuxWithBazAsync(
					context.Background(),
					&clienttaskqueuepb.QuxInput{Id: "1"},
					&clienttaskqueuepb.BazInput{Id: "2"},
				)
				require.NoError(t, err)
			},
		},
		"execute workflow invocation override takes precedence over client task queue": {
			assert: func(t *testing.T, c *clientmocks.MockClient) {
				sc := clienttaskqueuepb.NewClientTaskQueueServiceClient(
					c, clienttaskqueuepb.NewClientTaskQueueServiceClientOptions().WithTaskQueue("baz"),
				)
				c.EXPECT().
					ExecuteWorkflow(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					RunAndReturn(func(
						ctx context.Context,
						options client.StartWorkflowOptions,
						workflow any,
						args ...any,
					) (client.WorkflowRun, error) {
						if !assert.Equal(t, "qux", options.TaskQueue) {
							return nil, errors.New("fail")
						}
						run := clientmocks.NewMockWorkflowRun(t)
						return run, nil
					})
				_, err := sc.QuxAsync(context.Background(), &clienttaskqueuepb.QuxInput{
					Id: "foo",
				}, clienttaskqueuepb.NewQuxOptions().WithTaskQueue("qux"))
				require.NoError(t, err)
			},
		},
		"signal with start workflow invocation override takes precedence over client task queue": {
			assert: func(t *testing.T, c *clientmocks.MockClient) {
				sc := clienttaskqueuepb.NewClientTaskQueueServiceClient(
					c, clienttaskqueuepb.NewClientTaskQueueServiceClientOptions().WithTaskQueue("baz"),
				)
				c.EXPECT().
					SignalWithStartWorkflow(
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
					).
					RunAndReturn(func(
						ctx context.Context,
						workflowID string,
						signalName string,
						signalArg any,
						options client.StartWorkflowOptions,
						workflow any,
						workflowArgs ...any,
					) (client.WorkflowRun, error) {
						if !assert.Equal(t, "qux", options.TaskQueue) {
							return nil, errors.New("fail")
						}
						run := clientmocks.NewMockWorkflowRun(t)
						return run, nil
					})
				_, err := sc.QuxWithBarAsync(
					context.Background(),
					&clienttaskqueuepb.QuxInput{Id: "1"},
					&clienttaskqueuepb.BarInput{Id: "2"},
					clienttaskqueuepb.NewQuxOptions().WithTaskQueue("qux"),
				)
				require.NoError(t, err)
			},
		},
		"update with start workflow invocation override takes precedence over client task queue": {
			assert: func(t *testing.T, c *clientmocks.MockClient) {
				sc := clienttaskqueuepb.NewClientTaskQueueServiceClient(
					c, clienttaskqueuepb.NewClientTaskQueueServiceClientOptions().WithTaskQueue("baz"),
				)
				c.EXPECT().
					GetWorkflow(mock.Anything, mock.Anything, mock.Anything).
					RunAndReturn(func(ctx context.Context, workflowID, runID string) client.WorkflowRun {
						return nil
					})
				c.EXPECT().
					NewWithStartWorkflowOperation(mock.Anything, mock.Anything, mock.Anything).
					RunAndReturn(func(
						options client.StartWorkflowOptions,
						workflow any,
						args ...any,
					) client.WithStartWorkflowOperation {
						run := clientmocks.NewMockWithStartWorkflowOperation(t)
						assert.Equal(t, "qux", options.TaskQueue)
						return run
					})
				c.EXPECT().
					UpdateWithStartWorkflow(mock.Anything, mock.Anything).
					RunAndReturn(func(
						ctx context.Context,
						options client.UpdateWithStartWorkflowOptions,
					) (client.WorkflowUpdateHandle, error) {
						h := clientmocks.NewMockWorkflowUpdateHandle(t)
						h.EXPECT().WorkflowID().Return("test-workflow-id")
						h.EXPECT().RunID().Return("test-run-id")
						return h, nil
					})
				_, _, err := sc.QuxWithBazAsync(
					context.Background(),
					&clienttaskqueuepb.QuxInput{Id: "1"},
					&clienttaskqueuepb.BazInput{Id: "2"},
					clienttaskqueuepb.NewQuxWithBazOptions().WithQuxOptions(
						clienttaskqueuepb.NewQuxOptions().WithTaskQueue("qux"),
					),
				)
				require.NoError(t, err)
			},
		},
	}

	for _, name := range workflow.DeterministicKeys(cases) {
		tc := cases[name]
		t.Run(name, func(t *testing.T) {
			c := clientmocks.NewMockClient(t)
			tc.assert(t, c)
		})
	}
}
