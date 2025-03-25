package repository

import (
	"context"
	"fmt"

	"github.com/artyom-kalman/user-api-go/internal/app/users"
	"github.com/artyom-kalman/user-api-go/pkg/logger"
)

func (r *UserRepository) Delete(u *users.User, ctx context.Context) error {
	query := fmt.Sprintf("DELETE FROM users WHERE id = %d", u.ID)
	logger.Debug("Executing query: %s", query)

	if _, err := r.db.Query(ctx, query); err != nil {
		logger.Error("error deleting user with id = %d: %v", u.ID, err)
		return err
	}

	return nil
}
