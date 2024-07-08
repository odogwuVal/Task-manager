package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Task Manager API Project")

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		// Render the HTML template
		c.String(200, "Welcome to Task API")
	})
	router.Run()
}
