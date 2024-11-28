// Code generated by mockery v2.51.0. DO NOT EDIT.

package examplev1

import (
	context "context"

	client "go.temporal.io/sdk/client"

	examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"

	mock "github.com/stretchr/testify/mock"
)

// MockCreateFooRun is an autogenerated mock type for the CreateFooRun type
type MockCreateFooRun struct {
	mock.Mock
}

type MockCreateFooRun_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCreateFooRun) EXPECT() *MockCreateFooRun_Expecter {
	return &MockCreateFooRun_Expecter{mock: &_m.Mock}
}

// Cancel provides a mock function with given fields: ctx
func (_m *MockCreateFooRun) Cancel(ctx context.Context) error {
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

// MockCreateFooRun_Cancel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Cancel'
type MockCreateFooRun_Cancel_Call struct {
	*mock.Call
}

// Cancel is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockCreateFooRun_Expecter) Cancel(ctx interface{}) *MockCreateFooRun_Cancel_Call {
	return &MockCreateFooRun_Cancel_Call{Call: _e.mock.On("Cancel", ctx)}
}

func (_c *MockCreateFooRun_Cancel_Call) Run(run func(ctx context.Context)) *MockCreateFooRun_Cancel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockCreateFooRun_Cancel_Call) Return(_a0 error) *MockCreateFooRun_Cancel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCreateFooRun_Cancel_Call) RunAndReturn(run func(context.Context) error) *MockCreateFooRun_Cancel_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx
func (_m *MockCreateFooRun) Get(ctx context.Context) (*examplev1.CreateFooResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *examplev1.CreateFooResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*examplev1.CreateFooResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *examplev1.CreateFooResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*examplev1.CreateFooResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCreateFooRun_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockCreateFooRun_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockCreateFooRun_Expecter) Get(ctx interface{}) *MockCreateFooRun_Get_Call {
	return &MockCreateFooRun_Get_Call{Call: _e.mock.On("Get", ctx)}
}

func (_c *MockCreateFooRun_Get_Call) Run(run func(ctx context.Context)) *MockCreateFooRun_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockCreateFooRun_Get_Call) Return(_a0 *examplev1.CreateFooResponse, _a1 error) *MockCreateFooRun_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCreateFooRun_Get_Call) RunAndReturn(run func(context.Context) (*examplev1.CreateFooResponse, error)) *MockCreateFooRun_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetFooProgress provides a mock function with given fields: ctx
func (_m *MockCreateFooRun) GetFooProgress(ctx context.Context) (*examplev1.GetFooProgressResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetFooProgress")
	}

	var r0 *examplev1.GetFooProgressResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*examplev1.GetFooProgressResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *examplev1.GetFooProgressResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*examplev1.GetFooProgressResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCreateFooRun_GetFooProgress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFooProgress'
type MockCreateFooRun_GetFooProgress_Call struct {
	*mock.Call
}

// GetFooProgress is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockCreateFooRun_Expecter) GetFooProgress(ctx interface{}) *MockCreateFooRun_GetFooProgress_Call {
	return &MockCreateFooRun_GetFooProgress_Call{Call: _e.mock.On("GetFooProgress", ctx)}
}

func (_c *MockCreateFooRun_GetFooProgress_Call) Run(run func(ctx context.Context)) *MockCreateFooRun_GetFooProgress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockCreateFooRun_GetFooProgress_Call) Return(_a0 *examplev1.GetFooProgressResponse, _a1 error) *MockCreateFooRun_GetFooProgress_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCreateFooRun_GetFooProgress_Call) RunAndReturn(run func(context.Context) (*examplev1.GetFooProgressResponse, error)) *MockCreateFooRun_GetFooProgress_Call {
	_c.Call.Return(run)
	return _c
}

// ID provides a mock function with no fields
func (_m *MockCreateFooRun) ID() string {
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

// MockCreateFooRun_ID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ID'
type MockCreateFooRun_ID_Call struct {
	*mock.Call
}

// ID is a helper method to define mock.On call
func (_e *MockCreateFooRun_Expecter) ID() *MockCreateFooRun_ID_Call {
	return &MockCreateFooRun_ID_Call{Call: _e.mock.On("ID")}
}

func (_c *MockCreateFooRun_ID_Call) Run(run func()) *MockCreateFooRun_ID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCreateFooRun_ID_Call) Return(_a0 string) *MockCreateFooRun_ID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCreateFooRun_ID_Call) RunAndReturn(run func() string) *MockCreateFooRun_ID_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with no fields
func (_m *MockCreateFooRun) Run() client.WorkflowRun {
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

// MockCreateFooRun_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockCreateFooRun_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
func (_e *MockCreateFooRun_Expecter) Run() *MockCreateFooRun_Run_Call {
	return &MockCreateFooRun_Run_Call{Call: _e.mock.On("Run")}
}

func (_c *MockCreateFooRun_Run_Call) Run(run func()) *MockCreateFooRun_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCreateFooRun_Run_Call) Return(_a0 client.WorkflowRun) *MockCreateFooRun_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCreateFooRun_Run_Call) RunAndReturn(run func() client.WorkflowRun) *MockCreateFooRun_Run_Call {
	_c.Call.Return(run)
	return _c
}

