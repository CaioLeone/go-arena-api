package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	port := getEnv("SERVER_PORT", "8080")
	env := getEnv("SERVER_ENV", "dev")

	// Setup Gin mode based on environment
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize router
	router := gin.Default()

	// Middleware
	// TODO: Add CORS, JWT, Rate Limiting middleware in Fase 2

	// Routes
	setupRoutes(router)

	// Start server
	log.Printf("Arena API starting on port %s (env: %s)\n", port, env)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// setupRoutes configures all API routes
func setupRoutes(router *gin.Engine) {
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "arena-api",
		})
	})

	// TODO: Add Auth routes (Fase 2)
	// router.POST("/auth/register", authHandler.Register)
	// router.POST("/auth/login", authHandler.Login)

	// TODO: Add Character routes (Fase 3)
	// router.POST("/characters", characterHandler.Create)
	// router.GET("/characters", characterHandler.List)
	// router.GET("/characters/:id", characterHandler.GetByID)
	// router.PUT("/characters/:id", characterHandler.Update)
	// router.DELETE("/characters/:id", characterHandler.Delete)

	// TODO: Add Battle routes (Fase 4)
	// router.POST("/battles", battleHandler.Create)
	// router.GET("/battles/history", battleHandler.GetHistory)

	// TODO: Add Ranking routes (Fase 5)
	// router.GET("/ranking", rankingHandler.GetUserRanking)
	// router.GET("/ranking/top", rankingHandler.GetTopPlayers)
}

// getEnv retrieves environment variable or returns default value
func getEnv(key, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultVal
}
