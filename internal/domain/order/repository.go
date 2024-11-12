package order

import (
	"context"
	"shop-app/internal/dao"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *dao.Order) error
	GetOrder(ctx context.Context, orderID string) (*dao.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID string, status string) error
}

type orderRepository struct {
	db *dao.DB
}

func NewOrderRepository(db *dao.DB) *orderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(ctx context.Context, order *dao.Order) error {
	return dao.CreateOrder(r.db, order)
}

func (r *orderRepository) GetOrder(ctx context.Context, orderID string) (*dao.Order, error) {
	return dao.GetOrder(r.db, orderID)
}

func (r *orderRepository) UpdateOrderStatus(ctx context.Context, orderID string, status string) error {
	return dao.UpdateOrderStatus(r.db, orderID, status)
}