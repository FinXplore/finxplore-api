package service

import (
	"errors"
	"time"

	"github.com/Dhyey3187/finxplore-api/api/dto"
	"github.com/Dhyey3187/finxplore-api/api/models"
	"github.com/Dhyey3187/finxplore-api/api/repository"
	"github.com/shopspring/decimal"
)

type MarketService interface {
	GetStockDetails(ticker string) (*models.Stock, error)
	GetStockChart(ticker string, period string) ([]models.StockPriceHistory, error)
	SearchStocks(query string) ([]models.Stock, error)
	GetIndices() ([]dto.MarketIndexResponse, error)
	GetMarketMovers(exchange string, limit int, moverType string) ([]dto.MarketMoverResponse, error)
	GetStockQuote(ticker string) (*dto.StockQuoteResponse, error)
	GetStockNews(ticker string, limit int) ([]dto.StockNewsResponse, error)
	GetStockFinancials(ticker string, statementType string) ([]dto.FinancialStatementResponse, error)
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

func (s *marketService) GetIndices() ([]dto.MarketIndexResponse, error) {
	return []dto.MarketIndexResponse{
		{Name: "NIFTY 50", Value: decimal.NewFromFloat(22000.50), Change: decimal.NewFromFloat(150.20), ChangePct: decimal.NewFromFloat(0.68), LastUpdated: time.Now()},
		{Name: "SENSEX", Value: decimal.NewFromFloat(72500.30), Change: decimal.NewFromFloat(400.10), ChangePct: decimal.NewFromFloat(0.55), LastUpdated: time.Now()},
		{Name: "NASDAQ", Value: decimal.NewFromFloat(16000.00), Change: decimal.NewFromFloat(-50.00), ChangePct: decimal.NewFromFloat(-0.31), LastUpdated: time.Now()},
	}, nil
}

func (s *marketService) GetMarketMovers(exchange string, limit int, moverType string) ([]dto.MarketMoverResponse, error) {
	return []dto.MarketMoverResponse{
		{Ticker: "MOCK1", Name: "Mock Company 1", Price: decimal.NewFromFloat(100.0), Change: decimal.NewFromFloat(5.0), ChangePct: decimal.NewFromFloat(5.0), Volume: 100000},
	}, nil
}

func (s *marketService) GetStockQuote(ticker string) (*dto.StockQuoteResponse, error) {
	stock, err := s.stockRepo.GetByTicker(ticker)
	if err != nil {
		return nil, err
	}
	history, _ := s.stockRepo.GetHistory(stock.ID, time.Now().AddDate(0, 0, -5), time.Now())
	var price decimal.Decimal
	if len(history) > 0 {
		price = history[len(history)-1].ClosePrice
	}
	return &dto.StockQuoteResponse{
		Ticker:      stock.Ticker,
		Price:       price,
		Change:      decimal.Zero,
		ChangePct:   decimal.Zero,
		Volume:      0,
		LastUpdated: time.Now(),
	}, nil
}

func (s *marketService) GetStockNews(ticker string, limit int) ([]dto.StockNewsResponse, error) {
	stock, err := s.stockRepo.GetByTicker(ticker)
	if err != nil {
		return nil, err
	}
	news, err := s.stockRepo.GetNews(stock.ID, limit)
	if err != nil {
		return nil, err
	}
	var res []dto.StockNewsResponse
	for _, n := range news {
		res = append(res, dto.StockNewsResponse{
			Headline:    n.Title,
			Source:      n.Publisher,
			URL:         n.URL,
			PublishedAt: n.PublishedAt,
		})
	}
	return res, nil
}

func (s *marketService) GetStockFinancials(ticker string, statementType string) ([]dto.FinancialStatementResponse, error) {
	stock, err := s.stockRepo.GetByTicker(ticker)
	if err != nil {
		return nil, err
	}

	var res []dto.FinancialStatementResponse

	switch statementType {
	case "income":
		stmts, err := s.stockRepo.GetIncomeStatements(stock.ID)
		if err != nil {
			return nil, err
		}
		for _, st := range stmts {
			res = append(res, dto.FinancialStatementResponse{
				PeriodEndDate: st.PeriodEndDate.Format("2006-01-02"),
				Metrics: map[string]interface{}{
					"total_revenue":    st.TotalRevenue,
					"gross_profit":     st.GrossProfit,
					"operating_income": st.OperatingIncome,
					"net_income":       st.NetIncome,
				},
			})
		}
	case "balance":
		stmts, err := s.stockRepo.GetBalanceSheets(stock.ID)
		if err != nil {
			return nil, err
		}
		for _, st := range stmts {
			res = append(res, dto.FinancialStatementResponse{
				PeriodEndDate: st.PeriodEndDate.Format("2006-01-02"),
				Metrics: map[string]interface{}{
					"total_assets":      st.TotalAssets,
					"total_liabilities": st.TotalLiabilities,
					"total_equity":      st.TotalEquity,
					"cash":              st.Cash,
					"debt":              st.Debt,
				},
			})
		}
	case "cashflow":
		stmts, err := s.stockRepo.GetCashFlows(stock.ID)
		if err != nil {
			return nil, err
		}
		for _, st := range stmts {
			res = append(res, dto.FinancialStatementResponse{
				PeriodEndDate: st.PeriodEndDate.Format("2006-01-02"),
				Metrics: map[string]interface{}{
					"operating_cash_flow": st.OperatingCashFlow,
					"investing_cash_flow": st.InvestingCashFlow,
					"financing_cash_flow": st.FinancingCashFlow,
					"free_cash_flow":      st.FreeCashFlow,
				},
			})
		}
	default:
		return nil, errors.New("invalid financial statement type")
	}

	return res, nil
}