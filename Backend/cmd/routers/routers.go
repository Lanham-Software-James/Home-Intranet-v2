// Package routers provides all the details of our chi router.
package routers

import (
	"Home-Intranet-v2-Backend/cmd/routers/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// SetupRouter is called to instantiate and attach all middleware and routes to the router
func SetupRouter() *chi.Mux {
	router := chi.NewRouter()

	registerMiddleware(router)
	registerRoutes(router)

	return router
}

func registerMiddleware(router *chi.Mux) {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middlewares.SetupCors())
}

func registerRoutes(router *chi.Mux) {
	RootRoutes(router)
	LibraryRoutes(router)
}
