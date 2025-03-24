package db

import (
	"context"
	"database/sql"
)

func (db *Database) Query(ctx context.Context, query string) (*sql.Rows, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	return db.conn.QueryContext(ctx, query)
}
