package xnserr

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"
	"time"

	xnsv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/xns/v1"
	xnserrv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/xnserr/v1"
	"github.com/cludden/protoc-gen-go-temporal/gen/test/xnserr/v1/xnserrv1xns"
	xnserrv1mocks "github.com/cludden/protoc-gen-go-temporal/mocks/github.com/cludden/protoc-gen-go-temporal/gen/test/xnserr/v1"
	"github.com/cludden/protoc-gen-go-temporal/pkg/xns"
	"github.com/oklog/run"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/api/filter/v1"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
	"google.golang.org/protobuf/types/known/durationpb"
)

type XnsErrSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite

	cancel  func()
	c       client.Client
	client  xnserrv1.ClientClient
	ctx     context.Context
	doneCh  chan error
	g       *run.Group
	require *require.Assertions
	sc      client.Client
	server  xnserrv1.ServerClient
	srv     *testsuite.DevServer
}

func TestXnsErrSuite(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	suite.Run(t, new(XnsErrSuite))
}

func registerNamespaceIfNotExists(ctx context.Context, t *testing.T, c client.Client) {
	retention := time.Hour * 24

	// fetch all namespaces
	var namespaces []*workflowservice.DescribeNamespaceResponse
	res, err := c.WorkflowService().ListNamespaces(ctx, &workflowservice.ListNamespacesRequest{})
	require.NoError(t, err)
	namespaces = append(namespaces, res.Namespaces...)

	for len(res.NextPageToken) > 0 {
		res, err := c.WorkflowService().ListNamespaces(ctx, &workflowservice.ListNamespacesRequest{NextPageToken: res.NextPageToken})
		require.NoError(t, err)
		namespaces = append(namespaces, res.Namespaces...)
	}

	// check if we already have xnserr-server and if so return
	for _, n := range namespaces {
		if n.NamespaceInfo.Name == "xnserr-server" {
			return
		}
	}

	// since we don't have this ns let's create it
	_, err = c.WorkflowService().RegisterNamespace(ctx, &workflowservice.RegisterNamespaceRequest{Namespace: "xnserr-server", WorkflowExecutionRetentionPeriod: durationpb.New(retention)})
	require.NoError(t, err)
}

func (s *XnsErrSuite) SetupSuite() {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.require = s.Require()

	var err error
	s.srv, err = testsuite.StartDevServer(s.ctx, testsuite.DevServerOptions{
		ClientOptions: &client.Options{
			HostPort: "0.0.0.0:7233",
			Logger:   log.NewStructuredLogger(slog.New(slog.NewTextHandler(io.Discard, nil))),
			// ContextPropagators: []workflow.ContextPropagator{
			// 	patch.NewContextPropagator(),
			// },
		},
		EnableUI: true,
		LogLevel: "error",
	})
	s.require.NoError(err)
	s.T().Logf("dev server running at %s", s.srv.FrontendHostPort())

	s.c = s.srv.Client()

	registerNamespaceIfNotExists(s.ctx, s.T(), s.c)

	s.g = &run.Group{}
	s.g.Add(
		func() error {
			<-s.ctx.Done()
			return s.ctx.Err()
		},
		func(error) {
			s.cancel()
		},
	)

	s.sc, err = client.NewClientFromExisting(s.c, client.Options{
		Namespace: "xnserr-server",
		// ContextPropagators: []workflow.ContextPropagator{
		// 	patch.NewContextPropagator(),
		// },
	})
	s.require.NoError(err)

	server := worker.New(s.sc, xnserrv1.ServerTaskQueue, worker.Options{})
	xnserrv1.RegisterServerWorkflows(server, &ServerWorkflows{})
	s.g.Add(
		func() error {
			return server.Run(nil)
		},
		func(error) {
			server.Stop()
		},
	)
	s.server, err = xnserrv1.NewServerClientWithOptions(s.sc, client.Options{
		Namespace: "xnserr-server",
		// ContextPropagators: []workflow.ContextPropagator{
		// 	patch.NewContextPropagator(),
		// },
	})
	s.require.NoError(err)

	client := worker.New(s.c, xnserrv1.ClientTaskQueue, worker.Options{})
	xnserrv1.RegisterClientWorkflows(client, &ClientWorkflows{})
	xnserrv1xns.RegisterServerActivities(client, xnserrv1.NewServerClient(s.sc))
	s.g.Add(
		func() error {
			return client.Run(nil)
		},
		func(error) {
			client.Stop()
		},
	)
	s.client = xnserrv1.NewClientClient(s.c)

	s.T().Cleanup(func() {
		defer s.srv.Stop()
		defer s.c.Close()
		s.cancel()
		s.require.ErrorIs(<-s.doneCh, context.Canceled)
	})

	s.doneCh = make(chan error)
	go func() {
		s.doneCh <- s.g.Run()
	}()
}

