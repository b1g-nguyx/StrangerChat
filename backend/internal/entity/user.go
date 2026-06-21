package entity

import "time"

type User struct {
	BaseEntity

	// Thông tin định danh cơ bản
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"` // Giữ bí mật, không trả về qua API
	DisplayName  string `json:"display_name"`
	AvatarURL    string `json:"avatar_url,omitempty"`

	// Thông tin trạng thái
	Status        string     `json:"status"`
	CurrentRoomID *string    `json:"current_room_id,omitempty"`
	RefreshToken  string     `json:"-"`
	IsBanned      bool       `json:"is_banned"`
	BannedAt      *time.Time `json:"banned_at,omitempty"`
}
