package main

import "github.com/gin-gonic/gin"

func main() {
	//关闭日志颜色输出
	//gin.DisableConsoleColor()

	//强制日志颜色输出
	//gin.ForceConsoleColor()

	//默认日志颜色输出由 TTY 决定
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Run(":8080")
}