func (s *XnsErrSuite) TestWorkflowExecutionError_ClientCanceled() {
	run, err := s.client.CallSleepAsync(s.ctx, &xnserrv1.CallSleepRequest{
		Sleep: durationpb.New(2 * time.Hour),
		StartWorkflowOptions: &xnsv1.StartWorkflowOptions{
			Id: "TestWorkflowExecutionError_Canceled_Server",
			RetryPolicy: &xnsv1.RetryPolicy{
				MaxInterval: durationpb.New(time.Second),
				MaxAttempts: 3,
			},
		},
	}, xnserrv1.NewCallSleepOptions().WithStartWorkflowOptions(client.StartWorkflowOptions{
		ID: "TestWorkflowExecutionError_Canceled_Client",
	}))
	s.require.NoError(err)

	go func() {
		<-time.After(time.Second * 3)
		s.require.NoError(s.client.CancelWorkflow(s.ctx, "TestWorkflowExecutionError_Canceled_Client", ""))
	}()

	err = run.Get(s.ctx)
	var cancelledErr *temporal.CanceledError
	s.require.ErrorContains(err, cancelledErr.Error())

	execs, err := s.sc.WorkflowService().ListClosedWorkflowExecutions(s.ctx, &workflowservice.ListClosedWorkflowExecutionsRequest{
		Namespace: "xnserr-server",
		Filters: &workflowservice.ListClosedWorkflowExecutionsRequest_ExecutionFilter{
			ExecutionFilter: &filter.WorkflowExecutionFilter{
				WorkflowId: "TestWorkflowExecutionError_Canceled_Server",
			},
		},
	})
	s.require.NoError(err)
	s.require.Eventually(func() bool {
		return len(execs.GetExecutions()) == 1
	}, time.Second*10, time.Second)
	s.require.Equal(enums.WORKFLOW_EXECUTION_STATUS_CANCELED, execs.GetExecutions()[0].GetStatus())
}

func (s *XnsErrSuite) TestWorkflowExecutionError_ServerCanceled() {
	run, err := s.client.CallSleepAsync(s.ctx, &xnserrv1.CallSleepRequest{
		Sleep: durationpb.New(time.Hour),
		StartWorkflowOptions: &xnsv1.StartWorkflowOptions{
			Id: "TestWorkflowExecutionError_Canceled_Server",
			RetryPolicy: &xnsv1.RetryPolicy{
				MaxInterval: durationpb.New(time.Second),
				MaxAttempts: 3,
			},
		},
	}, xnserrv1.NewCallSleepOptions().WithStartWorkflowOptions(client.StartWorkflowOptions{
		ID: "TestWorkflowExecutionError_Canceled_Client",
	}))
	s.require.NoError(err)

	go func() {
		<-time.After(time.Second * 3)
		s.require.NoError(s.server.CancelWorkflow(s.ctx, "TestWorkflowExecutionError_Canceled_Server", ""))
	}()

	err = run.Get(s.ctx)
	var cancelledErr *temporal.CanceledError
	s.require.ErrorAs(err, &cancelledErr)

	execs, err := s.sc.WorkflowService().ListClosedWorkflowExecutions(s.ctx, &workflowservice.ListClosedWorkflowExecutionsRequest{
		Namespace: "xnserr-server",
		Filters: &workflowservice.ListClosedWorkflowExecutionsRequest_ExecutionFilter{
			ExecutionFilter: &filter.WorkflowExecutionFilter{
				WorkflowId: "TestWorkflowExecutionError_Canceled_Server",
			},
		},
	})
	s.require.NoError(err)
	s.require.Len(execs.GetExecutions(), 1)
	s.require.Equal(enums.WORKFLOW_EXECUTION_STATUS_CANCELED, execs.GetExecutions()[0].GetStatus())
}

