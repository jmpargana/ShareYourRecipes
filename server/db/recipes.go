package db

import (
	"context"
	"fmt"
	"server/models"
	"time"
)

func (w *DBWrapper) FindRecipeByID(id int) (*models.Recipe, error) {
	r := new(models.Recipe)
	r.ID = id

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var private int
	err := w.db.QueryRowContext(ctx, selectRecipeByIDQuery, id).
		Scan(&r.UserID, &private, &r.Title, &r.Time, &r.Method)

	if err != nil {
		return nil, err
	}

	if private == 1 {
		r.Private = true
	}

	if err := w.FindIngridientsByRecipeID(r); err != nil {
		return nil, err
	}

	if err := w.FindTagsByRecipeID(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (w *DBWrapper) FindRecipesByTags(tags []models.Tag) (r []*models.Recipe, err error) {
	for _, tag := range tags {
		recipes, err := w.FindRecipesByTag(tag)
		if err != nil {
			return nil, fmt.Errorf("failed fetching recipes by tags with tag: %s with: %s", tag, err)
		}
		r = append(r, recipes...)
	}

	return
}

func (w *DBWrapper) FindAllUsersRecipes(userID int) ([]*models.Recipe, error) {
	return nil, nil
}

func (w *DBWrapper) FindRecipesByTag(tag models.Tag) ([]*models.Recipe, error) {
	return nil, nil
}

func (w *DBWrapper) FindRecipesByIngridient(ingridient models.Ingridient) ([]*models.Recipe, error) {
	return nil, nil
}

func (w *DBWrapper) FindAllPublicRecipes() (r []*models.Recipe, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := w.db.QueryContext(ctx, selectAllRecipes)
	if err != nil {
		return nil, fmt.Errorf("failed fetching all recipes with: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
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

		if err := w.FindIngridientsByRecipeID(recipe); err != nil {
			return nil, err
		}

		if err := w.FindTagsByRecipeID(recipe); err != nil {
			return nil, err
		}

		r = append(r, recipe)
	}
	return
}

func (w *DBWrapper) InsertRecipe(r *models.Recipe) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := w.InsertTags(r.ID, r.Tags); err != nil {
		return fmt.Errorf("failed creating tags for recipe with: %s", err)
	}

	if err := w.InsertIngridients(r.ID, r.Ingridients); err != nil {
		return fmt.Errorf("failed creating tags for recipe with: %s", err)
	}

	stmt, err := w.db.PrepareContext(ctx, insertRecipeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, r.ID, r.UserID, r.Private, r.Title, r.Time, r.Method)

	return err
}

func (w *DBWrapper) UpdateRecipe(r *models.Recipe) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := w.db.PrepareContext(ctx, updateRecipeQuery)
	if err != nil {
		return fmt.Errorf("failed updating recipe: %v with: %s", r, err)
	}

	_, err = stmt.ExecContext(ctx, &r.Private, &r.Title, &r.Time, &r.Method, &r.ID)
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
