package models

import (
	"time"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	SessionTitle string    `gorm:"column:session_title"`
	Context      string    `gorm:"type:jsonb"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UserID       uint      // Link session to a user
}

// CreateSession creates a new session in the database
func CreateSession(db *gorm.DB, sessionTitle string, context string, userID uint) (*Session, error) {
	session := Session{
		SessionTitle: sessionTitle,
		Context:      context,
		UserID:       userID,
	}
	result := db.Create(&session)
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}

// GetSessionByID retrieves a session by its ID
func GetSessionByID(db *gorm.DB, sessionID uint) (*Session, error) {
	var session Session
	result := db.First(&session, sessionID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}