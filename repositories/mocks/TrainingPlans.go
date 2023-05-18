// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/fiufit/trainings/models"
	mock "github.com/stretchr/testify/mock"

	trainings "github.com/fiufit/trainings/contracts/trainings"
)

// TrainingPlans is an autogenerated mock type for the TrainingPlans type
type TrainingPlans struct {
	mock.Mock
}

// CreateTrainingPlan provides a mock function with given fields: ctx, training
func (_m *TrainingPlans) CreateTrainingPlan(ctx context.Context, training models.TrainingPlan) (models.TrainingPlan, error) {
	ret := _m.Called(ctx, training)

	var r0 models.TrainingPlan
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.TrainingPlan) (models.TrainingPlan, error)); ok {
		return rf(ctx, training)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.TrainingPlan) models.TrainingPlan); ok {
		r0 = rf(ctx, training)
	} else {
		r0 = ret.Get(0).(models.TrainingPlan)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.TrainingPlan) error); ok {
		r1 = rf(ctx, training)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTrainingPlan provides a mock function with given fields: ctx, trainingID
func (_m *TrainingPlans) DeleteTrainingPlan(ctx context.Context, trainingID uint) error {
	ret := _m.Called(ctx, trainingID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, trainingID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTrainingByID provides a mock function with given fields: ctx, trainingID
func (_m *TrainingPlans) GetTrainingByID(ctx context.Context, trainingID uint) (models.TrainingPlan, error) {
	ret := _m.Called(ctx, trainingID)

	var r0 models.TrainingPlan
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) (models.TrainingPlan, error)); ok {
		return rf(ctx, trainingID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) models.TrainingPlan); ok {
		r0 = rf(ctx, trainingID)
	} else {
		r0 = ret.Get(0).(models.TrainingPlan)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, trainingID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTrainingPlans provides a mock function with given fields: ctx, req
func (_m *TrainingPlans) GetTrainingPlans(ctx context.Context, req trainings.GetTrainingsRequest) (trainings.GetTrainingsResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 trainings.GetTrainingsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, trainings.GetTrainingsRequest) (trainings.GetTrainingsResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, trainings.GetTrainingsRequest) trainings.GetTrainingsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(trainings.GetTrainingsResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, trainings.GetTrainingsRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTrainingPlan provides a mock function with given fields: ctx, training
func (_m *TrainingPlans) UpdateTrainingPlan(ctx context.Context, training models.TrainingPlan) (models.TrainingPlan, error) {
	ret := _m.Called(ctx, training)

	var r0 models.TrainingPlan
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.TrainingPlan) (models.TrainingPlan, error)); ok {
		return rf(ctx, training)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.TrainingPlan) models.TrainingPlan); ok {
		r0 = rf(ctx, training)
	} else {
		r0 = ret.Get(0).(models.TrainingPlan)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.TrainingPlan) error); ok {
		r1 = rf(ctx, training)
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
