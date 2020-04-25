package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Gin provides two sets of methods for binding:
//
//Type - Must bind
//Methods - Bind, BindJSON, BindXML, BindQuery, BindYAML, BindHeader
//Behavior - These methods use MustBindWith under the hood. If there is a binding error, the request is aborted with c.AbortWithError(400, err).SetType(ErrorTypeBind). This sets the response status code to 400 and the Content-Type header is set to text/plain; charset=utf-8. Note that if you try to set the response code after this, it will result in a warning [GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422. If you wish to have greater control over the behavior, consider using the ShouldBind equivalent method.
//Type - Should bind
//Methods - ShouldBind, ShouldBindJSON, ShouldBindXML, ShouldBindQuery, ShouldBindYAML, ShouldBindHeader
//Behavior - These methods use ShouldBindWith under the hood. If there is a binding error, the error is returned and it is the developer's responsibility to handle the request and error appropriately.

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	gin.ForceConsoleColor()
	r := gin.Default()

	// Example for binding JSON ({"user": "manu", "password": "123"})
	r.GET("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if json.User != "zbb" || json.Password != "123456" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "unauthorized.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "you are logged in.",
		})
	})

	// Example for binding XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>user</user>
	//		<password>123</password>
	//	</root>)
	r.GET("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if xml.User != "zbb" || xml.Password != "123456" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "unauthorized."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "logged in."})
	})

	// Example for binding a HTML form (user=manu&password=123)
	r.GET("/loginForm", func(c *gin.Context) {
		var form Login
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if form.User != "zbb" || form.Password != "123456" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "unauthorized."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "logged in."})
	})

	r.Run(":8080")

}
