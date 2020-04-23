package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//匹配 /usr/kitty ，不匹配 /user/ 或者 /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	//匹配 /user/kitty/ , /user/kitty/send
	//如果没有其它routers匹配 /user/kitty，将被重定向至 /user/kitty/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	//对于每个匹配的请求，上下文将保留路由定义
	router.POST("/user/:name/*action", func(c *gin.Context) {
		if c.FullPath() == "/user/:name/*action" {
			c.String(http.StatusOK, "url matched.")
		}
	})

	router.Run(":8080")
}
