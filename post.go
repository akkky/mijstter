package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mijstter/model"
	"net/http"
	"strconv"
)

const defaultPostLimit = 20

func GetPosts(c *gin.Context) {
	err := getPosts(c)
	if err != nil {
		c.JSON(err.Status, err)
	}
}

func getPosts(c *gin.Context) *Error {
	limitStr := c.Query("limit")

	var limit int
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = defaultPostLimit
	}

	posts, err := database.ReadPosts(limit)
	if err != nil {
		return NewError(http.StatusInternalServerError, "Can not read posts.", &err)
	}

	c.JSON(http.StatusOK, posts)

	return nil
}

func NewPost(c *gin.Context) {
	err := newPost(c)
	if err != nil {
		c.JSON(err.Status, err)
	}
}

func newPost(c *gin.Context) *Error {
	var post model.Post
	err := c.Bind(&post)
	if err != nil {
		return NewError(http.StatusBadRequest, "Bad Request", &err)
	}

	log.Printf("Post : %v\n", post)

	var user *model.User
	if post.SessionId != "" {
		user = GetUserBySessionId(post.SessionId)
	}
	if user == nil {
		user = &model.User{UserName: post.UserName}
	}
	if !user.IsValidUserName() {
		return NewError(http.StatusBadRequest, "user_name is not valid.", nil)
	}

	post.UserId = user.Id
	post.UserName = user.UserName

	log.Printf("Post (got user) : %v\n", post)

	// 他の書き込みを待機
	writeSemaphore <- 0
	defer func() { <-writeSemaphore }()

	err = database.WritePost(&post)
	if err != nil {
		return NewError(http.StatusInternalServerError, "Post can not be written.", &err)
	}

	c.JSON(http.StatusCreated, post.Id)

	return nil
}
