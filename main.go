package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func _main() (int, error) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Run(":8080")

	return 0, nil
}

func main() {
	if status, err := _main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(status)
	}
}
