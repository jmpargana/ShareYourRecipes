package db

const (
	insertRecipeQuery = `
INSERT recipe
SET 
	ID = ?,
	Title = ?,
	Ingridients = ?,
	Method = ?,
	Time = ?,
	Private = ?,
	Tags = ?
`
	selectRecipeByID = `
SELECT *
FROM recipe
WHERE ID = ?
`
	selectRecipeByTags = `
SELECT * 
FROM recipe
JOIN tag ON id = id
WHERE tag.ID IN ('?')
`
	updatePrivateRecipeQuery = `
UPDATE recipe 
SET Private = ? 
WHERE ID = ?
`
)
