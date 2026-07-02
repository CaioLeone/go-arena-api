package dto

import "github.com/google/uuid"

type UserRankingResponse struct {
	CharacterID uuid.UUID `json:"character_id"`
	Name        string    `json:"name" validate:"required"`
	Class       string    `json:"class" validate:"required"`
	Level       int       `json:"level" validate:"required,min=1"`
	Rank        int64     `json:"rank"`
	Score       int64     `json:"score"`
}

type TopPlayersResponse struct {
	Rank        int64     `json:"rank"`
	CharacterID uuid.UUID `json:"character_id"`
	Name        string    `json:"name" validate:"required"`
	Class       string    `json:"class" validate:"required"`
    Level       int       `json:"level" validate:"required,min=1"`
	Score       int64     `json:"score"`
}

type LeaderboardResponse struct {
	Players []TopPlayersResponse `json:"players"`
	Total   int                  `json:"total"`
}
