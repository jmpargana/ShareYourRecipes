package db

import (
	"context"
	"fmt"
	"time"
)

func (w *DBWrapper) InsertRecipeTag(recipeID, tagID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create only if  not exists
	// if rows := w.FindRecipeTag(recipeID, tagID); !rows.Next() {
	stmt, err := w.db.PrepareContext(ctx, insertRecipeTagQuery)
	if err != nil {
		return fmt.Errorf("failed preparing context for query: %s with: %s", insertRecipeTagQuery, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, recipeID, tagID)
	return err
	// }
	// return nil
}

// func (w *DBWrapper) FindRecipeTag(recipeID, tagID int) *sql.Rows {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	rows, _ := w.db.QueryContext(ctx, selectRecipeTag, recipeID, tagID)
// 	return rows
// }
