package user

import (
	"context"
	"shop-app/api/v1/dto"
	"shop-app/domain/user"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, req *dto.RegisterUserRequest) error
	LoginUser(ctx context.Context, req *dto.LoginUserRequest) (*dto.UserResponse, error)
}

type userUsecase struct {
	repo user.UserRepository
}

func NewUserUsecase(repo user.UserRepository) *userUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) RegisterUser(ctx context.Context, req *dto.RegisterUserRequest) error {
	return u.repo.RegisterUser(ctx, &req.User)
}

func (u *userUsecase) LoginUser(ctx context.Context, req *dto.LoginUserRequest) (*dto.UserResponse, error) {
	user, err := u.repo.LoginUser(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return dto.ConvertToUserResponse(user), nil
}