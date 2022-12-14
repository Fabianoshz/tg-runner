// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PlanExecutionInterface is an autogenerated mock type for the PlanExecutionInterface type
type PlanExecutionInterface struct {
	mock.Mock
}

// PlanExecution provides a mock function with given fields: _a0
func (_m *PlanExecutionInterface) PlanExecution(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewPlanExecutionInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewPlanExecutionInterface creates a new instance of PlanExecutionInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPlanExecutionInterface(t mockConstructorTestingTNewPlanExecutionInterface) *PlanExecutionInterface {
	mock := &PlanExecutionInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
