package usecase

import (
	"context"
	"time"

	"github.com/b1g-nguyx/strangerchat-backend/internal/broker"
	"github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/entity"
	"github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/repository"
)

type ChatUsecase interface {
	SendMessage(ctx context.Context, roomID string, message entity.ChatMessage) error
	ReportUser(ctx context.Context, reporterID, reportedID, roomID, reason string) error
}

type chatUsecaseImpl struct {
	redisRepo repository.RedisRoomRepo
}

func NewChatUsecase(redisRepo repository.RedisRoomRepo) ChatUsecase {
	return &chatUsecaseImpl{redisRepo: redisRepo}
}

func (u *chatUsecaseImpl) SendMessage(ctx context.Context, roomID string, message entity.ChatMessage) error {
	// TTL for chat logs: 1 hour
	return u.redisRepo.AppendChatLog(ctx, roomID, message, time.Hour)
}

func (u *chatUsecaseImpl) ReportUser(ctx context.Context, reporterID, reportedID, roomID, reason string) error {
	// 1. Get current chat log from Redis
	logs, err := u.redisRepo.GetChatLog(ctx, roomID)
	if err != nil {
		return err // In real app, might want to handle this gracefully if logs expired
	}

	// 2. Publish to RabbitMQ
	eventPayload := map[string]interface{}{
		"reporter_id": reporterID,
		"reported_id": reportedID,
		"room_id":     roomID,
		"reason":      reason,
		"chat_logs":   logs, // array of strings
	}

	if broker.RMQ != nil {
		return broker.RMQ.PublishMessage(ctx, eventPayload)
	}
	return nil
}
