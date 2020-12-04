package main

import (
	"os"
	"server/db"
)

var (
	user     = os.Getenv("MYSQL_ROOT_USERNAME")
	password = os.Getenv("MYSQL_ROOT_PASSWORD")
	endPoint = os.Getenv("ENDPOINT")
	database = os.Getenv("DATABASE")
)

func main() {
	_, err := db.New("mysql", db.GenerateDSN(user, password, endPoint, database), 0, 0)
	if err != nil {
		panic(err)
	}
}
