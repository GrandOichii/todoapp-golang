package dto

import "github.com/GrandOichii/todoapp-golang/api/models"

type CreateTask struct {
	Title string `json:"title" validate:"required,gt=3,lt=20"`
	Text  string `json:"text" validate:"lt=30"`
}

func (t CreateTask) ToTask() *models.Task {
	return &models.Task{
		Id:        "",
		Title:     t.Title,
		Text:      t.Text,
		Completed: false,
	}
}
