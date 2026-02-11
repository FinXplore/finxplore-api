package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// ---------------------------------------------------------
// 1. Core Company Profile
// ---------------------------------------------------------
type Stock struct {
	ID        uint   `gorm:"primaryKey"`
	Ticker    string `gorm:"size:20;uniqueIndex;not null"`
	Exchange  string `gorm:"size:50"`
	Currency  string `gorm:"size:10"`
	QuoteType string `gorm:"size:20"` // Equity, ETF, etc.

	// Profile
	LongName          string
	ShortName         string
	Sector            string
	Industry          string
	Country           string
	City              string
	Website           string
	BusinessSummary   string
	FullTimeEmployees int

	// Market Metadata
	MarketCap     int64
	Beta          decimal.Decimal `gorm:"type:numeric(6,3)"`
	TrailingPE    decimal.Decimal `gorm:"type:numeric(10,3)"`
	ForwardPE     decimal.Decimal `gorm:"type:numeric(10,3)"`
	DividendYield decimal.Decimal `gorm:"type:numeric(10,5)"`

	// Holders
	InsidersPct     decimal.Decimal `gorm:"type:numeric(6,3)"`
	InstitutionsPct decimal.Decimal `gorm:"type:numeric(6,3)"`
	FloatPct        decimal.Decimal `gorm:"type:numeric(6,3)"`

	LastUpdated time.Time `gorm:"autoUpdateTime"`

	// --- Relationships ---
	PriceHistory []StockPriceHistory `gorm:"foreignKey:StockID"`
	Dividends    []StockDividends    `gorm:"foreignKey:StockID"`
	Splits       []StockSplits       `gorm:"foreignKey:StockID"`
	
	// Financials
	IncomeStatements []IncomeStatement `gorm:"foreignKey:StockID"`
	BalanceSheets    []BalanceSheet    `gorm:"foreignKey:StockID"`
	CashFlows        []CashFlow        `gorm:"foreignKey:StockID"`

	News []News `gorm:"many2many:news_stock_map;"`
}

// ---------------------------------------------------------
// 2. Historical Data
// ---------------------------------------------------------
type StockPriceHistory struct {
	StockID   uint      `gorm:"primaryKey;autoIncrement:false"`
	TradeDate time.Time `gorm:"primaryKey;type:date"`

	OpenPrice  decimal.Decimal `gorm:"type:numeric(14,4)"`
	HighPrice  decimal.Decimal `gorm:"type:numeric(14,4)"`
	LowPrice   decimal.Decimal `gorm:"type:numeric(14,4)"`
	ClosePrice decimal.Decimal `gorm:"type:numeric(14,4)"`
	AdjClose   decimal.Decimal `gorm:"type:numeric(14,4)"`
	Volume     int64
}

type StockDividends struct {
	StockID uint      `gorm:"primaryKey;autoIncrement:false"`
	ExDate  time.Time `gorm:"primaryKey;type:date"`
	Dividend decimal.Decimal `gorm:"type:numeric(10,4)"`
}

type StockSplits struct {
	StockID    uint      `gorm:"primaryKey;autoIncrement:false"`
	SplitDate  time.Time `gorm:"primaryKey;type:date"`
	SplitRatio decimal.Decimal `gorm:"type:numeric(10,4)"`
}

// ---------------------------------------------------------
// 3. Financial Statements
// ---------------------------------------------------------
type IncomeStatement struct {
	StockID       uint      `gorm:"primaryKey;autoIncrement:false"`
	PeriodEndDate time.Time `gorm:"primaryKey;type:date"`
	PeriodType    string    `gorm:"primaryKey;size:10"` // quarterly, annual, ttm

	TotalRevenue    int64
	CostOfRevenue   int64
	GrossProfit     int64
	OperatingIncome int64
	NetIncome       int64
	EPS             decimal.Decimal `gorm:"type:numeric(10,4)"`
}

type BalanceSheet struct {
	StockID       uint      `gorm:"primaryKey;autoIncrement:false"`
	PeriodEndDate time.Time `gorm:"primaryKey;type:date"`
	PeriodType    string    `gorm:"primaryKey;size:10"`

	TotalAssets      int64
	TotalLiabilities int64
	TotalEquity      int64
	Cash             int64
	Debt             int64
}

type CashFlow struct {
	StockID       uint      `gorm:"primaryKey;autoIncrement:false"`
	PeriodEndDate time.Time `gorm:"primaryKey;type:date"`
	PeriodType    string    `gorm:"primaryKey;size:10"`

	OperatingCashFlow int64
	InvestingCashFlow int64
	FinancingCashFlow int64
	FreeCashFlow      int64
}

// ---------------------------------------------------------
// 4. Analysis & News
// ---------------------------------------------------------
type StockAnalystRatings struct {
	StockID    uint      `gorm:"primaryKey;autoIncrement:false"`
	RatingDate time.Time `gorm:"primaryKey;type:date"`
	Firm       string    `gorm:"primaryKey"` 

	FromGrade string
	ToGrade   string
	Action    string
}

type StockPriceTargets struct {
	StockID    uint      `gorm:"primaryKey;autoIncrement:false"`
	TargetDate time.Time `gorm:"primaryKey;type:date"`

	LowTarget    decimal.Decimal `gorm:"type:numeric(14,4)"`
	MeanTarget   decimal.Decimal `gorm:"type:numeric(14,4)"`
	HighTarget   decimal.Decimal `gorm:"type:numeric(14,4)"`
	AnalystCount int
}

type News struct {
	ID          uint64 `gorm:"primaryKey"`
	PublishedAt time.Time
	Title       string
	Publisher   string
	NewsType    string `gorm:"size:20"`
	URL         string
	Summary     string
}