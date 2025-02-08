// Package main contains the library home intarnet functionality
package main

import (
	"fmt"
	"net/http"
)

func main() {

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// Run on localhost:3000
	http.ListenAndServe(":3000", h)
}
