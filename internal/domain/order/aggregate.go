package order

import (
	"context"
	"shop-app/domain/order"
)

type OrderAggregate struct {
	repo order.OrderRepository
}

func NewOrderAggregate(repo order.OrderRepository) *OrderAggregate {
	return &OrderAggregate{repo: repo}
}

func (a *OrderAggregate) CreateOrder(ctx context.Context, order *order.Order) error {
	return a.repo.CreateOrder(ctx, order)
}

func (a *OrderAggregate) GetOrder(ctx context.Context, orderID string) (*order.Order, error) {
	return a.repo.GetOrder(ctx, orderID)
}

func (a *OrderAggregate) UpdateOrderStatus(ctx context.Context, orderID string, status string) error {
	return a.repo.UpdateOrderStatus(ctx, orderID, status)
}