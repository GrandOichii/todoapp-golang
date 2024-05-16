package services_test

import (
	"testing"

	dto "github.com/GrandOichii/todoapp-golang/api/dto/user"
	"github.com/GrandOichii/todoapp-golang/api/models"
	repositories "github.com/GrandOichii/todoapp-golang/api/repositories/user"
	"github.com/GrandOichii/todoapp-golang/api/security"
	services "github.com/GrandOichii/todoapp-golang/api/services/user"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	repositories.UserRepository
	mock.Mock
}

func (m *MockUserRepository) FindByUsername(username string) *models.User {
	args := m.Called(username)
	switch user := args.Get(0).(type) {
	case *models.User:
		return user
	case nil:
		return nil
	}
	return nil
}

func (m *MockUserRepository) Save(user *models.User) *models.User {
	args := m.Called(user)
	switch user := args.Get(0).(type) {
	case *models.User:
		return user
	case nil:
		return nil
	}
	return nil
}

func createUserRepository() *MockUserRepository {
	return new(MockUserRepository)
}

func createUserService(repo *MockUserRepository) services.UserService {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return services.CreateUserServiceImpl(
		repo,
		validate,
	)
}

func Test_ShouldRegister(t *testing.T) {
	repo := createUserRepository()
	service := createUserService(repo)
	data := dto.PostUser{
		Username: "user",
		Password: "password",
	}

	repo.On("Save", mock.Anything).Return(&models.User{})
	repo.On("FindByUsername", data.Username).Return(nil)

	err := service.Register(&data)
	assert.Nil(t, err)
}

func Test_ShouldNotRegisterUsernameTaken(t *testing.T) {
	repo := createUserRepository()
	service := createUserService(repo)
	data := dto.PostUser{
		Username: "user",
		Password: "password",
	}

	// repo.On("Save", mock.Anything).Return(&models.User{})
	repo.On("FindByUsername", data.Username).Return(&models.User{})

	err := service.Register(&data)
	assert.NotNil(t, err)
}

func Test_ShouldNotLogin(t *testing.T) {
	repo := createUserRepository()
	service := createUserService(repo)
	data := dto.PostUser{
		Username: "user",
		Password: "password",
	}

	// repo.On("Save", mock.Anything).Return(&models.User{})
	repo.On("FindByUsername", data.Username).Return(nil)

	login, err := service.Login(&data)
	assert.Nil(t, login)
	assert.NotNil(t, err)
}

func Test_ShouldNotLoginIncorrectPassword(t *testing.T) {
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

	login, err := service.Login(&data)
	assert.Nil(t, login)
	assert.NotNil(t, err)
}

func Test_ShouldLogin(t *testing.T) {
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

	login, err := service.Login(&data)
	assert.NotNil(t, login)
	assert.Nil(t, err)
}
