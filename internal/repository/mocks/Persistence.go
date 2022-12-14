// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Persistence is an autogenerated mock type for the Persistence type
type Persistence struct {
	mock.Mock
}

// GetPlanfiles provides a mock function with given fields: _a0
func (_m *Persistence) GetPlanfiles(_a0 uuid.UUID) {
	_m.Called(_a0)
}

// SavePlanfile provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Persistence) SavePlanfile(_a0 string, _a1 string, _a2 uuid.UUID, _a3 string) {
	_m.Called(_a0, _a1, _a2, _a3)
}

type mockConstructorTestingTNewPersistence interface {
	mock.TestingT
	Cleanup(func())
}

// NewPersistence creates a new instance of Persistence. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPersistence(t mockConstructorTestingTNewPersistence) *Persistence {
	mock := &Persistence{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
