package example

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"
	clientMock "github.com/cludden/protoc-gen-go-temporal/mocks/go.temporal.io/sdk/client"
	"github.com/cludden/protoc-gen-go-temporal/mocks/go.temporal.io/sdk/clientutils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/api/operatorservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
)

func TestCreateFooStartWorkflowOptions(t *testing.T) {
	ctx, require := context.Background(), require.New(t)

	c := clientMock.NewMockClient(t)
	c.On("ExecuteWorkflow", mock.Anything, mock.Anything, examplev1.CreateFooWorkflowName, mock.Anything).Return(
		func(ctx context.Context, opts client.StartWorkflowOptions, workflow any, params ...any) (run client.WorkflowRun, err error) {
			if opts.ID != "create-foo/bar" {
				err = errors.Join(err, fmt.Errorf("expected workflow id to equal 'create-foo/bar', got: %q", opts.ID))
			}
			if opts.WorkflowExecutionTimeout != time.Hour {
				err = errors.Join(err, fmt.Errorf("expected workflow execution timeout to equal 1h, got: %s", opts.WorkflowExecutionTimeout))
			}
			if err != nil {
				return nil, err
			}
			return clientutils.NewMockWorkflowRun(t), nil
		},
	)
	example := examplev1.NewExampleClient(c)
	_, err := example.CreateFooAsync(ctx, &examplev1.CreateFooRequest{Name: "bar"})
	require.NoError(err)
}

func TestUpdate(t *testing.T) {
	var suite testsuite.WorkflowTestSuite
	env := suite.NewTestWorkflowEnvironment()
	require, ctx := require.New(t), context.Background()
	client := examplev1.NewTestExampleClient(env, &Workflows{}, &Activities{})

	run, err := client.CreateFooAsync(ctx, &examplev1.CreateFooRequest{Name: "test"})
	require.NoError(err)

	env.OnActivity(examplev1.NotifyActivityName, mock.Anything, mock.Anything).Return(nil)

	env.RegisterDelayedCallback(func() {
		require.NoError(run.SetFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 5.7}))
	}, time.Second)

	env.RegisterDelayedCallback(func() {
		progress, err := run.GetFooProgress(ctx)
		require.NoError(err)
		require.Equal(float32(5.7), progress.GetProgress())
		require.Equal(examplev1.Foo_FOO_STATUS_CREATING, progress.GetStatus())
	}, time.Second*30)

	var handle examplev1.UpdateFooProgressHandle
	env.RegisterDelayedCallback(func() {
		handle, err = run.UpdateFooProgressAsync(ctx, &examplev1.SetFooProgressRequest{Progress: 100})
		require.NoError(err)
	}, time.Second*45)

	out, err := run.Get(ctx)
	require.NoError(err)
	require.Equal("test", out.GetFoo().GetName())
	require.Equal(examplev1.Foo_FOO_STATUS_READY, out.GetFoo().GetStatus())

	progress, err := handle.Get(ctx)
	require.NoError(err)
	require.Equal(float32(100), progress.GetProgress())
	require.Equal(examplev1.Foo_FOO_STATUS_READY, progress.GetStatus())
}

func TestUpdateIntegration(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	require, ctx := require.New(t), context.Background()

	srv, err := testsuite.StartDevServer(ctx, testsuite.DevServerOptions{
		LogLevel: "fatal",
		ExtraArgs: []string{
			"--dynamic-config-value", "frontend.enableUpdateWorkflowExecution=true",
			"--dynamic-config-value", "frontend.enableUpdateWorkflowExecutionAsyncAccepted=true",
		},
	})
	require.NoError(err)
	defer srv.Stop()

	c := srv.Client()
	example := examplev1.NewExampleClient(c)

	_, err = c.OperatorService().AddSearchAttributes(ctx, &operatorservice.AddSearchAttributesRequest{
		Namespace: "default",
		SearchAttributes: map[string]enums.IndexedValueType{
			"foo":        enums.INDEXED_VALUE_TYPE_TEXT,
			"created_at": enums.INDEXED_VALUE_TYPE_DATETIME,
		},
	})
	require.NoError(err)

	w := worker.New(c, examplev1.ExampleTaskQueue, worker.Options{})
	examplev1.RegisterExampleActivities(w, &Activities{})
	examplev1.RegisterExampleWorkflows(w, &Workflows{})
	require.NoError(w.Start())
	defer w.Stop()
	defer c.Close()

	run, err := example.CreateFooAsync(ctx, &examplev1.CreateFooRequest{Name: "test"})
	require.NoError(err)

	require.NoError(run.SetFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 5.7}))

	query, err := run.GetFooProgress(ctx)
	require.NoError(err)
	require.Equal(float32(5.7), query.GetProgress())
	require.Equal(examplev1.Foo_FOO_STATUS_CREATING.String(), query.GetStatus().String())

	handle, err := run.UpdateFooProgressAsync(ctx, &examplev1.SetFooProgressRequest{Progress: 100})
	require.NoError(err)

	update, err := handle.Get(ctx)
	require.NoError(err)
	require.Equal(float32(100), update.GetProgress())
	require.Equal(examplev1.Foo_FOO_STATUS_READY.String(), update.GetStatus().String())

	resp, err := run.Get(ctx)
	require.NoError(err)
	require.Equal(examplev1.Foo_FOO_STATUS_READY.String(), resp.GetFoo().GetStatus().String())
}
