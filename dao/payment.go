package dao

import (
	"shop-app/domain/payment"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	OrderID       string  `gorm:"index"`
	PaymentMethod string  `gorm:"index"`
	Amount        float64 `gorm:"index"`
	Status        string  `gorm:"index"`
	CreatedAt     string  `gorm.io/gorm:"index"`
	UpdatedAt     string  `gorm.io/gorm:"index"`
}

func (p *Payment) Record(db *gorm.DB) error {
	return db.Create(p).Error
}

func (p *Payment) Confirm(db *gorm.DB) error {
	return db.Model(p).Update("status", "succeeded").Error
}

func (p *Payment) Refund(db *gorm.DB) error {
	return db.Model(p).Update("status", "refunded").Error
}