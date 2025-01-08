package domain

import (
	"time"

	"github.com/google/uuid"
)

type Example struct {
	ID      uuid.UUID
	Example string

	CreatedAt time.Time
	UpdatedAt time.Time

	Optional *string
}
