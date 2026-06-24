package entity

import "time"

type Room struct {
	BaseEntity

	User1ID  *string    `json:"user1_id"`
	User2ID  *string    `json:"user2_id"`
	Status   string     `json:"status"`
	ClosedAt *time.Time `json:"closed_at,omitempty"`
}
