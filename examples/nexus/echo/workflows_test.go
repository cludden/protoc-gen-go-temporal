package echo

import (
	"context"
	"testing"

	"github.com/cludden/protoc-gen-go-temporal/examples/nexus/greeting"
	nexusv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/nexus/v1"
	"github.com/cludden/protoc-gen-go-temporal/gen/example/nexus/v1/nexusv1nexustemporal"
	"github.com/cludden/protoc-gen-go-temporal/internal/testutil"
	"github.com/stretchr/testify/require"
	"go.temporal.io/api/nexus/v1"
	"go.temporal.io/api/operatorservice/v1"
	"go.temporal.io/sdk/worker"
)

func TestE2E(t *testing.T) {
	c, ctx := testutil.NewIntegrationEnv(t, nil), context.Background()

	_, err := c.OperatorService().CreateNexusEndpoint(ctx, &operatorservice.CreateNexusEndpointRequest{
		Spec: &nexus.EndpointSpec{
			Name: "greeting",
			Target: &nexus.EndpointTarget{
				Variant: &nexus.EndpointTarget_Worker_{
					Worker: &nexus.EndpointTarget_Worker{
						Namespace: "default",
						TaskQueue: nexusv1.GreetingServiceTaskQueue,
					},
				},
			},
		},
	})
	require.NoError(t, err)

	greetingW := worker.New(c, nexusv1.GreetingServiceTaskQueue, worker.Options{})
	require.NoError(t, greeting.Register(greetingW))
	require.NoError(t, greetingW.Start())
	t.Cleanup(greetingW.Stop)

	echoW := worker.New(c, nexusv1.EchoServiceTaskQueue, worker.Options{})
	require.NoError(t, Register(echoW, nexusv1nexustemporal.NewGreetingServiceNexusClient("greeting")))
	require.NoError(t, echoW.Start())
	t.Cleanup(echoW.Stop)

	echo := nexusv1.NewEchoServiceClient(c)

	out, err := echo.Echo(ctx, &nexusv1.EchoInput{
		Name:     "World",
		Language: nexusv1.Language_LANGUAGE_SPANISH,
	})
	require.NoError(t, err)
	require.Equal(t, "¡Hola! World 👋", out.GetMessage())
}
