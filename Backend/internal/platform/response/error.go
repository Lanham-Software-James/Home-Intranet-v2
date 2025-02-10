// Package response contains the templates for building our responses to the user
package response

import (
	"encoding/json"
	"net/http"
)

// InternalServerError is used to send a 500 response to the user
func InternalServerError(w http.ResponseWriter, data interface{}) interface{} {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "internal server error",
		"data":    &data,
	})
}
