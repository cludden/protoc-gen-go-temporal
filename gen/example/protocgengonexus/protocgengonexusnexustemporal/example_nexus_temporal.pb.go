// Code generated by protoc-gen-nexus-temporal. DO NOT EDIT.
// versions:
//
//	protoc-gen-nexus-temporal dev (latest)
//	go go1.23.5
//	protoc (unknown)
//
// source: example/protocgengonexus/example.proto
package protocgengonexusnexustemporal

import (
	protocgengonexus "github.com/cludden/protoc-gen-go-temporal/gen/example/protocgengonexus"
	protocgengonexusnexus "github.com/cludden/protoc-gen-go-temporal/gen/example/protocgengonexus/protocgengonexusnexus"
	workflow "go.temporal.io/sdk/workflow"
)

type GreetingNexusClient struct {
	client workflow.NexusClient
}

// NewGreetingNexusClient initializes a new GreetingNexusClient.
func NewGreetingNexusClient(endpoint string) *GreetingNexusClient {
	return &GreetingNexusClient{
		client: workflow.NewNexusClient(endpoint, protocgengonexusnexus.GreetingServiceName),
	}
}

type GenerateGreetingFuture struct {
	workflow.NexusOperationFuture
}

func (f *GenerateGreetingFuture) GetTyped(ctx workflow.Context) (*protocgengonexus.GenerateGreetingOutput, error) {
	var output protocgengonexus.GenerateGreetingOutput
	err := f.Get(ctx, &output)
	return &output, err
}
func (c *GreetingNexusClient) GenerateGreetingAsync(ctx workflow.Context, input *protocgengonexus.GenerateGreetingInput, options workflow.NexusOperationOptions) GenerateGreetingFuture {
	fut := c.client.ExecuteOperation(ctx, protocgengonexusnexus.GreetingGenerateGreetingOperationName, input, options)
	return GenerateGreetingFuture{
		fut,
	}
}
func (c *GreetingNexusClient) GenerateGreeting(ctx workflow.Context, input *protocgengonexus.GenerateGreetingInput, options workflow.NexusOperationOptions) (*protocgengonexus.GenerateGreetingOutput, error) {
	var output protocgengonexus.GenerateGreetingOutput
	fut := c.client.ExecuteOperation(ctx, protocgengonexusnexus.GreetingGenerateGreetingOperationName, input, options)
	err := fut.Get(ctx, &output)
	return &output, err
}

type GreetFuture struct {
	workflow.NexusOperationFuture
}

func (f *GreetFuture) GetTyped(ctx workflow.Context) (*protocgengonexus.GreetOutput, error) {
	var output protocgengonexus.GreetOutput
	err := f.Get(ctx, &output)
	return &output, err
}
func (c *GreetingNexusClient) GreetAsync(ctx workflow.Context, input *protocgengonexus.GreetInput, options workflow.NexusOperationOptions) GreetFuture {
	fut := c.client.ExecuteOperation(ctx, protocgengonexusnexus.GreetingGreetOperationName, input, options)
	return GreetFuture{
		fut,
	}
}
func (c *GreetingNexusClient) Greet(ctx workflow.Context, input *protocgengonexus.GreetInput, options workflow.NexusOperationOptions) (*protocgengonexus.GreetOutput, error) {
	var output protocgengonexus.GreetOutput
	fut := c.client.ExecuteOperation(ctx, protocgengonexusnexus.GreetingGreetOperationName, input, options)
	err := fut.Get(ctx, &output)
	return &output, err
}

type CallerNexusClient struct {
	client workflow.NexusClient
}

// NewCallerNexusClient initializes a new CallerNexusClient.
func NewCallerNexusClient(endpoint string) *CallerNexusClient {
	return &CallerNexusClient{
		client: workflow.NewNexusClient(endpoint, protocgengonexusnexus.CallerServiceName),
	}
}

type CallGreetFuture struct {
	workflow.NexusOperationFuture
}

func (f *CallGreetFuture) GetTyped(ctx workflow.Context) (*protocgengonexus.CallGreetOutput, error) {
	var output protocgengonexus.CallGreetOutput
	err := f.Get(ctx, &output)
	return &output, err
}
func (c *CallerNexusClient) CallGreetAsync(ctx workflow.Context, input *protocgengonexus.CallGreetInput, options workflow.NexusOperationOptions) CallGreetFuture {
	fut := c.client.ExecuteOperation(ctx, protocgengonexusnexus.CallerCallGreetOperationName, input, options)
	return CallGreetFuture{
		fut,
	}
}
func (c *CallerNexusClient) CallGreet(ctx workflow.Context, input *protocgengonexus.CallGreetInput, options workflow.NexusOperationOptions) (*protocgengonexus.CallGreetOutput, error) {
	var output protocgengonexus.CallGreetOutput
	fut := c.client.ExecuteOperation(ctx, protocgengonexusnexus.CallerCallGreetOperationName, input, options)
	err := fut.Get(ctx, &output)
	return &output, err
}
