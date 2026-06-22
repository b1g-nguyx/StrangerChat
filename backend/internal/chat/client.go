package chat

import (
	"context"
	"log"

	"github.com/b1g-nguyx/strangerchat-backend/internal/broker"
	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	Hub       *Hub
	Conn      *websocket.Conn
	Send      chan []byte
	SessionID string
	UserID    string
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

			// 2. Publish to RabbitMQ for logging
			logMsg := map[string]interface{}{
				"session_id": c.SessionID,
				"user_id":    c.UserID,
				"content":    string(message),
			}

			if broker.RMQ != nil {
				if err := broker.RMQ.PublishMessage(context.Background(), logMsg); err != nil {
					log.Printf("Failed to publish to RabbitMQ: %v", err)
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
