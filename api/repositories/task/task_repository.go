package repositories

import "github.com/GrandOichii/todoapp-golang/api/models"

type TaskRepository interface {
	FindAll() []*models.Task
	FindByOwnerId(ownerId string) []*models.Task
	Save(*models.Task) bool
	FindById(id string) *models.Task
	UpdateById(id string, updateF func(*models.Task) *models.Task) *models.Task
	Remove(id string) bool
}
