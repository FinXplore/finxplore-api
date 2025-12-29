package routes

import (
	"github.com/Dhyey3187/finxplore-api/api/handler"
	"github.com/gin-gonic/gin"
)

type StockRoutes struct {
	marketHandler  *handler.MarketHandler
}

func NewStockRoutes(marketHandler *handler.MarketHandler) *StockRoutes {
	return &StockRoutes{
		marketHandler:  marketHandler,
	}
}

func (r *StockRoutes) Register(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "Route Working",
			"user_code": "Test",
		})
	})
}


func (r *StockRoutes) RegisterProtected(rg *gin.RouterGroup) {
	rg.GET("/search", r.marketHandler.Search)
	// Final URL: GET /api/v1/stock/AAPL
	rg.GET("/:ticker", r.marketHandler.GetStock)
		
	// Final URL: GET /api/v1/stock/AAPL/chart
	rg.GET("/:ticker/chart", r.marketHandler.GetChart)
}