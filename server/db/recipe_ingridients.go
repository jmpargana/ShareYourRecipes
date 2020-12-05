package db

func (w *DBWrapper) InsertRecipeIngridient(recipeID, ingridientID int) error {
	return w.execQueryWithPrepare(
		insertRecipeIngridientQuery,
		recipeID,
		ingridientID,
	)
}
