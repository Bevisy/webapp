package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//静态文件目录
	r.Static("assets", "./assets")
	//静态文件系统，例如文件夹并显示文件夹内容
	r.StaticFS("/more_static", http.Dir("my_files"))
	//静态文件
	r.StaticFile("screenshot.jpg", "./resource/screenshot.jpg")

	r.Run("127.0.0.1:8080")
}
