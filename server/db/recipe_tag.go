package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (w *DBWrapper) InsertRecipeTag(recipeID, tagID int) error {
	return w.withContext(
		performExecQueryWithPrepare,
		insertRecipeTagQuery,
		recipeID,
		tagID,
	)
}

// func (w *DBWrapper) InsertRecipeTag(recipeID, tagID int) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	stmt, err := w.db.PrepareContext(ctx, insertRecipeTagQuery)
// 	if err != nil {
// 		return fmt.Errorf("failed preparing context for query: %s with: %s", insertRecipeTagQuery, err)
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.ExecContext(ctx, recipeID, tagID)
// 	return err
// }

func (w *DBWrapper) FindRecipeTagByTagID(tagID int) ([]int, error) {
	var ids []int

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := w.db.QueryContext(ctx, selectRecipeTagByTagID, tagID)
	if err != nil {
		return nil, fmt.Errorf("failed fetching recipetags by tagid with: %s", err)
	}

	for rows.Next() {
		var recipeid int
		if err := rows.Scan(&recipeid); err != nil {
			return nil, fmt.Errorf("failed scanning recipetags with %s", err)
		}
		ids = append(ids, recipeid)
	}

	return ids, nil
}

func (w *DBWrapper) FindRecipeTag(recipeID, tagID int) *sql.Row {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return w.db.QueryRowContext(ctx, selectRecipeTag, recipeID, tagID)
}
