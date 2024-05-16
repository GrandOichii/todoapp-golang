package services_test

import (
	"github.com/GrandOichii/todoapp-golang/api/models"
	taskrepositories "github.com/GrandOichii/todoapp-golang/api/repositories/task"
	userrepositories "github.com/GrandOichii/todoapp-golang/api/repositories/user"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	taskrepositories.TaskRepository
	mock.Mock
}

func createTaskRepository() *MockTaskRepository {
	return new(MockTaskRepository)
}

func (m *MockTaskRepository) FindByOwnerId(owner string) []*models.Task {
	return m.Called(owner).Get(0).([]*models.Task)
}

func (m *MockTaskRepository) Save(task *models.Task) bool {
	return m.Called(task).Bool(0)
}

func (m *MockTaskRepository) UpdateById(id string, updateF func(*models.Task) *models.Task) *models.Task {
	args := m.Called(id, updateF)
	switch task := args.Get(0).(type) {
	case *models.Task:
		return task
	case nil:
		return nil
	}
	return nil
}

func (m *MockTaskRepository) Remove(id string) bool {
	return m.Called(id).Bool(0)
}

func (m *MockTaskRepository) FindById(id string) *models.Task {
	args := m.Called(id)
	switch task := args.Get(0).(type) {
	case *models.Task:
		return task
	case nil:
		return nil
	}
	return nil

}

type MockUserRepository struct {
	userrepositories.UserRepository
	mock.Mock
}

func createUserRepository() *MockUserRepository {
	return new(MockUserRepository)
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
