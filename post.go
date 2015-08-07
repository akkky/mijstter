package main

import (
	_ "fmt"
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
	posts := make([]model.Post, 0, 10)

	for i := 0; i < 10; i++ {
		post := model.Post{
			Model:    model.Model{Id: model.ID(i + 1)},
			UserId:   model.ID(1),
			Message:  "こんにちは！",
			Url:      "https://www.mijs.jp/wp-content/themes/mijs-site/images/common/h1-logo.jpg",
			UserName: "mijs",
		}
		posts = append(posts, post)
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
	err := c.BindJSON(&post)
	if err != nil {
		return NewError(http.StatusBadRequest, "Bad Request", &err)
	}

	//	err = database.WriteUser(&user)
	//	if err != nil {
	//		return NewError(http.StatusInternalServerError, "User can not be written.", &err)
	//	}

	c.JSON(http.StatusCreated, post.Id)

	return nil
}
