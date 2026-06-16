package middleware

import (
	"net/http"

	"github.com/caioLeone/go-arena-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Error("Panic recuperado: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"error":   "Erro interno do servidor",
				})
				c.Abort()
			}
		}()

		c.Next()

		//Verificar se houve erro na requisição
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			log.Error("Erro na requisicao: %v", err)

			statusCode := http.StatusInternalServerError
			if c.Writer.Status() != http.StatusOK {
				statusCode = c.Writer.Status()
			}

			c.JSON(statusCode, gin.H{
				"success": false,
				"error":   err.Error(),
			})
		}
	}
}

// Valida erro e retorna resposta apropriada
func ValidateError(c *gin.Context, err error, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"success": false,
		"error":   message,
		"details": err.Error(),
	})
}
