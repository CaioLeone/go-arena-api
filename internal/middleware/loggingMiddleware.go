package middleware

import (
	"time"

	"github.com/caioLeone/go-arena-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

func LoggingMiddlewareWithLogger(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		method := c.Request.Method
		path := c.Request.RequestURI
		ip := c.ClientIP()

		//processar requisicao
		c.Next()

		//Informacoes de respostas
		statusCode := c.Writer.Status()
		duration := time.Since(startTime)

		//log formatado
		log.Info(
			"[%s] %s %s - Status: %d - Duracao: %v - IP: %s",
			method,
			path,
			c.Request.Proto,
			statusCode,
			duration,
			ip,
		)

		if statusCode >= 400 {
			log.Warm()
		}
	}
}
