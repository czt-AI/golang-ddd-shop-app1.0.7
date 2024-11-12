package logistics

import (
	"context"
	"shop-app/domain/logistics"
)

type LogisticsService interface {
	CreateLogistics(ctx context.Context, logisticsOrder *logistics.LogisticsOrder) error
	UpdateLogisticsStatus(ctx context.Context, logisticsID string, status string, description string) error
	GetLogisticsDetails(ctx context.Context, logisticsID string) (*logistics.LogisticsOrder, error)
}

type logisticsService struct {
	repo logistics.LogisticsRepository
}

func NewLogisticsService(repo logistics.LogisticsRepository) *logisticsService {
	return &logisticsService{repo: repo}
}

func (s *logisticsService) CreateLogistics(ctx context.Context, logisticsOrder *logistics.LogisticsOrder) error {
	// 创建物流订单前，检查订单状态
	order, err := s.repo.GetOrder(ctx, logisticsOrder.OrderID)
	if err != nil {
		return err
	}
	if order.Status != "shipped" {
		return errors.New("order status is not shipped")
	}
	return s.repo.CreateLogistics(ctx, logisticsOrder)
}

func (s *logisticsService) UpdateLogisticsStatus(ctx context.Context, logisticsID string, status string, description string) error {
	// 更新物流状态前，检查物流状态是否可以更新
	logistics, err := s.repo.GetLogistics(ctx, logisticsID)
	if err != nil {
		return err
	}
	if !canUpdateStatus(logistics.Status, status) {
		return errors.New("logistics status cannot be updated")
	}
	return s.repo.UpdateLogisticsStatus(ctx, logisticsID, status, description)
}

func (s *logisticsService) GetLogisticsDetails(ctx context.Context, logisticsID string) (*logistics.LogisticsOrder, error) {
	return s.repo.GetLogisticsDetails(ctx, logisticsID)
}

// canUpdateStatus 检查物流状态是否可以更新
func canUpdateStatus(currentStatus, newStatus string) bool {
	// 这里定义物流状态更新规则
	// ...
	return canUpdate
}