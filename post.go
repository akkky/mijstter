package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mijstter/model"
	"net/http"
)

func GetPosts(c *gin.Context) {
	err := getPosts(c)
	if err != nil {
		c.JSON(err.Status, err)
	}
}

func getPosts(c *gin.Context) *Error {
	posts, err := database.ReadPosts(20)
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

	fmt.Printf("Post : %v\n", post)

	if post.SessionId != "" {
		user := GetUserBySessionId(post.SessionId)
		post.UserId = user.Id
		post.UserName = user.UserName
	}

	err = database.WritePost(&post)
	if err != nil {
		return NewError(http.StatusInternalServerError, "Post can not be written.", &err)
	}

	c.JSON(http.StatusCreated, post.Id)

	return nil
}
