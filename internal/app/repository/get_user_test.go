package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/artyom-kalman/user-api-go/internal/app/users"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	repo := NewUserRepository(db)

	t.Run("found", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "email", "password"}).
			AddRow(1, "john@example.com", "hashed_password")

		mock.ExpectQuery("SELECT id, email, password FROM users WHERE id = 1").
			WillReturnRows(rows)

		user, err := repo.GetUserById("1", context.Background())
		assert.NoError(t, err)
		assert.Equal(t, &users.User{ID: 1, Email: "john@example.com", Password: "hashed_password"}, user)
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, email, password FROM users WHERE id = 100").
			WillReturnError(sql.ErrNoRows)

		user, err := repo.GetUserById("100", context.Background())
		assert.Error(t, err)
		assert.Nil(t, user)
	})

	t.Run("negative id", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, email, password FROM users WHERE id = -1").
			WillReturnError(sql.ErrNoRows)

		user, err := repo.GetUserById("-1", context.Background())
		assert.Error(t, err)
		assert.Nil(t, user)
	})

	t.Run("not int", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, email, password FROM users WHERE id = 'abc'").
			WillReturnError(sql.ErrNoRows)

		user, err := repo.GetUserById("abc", context.Background())
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}
