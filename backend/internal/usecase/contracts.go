// Package usecase implements application business logic.
package usecase

import (
	"context"

	"github.com/b1g-nguyx/strangerchat-backend/internal/entity"
)

type (
	// User -.
	User interface {
		Register(ctx context.Context, username, email, password string) (entity.User, error)
		Login(ctx context.Context, email, password string) (string, error)
		GetUser(ctx context.Context, userID string) (entity.User, error)
	}
)
