package handler

import (
	"net/http"

	"github.com/Dhyey3187/finxplore-api/api/dto"
	"github.com/Dhyey3187/finxplore-api/api/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.UserService
}

func NewAuthHandler(s service.UserService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	// 1. Validation (Middleware logic inside Gin)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Call Service
	accessToken, refreshToken, user, err := h.service.RegisterUser(req.Email, req.Password, req.FirstName, req.LastName, req.DialingCode, req.MobileNumber, req.Currency)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	// 3. Response (Resource Transformation)
	response := dto.LoginResponse{
		FullName: user.FirstName + " " + user.LastName,
		Role:     user.Role,
		UserCode:     user.UserCode,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	// 1. Validation (Middleware logic inside Gin)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Call Service
	accessToken, refreshToken, user, err := h.service.LoginUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	// 3. Response (Resource Transformation)
	response := dto.LoginResponse{
		FullName: user.FirstName + " " + user.LastName,
		Role:     user.Role,
		UserCode:     user.UserCode,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	var req dto.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := h.service.RefreshAccessToken(req.Email, req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.RefreshResponse{
		AccessToken: accessToken,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	userCode, _ := c.Get("user_code")
	if err := h.service.LogoutUser(userCode.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to logout"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}

func (h *AuthHandler) GetMe(c *gin.Context) {
	userCode, _ := c.Get("user_code")
	user, err := h.service.GetProfile(userCode.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ProfileResponse{
		UserCode:      user.UserCode,
		Email:         user.Email,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		AvatarURL:     user.AvatarURL,
		Currency:      user.Currency,
		Role:          user.Role,
		RiskTolerance: user.RiskTolerance,
	})
}

func (h *AuthHandler) UpdateMe(c *gin.Context) {
	userCode, _ := c.Get("user_code")
	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.UpdateProfile(userCode.(string), req.FirstName, req.LastName, req.AvatarURL, req.Currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ProfileResponse{
		UserCode:      user.UserCode,
		Email:         user.Email,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		AvatarURL:     user.AvatarURL,
		Currency:      user.Currency,
		Role:          user.Role,
		RiskTolerance: user.RiskTolerance,
	})
}

func (h *AuthHandler) ChangePassword(c *gin.Context) {
	userCode, _ := c.Get("user_code")
	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.ChangePassword(userCode.(string), req.CurrentPassword, req.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password updated successfully"})
}