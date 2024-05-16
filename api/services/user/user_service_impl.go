package services

import (
	dto "github.com/GrandOichii/todoapp-golang/api/dto/user"
	repositories "github.com/GrandOichii/todoapp-golang/api/repositories/user"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserService

	validate *validator.Validate
	repo     repositories.UserRepository
}

func CreateUserServiceImpl(repo repositories.UserRepository, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		repo:     repo,
		validate: validate,
	}
}

func (ser UserServiceImpl) Register(user *dto.PostUser) error {
	// TODO
	return nil
}

func (ser UserServiceImpl) Login(user *dto.PostUser) (*LoginResult, error) {
	// TODO
	return nil, nil
}
