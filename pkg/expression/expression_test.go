package expression_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	simplepb "github.com/cludden/protoc-gen-go-temporal/gen/simple"
	"github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	"github.com/stretchr/testify/require"
)

func TestExpression(t *testing.T) {
	require := require.New(t)

	cases := []struct {
		expr     string
		msg      *simplepb.SomeWorkflow1Request
		assert   func(result string, err error)
		expected string
		err      string
	}{
		{
			expr: "test/${!id}",
			msg: &simplepb.SomeWorkflow1Request{
				Id: "foo",
			},
			expected: "test/foo",
		},
		{
			expr: "test/${!id.uppercase()}",
			msg: &simplepb.SomeWorkflow1Request{
				Id: "foo",
			},
			expected: "test/FOO",
		},
		{
			expr: "test/${!intField.string()}",
			msg: &simplepb.SomeWorkflow1Request{
				IntField: 30,
			},
			expected: "test/30",
		},
		{
			expr:     `test/${!intField.or("unknown")}`,
			msg:      &simplepb.SomeWorkflow1Request{},
			expected: "test/unknown",
		},
		{
			expr: "test/${!bytesField}",
			msg: &simplepb.SomeWorkflow1Request{
				BytesField: []byte("foo"),
			},
			expected: fmt.Sprintf("test/%s", base64.StdEncoding.EncodeToString([]byte("foo"))),
		},
		{
			expr: `test/${!bytesField.decode("base64").string()}`,
			msg: &simplepb.SomeWorkflow1Request{
				BytesField: []byte("foo"),
			},
			expected: "test/foo",
		},
		{
			expr: `test/${! boolField.string() }`,
			msg: &simplepb.SomeWorkflow1Request{
				BoolField: true,
			},
			expected: "test/true",
		},
		{
			expr: `test/${!outerSingle.foo}/${!outerSingle.innerSingle.bar}`,
			msg: &simplepb.SomeWorkflow1Request{
				OuterSingle: &simplepb.SomeWorkflow1Request_OuterNested{
					Foo: "bar",
					InnerSingle: &simplepb.SomeWorkflow1Request_OuterNested_InnerNested{
						Bar: "baz",
					},
				},
			},
			expected: "test/bar/baz",
		},
		{
			expr: `test/${!outerList.0.foo}/${!outerList.0.innerList.0.bar}`,
			msg: &simplepb.SomeWorkflow1Request{
				OuterList: []*simplepb.SomeWorkflow1Request_OuterNested{{
					Foo: "bar",
					InnerList: []*simplepb.SomeWorkflow1Request_OuterNested_InnerNested{{
						Bar: "baz",
					}},
				}},
			},
			expected: "test/bar/baz",
		},
	}

	for _, c := range cases {
		// parse expression
		expr, err := expression.ParseExpression(c.expr)
		require.NoError(err)
		require.NotNil(expr)
		require.GreaterOrEqual(len(expr.Fragments), 1)

		actual, err := expression.EvalExpression(expr, c.msg.ProtoReflect())
		if c.err != "" {
			require.ErrorContains(err, c.err)
		} else if c.expected != "" {
			require.NoError(err)
			require.Equal(c.expected, actual)
		} else {
			c.assert(actual, err)
		}
	}
}
