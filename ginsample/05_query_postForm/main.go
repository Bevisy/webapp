package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultPostForm("page", "0")
		name := c.PostForm("name")
		msg := c.PostForm("msg")

		c.JSON(http.StatusOK, gin.H{
			"id":   id,
			"name": name,
			"page": page,
			"msg":  msg,
		})
	})

	r.Run(":8080")
}
