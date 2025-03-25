package repository

import (
	"context"
	"fmt"

	"github.com/artyom-kalman/user-api-go/internal/app/users"
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

	query := fmt.Sprintf("UPDATE users SET email = '%s' WHERE id = %d", u.Email, u.ID)
	_, err := r.db.Query(ctx, query)
	if err != nil {
		return fmt.Errorf("error updating user email: %v", err)
	}

	return nil
}

func (r *UserRepository) updatePassword(u *users.User, ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	query := fmt.Sprintf("UPDATE users SET password = '%s' WHERE id = %d", u.Password, u.ID)
	_, err := r.db.Query(ctx, query)
	if err != nil {
		return fmt.Errorf("error updating user password: %v", err)
	}

	return nil
}
