// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/jfelipearaujo-org/ms-order-management/internal/entity"
	mock "github.com/stretchr/testify/mock"

	service "github.com/jfelipearaujo-org/ms-order-management/internal/service"
)

// MockGetOrdersService is an autogenerated mock type for the GetOrdersService type
type MockGetOrdersService struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, request
func (_m *MockGetOrdersService) Handle(ctx context.Context, request service.GetOrdersDto) ([]entity.Order, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 []entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, service.GetOrdersDto) ([]entity.Order, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, service.GetOrdersDto) []entity.Order); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, service.GetOrdersDto) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockGetOrdersService creates a new instance of MockGetOrdersService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGetOrdersService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGetOrdersService {
	mock := &MockGetOrdersService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
