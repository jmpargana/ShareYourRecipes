package db

import "server/models"

var (
	selectRecipeByIDQueryTest     = "SELECT id, private, title, ingridients, time, method FROM recipes WHERE id = \\?"
	insertRecipeQueryTest         = "INSERT INTO recipes \\(id, private, title, ingridients, time, method\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?\\)"
	insertRecipeQueryFailTest     = "INSERT INTO recipe \\(id, private, title, ingridients, time, method\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?\\)"
	updateRecipeQueryTest         = "UPDATE recipes SET private = \\?, title = \\?, ingridients = \\?, time = \\?, method = \\?  WHERE id = \\?"
	updateRecipeQueryFailTest     = "UPDATE recipe SET private = \\?, title = \\?, ingridients = \\?, time = \\?, method = \\?  WHERE id = \\?"
	deleteRecipeByIDQueryTest     = "DELETE FROM recipes WHERE id = \\?"
	deleteRecipeByIDQueryFailTest = "DELETE FROM recipe WHERE id = \\?"

	recipeTableTop = []string{"id", "private", "title", "ingridients", "time", "method"}
	r              = models.Recipe{
		ID:          1,
		Private:     1,
		Title:       "curry",
		Ingridients: "rice",
		Time:        30,
		Method:      "cook",
	}
)
