package logistics

import (
	"context"
	"shop-app/api/v1/dto"
	"shop-app/domain/logistics"
)

type LogisticsUsecase interface {
	CreateLogistics(ctx context.Context, req *dto.CreateLogisticsRequest) error
	UpdateLogisticsStatus(ctx context.Context, req *dto.UpdateLogisticsStatusRequest) error
	GetLogisticsDetails(ctx context.Context, logisticsID string) (*dto.LogisticsOrderResponse, error)
}

type logisticsUsecase struct {
	repo logistics.LogisticsRepository
}

func NewLogisticsUsecase(repo logistics.LogisticsRepository) *logisticsUsecase {
	return &logisticsUsecase{repo: repo}
}

func (u *logisticsUsecase) CreateLogistics(ctx context.Context, req *dto.CreateLogisticsRequest) error {
	return u.repo.CreateLogistics(ctx, &req.LogisticsOrder)
}

func (u *logisticsUsecase) UpdateLogisticsStatus(ctx context.Context, req *dto.UpdateLogisticsStatusRequest) error {
	return u.repo.UpdateLogisticsStatus(ctx, req.LogisticsID, req.Status, req.Description)
}

func (u *logisticsUsecase) GetLogisticsDetails(ctx context.Context, logisticsID string) (*dto.LogisticsOrderResponse, error) {
	logisticsOrder, err := u.repo.GetLogisticsDetails(ctx, logisticsID)
	if err != nil {
		return nil, err
	}
	return dto.ConvertToLogisticsOrderResponse(logisticsOrder), nil
}