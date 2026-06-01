package model

import (
	"time"

	"github.com/google/uuid"
)

type CharacterModel struct {
	ID            uuid.UUID
	UserId        uuid.UUID
	Name          string
	Class         string
	Level         int
	Hp            int
	Attack        int
	Defense       int
	RankingPoints int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
