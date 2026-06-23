package models

import (
    "time"

	"github.com/google/uuid"
)

type Room struct{
	ID uuid.UUID `gorm:"type:uuid;primarykey"`
	Username string
	Name string
	Created_At time.Time
}