package main

import "server/db"

func main() {
	_, err := db.Setup()
	if err != nil {
		panic(err)
	}
}
