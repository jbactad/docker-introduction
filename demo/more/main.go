package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!\n")

		greetings := fmt.Sprintf("Hello, %s!", os.Getenv("PERSON_NAME"))

		c.String(200, greetings)
	})

	r.Run(":3000")
}