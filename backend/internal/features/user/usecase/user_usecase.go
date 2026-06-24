package user

import (
	"context"
	"fmt"

	"github.com/b1g-nguyx/strangerchat-backend/internal/common/filter"
	"github.com/b1g-nguyx/strangerchat-backend/internal/common/jwt"
	"github.com/b1g-nguyx/strangerchat-backend/internal/core/entity"
	"github.com/b1g-nguyx/strangerchat-backend/internal/repo"
)

// UseCase -.
type UseCase struct {
	repo repo.UserRepo
	jwt  *jwt.Manager
}

// New -.
func New(r repo.UserRepo, j *jwt.Manager) *UseCase {
	return &UseCase{
		repo: r,
		jwt:  j,
	}
}

// GetUser -.
func (uc *UseCase) GetUser(ctx context.Context, userID string) (entity.User, error) {
	user, err := uc.repo.GetByID(ctx, userID)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserUseCase - GetUser - uc.repo.GetByID: %w", err)
	}

	return user, nil
}

// GetUsers processes raw input filters, sanitizes them, and fetches users from the repository.
func (uc *UseCase) GetUsers(ctx context.Context, inputFilters map[string]any) ([]entity.User, error) {
	// 1. Define allowed columns for querying to prevent SQL Column Injection
	allowedFields := []string{
		"username",
		"email",
		"status",
		"is_banned",
	}

	// 2. Sanitize the filters using our utility function
	safeFilters := filter.AllowedKeys(inputFilters, allowedFields)

	// 3. Pass the strictly safe filters to the Repository layer
	users, err := uc.repo.GetUsers(ctx, safeFilters)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetUsers - uc.repo.GetUsers: %w", err)
	}

	return users, nil
}
