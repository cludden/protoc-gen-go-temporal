package main

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
	"go.temporal.io/sdk/client"
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
			if len(opts.SearchAttributes) != 2 {
				err = errors.Join(err, fmt.Errorf("expected 2 search attributes, got: %d", len(opts.SearchAttributes)))
			}
			if raw, ok := opts.SearchAttributes["foo"]; !ok {
				err = errors.Join(err, fmt.Errorf("expected search attributes to contain 'foo' attribute"))
			} else if v, ok := raw.(string); !ok {
				err = errors.Join(err, fmt.Errorf("expected 'foo' attribute to be string, got: %T", raw))
			} else if v != "bar" {
				err = errors.Join(err, fmt.Errorf("expected 'foo' to equal 'bar', got: %q", v))
			}
			if raw, ok := opts.SearchAttributes["created_at"]; !ok {
				err = errors.Join(err, fmt.Errorf("expected search attributes to contain 'created_at' attribute"))
			} else if v, ok := raw.(time.Time); !ok {
				err = errors.Join(err, fmt.Errorf("expected 'created_at' attribute to be string, got: %T", raw))
			} else if time.Since(v) > time.Second {
				err = errors.Join(err, fmt.Errorf("expected 'created_at' to be within 1s: %q", v))
			}
			if err != nil {
				return nil, err
			}
			return clientutils.NewMockWorkflowRun(t), nil
		},
	)
	example := examplev1.NewExampleClient(c)
	_, err := example.CreateFooAsync(ctx, &examplev1.CreateFooInput{RequestName: "bar"})
	require.NoError(err)
}
