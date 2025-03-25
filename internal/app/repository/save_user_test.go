package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/artyom-kalman/user-api-go/internal/app/users"
	"github.com/stretchr/testify/assert"
)

func TestSaveUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	assert.NoError(t, err)

	repo := NewUserRepository(db)

	t.Run("added new user", func(t *testing.T) {
		user := &users.User{Email: "john@example.com", Password: "hashed_password"}

		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(1)

		mock.ExpectQuery("INSERT INTO users \\(email, password\\) VALUES \\(\\$1, \\$2\\) RETURNING id").
			WithArgs(user.Email, user.Password).
			WillReturnRows(rows)

		err := repo.Save(user, context.Background())
		assert.Equal(t, 1, user.ID)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("handles duplicate email", func(t *testing.T) {
		user := &users.User{Email: "existing@example.com", Password: "hashed_password"}

		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(1)

		mock.ExpectQuery("INSERT INTO users \\(email, password\\) VALUES \\(\\$1, \\$2\\) RETURNING id").
			WithArgs(user.Email, user.Password).
			WillReturnRows(rows)

		err := repo.Save(user, context.Background())
		assert.NoError(t, err)

		mock.ExpectQuery("INSERT INTO users \\(email, password\\) VALUES \\(\\$1, \\$2\\) RETURNING id").
			WithArgs(user.Email, user.Password).
			WillReturnError(sql.ErrNoRows)

		err = repo.Save(user, context.Background())
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
