package models

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"gorm.io/gorm"
)



type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

// CreateUser inserts a new user into the database
// func CreateUser(db *sql.DB, user *User) error {
// 	query := `
// 		INSERT INTO users (name, email, password)
// 		VALUES ($1, $2, $3)
// 		RETURNING id
// 	`
// 	err := db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func CreateUser(db *sql.DB, user *User) error {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		log.Printf("Error executing SQL query: %v\n", err) // Log the error
		return err
	}
	return nil
}

// GetUserByEmail retrieves a user from the database by email
func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	user := &User{}
	query := `
		SELECT id, name, email, password
		FROM users
		WHERE email = $1
	`
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

// ValidateUserCredentials checks if the provided email and password match a user in the database
func ValidateUserCredentials(db *sql.DB, email, password string) (*User, error) {
	user, err := GetUserByEmail(db, email)
	if err != nil {
		return nil, err
	}

	// Compare the provided password with the hashed password in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
// GetAllUsers retrieves all users from the database
func GetAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
