package db

const (
	// Recipes
	selectRecipeByIDQuery = "SELECT id, private, title, ingridients, time, method FROM recipes WHERE id = ?"
	selectAllRecipes      = "SELECT id, private, title, ingridients, time, method FROM recipes"
	insertRecipeQuery     = "INSERT INTO recipes (id, private, title, ingridients, time, method) VALUES (?, ?, ?, ?, ?, ?)"
	updateRecipeQuery     = "UPDATE recipes SET private = ?, title = ?, ingridients = ?, time = ?, method = ? WHERE id = ?"
	deleteRecipeByIDQuery = "DELETE FROM recipes WHERE id = ?"

	// Tags
	selectTagByName = "SELECT id FROM tags WHERE name = ?"
	insertTagQuery  = "INSERT INTO tags (id, name) VALUE (NULL, ?)"

	// RecipeTags
	insertRecipeTagQuery = "INSERT INTO recipetags (recipeid, tagid) VALUES (?, ?)"

	// Ingridients
	selectIngridientByName = "SELECT id FROM ingridients WHERE name = ?"
	insertIngridientQuery  = "INSERT INTO ingridients (id, name) VALUES (NULL, ?)"

	// RecipeIngridients
	insertRecipeIngridientQuery = "INSERT INTO recipeingridients (recipeid, ingridientID) VALUES (?, ?)"
)
