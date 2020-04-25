package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong.")
	})
	//服务启动添加TLS证书
	r.RunTLS("127.0.0.1:8080", "./cert/sample.crt", "./cert/sample.key")
}
