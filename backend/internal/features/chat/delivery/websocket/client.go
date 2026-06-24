package websocket

import (
	"context"
	"log"
	"time"

	"github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/entity"
	"github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/usecase"
	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	Hub       *Hub
	Conn      *websocket.Conn
	Send      chan []byte
	SessionID string
	UserID    string
	Usecase   usecase.ChatUsecase
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		if messageType == websocket.TextMessage {
			// 1. Broadcast message to other clients in Hub
			c.Hub.Broadcast <- message

			// 2. Save to Redis via Usecase
			if c.Usecase != nil {
				chatMsg := entity.ChatMessage{
					SenderID:  c.UserID,
					Content:   string(message),
					Timestamp: time.Now(),
				}
				if err := c.Usecase.SendMessage(context.Background(), c.SessionID, chatMsg); err != nil {
					log.Printf("Failed to save chat log: %v", err)
				}
			}
		}
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Send
		if !ok {
			c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		c.Conn.WriteMessage(websocket.TextMessage, message)
	}
}
