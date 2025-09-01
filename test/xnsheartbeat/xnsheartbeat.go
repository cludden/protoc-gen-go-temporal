package xnsheartbeat

import (
	xnsheartbeatv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/xnsheartbeat/v1"
	"github.com/cludden/protoc-gen-go-temporal/gen/test/xnsheartbeat/v1/xnsheartbeatv1xns"
	"go.temporal.io/sdk/workflow"
)

type (
	Workflows struct{}

	TestWorkflow struct {
		*Workflows
		*xnsheartbeatv1.TestWorkflowWorkflowInput
		signalReceived bool
		updateReceived bool
	}
)

func (w *Workflows) TestWorkflow(
	ctx workflow.Context,
	input *xnsheartbeatv1.TestWorkflowWorkflowInput,
) (xnsheartbeatv1.TestWorkflowWorkflow, error) {
	return &TestWorkflow{w, input, false, false}, nil
}

func (w *TestWorkflow) Execute(ctx workflow.Context) (*xnsheartbeatv1.TestWorkflowOutput, error) {
	_, _ = w.TestSignal.Receive(ctx)
	w.signalReceived = true
	if err := workflow.Await(ctx, func() bool {
		return w.updateReceived && workflow.AllHandlersFinished(ctx)
	}); err != nil {
		return nil, err
	}
	return &xnsheartbeatv1.TestWorkflowOutput{}, nil
}

func (w *TestWorkflow) TestUpdate(
	ctx workflow.Context,
	input *xnsheartbeatv1.TestUpdateInput,
) (*xnsheartbeatv1.TestUpdateOutput, error) {
	if err := workflow.Await(ctx, func() bool {
		return w.signalReceived
	}); err != nil {
		return nil, err
	}
	w.updateReceived = true
	return &xnsheartbeatv1.TestUpdateOutput{}, nil
}

type CallerWorkflows struct{}

type CallTestWorkflow struct {
	*CallerWorkflows
	*xnsheartbeatv1.CallTestWorkflowWorkflowInput
}

func (w *CallerWorkflows) CallTestWorkflow(
	ctx workflow.Context,
	input *xnsheartbeatv1.CallTestWorkflowWorkflowInput,
) (xnsheartbeatv1.CallTestWorkflowWorkflow, error) {
	return &CallTestWorkflow{w, input}, nil
}

func (w *CallTestWorkflow) Execute(ctx workflow.Context) error {
	_, run, err := xnsheartbeatv1xns.TestWorkflowWithTestUpdate(
		ctx,
		&xnsheartbeatv1.TestWorkflowInput{},
		&xnsheartbeatv1.TestUpdateInput{},
	)
	if err != nil {
		return err
	}
	_, err = run.Get(ctx)
	if err != nil {
		return err
	}
	return nil
}