func (s *XnsErrSuite) TestWorkflowExecutionError_Terminated() {
	run, err := s.client.CallSleepAsync(s.ctx, &xnserrv1.CallSleepRequest{
		Sleep: durationpb.New(time.Hour),
		StartWorkflowOptions: &xnsv1.StartWorkflowOptions{
			Id: "TestWorkflowExecutionError_Terminated_Server",
			RetryPolicy: &xnsv1.RetryPolicy{
				MaxInterval: durationpb.New(time.Second),
				MaxAttempts: 3,
			},
		},
	}, xnserrv1.NewCallSleepOptions().WithStartWorkflowOptions(client.StartWorkflowOptions{
		ID: "TestWorkflowExecutionError_Terminated_Client",
	}))
	s.require.NoError(err)

	go func() {
		<-time.After(time.Second * 3)
		s.require.NoError(s.server.TerminateWorkflow(s.ctx, "TestWorkflowExecutionError_Terminated_Server", "", "test-termination"))
	}()

	err = run.Get(s.ctx)
	s.require.Error(err)

	terr := xns.Unwrap(err)
	s.require.NotNil(terr)
	s.require.Equal("TerminatedError", xns.Code(terr))
	s.require.True(xns.IsNonRetryable(err))

	execs, err := s.sc.WorkflowService().ListClosedWorkflowExecutions(s.ctx, &workflowservice.ListClosedWorkflowExecutionsRequest{
		Namespace: "xnserr-server",
		Filters: &workflowservice.ListClosedWorkflowExecutionsRequest_ExecutionFilter{
			ExecutionFilter: &filter.WorkflowExecutionFilter{
				WorkflowId: "TestWorkflowExecutionError_Terminated_Server",
			},
		},
	})
	s.require.NoError(err)
	s.require.Len(execs.GetExecutions(), 1)
	s.require.Equal(enums.WORKFLOW_EXECUTION_STATUS_TERMINATED, execs.GetExecutions()[0].GetStatus())
}

func (s *XnsErrSuite) TestWorkflowExecutionError_Timeout() {
	run, err := s.client.CallSleepAsync(s.ctx, &xnserrv1.CallSleepRequest{
		Sleep: durationpb.New(time.Hour),
		StartWorkflowOptions: &xnsv1.StartWorkflowOptions{
			Id:               "TestWorkflowExecutionError_Timeout_Server",
			ExecutionTimeout: durationpb.New(time.Second),
			RetryPolicy: &xnsv1.RetryPolicy{
				MaxInterval: durationpb.New(time.Second),
				MaxAttempts: 3,
			},
		},
	}, xnserrv1.NewCallSleepOptions().WithStartWorkflowOptions(client.StartWorkflowOptions{
		ID: "TestWorkflowExecutionError_Timeout_Client",
	}))
	s.require.NoError(err)

	err = run.Get(s.ctx)
	s.require.Error(err)

	terr := xns.Unwrap(err)
	s.require.NotNil(terr)
	s.require.Equal("TimeoutError", xns.Code(terr))
	s.require.True(xns.IsNonRetryable(err))

	execs, err := s.sc.WorkflowService().ListClosedWorkflowExecutions(s.ctx, &workflowservice.ListClosedWorkflowExecutionsRequest{
		Namespace: "xnserr-server",
		Filters: &workflowservice.ListClosedWorkflowExecutionsRequest_ExecutionFilter{
			ExecutionFilter: &filter.WorkflowExecutionFilter{
				WorkflowId: "TestWorkflowExecutionError_Timeout_Server",
			},
		},
	})
	s.require.NoError(err)
	s.require.Len(execs.GetExecutions(), 1)
	s.require.Equal(enums.WORKFLOW_EXECUTION_STATUS_TIMED_OUT, execs.GetExecutions()[0].GetStatus())
}

