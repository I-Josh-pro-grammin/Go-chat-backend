package handlers

import (
	// "github.com/google/uuid"
	"time"
)

type UserRequest struct{
	// ID uuid.UUID `json:"id"`
	Username string `json:"username"`
	Created_At time.Time `json:"created_at"`
}