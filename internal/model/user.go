package main

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	id        uuid.UUID
	name      string
	email     string
	password  string
	createdAt time.Time
}
