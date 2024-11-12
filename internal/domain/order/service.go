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
}

func NewOrderService(repo order.OrderRepository) *orderService {
	return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(ctx context.Context, order *order.Order) error {
	// 在创建订单前，先检查库存是否充足
	if !checkStock(ctx, order.ProductList) {
		return errors.New("insufficient stock")
	}
	return s.repo.CreateOrder(ctx, order)
}

func (s *orderService) GetOrder(ctx context.Context, orderID string) (*order.Order, error) {
	return s.repo.GetOrder(ctx, orderID)
}

// checkStock 检查库存是否充足
func checkStock(ctx context.Context, productList []order.OrderItem) bool {
	// 这里遍历订单中的商品，检查库存
	// ...
	return hasStock
}