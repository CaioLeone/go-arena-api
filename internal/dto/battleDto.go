package dto

import (
	"github.com/google/uuid"
)

type BattleCreateRequest struct {
	AttackerCharacterID string `json:"attacker_character_id" validate:"required,uuid"`
	DefenderCharacterID string `json:"defender_character_id" validate:"required,uuid"`
}

type BattleRound struct {
	AttackerDamage int    `json:"attacker_damage"`
	DefenderHP     int    `json:"defender_hp"`
	IsCritical     bool   `json:"is_critical"`
	Message        string `json:"message"`
}

// RESPONSE
type BattleResponse struct {
	ID              uuid.UUID     `json:"id"`
	AttackerID      uuid.UUID     `json:"attacker_id"`
	AttackerName    string        `json:"attacker_name"`
	DefenderID      uuid.UUID     `json:"defender_id"`
	DefenderName    string        `json:"defender_name"`
	WinnerID        uuid.UUID     `json:"winner_id"`
	WinnerName      string        `json:"winner_name"`
	AttackerHPFinal int           `json:"attacker_hp_final"`
	DefenderHPFinal int           `json:"defender_hp_final"`
	RoundsCount     int           `json:"rounds_count"`
	Rounds          []BattleRound `json:"rounds"`
	CreatedAt       string        `json:"created_at"`
}

type BattleHistoryResponse struct {
	ID              uuid.UUID `json:"id"`
	AttackerID      uuid.UUID `json:"attacker_id"`
	AttackerName    string    `json:"attacker_name"`
	DefenderID      uuid.UUID `json:"defender_id"`
	DefenderName    string    `json:"defender_name"`
	WinnerID        uuid.UUID `json:"winner_id"`
	WinnerName      string    `json:"winner_name"`
	AttackerHPFinal int       `json:"attacker_hp_final"`
	DefenderHPFinal int       `json:"defender_hp_final"`
	RoundsCount     int       `json:"rounds_count"`
	CreatedAt       string    `json:"created_at"`
}
