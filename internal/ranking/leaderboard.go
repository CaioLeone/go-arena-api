package ranking

import (
    "context"
    "encoding/json"
    "fmt"
    "log"

    "github.com/caioLeone/go-arena-api/pkg/redis"
    "github.com/redis/go-redis/v9"
)

const LeaderboardKey = "leaderboard"

type LeaderboardService struct{
	redisClient *redis.Client
}

func NewLeaderboardService(redisClient *redis.Client) * LeaderboardService{
	return &LeaderboardService{
		redisClient: redisClient,
	}
}

//UpdatePlayerScore atualiza score de uma jogador no leaderboard
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

//GetPlayerRank retorna o rank de um jogador
func (ls *LeaderboardService) GetPlayerRank(characterID string, characterName string) (int64, error) {
    ctx := context.Background()

    memberKey := fmt.Sprintf("%s:%s", characterID, characterName)

    rawClient := ls.redisClient.GetRawClient()
    rank, err := rawClient.ZRevRank(ctx, LeaderboardKey, memberKey).Result()
    if err != nil {
        if err == redis.Nil {
            return -1, nil // Jogador não encontrado
        }
        return -1, fmt.Errorf("Erro ao buscar rank: %w", err)
    }

    return rank + 1, nil // Redis retorna 0-indexed, convertemos para 1-indexed
}