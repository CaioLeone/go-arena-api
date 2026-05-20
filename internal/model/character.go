package model

import "github.com/google/uuid"

type CharacterModel struct {
	id            uuid.UUID
	userId        uuid.UUID
	name          string
	class         string
	level         int
	hp            int
	attack        int
	defense       int
	rankingPoints int
}
