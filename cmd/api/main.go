package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/caioLeone/go-arena-api/internal/config"
	"github.com/caioLeone/go-arena-api/internal/handler"
	"github.com/caioLeone/go-arena-api/internal/middleware"
	"github.com/caioLeone/go-arena-api/internal/ranking"
	"github.com/caioLeone/go-arena-api/internal/repository"
	"github.com/caioLeone/go-arena-api/internal/service"
	"github.com/caioLeone/go-arena-api/pkg/database"
	"github.com/caioLeone/go-arena-api/pkg/logger"
	"github.com/caioLeone/go-arena-api/pkg/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	//0. Criar Logger
	log := logger.NewLogger(logger.INFO)
	log.Info("Iniciando Arena API")

	//1. Carregar Configuracao
	cfg := config.Load()
	log.Info("Configuracoes Carregadas (env: %s)", cfg.ServerEnv)

	//2. Conectar PostgreSQL
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Erro ao conectar PostgreSQL: %v", err)
	}
	defer db.Close()

	//3. Rodar Migrations
	if err := database.RunMigrations(db, "migrations/"); err != nil {
		log.Fatal("Erro ao rodar migrations: %v", err)
	}
	log.Info("Migrations Executadas Com Sucesso")

	//4. Conectar Redis
	redisClient := redis.Connect(cfg)
	defer redisClient.Close()
	log.Info("Conectado Ao Redis Com Sucesso")

	//5. Setup Gin
	if cfg.ServerEnv == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	//6. Aplicar Middlewares Globais
	router.Use(middleware.CORSMiddleware(cfg))
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.ErrorHandlingMiddleware(log))

	//Rate Limiting: 100 Requisicoes por minuto
	rateLimiter := middleware.NewRateLimiter(100, 1*time.Minute)
	router.Use(rateLimiter.Middleware())

	router.Use(gin.Recovery())

	//7. Inicializar dependencias
	initializeDependencies(router, db, cfg, redisClient, log)

	//8. Iniciar Servidor
	log.Info("Arene API Iniciada na porta %s (env: %s)", cfg.ServerPort, cfg.ServerEnv)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Erro ao iniciar servidor: %v", err)
	}
}

func initializeDependencies(router *gin.Engine, db *sql.DB, cfg *config.Config, redisClient *redis.Client, log *logger.Logger) {
	//Health Check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "arena-api",
		})
	})

	//Repositories
	userRepo := repository.NewUserRepository(db)
	characterRepo := repository.NewCharacterRepository(db)
	battleRepo := repository.NewBattleRepository(db)

	//Services
	authService := service.NewAuthService(userRepo, cfg)
	characterService := service.NewCharacterService(characterRepo)
	leaderboardService := ranking.NewLeaderboardService(redisClient)
	battleService := service.NewBattleService(battleRepo, characterRepo, leaderboardService)

	//Handlers
	authHandler := handler.NewAuthHandler(authService)
	characterHandler := handler.NewCharacterHandler(characterService)
	battleHandler := handler.NewBattleHandler(battleService)
	rankingHandler := handler.NewRankingHandler(leaderboardService, characterRepo)

	//Routes - Auth (publicas)
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.Refresh)
	}

	//Routes - Protegidas (Exemplo)
	protected := router.Group("/api")
	protected.Use(middleware.JWTMiddleware(cfg))
	{
		protected.GET("/me", func(c *gin.Context) {
			userID := c.GetString("user_id")
			c.JSON(http.StatusOK, gin.H{
				"user_id": userID,
				"message": "Voce esta autenticado",
			})
		})
	}

	//Routes - Characters
	characters := router.Group("/characters")
	characters.Use(middleware.JWTMiddleware(cfg))
	{
		characters.POST("", characterHandler.Create)
		characters.GET("", characterHandler.GetAll)
		characters.GET("/:id", characterHandler.GetByID)
		characters.PUT("/:id", characterHandler.Update)
		characters.DELETE("/:id", characterHandler.Delete)
	}

	//Routes - Battles
	battles := router.Group("/battles")
	battles.Use(middleware.JWTMiddleware(cfg))
	{
		battles.POST("", battleHandler.StartBattle)
		battles.GET("/history", battleHandler.GetHistory)
	}

	//Routes - Ranking (públicas)
	rankingRoutes := router.Group("/ranking")
	{
		rankingRoutes.GET("", middleware.JWTMiddleware(cfg), rankingHandler.GetUserRanking)
		rankingRoutes.GET("/top", rankingHandler.GetTopPlayers) // Pública
	}

	log.Info("Todas As Rotas Registradas Com Sucesso")
}