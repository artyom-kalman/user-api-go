package app

import (
	"context"
	"fmt"

	"github.com/artyom-kalman/user-api-go/pkg/db"
)

func (u *User) Save(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	db := db.GetDatabase()

	query := fmt.Sprintf("INSERT INTO users (email, password) VALUES ('%s', '%s') RETURNING id", u.Email, u.Password)
	rows, err := db.Query(ctx, query)
	if err != nil || !rows.Next() {
		return fmt.Errorf("error inserting new user: %v", err)
	}

	err = rows.Scan(&u.ID)
	if err != nil {
		return fmt.Errorf("error scanning user ID: %v", err)
	}

	return nil
}
