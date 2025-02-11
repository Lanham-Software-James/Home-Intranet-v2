package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestSuccessResponse(t *testing.T) {
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
					data: "Operation successful",
				}
			},
			want1:    nil,
			wantCode: http.StatusOK,
			wantBody: map[string]interface{}{
				"message": "success",
				"data":    "Operation successful",
			},
		},
		{
			name: "Struct data",
			args: func(_ *testing.T) args {
				return args{
					w: httptest.NewRecorder(),
					data: struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
					}{
						ID:   1,
						Name: "John Doe",
					},
				}
			},
			want1:    nil,
			wantCode: http.StatusOK,
			wantBody: map[string]interface{}{
				"message": "success",
				"data": map[string]interface{}{
					"id":   float64(1), // JSON numbers are floats
					"name": "John Doe",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := SuccessResponse(tArgs.w, tArgs.data)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SuccessResponse got1 = %v, want1: %v", got1, tt.want1)
			}

			rec, ok := tArgs.w.(*httptest.ResponseRecorder)
			if !ok {
				t.Fatal("ResponseRecorder not found")
			}

			if rec.Code != tt.wantCode {
				t.Errorf("SuccessResponse status code = %v, want: %v", rec.Code, tt.wantCode)
			}

			if rec.Header().Get("Content-Type") != "application/json" {
				t.Errorf("SuccessResponse Content-Type = %v, want: application/json", rec.Header().Get("Content-Type"))
			}

			var gotBody map[string]interface{}
			if err := json.Unmarshal(rec.Body.Bytes(), &gotBody); err != nil {
				t.Fatalf("Failed to unmarshal response body: %v", err)
			}

			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("SuccessResponse body = %v, want: %v", gotBody, tt.wantBody)
			}
		})
	}
}
