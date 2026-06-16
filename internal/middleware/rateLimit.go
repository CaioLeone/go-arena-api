package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimitStore struct {
	requests map[string][]time.Time
	mu       sync.RWMutex
}

type RateLimiter struct {
	store       *RateLimitStore
	maxRequests int
	window      time.Duration
}

func NewRateLimiter(maxRequests int, window time.Duration) *RateLimiter {
	limiter := &RateLimiter{
		store: &RateLimitStore{
			requests: make(map[string][]time.Time),
		},
		maxRequests: maxRequests,
		window:      window,
	}
	//Limpar registros expirados periodicamente
	go limiter.cleanup()
	return limiter
}

func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		rl.store.mu.Lock()
		now := time.Now()
		for ip, times := range rl.store.requests {
			//remover timestamps expirados
			validTimes := []time.Time{}
			for _, t := range times {
				if now.Sub(t) < rl.window {
					validTimes = append(validTimes, t)
				}
			}
			if len(validTimes) == 0 {
				delete(rl.store.requests, ip)
			} else {
				rl.store.requests[ip] = validTimes
			}
		}
		rl.store.mu.Unlock()
	}
}

// Verifica se o IP pode fazer a requisição
func (rl *RateLimiter) isAllowed(ip string) bool {
	rl.store.mu.Lock()
	defer rl.store.mu.Unlock()

	now := time.Now()
	requests := rl.store.requests[ip]

	//remover timestamps expirados
	validRequests := []time.Time{}
	for _, t := range requests {
		if now.Sub(t) < rl.window {
			validRequests = append(validRequests, t)
		}
	}

	//Verifica se pode fazer nova requisição
	if len(validRequests) < rl.maxRequests {
		validRequests = append(validRequests, now)
		rl.store.requests[ip] = validRequests
		return true
	}

	rl.store.requests[ip] = validRequests
	return false
}

// Middleware retorna o middleware de rate limiting do gin
func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		if !rl.isAllowed(ip) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"error":   "Rate Limit Excedido. Tente Novamente Mais Tarde",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
