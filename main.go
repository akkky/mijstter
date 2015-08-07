package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mijstter/db"
	"os"
)

const (
	databaseFile = "mijstter.sqlite3"
)

var (
	database *db.Database
)

func getPort() string {
	port := os.Getenv("MIJSTTER_PORT")
	if port != "" {
		return port
	}

	return "localhost:8080"
}

func getBaseUrl() string {
	baseUrl := os.Getenv("MIJSTTER_BASE_URL")
	if baseUrl != "" {
		return baseUrl
	}

	return "http://127.0.0.1:8080"
}

func AcceptCrossOrigin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

		c.Next()
	}
}

func _main() (int, error) {
	db, err := db.NewDatabase(databaseFile)
	if err != nil {
		return 1, err
	}
	defer db.Close()
	database = db

	r := gin.Default()

	r.Use(AcceptCrossOrigin())
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// スタティックリソース
	r.Static("/files", "./files")

	// POST /users
	r.POST("/users", NewUser)
	// POST /login
	r.POST("/login", Login)
	// GET /posts
	r.GET("/posts", GetPosts)
	// POST /posts
	r.POST("/posts", NewPost)
	// POST /images
	r.POST("/images", NewImage)

	r.Run(getPort())

	return 0, nil
}

func main() {
	if status, err := _main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(status)
	}
}
