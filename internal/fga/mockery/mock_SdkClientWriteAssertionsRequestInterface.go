// Code generated by mockery v2.40.1. DO NOT EDIT.

package client

import (
	context "context"

	client "github.com/openfga/go-sdk/client"

	mock "github.com/stretchr/testify/mock"
)

// MockSdkClientWriteAssertionsRequestInterface is an autogenerated mock type for the SdkClientWriteAssertionsRequestInterface type
type MockSdkClientWriteAssertionsRequestInterface struct {
	mock.Mock
}

type MockSdkClientWriteAssertionsRequestInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSdkClientWriteAssertionsRequestInterface) EXPECT() *MockSdkClientWriteAssertionsRequestInterface_Expecter {
	return &MockSdkClientWriteAssertionsRequestInterface_Expecter{mock: &_m.Mock}
}

// Body provides a mock function with given fields: body
func (_m *MockSdkClientWriteAssertionsRequestInterface) Body(body []client.ClientAssertion) client.SdkClientWriteAssertionsRequestInterface {
	ret := _m.Called(body)

	if len(ret) == 0 {
		panic("no return value specified for Body")
	}

	var r0 client.SdkClientWriteAssertionsRequestInterface
	if rf, ok := ret.Get(0).(func([]client.ClientAssertion) client.SdkClientWriteAssertionsRequestInterface); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SdkClientWriteAssertionsRequestInterface)
		}
	}

	return r0
}

// MockSdkClientWriteAssertionsRequestInterface_Body_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Body'
type MockSdkClientWriteAssertionsRequestInterface_Body_Call struct {
	*mock.Call
}

// Body is a helper method to define mock.On call
//   - body []client.ClientAssertion
func (_e *MockSdkClientWriteAssertionsRequestInterface_Expecter) Body(body interface{}) *MockSdkClientWriteAssertionsRequestInterface_Body_Call {
	return &MockSdkClientWriteAssertionsRequestInterface_Body_Call{Call: _e.mock.On("Body", body)}
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_Body_Call) Run(run func(body []client.ClientAssertion)) *MockSdkClientWriteAssertionsRequestInterface_Body_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]client.ClientAssertion))
	})
	return _c
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_Body_Call) Return(_a0 client.SdkClientWriteAssertionsRequestInterface) *MockSdkClientWriteAssertionsRequestInterface_Body_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_Body_Call) RunAndReturn(run func([]client.ClientAssertion) client.SdkClientWriteAssertionsRequestInterface) *MockSdkClientWriteAssertionsRequestInterface_Body_Call {
	_c.Call.Return(run)
	return _c
}

// Execute provides a mock function with given fields:
func (_m *MockSdkClientWriteAssertionsRequestInterface) Execute() (*client.ClientWriteAssertionsResponse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *client.ClientWriteAssertionsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*client.ClientWriteAssertionsResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *client.ClientWriteAssertionsResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientWriteAssertionsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSdkClientWriteAssertionsRequestInterface_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockSdkClientWriteAssertionsRequestInterface_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
func (_e *MockSdkClientWriteAssertionsRequestInterface_Expecter) Execute() *MockSdkClientWriteAssertionsRequestInterface_Execute_Call {
	return &MockSdkClientWriteAssertionsRequestInterface_Execute_Call{Call: _e.mock.On("Execute")}
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_Execute_Call) Run(run func()) *MockSdkClientWriteAssertionsRequestInterface_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_Execute_Call) Return(_a0 *client.ClientWriteAssertionsResponse, _a1 error) *MockSdkClientWriteAssertionsRequestInterface_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_Execute_Call) RunAndReturn(run func() (*client.ClientWriteAssertionsResponse, error)) *MockSdkClientWriteAssertionsRequestInterface_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// GetAuthorizationModelIdOverride provides a mock function with given fields:
func (_m *MockSdkClientWriteAssertionsRequestInterface) GetAuthorizationModelIdOverride() *string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAuthorizationModelIdOverride")
	}

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

// MockSdkClientWriteAssertionsRequestInterface_GetAuthorizationModelIdOverride_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthorizationModelIdOverride'
type MockSdkClientWriteAssertionsRequestInterface_GetAuthorizationModelIdOverride_Call struct {
	*mock.Call
}

// GetAuthorizationModelIdOverride is a helper method to define mock.On call
func (_e *MockSdkClientWriteAssertionsRequestInterface_Expecter) GetAuthorizationModelIdOverride() *MockSdkClientWriteAssertionsRequestInterface_GetAuthorizationModelIdOverride_Call {
	return &MockSdkClientWriteAssertionsRequestInterface_GetAuthorizationModelIdOverride_Call{Call: _e.mock.On("GetAuthorizationModelIdOverride")}
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_GetAuthorizationModelIdOverride_Call) Run(run func()) *MockSdkClientWriteAssertionsRequestInterface_GetAuthorizationModelIdOverride_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_GetAuthorizationModelIdOverride_Call) Return(_a0 *string) *MockSdkClientWriteAssertionsRequestInterface_GetAuthorizationModelIdOverride_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_GetAuthorizationModelIdOverride_Call) RunAndReturn(run func() *string) *MockSdkClientWriteAssertionsRequestInterface_GetAuthorizationModelIdOverride_Call {
	_c.Call.Return(run)
	return _c
}

// GetBody provides a mock function with given fields:
func (_m *MockSdkClientWriteAssertionsRequestInterface) GetBody() *[]client.ClientAssertion {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBody")
	}

	var r0 *[]client.ClientAssertion
	if rf, ok := ret.Get(0).(func() *[]client.ClientAssertion); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]client.ClientAssertion)
		}
	}

	return r0
}

