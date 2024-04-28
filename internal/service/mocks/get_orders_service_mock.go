// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/jfelipearaujo-org/ms-order-management/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockGetOrdersService is an autogenerated mock type for the GetOrdersService type
type MockGetOrdersService[T interface{}] struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, request
func (_m *MockGetOrdersService[T]) Handle(ctx context.Context, request T) (int, []entity.Order, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 int
	var r1 []entity.Order
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, T) (int, []entity.Order, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, T) int); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, T) []entity.Order); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]entity.Order)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, T) error); ok {
		r2 = rf(ctx, request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewMockGetOrdersService creates a new instance of MockGetOrdersService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGetOrdersService[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGetOrdersService[T] {
	mock := &MockGetOrdersService[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
