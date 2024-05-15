package repositories

import (
	"strconv"

	"github.com/GrandOichii/todoapp-golang/api/models"
)

type TaskRepositoryImpl struct {
	TaskRepository

	tasks []*models.Task
}

func (repo TaskRepositoryImpl) FindAll() []*models.Task {
	return repo.tasks
}

func (repo *TaskRepositoryImpl) Save(task *models.Task) error {
	task.Id = strconv.Itoa(len(repo.tasks))
	repo.tasks = append(repo.tasks, task)

	return nil
}

func (repo TaskRepositoryImpl) FindById(id string) *models.Task {
	for _, task := range repo.tasks {
		if task.Id == id {
			return task
		}
	}
	return nil
}
