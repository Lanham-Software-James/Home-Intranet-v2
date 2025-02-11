package routers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRootRoutes(t *testing.T) {
	type args struct {
		r *chi.Mux
	}
	tests := []struct {
		name string
		args func(t *testing.T) args
	}{
		{
			name: "Root route setup",
			args: func(_ *testing.T) args {
				return args{
					r: chi.NewRouter(),
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			RootRoutes(tArgs.r)

			// Test the root route
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			tArgs.r.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			// Check the response body
			expected := map[string]interface{}{
				"message": "success",
				"data":    "alive ok",
			}
			var got map[string]interface{}
			err = json.Unmarshal(rr.Body.Bytes(), &got)
			if err != nil {
				t.Fatalf("Failed to parse response body: %v", err)
			}
			if !reflect.DeepEqual(got, expected) {
				t.Errorf("handler returned unexpected body: got %v want %v",
					got, expected)
			}

			// Check the content type
			if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
				t.Errorf("content type header does not match: got %v want %v",
					ctype, "application/json")
			}
		})
	}
}
