package db

import (
	"context"
	"fmt"
	"server/models"
	"time"
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
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		stmt, err := w.db.PrepareContext(ctx, insertIngridientQuery)
		if err != nil {
			return fmt.Errorf("failed preparing query: %s with: %s", insertIngridientQuery, err)
		}
		defer stmt.Close()

		if _, err := stmt.ExecContext(ctx, ingridient); err != nil {
			return fmt.Errorf("failed inserting ingridient with: %s", err)
		}

		ingridientID, err = w.FindIngridient(ingridient)
	} else if err != nil {
		return fmt.Errorf("nothing returned from select tag query: %s", err)
	}

	return w.InsertRecipeIngridient(id, ingridientID)
}

/// Insert Tag if non existent with NULL id and retrieve ID for recipetags table.
func (w *DBWrapper) FindIngridient(ingridient models.Ingridient) (id int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = w.db.QueryRowContext(ctx, selectIngridientByName, ingridient).
		Scan(&id)

	if err != nil {
		return 0, ErrNoID
	}

	return
}

func (w *DBWrapper) FindIngridientsByRecipeID(r *models.Recipe) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := w.db.QueryContext(ctx, selectIngridientsByRecipeID, r.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var ingridient string
		err := rows.Scan(&ingridient)
		if err != nil {
			return fmt.Errorf("failed scanning ingridient with: %s", err)
		}

		r.Ingridients = append(r.Ingridients, models.Ingridient(ingridient))
	}

	return nil
}
