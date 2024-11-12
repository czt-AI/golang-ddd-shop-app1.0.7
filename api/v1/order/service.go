package v1

import (
	"context"
	"shop-app/api/v1/dto"
	"shop-app/domain/order"
)

type OrderService interface {
	CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) error
	GetOrder(ctx context.Context, orderID string) (*dto.OrderResponse, error)
}

type orderService struct {
	usecase order.OrderUsecase
}

func NewOrderService(usecase order.OrderUsecase) *orderService {
	return &orderService{usecase: usecase}
}

func (s *orderService) CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) error {
	return s.usecase.CreateOrder(ctx, req)
}

func (s *orderService) GetOrder(ctx context.Context, orderID string) (*dto.OrderResponse, error) {
	return s.usecase.GetOrder(ctx, orderID)
}