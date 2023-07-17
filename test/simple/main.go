package main

import (
	"context"
	"log"
	"os"
	"strings"

	simplepb "github.com/cludden/protoc-gen-go-temporal/gen/simple"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

type Workflows struct {
	simplepb.SimpleWorkflows
}

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

	// Handle input
	s.SomeSignal1.Select(s.sel, func() {
		s.events = append(s.events, "some signal 1")
	})
	s.SomeSignal2.Select(s.sel, func(req *simplepb.SomeSignal2Request) {
		s.events = append(s.events, "some signal 2 with param "+req.RequestVal)
	})

	// Run until done
	for s.sel.HasPending() {
		s.sel.Select(ctx)
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
