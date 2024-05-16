package dto

import (
	"github.com/GrandOichii/todoapp-golang/api/models"
	"github.com/GrandOichii/todoapp-golang/api/security"
)

type PostUser struct {
	Username string `json:"username" validate:"required,gte=4,lt=20"`
	Password string `json:"password" validate:"required,gte=8,lt=20"`
}

func (u PostUser) ToUser() (*models.User, error) {
	hash, err := security.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	return &models.User{
		Username:     u.Username,
		PasswordHash: hash,
	}, nil
}
