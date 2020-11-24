package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"server/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	username = os.Getenv("MYSQL_ROOT_USERNAME")
	password = os.Getenv("MYSQL_ROOT_PASSWORD")
	hostname = "127.0.0.1:3306"
	dbname   = "StoreAndShareYourRecipes"
)

type DB struct {
	*sql.DB
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func New(dataSourceName string) (*DB, error) {
	db, err := sql.Open("mysql", dsn(dbname))
	if err != nil {
		return nil, fmt.Errorf("Couldn't open DB: %s", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Couldn't reach DB: %s", err)
	}

	log.Printf("Opened MYSQL connection to %s\n", dsn(dbname))
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)

	return &DB{db}, nil
}

func (db *DB) Close() {
	db.Close()
}

func (db *DB) FetchRecipes() []models.Recipe {
	sel, err := db.Query("SELECT * FROM recipe")
	if err != nil {
		log.Fatalf("failed fetching recipes: %s", err)
		return nil
	}

	res := []models.Recipe{}

	for sel.Next() {
		var rec models.Recipe

		if err := sel.Scan(&rec); err != nil {
			log.Fatalf("failed scanning row: %s", err)
			return nil
		}

		res = append(res, rec)
	}

	return res
}

// This is the recipe for each query in the CRUD
func (db *DB) Insert(r *models.Recipe) (err error) {
	stmt, err := db.Prepare(insertRecipeQuery)
	if err != nil {
		return checkPrepareError(insertRecipeQuery, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		r.ID,
		r.Title,
		r.Ingridients,
		r.Method,
		r.Time,
		r.Private,
		r.Tags,
	)

	return
}

func (db *DB) SelectRecipeByID(ID int) (r *models.Recipe, err error) {
	stmt, err := db.Prepare(selectRecipeByID)
	if err != nil {
		return nil, checkPrepareError(selectRecipeByID, err)
	}
	defer stmt.Close()

	res, err := stmt.Query(ID)
	defer res.Close()

	if res.Next() {
		res.Scan(
			r.ID,
			r.Title,
			r.Ingridients,
			r.Method,
			r.Time,
			r.Private,
			r.Tags,
		)
	}
	return
}

func (db *DB) SelectRecipesByTags(tags []string) (r []*models.Recipe, err error) {
	stmt, err := db.Prepare(selectRecipeByTags)
	if err != nil {
		return nil, checkPrepareError(selectRecipeByID, err)
	}
	defer stmt.Close()

	res, err := stmt.Query(tags)
	defer res.Close()

	for res.Next() {
		var recipe *models.Recipe

		res.Scan(
			recipe.ID,
			recipe.Title,
			recipe.Ingridients,
			recipe.Method,
			recipe.Time,
			recipe.Private,
			recipe.Tags,
		)

		r = append(r, recipe)
	}
	return
}

func (db *DB) UpdatePrivateRecipe(ID, private int) error {
	stmt, err := db.Prepare(updatePrivateRecipeQuery)
	if err != nil {
		return checkPrepareError(updatePrivateRecipeQuery, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(ID, private)
	return err
}

func checkPrepareError(query string, err error) error {
	return fmt.Errorf("failed preparing query: %s: %s", query, err)
}
