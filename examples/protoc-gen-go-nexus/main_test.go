package main

import (
	"context"
	"testing"
	"time"

	examplepb "github.com/cludden/protoc-gen-go-temporal/gen/example/protocgengonexus"
	"github.com/cludden/protoc-gen-go-temporal/internal/testutil"
	"github.com/stretchr/testify/require"
	nexusv1 "go.temporal.io/api/nexus/v1"
	operatorservicev1 "go.temporal.io/api/operatorservice/v1"
	workflowservicev1 "go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
	"google.golang.org/protobuf/types/known/durationpb"
)

func TestE2E(t *testing.T) {
	// start dev server with nexus enabled
	defaultClient := testutil.StartExistingDevServerOrSkipNow(t, testsuite.DevServerOptions{
		ClientOptions: &client.Options{
			HostPort: "0.0.0.0:7233",
		},
		ExtraArgs: []string{
			"--http-port", "7243",
			"--dynamic-config-value", "system.enableNexus=true",
		},
	})
	ctx, r := context.Background(), require.New(t)

	// register greeting namespace
	_, err := defaultClient.WorkflowService().RegisterNamespace(ctx, &workflowservicev1.RegisterNamespaceRequest{
		Namespace:                        "greeting",
		WorkflowExecutionRetentionPeriod: durationpb.New(time.Hour * 24),
	})
	r.NoError(err)
	greetingClient, err := client.NewClientFromExisting(defaultClient, client.Options{Namespace: "greeting"})
	r.NoError(err)

	// start greeting worker
	greetingWorker := worker.New(greetingClient, examplepb.GreetingTaskQueue, worker.Options{})
	r.NoError(RegisterGreeting(greetingWorker))
	r.NoError(greetingWorker.Start())
	t.Cleanup(greetingWorker.Stop)

	// create enxus endpoint targetting greeting worker
	_, err = defaultClient.OperatorService().CreateNexusEndpoint(ctx, &operatorservicev1.CreateNexusEndpointRequest{
		Spec: &nexusv1.EndpointSpec{
			Name: "greeting",
			Target: &nexusv1.EndpointTarget{
				Variant: &nexusv1.EndpointTarget_Worker_{
					Worker: &nexusv1.EndpointTarget_Worker{
						Namespace: "greeting",
						TaskQueue: examplepb.GreetingTaskQueue,
					},
				},
			},
		},
	})
	r.NoError(err)

	// start caller worker
	callerWorker := worker.New(defaultClient, examplepb.CallerTaskQueue, worker.Options{})
	r.NoError(RegisterCaller(callerWorker, "greeting"))
	r.NoError(callerWorker.Start())
	t.Cleanup(callerWorker.Stop)

	// call greeting
	caller := examplepb.NewCallerClient(defaultClient)
	out, err := caller.CallGreet(ctx, &examplepb.CallGreetInput{Name: "test"})
	r.NoError(err)
	r.Equal("Hello, test!", out.GetGreeting())
}
