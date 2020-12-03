package db

import "server/models"

var (
	selectRecipeByIDQueryTest     = "SELECT id, private, title, time, method FROM recipes WHERE id = \\?"
	insertRecipeQueryTest         = "INSERT INTO recipes \\(id, userid, private, title, time, method\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?\\)"
	insertRecipeQueryFailTest     = "INSERT INTO recipe \\(id, private, title, time, method\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?\\)"
	updateRecipeQueryTest         = "UPDATE recipes SET private = \\?, title = \\?, time = \\?, method = \\?  WHERE id = \\?"
	updateRecipeQueryFailTest     = "UPDATE recipe SET private = \\?, title = \\?, time = \\?, method = \\?  WHERE id = \\?"
	deleteRecipeByIDQueryTest     = "DELETE FROM recipes WHERE id = \\?"
	deleteRecipeByIDQueryFailTest = "DELETE FROM recipe WHERE id = \\?"

	recipeTableTop = []string{"id", "userid", "private", "title", "time", "method"}

	// Mock Recipe
	r = models.Recipe{
		ID:          1,
		UserID:      1,
		Private:     true,
		Title:       "curry",
		Ingridients: []models.Ingridient{"rice"},
		Time:        30,
		Method:      "cook",
		Tags:        []models.Tag{"indian", "curry"},
	}
)
