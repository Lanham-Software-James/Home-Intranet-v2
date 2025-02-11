package middlewares

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-chi/cors"
)

func TestSetupCors(t *testing.T) {
	tests := []struct {
		name         string
		allowedHosts string
		want1        func(http.Handler) http.Handler
	}{
		{
			name:         "Default CORS setup",
			allowedHosts: "http://localhost:3000",
			want1: cors.Handler(cors.Options{
				AllowedOrigins:   []string{"http://localhost:3000"},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: false,
				MaxAge:           300,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("BACKEND_ALLOWED_HOSTS", tt.allowedHosts)
			got1 := SetupCors()

			// We can't directly compare functions, so we'll test the behavior
			testHandler := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})
			got1Handler := got1(testHandler)
			want1Handler := tt.want1(testHandler)

			// Create a test request
			req, _ := http.NewRequest("GET", "http://example.com", nil)
			req.Header.Set("Origin", "http://localhost:3000")

			// Create response recorders
			gotRec := httptest.NewRecorder()
			wantRec := httptest.NewRecorder()

			// Serve the request using both handlers
			got1Handler.ServeHTTP(gotRec, req)
			want1Handler.ServeHTTP(wantRec, req)

			// Compare the responses
			if !reflect.DeepEqual(gotRec.Header(), wantRec.Header()) {
				t.Errorf("SetupCors got headers = %v, want headers: %v", gotRec.Header(), wantRec.Header())
			}
		})
	}
}
