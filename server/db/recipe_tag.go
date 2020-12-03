package db

import (
	"context"
	"fmt"
	"time"
)

func (w *DBWrapper) InsertRecipeTag(recipeID, tagID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := w.db.PrepareContext(ctx, insertRecipeTagQuery)
	if err != nil {
		return fmt.Errorf("failed preparing context for query: %s with: %s", insertRecipeTagQuery, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, recipeID, tagID)
	return fmt.Errorf("failed inserting recipetag with: %s", err)
}
