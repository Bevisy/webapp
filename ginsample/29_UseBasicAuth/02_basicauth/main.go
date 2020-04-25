package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//添加basicauth认证插件，插件指定认证信息;
	r.GET("/test", gin.BasicAuth(gin.Accounts{"zbb": "123"}), func(c *gin.Context) {
		c.JSON(200, gin.H{c.MustGet("user").(string): "passed"})
	})
	r.Run("127.0.0.1:8080")
}
