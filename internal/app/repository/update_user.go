package repository

import (
	"context"
	"fmt"

	"github.com/artyom-kalman/user-api-go/internal/app/users"
	"github.com/artyom-kalman/user-api-go/pkg/logger"
)

func (r *UserRepository) Update(u *users.User, ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	if u.Email != "" {
		err := r.updateEmail(u, ctx)
		if err != nil {
			return fmt.Errorf("error updating user email: %v", err)
		}
	}

	if u.Password != "" {
		err := r.updatePassword(u, ctx)
		if err != nil {
			return fmt.Errorf("error updating user password: %v", err)
		}
	}

	return nil
}

func (r *UserRepository) updateEmail(u *users.User, ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	query := "UPDATE users SET email = $1 WHERE id = $2"
	logger.Debug("Executing query: %s with values %s %d", query, u.Email, u.ID)

	_, err := r.conn.QueryContext(ctx, query, u.Email, u.ID)
	if err != nil {
		return fmt.Errorf("error updating user email: %v", err)
	}

	return nil
}

func (r *UserRepository) updatePassword(u *users.User, ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	query := "UPDATE users SET password = $1 WHERE id = $2"
	logger.Debug("Executing query: %s with values %s %d", query, u.Password, u.ID)

	_, err := r.conn.QueryContext(ctx, query, u.Password, u.ID)
	if err != nil {
		return fmt.Errorf("error updating user password: %v", err)
	}

	return nil
}
