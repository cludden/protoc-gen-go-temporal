package main

import (
	"context"
	"fmt"

	patchv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/patch/v1"
	"github.com/cludden/protoc-gen-go-temporal/pkg/patch"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

type (
	Pv77Workflows struct{}

	Pv77Activities struct{}

	Pv77FooWorkflow struct {
		*Pv77Workflows
		*patchv1.Pv77FooWorkflowInput
	}

	Pv77BarWorkflow struct {
		*Pv77Workflows
		*patchv1.Pv77BarWorkflowInput
	}

	Pv77BazWorkflow struct {
		*Pv77Workflows
		*patchv1.Pv77BazWorkflowInput
	}

	Pv77QuxWorkflow struct {
		*Pv77Workflows
		*patchv1.Pv77QuxWorkflowInput
	}
)

func (w *Pv77Workflows) Pv77Foo(ctx workflow.Context, input *patchv1.Pv77FooWorkflowInput) (patchv1.Pv77FooWorkflow, error) {
	return &Pv77FooWorkflow{w, input}, nil
}

func (w *Pv77FooWorkflow) Execute(ctx workflow.Context) (*patchv1.Pv77FooOutput, error) {
	out, err := w.execute(ctx, &patchv1.Pv77Input{Next: w.Req.GetNext()})
	return &patchv1.Pv77FooOutput{TaskQueues: out.GetTaskQueues(), Defaults: out.GetDefaults()}, err
}

func (w *Pv77Workflows) Pv77Bar(ctx workflow.Context, input *patchv1.Pv77BarWorkflowInput) (patchv1.Pv77BarWorkflow, error) {
	return &Pv77BarWorkflow{w, input}, nil
}

func (w *Pv77BarWorkflow) Execute(ctx workflow.Context) (*patchv1.Pv77BarOutput, error) {
	out, err := w.execute(ctx, &patchv1.Pv77Input{Next: w.Req.GetNext()})
	return &patchv1.Pv77BarOutput{TaskQueues: out.GetTaskQueues(), Defaults: out.GetDefaults()}, err
}

func (w *Pv77Workflows) Pv77Baz(ctx workflow.Context, input *patchv1.Pv77BazWorkflowInput) (patchv1.Pv77BazWorkflow, error) {
	return &Pv77BazWorkflow{w, input}, nil
}

func (w *Pv77BazWorkflow) Execute(ctx workflow.Context) (*patchv1.Pv77BazOutput, error) {
	out, err := w.execute(ctx, &patchv1.Pv77Input{Next: w.Req.GetNext()})
	return &patchv1.Pv77BazOutput{TaskQueues: out.GetTaskQueues(), Defaults: out.GetDefaults()}, err
}

func (w *Pv77Workflows) Pv77Qux(ctx workflow.Context, input *patchv1.Pv77QuxWorkflowInput) (patchv1.Pv77QuxWorkflow, error) {
	return &Pv77QuxWorkflow{w, input}, nil
}

func (w *Pv77QuxWorkflow) Execute(ctx workflow.Context) (*patchv1.Pv77QuxOutput, error) {
	out, err := w.execute(ctx, &patchv1.Pv77Input{Next: w.Req.GetNext()})
	return &patchv1.Pv77QuxOutput{TaskQueues: out.GetTaskQueues(), Defaults: out.GetDefaults()}, err
}

func (a *Pv77Activities) Pv77Foo(ctx context.Context, input *patchv1.Pv77FooInput) (*patchv1.Pv77FooOutput, error) {
	return &patchv1.Pv77FooOutput{TaskQueues: []string{activity.GetInfo(ctx).TaskQueue}}, nil
}

func (a *Pv77Activities) Pv77Bar(ctx context.Context, input *patchv1.Pv77BarInput) (*patchv1.Pv77BarOutput, error) {
	return &patchv1.Pv77BarOutput{TaskQueues: []string{activity.GetInfo(ctx).TaskQueue}}, nil
}

func (a *Pv77Activities) Pv77Baz(ctx context.Context, input *patchv1.Pv77BazInput) (*patchv1.Pv77BazOutput, error) {
	return &patchv1.Pv77BazOutput{TaskQueues: []string{activity.GetInfo(ctx).TaskQueue}}, nil
}

func (a *Pv77Activities) Pv77Qux(ctx context.Context, input *patchv1.Pv77QuxInput) (*patchv1.Pv77QuxOutput, error) {
	return &patchv1.Pv77QuxOutput{TaskQueues: []string{activity.GetInfo(ctx).TaskQueue}}, nil
}

type Pv77Output interface {
	GetDefaults() map[string]string
	GetTaskQueues() []string
}

func (w *Pv77Workflows) execute(ctx workflow.Context, input *patchv1.Pv77Input) (*patchv1.Pv77Output, error) {
	result := &patchv1.Pv77Output{
		TaskQueues: []string{workflow.GetInfo(ctx).TaskQueueName},
		Defaults:   patch.DefaultTaskQueues(ctx),
	}
	var out Pv77Output
	var err error
	var n string
	next := input.GetNext()
	for len(next) > 0 {
		n, next = next[0], next[1:]
		switch n {
		case patchv1.Pv77FooActivityName:
			out, err = patchv1.Pv77Foo(ctx, &patchv1.Pv77FooInput{Next: next})
		case patchv1.Pv77BarActivityName:
			out, err = patchv1.Pv77Bar(ctx, &patchv1.Pv77BarInput{Next: next})
		case patchv1.Pv77BazActivityName:
			out, err = patchv1.Pv77Baz(ctx, &patchv1.Pv77BazInput{Next: next})
		case patchv1.Pv77QuxActivityName:
			out, err = patchv1.Pv77Qux(ctx, &patchv1.Pv77QuxInput{Next: next})
		case patchv1.Pv77FooWorkflowName:
			out, err = patchv1.Pv77FooChild(ctx, &patchv1.Pv77FooInput{Next: next})
			next = nil
		case patchv1.Pv77BarWorkflowName:
			out, err = patchv1.Pv77BarChild(ctx, &patchv1.Pv77BarInput{Next: next})
			next = nil
		case patchv1.Pv77BazWorkflowName:
			out, err = patchv1.Pv77BazChild(ctx, &patchv1.Pv77BazInput{Next: next})
			next = nil
		case patchv1.Pv77QuxWorkflowName:
			out, err = patchv1.Pv77QuxChild(ctx, &patchv1.Pv77QuxInput{Next: next})
			next = nil
		default:
			return nil, fmt.Errorf("unknown workflow: %q", next[0])
		}
		if out != nil {
			result.TaskQueues = append(result.TaskQueues, out.GetTaskQueues()...)
			for k, v := range out.GetDefaults() {
				result.Defaults[k] = v
			}
		}
	}
	return result, err
}
