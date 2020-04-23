package main

import (
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func pass() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
func main() {
	//Default with the Logger and Recovery middleware already attached
	//r := gin.Default()

	//不使用默认的中间件创建 router
	r := gin.New()

	// Logger instances a Logger middleware that will write the logs to gin.DefaultWriter.
	// By default gin.DefaultWriter = os.Stdout.
	// Recovery returns a middleware that recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Logger(), gin.Recovery()) //使用中间件

	r.GET("/benchmark", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": "release.200",
			"msg":  "debug msg",
			"body": "reponse body",
		})
	})

	auth := r.Group("/")
	auth.Use(AuthRequired()) //使用中间件
	{
		auth.POST("/login", pass())
		auth.POST("/submit", pass())
		auth.POST("/read", pass())

		testing := auth.Group("testing")
		testing.GET("/analytics", pass())
	}

	r.Run(":8080")
}
