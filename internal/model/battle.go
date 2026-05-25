package model

import (
	"time"

	"github.com/google/uuid"
)

type BattleModel struct {
	ID          uuid.UUID
	AttackerId  uuid.UUID
	DefenderId  uuid.UUID
	WinnerId    uuid.UUID
	DamageDealt int
	CreatedAt   time.Time
}
