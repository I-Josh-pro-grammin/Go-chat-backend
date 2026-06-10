package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct{
	ID uuid.UUID `gorm:"uuid; primaryKey"`
	UserID uuid.UUID
	RoomID uuid.UUID

	Content string
	Created_At time.Time
}