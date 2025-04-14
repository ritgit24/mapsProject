package database

import (
	"fmt"
	"log"
	"os"

	 "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global GORM database connection
var DB *gorm.DB

// InitDB initializes the GORM database connection
func InitDB() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	log.Println("Connected to the database")
}