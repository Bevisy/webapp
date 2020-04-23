package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// API分组
	v1 := r.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "login.")
		})
		v1.POST("/submit", func(c *gin.Context) {
			c.String(http.StatusOK, "submit.")
		})
		v1.POST("/read", func(c *gin.Context) {
			c.String(http.StatusOK, "read.")
		})
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "login.")
		})
		v2.POST("/submit", func(c *gin.Context) {
			c.String(http.StatusOK, "submit.")
		})
		v2.POST("/read", func(c *gin.Context) {
			c.String(http.StatusOK, "read.")
		})
	}

	r.Run(":8080")
}
