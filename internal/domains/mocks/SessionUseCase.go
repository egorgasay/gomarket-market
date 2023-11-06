// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SessionUseCase is an autogenerated mock type for the SessionUseCase type
type SessionUseCase struct {
	mock.Mock
}

// Generate provides a mock function with given fields:
func (_m *SessionUseCase) Generate() (string, error) {
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

type mockConstructorTestingTNewSessionUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewSessionUseCase creates a new instance of SessionUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSessionUseCase(t mockConstructorTestingTNewSessionUseCase) *SessionUseCase {
	mock := &SessionUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
