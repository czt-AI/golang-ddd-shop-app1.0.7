package dao

import (
	"shop-app/domain/logistics"
	"gorm.io/gorm"
)

type LogisticsOrder struct {
	gorm.Model
	OrderID        string `gorm:"index"`
	LogisticsCompany string `gorm:"index"`
	ShippingTime   string `gorm:"index"`
	DeliveryTime   string `gorm:"index"`
	Status         string `gorm:"index"`
	Description    string `gorm:"index"`
	CreatedAt      string `gorm:"index"`
	UpdatedAt      string `gorm:"index"`
}

func (lo *LogisticsOrder) Create(db *gorm.DB) error {
	return db.Create(lo).Error
}

func (lo *LogisticsOrder) UpdateStatus(db *gorm.DB, status, description string) error {
	return db.Model(lo).Update("status", status).Error
}

func (lo *LogisticsOrder) GetDetails(db *gorm.DB, logisticsID string) (*LogisticsOrder, error) {
	var logisticsOrder LogisticsOrder
	err := db.Where("id = ?", logisticsID).First(&logisticsOrder).Error
	return &logisticsOrder, err
}