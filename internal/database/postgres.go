package database

import (
	"fmt"

	"github.com/Dhyey3187/finxplore-api/api/models"
	"github.com/Dhyey3187/finxplore-api/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(cfg *config.Config, zapLogger *zap.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Kolkata",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	zapLogger.Info("Running Database Migrations...")
	err = db.AutoMigrate(
		&models.User{}, 
		// &models.Wallet{},      // Add future models here
		// &models.Transaction{}, // Add future models here
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	zapLogger.Info("✅ Database Migration Complete")

	return db, nil
}