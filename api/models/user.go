package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	// -- Auth Data --
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	DialingCode string `gorm:"not null"`
	MobileNumber string `gorm:"uniqueIndex;not null"`
	Email        string    `gorm:"not null"`
	Password string    `gorm:"not null"`
	Role         string    `gorm:"default:'user'"` // 'user' or 'admin'
	
	// --- Profile Data (Merged) ---
	FirstName string `gorm:"not null"`
	LastName string `gorm:"not null"`
	RiskTolerance string `gorm:"default:'moderate'"`
	Currency      string `gorm:"default:'USD'"`
	AvatarURL     string 
	
	// --- System Fields ---
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Ensure that id is not null
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}