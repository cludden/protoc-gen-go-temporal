// Code generated by mockery. DO NOT EDIT.

package xnserrv1mocks

import (
	context "context"

	xnserrv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/xnserr/v1"
	mock "github.com/stretchr/testify/mock"
)

// MockServerClient is an autogenerated mock type for the ServerClient type
type MockServerClient struct {
	mock.Mock
}

type MockServerClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockServerClient) EXPECT() *MockServerClient_Expecter {
	return &MockServerClient_Expecter{mock: &_m.Mock}
}

// CancelWorkflow provides a mock function with given fields: ctx, workflowID, runID
func (_m *MockServerClient) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	ret := _m.Called(ctx, workflowID, runID)

	if len(ret) == 0 {
		panic("no return value specified for CancelWorkflow")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, workflowID, runID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServerClient_CancelWorkflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CancelWorkflow'
type MockServerClient_CancelWorkflow_Call struct {
	*mock.Call
}

// CancelWorkflow is a helper method to define mock.On call
//   - ctx context.Context
//   - workflowID string
//   - runID string
func (_e *MockServerClient_Expecter) CancelWorkflow(ctx interface{}, workflowID interface{}, runID interface{}) *MockServerClient_CancelWorkflow_Call {
	return &MockServerClient_CancelWorkflow_Call{Call: _e.mock.On("CancelWorkflow", ctx, workflowID, runID)}
}

func (_c *MockServerClient_CancelWorkflow_Call) Run(run func(ctx context.Context, workflowID string, runID string)) *MockServerClient_CancelWorkflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockServerClient_CancelWorkflow_Call) Return(_a0 error) *MockServerClient_CancelWorkflow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServerClient_CancelWorkflow_Call) RunAndReturn(run func(context.Context, string, string) error) *MockServerClient_CancelWorkflow_Call {
	_c.Call.Return(run)
	return _c
}

// GetSleep provides a mock function with given fields: ctx, workflowID, runID
func (_m *MockServerClient) GetSleep(ctx context.Context, workflowID string, runID string) xnserrv1.SleepRun {
	ret := _m.Called(ctx, workflowID, runID)

	if len(ret) == 0 {
		panic("no return value specified for GetSleep")
	}

	var r0 xnserrv1.SleepRun
	if rf, ok := ret.Get(0).(func(context.Context, string, string) xnserrv1.SleepRun); ok {
		r0 = rf(ctx, workflowID, runID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(xnserrv1.SleepRun)
		}
	}

	return r0
}

// MockServerClient_GetSleep_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSleep'
type MockServerClient_GetSleep_Call struct {
	*mock.Call
}

// GetSleep is a helper method to define mock.On call
//   - ctx context.Context
//   - workflowID string
//   - runID string
func (_e *MockServerClient_Expecter) GetSleep(ctx interface{}, workflowID interface{}, runID interface{}) *MockServerClient_GetSleep_Call {
	return &MockServerClient_GetSleep_Call{Call: _e.mock.On("GetSleep", ctx, workflowID, runID)}
}

func (_c *MockServerClient_GetSleep_Call) Run(run func(ctx context.Context, workflowID string, runID string)) *MockServerClient_GetSleep_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockServerClient_GetSleep_Call) Return(_a0 xnserrv1.SleepRun) *MockServerClient_GetSleep_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServerClient_GetSleep_Call) RunAndReturn(run func(context.Context, string, string) xnserrv1.SleepRun) *MockServerClient_GetSleep_Call {
	_c.Call.Return(run)
	return _c
}

// Sleep provides a mock function with given fields: ctx, req, opts
func (_m *MockServerClient) Sleep(ctx context.Context, req *xnserrv1.SleepRequest, opts ...*xnserrv1.SleepOptions) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Sleep")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *xnserrv1.SleepRequest, ...*xnserrv1.SleepOptions) error); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServerClient_Sleep_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sleep'
type MockServerClient_Sleep_Call struct {
	*mock.Call
}

// Sleep is a helper method to define mock.On call
//   - ctx context.Context
//   - req *xnserrv1.SleepRequest
//   - opts ...*xnserrv1.SleepOptions
func (_e *MockServerClient_Expecter) Sleep(ctx interface{}, req interface{}, opts ...interface{}) *MockServerClient_Sleep_Call {
	return &MockServerClient_Sleep_Call{Call: _e.mock.On("Sleep",
		append([]interface{}{ctx, req}, opts...)...)}
}

