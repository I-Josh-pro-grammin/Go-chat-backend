package models

import (
	"time"
	
	"github.com/google/uuid"
)

type User struct{
	ID uuid.UUID `gorm:"type:uuid;primarykey"`
	Username string
	Created_At time.Time
}