// RunID provides a mock function with no fields
func (_m *MockCreateFooRun) RunID() string {
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

// MockCreateFooRun_RunID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunID'
type MockCreateFooRun_RunID_Call struct {
	*mock.Call
}

// RunID is a helper method to define mock.On call
func (_e *MockCreateFooRun_Expecter) RunID() *MockCreateFooRun_RunID_Call {
	return &MockCreateFooRun_RunID_Call{Call: _e.mock.On("RunID")}
}

func (_c *MockCreateFooRun_RunID_Call) Run(run func()) *MockCreateFooRun_RunID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCreateFooRun_RunID_Call) Return(_a0 string) *MockCreateFooRun_RunID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCreateFooRun_RunID_Call) RunAndReturn(run func() string) *MockCreateFooRun_RunID_Call {
	_c.Call.Return(run)
	return _c
}

// SetFooProgress provides a mock function with given fields: ctx, req
func (_m *MockCreateFooRun) SetFooProgress(ctx context.Context, req *examplev1.SetFooProgressRequest) error {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for SetFooProgress")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *examplev1.SetFooProgressRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCreateFooRun_SetFooProgress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetFooProgress'
type MockCreateFooRun_SetFooProgress_Call struct {
	*mock.Call
}

// SetFooProgress is a helper method to define mock.On call
//   - ctx context.Context
//   - req *examplev1.SetFooProgressRequest
func (_e *MockCreateFooRun_Expecter) SetFooProgress(ctx interface{}, req interface{}) *MockCreateFooRun_SetFooProgress_Call {
	return &MockCreateFooRun_SetFooProgress_Call{Call: _e.mock.On("SetFooProgress", ctx, req)}
}

func (_c *MockCreateFooRun_SetFooProgress_Call) Run(run func(ctx context.Context, req *examplev1.SetFooProgressRequest)) *MockCreateFooRun_SetFooProgress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*examplev1.SetFooProgressRequest))
	})
	return _c
}

func (_c *MockCreateFooRun_SetFooProgress_Call) Return(_a0 error) *MockCreateFooRun_SetFooProgress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCreateFooRun_SetFooProgress_Call) RunAndReturn(run func(context.Context, *examplev1.SetFooProgressRequest) error) *MockCreateFooRun_SetFooProgress_Call {
	_c.Call.Return(run)
	return _c
}

// Terminate provides a mock function with given fields: ctx, reason, details
func (_m *MockCreateFooRun) Terminate(ctx context.Context, reason string, details ...interface{}) error {
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

// MockCreateFooRun_Terminate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Terminate'
type MockCreateFooRun_Terminate_Call struct {
	*mock.Call
}

// Terminate is a helper method to define mock.On call
//   - ctx context.Context
//   - reason string
//   - details ...interface{}
func (_e *MockCreateFooRun_Expecter) Terminate(ctx interface{}, reason interface{}, details ...interface{}) *MockCreateFooRun_Terminate_Call {
	return &MockCreateFooRun_Terminate_Call{Call: _e.mock.On("Terminate",
		append([]interface{}{ctx, reason}, details...)...)}
}

func (_c *MockCreateFooRun_Terminate_Call) Run(run func(ctx context.Context, reason string, details ...interface{})) *MockCreateFooRun_Terminate_Call {
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

func (_c *MockCreateFooRun_Terminate_Call) Return(_a0 error) *MockCreateFooRun_Terminate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCreateFooRun_Terminate_Call) RunAndReturn(run func(context.Context, string, ...interface{}) error) *MockCreateFooRun_Terminate_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateFooProgress provides a mock function with given fields: ctx, req, opts
func (_m *MockCreateFooRun) UpdateFooProgress(ctx context.Context, req *examplev1.SetFooProgressRequest, opts ...*examplev1.UpdateFooProgressOptions) (*examplev1.GetFooProgressResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateFooProgress")
	}

	var r0 *examplev1.GetFooProgressResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *examplev1.SetFooProgressRequest, ...*examplev1.UpdateFooProgressOptions) (*examplev1.GetFooProgressResponse, error)); ok {
		return rf(ctx, req, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *examplev1.SetFooProgressRequest, ...*examplev1.UpdateFooProgressOptions) *examplev1.GetFooProgressResponse); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*examplev1.GetFooProgressResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *examplev1.SetFooProgressRequest, ...*examplev1.UpdateFooProgressOptions) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCreateFooRun_UpdateFooProgress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateFooProgress'
