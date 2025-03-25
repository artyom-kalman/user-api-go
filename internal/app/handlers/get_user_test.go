package handlers

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/artyom-kalman/user-api-go/internal/app/repository"
	"github.com/stretchr/testify/assert"
)

func TestHandleGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	userRepo := repository.NewUserRepository(db)
	handler := NewUserHandler(userRepo)

	t.Run("got response, no user", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, email, password FROM users WHERE id = \\$1").
			WithArgs("1").
			WillReturnError(sql.ErrNoRows)

		req := httptest.NewRequest("GET", "http://localhost/users?id=1", nil)
		w := httptest.NewRecorder()
		handler.handleGetUser(w, req)

		resp := w.Result()
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

}
