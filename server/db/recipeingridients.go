package db

import (
	"context"
	"fmt"
	"time"
)

func (w *DBWrapper) InsertRecipeIngridient(recipeID, ingridientID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := w.db.PrepareContext(ctx, insertRecipeIngridientQuery)
	if err != nil {
		return fmt.Errorf("failed preparing context for query: %s with: %s", insertRecipeIngridientQuery, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, recipeID, ingridientID)
	return fmt.Errorf("failed inserting recipeingridient with: %s", err)
}
