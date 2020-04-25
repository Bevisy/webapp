package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/someFileFromReader", func(c *gin.Context) {
		reponse, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || reponse.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		reader := reponse.Body
		contentLength := reponse.ContentLength
		contentType := reponse.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(200, contentLength, contentType, reader, extraHeaders)
	})

	r.Run("127.0.0.1:8080")
}
