package repository

import (
	"context"
	"database/sql"
)

type database interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

type UserRepository struct {
	conn database
}

func NewUserRepository(conn database) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}
