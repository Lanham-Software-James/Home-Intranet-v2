// Package library contains the library home intarnet functionality
package library

import (
	"fmt"
	"net/http"
)

func main() {

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// Run on localhost:8050
	http.ListenAndServe(":8050", h)
}
