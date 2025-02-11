// Package middlewares contains all of our custom defined or configured middleware for the go-chi router
package middlewares

import (
	"Home-Intranet-v2-Backend/internal/platform/config"
	"net/http"

	"github.com/go-chi/cors"
)

// SetupCors is used to configure the go-chi CORS middleware
func SetupCors() func(http.Handler) http.Handler {
	allowedHosts := config.GetAllowedHosts()
	return cors.Handler(cors.Options{
		AllowedOrigins:   []string{allowedHosts},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
}
