// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SessionService is an autogenerated mock type for the SessionService type
type SessionService struct {
	mock.Mock
}

// Generate provides a mock function with given fields:
func (_m *SessionService) Generate() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSessionService interface {
	mock.TestingT
	Cleanup(func())
}

// NewSessionService creates a new instance of SessionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSessionService(t mockConstructorTestingTNewSessionService) *SessionService {
	mock := &SessionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
