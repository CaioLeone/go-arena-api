package handler

import (
	"net/http"
	"strconv"

	"github.com/caioLeone/go-arena-api/internal/dto"
	"github.com/caioLeone/go-arena-api/internal/ranking"
	"github.com/caioLeone/go-arena-api/internal/repository"
	"github.com/gin-gonic/gin"
)

type RankingHandler struct {
	leaderboardService *ranking.LeaderboardService
	characterRepo      repository.CharacterRepository
}

func NewRankingHandler(leaderboardService *ranking.LeaderboardService, characterRepo repository.CharacterRepository) *RankingHandler {
	return &RankingHandler{
		leaderboardService: leaderboardService,
		characterRepo:      characterRepo,
	}
}

// GetUserRanking godoc
//
// @Summary Ranking do usuário
// @Description Retorna o ranking de todos os personagens pertencentes ao usuário autenticado.
// @Tags Ranking
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /ranking [get]
func (h *RankingHandler) GetUserRanking(c *gin.Context) {
	userID := c.GetString("user_id")

	//Buscar Personagens do Usuario
	characters, err := h.characterRepo.GetAllByUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	rankings := make([]dto.UserRankingResponse, 0, len(characters))
	for _, char := range characters {

		rank, err :=
			h.leaderboardService.GetPlayerRank(
				char.ID.String(),
				char.Name,
			)

		if err != nil {
			rank = -1
		}

		score, err :=
			h.leaderboardService.GetPlayerScore(
				char.ID.String(),
				char.Name,
			)

		if err != nil {
			score = 0
		}

		rankings = append(
			rankings,
			dto.UserRankingResponse{
				CharacterID: char.ID,
				Name:        char.Name,
				Class:       char.Class,
				Level:       char.Level,
				Rank:        rank,
				Score:       int64(score),
			},
		)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rankings,
	})
}

// GetTopPlayers godoc
//
// @Summary Top Ranking
// @Description Retorna os jogadores com maior pontuação no leaderboard.
// @Tags Ranking
// @Produce json
// @Param limit query int false "Quantidade de jogadores (máximo 100)" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /ranking/top [get]
func (h *RankingHandler) GetTopPlayers(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil || limit <= 0 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	rankingPlayers, err := h.leaderboardService.GetTopPlayers(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	players := make(
		[]dto.TopPlayersResponse,
		0,
		len(rankingPlayers),
	)

	for _, rankingPlayer := range rankingPlayers {
		character, err := h.characterRepo.GetByIDNoUserFilter(
			rankingPlayer.CharacterID,
		)

		if err != nil {
			continue
		}

		players = append(
			players,
			dto.TopPlayersResponse{
				Rank:        rankingPlayer.Rank,
				CharacterID: character.ID,
				Name:        character.Name,
				Class:       character.Class,
				Level:       character.Level,
				Score:       rankingPlayer.Score,
			},
		)
	}

	response := dto.LeaderboardResponse{
		Players: players,
		Total:   len(players),
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}
