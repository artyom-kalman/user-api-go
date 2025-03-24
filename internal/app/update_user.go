package app

import (
	"context"
	"fmt"

	"github.com/artyom-kalman/user-api-go/pkg/db"
)

func (u *User) Update(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	if u.Email != "" {
		err := u.updateEmail(ctx)
		if err != nil {
			return fmt.Errorf("error updating user email: %v", err)
		}
	}

	if u.Password != "" {
		err := u.updatePassword(ctx)
		if err != nil {
			return fmt.Errorf("error updating user password: %v", err)
		}
	}

	return nil
}

func (u *User) updateEmail(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	db := db.GetDatabase()

	query := fmt.Sprintf("UPDATE users SET email = '%s' WHERE id = %d", u.Email, u.ID)
	_, err := db.Query(ctx, query)
	if err != nil {
		return fmt.Errorf("error updating user email: %v", err)
	}

	return nil
}

func (u *User) updatePassword(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	db := db.GetDatabase()

	query := fmt.Sprintf("UPDATE users SET password = '%s' WHERE id = %d", u.Password, u.ID)
	_, err := db.Query(ctx, query)
	if err != nil {
		return fmt.Errorf("error updating user password: %v", err)
	}

	return nil
}
