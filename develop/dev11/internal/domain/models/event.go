package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	OwnerID     uuid.UUID `json:"owner_id"`
}
