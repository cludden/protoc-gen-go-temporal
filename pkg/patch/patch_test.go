package patch

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultTaskQueue(t *testing.T) {
	r, ctx := require.New(t), context.Background()

	r.Equal("foo", DefaultTaskQueue(ctx, "foo"))
	ctx = WithDefaultTaskQueue(context.WithValue, ctx, "foo")
	r.Equal("foo", DefaultTaskQueue(ctx, "foo"))
	ctx = WithDefaultTaskQueue(context.WithValue, ctx, "foo", "bar")
	r.Equal("foo", DefaultTaskQueue(ctx, "foo"))

	r.Equal("bar", DefaultTaskQueue(ctx, "bar"))
	ctx = WithDefaultTaskQueue(context.WithValue, ctx, "bar", "baz")
	r.Equal("baz", DefaultTaskQueue(ctx, "bar"))
	ctx = WithDefaultTaskQueue(context.WithValue, ctx, "foo", "qux")
	r.Equal("baz", DefaultTaskQueue(ctx, "bar"))
	r.Equal("foo", DefaultTaskQueue(ctx, "foo"))

	r.Equal("baz", DefaultTaskQueue(ctx, "baz"))
	ctx = WithDefaultTaskQueue(context.WithValue, ctx, "baz", "baz")
	r.Equal("baz", DefaultTaskQueue(ctx, "baz"), "baz")
	v, _ := ctx.Value(defaultTaskQueues).(map[string]string)
	r.Equal(map[string]string{
		"foo": "",
		"bar": "baz",
		"baz": "",
	}, v)
}
