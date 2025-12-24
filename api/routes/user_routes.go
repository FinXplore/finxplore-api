package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Dhyey3187/finxplore-api/api/handler"
)

type UserRoutes struct {
	handler *handler.AuthHandler
}

func NewUserRoutes(handler *handler.AuthHandler) *UserRoutes {
	return &UserRoutes{handler: handler}
}

func (r *UserRoutes) Register(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	auth.POST("/login", r.handler.Login)
	auth.POST("/register", r.handler.Register)
	auth.POST("/refresh",r.handler.Refresh)
}

// RegisterProtected handles /me (and user profile updates)
func (r *UserRoutes) RegisterProtected(rg *gin.RouterGroup) {
	rg.GET("/me", func(c *gin.Context) {
		// Just a test handler for now
		userCode, _ := c.Get("user_code")
		role, _ := c.Get("role")
		c.JSON(200, gin.H{
			"message":   "You are authorized!",
			"user_code": userCode,
			"role":      role,
		})
	})
}