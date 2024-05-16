package services

import dto "github.com/GrandOichii/todoapp-golang/api/dto/user"

type LoginResult struct {
	Token string `json:"token"`
}

type UserService interface {
	Login(*dto.PostUser) (*LoginResult, error)
	Register(*dto.PostUser) error
}
