package models

type User struct {
	Id           string `json:"id" bson:"_id,omitempty"`
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
}
