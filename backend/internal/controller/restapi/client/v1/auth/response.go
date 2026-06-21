package auth

import "github.com/b1g-nguyx/strangerchat-backend/internal/entity"

// RegisterResponse defines the standard success payload for registration.
type AuthResponse struct {
	Message      string       `json:"message"`
	Data         *entity.User `json:"data"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
}

// ErrorResponse defines the standard error payload.
type ErrorResponse struct {
	Error string `json:"error"`
}
