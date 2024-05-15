package models

type Task struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
