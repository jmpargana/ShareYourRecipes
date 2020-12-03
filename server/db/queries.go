package db

const (
	// Recipes
	selectRecipeByIDQuery = "SELECT id, private, title, ingridients, time, method FROM recipes WHERE id = ?"
	selectAllRecipes      = "SELECT id, private, title, ingridients, time, method FROM recipes WHERE private = 0"
	insertRecipeQuery     = "INSERT INTO recipes (id, private, title, ingridients, time, method) VALUES (?, ?, ?, ?, ?, ?)"
	updateRecipeQuery     = "UPDATE recipes SET private = ?, title = ?, ingridients = ?, time = ?, method = ? WHERE id = ?"
	deleteRecipeByIDQuery = "DELETE FROM recipes WHERE id = ?"

	// Tags
	selectTagByName      = "SELECT id FROM tags WHERE name = ?"
	selectTagsByRecipeID = "SELECT name FROM tags JOIN recipetags ON id=tagid WHERE recipeid = ?"
	insertTagQuery       = "INSERT INTO tags (id, name) VALUE (NULL, ?)"

	// RecipeTags
	insertRecipeTagQuery = "INSERT INTO recipetags (recipeid, tagid) VALUES (?, ?)"

	// Ingridients
	selectIngridientByName      = "SELECT id FROM ingridients WHERE name = ?"
	selectIngridientsByRecipeID = "SELECT name FROM ingridients JOIN recipeingridients ON id=ingridientid WHERE recipeid = ?"
	insertIngridientQuery       = "INSERT INTO ingridients (id, name) VALUES (NULL, ?)"

	// RecipeIngridients
	insertRecipeIngridientQuery = "INSERT INTO recipeingridients (recipeid, ingridientID) VALUES (?, ?)"
)
