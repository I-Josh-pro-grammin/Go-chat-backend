package handlers

type RoomRequest struct {
	Name string `json:"name"`
	UserID string `json:"userId"`
}