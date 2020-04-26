package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

//TODO
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong.")
	})

	//借助 endless 实现优雅退出
	//windows 下无法执行
	endless.ListenAndServe("127.0.0.1:8080", r)
}
