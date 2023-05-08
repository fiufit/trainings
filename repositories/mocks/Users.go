// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	users "github.com/fiufit/trainings/contracts/users"
)

// Users is an autogenerated mock type for the Users type
type Users struct {
	mock.Mock
}

// GetUserByID provides a mock function with given fields: ctx, userID
func (_m *Users) GetUserByID(ctx context.Context, userID string) (users.GetUserResponse, error) {
	ret := _m.Called(ctx, userID)

	var r0 users.GetUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (users.GetUserResponse, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) users.GetUserResponse); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(users.GetUserResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUsers interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsers creates a new instance of Users. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsers(t mockConstructorTestingTNewUsers) *Users {
	mock := &Users{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}