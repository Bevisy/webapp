package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name string `form:"name"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2020-01-01" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime time.Time `form:"unixTime" time_format:"unixNano"`
}

func main() {
	r := gin.Default()

	r.GET("/person", testPerson)

	r.Run("127.0.0.1:8080")
}

func testPerson(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err == nil {
		log.Printf("%s who lived in %s was born in %s.", person.Name, person.Address, person.Birthday)
	}
	c.String(http.StatusOK, "OK.")
}
