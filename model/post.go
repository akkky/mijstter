package model

type Post struct {
	Model
	UserId  ID
	Message string
	Url     string
}
