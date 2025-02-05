package testutil

import (
	"context"
	"testing"

	"github.com/hairyhenderson/go-which"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/testsuite"
)

func StartExistingDevServer(t *testing.T, opts ...testsuite.DevServerOptions) (client.Client, bool) {
	existingPath := which.Which("temporal")
	if existingPath == "" {
		return nil, false
	}

	var o testsuite.DevServerOptions
	if len(opts) > 0 {
		o = opts[0]
	}
	o.ExistingPath = existingPath
	if o.ClientOptions == nil {
		o.ClientOptions = &client.Options{}
	}

	srv, err := testsuite.StartDevServer(context.Background(), o)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, srv.Stop()) })

	c := srv.Client()
	t.Cleanup(c.Close)
	return c, true
}

func StartExistingDevServerOrSkipNow(t *testing.T, opts ...testsuite.DevServerOptions) client.Client {
	c, ok := StartExistingDevServer(t, opts...)
	if !ok {
		t.Skipf("temporal cli not available in path")
		return nil
	}
	return c
}
