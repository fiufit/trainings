// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/fiufit/trainings/models"
	mock "github.com/stretchr/testify/mock"
)

// Exercises is an autogenerated mock type for the Exercises type
type Exercises struct {
	mock.Mock
}

// CreateExercise provides a mock function with given fields: ctx, exercise
func (_m *Exercises) CreateExercise(ctx context.Context, exercise models.Exercise) (models.Exercise, error) {
	ret := _m.Called(ctx, exercise)

	var r0 models.Exercise
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Exercise) (models.Exercise, error)); ok {
		return rf(ctx, exercise)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.Exercise) models.Exercise); ok {
		r0 = rf(ctx, exercise)
	} else {
		r0 = ret.Get(0).(models.Exercise)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.Exercise) error); ok {
		r1 = rf(ctx, exercise)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteExercise provides a mock function with given fields: ctx, exerciseID
func (_m *Exercises) DeleteExercise(ctx context.Context, exerciseID string) error {
	ret := _m.Called(ctx, exerciseID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, exerciseID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetExerciseByID provides a mock function with given fields: ctx, exerciseID
func (_m *Exercises) GetExerciseByID(ctx context.Context, exerciseID string) (models.Exercise, error) {
	ret := _m.Called(ctx, exerciseID)

	var r0 models.Exercise
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (models.Exercise, error)); ok {
		return rf(ctx, exerciseID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Exercise); ok {
		r0 = rf(ctx, exerciseID)
	} else {
		r0 = ret.Get(0).(models.Exercise)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, exerciseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateExercise provides a mock function with given fields: ctx, exercise
func (_m *Exercises) UpdateExercise(ctx context.Context, exercise models.Exercise) (models.Exercise, error) {
	ret := _m.Called(ctx, exercise)

	var r0 models.Exercise
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Exercise) (models.Exercise, error)); ok {
		return rf(ctx, exercise)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.Exercise) models.Exercise); ok {
		r0 = rf(ctx, exercise)
	} else {
		r0 = ret.Get(0).(models.Exercise)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.Exercise) error); ok {
		r1 = rf(ctx, exercise)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewExercises interface {
	mock.TestingT
	Cleanup(func())
}

// NewExercises creates a new instance of Exercises. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExercises(t mockConstructorTestingTNewExercises) *Exercises {
	mock := &Exercises{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
