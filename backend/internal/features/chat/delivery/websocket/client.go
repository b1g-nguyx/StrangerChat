package websocket

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/entity"
	"github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/usecase"
	"github.com/gofiber/contrib/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 32768
)

type Client struct {
	Hub       *Hub
	Conn      *websocket.Conn
	Send      chan []byte
	SessionID string
	UserID    string
	RoomID    string // Track which room the client is currently in
	Usecase   usecase.ChatUsecase
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		if c.RoomID != "" && c.Usecase != nil {
			c.Usecase.LeaveRoom(context.Background(), c.RoomID, c.UserID)
		} else if c.Usecase != nil {
			// If not in a room, maybe in queue. Remove them.
			c.Usecase.CancelMatchmaking(context.Background(), c.UserID)
		}
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		messageType, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		if messageType == websocket.TextMessage {
			var wsMsg WSMessage
			if err := json.Unmarshal(message, &wsMsg); err != nil {
				log.Printf("Invalid message format: %v", err)
				continue
			}

			switch wsMsg.Type {
			case MsgTypeFindMatch:
				if c.Usecase != nil {
					c.Usecase.FindMatch(context.Background(), c.UserID)
				}
			case MsgTypeChat:
				if c.RoomID == "" {
					continue
				}
				
				if c.Usecase != nil {
					chatMsg := entity.ChatMessage{
						SenderID:  c.UserID,
						Content:   wsMsg.Content,
						Timestamp: time.Now(),
					}
					if err := c.Usecase.SendMessage(context.Background(), c.RoomID, chatMsg); err != nil {
						log.Printf("Failed to send message: %v", err)
					}
				}
			case MsgTypeLeave:
				if c.RoomID != "" && c.Usecase != nil {
					c.Usecase.LeaveRoom(context.Background(), c.RoomID, c.UserID)
					c.RoomID = "" // clear room
				}
				c.Hub.Unregister <- c
			case MsgTypeReport:
				if c.Usecase != nil {
					if err := c.Usecase.ReportUser(context.Background(), c.UserID, wsMsg.ReportedID, wsMsg.RoomID, wsMsg.Content); err != nil {
						log.Printf("Failed to report user: %v", err)
					}
				}
			case MsgTypeWebRTCOffer, MsgTypeWebRTCAnswer, MsgTypeWebRTCICECandid:
				if c.RoomID != "" && c.Usecase != nil {
					if err := c.Usecase.SendWebRTCSignal(context.Background(), c.RoomID, c.UserID, string(wsMsg.Type), wsMsg.Payload); err != nil {
						log.Printf("Failed to send WebRTC signal: %v", err)
					}
				}
			}
		}
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Conn.WriteMessage(websocket.TextMessage, message)
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
