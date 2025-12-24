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
	user, err := h.service.RegisterUser(req.Email, req.Password, req.FirstName, req.LastName, req.DialingCode, req.MobileNumber, req.Currency)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	// 3. Response (Resource Transformation)
	response := dto.UserResponse{
		FullName: user.FirstName + " " + user.LastName,
		Role:     user.Role,
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
	accessToken, refreshToken, user, err := h.service.LoginUser(req.DialingCode, req.MobileNumber, req.Password)
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

	accessToken, err := h.service.RefreshAccessToken(req.DialingCode, req.MobileNumber, req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.RefreshResponse{
		AccessToken: accessToken,
	})
}