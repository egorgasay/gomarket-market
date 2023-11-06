// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	model "go-rest-api/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// UserUseCase is an autogenerated mock type for the UserUseCase type
type UserUseCase struct {
	mock.Mock
}

// SignUp provides a mock function with given fields: user
func (_m *UserUseCase) SignUp(user model.User) (string, error) {
	ret := _m.Called(user)

	var r0 string
	if rf, ok := ret.Get(0).(func(model.User) string); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUseCase creates a new instance of UserUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUseCase(t mockConstructorTestingTNewUserUseCase) *UserUseCase {
	mock := &UserUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
