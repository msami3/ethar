package main

import (
	database "ethar/db"
	"ethar/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// Import any other necessary packages
)

// User represents a user in the system (employer or job seeker)
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"` // "employer" or "job_seeker"
}

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

	// Start the server
	r.Run(":8080")
}
