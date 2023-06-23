// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	metrics "github.com/fiufit/trainings/contracts/metrics"
	mock "github.com/stretchr/testify/mock"
)

// Metrics is an autogenerated mock type for the Metrics type
type Metrics struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, req
func (_m *Metrics) Create(ctx context.Context, req metrics.CreateMetricRequest) {
	_m.Called(ctx, req)
}

type mockConstructorTestingTNewMetrics interface {
	mock.TestingT
	Cleanup(func())
}

// NewMetrics creates a new instance of Metrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMetrics(t mockConstructorTestingTNewMetrics) *Metrics {
	mock := &Metrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
