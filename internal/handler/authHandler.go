package handler

import (
	"net/http"

	"github.com/caioLeone/go-arena-api/internal/dto"
	"github.com/caioLeone/go-arena-api/internal/middleware"
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

// Register godoc
//
// @Summary Registrar usuário
// @Description Cria um novo usuário e retorna os tokens JWT.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.UserCreateRequest true "Dados do usuário"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.UserCreateRequest

	//PARSE JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			//"error":   "dados invalidos",
			"error": err.Error(),
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

// Login godoc
//
// @Summary Login
// @Description Autentica um usuário.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.UserLoginRequest true "Credenciais"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.UserLoginRequest

	//Parse JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			//"error":   "dados invalidos",
			"error": err.Error(),
		})
		return
	}

	//Validar
	erros := middleware.ValidRequest(h.validator, req)
	if len(erros) > 0 {
		middleware.ValidationErrorResponse(c, erros)
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

// Refresh godoc
//
// @Summary Renovar Access Token
// @Description Gera um novo Access Token usando Refresh Token.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest true "Refresh Token"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /auth/refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {
	var req dto.RefreshTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "refresh_token obrigatorio",
		})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Validacao Falhou",
		})
		return
	}

	accessToken, err := h.authService.Refresh(req.RefreshToken)
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
			"access_token": accessToken,
		},
	})
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID := c.GetString("user_id")

	user, err := h.authService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
