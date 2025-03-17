package main

import (
	database "ethar/db"
	"ethar/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Init()

	// Initialize Gin router
	r := gin.Default()

	// Add a simple route to respond to GET requests at the root ("/")
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	// Registration route
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Start the server
	r.Run(":8080")
}
