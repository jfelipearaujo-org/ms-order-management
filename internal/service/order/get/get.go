package get

import (
	"context"

	"github.com/jfelipearaujo-org/ms-order-management/internal/entity"
	"github.com/jfelipearaujo-org/ms-order-management/internal/repository"
)

type Service struct {
	repository repository.OrderRepository
}

func NewService(repository repository.OrderRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Handle(ctx context.Context, request GetOrderDto) (entity.Order, error) {
	if err := request.Validate(); err != nil {
		return entity.Order{}, err
	}

	if request.FindViaID() {
		order, err := s.repository.GetByID(ctx, request.UUID)
		if err != nil {
			return entity.Order{}, err
		}

		return order, nil
	}

	order, err := s.repository.GetByTrackID(ctx, request.TrackId)
	if err != nil {
		return entity.Order{}, err
	}

	return order, nil
}