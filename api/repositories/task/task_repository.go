package repositories

import "github.com/GrandOichii/todoapp-golang/api/models"

type TaskRepository interface {
	FindAll() []*models.Task
	Save(*models.Task) error
	FindById(id string) *models.Task
	UpdateById(id string, updateF func(*models.Task) *models.Task) *models.Task
}
