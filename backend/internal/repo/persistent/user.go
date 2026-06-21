package persistent

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/b1g-nguyx/strangerchat-backend/internal/entity"
	"github.com/b1g-nguyx/strangerchat-backend/internal/repo"
)

type userRepoImpl struct {
	BaseRepo
}

func NewUserRepo(db *sql.DB) repo.UserRepo {
	return &userRepoImpl{
		BaseRepo: NewBaseRepo(db),
	}
}

func (r *userRepoImpl) GetByID(ctx context.Context, id string) (entity.User, error) {
	var user entity.User

	err := r.Builder.Select(
		"id", "created_at", "updated_at", "deleted_at",
		"username", "email", "password_hash", "display_name", "avatar_url",
		"status", "current_room_id", "refresh_token", "is_banned", "banned_at",
	).
		From("users").
		Where(squirrel.Eq{"id": id}).
		QueryRowContext(ctx).
		Scan(
			&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
			&user.Username, &user.Email, &user.PasswordHash, &user.DisplayName, &user.AvatarURL,
			&user.Status, &user.CurrentRoomID, &user.RefreshToken, &user.IsBanned, &user.BannedAt,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, nil
		}
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepoImpl) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User

	err := r.Builder.Select(
		"id", "created_at", "updated_at", "deleted_at",
		"username", "email", "password_hash", "display_name", "avatar_url",
		"status", "current_room_id", "refresh_token", "is_banned", "banned_at",
	).
		From("users").
		Where(squirrel.Eq{"email": email}).
		QueryRowContext(ctx).
		Scan(
			&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
			&user.Username, &user.Email, &user.PasswordHash, &user.DisplayName, &user.AvatarURL,
			&user.Status, &user.CurrentRoomID, &user.RefreshToken, &user.IsBanned, &user.BannedAt,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, entity.ErrInvalidCredentials
		}
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepoImpl) Insert(ctx context.Context, user *entity.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := r.Builder.Insert("users").
		Columns(
			"id", "created_at", "updated_at", "deleted_at",
			"username", "email", "password_hash", "display_name", "avatar_url",
			"status", "current_room_id", "refresh_token", "is_banned", "banned_at",
		).
		Values(
			user.ID, user.CreatedAt, user.UpdatedAt, user.DeletedAt,
			user.Username, user.Email, user.PasswordHash, user.DisplayName, user.AvatarURL,
			user.Status, user.CurrentRoomID, user.RefreshToken, user.IsBanned, user.BannedAt,
		).
		ExecContext(ctx)

	return err
}

func (r *userRepoImpl) Update(ctx context.Context, user *entity.User) error {
	user.UpdatedAt = time.Now()

	_, err := r.Builder.Update("users").
		Set("updated_at", user.UpdatedAt).
		Set("username", user.Username).
		Set("email", user.Email).
		Set("password_hash", user.PasswordHash).
		Set("display_name", user.DisplayName).
		Set("avatar_url", user.AvatarURL).
		Set("status", user.Status).
		Set("current_room_id", user.CurrentRoomID).
		Set("refresh_token", user.RefreshToken).
		Set("is_banned", user.IsBanned).
		Set("banned_at", user.BannedAt).
		Where(squirrel.Eq{"id": user.ID}).
		ExecContext(ctx)

	return err
}
