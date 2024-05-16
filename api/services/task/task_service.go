package services

import dto "github.com/GrandOichii/todoapp-golang/api/dto/task"

type TaskService interface {
	GetAll(string) []*dto.GetTask
	Add(string, *dto.CreateTask) (*dto.GetTask, error)
	GetById(userId, id string) (*dto.GetTask, error)
	ToggleCompleted(userId, id string) (*dto.GetTask, error)
	Delete(userId, id string) error
}
