package handler

import (
	"net/http"

	"github.com/caioLeone/go-arena-api/internal/dto"
	"github.com/caioLeone/go-arena-api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CharacterHandler struct {
	characterService service.CharacterService
	validator        *validator.Validate
}

func NewCharacterHandler(characterService service.CharacterService) *CharacterHandler {
	return &CharacterHandler{
		characterService: characterService,
		validator:        validator.New(),
	}
}

// CREATE CRIA NOVO PERSONAGEM
// POST /characters
func (h *CharacterHandler) Create(c *gin.Context) {
	var req dto.CharacterCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "dados invalidos",
		})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Validacao Falhou: naome (3-100), Classe Obrigatoria",
		})
		return
	}

	userID := c.GetString("user_id")
	character, err := h.characterService.Create(userID, &req)
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    character,
	})
}

// GET ALL retorna todos os personagens do usuario
// GET /characters
func (h *CharacterHandler) GetAll(c *gin.Context) {
	userID := c.GetString("user_id")
	characters, err := h.characterService.GetAll(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    characters,
	})
}

// GET BY ID retorna um personagem especifico do usuario
// GET /characters/:id
func (h *CharacterHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("user_id")

	character, err := h.characterService.GetByID(id, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    character,
	})
}

// Update atualiza um personagem
// PUT /characters/:id
func (h *CharacterHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req dto.CharacterUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Dados Invalidos",
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

	userID := c.GetString("user_id")

	character, err := h.characterService.Update(id, userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    character,
	})
}

// Delete deleta um personagem
// DELETE /characters/:id
func (h *CharacterHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("user_id")

	err := h.characterService.Delete(id, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Personagem Deletado Com Sucesso",
	})
}
