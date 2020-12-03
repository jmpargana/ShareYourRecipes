package db

// func TestFindRecipeByID(t *testing.T) {
// 	w, mock, cancel := setupMock()
// 	defer cancel()

// 	rows := sqlmock.NewRows([]string{"id", "private", "title", "time", "method"}).
// 		AddRow(r.ID, r.Private, r.Title, r.Time, r.Method)

// 	mock.ExpectQuery(selectRecipeByIDQueryTest).WithArgs(r.ID).WillReturnRows(rows)

// 	recipe, err := w.FindRecipeByID(r.ID)
// 	checkNil(t, err)
// 	compareRecipes(t, recipe, &r)
// }

// func TestFindRecipeByIDFail(t *testing.T) {
// 	w, mock, cancel := setupMock()
// 	defer cancel()

// 	rows := sqlmock.NewRows([]string{"id", "private", "title", "time", "method"})

// 	mock.ExpectQuery(selectRecipeByIDQueryTest).WithArgs(r.ID).WillReturnRows(rows)

// 	recipe, err := w.FindRecipeByID(r.ID)

// 	recipeNotNil(t, recipe)
// 	shouldntBeNil(t, err)
// }

// func TestInsertRecipe(t *testing.T) {
// 	w, mock, cancel := setupMock()
// 	defer cancel()

// 	private := 0
// 	if r.Private {
// 		private = 1
// 	}

// 	prep := mock.ExpectPrepare(insertRecipeQueryTest)
// 	prep.ExpectExec().WithArgs(r.ID, r.UserID, private, r.Title, r.Time, r.Method).
// 		WillReturnResult(sqlmock.NewResult(0, 1))

// 	err := w.InsertRecipe(&r)
// 	checkNil(t, err)
// }

// func TestInsertRecipeFail(t *testing.T) {
// 	w, mock, cancel := setupMock()
// 	defer cancel()

// 	prep := mock.ExpectPrepare(insertRecipeQueryFailTest)
// 	prep.ExpectExec().WithArgs(r.ID, r.Private, r.Title, r.Time, r.Method).
// 		WillReturnResult(sqlmock.NewResult(0, 0))

// 	err := w.InsertRecipe(&r)
// 	shouldntBeNil(t, err)
// }

// func TestUpdateRecipe(t *testing.T) {
// 	w, mock, cancel := setupMock()
// 	defer cancel()

// 	prep := mock.ExpectPrepare(updateRecipeQueryTest)
// 	prep.ExpectExec().WithArgs(r.Private, r.Title, r.Time, r.Method, r.ID).
// 		WillReturnResult(sqlmock.NewResult(0, 1))

// 	err := w.UpdateRecipe(&r)
// 	checkNil(t, err)
// }

// func TestUpdateRecipeFail(t *testing.T) {
// 	w, mock, cancel := setupMock()
// 	defer cancel()

// 	prep := mock.ExpectPrepare(updateRecipeQueryFailTest)
// 	prep.ExpectExec().WithArgs(r.Private, r.Title, r.Time, r.Method, r.ID).
// 		WillReturnResult(sqlmock.NewResult(0, 0))

// 	err := w.UpdateRecipe(&r)
// 	shouldntBeNil(t, err)
// }

// func TestDeleteRecipeByID(t *testing.T) {
// 	w, mock, cancel := setupMock()
// 	defer cancel()

// 	prep := mock.ExpectPrepare(deleteRecipeByIDQueryTest)
// 	prep.ExpectExec().WithArgs(r.ID).WillReturnResult(sqlmock.NewResult(0, 1))

// 	err := w.DeleteRecipe(r.ID)
// 	checkNil(t, err)
// }

// func TestDeleteRecipeByIDFail(t *testing.T) {
// 	w, mock, cancel := setupMock()
// 	defer cancel()

// 	prep := mock.ExpectPrepare(deleteRecipeByIDQueryFailTest)
// 	prep.ExpectExec().WithArgs(r.ID).WillReturnResult(sqlmock.NewResult(0, 0))

// 	err := w.DeleteRecipe(r.ID)
// 	shouldntBeNil(t, err)
// }
