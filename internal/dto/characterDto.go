package dto

import "github.com/google/uuid"

type CharacterCreateRequest struct {
	Name  string `json:"name" validate:"required,min=1,max=25"`
	Class string `json:"class" validate:"required,oneof=Barbaro Mago Arqueiro Assassino"`
}

type CharacterUpdateRequest struct {
	Name  string `json:"name" validate:"required,min=1,max=25"`
	Class string `json:"class" validate:"required,oneof=Barbaro Mago Arqueiro Assassino"`
}

// RESPONSE

type CharacterResponse struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	Name          string    `json:"name"`
	Class         string    `json:"class"`
	Level         int       `json:"level"`
	HP            int       `json:"hp"`
	Attack        int       `json:"attack"`
	Defense       int       `json:"defense"`
	RankingPoints int       `json:"ranking_points"`
}

type CharacterListResponse struct {
	Characters []CharacterResponse `json:"characters"`
	Total      int                 `json:"total"`
}
