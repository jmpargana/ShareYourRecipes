package db

import (
	"context"
	"time"
)

func (w *DBWrapper) InsertTags(id int, tags []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Check if tags exist

	// 2. Insert Tags with new IDs

	// 3. Associate IDs with recipe

	stmt, err := w.db.PrepareContext(ctx, insertRecipeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, &r.ID, &r.Private, &r.Title, &r.Time, &r.Method)

	return err
}
