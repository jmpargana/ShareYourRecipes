package db

import (
	"context"
	"fmt"
	"time"
)

func (w *DBWrapper) withContext(
	f func(*DBWrapper, context.Context, string, ...interface{}) error,
	query string,
	args ...interface{},
) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return f(w, ctx, query, args...)
}

func performExecQueryWithPrepare(w *DBWrapper, ctx context.Context, query string, args ...interface{}) error {
	stmt, err := w.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed preparing context for query: %s with: %s", query, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, args...)
	return err
}
