package db

import (
	"database/sql"
	"fmt"
)

func (w *DBWrapper) InsertRecipeTag(recipeID, tagID int) error {
	return w.execQueryWithPrepare(
		insertRecipeTagQuery,
		recipeID,
		tagID,
	)
}

func (w *DBWrapper) FindRecipeTagByTagID(tagID int) ([]int, error) {
	var ids []int

	err := w.queryRows(
		selectRecipeTagByTagID,

		func(rows *sql.Rows) error {
			for rows.Next() {
				var recipeid int
				if err := rows.Scan(&recipeid); err != nil {
					return fmt.Errorf("failed scanning with: %s", err)
				}
				ids = append(ids, recipeid)
			}
			return nil
		},
		tagID,
	)

	return ids, err
}
