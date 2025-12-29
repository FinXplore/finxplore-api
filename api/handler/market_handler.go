package handler

import (
	"net/http"

	"github.com/Dhyey3187/finxplore-api/api/service"
	"github.com/Dhyey3187/finxplore-api/api/dto"
	"github.com/gin-gonic/gin"
)

type MarketHandler struct {
	marketService service.MarketService
}

func NewMarketHandler(marketService service.MarketService) *MarketHandler {
	return &MarketHandler{marketService: marketService}
}

// GET /stocks/:ticker
func (h *MarketHandler) GetStock(c *gin.Context) {
	ticker := c.Param("ticker")
	stock, err := h.marketService.GetStockDetails(ticker)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	// Map Model -> DTO
	response := dto.StockDetailResponse{
		Ticker:        stock.Ticker,
		Name:          stock.ShortName,
		Exchange:      stock.Exchange,
		Currency:      stock.Currency,
		Sector:        stock.Sector,
		Industry:      stock.Industry,
		Description:   stock.BusinessSummary,
		Website:       stock.Website,
		MarketCap:     stock.MarketCap,
		PE:            stock.TrailingPE,
		DividendYield: stock.DividendYield,
		LastUpdated:   stock.LastUpdated,
	}

	c.JSON(http.StatusOK, response)
}

// GET /stocks/:ticker/chart
func (h *MarketHandler) GetChart(c *gin.Context) {
	ticker := c.Param("ticker")
	period := c.DefaultQuery("period", "1m")

	history, err := h.marketService.GetStockChart(ticker, period)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Map List of Models -> List of DTOs
	var response []dto.StockCandleResponse
	for _, day := range history {
		response = append(response, dto.StockCandleResponse{
			Date:   day.TradeDate.Format("2006-01-02"), // YYYY-MM-DD
			Open:   day.OpenPrice,
			High:   day.HighPrice,
			Low:    day.LowPrice,
			Close:  day.ClosePrice,
			Volume: day.Volume,
		})
	}

	c.JSON(http.StatusOK, response)
}

// GET /stocks/search
func (h *MarketHandler) Search(c *gin.Context) {
	query := c.Query("q")
	results, err := h.marketService.SearchStocks(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Search failed"})
		return
	}

	var response []dto.StockSearchResponse
	for _, s := range results {
		response = append(response, dto.StockSearchResponse{
			Ticker:   s.Ticker,
			Name:     s.ShortName,
			Type:     s.QuoteType,
			Exchange: s.Exchange,
		})
	}

	c.JSON(http.StatusOK, response)
}