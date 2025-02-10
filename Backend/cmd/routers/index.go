// Package routers provides all the details of our chi router.
package routers

import (
	"Home-Intranet-v2-Backend/internal/platform/response"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// RootRoutes is used to declare routes related to the application root
func RootRoutes(r *chi.Mux) {

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		response.SuccessResponse(w, "alive ok")
	})
}
