package testutil

import (
	"context"

	"go.temporal.io/sdk/temporal"
)

const (
	ErrCodeUpdateInvalidResponseType = "UpdateInvalidResponseType"
	ErrCodeUpdateValidationFailed    = "UpdateValidationFailed"
)

type UpdateCallbacks struct {
	accepted bool
	result   chan any
	err      chan error
}

func NewUpdateCallbacks() *UpdateCallbacks {
	return &UpdateCallbacks{
		result: make(chan any, 1),
		err:    make(chan error, 1),
	}
}

func (uc *UpdateCallbacks) Accept() {
	uc.accepted = true
}

func (uc *UpdateCallbacks) Complete(result any, err error) {
	if err != nil {
		uc.err <- err
	} else {
		uc.result <- result
	}
}

func (uc *UpdateCallbacks) Get(ctx context.Context) (any, error) {
	defer func() {
		close(uc.err)
		close(uc.result)
	}()

	select {
	case result := <-uc.result:
		return result, nil
	case err := <-uc.err:
		return nil, err
	case <-ctx.Done():
		return nil, context.Canceled
	}
}

func (uc *UpdateCallbacks) Reject(err error) {
	uc.err <- temporal.NewNonRetryableApplicationError("update validation failed", ErrCodeUpdateValidationFailed, err)
}
