package main

import (
	workerversioningv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/workerversioning/v1"
	"go.temporal.io/sdk/worker"
)

type Workflows struct {
	workerversioningv1.ExampleWorkflows
}

func Register(r worker.Registry) {
	workerversioningv1.RegisterExampleWorkflows(r, &Workflows{})
}
