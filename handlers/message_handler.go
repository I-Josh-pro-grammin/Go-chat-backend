package handlers

type CreateMessageRequest struct {
   UserID string `json:"userId"`
   RoomID string `json:"roomId"`
   Content string `json:"content"`
}