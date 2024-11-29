package main

import (
	"exptrack/models"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Main function to set up the server, database, and routes
func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the models (create tables with relationships and auto-incrementing IDs)
	err = db.AutoMigrate(&models.User{}, &models.Expense{}, &models.Category{}, &models.Budget{})
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Seed the database with initial data
	seedDatabase()

	// Initialize Gin router
	r := gin.Default()

	// Define routes
	r.GET("/users", getUsers)
	r.POST("/users", createUser)
	r.GET("/expenses", getExpenses)
	r.POST("/expenses", createExpense)

	// Start the server
	r.Run(":8080")
}

// SeedDatabase is a simple function to seed initial data
func seedDatabase() {
	// Seed Users
	db.Create(&models.User{Name: "John Doe", Email: "johndoe@example.com", Password: "password123", Currency: "USD"})

	// Seed Categories
	db.Create(&models.Category{Name: "Food", Description: "All food-related expenses"})
	db.Create(&models.Category{Name: "Transport", Description: "Transportation-related expenses"})
}

// GetUsers retrieves all users from the database
func getUsers(c *gin.Context) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(200, users)
}

// CreateUser creates a new user in the database
func createUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	if err := db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(201, user)
}

// GetExpenses retrieves all expenses from the database
func getExpenses(c *gin.Context) {
	var expenses []models.Expense
	if err := db.Find(&expenses).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch expenses"})
		return
	}
	c.JSON(200, expenses)
}

// CreateExpense creates a new expense in the database
func createExpense(c *gin.Context) {
	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	if err := db.Create(&expense).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create expense"})
		return
	}
	c.JSON(201, expense)
}
