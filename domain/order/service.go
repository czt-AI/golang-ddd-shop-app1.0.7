package order

import (
	"context"
	"shop-app/domain/order"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *order.Order) error
	GetOrder(ctx context.Context, orderID string) (*order.Order, error)
}

type orderService struct {
	repo order.OrderRepository
	rmq  *messagequeue.RabbitMQClient
}

func NewOrderService(repo order.OrderRepository, rmq *messagequeue.RabbitMQClient) *orderService {
	return &orderService{
		repo: repo,
		rmq:  rmq,
	}
}

func (s *orderService) CreateOrder(ctx context.Context, order *order.Order) error {
	// Create order in the database
	if err := s.repo.CreateOrder(ctx, order); err != nil {
		return err
	}

	// Publish a message to a queue to handle further processing
	if err := s.rmq.Publish("order", "order_created", []byte(order.OrderID)); err != nil {
		return err
	}

	return nil
}

func (s *orderService) GetOrder(ctx context.Context, orderID string) (*order.Order, error) {
	return s.repo.GetOrder(ctx, orderID)
}