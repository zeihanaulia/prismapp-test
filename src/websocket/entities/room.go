package entity

import (
	"time"
)

// Room entity
type Room struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
