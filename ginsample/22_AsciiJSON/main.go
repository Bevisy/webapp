package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/asc2json", func(c *gin.Context) {
		data := gin.H{
			"lang": "GO语言",
			"string": "!@#$%^&*()",
			"tag": "<br>",
		}
		// output:
		// 	{"lang":"GO\u8bed\u8a00","string":"!@#$%^\u0026*()","tag":"\u003cbr\u003e"}
		c.AsciiJSON(200, data)
	})
	r.Run("127.0.0.1:8080")
}
