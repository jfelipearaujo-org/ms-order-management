package service

import (
	"context"

	"github.com/jfelipearaujo-org/ms-order-management/internal/common"
	"github.com/jfelipearaujo-org/ms-order-management/internal/entity"
)

type CreateOrderService[T any] interface {
	Handle(ctx context.Context, request T) (*entity.Order, error)
}

type GetOrderService[T any] interface {
	Handle(ctx context.Context, request T) (entity.Order, error)
}

type GetOrdersDto struct {
	State int `query:"state"`

	common.Pagination
}

type GetOrdersService interface {
	Handle(ctx context.Context, request GetOrdersDto) ([]entity.Order, error)
}

type UpdateOrderDto struct {
	UUID string `param:"id"`

	State int           `json:"state"`
	Items []entity.Item `json:"items"`
}

type UpdateOrderService interface {
	Handle(ctx context.Context, request UpdateOrderDto) (*entity.Order, error)
}
