// Package main contains the library home intarnet functionality
package main

import (
	"Home-Intranet-v2-Backend/cmd/routers"
	"Home-Intranet-v2-Backend/internal/platform/config"

	"net/http"
)

func main() {
	router := routers.SetupRouter()
	http.ListenAndServe(config.GetServerHost(), router)
}
