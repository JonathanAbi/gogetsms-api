package main

import (
	"gogetsms-api/handler"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Logger)

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	r.Route("/rentlist", func(r chi.Router) {
		r.Get("/", handler.GetAllRentList)
		r.Get("/filter", handler.GetFilteredRentList)
	})

	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
