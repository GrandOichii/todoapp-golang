package services

import (
	"errors"
	"fmt"

	dto "github.com/GrandOichii/todoapp-golang/api/dto/user"
	repositories "github.com/GrandOichii/todoapp-golang/api/repositories/user"
	"github.com/GrandOichii/todoapp-golang/api/security"
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

	newUser, err := user.ToUser()
	if err != nil {
		return err
	}

	ser.repo.Save(newUser)

	return nil
}

func (ser UserServiceImpl) Login(user *dto.PostUser) (*dto.GetUser, error) {
	if err := ser.validate.Struct(user); err != nil {
		return nil, err
	}

	existing := ser.repo.FindByUsername(user.Username)
	if existing == nil {
		return nil, errors.New("incorrect username or password")
	}

	if !security.CheckPasswordHash(user.Password, existing.PasswordHash) {
		return nil, errors.New("incorrect username or password")
	}

	return dto.CreateGetUser(existing), nil
}
