package logistics

import (
	"context"
	"shop-app/internal/dao"
)

type LogisticsRepository interface {
	CreateLogistics(ctx context.Context, logisticsOrder *dao.LogisticsOrder) error
	UpdateLogisticsStatus(ctx context.Context, logisticsID string, status string, description string) error
	GetLogisticsDetails(ctx context.Context, logisticsID string) (*dao.LogisticsOrder, error)
}

type logisticsRepository struct {
	db *dao.DB
}

func NewLogisticsRepository(db *dao.DB) *logisticsRepository {
	return &logisticsRepository{db: db}
}

func (r *logisticsRepository) CreateLogistics(ctx context.Context, logisticsOrder *dao.LogisticsOrder) error {
	return dao.CreateLogistics(r.db, logisticsOrder)
}

func (r *logisticsRepository) UpdateLogisticsStatus(ctx context.Context, logisticsID string, status string, description string) error {
	return dao.UpdateLogisticsStatus(r.db, logisticsID, status, description)
}

func (r *logisticsRepository) GetLogisticsDetails(ctx context.Context, logisticsID string) (*dao.LogisticsOrder, error) {
	return dao.GetLogisticsDetails(r.db, logisticsID)
}