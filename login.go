package main

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/gin-gonic/gin"
	"log"
	"mijstter/model"
	"net/http"
	"time"
)

var (
	sessionMap     = make(map[string]*model.LoginSession)
	sessionTimeout = time.Duration(30) * time.Minute
)

func CreateSessionId(user *model.User) string {
	session := model.LoginSession{
		SessionId: uuid.New(),
		User:      user,
		Expire:    time.Now().Add(sessionTimeout),
	}
	sessionMap[session.SessionId] = &session

	// セキュリティ対策のため3秒待つ
	time.Sleep(3 * time.Second)

	return session.SessionId
}

func GetUserBySessionId(sessionId string) *model.User {
	session, ok := sessionMap[sessionId]
	if !ok {
		return nil
	}

	if time.Now().After(session.Expire) {
		delete(sessionMap, sessionId)
		return nil
	}

	session.Expire = time.Now().Add(sessionTimeout)

	return session.User
}

func Login(c *gin.Context) {
	err := login(c)
	if err != nil {
		c.JSON(err.Status, err)
	}
}

func login(c *gin.Context) *Error {
	var login model.Login
	err := c.Bind(&login)
	if err != nil {
		return NewError(http.StatusBadRequest, "Bad Request", &err)
	}

	log.Printf("Login : %v\n", login)

	exist, err := database.IsUserNameExist(login.UserName)
	if err != nil {
		return NewError(http.StatusInternalServerError, "Can not read users.", &err)
	}
	if !exist {
		return NewError(http.StatusUnauthorized, "user_name or password is not valid.", nil)
	}

	user, err := database.ReadUserByUserName(login.UserName)
	if err != nil {
		return NewError(http.StatusInternalServerError, "Can not read users.", &err)
	}

	if !user.Authorize(login.Password) {
		return NewError(http.StatusUnauthorized, "user_name or password is not valid.", nil)
	}

	result := model.LoginResult{SessionId: CreateSessionId(user)}

	c.JSON(http.StatusOK, result)

	return nil
}
