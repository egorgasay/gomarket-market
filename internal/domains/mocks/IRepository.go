// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	model "go-rest-api/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// IRepository is an autogenerated mock type for the IRepository type
type IRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: user
func (_m *IRepository) CreateUser(user model.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRepository creates a new instance of IRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRepository(t mockConstructorTestingTNewIRepository) *IRepository {
	mock := &IRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
