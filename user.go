package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mijstter/model"
	"net/http"
)

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

	fmt.Printf("Post : %v\n", user)

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
	users, err := database.ReadUsers(30)
	if err != nil {
		return NewError(http.StatusInternalServerError, "Can not read userslist.", &err)
	}

	c.JSON(http.StatusOK, users)

	return nil
}
