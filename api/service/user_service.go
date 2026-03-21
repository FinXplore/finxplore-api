package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/Dhyey3187/finxplore-api/api/models"
	"github.com/Dhyey3187/finxplore-api/api/repository"
	"github.com/Dhyey3187/finxplore-api/internal/config"
	"github.com/Dhyey3187/finxplore-api/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(email, password, firstName, lastName, dialingCode, mobileNumber, currency string) (string, string, *models.User, error)
	LoginUser(email, password string) (string, string, *models.User, error)
	RefreshAccessToken(email, refreshToken string) (string, error)
	LogoutUser(userCode string) error
	GetProfile(userCode string) (*models.User, error)
	UpdateProfile(userCode, firstName, lastName, avatarURL, currency string) (*models.User, error)
	ChangePassword(userCode, currentPassword, newPassword string) error
}

type userService struct {
	repo      repository.UserRepository
	cacheRepo repository.CacheRepository
	cfg       *config.Config
}

func NewUserService(repo repository.UserRepository, cacheRepo repository.CacheRepository, cfg *config.Config) UserService {

	return &userService{
		repo:      repo,
		cacheRepo: cacheRepo,
		cfg:       cfg,
	}
}

func (s *userService) RegisterUser(email, password, firstName, lastName, dialingCode, mobileNumber, currency string) (string, string, *models.User, error) {
	// 1. Check if user exists
	existingUser, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", "", nil, err
	}
	if existingUser != nil {
		return "", "", nil, fmt.Errorf(
			"email %s already linked with other account",
			email,
		)
	}

	// 2. Hash Password
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", "", nil, err
	}

	// 3. Create User Model
	newUser := &models.User{
		Email:         email,
		Password:      string(hashedBytes),
		FirstName:     firstName,
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
		return "", "", nil, err
	}

	// 5. Generate Tokens
	accessToken, err := utils.CreateAccessToken(newUser.UserCode, newUser.Role, s.cfg.JWTSecret)
	if err != nil {
		return "", "", nil, err
	}
	refreshToken := utils.CreateRefreshToken()

	redisKey := "refresh:" + newUser.UserCode
	err = s.cacheRepo.SetSession(redisKey, refreshToken, 7*24*time.Hour)
	if err != nil {
		return "", "", nil, errors.New("failed to save session")
	}

	return accessToken, refreshToken, newUser, nil
}

func (s *userService) LoginUser(email, password string) (string, string, *models.User, error) {
	// 1. Find User & Verify Password
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", "", nil, errors.New("invalid credentials")
	}
	// Check if user exists (repository returns nil, nil when not found)
	if user == nil {
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

func (s *userService) RefreshAccessToken(email, refreshToken string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}
	// Check if user exists (repository returns nil, nil when not found)
	if user == nil {
		return "", errors.New("user not found")
	}

	// Redis Key is still based on UserCode (refresh:FX...)
	redisKey := "refresh:" + user.UserCode
	storedToken, err := s.cacheRepo.GetSession(redisKey)
	if err != nil {
		return "", errors.New("session expired or invalid")
	}

	if storedToken != refreshToken {
		return "", errors.New("invalid refresh token")
	}

	// 4. Generate NEW Access Token
	newAccessToken, err := utils.CreateAccessToken(user.UserCode, user.Role, s.cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}

func (s *userService) LogoutUser(userCode string) error {
	redisKey := "refresh:" + userCode
	return s.cacheRepo.DeleteSession(redisKey)
}

func (s *userService) GetProfile(userCode string) (*models.User, error) {
	user, err := s.repo.GetUserByCode(userCode)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *userService) UpdateProfile(userCode, firstName, lastName, avatarURL, currency string) (*models.User, error) {
	user, err := s.repo.GetUserByCode(userCode)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	if firstName != "" {
		user.FirstName = firstName
	}
	user.LastName = lastName
	user.AvatarURL = avatarURL
	user.Currency = currency

	if err := s.repo.UpdateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) ChangePassword(userCode, currentPassword, newPassword string) error {
	user, err := s.repo.GetUserByCode(userCode)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(currentPassword))
	if err != nil {
		return errors.New("invalid current password")
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedBytes)
	return s.repo.UpdateUser(user)
}
