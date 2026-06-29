package usecase

import (
	"context"
	"time"

	"encoding/json"
	"log"

	"github.com/b1g-nguyx/strangerchat-backend/internal/broker"
	"github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/entity"
	"github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/repository"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type ChatUsecase interface {
	SendMessage(ctx context.Context, roomID string, message entity.ChatMessage) error
	ReportUser(ctx context.Context, reporterID, reportedID, roomID, reason string) error
	FindMatch(ctx context.Context, userID string) error
	CancelMatchmaking(ctx context.Context, userID string) error
	LeaveRoom(ctx context.Context, roomID, userID string) error
	StartMatchmakingWorker(ctx context.Context)
	SendWebRTCSignal(ctx context.Context, roomID, userID, signalType string, payload interface{}) error
}

type chatUsecaseImpl struct {
	redisRepo repository.RedisRoomRepo
}

func NewChatUsecase(redisRepo repository.RedisRoomRepo) ChatUsecase {
	return &chatUsecaseImpl{redisRepo: redisRepo}
}

func (u *chatUsecaseImpl) SendMessage(ctx context.Context, roomID string, message entity.ChatMessage) error {
	eventPayload := map[string]interface{}{
		"type":      "chat",
		"room_id":   roomID,
		"content":   message.Content,
		"sender_id": message.SenderID,
		// Assuming we broadcast to everyone in the room. Real implementation might look up room state to get receiver_id.
		"receiver_id": "broadcast",
	}
	u.redisRepo.PublishEvent(ctx, "global_chat_events", eventPayload)
	return u.redisRepo.AppendChatLog(ctx, roomID, message, time.Hour)
}

func (u *chatUsecaseImpl) ReportUser(ctx context.Context, reporterID, reportedID, roomID, reason string) error {
	logs, err := u.redisRepo.GetChatLog(ctx, roomID)
	if err != nil {
		return err
	}

	eventPayload := map[string]interface{}{
		"reporter_id": reporterID,
		"reported_id": reportedID,
		"room_id":     roomID,
		"reason":      reason,
		"chat_logs":   logs,
	}

	if broker.RMQ != nil {
		return broker.RMQ.PublishMessage(ctx, eventPayload)
	}
	return nil
}

func (u *chatUsecaseImpl) FindMatch(ctx context.Context, userID string) error {
	return u.redisRepo.EnqueueMatchmaking(ctx, userID)
}

func (u *chatUsecaseImpl) CancelMatchmaking(ctx context.Context, userID string) error {
	return u.redisRepo.RemoveFromQueue(ctx, userID)
}

func (u *chatUsecaseImpl) LeaveRoom(ctx context.Context, roomID, userID string) error {
	// Publish an event that this user left the room
	eventPayload := map[string]interface{}{
		"type":    "partner_left",
		"room_id": roomID,
		"user_id": userID,
	}
	u.redisRepo.DeleteRoom(ctx, roomID) // Simple implementation: delete room on leave
	return u.redisRepo.PublishEvent(ctx, "global_chat_events", eventPayload)
}

func (u *chatUsecaseImpl) SendWebRTCSignal(ctx context.Context, roomID, userID, signalType string, payload interface{}) error {
	eventPayload := map[string]interface{}{
		"type":        "webrtc_signal",
		"room_id":     roomID,
		"sender_id":   userID,
		"signal_type": signalType,
		"payload":     payload,
	}
	return u.redisRepo.PublishEvent(ctx, "global_chat_events", eventPayload)
}

func (u *chatUsecaseImpl) StartMatchmakingWorker(ctx context.Context) {
	go func() {
		log.Println("Matchmaking worker started...")
		for {
			select {
			case <-ctx.Done():
				return
			default:
				user1, err := u.redisRepo.DequeueMatchmaking(ctx, 2*time.Second)
				if err != nil {
					if err != redis.Nil && err != context.DeadlineExceeded && err != context.Canceled {
						log.Printf("Matchmaking dequeue error (user1): %v", err)
					}
					continue
				}

				user2, err := u.redisRepo.DequeueMatchmaking(ctx, 5*time.Second)
				if err != nil {
					// Put user1 back if we couldn't find user2 within timeout
					u.redisRepo.EnqueueMatchmaking(ctx, user1)
					continue
				}

				if user1 == user2 {
					// Prevent matching with self if user sent multiple requests
					u.redisRepo.EnqueueMatchmaking(ctx, user1)
					continue
				}

				// We have a match!
				roomID := uuid.New().String()
				
				// Save room state
				roomData, _ := json.Marshal(map[string]string{"user1": user1, "user2": user2})
				u.redisRepo.SaveRoom(ctx, roomID, string(roomData), 24*time.Hour)

				// Publish match event
				matchEvent := map[string]interface{}{
					"type":    "matched",
					"room_id": roomID,
					"user1":   user1,
					"user2":   user2,
				}
				if err := u.redisRepo.PublishEvent(ctx, "global_chat_events", matchEvent); err != nil {
					log.Printf("Failed to publish match event: %v", err)
				}
			}
		}
	}()
}
