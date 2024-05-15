package services

import (
	"fmt"

	dto "github.com/GrandOichii/todoapp-golang/api/dto/task"
	"github.com/GrandOichii/todoapp-golang/api/models"
	repositories "github.com/GrandOichii/todoapp-golang/api/repositories/task"
	"github.com/GrandOichii/todoapp-golang/api/utility"
	"github.com/go-playground/validator/v10"
)

type TaskServiceImpl struct {
	TaskService

	repo     repositories.TaskRepository
	validate *validator.Validate
}

func CreateTaskServiceImpl(repo repositories.TaskRepository, validate *validator.Validate) *TaskServiceImpl {
	return &TaskServiceImpl{
		repo:     repo,
		validate: validate,
	}
}

func (ser TaskServiceImpl) GetAll() []*dto.GetTask {
	return utility.MapSlice(
		ser.repo.FindAll(),
		func(task *models.Task) *dto.GetTask {
			return dto.CreateGetTask(task)
		},
	)
}

func (ser TaskServiceImpl) Add(task *dto.CreateTask) (*dto.GetTask, error) {
	if err := ser.validate.Struct(task); err != nil {
		return nil, err
	}

	newTask := task.ToTask()

	if err := ser.repo.Save(newTask); err != nil {
		return nil, err
	}

	return dto.CreateGetTask(newTask), nil
}

func (ser TaskServiceImpl) GetById(id string) (*dto.GetTask, error) {
	result := ser.repo.FindById(id)
	if result == nil {
		return nil, fmt.Errorf("no task with id %s", id)
	}
	return dto.CreateGetTask(result), nil
}

func (ser TaskServiceImpl) ToggleCompleted(id string) (*dto.GetTask, error) {
	result := ser.repo.UpdateById(id, func(task *models.Task) *models.Task {
		task.Completed = !task.Completed
		return task
	})
	if result == nil {
		return nil, fmt.Errorf("no task with id %s", id)
	}
	return dto.CreateGetTask(result), nil
}
