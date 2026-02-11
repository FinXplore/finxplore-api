package routes

import "github.com/gin-gonic/gin"

type Routes struct {
	UserRoutes *UserRoutes
	StockRoutes *StockRoutes
	AuthMiddleware gin.HandlerFunc
}

func NewRoutes(userRoutes *UserRoutes,stockRoutes *StockRoutes, authMiddleware gin.HandlerFunc) *Routes {
	return &Routes{
		UserRoutes:     userRoutes,
		StockRoutes:    stockRoutes,
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
	auth := api.Group("/auth")
	r.UserRoutes.Register(auth)
	stock := api.Group("/stock")
	r.StockRoutes.Register(stock)

	// 2. Protected Routes (Token Required)
	protected := auth.Group("/")
	stock_protected := stock.Group("/")
	protected.Use(r.AuthMiddleware)
	{
		r.UserRoutes.RegisterProtected(protected)
		// r.WalletRoutes.Register(protected) // Future
	}
	stock_protected.Use(r.AuthMiddleware)
	{
		r.StockRoutes.RegisterProtected(stock_protected)
	}
}
