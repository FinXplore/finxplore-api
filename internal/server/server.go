package server
import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/Dhyey3187/finxplore-api/internal/config"
)
type Server struct {
	cfg    *config.Config
	logger *zap.Logger // Add this
	router *gin.Engine
	db     *gorm.DB
	redis  *redis.Client
}

// Update constructor to accept logger
func NewServer(cfg *config.Config, logger *zap.Logger, db *gorm.DB, rdb *redis.Client) *Server {
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	return &Server{
		cfg:    cfg,
		logger: logger, // Store it
		router: gin.Default(),
		db:     db,
		redis:  rdb,
	}
}

func (s *Server) Run() error {
	// Call the method to setup routes
	s.RegisterRoutes()

	s.logger.Info("🚀 Server starting",
		zap.String("port", s.cfg.ServerPort),
		zap.String("env", s.cfg.AppEnv),
	)

	return s.router.Run(":" + s.cfg.ServerPort)
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
}