package controllers_test

import (
	"errors"
	"testing"

	"github.com/GrandOichii/todoapp-golang/api/controllers"
	dto "github.com/GrandOichii/todoapp-golang/api/dto/user"
	"github.com/GrandOichii/todoapp-golang/api/middleware"
	services "github.com/GrandOichii/todoapp-golang/api/services/user"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func createAuthController(service services.UserService) *controllers.AuthController {
	return controllers.CreateAuthController(
		service,
		middleware.CreateJwtMiddleware(service).Middle.LoginHandler,
	)
}

func Test_ShouldRegister(t *testing.T) {
	// arrange
	service := createMockUserService()
	controller := createAuthController(service)
	data := dto.PostUser{
		Username: "user",
		Password: "password",
	}
	service.On("Register", mock.Anything).Return(nil)

	c, w := createTestContext(data)

	// act
	controller.Register(c)

	// assert
	assert.Equal(t, w.Code, 200)
}

func Test_ShouldNotRegister(t *testing.T) {
	// arrange
	service := createMockUserService()
	controller := createAuthController(service)
	data := dto.PostUser{
		Username: "user",
		Password: "password",
	}
	service.On("Register", mock.Anything).Return(errors.New(""))

	c, w := createTestContext(data)

	// act
	controller.Register(c)

	// assert
	assert.Equal(t, w.Code, 400)
}

func Test_ShouldLogin(t *testing.T) {
	// arrange
	service := createMockUserService()
	controller := createAuthController(service)
	data := dto.PostUser{
		Username: "user",
		Password: "password",
	}
	service.On("Login", mock.Anything).Return(&dto.GetUser{
		Id:       "userId",
		Username: "user",
	}, nil)

	c, w := createTestContext(data)

	// act
	controller.Login(c)

	// assert
	assert.Equal(t, w.Code, 200)
}

func Test_ShouldNotLogin(t *testing.T) {
	// arrange
	service := createMockUserService()
	controller := createAuthController(service)
	data := dto.PostUser{
		Username: "user",
		Password: "password",
	}
	service.On("Login", mock.Anything).Return(nil, errors.New(""))

	c, w := createTestContext(data)

	// act
	controller.Login(c)

	// assert
	assert.Equal(t, w.Code, 401)
}
