// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/jfelipearaujo-org/ms-order-management/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockGetOrderService is an autogenerated mock type for the GetOrderService type
type MockGetOrderService[T interface{}] struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, request
func (_m *MockGetOrderService[T]) Handle(ctx context.Context, request T) (entity.Order, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, T) (entity.Order, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, T) entity.Order); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(entity.Order)
	}

	if rf, ok := ret.Get(1).(func(context.Context, T) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockGetOrderService creates a new instance of MockGetOrderService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGetOrderService[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGetOrderService[T] {
	mock := &MockGetOrderService[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
