package model

import (
	"time"

	"github.com/google/uuid"
)

type CharacterModel struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	Name          string
	Class         string
	Level         int
	HP            int
	Attack        int
	Defense       int
	RankingPoints int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
