package db

import (
	"database/sql"
	"fmt"
	"server/models"
)

func (w *DBWrapper) appendTagsAndIngridients(r *models.Recipe) error {
	if err := w.FindIngridientsByRecipeID(r); err != nil {
		return err
	}

	return w.FindTagsByRecipeID(r)
}

func (w *DBWrapper) insertOrUpdateTagsAndIngridients(r *models.Recipe) error {
	if err := w.InsertTags(r.ID, r.Tags); err != nil {
		return err
	}
	return w.InsertIngridients(r.ID, r.Ingridients)
}

func (w *DBWrapper) scanRow(rows *sql.Rows) (r *models.Recipe, err error) {
	recipe := new(models.Recipe)
	var private int
	err = rows.Scan(
		&recipe.ID,
		&recipe.UserID,
		&recipe.Private,
		&recipe.Title,
		&recipe.Time,
		&recipe.Method,
	)

	if err != nil {
		return nil, fmt.Errorf("failed scanning recipe with: %s", err)
	}
	if private == 1 {
		recipe.Private = true
	}
	if err := w.appendTagsAndIngridients(recipe); err != nil {
		return nil, err
	}
	return
}
