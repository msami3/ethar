package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Global DB variable
var DB *gorm.DB

// User represents a user model
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"` // "employer" or "job_seeker"
}

// Job represents a job model
type Job struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	EmployerID  uint   `json:"employer_id"`
}

// Initialize the database connection and run migrations
func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("db/ethar.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Run migrations
	err = DB.AutoMigrate(&User{}, &Job{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
