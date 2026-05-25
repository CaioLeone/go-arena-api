package model

import "github.com/google/uuid"

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
}
