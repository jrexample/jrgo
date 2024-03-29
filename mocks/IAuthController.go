// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// IAuthController is an autogenerated mock type for the IAuthController type
type IAuthController struct {
	mock.Mock
}

// ControllerLogin provides a mock function with given fields: c
func (_m *IAuthController) ControllerLogin(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for ControllerLogin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ControllerRefreshToken provides a mock function with given fields: c
func (_m *IAuthController) ControllerRefreshToken(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for ControllerRefreshToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ControllerRegister provides a mock function with given fields: c
func (_m *IAuthController) ControllerRegister(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for ControllerRegister")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIAuthController creates a new instance of IAuthController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAuthController(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAuthController {
	mock := &IAuthController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
