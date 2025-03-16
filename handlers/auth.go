package handlers

import (
	"net/http"

	"ethar/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey = []byte("your_secret_key") // Change this to a secure key in production

// Register handles the user registration process
func Register(c *gin.Context) {
	var user db.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Save user to the database
	database := db.GetDB() // Assumes you have a function to get the DB instance
	if err := database.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
