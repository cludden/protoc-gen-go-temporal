package patch

import (
	"github.com/cludden/protoc-gen-go-temporal/pkg/ctxkey"
	"github.com/cludden/protoc-gen-go-temporal/pkg/ctxpropagation"
	"go.temporal.io/sdk/workflow"
)

const (
	PV_64_ExpressionEvaluationLocalActivity = "cludden_protoc-gen-go-temporal_64_expression-evaluation-local-activity"
	PV_77_UseParentTaskQueue                = "cludden_protoc-gen-go-temporal_77_use-parent-task-queue"
)

var Keys = []ctxpropagation.Key{
	defaultTaskQueues,
}

var (
	defaultTaskQueues = ctxkey.Key[map[string]string]("github.com/cludden/protoc-gen-go-temporal#DefaultTaskQueues", true)
)

func NewContextPropagator() workflow.ContextPropagator {
	return ctxpropagation.New(ctxpropagation.WithKeys(Keys...))
}

func DefaultTaskQueue(ctx ctxkey.Context, tq string) string {
	v := DefaultTaskQueues(ctx)
	if vv, ok := v[tq]; !ok {
		return tq
	} else if vv == "" {
		return tq
	} else {
		return vv
	}
}

func DefaultTaskQueues(ctx ctxkey.Context) map[string]string {
	v, _ := ctx.Value(defaultTaskQueues).(map[string]string)
	if v == nil {
		v = map[string]string{}
	}
	return v
}

func WithDefaultTaskQueue[K ctxkey.Context, Setter func(K, any, any) K](set Setter, ctx K, taskQueues ...string) K {
	if len(taskQueues) == 0 {
		return ctx
	}
	tq := taskQueues[0]
	var override string
	if len(taskQueues) == 2 && taskQueues[1] != tq {
		override = taskQueues[1]
	}
	v, _ := ctx.Value(defaultTaskQueues).(map[string]string)
	if v == nil {
		v = map[string]string{}
	}
	if _, ok := v[tq]; !ok {
		v[tq] = override
		if override != "" {
			if _, ok = v[override]; !ok {
				v[override] = ""
			}
		}
	}
	return set(ctx, defaultTaskQueues, v)
}
