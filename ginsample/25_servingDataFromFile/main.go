package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/file", func(c *gin.Context) {
		c.File("test/test1.txt")
	})

	//var fs http.FileSystem
	//r.GET("/file2", func(c *gin.Context) {
	//	c.FileFromFS("test/test.txt", fs)
	//})

	r.Run("127.0.0.1:8080")
}
