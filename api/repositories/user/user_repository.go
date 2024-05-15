package repositories

import "github.com/GrandOichii/todoapp-golang/api/models"

type UserRepository interface {
	FindByUsername(username string) *models.User
	Save(*models.User) bool
}
