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
	rmq  *messagequeue.RabbitMQClient
}

func NewLogisticsService(repo logistics.LogisticsRepository, rmq *messagequeue.RabbitMQClient) *logisticsService {
	return &logisticsService{
		repo: repo,
		rmq:  rmq,
	}
}

func (s *logisticsService) CreateLogistics(ctx context.Context, logisticsOrder *logistics.LogisticsOrder) error {
	// Create logistics order in the database
	if err := s.repo.CreateLogistics(ctx, logisticsOrder); err != nil {
		return err
	}

	// Publish a message to a queue to handle further processing
	if err := s.rmq.Publish("logistics", "logistics_created", []byte(logisticsOrder.OrderID)); err != nil {
		return err
	}

	return nil
}

func (s *logisticsService) UpdateLogisticsStatus(ctx context.Context, logisticsID string, status string, description string) error {
	// Update logistics status in the database
	if err := s.repo.UpdateLogisticsStatus(ctx, logisticsID, status, description); err != nil {
		return err
	}

	// Publish a message to a queue to handle further processing
	if err := s.rmq.Publish("logistics", "logistics_status_updated", []byte(logisticsID)); err != nil {
		return err
	}

	return nil
}

func (s *logisticsService) GetLogisticsDetails(ctx context.Context, logisticsID string) (*logistics.LogisticsOrder, error) {
	return s.repo.GetLogisticsDetails(ctx, logisticsID)
}