package dto

import "github.com/GrandOichii/todoapp-golang/api/models"

type PostUser struct {
	Username string `json:"username" validate:"required,gt=4,lt=20"`
	Password string `json:"password" validate:"required,gte=8,lt=20"`
}

func (u PostUser) ToUser() *models.User {
	// TODO hash password
	return &models.User{
		Username:     u.Username,
		PasswordHash: u.Password,
	}
}
