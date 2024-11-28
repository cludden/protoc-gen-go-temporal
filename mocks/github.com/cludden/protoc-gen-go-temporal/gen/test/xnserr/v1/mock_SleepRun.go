// Code generated by mockery v2.51.0. DO NOT EDIT.

package xnserrv1mocks

import (
	context "context"

	client "go.temporal.io/sdk/client"

	mock "github.com/stretchr/testify/mock"
)

// MockSleepRun is an autogenerated mock type for the SleepRun type
type MockSleepRun struct {
	mock.Mock
}

type MockSleepRun_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSleepRun) EXPECT() *MockSleepRun_Expecter {
	return &MockSleepRun_Expecter{mock: &_m.Mock}
}

// Cancel provides a mock function with given fields: ctx
func (_m *MockSleepRun) Cancel(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Cancel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSleepRun_Cancel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Cancel'
type MockSleepRun_Cancel_Call struct {
	*mock.Call
}

// Cancel is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSleepRun_Expecter) Cancel(ctx interface{}) *MockSleepRun_Cancel_Call {
	return &MockSleepRun_Cancel_Call{Call: _e.mock.On("Cancel", ctx)}
}

func (_c *MockSleepRun_Cancel_Call) Run(run func(ctx context.Context)) *MockSleepRun_Cancel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSleepRun_Cancel_Call) Return(_a0 error) *MockSleepRun_Cancel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSleepRun_Cancel_Call) RunAndReturn(run func(context.Context) error) *MockSleepRun_Cancel_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx
func (_m *MockSleepRun) Get(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSleepRun_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockSleepRun_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSleepRun_Expecter) Get(ctx interface{}) *MockSleepRun_Get_Call {
	return &MockSleepRun_Get_Call{Call: _e.mock.On("Get", ctx)}
}

func (_c *MockSleepRun_Get_Call) Run(run func(ctx context.Context)) *MockSleepRun_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSleepRun_Get_Call) Return(_a0 error) *MockSleepRun_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSleepRun_Get_Call) RunAndReturn(run func(context.Context) error) *MockSleepRun_Get_Call {
	_c.Call.Return(run)
	return _c
}

// ID provides a mock function with no fields
func (_m *MockSleepRun) ID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockSleepRun_ID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ID'
type MockSleepRun_ID_Call struct {
	*mock.Call
}

// ID is a helper method to define mock.On call
func (_e *MockSleepRun_Expecter) ID() *MockSleepRun_ID_Call {
	return &MockSleepRun_ID_Call{Call: _e.mock.On("ID")}
}

func (_c *MockSleepRun_ID_Call) Run(run func()) *MockSleepRun_ID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSleepRun_ID_Call) Return(_a0 string) *MockSleepRun_ID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSleepRun_ID_Call) RunAndReturn(run func() string) *MockSleepRun_ID_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with no fields
func (_m *MockSleepRun) Run() client.WorkflowRun {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 client.WorkflowRun
	if rf, ok := ret.Get(0).(func() client.WorkflowRun); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.WorkflowRun)
		}
	}

	return r0
}

// MockSleepRun_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockSleepRun_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
func (_e *MockSleepRun_Expecter) Run() *MockSleepRun_Run_Call {
	return &MockSleepRun_Run_Call{Call: _e.mock.On("Run")}
}

func (_c *MockSleepRun_Run_Call) Run(run func()) *MockSleepRun_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSleepRun_Run_Call) Return(_a0 client.WorkflowRun) *MockSleepRun_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSleepRun_Run_Call) RunAndReturn(run func() client.WorkflowRun) *MockSleepRun_Run_Call {
	_c.Call.Return(run)
	return _c
}

// RunID provides a mock function with no fields
func (_m *MockSleepRun) RunID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RunID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockSleepRun_RunID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunID'
type MockSleepRun_RunID_Call struct {
	*mock.Call
}

// RunID is a helper method to define mock.On call
func (_e *MockSleepRun_Expecter) RunID() *MockSleepRun_RunID_Call {
	return &MockSleepRun_RunID_Call{Call: _e.mock.On("RunID")}
}

func (_c *MockSleepRun_RunID_Call) Run(run func()) *MockSleepRun_RunID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSleepRun_RunID_Call) Return(_a0 string) *MockSleepRun_RunID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSleepRun_RunID_Call) RunAndReturn(run func() string) *MockSleepRun_RunID_Call {
	_c.Call.Return(run)
	return _c
}

// Terminate provides a mock function with given fields: ctx, reason, details
func (_m *MockSleepRun) Terminate(ctx context.Context, reason string, details ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, reason)
	_ca = append(_ca, details...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Terminate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) error); ok {
		r0 = rf(ctx, reason, details...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSleepRun_Terminate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Terminate'
type MockSleepRun_Terminate_Call struct {
	*mock.Call
}

// Terminate is a helper method to define mock.On call
//   - ctx context.Context
//   - reason string
//   - details ...interface{}
func (_e *MockSleepRun_Expecter) Terminate(ctx interface{}, reason interface{}, details ...interface{}) *MockSleepRun_Terminate_Call {
	return &MockSleepRun_Terminate_Call{Call: _e.mock.On("Terminate",
		append([]interface{}{ctx, reason}, details...)...)}
}

func (_c *MockSleepRun_Terminate_Call) Run(run func(ctx context.Context, reason string, details ...interface{})) *MockSleepRun_Terminate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSleepRun_Terminate_Call) Return(_a0 error) *MockSleepRun_Terminate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSleepRun_Terminate_Call) RunAndReturn(run func(context.Context, string, ...interface{}) error) *MockSleepRun_Terminate_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSleepRun creates a new instance of MockSleepRun. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSleepRun(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSleepRun {
	mock := &MockSleepRun{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
