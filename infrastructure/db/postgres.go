package db

import (
	"gorm.io/gorm"
	"log"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=user dbname=dbname sslmode=disable password=password"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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