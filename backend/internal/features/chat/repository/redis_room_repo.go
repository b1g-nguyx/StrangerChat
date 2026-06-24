package repository

import (
	"context"
	"time"

	"encoding/json"

	"github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/entity"
	"github.com/redis/go-redis/v9"
)

type RedisRoomRepo interface {
	SaveRoom(ctx context.Context, roomID string, data string, ttl time.Duration) error
	DeleteRoom(ctx context.Context, roomID string) error
	AppendChatLog(ctx context.Context, roomID string, message entity.ChatMessage, ttl time.Duration) error
	GetChatLog(ctx context.Context, roomID string) ([]entity.ChatMessage, error)
}

type redisRoomRepoImpl struct {
	client *redis.Client
}

func NewRedisRoomRepo(client *redis.Client) RedisRoomRepo {
	return &redisRoomRepoImpl{client: client}
}

func (r *redisRoomRepoImpl) SaveRoom(ctx context.Context, roomID string, data string, ttl time.Duration) error {
	key := "room:" + roomID
	return r.client.Set(ctx, key, data, ttl).Err()
}

func (r *redisRoomRepoImpl) DeleteRoom(ctx context.Context, roomID string) error {
	key := "room:" + roomID
	return r.client.Del(ctx, key).Err()
}

func (r *redisRoomRepoImpl) AppendChatLog(ctx context.Context, roomID string, message entity.ChatMessage, ttl time.Duration) error {
	key := "chat_log:" + roomID
	
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	pipe := r.client.Pipeline()
	pipe.RPush(ctx, key, string(bytes))
	pipe.Expire(ctx, key, ttl)
	_, err = pipe.Exec(ctx)
	return err
}

func (r *redisRoomRepoImpl) GetChatLog(ctx context.Context, roomID string) ([]entity.ChatMessage, error) {
	key := "chat_log:" + roomID
	rawLogs, err := r.client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var messages []entity.ChatMessage
	for _, raw := range rawLogs {
		var msg entity.ChatMessage
		if err := json.Unmarshal([]byte(raw), &msg); err == nil {
			messages = append(messages, msg)
		}
	}
	return messages, nil
}
