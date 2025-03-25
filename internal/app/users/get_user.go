package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/artyom-kalman/user-api-go/pkg/db"
	"github.com/artyom-kalman/user-api-go/pkg/logger"
)

var ErrUserNotFound = errors.New("user not found")

func GetUserById(ctx context.Context, userId string) (*User, error) {
	if userId == "" {
		return nil, errors.New("User ID is required")
	}

	if ctx == nil {
		ctx = context.Background()
	}

	db := db.GetDatabase()

	query := fmt.Sprintf("SELECT id, email, password FROM users WHERE id = %s", userId)
	logger.Debug("Executing query: %s", query)

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error finding user with id = %s: %v", userId, err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, ErrUserNotFound
	}

	var user User
	if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		return nil, fmt.Errorf("error scanning user data: %v", err)
	}

	return &user, nil
}
