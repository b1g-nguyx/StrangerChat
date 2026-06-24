package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/b1g-nguyx/strangerchat-backend/internal/common/jwt"
	"github.com/b1g-nguyx/strangerchat-backend/internal/core/entity"
	"github.com/b1g-nguyx/strangerchat-backend/internal/repo"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// AuthUseCase -.
type AuthUseCase struct {
	repo repo.UserRepo
	jwt  *jwt.Manager
}

// NewAuthUseCase -.
func NewAuthUseCase(r repo.UserRepo, j *jwt.Manager) *AuthUseCase {
	return &AuthUseCase{
		repo: r,
		jwt:  j,
	}
}

// Register -.
func (uc *AuthUseCase) Register(ctx context.Context, username, email, password string) (entity.User, string, string, error) {
	//check is exit by email
	existingUser, err := uc.repo.GetByEmail(ctx, email)
	if err == nil && existingUser.ID != "" {
		return entity.User{}, "", "", fmt.Errorf("email %s is exit", email)
	}
	//bcrypt password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, "", "", fmt.Errorf("AuthUseCase - Register - bcrypt.GenerateFromPassword: %w", err)
	}

	now := time.Now().UTC()
	userID := uuid.New().String()

	// Generate Tokens
	accessToken, err := uc.jwt.GenerateToken(userID)
	if err != nil {
		return entity.User{}, "", "", fmt.Errorf("AuthUseCase - Register - GenerateToken: %w", err)
	}

	refreshToken, err := uc.jwt.GenerateRefreshToken()
	if err != nil {
		return entity.User{}, "", "", fmt.Errorf("AuthUseCase - Register - GenerateRefreshToken: %w", err)
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
		return entity.User{}, "", "", fmt.Errorf("AuthUseCase - Register - uc.repo.Insert: %w", err)
	}

	return user, accessToken, refreshToken, nil
}

// Login -.
func (uc *AuthUseCase) Login(ctx context.Context, email, password string) (entity.User, string, string, error) {
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
		return entity.User{}, "", "", fmt.Errorf("AuthUseCase - Login - GenerateToken: %w", err)
	}

	refreshToken, err := uc.jwt.GenerateRefreshToken()
	if err != nil {
		return entity.User{}, "", "", fmt.Errorf("AuthUseCase - Login - GenerateRefreshToken: %w", err)
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
