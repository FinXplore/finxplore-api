package routes

import "github.com/gin-gonic/gin"

type Routes struct {
	UserRoutes *UserRoutes
}

func NewRoutes(userRoutes *UserRoutes) *Routes {
	return &Routes{
		UserRoutes: userRoutes,
	}
}

func (r *Routes) Register(router *gin.Engine) {
	// Health check can also live here
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "up",
		})
	})

	api := router.Group("/api/v1")
	r.UserRoutes.Register(api)
}
