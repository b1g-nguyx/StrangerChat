package websocket

import (
	"context"
	"encoding/json"
	"log"

	"github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/repository"
)

type Hub struct {
	// Local state for connections to THIS instance only
	Clients map[string]*Client // map UserID to Client

	Register   chan *Client
	Unregister chan *Client

	// Redis dependency for listening to global events
	redisRepo repository.RedisRoomRepo
}

func NewHub(redisRepo repository.RedisRoomRepo) *Hub {
	return &Hub{
		Clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		redisRepo:  redisRepo,
	}
}

func (h *Hub) Run(ctx context.Context) {
	// Subscribe to global chat events
	pubsub := h.redisRepo.SubscribeEvent(ctx, "global_chat_events")
	defer pubsub.Close()
	ch := pubsub.Channel()

	for {
		select {
		case <-ctx.Done():
			return
		case client := <-h.Register:
			h.Clients[client.UserID] = client
			log.Printf("Client registered: %s", client.UserID)
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.UserID]; ok {
				delete(h.Clients, client.UserID)
				close(client.Send)
				log.Printf("Client unregistered: %s", client.UserID)
			}
		case msg := <-ch:
			// Process global events
			var event map[string]interface{}
			if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
				log.Printf("Failed to unmarshal event: %v", err)
				continue
			}

			eventType, ok := event["type"].(string)
			if !ok {
				continue
			}

			switch eventType {
			case "matched":
				roomID := event["room_id"].(string)
				user1 := event["user1"].(string)
				user2 := event["user2"].(string)

				matchMsg := WSMessage{
					Type:   MsgTypeMatched,
					RoomID: roomID,
				}
				matchMsgBytes, _ := json.Marshal(matchMsg)

				// If user1 is connected to this instance
				if c1, ok := h.Clients[user1]; ok {
					c1.RoomID = roomID
					c1.Send <- matchMsgBytes
				}
				// If user2 is connected to this instance
				if c2, ok := h.Clients[user2]; ok {
					c2.RoomID = roomID
					c2.Send <- matchMsgBytes
				}
			case "chat":
				roomID := event["room_id"].(string)
				content := event["content"].(string)
				senderID := event["sender_id"].(string)
				receiverID := event["receiver_id"].(string)

				chatMsg := WSMessage{
					Type:    MsgTypeChat,
					RoomID:  roomID,
					Content: content,
				}
				chatMsgBytes, _ := json.Marshal(chatMsg)

				if receiverID == "broadcast" {
					for _, client := range h.Clients {
						// Broadcast to all clients in the same room except the sender
						if client.RoomID == roomID && client.UserID != senderID {
							client.Send <- chatMsgBytes
						}
					}
				} else {
					if receiver, ok := h.Clients[receiverID]; ok {
						receiver.Send <- chatMsgBytes
					}
				}
			case "webrtc_signal":
				roomID := event["room_id"].(string)
				senderID := event["sender_id"].(string)
				signalType := event["signal_type"].(string)
				
				signalMsg := WSMessage{
					Type:    WSMessageType(signalType),
					RoomID:  roomID,
					Payload: event["payload"],
				}
				signalMsgBytes, _ := json.Marshal(signalMsg)

				for _, client := range h.Clients {
					if client.RoomID == roomID && client.UserID != senderID {
						// Run in goroutine to not block chat messages
						go func(c *Client, msg []byte) {
							select {
							case c.Send <- msg:
							default:
							}
						}(client, signalMsgBytes)
					}
				}
			case "partner_left":
				roomID := event["room_id"].(string)
				userID := event["user_id"].(string) // The user who left

				leaveMsg := WSMessage{
					Type:   MsgTypePartnerLeft,
					RoomID: roomID,
				}
				leaveMsgBytes, _ := json.Marshal(leaveMsg)

				for _, client := range h.Clients {
					if client.RoomID == roomID && client.UserID != userID {
						client.RoomID = "" // Clear room id since partner left
						select {
						case client.Send <- leaveMsgBytes:
						default:
							close(client.Send)
							delete(h.Clients, client.UserID)
						}
					}
				}
			}
		}
	}
}
