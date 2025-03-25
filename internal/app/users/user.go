package users

import (
	"context"
	"fmt"

	"github.com/artyom-kalman/user-api-go/pkg/db"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) isExists(ctx context.Context) bool {
	if ctx == nil {
		ctx = context.Background()
	}

	db := db.GetDatabase()

	query := fmt.Sprintf("SELECT COUNT(*) FROM users WHERE id = %d", u.ID)
	rows, err := db.Query(ctx, query)
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
