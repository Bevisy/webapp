package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	gin.DisableConsoleColor()

	//logging to file
	f, _ := os.Create("custom.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ping": "pong",
		})
	})

	r.Run(":8080")
}
