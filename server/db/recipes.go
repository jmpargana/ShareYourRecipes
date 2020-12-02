package db

import (
	"context"
	"fmt"
	"log"
	"server/models"
	"time"
)

func (w *DBWrapper) FindRecipeByID(id int) (*models.Recipe, error) {
	r := new(models.Recipe)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := w.db.QueryRowContext(ctx, selectRecipeByIDQuery, id).
		Scan(&r.ID, &r.Private, &r.Title, &r.Ingridients, &r.Time, &r.Method)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (w *DBWrapper) FindRecipesByTags(tags []string) (r []*models.Recipe, err error) {
	for _, tag := range tags {
		recipes, err := w.FindRecipesByTag(tag)
		if err != nil {
			return nil, fmt.Errorf("failed fetching recipes by tags with tag: %s with: %s", tag, err)
		}
		r = append(r, recipes...)
	}

	return
}

func (w *DBWrapper) FindRecipesByTag(tag string) ([]*models.Recipe, error) {
	return nil, nil
}

func (w *DBWrapper) FindAllRecipes() (r []*models.Recipe, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := w.db.QueryContext(ctx, selectAllRecipes)
	if err != nil {
		return nil, fmt.Errorf("failed fetching all recipes with: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		recipe := new(models.Recipe)
		err = rows.Scan(
			&recipe.ID,
			&recipe.Private,
			&recipe.Title,
			&recipe.Ingridients,
			&recipe.Time,
			&recipe.Method,
		)

		if err != nil {
			return nil, fmt.Errorf("failed scanning recipe with: %s", err)
		}

		r = append(r, recipe)
	}
	return
}

func (w *DBWrapper) InsertRecipe(r *models.Recipe) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := w.db.PrepareContext(ctx, insertRecipeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, &r.ID, &r.Private, &r.Title, &r.Ingridients, &r.Time, &r.Method)

	log.Println(result)
	return err
}

func (w *DBWrapper) UpdateRecipe(r *models.Recipe) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := w.db.PrepareContext(ctx, updateRecipeQuery)
	if err != nil {
		return fmt.Errorf("failed updating recipe: %v with: %s", r, err)
	}

	_, err = stmt.ExecContext(ctx, &r.Private, &r.Title, &r.Ingridients, &r.Time, &r.Method, &r.ID)
	return err
}

func (w *DBWrapper) DeleteRecipe(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := w.db.PrepareContext(ctx, deleteRecipeByIDQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	return err
}
