package repository

import (
	"context"
	"fmt"

	"github.com/artyom-kalman/user-api-go/internal/app/users"
)

func (r *UserRepository) Save(u *users.User, ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	query := fmt.Sprintf("INSERT INTO users (email, password) VALUES ('%s', '%s') RETURNING id", u.Email, u.Password)
	rows, err := r.conn.Query(ctx, query)
	if err != nil || !rows.Next() {
		return fmt.Errorf("error inserting new user: %v", err)
	}

	err = rows.Scan(&u.ID)
	if err != nil {
		return fmt.Errorf("error scanning user ID: %v", err)
	}

	return nil
}
