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

func TestInsertRecipe(t *testing.T) {
	w, mock, cancel := setupMock()
	defer cancel()

	prep := mock.ExpectPrepare(insertRecipeQueryTest)
	prep.ExpectExec().WithArgs(r.ID, r.Private, r.Title, r.Ingridients, r.Time, r.Method).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := w.InsertRecipe(&r)
	checkNil(t, err)
}

func TestInsertRecipeFail(t *testing.T) {
	w, mock, cancel := setupMock()
	defer cancel()

	prep := mock.ExpectPrepare(insertRecipeQueryFailTest)
	prep.ExpectExec().WithArgs(r.ID, r.Private, r.Title, r.Ingridients, r.Time, r.Method).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := w.InsertRecipe(&r)
	shouldntBeNil(t, err)
}
