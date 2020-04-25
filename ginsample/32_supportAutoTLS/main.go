package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

//TODO
// 没看明白这里的自动tls是怎么做的，以及启动的服务如何访问
func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	autotls.Run(r, "example01.com")

}