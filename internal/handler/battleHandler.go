package handler

import (
	"net/http"
	"strconv"

	"github.com/caioLeone/go-arena-api/internal/dto"
	"github.com/caioLeone/go-arena-api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BattleHandler struct {
	battleService service.BattleService
	validator     *validator.Validate
}

func NewBattleHandler(battleService service.BattleService) *BattleHandler {
	return &BattleHandler{
		battleService: battleService,
		validator:     validator.New(),
	}
}

// StartBattle godoc
//
// @Summary Iniciar batalha
// @Description Inicia uma batalha entre dois personagens e salva o resultado.
// @Tags Battles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.BattleCreateRequest true "Dados da batalha"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /battles [post]
func (h *BattleHandler) StartBattle(c *gin.Context) {
	var req dto.BattleCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "dados inválidos",
		})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "IDs de personagens inválidos",
		})
		return
	}

	userID := c.GetString("user_id")

	battle, err := h.battleService.StartBattle(userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    battle,
	})
}

// GetHistory godoc
//
// @Summary Histórico de batalhas
// @Description Retorna o histórico de batalhas do usuário autenticado.
// @Tags Battles
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Quantidade de registros" default(10)
// @Param offset query int false "Offset da paginação" default(0)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /battles/history [get]
func (h *BattleHandler) GetHistory(c *gin.Context) {
	userID := c.GetString("user_id")

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	history, err := h.battleService.GetBattleHistory(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    history,
	})
}
