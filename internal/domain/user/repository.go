package user

import (
	"context"
	"shop-app/internal/dao"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user *dao.User) error
	LoginUser(ctx context.Context, username, password string) (*dao.User, error)
	UpdateUser(ctx context.Context, user *dao.User) error
	DeleteUser(ctx context.Context, userID int64) error
}

type userRepository struct {
	db *dao.DB
}

func NewUserRepository(db *dao.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) RegisterUser(ctx context.Context, user *dao.User) error {
	return dao.RegisterUser(r.db, user)
}

func (r *userRepository) LoginUser(ctx context.Context, username, password string) (*dao.User, error) {
	return dao.LoginUser(r.db, username, password)
}

func (r *userRepository) UpdateUser(ctx context.Context, user *dao.User) error {
	return dao.UpdateUser(r.db, user)
}

func (r *userRepository) DeleteUser(ctx context.Context, userID int64) error {
	return dao.DeleteUser(r.db, userID)
}