package logistics

import (
	"context"
	"shop-app/domain/logistics"
)

type LogisticsAggregate struct {
	repo logistics.LogisticsRepository
}

func NewLogisticsAggregate(repo logistics.LogisticsRepository) *LogisticsAggregate {
	return &LogisticsAggregate{repo: repo}
}

func (a *LogisticsAggregate) CreateLogistics(ctx context.Context, logisticsOrder *logistics.LogisticsOrder) error {
	return a.repo.CreateLogistics(ctx, logisticsOrder)
}

func (a *LogisticsAggregate) UpdateLogisticsStatus(ctx context.Context, logisticsID string, status string, description string) error {
	return a.repo.UpdateLogisticsStatus(ctx, logisticsID, status, description)
}

func (a *LogisticsAggregate) GetLogisticsDetails(ctx context.Context, logisticsID string) (*logistics.LogisticsOrder, error) {
	return a.repo.GetLogisticsDetails(ctx, logisticsID)
}