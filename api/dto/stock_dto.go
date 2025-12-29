package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

// 1. Full Stock Profile Response
type StockDetailResponse struct {
	Ticker            string          `json:"ticker"`
	Name              string          `json:"name"` // Mapped from ShortName
	Exchange          string          `json:"exchange"`
	Currency          string          `json:"currency"`
	Sector            string          `json:"sector"`
	Industry          string          `json:"industry"`
	Description       string          `json:"description"` // Mapped from BusinessSummary
	Website           string          `json:"website"`
	MarketCap         int64           `json:"market_cap"`
	PE                decimal.Decimal `json:"pe_ratio"`       // TrailingPE
	DividendYield     decimal.Decimal `json:"dividend_yield"` // In percentage
	LastUpdated       time.Time       `json:"last_updated"`
}

// 2. Chart Data (Candlestick)
type StockCandleResponse struct {
	Date   string          `json:"date"` // Format "2024-01-01"
	Open   decimal.Decimal `json:"open"`
	High   decimal.Decimal `json:"high"`
	Low    decimal.Decimal `json:"low"`
	Close  decimal.Decimal `json:"close"`
	Volume int64           `json:"volume"`
}

// 3. Search Result (Lightweight)
type StockSearchResponse struct {
	Ticker   string `json:"ticker"`
	Name     string `json:"name"`
	Type     string `json:"type"`     // Equity, ETF
	Exchange string `json:"exchange"`
}