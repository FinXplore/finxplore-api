package cron

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

// Scheduler holds the config and shared dependencies
type Scheduler struct {
	WorkerURL    string
	WorkerAPIKey string
	Logger       *zap.Logger
}

// NewScheduler creates the instance
func NewScheduler(url, apiKey string, logger *zap.Logger) *Scheduler {
	return &Scheduler{
		WorkerURL:    url,
		WorkerAPIKey: apiKey,
		Logger:       logger,
	}
}

// Start registers all jobs and starts the timer
func (s *Scheduler) Start() {
	c := cron.New()

	// --- Job 1: Market Data Sync ---
	// Schedule: Every 12 hours
	// Note: Use "@every 1m" for testing
	_, err := c.AddFunc("@every 12h", func() {
		s.TriggerMarketSync() // <--- Calls the method in the other file
	})

	if err != nil {
		s.Logger.Error("❌ Failed to register Market Sync job", zap.Error(err))
	} else {
		s.Logger.Info("✅ Registered Job: Market Data Sync (Every 12h)")
	}

	// Future jobs can go here...
	// c.AddFunc("@midnight", s.TriggerDividendPayout)

	c.Start()
	s.Logger.Info("⏳ Scheduler Started")
}