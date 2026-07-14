package model

import (
	"time"

	"github.com/google/uuid"
)

type CharacterModel struct {
	ID              uuid.UUID
	UserID          uuid.UUID
	Name            string
	Class           string
	Level           int
	Experience      int
	HP              int
	Attack          int
	Defense         int
	AttributePoints int
	RankingPoints   int
	CriticalChance  int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
