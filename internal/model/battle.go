package model

import (
	"time"

	"github.com/google/uuid"
)

type BattleModel struct {
	id          uuid.UUID
	attackerId  uuid.UUID
	defenderId  uuid.UUID
	winnerId    uuid.UUID
	damageDealt int
	createdAt   time.Time
}
