package model

type Post struct {
	Model
	UserId   ID     `json:"user_id"`
	UserName string `json:"user_name"`
	Message  string `json:"message"`
	Url      string `json:"url"`
}
