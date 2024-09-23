package errs

import "go.temporal.io/sdk/workflow"

type futureError struct {
	err error
}

func NewFutureError(err error) workflow.Future {
	return &futureError{err}
}

func (e *futureError) Get(workflow.Context, any) error {
	return e.err
}

func (e *futureError) IsReady() bool {
	return true
}

type nexusOperationFutureError struct {
	futureError
}

func NewNexusOperationFutureError(err error) workflow.NexusOperationFuture {
	return &nexusOperationFutureError{futureError: futureError{err}}
}

func (e *nexusOperationFutureError) GetNexusOperationExecution() workflow.Future {
	return &e.futureError
}
