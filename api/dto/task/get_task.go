package dto

import "github.com/GrandOichii/todoapp-golang/api/models"

type GetTask struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

func CreateGetTask(task *models.Task) *GetTask {
	return &GetTask{
		Id:        task.Id,
		Title:     task.Title,
		Text:      task.Text,
		Completed: task.Completed,
	}
}