func (s *XnsErrSuite) TestWorkflowExecutionError_Application_NonRetryable() {
	run, err := s.client.CallSleepAsync(s.ctx, &xnserrv1.CallSleepRequest{
		Failure: &xnserrv1.Failure{
			Message:      "foo",
			Info:         xnserrv1.FailureInfo_FAILURE_INFO_APPLICATION_ERROR,
			NonRetryable: true,
		},
		StartWorkflowOptions: &xnsv1.StartWorkflowOptions{
			Id: "TestWorkflowExecutionError_Application_NonRetryable",
			RetryPolicy: &xnsv1.RetryPolicy{
				MaxInterval: durationpb.New(time.Second),
				MaxAttempts: 3,
			},
		},
	}, xnserrv1.NewCallSleepOptions().WithStartWorkflowOptions(client.StartWorkflowOptions{
		ID: "TestWorkflowExecutionError_Application_NonRetryable",
	}))
	s.require.NoError(err)

	err = run.Get(s.ctx)
	s.require.Error(err)

	terr := xns.Unwrap(err)
	s.require.NotNil(terr)
	s.require.Equal("SleepError", xns.Code(terr))
	s.require.True(xns.IsNonRetryable(err))

	execs, err := s.sc.WorkflowService().ListClosedWorkflowExecutions(s.ctx, &workflowservice.ListClosedWorkflowExecutionsRequest{
		Namespace: "xnserr-server",
		Filters: &workflowservice.ListClosedWorkflowExecutionsRequest_ExecutionFilter{
			ExecutionFilter: &filter.WorkflowExecutionFilter{
				WorkflowId: "TestWorkflowExecutionError_Application_NonRetryable",
			},
		},
	})
	s.require.NoError(err)
	s.require.Len(execs.GetExecutions(), 1)
}

func (s *XnsErrSuite) TestWorkflowExecutionError_Application_Retryable() {
	run, err := s.client.CallSleepAsync(s.ctx, &xnserrv1.CallSleepRequest{
		Failure: &xnserrv1.Failure{
			Message: "foo",
			Info:    xnserrv1.FailureInfo_FAILURE_INFO_APPLICATION_ERROR,
		},
		StartWorkflowOptions: &xnsv1.StartWorkflowOptions{
			Id: "TestWorkflowExecutionError_Application_Retryable",
			RetryPolicy: &xnsv1.RetryPolicy{
				MaxInterval: durationpb.New(time.Second),
				MaxAttempts: 3,
			},
		},
	}, xnserrv1.NewCallSleepOptions().WithStartWorkflowOptions(client.StartWorkflowOptions{
		ID: "TestWorkflowExecutionError_Application_Retryable",
	}))
	s.require.NoError(err)

	err = run.Get(s.ctx)
	s.require.Error(err)

	terr := xns.Unwrap(err)
	s.require.NotNil(terr)
	s.require.Equal("SleepError", xns.Code(terr))
	s.require.True(xns.IsNonRetryable(err))

	execs, err := s.sc.WorkflowService().ListClosedWorkflowExecutions(s.ctx, &workflowservice.ListClosedWorkflowExecutionsRequest{
		Namespace: "xnserr-server",
		Filters: &workflowservice.ListClosedWorkflowExecutionsRequest_ExecutionFilter{
			ExecutionFilter: &filter.WorkflowExecutionFilter{
				WorkflowId: "TestWorkflowExecutionError_Application_Retryable",
			},
		},
	})
	s.require.NoError(err)
	s.require.Len(execs.GetExecutions(), 3)
}

