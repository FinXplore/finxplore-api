package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/Dhyey3187/finxplore-api/api/handler"
	"github.com/Dhyey3187/finxplore-api/internal/config"
)
type Server struct {
	cfg    *config.Config
	logger *zap.Logger // Add this
	router *gin.Engine
	db     *gorm.DB
	redis  *redis.Client
	authHandler *handler.AuthHandler
}

func NewServer(cfg *config.Config,logger *zap.Logger, db *gorm.DB, rdb *redis.Client, authHandler *handler.AuthHandler) *Server {

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(gin.Recovery()) 
	router.Use(gin.Logger())
	// router.Use(middleware.CorsMiddleware()) // We will add this later

	s := &Server{
		cfg:    cfg,
		db:     db,
		redis:  rdb,
		router: router,
		logger: logger,
		authHandler: authHandler,
	}

	return s
}

func (s *Server) Run() error {
	s.RegisterRoutes()

	s.logger.Info("🚀 Server starting",
		zap.Int("port", s.cfg.ServerPort),
		zap.String("env", s.cfg.AppEnv),
	)

	addr := fmt.Sprintf(":%d", s.cfg.ServerPort)
	return s.router.Run(addr)
}

func (s *Server) RegisterRoutes() {
	// Health Check
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "up",
			"env":    s.cfg.AppEnv,
			"db":     "likely connected",
		})
	})

	api := s.router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			// Use the injected handler
			auth.POST("/register", s.authHandler.Register)
			auth.POST("/login", s.authHandler.Login)
		}
	}
}