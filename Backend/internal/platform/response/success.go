// Package response contains the templates for building our responses to the user
package response

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse is used to send a 200 response to the user
func SuccessResponse(w http.ResponseWriter, data interface{}) interface{} {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "success",
		"data":    &data,
	})
}
