// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/jfelipearaujo-org/ms-order-management/internal/common"

	mock "github.com/stretchr/testify/mock"

	order_entity "github.com/jfelipearaujo-org/ms-order-management/internal/entity/order_entity"

	repository "github.com/jfelipearaujo-org/ms-order-management/internal/repository"
)

// MockOrderRepository is an autogenerated mock type for the OrderRepository type
type MockOrderRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, order
func (_m *MockOrderRepository) Create(ctx context.Context, order *order_entity.Order) error {
	ret := _m.Called(ctx, order)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *order_entity.Order) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, pagination, filter
func (_m *MockOrderRepository) GetAll(ctx context.Context, pagination common.Pagination, filter repository.GetAllOrdersFilter) (int, []order_entity.Order, error) {
	ret := _m.Called(ctx, pagination, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 int
	var r1 []order_entity.Order
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Pagination, repository.GetAllOrdersFilter) (int, []order_entity.Order, error)); ok {
		return rf(ctx, pagination, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Pagination, repository.GetAllOrdersFilter) int); ok {
		r0 = rf(ctx, pagination, filter)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Pagination, repository.GetAllOrdersFilter) []order_entity.Order); ok {
		r1 = rf(ctx, pagination, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]order_entity.Order)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, common.Pagination, repository.GetAllOrdersFilter) error); ok {
		r2 = rf(ctx, pagination, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByCustomerID provides a mock function with given fields: ctx, customerId
func (_m *MockOrderRepository) GetByCustomerID(ctx context.Context, customerId string) (order_entity.Order, error) {
	ret := _m.Called(ctx, customerId)

	if len(ret) == 0 {
		panic("no return value specified for GetByCustomerID")
	}

	var r0 order_entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (order_entity.Order, error)); ok {
		return rf(ctx, customerId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) order_entity.Order); ok {
		r0 = rf(ctx, customerId)
	} else {
		r0 = ret.Get(0).(order_entity.Order)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, customerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *MockOrderRepository) GetByID(ctx context.Context, id string) (order_entity.Order, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 order_entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (order_entity.Order, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) order_entity.Order); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(order_entity.Order)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTrackID provides a mock function with given fields: ctx, trackId
func (_m *MockOrderRepository) GetByTrackID(ctx context.Context, trackId string) (order_entity.Order, error) {
	ret := _m.Called(ctx, trackId)

	if len(ret) == 0 {
		panic("no return value specified for GetByTrackID")
	}

	var r0 order_entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (order_entity.Order, error)); ok {
		return rf(ctx, trackId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) order_entity.Order); ok {
		r0 = rf(ctx, trackId)
	} else {
		r0 = ret.Get(0).(order_entity.Order)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, trackId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, order, updateItems
func (_m *MockOrderRepository) Update(ctx context.Context, order *order_entity.Order, updateItems bool) error {
	ret := _m.Called(ctx, order, updateItems)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *order_entity.Order, bool) error); ok {
		r0 = rf(ctx, order, updateItems)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockOrderRepository creates a new instance of MockOrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOrderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOrderRepository {
	mock := &MockOrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
