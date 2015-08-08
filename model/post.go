package model

type Post struct {
	Model
	SessionId string `form:"session_id" json:"session_id,omitempty"`
	UserId    ID     `form:"user_id" json:"user_id"`
	UserName  string `form:"user_name" json:"user_name"`
	Message   string `form:"message" json:"message"`
	Url       string `form:"url" json:"url"`
}
