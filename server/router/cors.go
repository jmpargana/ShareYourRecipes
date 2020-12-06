package router

import (
	"github.com/rs/cors"
)

func createCorsWrapper() *cors.Cors {
	return cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
}
