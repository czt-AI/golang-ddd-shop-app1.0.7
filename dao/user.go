package dao

import (
	"shop-app/domain/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(255);unique_index"`
	Password  string `gorm:"type:varchar(255)"`
	Email     string `gorm:"type:varchar(255);unique_index"`
	Phone     string `gorm:"type:varchar(20);unique_index"`
	Avatar    string `gorm:"type:varchar(255)"`
	Role      string `gorm:"type:varchar(50)"`
	CreateTime string `gorm:"type:datetime"`
	LastLogin string `gorm:"type:datetime"`
}

func (u *User) Register(db *gorm.DB) error {
	return db.Create(u).Error
}

func (u *User) Login(db *gorm.DB) (*User, error) {
	var user User
	err := db.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return &user, err
}

func (u *User) Update(db *gorm.DB) error {
	return db.Save(u).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Delete(u).Error
}