package main

import (
	"net/http"
	"os"
	"server/db"
	"server/router"
)

var (
	user     = os.Getenv("MYSQL_ROOT_USERNAME")
	password = os.Getenv("MYSQL_ROOT_PASSWORD")
	endPoint = os.Getenv("ENDPOINT")
	database = os.Getenv("DATABASE")
)

func main() {
	dsn := db.GenerateDSN(user, password, endPoint, database)
	sql, err := db.New("mysql", dsn, 0, 0)
	if err != nil {
		panic(err)
	}

	r := router.New(sql)

	// TODO: dothe setup for HTTPS
	http.ListenAndServe(":8080", r)
}
