package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Name    string `form:"name"` // binding:"required"
	Address string `form:"address"`
}

//curl http://127.0.0.1:8080/person?name=zbb&address=Sichuang

func main() {
	r := gin.Default()

	r.GET("/person", testPerson)

	r.Run("127.0.0.1:8080")
}

func testPerson(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		c.JSON(http.StatusOK, gin.H{
			"name":    person.Name,
			"address": person.Address,
		})
	} else {
		c.JSON(http.StatusBadRequest, "Bad Request.")
	}
}
