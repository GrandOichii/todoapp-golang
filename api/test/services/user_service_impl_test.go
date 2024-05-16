package services_test

import (
	"testing"

	dto "github.com/GrandOichii/todoapp-golang/api/dto/user"
	"github.com/GrandOichii/todoapp-golang/api/models"
	"github.com/GrandOichii/todoapp-golang/api/security"
	services "github.com/GrandOichii/todoapp-golang/api/services/user"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createUserService(repo *MockUserRepository) services.UserService {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return services.CreateUserServiceImpl(
		repo,
		validate,
	)
}

func Test_ShouldRegister(t *testing.T) {
	// arrange
	repo := createUserRepository()
	service := createUserService(repo)
	data := dto.PostUser{
		Username: "user",
		Password: "password",
	}

	repo.On("Save", mock.Anything).Return(&models.User{})
	repo.On("FindByUsername", data.Username).Return(nil)

	// act
	err := service.Register(&data)

	// assert
	assert.Nil(t, err)
}

func Test_ShouldNotRegisterUsernameTaken(t *testing.T) {
	// arrange
	repo := createUserRepository()
	service := createUserService(repo)
	data := dto.PostUser{
		Username: "user",
		Password: "password",
	}

	repo.On("FindByUsername", data.Username).Return(&models.User{})

	// act
	err := service.Register(&data)

	// assert
	assert.NotNil(t, err)
}

func Test_ShouldNotLogin(t *testing.T) {
	// arrange
	repo := createUserRepository()
	service := createUserService(repo)
	data := dto.PostUser{
		Username: "user",
		Password: "password",
	}

	repo.On("FindByUsername", data.Username).Return(nil)

	// act
	login, err := service.Login(&data)

	// assert
	assert.Nil(t, login)
	assert.NotNil(t, err)
}

func Test_ShouldNotLoginIncorrectPassword(t *testing.T) {
	// arrange
	repo := createUserRepository()
	service := createUserService(repo)
	data := dto.PostUser{
		Username: "user",
		Password: "password",
	}
	existing := models.User{
		Username:     data.Username,
		PasswordHash: "passwordHash",
	}

	repo.On("FindByUsername", data.Username).Return(&existing)

	// act
	login, err := service.Login(&data)

	// assert
	assert.Nil(t, login)
	assert.NotNil(t, err)
}

func Test_ShouldLogin(t *testing.T) {
	// arrange
	repo := createUserRepository()
	service := createUserService(repo)
	data := dto.PostUser{
		Username: "user",
		Password: "password",
	}
	hash, _ := security.HashPassword(data.Password)
	existing := models.User{
		Username:     data.Username,
		PasswordHash: hash,
	}

	repo.On("FindByUsername", data.Username).Return(&existing)

	// act
	login, err := service.Login(&data)

	// assert
	assert.NotNil(t, login)
	assert.Nil(t, err)
}
