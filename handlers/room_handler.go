package handlers

type RoomRequest struct {
	Username string `json:"username"`
	Name string `json:"name"`
	// UserID string `json:"userId"`
}