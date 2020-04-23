package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Request --
//curl --location --request POST 'http://127.0.0.1:8080/post?ids[a]=123&ids[b]=hello' \
//--header 'Content-Type: application/x-www-form-urlencoded' \
//--data-urlencode 'names[first]=first' \
//--data-urlencode 'names[second]=second'

func main() {
	r := gin.Default()

	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
		//fmt.Printf("ids: %v, names: %v.", ids, names)
	})

	r.Run(":8080")
}
