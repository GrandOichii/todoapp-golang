package services

import (
	"errors"
	"fmt"

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
	if err := ser.validate.Struct(user); err != nil {
		return err
	}

	existing := ser.repo.FindByUsername(user.Username)
	if existing != nil {
		return fmt.Errorf("user with username %s already exists", user.Username)
	}

	newUser := user.ToUser()
	ser.repo.Save(newUser)

	return nil
}

func (ser UserServiceImpl) Login(user *dto.PostUser) (*LoginResult, error) {
	if err := ser.validate.Struct(user); err != nil {
		return nil, err
	}

	existing := ser.repo.FindByUsername(user.Username)
	if existing == nil {
		return nil, errors.New("incorrect username or password")
	}

	// TODO check hash
	if existing.PasswordHash != user.Password {
		return nil, errors.New("incorrect username or password")
	}

	// TODO generate token
	token := "jwt token"

	return &LoginResult{
		Token: token,
	}, nil
}
