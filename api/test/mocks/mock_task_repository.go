package mocks

import (
	"strconv"

	"github.com/GrandOichii/todoapp-golang/api/models"
)

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

type MockTaskRepository struct {
	tasks []*models.Task
}

func CreateMockTaskRepository() *MockTaskRepository {
	return &MockTaskRepository{}
}

func (repo MockTaskRepository) FindAll() []*models.Task {
	return repo.tasks
}

func (repo MockTaskRepository) FindByOwnerId(ownerId string) []*models.Task {
	return filter(repo.tasks, func(task *models.Task) bool {
		return task.OwnerId == ownerId
	})
}

func (repo *MockTaskRepository) Save(task *models.Task) error {
	task.Id = strconv.Itoa(len(repo.tasks))
	repo.tasks = append(repo.tasks, task)

	return nil
}

func (repo MockTaskRepository) FindById(id string) *models.Task {
	for _, task := range repo.tasks {
		if task.Id == id {
			return task
		}
	}
	return nil
}
