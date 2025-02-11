package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestBadRequest(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		data interface{}
	}
	tests := []struct {
		name     string
		args     func(t *testing.T) args
		want1    interface{}
		wantCode int
		wantBody map[string]interface{}
	}{
		{
			name: "Simple string data",
			args: func(_ *testing.T) args {
				return args{
					w:    httptest.NewRecorder(),
					data: "Invalid input",
				}
			},
			want1:    nil,
			wantCode: http.StatusBadRequest,
			wantBody: map[string]interface{}{
				"message": "data validation failed",
				"data":    "Invalid input",
			},
		},
		{
			name: "Struct data",
			args: func(_ *testing.T) args {
				return args{
					w: httptest.NewRecorder(),
					data: struct {
						Field string `json:"field"`
					}{
						Field: "Invalid",
					},
				}
			},
			want1:    nil,
			wantCode: http.StatusBadRequest,
			wantBody: map[string]interface{}{
				"message": "data validation failed",
				"data": map[string]interface{}{
					"field": "Invalid",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := BadRequest(tArgs.w, tArgs.data)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BadRequest got1 = %v, want1: %v", got1, tt.want1)
			}

			rec, ok := tArgs.w.(*httptest.ResponseRecorder)
			if !ok {
				t.Fatal("ResponseRecorder not found")
			}

			if rec.Code != tt.wantCode {
				t.Errorf("BadRequest status code = %v, want: %v", rec.Code, tt.wantCode)
			}

			if rec.Header().Get("Content-Type") != "application/json" {
				t.Errorf("BadRequest Content-Type = %v, want: application/json", rec.Header().Get("Content-Type"))
			}

			var gotBody map[string]interface{}
			if err := json.Unmarshal(rec.Body.Bytes(), &gotBody); err != nil {
				t.Fatalf("Failed to unmarshal response body: %v", err)
			}

			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("BadRequest body = %v, want: %v", gotBody, tt.wantBody)
			}
		})
	}
}
