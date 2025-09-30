package main

import (
	"fmt"

	issue_125v1 "github.com/cludden/protoc-gen-go-temporal/gen/test/issue-125/v1"
	"go.temporal.io/sdk/workflow"
)

type (
	Workflows struct{}

	FooWorkflow struct {
		*issue_125v1.FooWorkflowInput
		bar bool
		baz bool
	}
)

func (w *Workflows) Foo(
	ctx workflow.Context,
	input *issue_125v1.FooWorkflowInput,
) (issue_125v1.FooWorkflow, error) {
	return &FooWorkflow{input, false, false}, nil
}

func (w *FooWorkflow) Execute(ctx workflow.Context) (*issue_125v1.FooOutput, error) {
	if err := workflow.Await(ctx, func() bool {
		return w.bar && w.baz
	}); err != nil {
		return nil, err
	}
	return &issue_125v1.FooOutput{
		Result: fmt.Sprintf("foo:%s", w.Req.GetId()),
	}, nil
}

func (w *FooWorkflow) Bar(
	ctx workflow.Context,
	input *issue_125v1.BarInput,
) (*issue_125v1.BarOutput, error) {
	w.bar = true
	return &issue_125v1.BarOutput{
		Result: fmt.Sprintf("bar:%s", input.GetId()),
	}, nil
}

func (w *FooWorkflow) Baz(
	ctx workflow.Context,
	input *issue_125v1.BazInput,
) (*issue_125v1.BazOutput, error) {
	w.baz = true
	return &issue_125v1.BazOutput{
		Result: fmt.Sprintf("baz:%s", input.GetId()),
	}, nil
}