func TestClientStopped(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	require, ctx := require.New(t), context.Background()

	// start dev server
	srv, err := testsuite.StartDevServer(ctx, testsuite.DevServerOptions{
		ClientOptions: &client.Options{
			HostPort: "0.0.0.0:7233",
			Logger:   log.NewStructuredLogger(slog.New(slog.NewTextHandler(io.Discard, nil))),
		},
		EnableUI: true,
		LogLevel: "error",
	})
	require.NoError(err)
	t.Cleanup(func() {
		srv.Stop()
	})

	// create server namespace
	c := srv.Client()
	t.Cleanup(func() {
		c.Close()
	})
	registerNamespaceIfNotExists(ctx, t, c)

	// create server namespace client
	sc, err := client.NewClientFromExisting(c, client.Options{Namespace: "xnserr-server"})
	require.NoError(err)

	// initialize server worker
	serverWorker := worker.New(sc, xnserrv1.ServerTaskQueue, worker.Options{})
	xnserrv1.RegisterServerWorkflows(serverWorker, &ServerWorkflows{})
	require.NoError(serverWorker.Start())
	t.Cleanup(func() {
		serverWorker.Stop()
	})

	// initialize client worker
	clientWorker := worker.New(c, xnserrv1.ClientTaskQueue, worker.Options{})
	xnserrv1.RegisterClientWorkflows(clientWorker, &ClientWorkflows{})
	xnserrv1xns.RegisterServerActivities(clientWorker, xnserrv1.NewServerClient(sc))
	clientClient := xnserrv1.NewClientClient(c)
	require.NoError(clientWorker.Start())

	// start xns workflow with long enough sleep
	run, err := clientClient.CallSleepAsync(ctx, &xnserrv1.CallSleepRequest{
		Sleep: durationpb.New(time.Second * 10),
		StartWorkflowOptions: &xnsv1.StartWorkflowOptions{
			Id: "TestClientStopped_Server",
			RetryPolicy: &xnsv1.RetryPolicy{
				MaxInterval: durationpb.New(time.Second),
				MaxAttempts: 3,
			},
		},
	}, xnserrv1.NewCallSleepOptions().WithStartWorkflowOptions(client.StartWorkflowOptions{
		ID: "TestClientStopped_Client",
	}))
	require.NoError(err)

	// sleep briefly and then stop client worker
	<-time.After(time.Second * 3)
	clientWorker.Stop()

	// sleep briefly and then restart server worker
	<-time.After(time.Second * 3)
	clientWorker = worker.New(c, xnserrv1.ClientTaskQueue, worker.Options{})
	xnserrv1.RegisterClientWorkflows(clientWorker, &ClientWorkflows{})
	xnserrv1xns.RegisterServerActivities(clientWorker, xnserrv1.NewServerClient(sc))
	require.NoError(clientWorker.Start())
	t.Cleanup(func() {
		clientWorker.Stop()
	})

	// await workflow completion
	require.NoError(run.Get(ctx))

	// verify server workflow status
	execs, err := sc.WorkflowService().ListClosedWorkflowExecutions(ctx, &workflowservice.ListClosedWorkflowExecutionsRequest{
		Namespace: "xnserr-server",
		Filters: &workflowservice.ListClosedWorkflowExecutionsRequest_ExecutionFilter{
			ExecutionFilter: &filter.WorkflowExecutionFilter{
				WorkflowId: "TestClientStopped_Server",
			},
		},
	})
	require.NoError(err)
	require.Len(execs.GetExecutions(), 1)
	require.Equal(enums.WORKFLOW_EXECUTION_STATUS_COMPLETED, execs.GetExecutions()[0].GetStatus())

	// verify client workflow status
	execs, err = c.WorkflowService().ListClosedWorkflowExecutions(ctx, &workflowservice.ListClosedWorkflowExecutionsRequest{
		Namespace: "default",
		Filters: &workflowservice.ListClosedWorkflowExecutionsRequest_ExecutionFilter{
			ExecutionFilter: &filter.WorkflowExecutionFilter{
				WorkflowId: "TestClientStopped_Client",
			},
		},
	})
	require.NoError(err)
	require.Len(execs.GetExecutions(), 1)
	require.Equal(enums.WORKFLOW_EXECUTION_STATUS_COMPLETED, execs.GetExecutions()[0].GetStatus())

	// verify "WorkerStopped" error as last failure for xns activity
	var found bool
	cursor := c.GetWorkflowHistory(ctx, "TestClientStopped_Client", execs.GetExecutions()[0].GetExecution().GetRunId(), false, enums.HISTORY_EVENT_FILTER_TYPE_ALL_EVENT)
	for cursor.HasNext() {
		e, err := cursor.Next()
		require.NoError(err)
		if e.GetEventType() == enums.EVENT_TYPE_ACTIVITY_TASK_STARTED {
			found = true
			attrs := e.GetActivityTaskStartedEventAttributes()
			require.Equal(int32(2), attrs.GetAttempt())
			require.Equal("worker is stopping", attrs.GetLastFailure().GetMessage())
			require.Equal("WorkerStopped", attrs.GetLastFailure().GetApplicationFailureInfo().GetType())
			break
		}
	}
	require.True(found, "expected to find %s event in workflow history", enums.EVENT_TYPE_ACTIVITY_TASK_STARTED.String())
}

