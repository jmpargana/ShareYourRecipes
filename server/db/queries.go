package db

var (
	// Schame
	tables = []string{
		"CREATE TABLE IF NOT EXISTS recipes (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, userid INT, private INT, title TEXT, time INT, method TEXT)",
		"CREATE TABLE IF NOT EXISTS tags (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, name TEXT)",
		"CREATE TABLE IF NOT EXISTS recipetags (recipeid INT NOT NULL PRIMARY KEY, tagid INT NOT NULL)",
		"CREATE TABLE IF NOT EXISTS ingridients (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, name TEXT)",
		"CREATE TABLE IF NOT EXISTS recipeingridients (recipeid INT, ingridientid INT)",
	}
)

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
