package repository

import (
	"context"
	"database/sql"
)

type database interface {
	Query(ctx context.Context, query string) (*sql.Rows, error)
}

type UserRepository struct {
	conn database
}

func NewUserRepository(conn database) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}
