// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/fabianoshz/tg-runner/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// LoadResourcesInterface is an autogenerated mock type for the LoadResourcesInterface type
type LoadResourcesInterface struct {
	mock.Mock
}

// LoadResources provides a mock function with given fields: _a0
func (_m *LoadResourcesInterface) LoadResources(_a0 string) []entity.Resource {
	ret := _m.Called(_a0)

	var r0 []entity.Resource
	if rf, ok := ret.Get(0).(func(string) []entity.Resource); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Resource)
		}
	}

	return r0
}

type mockConstructorTestingTNewLoadResourcesInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewLoadResourcesInterface creates a new instance of LoadResourcesInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLoadResourcesInterface(t mockConstructorTestingTNewLoadResourcesInterface) *LoadResourcesInterface {
	mock := &LoadResourcesInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