func (_c *MockServerClient_Sleep_Call) Run(run func(ctx context.Context, req *xnserrv1.SleepRequest, opts ...*xnserrv1.SleepOptions)) *MockServerClient_Sleep_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*xnserrv1.SleepOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(*xnserrv1.SleepOptions)
			}
		}
		run(args[0].(context.Context), args[1].(*xnserrv1.SleepRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockServerClient_Sleep_Call) Return(_a0 error) *MockServerClient_Sleep_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServerClient_Sleep_Call) RunAndReturn(run func(context.Context, *xnserrv1.SleepRequest, ...*xnserrv1.SleepOptions) error) *MockServerClient_Sleep_Call {
	_c.Call.Return(run)
	return _c
}

// SleepAsync provides a mock function with given fields: ctx, req, opts
func (_m *MockServerClient) SleepAsync(ctx context.Context, req *xnserrv1.SleepRequest, opts ...*xnserrv1.SleepOptions) (xnserrv1.SleepRun, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SleepAsync")
	}

	var r0 xnserrv1.SleepRun
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *xnserrv1.SleepRequest, ...*xnserrv1.SleepOptions) (xnserrv1.SleepRun, error)); ok {
		return rf(ctx, req, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *xnserrv1.SleepRequest, ...*xnserrv1.SleepOptions) xnserrv1.SleepRun); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(xnserrv1.SleepRun)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *xnserrv1.SleepRequest, ...*xnserrv1.SleepOptions) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServerClient_SleepAsync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SleepAsync'
type MockServerClient_SleepAsync_Call struct {
	*mock.Call
}

// SleepAsync is a helper method to define mock.On call
//   - ctx context.Context
//   - req *xnserrv1.SleepRequest
//   - opts ...*xnserrv1.SleepOptions
func (_e *MockServerClient_Expecter) SleepAsync(ctx interface{}, req interface{}, opts ...interface{}) *MockServerClient_SleepAsync_Call {
	return &MockServerClient_SleepAsync_Call{Call: _e.mock.On("SleepAsync",
		append([]interface{}{ctx, req}, opts...)...)}
}

func (_c *MockServerClient_SleepAsync_Call) Run(run func(ctx context.Context, req *xnserrv1.SleepRequest, opts ...*xnserrv1.SleepOptions)) *MockServerClient_SleepAsync_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*xnserrv1.SleepOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(*xnserrv1.SleepOptions)
			}
		}
		run(args[0].(context.Context), args[1].(*xnserrv1.SleepRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockServerClient_SleepAsync_Call) Return(_a0 xnserrv1.SleepRun, _a1 error) *MockServerClient_SleepAsync_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServerClient_SleepAsync_Call) RunAndReturn(run func(context.Context, *xnserrv1.SleepRequest, ...*xnserrv1.SleepOptions) (xnserrv1.SleepRun, error)) *MockServerClient_SleepAsync_Call {
	_c.Call.Return(run)
	return _c
}

// TerminateWorkflow provides a mock function with given fields: ctx, workflowID, runID, reason, details
func (_m *MockServerClient) TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, workflowID, runID, reason)
	_ca = append(_ca, details...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TerminateWorkflow")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, ...interface{}) error); ok {
		r0 = rf(ctx, workflowID, runID, reason, details...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServerClient_TerminateWorkflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TerminateWorkflow'
type MockServerClient_TerminateWorkflow_Call struct {
	*mock.Call
}

// TerminateWorkflow is a helper method to define mock.On call
//   - ctx context.Context
//   - workflowID string
//   - runID string
//   - reason string
//   - details ...interface{}
func (_e *MockServerClient_Expecter) TerminateWorkflow(ctx interface{}, workflowID interface{}, runID interface{}, reason interface{}, details ...interface{}) *MockServerClient_TerminateWorkflow_Call {
	return &MockServerClient_TerminateWorkflow_Call{Call: _e.mock.On("TerminateWorkflow",
		append([]interface{}{ctx, workflowID, runID, reason}, details...)...)}
}

func (_c *MockServerClient_TerminateWorkflow_Call) Run(run func(ctx context.Context, workflowID string, runID string, reason string, details ...interface{})) *MockServerClient_TerminateWorkflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-4)
		for i, a := range args[4:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockServerClient_TerminateWorkflow_Call) Return(_a0 error) *MockServerClient_TerminateWorkflow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServerClient_TerminateWorkflow_Call) RunAndReturn(run func(context.Context, string, string, string, ...interface{}) error) *MockServerClient_TerminateWorkflow_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockServerClient creates a new instance of MockServerClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockServerClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockServerClient {
	mock := &MockServerClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
