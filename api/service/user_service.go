package service

import (
	"fmt"
	"errors"
	"time"

	"github.com/Dhyey3187/finxplore-api/api/models"
	"github.com/Dhyey3187/finxplore-api/api/repository"
	"github.com/Dhyey3187/finxplore-api/internal/utils"
	"github.com/Dhyey3187/finxplore-api/internal/config"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(email, password, firstName, lastName, dialingCode, mobileNumber, currency string) (*models.User, error)
	LoginUser(dialingCode, mobileNumber, password string) (string, string, *models.User, error)
}

type userService struct {
	repo repository.UserRepository
	cacheRepo repository.CacheRepository
	cfg       *config.Config
}

func NewUserService(repo repository.UserRepository, cacheRepo repository.CacheRepository,cfg *config.Config) UserService {
	
	return &userService{
		repo:      repo,
		cacheRepo: cacheRepo,
		cfg:       cfg,
	}
}

func (s *userService) RegisterUser(email, password, firstName, lastName, dialingCode, mobileNumber, currency string) (*models.User, error) {
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
		RiskTolerance: "moderate",
		Currency:      currency,
	}

	// 4. Save to DB
	err = s.repo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *userService) LoginUser(dialingCode, mobileNumber, password string) (string, string, *models.User, error) {
	// 1. Find User & Verify Password
	user, err := s.repo.GetUserByMobileNumber(dialingCode, mobileNumber)
	if err != nil {
		return "", "", nil, errors.New("invalid credentials")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", nil, errors.New("invalid credentials")
	}

	// 2. Generate Tokens
	accessToken, err := utils.CreateAccessToken(user.UserCode, user.Role, s.cfg.JWTSecret)
	if err != nil {
		return "", "", nil, err
	}
	refreshToken := utils.CreateRefreshToken()

	// 3. Use CacheRepository to save session
	// Notice how clean this is! No contexts, no redis commands.
	redisKey := "refresh:" + user.UserCode
	err = s.cacheRepo.SetSession(redisKey, refreshToken, 7*24*time.Hour)
	if err != nil {
		return "", "", nil, errors.New("failed to save session")
	}

	return accessToken, refreshToken, user, nil
}