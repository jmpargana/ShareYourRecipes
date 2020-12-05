package db

import (
	"database/sql"
	"fmt"
	"server/models"
)

func (w *DBWrapper) InsertIngridients(id int, ingridients []models.Ingridient) error {
	for _, ingridient := range ingridients {
		if err := w.InsertIngridient(id, ingridient); err != nil {
			return fmt.Errorf("failed inserting tag: %s with: %s", ingridient, err)
		}
	}
	return nil
}

func (w *DBWrapper) InsertIngridient(id int, ingridient models.Ingridient) error {
	ingridientID, err := w.FindIngridient(ingridient)

	if err == ErrNoID {
		if err := w.execQueryWithPrepare(insertIngridientQuery, ingridient); err != nil {
			return err
		}
		ingridientID, err = w.FindIngridient(ingridient)
	}

	return w.InsertRecipeIngridient(id, ingridientID)
}

/// Insert Tag if non existent with NULL id and retrieve ID for recipetags table.
func (w *DBWrapper) FindIngridient(ingridient models.Ingridient) (id int, err error) {
	return w.queryID(
		selectIngridientByName,
		string(ingridient),
	)
}

func (w *DBWrapper) FindIngridientsByRecipeID(r *models.Recipe) error {
	return w.queryRows(
		selectIngridientsByRecipeID,

		func(rows *sql.Rows) error {
			for rows.Next() {
				var ingridient string
				if err := rows.Scan(&ingridient); err != nil {
					return err
				}
				r.Ingridients = append(r.Ingridients, models.Ingridient(ingridient))
			}
			return nil
		},
		r.ID,
	)
}
