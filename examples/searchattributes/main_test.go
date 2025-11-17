package main

import (
	"context"
	"testing"
	"time"

	searchattributesv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/searchattributes/v1"
	"github.com/hairyhenderson/go-which"
	"github.com/stretchr/testify/require"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/api/operatorservice/v1"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestTypedSearchAttributes(t *testing.T) {
	existingPath := which.Which("temporal")
	if existingPath == "" {
		t.Skip("temporal binary not found in PATH, skipping test")
	}
	srv, err := testsuite.StartDevServer(context.Background(), testsuite.DevServerOptions{})
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, srv.Stop()) })

	c := srv.Client()
	_, err = c.OperatorService().AddSearchAttributes(context.Background(), &operatorservice.AddSearchAttributesRequest{
		Namespace: "default",
		SearchAttributes: map[string]enums.IndexedValueType{
			"CustomKeywordField":     enums.INDEXED_VALUE_TYPE_KEYWORD,
			"CustomKeywordListField": enums.INDEXED_VALUE_TYPE_KEYWORD_LIST,
			"CustomTextField":        enums.INDEXED_VALUE_TYPE_TEXT,
			"CustomIntField":         enums.INDEXED_VALUE_TYPE_INT,
			"CustomDoubleField":      enums.INDEXED_VALUE_TYPE_DOUBLE,
			"CustomBoolField":        enums.INDEXED_VALUE_TYPE_BOOL,
			"CustomDatetimeField":    enums.INDEXED_VALUE_TYPE_DATETIME,
		},
	})
	require.NoError(t, err)

	client := searchattributesv1.NewExampleClient(c)
	w := worker.New(c, searchattributesv1.ExampleTaskQueue, worker.Options{})
	searchattributesv1.RegisterExampleWorkflows(w, &Workflows{})
	require.NoError(t, w.Start())
	t.Cleanup(w.Stop)

	now := time.Now().UTC()
	out, err := client.TypedSearchAttributes(context.Background(), &searchattributesv1.TypedSearchAttributesInput{
		CustomKeywordField:     "test",
		CustomTextField:        "test",
		CustomIntField:         123,
		CustomDoubleField:      123.456,
		CustomBoolField:        true,
		CustomDatetimeField:    timestamppb.New(now),
		CustomKeywordListField: []string{"test"},
	})
	require.NoError(t, err)
	require.Equal(t, "test", out.CustomKeywordField)
	require.Equal(t, "test", out.CustomTextField)
	require.Equal(t, int64(123), out.CustomIntField)
	require.Equal(t, 123.456, out.CustomDoubleField)
	require.Equal(t, true, out.CustomBoolField)
	require.Equal(t, now, out.CustomDatetimeField.AsTime())
	require.Equal(t, []string{"test"}, out.CustomKeywordListField)
}
