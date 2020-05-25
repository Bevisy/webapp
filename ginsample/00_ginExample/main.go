package main

import "github.com/gin-gonic/gin"

func main() {
	//Create a gin router with default middleware:
	//logger and recovery (crash-free) middleware
	router := gin.Default()

	//支持的请求方法包括：GET、POST、PUT、DELETE、PATCH、HEAD、OPTIONS
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "GET",
		})
	})

	router.POST("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "POST",
		})
	})

	router.PUT("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "PUT",
		})
	})

	router.DELETE("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "DELETE",
		})
	})

	router.PATCH("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "PATCH",
		})
	})

	router.HEAD("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "HEAD",
		})
	})

	router.OPTIONS("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "OPTIONS",
		})
	})

	//默认监听 ":8080"
	//router.Run()
	router.Run("0.0.0.0:8080")
}
