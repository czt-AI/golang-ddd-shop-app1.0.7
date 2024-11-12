package user

import (
	"context"
	"shop-app/domain/user"
)

type UserAggregate struct {
	repo user.UserRepository
}

func NewUserAggregate(repo user.UserRepository) *UserAggregate {
	return &UserAggregate{repo: repo}
}

func (a *UserAggregate) RegisterUser(ctx context.Context, user *user.User) error {
	return a.repo.RegisterUser(ctx, user)
}

func (a *UserAggregate) LoginUser(ctx context.Context, username, password string) (*user.User, error) {
	return a.repo.LoginUser(ctx, username, password)
}

func (a *UserAggregate) UpdateUser(ctx context.Context, user *user.User) error {
	return a.repo.UpdateUser(ctx, user)
}

func (a *UserAggregate) DeleteUser(ctx context.Context, userID int64) error {
	return a.repo.DeleteUser(ctx, userID)
}