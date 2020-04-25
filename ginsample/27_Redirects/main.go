package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//方案一
	r.GET("/test1", func(c *gin.Context) {
		//重定向
		c.Redirect(http.StatusMovedPermanently, "http://google.com")
	})

	//方案二
	r.GET("/test2", func(c *gin.Context) {
		c.Request.URL.Path = "/test3"
		r.HandleContext(c)
	})
	r.GET("/test3", func(c *gin.Context) {
		c.JSON(200, gin.H{"Hello": "World."})
	})

	r.Run("127.0.0.1:8080")
}
