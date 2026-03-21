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

func (h *MarketHandler) GetIndices(c *gin.Context) {
	data, err := h.marketService.GetIndices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *MarketHandler) GetGainers(c *gin.Context) {
	exchange := c.DefaultQuery("exchange", "NSE")
	data, err := h.marketService.GetMarketMovers(exchange, 10, "gainers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *MarketHandler) GetLosers(c *gin.Context) {
	exchange := c.DefaultQuery("exchange", "NSE")
	data, err := h.marketService.GetMarketMovers(exchange, 10, "losers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *MarketHandler) GetVolumeLeaders(c *gin.Context) {
	exchange := c.DefaultQuery("exchange", "NSE")
	data, err := h.marketService.GetMarketMovers(exchange, 10, "volume_leaders")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *MarketHandler) GetQuote(c *gin.Context) {
	ticker := c.Param("ticker")
	data, err := h.marketService.GetStockQuote(ticker)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *MarketHandler) GetNews(c *gin.Context) {
	ticker := c.Param("ticker")
	data, err := h.marketService.GetStockNews(ticker, 20)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *MarketHandler) GetFinancials(c *gin.Context) {
	ticker := c.Param("ticker")
	stmtType := c.DefaultQuery("type", "income")
	data, err := h.marketService.GetStockFinancials(ticker, stmtType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}