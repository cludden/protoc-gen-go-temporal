package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	simplepb "github.com/cludden/protoc-gen-go-temporal/gen/simple"
	"github.com/urfave/cli/v2"
	logger "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

type Workflows struct{}

// ============================================================================

type someWorkflow1 struct {
	*simplepb.SomeWorkflow1Input
	sel    workflow.Selector
	events []string
}

func Register(r worker.Registry) {
	simplepb.RegisterSimpleWorkflows(r, &Workflows{})
	simplepb.RegisterSimpleActivities(r, &Activities{})
}

func (w *Workflows) SomeWorkflow1(ctx workflow.Context, in *simplepb.SomeWorkflow1Input) (simplepb.SomeWorkflow1Workflow, error) {
	return &someWorkflow1{SomeWorkflow1Input: in, sel: workflow.NewSelector(ctx)}, nil
}

func (s *someWorkflow1) Execute(ctx workflow.Context) (*simplepb.SomeWorkflow1Response, error) {
	s.events = append(s.events, "started with param "+s.Req.RequestVal)

	// Call regular activity
	resp, err := simplepb.SomeActivity3(ctx, &simplepb.SomeActivity3Request{
		RequestVal: "some activity param",
	})
	if err != nil {
		return nil, err
	}
	s.events = append(s.events, "some activity 3 with response "+resp.ResponseVal)

	// Call local activity
	resp, err = simplepb.SomeActivity3Local(ctx, &simplepb.SomeActivity3Request{
		RequestVal: "some local activity param",
	})
	if err != nil {
		return nil, err
	}
	s.events = append(s.events, "some local activity 3 with response "+resp.ResponseVal)

	var signal1, signal2 int
	for {
		workflow.NewSelector(ctx).
			AddReceive(s.SomeSignal1.Channel, func(workflow.ReceiveChannel, bool) {
				s.SomeSignal1.ReceiveAsync()
				s.events = append(s.events, "some signal 1")
				signal1++
			}).
			AddReceive(s.SomeSignal2.Channel, func(workflow.ReceiveChannel, bool) {
				req := s.SomeSignal2.ReceiveAsync()
				s.events = append(s.events, "some signal 2 with param "+req.RequestVal)
				signal2++
			}).
			Select(ctx)

		if signal1 > 0 && signal2 > 1 {
			break
		}
	}

	return &simplepb.SomeWorkflow1Response{ResponseVal: strings.Join(s.events, "\n")}, nil
}

func (s *someWorkflow1) SomeQuery1() (*simplepb.SomeQuery1Response, error) {
	return &simplepb.SomeQuery1Response{
		ResponseVal: strings.Join(s.events, "\n") + "\nsome query 1",
	}, nil
}

func (s *someWorkflow1) SomeQuery2(req *simplepb.SomeQuery2Request) (*simplepb.SomeQuery2Response, error) {
	return &simplepb.SomeQuery2Response{
		ResponseVal: strings.Join(s.events, "\n") + "\nsome query 2 with param " + req.RequestVal,
	}, nil
}

// ============================================================================

type someWorkflow2 struct {
	*simplepb.SomeWorkflow2Input
	log     logger.Logger
	updates int
}

func (w *Workflows) SomeWorkflow2(ctx workflow.Context, input *simplepb.SomeWorkflow2Input) (simplepb.SomeWorkflow2Workflow, error) {
	wf := &someWorkflow2{SomeWorkflow2Input: input, log: workflow.GetLogger(ctx)}
	return wf, nil
}

func (wf *someWorkflow2) Execute(ctx workflow.Context) error {
	return workflow.Await(ctx, func() bool {
		fmt.Printf("updates: %d\n", wf.updates)
		return wf.updates > 0
	})
}

func (wf *someWorkflow2) SomeUpdate1(ctx workflow.Context, req *simplepb.SomeUpdate1Request) (*simplepb.SomeUpdate1Response, error) {
	wf.log.Info("SomeUpdate1", "req", req.String())
	wf.updates++
	return &simplepb.SomeUpdate1Response{ResponseVal: strings.ToUpper(req.GetRequestVal())}, nil
}

func (wf *someWorkflow2) ValidateSomeUpdate1(ctx workflow.Context, req *simplepb.SomeUpdate1Request) error {
	if l := len(req.GetRequestVal()); l < 3 || l > 10 {
		return fmt.Errorf("request val length must be between 3 and 10")
	}
	return nil
}

// ============================================================================

type someWorkflow3 struct {
	*simplepb.SomeWorkflow3Input
	log     logger.Logger
	signals int
	updates int
}

func (w *Workflows) SomeWorkflow3(ctx workflow.Context, input *simplepb.SomeWorkflow3Input) (simplepb.SomeWorkflow3Workflow, error) {
	return &someWorkflow3{input, workflow.GetLogger(ctx), 0, 0}, nil
}

func (wf *someWorkflow3) Execute(ctx workflow.Context) error {
	wf.SomeSignal2.Receive(ctx)
	return nil
}

// ============================================================================

type Activities struct{}

var ActivityEvents []string

func (Activities) SomeActivity1(context.Context) error {
	ActivityEvents = append(ActivityEvents, "some activity 1")
	return nil
}

func (Activities) SomeActivity2(ctx context.Context, req *simplepb.SomeActivity2Request) error {
	ActivityEvents = append(ActivityEvents, "some activity 2 with param "+req.RequestVal)
	return nil
}

func (Activities) SomeActivity3(ctx context.Context, req *simplepb.SomeActivity3Request) (*simplepb.SomeActivity3Response, error) {
	ActivityEvents = append(ActivityEvents, "some activity 3 with param "+req.RequestVal)
	return &simplepb.SomeActivity3Response{ResponseVal: "some response"}, nil
}

func (Activities) SomeSignal1(ctx context.Context) error {
	return nil
}

func (Activities) SomeSignal2(ctx context.Context, req *simplepb.SomeSignal2Request) error {
	return nil
}

func (Activities) SomeSignal3(ctx context.Context, req *simplepb.SomeSignal3Request) (*simplepb.SomeSignal3Response, error) {
	return &simplepb.SomeSignal3Response{ResponseVal: req.GetRequestVal()}, nil
}

func (Activities) SomeUpdate1(ctx context.Context, req *simplepb.SomeUpdate1Request) (*simplepb.SomeUpdate1Response, error) {
	return &simplepb.SomeUpdate1Response{ResponseVal: req.GetRequestVal()}, nil
}

func newCli() (*cli.App, error) {
	simpleCmd, err := simplepb.NewSimpleCliCommand()
	if err != nil {
		return nil, err
	}
	otherCmd, err := simplepb.NewOtherCliCommand()
	if err != nil {
		return nil, err
	}
	return &cli.App{
		Name:     "test",
		Commands: []*cli.Command{simpleCmd, otherCmd},
	}, nil
}

func main() {
	app, err := newCli()
	if err != nil {
		log.Fatal(err)
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
