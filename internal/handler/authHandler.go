package handler

import (
	"net/http"

	"github.com/caioLeone/go-arena-api/internal/dto"
	"github.com/caioLeone/go-arena-api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AuthHandler struct para handles de autenticacao
type AuthHandler struct {
	authService service.AuthService
	validator   *validator.Validate
}

// NewAuthHandler cria novo handler de auth
func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   validator.New(),
	}
}

// Register Registra novo usuario
// POST /auth/register
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.UserCreateRequest

	//PARSE JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "dados invalidos",
		})
		return
	}

	//validar com go-playground/validator
	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Validacao falhou: email deve ser valido e senha minima 6 caracteres",
		})
		return
	}

	//Chamar Service
	userResp, loginResp, err := h.authService.Register(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": gin.H{
			"user":          userResp,
			"access_token":  loginResp.AccessToken,
			"refresh_token": loginResp.RefreshToken,
		},
	})
}

// Login faz login do Usuario
// POST /auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.UserLoginRequest

	//Parse JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "dados invalidos",
		})
		return
	}

	//Validar
	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Email e Senha Obrigatorios",
		})
		return
	}

	//Chamar Service
	loginResp, err := h.authService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"user":          loginResp.User,
			"access_token":  loginResp.AccessToken,
			"refresh_token": loginResp.RefreshToken,
		},
	})
}
