package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) GetUserByID(ctx context.Context, userID int) (string, error) {
	var username string
	err := ur.db.QueryRow(ctx, "SELECT username FROM users WHERE id = $1", userID).Scan(&username)
	if err != nil {
		return "", fmt.Errorf("failed to get user by ID: %w", err)
	}
	return username, nil
}
