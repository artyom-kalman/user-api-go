package repository

import (
	"context"
	"fmt"

	"github.com/artyom-kalman/user-api-go/internal/app/users"
)

func (r *UserRepository) isExists(u *users.User, ctx context.Context) bool {
	if ctx == nil {
		ctx = context.Background()
	}

	query := fmt.Sprintf("SELECT COUNT(*) FROM users WHERE id = %d", u.ID)
	rows, err := r.conn.QueryContext(ctx, query)
	if err != nil || !rows.Next() {
		return false
	}
	defer rows.Close()

	var count int
	err = rows.Scan(&count)
	if err != nil {
		return false
	}

	return count > 0
}
