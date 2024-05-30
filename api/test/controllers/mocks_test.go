// ! because of alphabetical file ordering, this file can't be called mocks.go - why on earth would is this even a thing
package controllers_test

import (
	taskdto "github.com/GrandOichii/todoapp-golang/api/dto/task"
	userdto "github.com/GrandOichii/todoapp-golang/api/dto/user"
	taskservices "github.com/GrandOichii/todoapp-golang/api/services/task"
	userservices "github.com/GrandOichii/todoapp-golang/api/services/user"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	userservices.UserService
	mock.Mock
}

func createMockUserService() *MockUserService {
	return new(MockUserService)
}

func (ser *MockUserService) Login(user *userdto.PostUser) (*userdto.GetUser, error) {
	args := ser.Called(user)
	switch user := args.Get(0).(type) {
	case *userdto.GetUser:
		return user, args.Error(1)
	case nil:
		return nil, args.Error(1)
	}
	return nil, args.Error(1)
}

func (ser *MockUserService) Register(user *userdto.PostUser) error {
	args := ser.Called(user)
	return args.Error(0)
}

type MockTaskService struct {
	taskservices.TaskService
	mock.Mock
}

func createMockTaskService() *MockTaskService {
	return new(MockTaskService)
}

func (ser *MockTaskService) GetAll(userId string) []*taskdto.GetTask {
	return ser.Called(userId).Get(0).([]*taskdto.GetTask)
}

func (ser *MockTaskService) Add(ownerId string, task *taskdto.CreateTask) (*taskdto.GetTask, error) {
	args := ser.Called(ownerId, task)
	switch task := args.Get(0).(type) {
	case *taskdto.GetTask:
		return task, args.Error(1)
	case nil:
		return nil, args.Error(1)
	}
	return nil, args.Error(1)

}

func (ser *MockTaskService) GetById(userId, id string) (*taskdto.GetTask, error) {
	args := ser.Called(userId, id)
	switch task := args.Get(0).(type) {
	case *taskdto.GetTask:
		return task, args.Error(1)
	case nil:
		return nil, args.Error(1)
	}
	return nil, args.Error(1)

}

func (ser *MockTaskService) ToggleCompleted(userId, id string) (*taskdto.GetTask, error) {
	args := ser.Called(userId, id)
	switch task := args.Get(0).(type) {
	case *taskdto.GetTask:
		return task, args.Error(1)
	case nil:
		return nil, args.Error(1)
	}
	return nil, args.Error(1)

}

func (ser *MockTaskService) Delete(userId, id string) error {
	return ser.Called(userId, id).Error(0)
}
