// Package main contains the library home intarnet functionality
package main

import (
	"Home-Intranet-v2-Backend/cmd/routers"
	"os"

	"net/http"
)

func main() {
	router := routers.SetupRouter()
	http.ListenAndServe(os.Getenv("BACKEND_HOST"), router)
}
