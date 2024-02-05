// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IProfileService is an autogenerated mock type for the IProfileService type
type IProfileService struct {
	mock.Mock
}

// ServiceUpdateProfileName provides a mock function with given fields: i, name
func (_m *IProfileService) ServiceUpdateProfileName(i int64, name string) error {
	ret := _m.Called(i, name)

	if len(ret) == 0 {
		panic("no return value specified for ServiceUpdateProfileName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, string) error); ok {
		r0 = rf(i, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIProfileService creates a new instance of IProfileService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIProfileService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IProfileService {
	mock := &IProfileService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
