package users

import (
	"context"
	"fmt"

	"github.com/artyom-kalman/user-api-go/pkg/db"
	"github.com/artyom-kalman/user-api-go/pkg/logger"
)

func (u *User) Delete(ctx context.Context) error {
	db := db.GetDatabase()

	query := fmt.Sprintf("DELETE FROM users WHERE id = %d", u.ID)
	logger.Debug("Executing query: %s", query)

	if _, err := db.Query(ctx, query); err != nil {
		logger.Error("error deleting user with id = %d: %v", u.ID, err)
		return err
	}

	return nil
}