func TestUnhandledError(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	ctx, require := context.Background(), require.New(t)

	serverClient := xnserrv1mocks.NewMockServerClient(t)
	xnserrv1xns.RegisterServerActivities(env, serverClient)
	client := xnserrv1.NewTestClientClient(env, &ClientWorkflows{}, nil)

	run, err := client.CallSleepAsync(ctx, &xnserrv1.CallSleepRequest{
		RetryPolicy: &xnsv1.RetryPolicy{
			MaxAttempts: 2,
		},
	})
	require.NoError(err)

	serverClient.EXPECT().SleepAsync(mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, input *xnserrv1.SleepRequest, opts ...*xnserrv1.SleepOptions) (xnserrv1.SleepRun, error) {
			run := xnserrv1mocks.NewMockSleepRun(t)
			run.EXPECT().Get(mock.Anything).Return(errors.New("unhandled"))
			return run, nil
		})

	err = run.Get(ctx)
	var workflowExecutionErr *temporal.WorkflowExecutionError
	require.ErrorAs(err, &workflowExecutionErr)
	var activityErr *temporal.ActivityError
	require.ErrorAs(workflowExecutionErr.Unwrap(), &activityErr)
	terr := xns.Unwrap(activityErr)
	require.NotNil(terr)
	require.Equal("", terr.Type())
	require.Equal("unhandled", terr.Message())
	require.False(terr.NonRetryable())
}

func TestErrorConverter(t *testing.T) {
	var s testsuite.WorkflowTestSuite
	env := s.NewTestWorkflowEnvironment()
	ctx, require := context.Background(), require.New(t)

	serverClient := xnserrv1mocks.NewMockServerClient(t)
	xnserrv1xns.RegisterServerActivities(env, serverClient, xnserrv1xns.NewServerOptions().WithErrorConverter(func(err error) error {
		return temporal.NewApplicationError("uh oh", "OVERRIDDEN", err)
	}))
	client := xnserrv1.NewTestClientClient(env, &ClientWorkflows{}, nil)

	run, err := client.CallSleepAsync(ctx, &xnserrv1.CallSleepRequest{
		RetryPolicy: &xnsv1.RetryPolicy{
			MaxAttempts: 2,
		},
	})
	require.NoError(err)

	serverClient.EXPECT().SleepAsync(mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, input *xnserrv1.SleepRequest, opts ...*xnserrv1.SleepOptions) (xnserrv1.SleepRun, error) {
			run := xnserrv1mocks.NewMockSleepRun(t)
			run.EXPECT().Get(mock.Anything).Return(errors.New("unhandled"))
			return run, nil
		})

	err = run.Get(ctx)
	var workflowExecutionErr *temporal.WorkflowExecutionError
	require.ErrorAs(err, &workflowExecutionErr)
	var activityErr *temporal.ActivityError
	require.ErrorAs(workflowExecutionErr.Unwrap(), &activityErr)
	terr := xns.Unwrap(activityErr)
	require.NotNil(terr)
	require.Equal("OVERRIDDEN", terr.Type())
	require.Equal("uh oh", terr.Message())
	require.False(terr.NonRetryable())
}
