package httpdelivery

import (
	"time"

	"github.com/b1g-nguyx/strangerchat-backend/internal/core/entity"
)

// --- REQUEST DTOs ---

// RegisterRequest defines the mandatory payload structure for registration.
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// LoginRequest defines the payload structure for login.
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// --- RESPONSE DTOs ---

// UserDTO (Data Transfer Object) chỉ chứa những trường an toàn muốn trả về cho client.
// Nó loại bỏ đi PasswordHash, DeletedAt, v.v. từ entity.User gốc.
type UserDTO struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name"`
	AvatarURL   string    `json:"avatar_url"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

// AuthData là dữ liệu sẽ được truyền vào tham số `data` của hàm response.Success()
type AuthData struct {
	User         UserDTO `json:"user"`
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
}

// ToUserDTO là một hàm helper để copy dữ liệu từ Entity sang DTO
func ToUserDTO(user entity.User) UserDTO {
	return UserDTO{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		AvatarURL:   user.AvatarURL,
		Status:      user.Status,
		CreatedAt:   user.CreatedAt,
	}
}
