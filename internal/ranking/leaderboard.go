package ranking

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	redisclient "github.com/caioLeone/go-arena-api/pkg/redis"
	goredis "github.com/redis/go-redis/v9"
)

const LeaderboardKey = "leaderboard"

type LeaderboardService struct {
	redisClient *redisclient.Client
}

func NewLeaderboardService(redisClient *redisclient.Client) *LeaderboardService {
	return &LeaderboardService{
		redisClient: redisClient,
	}
}

type PlayerRanking struct {
	Rank  int64  `json:"rank"`
	Name  string `json:"name"`
	Score int64  `json:"score"`
}

// UpdatePlayerScore atualiza score de uma jogador no leaderboard
func (ls *LeaderboardService) UpdatePlayerScore(characterID string, characterName string, score float64) error {
	ctx := context.Background()

	memberKey := fmt.Sprintf("%s:%s", characterID, characterName)

	err := ls.redisClient.ZAdd(ctx, LeaderboardKey, memberKey, score)
	if err != nil {
		log.Printf("Erro ao atualizar leaderboard: %v", err)
		return fmt.Errorf("Erro ao atualizar leaderboard: %w", err)
	}
	return nil
}

// GetPlayerRank retorna o rank de um jogador
func (ls *LeaderboardService) GetPlayerRank(characterID string, characterName string) (int64, error) {
	ctx := context.Background()

	memberKey := fmt.Sprintf("%s:%s", characterID, characterName)

	rawClient := ls.redisClient.GetRawClient()
	rank, err := rawClient.ZRevRank(ctx, LeaderboardKey, memberKey).Result()
	if err != nil {
		if err == goredis.Nil {
			return -1, nil // Jogador não encontrado
		}
		return -1, fmt.Errorf("Erro ao buscar rank: %w", err)
	}

	return rank + 1, nil // Redis retorna 0-indexed, convertemos para 1-indexed
}

func (ls *LeaderboardService) GetPlayerScore(characterID string, characterName string) (float64, error) {
	ctx := context.Background()

	memberKey := fmt.Sprintf("%s:%s", characterID, characterName)

	score, err := ls.redisClient.ZScore(ctx, LeaderboardKey, memberKey)
	if err != nil {
		if err == goredis.Nil {
			return 0, nil
		}
		return 0, fmt.Errorf("Erro ao buscar score: %w", err)
	}
	return score, nil
}

func (ls *LeaderboardService) GetTopPlayers(limit int64) ([]PlayerRanking, error) {
	ctx := context.Background()

	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	scores, err := ls.redisClient.ZRangeWithScores(ctx, LeaderboardKey, 0, limit-1, true)
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar top players: %w", err)
	}
	var players []PlayerRanking
	for i, z := range scores {
		memberName, ok := z.Member.(string)
		if !ok {
			log.Printf("Aviso: Membro do leaderboard nao e string: %v", z.Member)
			continue
		}
		player := PlayerRanking{
			Rank:  int64(i + 1),
			Score: int64(z.Score),
			Name:  memberName,
		}
		players = append(players, player)
	}
	return players, nil
}

func (ls *LeaderboardService) GetLeaderboardJSON(limit int64) (string, error) {
	players, err := ls.GetTopPlayers(limit)
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(players)
	if err != nil {
		return "", fmt.Errorf("Erro ao serializar leaderboard: %w", err)
	}
	return string(data), nil
}

func (ls *LeaderboardService) ClearLeaderboard() error {
	ctx := context.Background()
	return ls.redisClient.Del(ctx, LeaderboardKey)
}
