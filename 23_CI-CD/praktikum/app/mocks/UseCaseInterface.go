// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	user "belajar-go-echo/features/user"

	mock "github.com/stretchr/testify/mock"
)

// UseCaseInterface is an autogenerated mock type for the UseCaseInterface type
type UseCaseInterface struct {
	mock.Mock
}

// GetAllUsers provides a mock function with given fields:
func (_m *UseCaseInterface) GetAllUsers() ([]user.UserCore, error) {
	ret := _m.Called()

	var r0 []user.UserCore
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]user.UserCore, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []user.UserCore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.UserCore)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: data
func (_m *UseCaseInterface) Insert(data user.UserCore) (int, error) {
	ret := _m.Called(data)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(user.UserCore) (int, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(user.UserCore) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(user.UserCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: email, password
func (_m *UseCaseInterface) Login(email string, password string) (user.UserCore, string, error) {
	ret := _m.Called(email, password)

	var r0 user.UserCore
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (user.UserCore, string, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) user.UserCore); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewUseCaseInterface creates a new instance of UseCaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUseCaseInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UseCaseInterface {
	mock := &UseCaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
