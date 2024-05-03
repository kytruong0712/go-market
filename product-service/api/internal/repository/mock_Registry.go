// Code generated by mockery v2.43.0. DO NOT EDIT.

package repository

import (
	context "context"

	category "github.com/kytruong0712/go-market/product-service/api/internal/repository/category"

	mock "github.com/stretchr/testify/mock"
)

// MockRegistry is an autogenerated mock type for the Registry type
type MockRegistry struct {
	mock.Mock
}

// Category provides a mock function with given fields:
func (_m *MockRegistry) Category() category.Repository {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Category")
	}

	var r0 category.Repository
	if rf, ok := ret.Get(0).(func() category.Repository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(category.Repository)
		}
	}

	return r0
}

// PingPG provides a mock function with given fields: _a0
func (_m *MockRegistry) PingPG(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for PingPG")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockRegistry creates a new instance of MockRegistry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRegistry(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRegistry {
	mock := &MockRegistry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}