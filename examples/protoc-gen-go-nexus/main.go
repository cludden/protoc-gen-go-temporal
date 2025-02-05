package main

import (
	"context"
	"fmt"
	"log"

	examplepb "github.com/cludden/protoc-gen-go-temporal/gen/example/protocgengonexus"
	examplenexustemporalpb "github.com/cludden/protoc-gen-go-temporal/gen/example/protocgengonexus/protocgengonexusnexustemporal"
	exampletemporalpb "github.com/cludden/protoc-gen-go-temporal/gen/example/protocgengonexus/protocgengonexustemporal"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	w := worker.New(c, examplepb.GreetingTaskQueue, worker.Options{})
	if err := RegisterGreeting(w); err != nil {
		log.Fatal(err)
	}

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatal(err)
	}
}

// =============================================================================

type (
	GreetingWorkflows struct{}

	GreetingActivities struct{}

	GreetWorkflow struct {
		*examplepb.GreetWorkflowInput
	}
)

func RegisterGreeting(r worker.Registry) error {
	examplepb.RegisterGreetingWorkflows(r, &GreetingWorkflows{})
	examplepb.RegisterGreetingActivities(r, &GreetingActivities{})
	return exampletemporalpb.RegisterGreetingNexusService(r)
}

func (w *GreetingWorkflows) Greet(ctx workflow.Context, input *examplepb.GreetWorkflowInput) (examplepb.GreetWorkflow, error) {
	return &GreetWorkflow{input}, nil
}

func (w *GreetWorkflow) Execute(ctx workflow.Context) (*examplepb.GreetOutput, error) {
	greeting, err := examplepb.GenerateGreeting(ctx, &examplepb.GenerateGreetingInput{Name: w.Req.GetName()})
	if err != nil {
		return nil, err
	}
	return &examplepb.GreetOutput{Greeting: greeting.GetGreeting()}, nil
}

func (a *GreetingActivities) GenerateGreeting(ctx context.Context, input *examplepb.GenerateGreetingInput) (*examplepb.GenerateGreetingOutput, error) {
	return &examplepb.GenerateGreetingOutput{Greeting: fmt.Sprintf("Hello, %s!", input.GetName())}, nil
}

// =============================================================================

type (
	CallerWorkflows struct {
		greeting *examplenexustemporalpb.GreetingNexusClient
	}

	CallerActivities struct{}

	CallGreetWorkflow struct {
		*CallerWorkflows
		*examplepb.CallGreetWorkflowInput
	}
)

func RegisterCaller(r worker.Registry, greetingEndpoint string) error {
	examplepb.RegisterCallerWorkflows(r, &CallerWorkflows{
		greeting: examplenexustemporalpb.NewGreetingNexusClient(greetingEndpoint),
	})
	examplepb.RegisterCallerActivities(r, &CallerActivities{})
	return nil
}

func (w *CallerWorkflows) CallGreet(ctx workflow.Context, input *examplepb.CallGreetWorkflowInput) (examplepb.CallGreetWorkflow, error) {
	return &CallGreetWorkflow{w, input}, nil
}

func (w *CallGreetWorkflow) Execute(ctx workflow.Context) (*examplepb.CallGreetOutput, error) {
	resp, err := w.greeting.Greet(ctx, &examplepb.GreetInput{Name: w.Req.GetName()}, workflow.NexusOperationOptions{})
	if err != nil {
		return nil, err
	}
	return &examplepb.CallGreetOutput{Greeting: resp.GetGreeting()}, nil
}
