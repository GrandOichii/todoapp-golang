// ! because of alphabetical file ordering, this file can't be called mocks.go - why on earth would is this even a thing
package controllers_test

import (
	dto "github.com/GrandOichii/todoapp-golang/api/dto/user"
	services "github.com/GrandOichii/todoapp-golang/api/services/user"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	services.UserService
	mock.Mock
}

func createMockUserService() *MockUserService {
	return new(MockUserService)
}

func (ser *MockUserService) Login(user *dto.PostUser) (*dto.GetUser, error) {
	args := ser.Called(user)
	switch user := args.Get(0).(type) {
	case *dto.GetUser:
		return user, args.Error(1)
	case nil:
		return nil, args.Error(1)
	}
	return nil, args.Error(1)
}

func (ser *MockUserService) Register(user *dto.PostUser) error {
	args := ser.Called(user)
	return args.Error(0)
}
