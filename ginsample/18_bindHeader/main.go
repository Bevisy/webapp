package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type testHeader struct {
	Name string `header:""name`
	Age  int    `header:"age"`
}

func main() {
	r := gin.Default()
	r.GET("/header", func(c *gin.Context) {
		var test testHeader
		if err := c.ShouldBindHeader(&test); err == nil {
			c.JSON(http.StatusOK, gin.H{"name": test.Name, "age": test.Age})
		} else {
			c.String(http.StatusBadRequest, "failure.")
		}
	})
	r.Run("127.0.0.1:8080")
}
