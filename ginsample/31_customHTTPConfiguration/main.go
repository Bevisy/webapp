package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong.")
	})

	//使用gin启动服务
	//r.Run("127.0.0.1:8080")

	//使用库net/http启动服务
	//http.ListenAndServe("127.0.0.1:8080", r)

	//配置 &http.Server 参数，启动服务
	s := &http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, //128 kb
	}
	s.ListenAndServe()
}
