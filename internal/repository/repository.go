package repository

import (
	"context"

	"github.com/jfelipearaujo-org/ms-order-management/internal/common"
	"github.com/jfelipearaujo-org/ms-order-management/internal/entity"
)

type OrderRepository interface {
	Create(ctx context.Context, order *entity.Order) error
	GetByID(ctx context.Context, id string) (entity.Order, error)
	GetByTrackID(ctx context.Context, trackId string) (entity.Order, error)
	GetAll(ctx context.Context, pagination common.Pagination, filter GetAllOrdersFilter) (int, []entity.Order, error)
	Update(ctx context.Context, order *entity.Order, updateItems bool) error
}
