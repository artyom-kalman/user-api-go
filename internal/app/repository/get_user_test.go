package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	repo := NewUserRepository(db)

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "email"}).
			AddRow(1, "John Doe", "john@example.com")

		mock.ExpectQuery("SELECT id, name, email FROM users WHERE id = ?").
			WithArgs(1).
			WillReturnRows(rows)

		user, err := repo.GetByID(1)
		assert.NoError(t, err)
		assert.Equal(t, &domain.User{ID: 1, Name: "John Doe", Email: "john@example.com"}, user)
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, name, email FROM users WHERE id = ?").
			WithArgs(2).
			WillReturnError(sql.ErrNoRows)

		user, err := repo.GetByID(2)
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}
