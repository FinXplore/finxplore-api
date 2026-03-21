package repository

import (
	"time"

	"github.com/Dhyey3187/finxplore-api/api/models"
	"gorm.io/gorm"
)

type StockRepository interface {
	GetByTicker(ticker string) (*models.Stock, error)
	GetHistory(stockID uint, start, end time.Time) ([]models.StockPriceHistory, error)
	Search(query string) ([]models.Stock, error)
	GetNews(stockID uint, limit int) ([]models.News, error)
	GetIncomeStatements(stockID uint) ([]models.IncomeStatement, error)
	GetBalanceSheets(stockID uint) ([]models.BalanceSheet, error)
	GetCashFlows(stockID uint) ([]models.CashFlow, error)
}

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &stockRepository{db: db}
}

// 1. Get Single Stock Profile
func (r *stockRepository) GetByTicker(ticker string) (*models.Stock, error) {
	var stock models.Stock
	// Preload latest price or just basic info?
	// Let's grab the basic info. Real-time price usually comes from Redis (later topic),
	// but for now, we rely on what Python synced.
	err := r.db.Where("ticker = ?", ticker).First(&stock).Error
	return &stock, err
}

// 2. Get Chart Data
func (r *stockRepository) GetHistory(stockID uint, start, end time.Time) ([]models.StockPriceHistory, error) {
	var history []models.StockPriceHistory
	err := r.db.Where("stock_id = ? AND trade_date BETWEEN ? AND ?", stockID, start, end).
		Order("trade_date ASC").
		Find(&history).Error
	return history, err
}

// 3. Search Stocks (Simple ILIKE for now)
func (r *stockRepository) Search(query string) ([]models.Stock, error) {
	var stocks []models.Stock
	searchPattern := "%" + query + "%"
	
	// Search by Ticker OR Company Name
	// Limit to 10 results for speed
	err := r.db.Where("ticker ILIKE ? OR short_name ILIKE ?", searchPattern, searchPattern).
		Limit(10).
		Find(&stocks).Error
		
	return stocks, err
}

func (r *stockRepository) GetNews(stockID uint, limit int) ([]models.News, error) {
	var stock models.Stock
	err := r.db.Preload("News", func(db *gorm.DB) *gorm.DB {
		return db.Order("published_at desc").Limit(limit)
	}).Where("id = ?", stockID).First(&stock).Error
	return stock.News, err
}

func (r *stockRepository) GetIncomeStatements(stockID uint) ([]models.IncomeStatement, error) {
	var statements []models.IncomeStatement
	err := r.db.Where("stock_id = ?", stockID).Order("period_end_date desc").Find(&statements).Error
	return statements, err
}

func (r *stockRepository) GetBalanceSheets(stockID uint) ([]models.BalanceSheet, error) {
	var statements []models.BalanceSheet
	err := r.db.Where("stock_id = ?", stockID).Order("period_end_date desc").Find(&statements).Error
	return statements, err
}

func (r *stockRepository) GetCashFlows(stockID uint) ([]models.CashFlow, error) {
	var flows []models.CashFlow
	err := r.db.Where("stock_id = ?", stockID).Order("period_end_date desc").Find(&flows).Error
	return flows, err
}