package order

import (
	"context"
	"shop-app/api/v1/dto"
	"shop-app/domain/order"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) error
	GetOrder(ctx context.Context, orderID string) (*dto.OrderResponse, error)
}

type orderUsecase struct {
	repo order.OrderRepository
}

func NewOrderUsecase(repo order.OrderRepository) *orderUsecase {
	return &orderUsecase{repo: repo}
}

func (u *orderUsecase) CreateOrder(ctx context.Context, req *dto.CreateOrderRequest) error {
	return u.repo.CreateOrder(ctx, &req.Order)
}

func (u *orderUsecase) GetOrder(ctx context.Context, orderID string) (*dto.OrderResponse, error) {
	dbOrder, err := u.repo.GetOrder(ctx, orderID)
	if err != nil {
		return nil, err
	}
	return dto.ConvertToOrderResponse(dbOrder), nil
}