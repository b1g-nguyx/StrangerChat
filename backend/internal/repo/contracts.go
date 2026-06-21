// Package repo implements application outer layer logic.
package repo

import (
	"context"

	"github.com/b1g-nguyx/strangerchat-backend/internal/entity"
)

type (
	// UserRepo -.
	UserRepo interface {
		Insert(ctx context.Context, user *entity.User) error
		Update(ctx context.Context, user *entity.User) error
		GetByID(ctx context.Context, id string) (entity.User, error)
		GetByEmail(ctx context.Context, email string) (entity.User, error)
		GetUsers(ctx context.Context, filters map[string]any) ([]entity.User, error)
	}
)
