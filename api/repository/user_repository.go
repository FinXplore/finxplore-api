package repository

import (
	"errors"

	"github.com/Dhyey3187/finxplore-api/api/models"
	"gorm.io/gorm"
)

// define the interface (best practice for testing)
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByMobileNumber(dialingCode, mobileNumber string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// CreateUser saves a new user to the database
func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

// GetUserByEmail finds a user by email (useful for login/duplicate check)
func (r *userRepository) GetUserByMobileNumber(dialingCode, mobileNumber string) (*models.User, error) {
	var user models.User
	// queries the 'users' table
	err := r.db.Where("dialing_code = ? AND mobile_number = ?", dialingCode, mobileNumber).First(&user).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Not an error, just no user found
		}
		return nil, err
	}
	
	return &user, nil
}