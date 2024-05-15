package models

type Task struct {
	Id        string `json:"id" bson:"_id,omitempty"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}
