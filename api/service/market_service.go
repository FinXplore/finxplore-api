package service

import (
	"errors"
	"time"

	"github.com/Dhyey3187/finxplore-api/api/models"
	"github.com/Dhyey3187/finxplore-api/api/repository"
)

type MarketService interface {
	GetStockDetails(ticker string) (*models.Stock, error)
	GetStockChart(ticker string, period string) ([]models.StockPriceHistory, error)
	SearchStocks(query string) ([]models.Stock, error)
}

type marketService struct {
	stockRepo repository.StockRepository
}

func NewMarketService(stockRepo repository.StockRepository) MarketService {
	return &marketService{
		stockRepo: stockRepo,
	}
}

func (s *marketService) GetStockDetails(ticker string) (*models.Stock, error) {
	return s.stockRepo.GetByTicker(ticker)
}

func (s *marketService) GetStockChart(ticker string, period string) ([]models.StockPriceHistory, error) {
	// 1. Find the Stock ID first
	stock, err := s.stockRepo.GetByTicker(ticker)
	if err != nil {
		return nil, errors.New("stock not found")
	}

	// 2. Determine Date Range
	now := time.Now()
	var startDate time.Time

	switch period {
	case "1w":
		startDate = now.AddDate(0, 0, -7)
	case "1m":
		startDate = now.AddDate(0, -1, 0)
	case "1y":
		startDate = now.AddDate(-1, 0, 0)
	case "5y":
		startDate = now.AddDate(-5, 0, 0)
	default: // Default to 1 month
		startDate = now.AddDate(0, -1, 0)
	}

	return s.stockRepo.GetHistory(stock.ID, startDate, now)
}

func (s *marketService) SearchStocks(query string) ([]models.Stock, error) {
	if len(query) < 1 {
		return nil, nil // Don't search for single letters
	}
	return s.stockRepo.Search(query)
}