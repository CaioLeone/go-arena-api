package dto

import "github.com/google/uuid"

type CharacterCreateRequest struct {
	Name    string `json:"name" validate:"required,min=1,max=25"`
	Class   string `json:"class" validate:"required,oneof=Barbaro Mago Arqueiro Assassino"`
	Level   int    `json:"level" validate:"omitempty,required,min=1,max=100"`
	HP      int    `json:"hp" validate:"omitempty,required,min=1"`
	Attack  int    `json:"attack" validate:"omitempty,required,min=1"`
	Defense int    `json:"defense" validate:"omitempty,required,min=1"`
	CriticalChance int    `json:"critical_chance" validate:"omitempty,min=0,max=100"`
}

type CharacterUpdateRequest struct {
	Name    string `json:"name" validate:"required,min=1,max=25"`
	Class   string `json:"class" validate:"required,oneof=Barbaro Mago Arqueiro Assassino"`
	Level   int    `json:"level" validate:"omitempty,required,min=1,max=100"`
	HP      int    `json:"hp" validate:"omitempty,required,min=1"`
	Attack  int    `json:"attack" validate:"omitempty,required,min=1"`
	Defense int    `json:"defense" validate:"omitempty,required,min=1"`
	CriticalChance int    `json:"critical_chance" validate:"omitempty,min=0,max=100"`
}

// RESPONSE

type CharacterResponse struct {
	ID             uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"user_id"`
	Name           string    `json:"name"`
	Class          string    `json:"class"`
	Level          int       `json:"level"`
	HP             int       `json:"hp"`
	Attack         int       `json:"attack"`
	Defense        int       `json:"defense"`
	RankingPoints  int       `json:"ranking_points"`
	CriticalChance int       `json:"critical_chance"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
}

type CharacterListResponse struct {
	Characters []CharacterResponse `json:"characters"`
	Total      int                 `json:"total"`
}
