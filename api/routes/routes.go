package routes

import "github.com/gin-gonic/gin"

type Routes struct {
	UserRoutes *UserRoutes
	AuthMiddleware gin.HandlerFunc
}

func NewRoutes(userRoutes *UserRoutes, authMiddleware gin.HandlerFunc) *Routes {
	return &Routes{
		UserRoutes:     userRoutes,
		AuthMiddleware: authMiddleware,
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

	// 2. Protected Routes (Token Required)
	protected := api.Group("/")
	protected.Use(r.AuthMiddleware)
	{
		r.UserRoutes.RegisterProtected(protected)
		// r.WalletRoutes.Register(protected) // Future
	}
}
