package repository

import (
	"context"
	"fmt"

	"github.com/artyom-kalman/user-api-go/internal/app/users"
	"github.com/artyom-kalman/user-api-go/pkg/logger"
)

func (r *UserRepository) Save(u *users.User, ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"
	logger.Debug("Running query: %s", query)

	rows, err := r.conn.QueryContext(ctx, query, u.Email, u.Password)
	if err != nil || !rows.Next() {
		return fmt.Errorf("error inserting new user: %v", err)
	}
	defer rows.Close()

	err = rows.Scan(&u.ID)
	if err != nil {
		return fmt.Errorf("error scanning user ID: %v", err)
	}
	logger.Debug("Id of created user: %d", u.ID)

	return nil
}
