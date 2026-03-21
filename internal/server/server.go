package server

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/Dhyey3187/finxplore-api/api/routes"
	"github.com/Dhyey3187/finxplore-api/internal/config"
	"github.com/Dhyey3187/finxplore-api/internal/cron"
)

type Server struct {
	cfg    *config.Config
	logger *zap.Logger
	router *gin.Engine
	db     *gorm.DB
	redis  *redis.Client
	routes *routes.Routes
}

func NewServer(
	cfg *config.Config,
	logger *zap.Logger,
	db *gorm.DB,
	rdb *redis.Client,
	routes *routes.Routes,
) *Server {

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	return &Server{
		cfg:    cfg,
		db:     db,
		redis:  rdb,
		router: router,
		logger: logger,
		routes: routes,
	}
}

func (s *Server) Run() error {
	// 🔑 Routes are registered here
	scheduler := cron.NewScheduler(s.cfg.DataWorkerURL, s.cfg.DataWorkerApiKey, s.logger)
	scheduler.Start()

	//Cors settings
	s.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5000", "http://192.168.1.31:5000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	s.routes.Register(s.router)

	s.logger.Info("🚀 Server starting",
		zap.Int("port", s.cfg.ServerPort),
		zap.String("env", s.cfg.AppEnv),
	)

	addr := fmt.Sprintf(":%d", s.cfg.ServerPort)
	return s.router.Run(addr)
}
