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
	user, err := h.service.RegisterUser(req.Email, req.Password, req.FirstName, req.LastName, req.DialingCode, req.MobileNumber)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	// 3. Response (Resource Transformation)
	response := dto.UserResponse{
		// ID:       user.ID.String(),
		Email:    user.Email,
		FullName: user.FirstName + " " + user.LastName,
		Role:     user.Role,
	}

	c.JSON(http.StatusCreated, response)
}