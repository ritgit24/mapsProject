package main

import (
	"chatbot-backend/database"
	"chatbot-backend/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Loading environment variables
	godotenv.Load()

	// Initialing the database
	database.InitDB()

	// set up the Gin router.The Gin router is a popular HTTP web framework for  Go.When we want a router with the default middleware (Logger and Recovery), we use gin.Default
	r := gin.Default()

	// to add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Allow requests from React frontend
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // No content for preflight requests
			return
		}
		c.Next()
	})

	// Define routes
	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)
	r.POST("/chat", handlers.GettingQueriesAnswer(database.DB, "your_session_secret_key", "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"))

	// to Start the server
	r.Run(":8080")
}