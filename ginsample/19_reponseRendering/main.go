package main

import (
	"github.com/gin-gonic/gin"
)

//返回的消息体渲染
func main() {
	r := gin.Default()
	//gin.H{} map[string]interface{}
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("xml", func(c *gin.Context) {
		c.XML(200, gin.H{"status": "ok"})
	})

	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(200, gin.H{"status": "ok"})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		var msg struct{
			Name string
			Age int
			Address string
		}
		msg.Name = "zbb"
		msg.Age = 12
		msg.Address = "Beijing"

		c.JSON(200, msg)
	})

	r.Run("127.0.0.1:8080")
}
