package middleware

import (
	"net/http"
	"strings"

	"github.com/Dhyey3187/finxplore-api/internal/config"
	"github.com/Dhyey3187/finxplore-api/internal/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware creates the Gin middleware
func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Get the Authorization Header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		// 2. Check Format ("Bearer <token>")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}
		tokenString := parts[1]

		// 3. Verify Token
		claims, err := utils.VerifyToken(tokenString, cfg.JWTSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// 4. Attach User Info to Context (So Handlers can use it)
		c.Set("user_code", claims.UserCode)
		c.Set("role", claims.Role)

		// 5. Allow request to proceed
		c.Next()
	}
}