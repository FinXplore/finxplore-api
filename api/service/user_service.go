package service

import (
	"fmt"

	"github.com/Dhyey3187/finxplore-api/api/models"
	"github.com/Dhyey3187/finxplore-api/api/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(email, password, firstName, lastName, dialingCode, mobileNumber string) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(email, password, firstName, lastName, dialingCode, mobileNumber string) (*models.User, error) {
	// 1. Check if user exists
	existingUser, err := s.repo.GetUserByMobileNumber(dialingCode, mobileNumber)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, fmt.Errorf(
			"mobile number %s already linked with other account",
			dialingCode+" "+mobileNumber,
		)
	}

	// 2. Hash Password
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 3. Create User Model
	newUser := &models.User{
		Email:         email,
		Password:  string(hashedBytes),
		FirstName:      firstName,
		LastName:      lastName,
		DialingCode:   dialingCode,
		MobileNumber:  mobileNumber,
		Role:          "user",
		// Set defaults explicitly (good practice)
		RiskTolerance: "moderate",
		Currency:      "INR",
	}

	// 4. Save to DB
	err = s.repo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}