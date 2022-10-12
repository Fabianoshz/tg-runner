// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/fabianoshz/iflantis/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// ExecutionPlanner is an autogenerated mock type for the ExecutionPlanner type
type ExecutionPlanner struct {
	mock.Mock
}

// PlanExecution provides a mock function with given fields: _a0
func (_m *ExecutionPlanner) PlanExecution(_a0 entity.Changelist) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(entity.Changelist) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewExecutionPlanner interface {
	mock.TestingT
	Cleanup(func())
}

// NewExecutionPlanner creates a new instance of ExecutionPlanner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExecutionPlanner(t mockConstructorTestingTNewExecutionPlanner) *ExecutionPlanner {
	mock := &ExecutionPlanner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}