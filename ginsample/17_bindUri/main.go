package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Name string `uri:"name" binding:"required"`
	ID string `uri:"id" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"name": person.Name,
				"id": person.ID,
			})
		} else {
			c.String(http.StatusBadRequest, err.Error())
		}
	})
	r.Run("127.0.0.1:8080")
}
