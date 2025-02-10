// Package response contains the templates for building our responses to the user
package response

import (
	"encoding/json"
	"net/http"
)

// BadRequest is used to send a 500 response to the user
func BadRequest(w http.ResponseWriter, data interface{}) interface{} {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "data validation failed",
		"data":    &data,
	})
}
