package expression_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	pb "github.com/cludden/protoc-gen-go-temporal/gen/test/expression/v1"
	"github.com/cludden/protoc-gen-go-temporal/pkg/expression"
	"github.com/stretchr/testify/require"
)

func TestExpression(t *testing.T) {
	require := require.New(t)

	cases := []struct {
		expr     string
		msg      *pb.Request
		assert   func(result string, err error)
		expected string
		err      string
	}{
		{
			expr: "test/${!id}",
			msg: &pb.Request{
				Id: "foo",
			},
			expected: "test/foo",
		},
		{
			expr: "test/${!id.uppercase()}",
			msg: &pb.Request{
				Id: "foo",
			},
			expected: "test/FOO",
		},
		{
			expr: "test/${!intField.string()}",
			msg: &pb.Request{
				IntField: 30,
			},
			expected: "test/30",
		},
		{
			expr:     `test/${!intField.or("unknown")}`,
			msg:      &pb.Request{},
			expected: "test/unknown",
		},
		{
			expr: `test/${!id.re_find_object("(?P<first>[^:]{3,10}):(?P<second>[^:]{3,10}):(?P<third>[^:]{3,10})").without("0").key_values().sort_by(pair -> pair.key).map_each(pair -> pair.value).join("/")}`,
			msg: &pb.Request{
				Id: "foo:bar:baz",
			},
			expected: "test/foo/bar/baz",
		},
		{
			expr: "test/${!bytesField}",
			msg: &pb.Request{
				BytesField: []byte("foo"),
			},
			expected: fmt.Sprintf("test/%s", base64.StdEncoding.EncodeToString([]byte("foo"))),
		},
		{
			expr: `test/${!bytesField.decode("base64").string()}`,
			msg: &pb.Request{
				BytesField: []byte("foo"),
			},
			expected: "test/foo",
		},
		{
			expr: `test/${! boolField.string() }`,
			msg: &pb.Request{
				BoolField: true,
			},
			expected: "test/true",
		},
		{
			expr: `test/${!outerSingle.foo}/${!outerSingle.innerSingle.bar}`,
			msg: &pb.Request{
				OuterSingle: &pb.Request_OuterNested{
					Foo: "bar",
					InnerSingle: &pb.Request_OuterNested_InnerNested{
						Bar: "baz",
					},
				},
			},
			expected: "test/bar/baz",
		},
		{
			expr: `test/${!outerList.0.foo}/${!outerList.0.innerList.0.bar}`,
			msg: &pb.Request{
				OuterList: []*pb.Request_OuterNested{{
					Foo: "bar",
					InnerList: []*pb.Request_OuterNested_InnerNested{{
						Bar: "baz",
					}},
				}},
			},
			expected: "test/bar/baz",
		},
		{
			expr: `test/${! ["svc", "region", "acc", "resource"].map_each(k -> id.re_find_object("arn:aws:(?P<svc>.+):(?P<region>.+):(?P<acc>.+):(?P<resource>.+)").get(k)).join("/") }`,
			msg: &pb.Request{
				Id: "arn:aws:ec2:us-east-1:123456789012:vpc/vpc-0e9801d129EXAMPLE",
			},
			expected: "test/ec2/us-east-1/123456789012/vpc/vpc-0e9801d129EXAMPLE",
		},
		{
			expr: `test/${! if id.contains("z") { id.uppercase() } else { id.lowercase() } }/${! id }`,
			msg: &pb.Request{
				Id: "zeb",
			},
			expected: "test/ZEB/zeb",
		},
		{
			expr: `this_is_a_\${!test/${! match id { "aaa" => "A", "bbb" => "B", _ => "Z" } }`,
			msg: &pb.Request{
				Id: "zeb",
			},
			expected: `this_is_a_${!test/Z`,
		},
		{
			expr: `${! id.re_find_object(".*\{(?P<g>[^\}]+)\}.*").g }`,
			msg: &pb.Request{
				Id: "foo{bar}baz",
			},
			expected: "bar",
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

func TestExpression_Error(t *testing.T) {
	require := require.New(t)

	cases := []struct {
		expr   string
		msg    *pb.Request
		errors []string
	}{
		{
			expr:   "test/${!blah}",
			msg:    &pb.Request{},
			errors: []string{"expected string result from `blah` query, got: <nil>"},
		},
		{
			expr: "test/${!blah}/${!blahz}",
			msg:  &pb.Request{},
			errors: []string{
				"expected string result from `blah` query, got: <nil>",
				"expected string result from `blahz` query, got: <nil>",
			},
		},
		{
			expr: `test/${!blah}/${!blahz.or(throw("uh oh"))}`,
			msg:  &pb.Request{},
			errors: []string{
				"expected string result from `blah` query, got: <nil>",
				"uh oh",
			},
		},
	}

	for _, c := range cases {
		// parse expression
		expr, err := expression.ParseExpression(c.expr)
		require.NoError(err)
		require.NotNil(expr)
		require.GreaterOrEqual(len(expr.Fragments), 1)

		_, err = expression.EvalExpression(expr, c.msg.ProtoReflect())
		for _, msg := range c.errors {
			require.ErrorContains(err, msg)
		}
	}
}
