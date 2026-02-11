package cron

import (
	"bytes"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// TriggerMarketSync handles the HTTP call to the Python worker
// Since it's in the same package, it can be a method of Scheduler
func (s *Scheduler) TriggerMarketSync() {
	s.Logger.Info("⏰ [Cron] Triggering Market Sync...")

	// 1. Construct URL
	targetURL := s.WorkerURL + "/sync"

	// 2. Create Request
	req, err := http.NewRequest("POST", targetURL, bytes.NewBuffer([]byte("{}")))
	if err != nil {
		s.Logger.Error("❌ [Cron] Failed to create request", zap.Error(err))
		return
	}

	// 3. Add Headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.WorkerAPIKey)

	// 4. Send Request with Timeout
	client := &http.Client{Timeout: 30 * time.Second} // Longer timeout for sync jobs
	resp, err := client.Do(req)

	if err != nil {
		s.Logger.Error("❌ [Cron] Connection failed", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	// 5. Log Result
	if resp.StatusCode == 200 {
		s.Logger.Info("✅ [Cron] Market Sync Triggered Successfully")
	} else {
		s.Logger.Warn("⚠️ [Cron] Market Sync failed", 
			zap.Int("status", resp.StatusCode),
			zap.String("target", targetURL),
		)
	}
}