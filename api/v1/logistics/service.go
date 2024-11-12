package v1

import (
	"context"
	"shop-app/api/v1/dto"
	"shop-app/domain/logistics"
)

type LogisticsService interface {
	CreateLogistics(ctx context.Context, req *dto.CreateLogisticsRequest) error
	UpdateLogisticsStatus(ctx context.Context, req *dto.UpdateLogisticsStatusRequest) error
	GetLogisticsDetails(ctx context.Context, logisticsID string) (*dto.LogisticsOrderResponse, error)
}

type logisticsService struct {
	usecase logistics.LogisticsUsecase
}

func NewLogisticsService(usecase logistics.LogisticsUsecase) *logisticsService {
	return &logisticsService{usecase: usecase}
}

func (s *logisticsService) CreateLogistics(ctx context.Context, req *dto.CreateLogisticsRequest) error {
	return s.usecase.CreateLogistics(ctx, req)
}

func (s *logisticsService) UpdateLogisticsStatus(ctx context.Context, req *dto.UpdateLogisticsStatusRequest) error {
	return s.usecase.UpdateLogisticsStatus(ctx, req)
}

func (s *logisticsService) GetLogisticsDetails(ctx context.Context, logisticsID string) (*dto.LogisticsOrderResponse, error) {
	return s.usecase.GetLogisticsDetails(ctx, logisticsID)
}