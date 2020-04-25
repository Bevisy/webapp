package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// Serves unicode entities
	//output: {"html":"\u003cb\u003eHello, world!\u003c/b\u003e"}
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// Serves literal characters
	// output: {"html":"<b>Hello, world!</b>"}
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run("127.0.0.1:8080")
}
