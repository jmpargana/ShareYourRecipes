package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFindRecipeByID(t *testing.T) {
	w, mock, cancel := setupMock()
	defer cancel()

	rows := sqlmock.NewRows([]string{"id", "private", "title", "ingridients", "time", "method"}).
		AddRow(r.ID, r.Private, r.Title, r.Ingridients, r.Time, r.Method)

	mock.ExpectQuery(selectRecipeByIDQueryTest).WithArgs(r.ID).WillReturnRows(rows)

	recipe, err := w.FindRecipeByID(r.ID)
	checkNil(t, err)
	compareRecipes(t, recipe, &r)
}

func TestFindRecipeByIDFail(t *testing.T) {
	w, mock, cancel := setupMock()
	defer cancel()

	rows := sqlmock.NewRows([]string{"id", "private", "title", "ingridients", "time", "method"})

	mock.ExpectQuery(selectRecipeByIDQueryTest).WithArgs(r.ID).WillReturnRows(rows)

	recipe, err := w.FindRecipeByID(r.ID)

	recipeNotNil(t, recipe)
	shouldntBeNil(t, err)
}
