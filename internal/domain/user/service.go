package user

import (
	"context"
	"shop-app/domain/user"
)

type UserService interface {
	RegisterUser(ctx context.Context, user *user.User) error
	LoginUser(ctx context.Context, username, password string) (*user.User, error)
}

type userService struct {
	repo user.UserRepository
}

func NewUserService(repo user.UserRepository) *userService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(ctx context.Context, user *user.User) error {
	// 使用密码哈希函数加密密码
	user.Password = encryptPassword(user.Password)
	return s.repo.RegisterUser(ctx, user)
}

func (s *userService) LoginUser(ctx context.Context, username, password string) (*user.User, error) {
	user, err := s.repo.LoginUser(ctx, username, password)
	if err != nil {
		return nil, err
	}
	// 检查密码是否匹配
	if !checkPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

// encryptPassword 使用哈希函数加密密码
func encryptPassword(password string) string {
	// 这里使用bcrypt算法进行密码加密
	// ...
	return encryptedPassword
}

// checkPasswordHash 检查密码是否与哈希值匹配
func checkPasswordHash(password, hash string) bool {
	// 这里使用bcrypt算法进行密码验证
	// ...
	return match
}