//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/Dhyey3187/finxplore-api/internal/config"
	"github.com/Dhyey3187/finxplore-api/internal/database"
	"github.com/Dhyey3187/finxplore-api/internal/logger"
	"github.com/Dhyey3187/finxplore-api/internal/server"
	"github.com/Dhyey3187/finxplore-api/api/handler"
	"github.com/Dhyey3187/finxplore-api/api/repository"
	"github.com/Dhyey3187/finxplore-api/api/service"
)

// InitializeApp is the blueprint for the Wire tool.
func InitializeApp() (*server.Server, error) {
	wire.Build(
		config.LoadConfig,        // Returns *Config
		logger.NewLogger,         // Returns *Logger, error
		database.ConnectPostgres, // Returns *gorm.DB, error
		database.ConnectRedis,    // Returns *redis.Client, error
		service.NewUserService,
		repository.NewUserRepository,
		repository.NewCacheRepository,
		handler.NewAuthHandler,
		server.NewServer,         // Returns *Server
	)
	return &server.Server{}, nil
}