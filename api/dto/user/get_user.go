package dto

import "github.com/GrandOichii/todoapp-golang/api/models"

type GetUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func CreateGetUser(user *models.User) *GetUser {
	return &GetUser{
		Id:       user.Id,
		Username: user.Username,
	}
}
