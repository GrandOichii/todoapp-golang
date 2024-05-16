package dto

type PostUser struct {
	Username string `json:"username" validate:"required,gt=4,lt=20"`
	Password string `json:"password" validate:"required,gt=8,lt=20"`
}
