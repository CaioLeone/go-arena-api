package model

import (
	"time"

	"github.com/google/uuid"
)

type BattleModel struct {
	ID              uuid.UUID
	AttackerID      uuid.UUID
	AttackerName    string
	DefenderID      uuid.UUID
	DefenderName    string
	DamageDealt     int
	WinnerID        *uuid.UUID
	WinnerName      string
	AttackerHPFinal int
	DefenderHPFinal int
	RoundsCount     int
	RoundsData      string // JSON armazenado no banco
	CreatedAt       time.Time
}
