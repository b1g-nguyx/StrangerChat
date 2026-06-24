package entity

import "time"

// ChatMessage represents a message within the chat bounded context.
// It is specifically tailored for WebSocket delivery and Redis persistence.
type ChatMessage struct {
	SenderID  string    `json:"sender_id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}
