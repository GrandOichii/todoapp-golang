package services

import dto "github.com/GrandOichii/todoapp-golang/api/dto/task"

type TaskService interface {
	GetAll() []*dto.GetTask
	Add(*dto.CreateTask) (*dto.GetTask, error)
	GetById(id string) (*dto.GetTask, error)
}
