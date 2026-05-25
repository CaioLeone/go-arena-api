package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/caioLeone/go-arena-api/internal/config"
	"github.com/caioLeone/go-arena-api/pkg/database"
	"github.com/caioLeone/go-arena-api/pkg/redis"
)

func main() {
	//1. Carregar Configuracao
	cfg := config.Load()
	log.Printf("Configuracoes Carregadas (env: %s)", cfg.ServerEnv)

	//2. Conectar PostgreSQL
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Erro ao conectar PostgreSQL: %v", err)
	}
	defer db.Close()

	//3. Rodar Migrations
	if err := database.RunMigrations(db, "migrations/"); err != nil {
		log.Fatalf("Erro ao rodar migrations: %v", err)
	}

	//4. Conectar Redis
	redisClient := redis.Connect(cfg)
	defer redisClient.Close()

	//5. Setup Gin
	if cfg.ServerEnv == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	//6. Aplicar Middlewares Globais
	router.Use(middleware.CORSMiddleware(cfg))
	router.Use(middleware.LoggingMiddleware())
	router.Use(gin.Recovery())

	//7. Inicializar dependencias
	initializeDependencies(router, db, cfg)

	//8. Iniciar Servidor
	log.Printf("Arene API Iniciada na porta %s (env: %s)", cfg.ServerPort, cfg.ServerEnv)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}

func initializeDependencies(router gin.Engine, db *sql.DB, cfg *config.Config) {
	//Health Check
	router.GET("/health", func(ctx *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "arena-api",
		})
	})

	//Repositories
	userRepo := repository.NewUserRepository(db)

	//Services
	authService := service.NewAuthService(userRepo, cfg)

	//Handlers
	authHandler := handler.NewAuthHandler(authService)

	//Routes - Auth (publicas)
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	//Routes - Protegidas (Exemplo)
	protected := router.Group("/api")
	protected.Use(middleware.JWTMiddleware(cfg))
	{
		userID := c.GetString("user_id")
		c.JSON(http.StatusOK, gin.H{
			"user_id": userID,
			"message": "Voce esta autenticado",
		})
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
