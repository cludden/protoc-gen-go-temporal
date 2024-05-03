package xns

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/temporal"
)

func TestErrorToApplicationError(t *testing.T) {
	require := require.New(t)

	err := ErrorToApplicationError(nil)
	require.Nil(err)

	err = errors.New("uh oh")
	require.ErrorIs(ErrorToApplicationError(err), err)

	err = temporal.NewNonRetryableApplicationError("uh oh", "Foo", nil)
	require.ErrorIs(ErrorToApplicationError(err), err)

	err = &temporal.WorkflowExecutionError{}
	require.NotNil(Unwrap(ErrorToApplicationError(err)))
	err = ErrorToApplicationError(err)
	require.Equal("WorkflowExecutionError", Code(err))
	require.True(IsNonRetryable(err))

	err = &temporal.CanceledError{}
	require.Equal(err, ErrorToApplicationError(err))
	require.True(temporal.IsCanceledError(ErrorToApplicationError(err)))

	err = &temporal.TerminatedError{}
	require.Equal(err, ErrorToApplicationError(err))
	require.True(temporal.IsTerminatedError(ErrorToApplicationError(err)))

	err = &temporal.ChildWorkflowExecutionError{}
	require.NotNil(Unwrap(ErrorToApplicationError(err)))
	err = ErrorToApplicationError(err)
	require.Equal("ChildWorkflowExecutionError", Code(err))
	require.True(IsNonRetryable(err))
}
