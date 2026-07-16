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

// Create godoc
//
// @Summary Criar personagem
// @Description Cria um novo personagem para o usuário autenticado.
// @Tags Characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CharacterCreateRequest true "Dados do personagem"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /characters [post]
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

// GetAll godoc
//
// @Summary Listar personagens
// @Description Retorna todos os personagens do usuário autenticado.
// @Tags Characters
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /characters [get]
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

// GetByID godoc
//
// @Summary Buscar personagem
// @Description Retorna um personagem específico do usuário autenticado.
// @Tags Characters
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID do personagem"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /characters/{id} [get]
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

// Update godoc
//
// @Summary Atualizar personagem
// @Description Atualiza os dados de um personagem.
// @Tags Characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID do personagem"
// @Param request body dto.CharacterUpdateRequest true "Novos dados do personagem"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /characters/{id} [put]
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

// Delete godoc
//
// @Summary Deletar personagem
// @Description Remove um personagem do usuário autenticado.
// @Tags Characters
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID do personagem"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /characters/{id} [delete]
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

// AddExperience godoc
//
// @Summary Adicionar experiência
// @Description Adiciona experiência ao personagem e realiza level up automaticamente.
// @Tags Characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID do personagem"
// @Param request body dto.AddExperienceRequest true "Experiência"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /characters/{id}/experience [post]
func (h *CharacterHandler) AddExperience(c *gin.Context) {
	id := c.Param("id")

	var req dto.AddExperienceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Dados Invalidos",
		})
		return
	}

	if err := h.characterService.AddExperience(id, req.Experience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Experiencia adicionada com sucesso",
	})
}

// SpendAttributePoints godoc
//
// @Summary Distribuir pontos de atributo
// @Description Gasta pontos de atributo obtidos ao subir de nível.
// @Tags Characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID do personagem"
// @Param request body dto.SpendAttributePointsRequest true "Distribuição dos pontos"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /characters/{id}/attributes [post]
func (h *CharacterHandler) SpendAttributePoints(c *gin.Context) {
	id := c.Param("id")

	var req dto.SpendAttributePointsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Dados invalidos",
		})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Distribuição de atributos invalida",
		})
		return
	}

	if err := h.characterService.SpendAttributePoints(id, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Pontos de atributo distribuídos com sucesso",
	})
}
