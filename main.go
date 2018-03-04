package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var message string

func main() {

	message = "Hello World"
	fmt.Println("Hello world")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, message)
	})

	r.POST("/post", func(c *gin.Context) {
		if len(c.Query("message")) > 0 {
			message = c.Query("message")
			fmt.Print(message)
			c.String(200, message)
			return
		}
		c.String(200, message)
	})

	r.Run(":8080")
}
