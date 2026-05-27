package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// CONTEM TODAS AS CONFIGURAÇOES DA APLICACAO
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	//REDIS
	RedisHost string
	RedisPort string

	//SERVER
	ServerPort string
	ServerEnv  string

	//JWT
	JWTSecret                string
	JWTAccessExpirationMin   int
	JWTRefreshExpirationDays int

	//CORS
	CORSAllowedOrigins string
}

// CARREGA AS CONFIGURACOES DO .ENV
func Load() *Config {
	// Carrega .env em ambiente de desenvolvimento
	// Em produção, variáveis devem estar no sistema

	_ = godotenv.Load()

	cfg := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "arena_user"),
		DBPassword: getEnv("DB_PASSWORD", "arena_password"),
		DBName:     getEnv("DB_NAME", "arena_db"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),

		//REDIS
		RedisHost: getEnv("REDIS_HOST", "localhost"),
		RedisPort: getEnv("REDIS_HOST", "6379"),

		//SERVER
		ServerPort: getEnv("SERVER_PORT", "8080"),
		ServerEnv:  getEnv("SERVER_ENV", "dev"),

		//JWT
		JWTSecret:                getEnv("JWT_SERVER", "dev-secret-key"),
		JWTAccessExpirationMin:   getEnvInt("JWT_ACCESS_EXPIRATION_MINUTES", 15),
		JWTRefreshExpirationDays: getEnvInt("JWT_REFRESH_EXPIRATION_DAYS", 7),

		//CORS
		CORSAllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000,http://localhost:5173"),
	}

	if cfg.JWTSecret == "dev-secret-key" {
		fmt.Println("AVISO: JWT_SECRET está usando valor padrão de desenvolvimento. NUNCA use em produção!")
	}
	return cfg
}

// GETENV RETORNA VARIAVEL DE AMBIENTE OU VALOR PADRAO
func getEnv(key, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if value, ok := os.LookupEnv(key); ok {
		var intVal int
		_, err := fmt.Sscanf(value, "%d", &intVal)
		if err == nil {
			return intVal
		}
	}

	return defaultVal
}
