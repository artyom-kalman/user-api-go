package repository

import (
	"context"

	"github.com/artyom-kalman/user-api-go/internal/app/users"
	"github.com/artyom-kalman/user-api-go/pkg/logger"
)

func (r *UserRepository) isExists(u *users.User, ctx context.Context) bool {
	if ctx == nil {
		ctx = context.Background()
	}

	query := "SELECT COUNT() FROM users WHERE id = $1"
	logger.Debug("Executing query: %s with value %d", query, u.ID)

	rows, err := r.conn.QueryContext(ctx, query, u.ID)
	defer rows.Close()
	if err != nil || !rows.Next() {
		if err != nil {
			logger.Error("Error executing query: %s", err.Error())
		}
		logger.Debug("exit")
		return false
	}

	var count int
	err = rows.Scan(&count)
	if err != nil {
		logger.Error("Error scanning result: %s", err.Error())
		return false
	}

	return count > 0
}
