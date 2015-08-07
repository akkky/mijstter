package model

import (
	"time"
)

type Login struct {
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
}

type LoginResult struct {
	SessionId string `json:"session_id"`
}

type LoginSession struct {
	SessionId string
	User      *User
	Expire    time.Time
}
