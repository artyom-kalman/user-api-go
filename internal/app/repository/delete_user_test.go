package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/artyom-kalman/user-api-go/internal/app/users"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	repo := NewUserRepository(db)

	t.Run("deleted", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "email", "password"}).
			AddRow(1, "john@example.com", "hashed_password")

		mock.ExpectQuery("DELETE FROM users WHERE id = 1").
			WillReturnRows(rows)

		user := &users.User{ID: 1, Email: "john@example.com", Password: "hashed_password"}
		err := repo.Delete(user, context.Background())
		assert.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery("DELETE FROM users WHERE id = 1").
			WillReturnError(sql.ErrNoRows)

		user := &users.User{ID: 1, Email: "john@example.com", Password: "hashed_password"}
		err := repo.Delete(user, context.Background())
		assert.Error(t, err)
	})
}
