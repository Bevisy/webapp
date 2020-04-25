package main

import "github.com/gin-gonic/gin"

// curl http://127.0.0.1:8080/JSONP?callback=x
// resp:
//	x({"foo":"bar"});
func main() {
	r := gin.Default()
	r.GET("/JSONP", func(c *gin.Context) {
		data := gin.H{
			"foo": "bar",
		}
		c.JSONP(200, data)
	})
	r.Run("127.0.0.1:8080")
}