// MockSdkClientWriteAssertionsRequestInterface_GetBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBody'
type MockSdkClientWriteAssertionsRequestInterface_GetBody_Call struct {
	*mock.Call
}

// GetBody is a helper method to define mock.On call
func (_e *MockSdkClientWriteAssertionsRequestInterface_Expecter) GetBody() *MockSdkClientWriteAssertionsRequestInterface_GetBody_Call {
	return &MockSdkClientWriteAssertionsRequestInterface_GetBody_Call{Call: _e.mock.On("GetBody")}
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_GetBody_Call) Run(run func()) *MockSdkClientWriteAssertionsRequestInterface_GetBody_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_GetBody_Call) Return(_a0 *[]client.ClientAssertion) *MockSdkClientWriteAssertionsRequestInterface_GetBody_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_GetBody_Call) RunAndReturn(run func() *[]client.ClientAssertion) *MockSdkClientWriteAssertionsRequestInterface_GetBody_Call {
	_c.Call.Return(run)
	return _c
}

// GetContext provides a mock function with given fields:
func (_m *MockSdkClientWriteAssertionsRequestInterface) GetContext() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetContext")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// MockSdkClientWriteAssertionsRequestInterface_GetContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContext'
type MockSdkClientWriteAssertionsRequestInterface_GetContext_Call struct {
	*mock.Call
}

// GetContext is a helper method to define mock.On call
func (_e *MockSdkClientWriteAssertionsRequestInterface_Expecter) GetContext() *MockSdkClientWriteAssertionsRequestInterface_GetContext_Call {
	return &MockSdkClientWriteAssertionsRequestInterface_GetContext_Call{Call: _e.mock.On("GetContext")}
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_GetContext_Call) Run(run func()) *MockSdkClientWriteAssertionsRequestInterface_GetContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_GetContext_Call) Return(_a0 context.Context) *MockSdkClientWriteAssertionsRequestInterface_GetContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_GetContext_Call) RunAndReturn(run func() context.Context) *MockSdkClientWriteAssertionsRequestInterface_GetContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetOptions provides a mock function with given fields:
func (_m *MockSdkClientWriteAssertionsRequestInterface) GetOptions() *client.ClientWriteAssertionsOptions {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetOptions")
	}

	var r0 *client.ClientWriteAssertionsOptions
	if rf, ok := ret.Get(0).(func() *client.ClientWriteAssertionsOptions); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientWriteAssertionsOptions)
		}
	}

	return r0
}

// MockSdkClientWriteAssertionsRequestInterface_GetOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOptions'
type MockSdkClientWriteAssertionsRequestInterface_GetOptions_Call struct {
	*mock.Call
}

// GetOptions is a helper method to define mock.On call
func (_e *MockSdkClientWriteAssertionsRequestInterface_Expecter) GetOptions() *MockSdkClientWriteAssertionsRequestInterface_GetOptions_Call {
	return &MockSdkClientWriteAssertionsRequestInterface_GetOptions_Call{Call: _e.mock.On("GetOptions")}
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_GetOptions_Call) Run(run func()) *MockSdkClientWriteAssertionsRequestInterface_GetOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_GetOptions_Call) Return(_a0 *client.ClientWriteAssertionsOptions) *MockSdkClientWriteAssertionsRequestInterface_GetOptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_GetOptions_Call) RunAndReturn(run func() *client.ClientWriteAssertionsOptions) *MockSdkClientWriteAssertionsRequestInterface_GetOptions_Call {
	_c.Call.Return(run)
	return _c
}

// Options provides a mock function with given fields: options
func (_m *MockSdkClientWriteAssertionsRequestInterface) Options(options client.ClientWriteAssertionsOptions) client.SdkClientWriteAssertionsRequestInterface {
	ret := _m.Called(options)

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 client.SdkClientWriteAssertionsRequestInterface
	if rf, ok := ret.Get(0).(func(client.ClientWriteAssertionsOptions) client.SdkClientWriteAssertionsRequestInterface); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SdkClientWriteAssertionsRequestInterface)
		}
	}

	return r0
}

// MockSdkClientWriteAssertionsRequestInterface_Options_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Options'
type MockSdkClientWriteAssertionsRequestInterface_Options_Call struct {
	*mock.Call
}

// Options is a helper method to define mock.On call
//   - options client.ClientWriteAssertionsOptions
func (_e *MockSdkClientWriteAssertionsRequestInterface_Expecter) Options(options interface{}) *MockSdkClientWriteAssertionsRequestInterface_Options_Call {
	return &MockSdkClientWriteAssertionsRequestInterface_Options_Call{Call: _e.mock.On("Options", options)}
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_Options_Call) Run(run func(options client.ClientWriteAssertionsOptions)) *MockSdkClientWriteAssertionsRequestInterface_Options_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(client.ClientWriteAssertionsOptions))
	})
	return _c
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_Options_Call) Return(_a0 client.SdkClientWriteAssertionsRequestInterface) *MockSdkClientWriteAssertionsRequestInterface_Options_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientWriteAssertionsRequestInterface_Options_Call) RunAndReturn(run func(client.ClientWriteAssertionsOptions) client.SdkClientWriteAssertionsRequestInterface) *MockSdkClientWriteAssertionsRequestInterface_Options_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSdkClientWriteAssertionsRequestInterface creates a new instance of MockSdkClientWriteAssertionsRequestInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSdkClientWriteAssertionsRequestInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSdkClientWriteAssertionsRequestInterface {
	mock := &MockSdkClientWriteAssertionsRequestInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
