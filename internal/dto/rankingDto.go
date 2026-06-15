package dto

import "github.com/google/uuid"

type UserRankingResponse struct {
	CharacterID uuid.UUID `json:"character_id"`
	Name        string    `json:"name"`
	Class       string    `json:"class"`
	Level       int       `json:"level"`
	Rank        int64     `json:"rank"`
	Score       int64     `json:"score"`
}

type TopPlayersResponse struct {
	Rank  int64  `json:"rank"`
	Name  string `json:"name"`
	Score int64  `json:"score"`
}

type LeaderboardResponse struct {
	Players []TopPlayersResponse `json:"players"`
	Total   int                  `json:"total"`
}