type MockCreateFooRun_UpdateFooProgress_Call struct {
	*mock.Call
}

// UpdateFooProgress is a helper method to define mock.On call
//   - ctx context.Context
//   - req *examplev1.SetFooProgressRequest
//   - opts ...*examplev1.UpdateFooProgressOptions
func (_e *MockCreateFooRun_Expecter) UpdateFooProgress(ctx interface{}, req interface{}, opts ...interface{}) *MockCreateFooRun_UpdateFooProgress_Call {
	return &MockCreateFooRun_UpdateFooProgress_Call{Call: _e.mock.On("UpdateFooProgress",
		append([]interface{}{ctx, req}, opts...)...)}
}

func (_c *MockCreateFooRun_UpdateFooProgress_Call) Run(run func(ctx context.Context, req *examplev1.SetFooProgressRequest, opts ...*examplev1.UpdateFooProgressOptions)) *MockCreateFooRun_UpdateFooProgress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*examplev1.UpdateFooProgressOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(*examplev1.UpdateFooProgressOptions)
			}
		}
		run(args[0].(context.Context), args[1].(*examplev1.SetFooProgressRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockCreateFooRun_UpdateFooProgress_Call) Return(_a0 *examplev1.GetFooProgressResponse, _a1 error) *MockCreateFooRun_UpdateFooProgress_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCreateFooRun_UpdateFooProgress_Call) RunAndReturn(run func(context.Context, *examplev1.SetFooProgressRequest, ...*examplev1.UpdateFooProgressOptions) (*examplev1.GetFooProgressResponse, error)) *MockCreateFooRun_UpdateFooProgress_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateFooProgressAsync provides a mock function with given fields: ctx, req, opts
func (_m *MockCreateFooRun) UpdateFooProgressAsync(ctx context.Context, req *examplev1.SetFooProgressRequest, opts ...*examplev1.UpdateFooProgressOptions) (examplev1.UpdateFooProgressHandle, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateFooProgressAsync")
	}

	var r0 examplev1.UpdateFooProgressHandle
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *examplev1.SetFooProgressRequest, ...*examplev1.UpdateFooProgressOptions) (examplev1.UpdateFooProgressHandle, error)); ok {
		return rf(ctx, req, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *examplev1.SetFooProgressRequest, ...*examplev1.UpdateFooProgressOptions) examplev1.UpdateFooProgressHandle); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(examplev1.UpdateFooProgressHandle)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *examplev1.SetFooProgressRequest, ...*examplev1.UpdateFooProgressOptions) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCreateFooRun_UpdateFooProgressAsync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateFooProgressAsync'
type MockCreateFooRun_UpdateFooProgressAsync_Call struct {
	*mock.Call
}

// UpdateFooProgressAsync is a helper method to define mock.On call
//   - ctx context.Context
//   - req *examplev1.SetFooProgressRequest
//   - opts ...*examplev1.UpdateFooProgressOptions
func (_e *MockCreateFooRun_Expecter) UpdateFooProgressAsync(ctx interface{}, req interface{}, opts ...interface{}) *MockCreateFooRun_UpdateFooProgressAsync_Call {
	return &MockCreateFooRun_UpdateFooProgressAsync_Call{Call: _e.mock.On("UpdateFooProgressAsync",
		append([]interface{}{ctx, req}, opts...)...)}
}

func (_c *MockCreateFooRun_UpdateFooProgressAsync_Call) Run(run func(ctx context.Context, req *examplev1.SetFooProgressRequest, opts ...*examplev1.UpdateFooProgressOptions)) *MockCreateFooRun_UpdateFooProgressAsync_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*examplev1.UpdateFooProgressOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(*examplev1.UpdateFooProgressOptions)
			}
		}
		run(args[0].(context.Context), args[1].(*examplev1.SetFooProgressRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockCreateFooRun_UpdateFooProgressAsync_Call) Return(_a0 examplev1.UpdateFooProgressHandle, _a1 error) *MockCreateFooRun_UpdateFooProgressAsync_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCreateFooRun_UpdateFooProgressAsync_Call) RunAndReturn(run func(context.Context, *examplev1.SetFooProgressRequest, ...*examplev1.UpdateFooProgressOptions) (examplev1.UpdateFooProgressHandle, error)) *MockCreateFooRun_UpdateFooProgressAsync_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCreateFooRun creates a new instance of MockCreateFooRun. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCreateFooRun(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCreateFooRun {
	mock := &MockCreateFooRun{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
