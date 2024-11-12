package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db

	// 自动迁移模式
	db.AutoMigrate(
		&user.User{},
		&order.Order{},
		&payment.Payment{},
		&logistics.LogisticsOrder{},
		&comment.Comment{},
	)
}