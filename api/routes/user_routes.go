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
	router.POST("/login", r.handler.Login)
	router.POST("/register", r.handler.Register)
	router.POST("/refresh",r.handler.Refresh)
}

// RegisterProtected handles /me (and user profile updates)
func (r *UserRoutes) RegisterProtected(rg *gin.RouterGroup) {
	rg.POST("/logout", r.handler.Logout)
	rg.GET("/me", r.handler.GetMe)
	rg.PUT("/me", r.handler.UpdateMe)
	rg.POST("/change-password", r.handler.ChangePassword)
}