package models

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"gorm.io/gorm"
)

type User struct {
	// -- Auth Data --
	ID uint `gorm:"primaryKey"`
	UserCode string `gorm:"uniqueIndex;not null;size:10"`
	DialingCode string `gorm:"not null"`
	MobileNumber string `gorm:"uniqueIndex;not null"`
	Email        string
	Password string    `gorm:"not null"`
	Role         string    `gorm:"default:'user'"` // 'user' or 'admin'
	
	// --- Profile Data (Merged) ---
	FirstName string `gorm:"not null"`
	LastName string
	RiskTolerance string `gorm:"default:'moderate'"`
	Currency      string `gorm:"default:'USD'"`
	AvatarURL     string 
	
	// --- System Fields ---
	CreatedAt time.Time
	UpdatedAt time.Time
	IsRistricted bool `gorm:"default:false"`
	RistrictedBy *uint
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Ensure that id is not null
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UserCode == "" {
		// Keep trying until we find a unique code (unlikely to collide, but safe)
		for {
			code := generateUserCode()
			// Check if this code already exists in DB
			var count int64
			tx.Model(&User{}).Where("user_code = ?", code).Count(&count)
			if count == 0 {
				u.UserCode = code
				break
			}
		}
	}
	return
}

func generateUserCode() string {
	const prefix = "FX" // FinXplore Prefix
	
	// Secure random number generation
	max := big.NewInt(99999999) // Max value for 8 digits
	n, _ := rand.Int(rand.Reader, max)
	
	// Format: FX + 8 digits (padded with zeros if number is small)
	return fmt.Sprintf("%s%08d", prefix, n.Int64())
}