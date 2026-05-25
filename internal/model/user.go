package model

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}
