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

type FakeTaskRepository struct {
	tasks []*models.Task
}

func CreateMockTaskRepository() *FakeTaskRepository {
	return &FakeTaskRepository{}
}

func (repo FakeTaskRepository) FindAll() []*models.Task {
	return repo.tasks
}

func (repo FakeTaskRepository) FindByOwnerId(ownerId string) []*models.Task {
	return filter(repo.tasks, func(task *models.Task) bool {
		return task.OwnerId == ownerId
	})
}

func (repo *FakeTaskRepository) Save(task *models.Task) error {
	task.Id = strconv.Itoa(len(repo.tasks))
	repo.tasks = append(repo.tasks, task)

	return nil
}

func (repo FakeTaskRepository) FindById(id string) *models.Task {
	for _, task := range repo.tasks {
		if task.Id == id {
			return task
		}
	}
	return nil
}
