// Code generated by protoc-gen-go_temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-go_temporal 1.15.1-next (f1e76430351366c0f5ba139a759e99d0ffa098d7)
//	go go1.23.5
//	protoc (unknown)
//
// source: example/protocgengonexus/example.proto
package protocgengonexustemporal

import (
	"context"
	protocgengonexus "github.com/cludden/protoc-gen-go-temporal/gen/example/protocgengonexus"
	protocgengonexusnexus "github.com/cludden/protoc-gen-go-temporal/gen/example/protocgengonexus/protocgengonexusnexus"
	nexus "github.com/nexus-rpc/sdk-go/nexus"
	temporalnexus "go.temporal.io/sdk/temporalnexus"
	worker "go.temporal.io/sdk/worker"
)

// GreetingNexusHandler is an implementation of the protoc-gen-go-nexus handler
type GreetingNexusHandler struct {
	protocgengonexusnexus.UnimplementedGreetingNexusHandler
}

// RegisterGreetingNexusService initializes a new Greeting nexus service and registers it with the provided registry
func RegisterGreetingNexusService(r worker.NexusServiceRegistry) error {
	svc, err := protocgengonexusnexus.NewGreetingNexusService(&GreetingNexusHandler{})
	if err != nil {
		return err
	}
	r.RegisterNexusService(svc)
	return nil
}

// returns a nexus operation for executing a example.protocgengonexus.Greeting.Greet workflow
func (h *GreetingNexusHandler) Greet(name string) nexus.Operation[*protocgengonexus.GreetInput, *protocgengonexus.GreetOutput] {
	return temporalnexus.MustNewWorkflowRunOperationWithOptions(temporalnexus.WorkflowRunOperationOptions[*protocgengonexus.GreetInput, *protocgengonexus.GreetOutput]{
		Handler: func(ctx context.Context, input *protocgengonexus.GreetInput, opts nexus.StartOperationOptions) (temporalnexus.WorkflowHandle[*protocgengonexus.GreetOutput], error) {
			o, err := protocgengonexus.NewGreetOptions().Build(input.ProtoReflect())
			if err != nil {
				return nil, err
			}
			return temporalnexus.ExecuteUntypedWorkflow[*protocgengonexus.GreetOutput](ctx, opts, o, protocgengonexus.GreetWorkflowName, input)
		},
		Name: name,
	})
}

// CallerNexusHandler is an implementation of the protoc-gen-go-nexus handler
type CallerNexusHandler struct {
	protocgengonexusnexus.UnimplementedCallerNexusHandler
}

// RegisterCallerNexusService initializes a new Caller nexus service and registers it with the provided registry
func RegisterCallerNexusService(r worker.NexusServiceRegistry) error {
	svc, err := protocgengonexusnexus.NewCallerNexusService(&CallerNexusHandler{})
	if err != nil {
		return err
	}
	r.RegisterNexusService(svc)
	return nil
}

// returns a nexus operation for executing a example.protocgengonexus.Caller.CallGreet workflow
func (h *CallerNexusHandler) CallGreet(name string) nexus.Operation[*protocgengonexus.CallGreetInput, *protocgengonexus.CallGreetOutput] {
	return temporalnexus.MustNewWorkflowRunOperationWithOptions(temporalnexus.WorkflowRunOperationOptions[*protocgengonexus.CallGreetInput, *protocgengonexus.CallGreetOutput]{
		Handler: func(ctx context.Context, input *protocgengonexus.CallGreetInput, opts nexus.StartOperationOptions) (temporalnexus.WorkflowHandle[*protocgengonexus.CallGreetOutput], error) {
			o, err := protocgengonexus.NewCallGreetOptions().Build(input.ProtoReflect())
			if err != nil {
				return nil, err
			}
			return temporalnexus.ExecuteUntypedWorkflow[*protocgengonexus.CallGreetOutput](ctx, opts, o, protocgengonexus.CallGreetWorkflowName, input)
		},
		Name: name,
	})
}
