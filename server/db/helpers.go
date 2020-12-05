package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (w *DBWrapper) execQueryWithPrepare(query string, args ...interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := w.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed preparing context for query: %s with: %s", query, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, args...)
	return err
}

func (w *DBWrapper) queryRows(query string, f func(*sql.Rows) error, args ...interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := w.db.QueryContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed querying: %s with: %s", query, err)
	}

	return f(rows)
}

func (w *DBWrapper) queryID(query, name string) (id int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = w.db.QueryRowContext(ctx, query, name).Scan(&id); err != nil {
		return 0, ErrNoID
	}
	return
}
