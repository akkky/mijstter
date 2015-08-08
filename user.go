package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mijstter/model"
	"net/http"
	"strconv"
)

const defaultUserLimit = 20

func NewUser(c *gin.Context) {
	err := newUser(c)
	if err != nil {
		c.JSON(err.Status, err)
	}
}

func newUser(c *gin.Context) *Error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return NewError(http.StatusBadRequest, "Bad Request", &err)
	}

	log.Printf("User : %v\n", user)

	if !user.IsValidUserName() {
		return NewError(http.StatusBadRequest, "user_name is not valid.", nil)
	}

	// ユーザーの存在チェック
	exist, err := database.IsUserNameExist(user.UserName)
	if err != nil {
		return NewError(http.StatusInternalServerError, "Can not read users.", &err)
	}
	if exist {
		return NewError(http.StatusConflict, "user_name is already exists.", nil)
	}

	if !user.CheckPassword() {
		return NewError(http.StatusBadRequest, "Password is not match.", nil)
	}

	user.SetPasswordHash()

	// 他の書き込みを待機
	writeSemaphore <- 0
	defer func() { <-writeSemaphore }()

	err = database.WriteUser(&user)
	if err != nil {
		return NewError(http.StatusInternalServerError, "User can not be written.", &err)
	}

	c.JSON(http.StatusCreated, user.Id)

	return nil
}

func GetUsers(c *gin.Context) {
	err := getUsers(c)
	if err != nil {
		c.JSON(err.Status, err)
	}
}

func getUsers(c *gin.Context) *Error {
	limitStr := c.Query("limit")

	var limit int
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = defaultUserLimit
	}

	users, err := database.ReadUsers(limit)
	if err != nil {
		return NewError(http.StatusInternalServerError, "Can not read userslist.", &err)
	}

	c.JSON(http.StatusOK, users)

	return nil
}
