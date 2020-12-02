package db

const (
	selectRecipeByIDQuery = "SELECT id, private, title, ingridients, time, method FROM recipes WHERE id = ?"
	selectAllRecipes      = "SELECT id, private, title, ingridients, time, method FROM recipes"
	insertRecipeQuery     = "INSERT INTO recipes (id, private, title, ingridients, time, method) VALUES (?, ?, ?, ?, ?, ?)"
	updateRecipeQuery     = "UPDATE recipes SET private = ?, title = ?, ingridients = ?, time = ?, method = ? WHERE id = ?"
	deleteRecipeByIDQuery = "DELETE FROM recipes WHERE id = ?"
)
