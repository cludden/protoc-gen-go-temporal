package testutil

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/hairyhenderson/go-which"
	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/testsuite"
)

func NewIntegrationEnv(t *testing.T, opts *testsuite.DevServerOptions) client.Client {
	t.Helper()
	var existingPath string
	if opts != nil && opts.ExistingPath == "" {
		existingPath = which.Which("temporal")
		if existingPath == "" {
			t.Skip("temporal CLI not found in PATH, skipping integration tests")
		}
	}

	if opts == nil {
		opts = &testsuite.DevServerOptions{}
	}
	if opts.ExistingPath == "" {
		opts.ExistingPath = existingPath
	}
	if opts.ClientOptions == nil {
		opts.ClientOptions = &client.Options{}
	}
	if opts.ClientOptions.Logger == nil {
		opts.ClientOptions.Logger = log.NewStructuredLogger(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})))
	}
	opts.EnableUI = true

	srv, err := testsuite.StartDevServer(context.Background(), *opts)
	if err != nil {
		t.Fatalf("failed to start dev server: %v", err)
	}
	t.Cleanup(func() { assert.NoError(t, srv.Stop(), "unexpected error stopping dev server") })

	c := srv.Client()
	t.Cleanup(c.Close)
	return c
}
