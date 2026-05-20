package dto

import (
	"time"

	"github.com/google/uuid"
)

type BattleCreateRequest struct {
	DefenderID uuid.UUID `json:"defender_id" validate:"required"`
}

// RESPONSE
type BattleResponse struct {
	ID          uuid.UUID `json:"id"`
	AttackerID  uuid.UUID `json:"attacker_id"`
	DefenderID  uuid.UUID `json:"defender_id"`
	WinnerID    uuid.UUID `json:"winner_id"`
	DamageDealt int       `json:"damage_dealt"`
	CreatedAt   time.Time `json:"created_at"`
}

type BattleHistoryResponse struct {
	Battles []BattleResponse `json:"battles"`
	Total   int              `json:"total"`
}
