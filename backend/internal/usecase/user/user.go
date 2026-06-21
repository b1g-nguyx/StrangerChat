package user

import (
	"context"
	"fmt"
	"time"

	"github.com/b1g-nguyx/strangerchat-backend/internal/entity"
	"github.com/b1g-nguyx/strangerchat-backend/internal/repo"
	"github.com/b1g-nguyx/strangerchat-backend/pkg/filter"
	"github.com/b1g-nguyx/strangerchat-backend/pkg/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

// Register -.
func (uc *UseCase) Register(ctx context.Context, username, email, password string) (entity.User, string, string, error) {
	//check is exit by email
	existingUser, err := uc.repo.GetByEmail(ctx, email)
	if err == nil && existingUser.ID != "" {
		return entity.User{}, "", "", fmt.Errorf("email %s is exit", email)
	}
	//bcrypt password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, "", "", fmt.Errorf("UserUseCase - Register - bcrypt.GenerateFromPassword: %w", err)
	}

	now := time.Now().UTC()
	userID := uuid.New().String()

	// Generate Tokens
	accessToken, err := uc.jwt.GenerateToken(userID)
	if err != nil {
		return entity.User{}, "", "", fmt.Errorf("UserUseCase - Register - GenerateToken: %w", err)
	}

	refreshToken, err := uc.jwt.GenerateRefreshToken()
	if err != nil {
		return entity.User{}, "", "", fmt.Errorf("UserUseCase - Register - GenerateRefreshToken: %w", err)
	}

	user := entity.User{
		BaseEntity: entity.BaseEntity{
			ID:        userID,
			CreatedAt: now,
			UpdatedAt: now,
		},
		Username:     username,
		Email:        email,
		PasswordHash: string(hash),
		RefreshToken: refreshToken,
	}

	err = uc.repo.Insert(ctx, &user)
	if err != nil {
		return entity.User{}, "", "", fmt.Errorf("UserUseCase - Register - uc.repo.Insert: %w", err)
	}

	return user, accessToken, refreshToken, nil
}

// Login -.
func (uc *UseCase) Login(ctx context.Context, email, password string) (entity.User, string, string, error) {
	user, err := uc.repo.GetByEmail(ctx, email)
	if err != nil {
		return entity.User{}, "", "", entity.ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return entity.User{}, "", "", entity.ErrInvalidCredentials
	}

	accessToken, err := uc.jwt.GenerateToken(user.ID)
	if err != nil {
		return entity.User{}, "", "", fmt.Errorf("UserUseCase - Login - GenerateToken: %w", err)
	}

	refreshToken, err := uc.jwt.GenerateRefreshToken()
	if err != nil {
		return entity.User{}, "", "", fmt.Errorf("UserUseCase - Login - GenerateRefreshToken: %w", err)
	}

	// Update the Refresh Token in the database
	user.RefreshToken = refreshToken
	err = uc.repo.Update(ctx, &user)
	if err != nil {
		// Log the error but continue to allow login
		fmt.Printf("Warning: failed to save refresh token: %v\n", err)
	}

	return user, accessToken, refreshToken, nil
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
