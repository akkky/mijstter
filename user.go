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
	err := c.BindJSON(&user)
	if err != nil {
		return NewError(http.StatusBadRequest, "Bad Request", &err)
	}

	if !user.CheckPassword() {
		return NewError(http.StatusBadRequest, "Password is not match.", nil)
	}

	user.SetPasswordHash()

	fmt.Println(database)

	err = database.WriteUser(&user)
	if err != nil {
		return NewError(http.StatusInternalServerError, "User can not be written.", &err)
	}

	c.JSON(http.StatusCreated, user.Id)

	return nil
}
