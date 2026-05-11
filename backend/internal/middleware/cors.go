package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	allowedOriginsStr := os.Getenv("CORS_ALLOWED_ORIGINS")
	frontendURL := os.Getenv("FRONTEND_URL")

	var origins []string
	if allowedOriginsStr != "" {
		origins = strings.Split(allowedOriginsStr, ",")
	}
	if frontendURL != "" {
		origins = append(origins, frontendURL)
	}

	// Clean up origins: trim spaces, remove empty strings
	var finalOrigins []string
	originMap := make(map[string]bool)
	for _, o := range origins {
		trimmed := strings.TrimSpace(o)
		if trimmed != "" && !originMap[trimmed] {
			finalOrigins = append(finalOrigins, trimmed)
			originMap[trimmed] = true
		}
	}

	// Fallback to localhost if no origins configured
	if len(finalOrigins) == 0 {
		finalOrigins = []string{"http://localhost:3000"}
	}

	return cors.New(cors.Config{
		AllowOrigins:     finalOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
