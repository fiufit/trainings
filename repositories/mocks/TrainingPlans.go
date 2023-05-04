// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/fiufit/trainings/models"
	mock "github.com/stretchr/testify/mock"

	training "github.com/fiufit/trainings/contracts/training"
)

// TrainingPlans is an autogenerated mock type for the TrainingPlans type
type TrainingPlans struct {
	mock.Mock
}

// CreateTrainingPlan provides a mock function with given fields: ctx, _a1
func (_m *TrainingPlans) CreateTrainingPlan(ctx context.Context, _a1 models.TrainingPlan) (models.TrainingPlan, error) {
	ret := _m.Called(ctx, _a1)

	var r0 models.TrainingPlan
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.TrainingPlan) (models.TrainingPlan, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.TrainingPlan) models.TrainingPlan); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(models.TrainingPlan)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.TrainingPlan) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTrainingPlans provides a mock function with given fields: ctx, req
func (_m *TrainingPlans) GetTrainingPlans(ctx context.Context, req training.GetTrainingsRequest) (training.GetTrainingsResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 training.GetTrainingsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, training.GetTrainingsRequest) (training.GetTrainingsResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, training.GetTrainingsRequest) training.GetTrainingsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(training.GetTrainingsResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, training.GetTrainingsRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTrainingPlans interface {
	mock.TestingT
	Cleanup(func())
}

// NewTrainingPlans creates a new instance of TrainingPlans. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTrainingPlans(t mockConstructorTestingTNewTrainingPlans) *TrainingPlans {
	mock := &TrainingPlans{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
