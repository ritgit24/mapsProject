package handlers

import (
	"chatbot-backend/database"
	"chatbot-backend/models"
	"chatbot-backend/utility"
	
	"log"
	"net/http"
"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// function Signup handles user registration
func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Create the user in the database
	if err := database.DB.Create(&user).Error; err != nil {
		log.Printf("Error creating user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	// Generate a JWT token. utility has a function Generate token which we are using right now to generate the token using the email,password and time
	token, err := utility.GenerateToken(user.Email, "your_session_secret_key")
	if err != nil {
		log.Printf("Error generating token: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Return the token to the client
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Login handles user authentication
func Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

// about ShouldBindJson
// Parses the JSON: It reads the JSON data from the request body.

// Binds to a struct: It maps the JSON fields to the corresponding fields in a Go struct based on their names and types.

// Handles errors: If the JSON is malformed or doesn't match the struct, it returns an error.
	if err := c.ShouldBindJSON(&credentials); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Retrieves the user from the database
	var user models.User
	if err := database.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		log.Printf("Error retrieving user: %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
    
	// this creates a session for the newly signed-up user
	sessionTitle := fmt.Sprintf("User %d Signup Session", user.ID) //using the user's ID in the session title
	if err := database.DB.Exec(`
		INSERT INTO sessions (session_title)
		VALUES (?)
	`, sessionTitle).Error; err != nil {
		log.Printf("Error creating session: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create session"})
		return
	}

	// Compare the provided password with the hashed password in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		log.Printf("Error comparing passwords: %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate a JWT token
	token, err := utility.GenerateToken(user.Email, "your_session_secret_key")
	if err != nil {
		log.Printf("Error generating token: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Return the token to the client
	// c.JSON(http.StatusOK, gin.H{"token": token})
	// c.JSON(http.StatusOK, gin.H{"name": user.Name}) this way 2 json objects were generated.which is wrong. so one json should be generated
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"name":  user.Name,
	  })
}