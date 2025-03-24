package db

import (
	"context"
	"database/sql"
)

func (db *Database) Query(ctx context.Context, query string) (*sql.Rows, error) {
	err := db.conn.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	if ctx == nil {
		ctx = context.Background()
	}
	return db.conn.QueryContext(ctx, query)
}
