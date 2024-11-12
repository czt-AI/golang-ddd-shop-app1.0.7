package v1

import (
	"context"
	"shop-app/api/v1/dto"
	"shop-app/domain/user"
)

type UserService interface {
	RegisterUser(ctx context.Context, req *dto.RegisterUserRequest) error
	LoginUser(ctx context.Context, req *dto.LoginUserRequest) (*dto.UserResponse, error)
}

type userService struct {
	usecase user.UserUsecase
}

func NewUserService(usecase user.UserUsecase) *userService {
	return &userService{usecase: usecase}
}

func (s *userService) RegisterUser(ctx context.Context, req *dto.RegisterUserRequest) error {
	return s.usecase.RegisterUser(ctx, req)
}

func (s *userService) LoginUser(ctx context.Context, req *dto.LoginUserRequest) (*dto.UserResponse, error) {
	return s.usecase.LoginUser(ctx, req)
}