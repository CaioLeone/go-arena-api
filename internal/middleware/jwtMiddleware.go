package middleware

import (
	"net/http"
	"strings"

	"github.com/caioLeone/go-arena-api/internal/auth"
	"github.com/caioLeone/go-arena-api/internal/config"
	"github.com/gin-gonic/gin"
)

// JWTMiddleware valida JWT em rotas protegidas
func JWTMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Extrair Token do Header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Token nao fornecido",
			})
			c.Abort()
			return
		}

		//Validar Formato "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Formato de Token Invalida",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		//validar Token
		claims, err := auth.ValidateToken(tokenString, cfg)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "token invalido ou expirado",
			})
			c.Abort()
			return
		}

		//Salvar user_id no context para usar nos handlers
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)

		c.Next()
	}
}

func CORSMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
