package dto

import "github.com/GrandOichii/todoapp-golang/api/models"

type CreateTask struct {
	Title string `json:"title" validate:"required,gt=3,lt=30"`
	Text  string `json:"text" validate:"required"`
}

func (t CreateTask) ToTask() *models.Task {
	return &models.Task{
		Title: t.Title,
		Text:  t.Text,
	}
}
