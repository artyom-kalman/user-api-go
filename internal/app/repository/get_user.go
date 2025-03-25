package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/artyom-kalman/user-api-go/internal/app/users"
	"github.com/artyom-kalman/user-api-go/pkg/logger"
)

var ErrUserNotFound = errors.New("user not found")

func (r *UserRepository) GetUserById(id string, ctx context.Context) (*users.User, error) {
	if id == "" {
		return nil, errors.New("User ID is required")
	}

	if ctx == nil {
		ctx = context.Background()
	}

	query := "SELECT id, email, password FROM users WHERE id = $1"
	logger.Debug("Executing query: %s with value %s", query, id)

	rows, err := r.conn.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("error finding user with id = %s: %v", id, err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, ErrUserNotFound
	}

	var user users.User
	if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		return nil, fmt.Errorf("error scanning user data: %v", err)
	}

	return &user, nil
}
