package db

import "server/models"

var (
	selectRecipeByIDQueryTest     = "SELECT id, private, title, time, method FROM recipes WHERE id = \\?"
	insertRecipeQueryTest         = "INSERT INTO recipes \\(id, private, title, time, method\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?\\)"
	insertRecipeQueryFailTest     = "INSERT INTO recipe \\(id, private, title, time, method\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?\\)"
	updateRecipeQueryTest         = "UPDATE recipes SET private = \\?, title = \\?, time = \\?, method = \\?  WHERE id = \\?"
	updateRecipeQueryFailTest     = "UPDATE recipe SET private = \\?, title = \\?, time = \\?, method = \\?  WHERE id = \\?"
	deleteRecipeByIDQueryTest     = "DELETE FROM recipes WHERE id = \\?"
	deleteRecipeByIDQueryFailTest = "DELETE FROM recipe WHERE id = \\?"

	recipeTableTop = []string{"id", "private", "title", "ingridients", "time", "method"}
	r              = models.Recipe{
		ID:          1,
		Private:     1,
		Title:       "curry",
		Ingridients: []string{"rice"},
		Time:        30,
		Method:      "cook",
		Tags:        []string{"indian", "curry"},
	}
)
