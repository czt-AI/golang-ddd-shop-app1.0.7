package dao

import (
	"shop-app/domain/order"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID      int64   `gorm:"index"`
	ProductList []OrderItem `gorm:"many2many:order_items;"`
	TotalAmount float64 `gorm:"index"`
	Status      string  `gorm:"index"`
	CreatedAt   string  `gorm:"index"`
	UpdatedAt   string  `gorm:"index"`
}

type OrderItem struct {
	gorm.Model
	OrderID      int64   `gorm:"index"`
	ProductID    int64   `gorm:"index"`
	Quantity     int    `gorm:"index"`
	Price        float64 `gorm:"index"`
}

func (o *Order) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

func (o *Order) Get(db *gorm.DB, orderID string) (*Order, error) {
	var order Order
	err := db.Where("id = ?", orderID).First(&order).Error
	return &order, err
}

func (o *Order) Update(db *gorm.DB) error {
	return db.Save(o).Error
}

func (o *Order) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}