package db

const (
	selectRecipeByIDQuery = "SELECT id, private, title, ingridients, time, method FROM recipes WHERE id = ?"
	insertRecipeQuery     = "INSERT INTO recipes (id, private, title, ingridients, time, method) VALUES (?, ?, ?, ?, ?, ?)"
)
