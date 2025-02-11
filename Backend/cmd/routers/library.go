// Package routers provides all the details of our chi router.
package routers

import (
	"Home-Intranet-v2-Backend/cmd/handlers/library"
	"Home-Intranet-v2-Backend/internal/platform/logger"
	"Home-Intranet-v2-Backend/internal/platform/repository"

	"github.com/go-chi/chi/v5"
)

// LibraryRoutes is used to declare routes related to the application root
func LibraryRoutes(r *chi.Mux) {

	mongo, err := repository.Connect()
	if err != nil {
		logger.Fatal("Could not connect to database")
	}

	handler := library.Handler{
		Repository: &repository.Repository{
			Mongo: mongo,
		},
	}

	r.Route("/v1", func(r chi.Router) {

		r.Route("/books", func(r chi.Router) {
			r.Get("/", handler.ListBooks)
			r.Post("/", handler.CreateBook)
		})
	})
}
