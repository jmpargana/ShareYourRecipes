package router

import (
	"net/http"
	"server/db"

	"github.com/gorilla/mux"
)

var sql *db.DBWrapper

func New(database *db.DBWrapper) http.Handler {
	jwtMiddleware := createJWTMiddleware()
	corsWrapper := createCorsWrapper()
	sql = database

	r := mux.NewRouter()

	// API end points
	r.Handle("/fetch", jwtMiddleware.Handler(http.HandlerFunc(GetRecipes))).Methods("GET")
	r.Handle("/upload", jwtMiddleware.Handler(http.HandlerFunc(CreateRecipe))).Methods("POST")
	r.Handle("/update", jwtMiddleware.Handler(http.HandlerFunc(UpdateRecipe))).Methods("POST")
	r.Handle("/delete", jwtMiddleware.Handler(http.HandlerFunc(DeleteRecipe))).Methods("GET")

	// React App

	return corsWrapper.Handler(r)
}
