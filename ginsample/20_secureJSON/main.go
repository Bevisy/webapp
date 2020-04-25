package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// output:
	//	while(1);[
	//		"a",
	//		"ab",
	//		"abc"
	//	]
	r.GET("/secureJSON", func(c *gin.Context) {
		names := []string{"a", "ab", "abc"}
		//防止json劫持；默认添加 "while(1);"
		c.SecureJSON(200, names)
	})
	r.Run("127.0.0.1:8080")
}
