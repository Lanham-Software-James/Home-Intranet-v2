package logger

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Test_newLogger(t *testing.T) {
	tests := []struct {
		name       string
		prodFlag   string
		want1      *zap.Logger
		wantErr    bool
		inspectErr func(err error, t *testing.T)
		inspectLog func(log *zap.Logger, t *testing.T)
	}{
		{
			name:     "Production Logger",
			prodFlag: "true",
			wantErr:  false,
			inspectLog: func(log *zap.Logger, t *testing.T) {
				if log.Core().Enabled(zap.DebugLevel) {
					t.Error("Production logger should not have debug level enabled")
				}
			},
		},
		{
			name:     "Development Logger",
			prodFlag: "false",
			wantErr:  false,
			inspectLog: func(log *zap.Logger, t *testing.T) {
				if !log.Core().Enabled(zap.DebugLevel) {
					t.Error("Development logger should have debug level enabled")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("BACKEND_PROD_FLAG", tt.prodFlag)

			got1, err := newLogger()

			if (err != nil) != tt.wantErr {
				t.Fatalf("newLogger error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}

			if !tt.wantErr && got1 == nil {
				t.Error("Expected non-nil logger, got nil")
			}

			if tt.inspectLog != nil && got1 != nil {
				tt.inspectLog(got1, t)
			}
		})
	}
}

func TestDebug(t *testing.T) {
	type args struct {
		msg    string
		fields []zap.Field
	}
	tests := []struct {
		name string
		args func(t *testing.T) args
		want string
	}{
		{
			name: "Simple debug message",
			args: func(_ *testing.T) args {
				return args{
					msg: "This is a debug message",
				}
			},
			want: "This is a debug message",
		},
		{
			name: "Debug message with fields",
			args: func(_ *testing.T) args {
				return args{
					msg: "Debug with fields",
					fields: []zap.Field{
						zap.String("key1", "value1"),
						zap.Int("key2", 42),
					},
				}
			},
			want: "Debug with fields",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			customLogger := zap.New(
				zapcore.NewCore(
					zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
					zapcore.AddSync(&buf),
					zapcore.DebugLevel,
				),
			)

			originalLogger := logger
			logger = customLogger
			defer func() { logger = originalLogger }()

			tArgs := tt.args(t)
			Debug(tArgs.msg, tArgs.fields...)

			t.Logf("Raw log output: %s", buf.String())

			if buf.Len() == 0 {
				t.Fatal("No log output captured")
			}

			var logEntry map[string]interface{}
			err := json.Unmarshal(buf.Bytes(), &logEntry)
			if err != nil {
				t.Fatalf("Failed to parse log output: %v", err)
			}

			// Check if the message is correct
			if msg, ok := logEntry["M"]; !ok || msg != tt.want {
				t.Errorf("Expected message %q, got %v", tt.want, msg)
			}

			// Check if the log level is debug
			if level, ok := logEntry["L"]; !ok || level != "DEBUG" {
				t.Errorf("Expected log level 'DEBUG', got %v", level)
			}

			// Check if fields are present (if any)
			for _, field := range tArgs.fields {
				if value, exists := logEntry[field.Key]; !exists {
					t.Errorf("Expected field %q not found in log output", field.Key)
				} else {
					switch field.Type {
					case zapcore.StringType:
						if value != field.String {
							t.Errorf("Expected field %q to have value %v, got %v", field.Key, field.String, value)
						}
					case zapcore.Int64Type:
						if int64(value.(float64)) != field.Integer {
							t.Errorf("Expected field %q to have value %v, got %v", field.Key, field.Integer, value)
						}
					default:
						t.Errorf("Unsupported field type for %q", field.Key)
					}
				}
			}
		})
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		msg    string
		fields []zap.Field
	}
	tests := []struct {
		name string
		args func(t *testing.T) args
		want string
	}{
		{
			name: "Simple info message",
			args: func(_ *testing.T) args {
				return args{
					msg: "This is an info message",
				}
			},
			want: "This is an info message",
		},
		{
			name: "Info message with fields",
			args: func(_ *testing.T) args {
				return args{
					msg: "Info with fields",
					fields: []zap.Field{
						zap.String("key1", "value1"),
						zap.Int("key2", 42),
					},
				}
			},
			want: "Info with fields",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			customLogger := zap.New(
				zapcore.NewCore(
					zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
					zapcore.AddSync(&buf),
					zapcore.InfoLevel,
				),
			)

			originalLogger := logger
			logger = customLogger
			defer func() { logger = originalLogger }()

			tArgs := tt.args(t)
			Info(tArgs.msg, tArgs.fields...)

			t.Logf("Raw log output: %s", buf.String())

			if buf.Len() == 0 {
				t.Fatal("No log output captured")
			}

			var logEntry map[string]interface{}
			err := json.Unmarshal(buf.Bytes(), &logEntry)
			if err != nil {
				t.Fatalf("Failed to parse log output: %v", err)
			}

			// Check if the message is correct
			if msg, ok := logEntry["M"]; !ok || msg != tt.want {
				t.Errorf("Expected message %q, got %v", tt.want, msg)
			}

			// Check if the log level is info
			if level, ok := logEntry["L"]; !ok || level != "INFO" {
				t.Errorf("Expected log level 'INFO', got %v", level)
			}

			// Check if fields are present (if any)
			for _, field := range tArgs.fields {
				if value, exists := logEntry[field.Key]; !exists {
					t.Errorf("Expected field %q not found in log output", field.Key)
				} else {
					switch field.Type {
					case zapcore.StringType:
						if value != field.String {
							t.Errorf("Expected field %q to have value %v, got %v", field.Key, field.String, value)
						}
					case zapcore.Int64Type:
						if int64(value.(float64)) != field.Integer {
							t.Errorf("Expected field %q to have value %v, got %v", field.Key, field.Integer, value)
						}
					default:
						t.Errorf("Unsupported field type for %q", field.Key)
					}
				}
			}
		})
	}
}

func TestWarn(t *testing.T) {
	type args struct {
		msg    string
		fields []zap.Field
	}
	tests := []struct {
		name string
		args func(t *testing.T) args
		want string
	}{
		{
			name: "Simple warn message",
			args: func(_ *testing.T) args {
				return args{
					msg: "This is a warning message",
				}
			},
			want: "This is a warning message",
		},
		{
			name: "Warn message with fields",
			args: func(_ *testing.T) args {
				return args{
					msg: "Warning with fields",
					fields: []zap.Field{
						zap.String("error_code", "W001"),
						zap.Int("affected_rows", 100),
					},
				}
			},
			want: "Warning with fields",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			customLogger := zap.New(
				zapcore.NewCore(
					zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
					zapcore.AddSync(&buf),
					zapcore.WarnLevel,
				),
			)

			originalLogger := logger
			logger = customLogger
			defer func() { logger = originalLogger }()

			tArgs := tt.args(t)
			Warn(tArgs.msg, tArgs.fields...)

			t.Logf("Raw log output: %s", buf.String())

			if buf.Len() == 0 {
				t.Fatal("No log output captured")
			}

			var logEntry map[string]interface{}
			err := json.Unmarshal(buf.Bytes(), &logEntry)
			if err != nil {
				t.Fatalf("Failed to parse log output: %v", err)
			}

			// Check if the message is correct
			if msg, ok := logEntry["M"]; !ok || msg != tt.want {
				t.Errorf("Expected message %q, got %v", tt.want, msg)
			}

			// Check if the log level is warn
			if level, ok := logEntry["L"]; !ok || level != "WARN" {
				t.Errorf("Expected log level 'WARN', got %v", level)
			}

			// Check if fields are present (if any)
			for _, field := range tArgs.fields {
				if value, exists := logEntry[field.Key]; !exists {
					t.Errorf("Expected field %q not found in log output", field.Key)
				} else {
					switch field.Type {
					case zapcore.StringType:
						if value != field.String {
							t.Errorf("Expected field %q to have value %v, got %v", field.Key, field.String, value)
						}
					case zapcore.Int64Type:
						if int64(value.(float64)) != field.Integer {
							t.Errorf("Expected field %q to have value %v, got %v", field.Key, field.Integer, value)
						}
					default:
						t.Errorf("Unsupported field type for %q", field.Key)
					}
				}
			}
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		msg    string
		fields []zap.Field
	}
	tests := []struct {
		name string
		args func(t *testing.T) args
		want string
	}{
		{
			name: "Simple error message",
			args: func(_ *testing.T) args {
				return args{
					msg: "This is an error message",
				}
			},
			want: "This is an error message",
		},
		{
			name: "Error message with fields",
			args: func(_ *testing.T) args {
				return args{
					msg: "Error with fields",
					fields: []zap.Field{
						zap.String("error_code", "E001"),
						zap.Int("status_code", 500),
						zap.Error(errors.New("sample error")),
					},
				}
			},
			want: "Error with fields",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			customLogger := zap.New(
				zapcore.NewCore(
					zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
					zapcore.AddSync(&buf),
					zapcore.ErrorLevel,
				),
			)

			originalLogger := logger
			logger = customLogger
			defer func() { logger = originalLogger }()

			tArgs := tt.args(t)
			Error(tArgs.msg, tArgs.fields...)

			t.Logf("Raw log output: %s", buf.String())

			if buf.Len() == 0 {
				t.Fatal("No log output captured")
			}

			var logEntry map[string]interface{}
			err := json.Unmarshal(buf.Bytes(), &logEntry)
			if err != nil {
				t.Fatalf("Failed to parse log output: %v", err)
			}

			// Check if the message is correct
			if msg, ok := logEntry["M"]; !ok || msg != tt.want {
				t.Errorf("Expected message %q, got %v", tt.want, msg)
			}

			// Check if the log level is error
			if level, ok := logEntry["L"]; !ok || level != "ERROR" {
				t.Errorf("Expected log level 'ERROR', got %v", level)
			}

			// Check if fields are present (if any)
			for _, field := range tArgs.fields {
				if value, exists := logEntry[field.Key]; !exists {
					t.Errorf("Expected field %q not found in log output", field.Key)
				} else {
					switch field.Type {
					case zapcore.StringType:
						if value != field.String {
							t.Errorf("Expected field %q to have value %v, got %v", field.Key, field.String, value)
						}
					case zapcore.Int64Type:
						if int64(value.(float64)) != field.Integer {
							t.Errorf("Expected field %q to have value %v, got %v", field.Key, field.Integer, value)
						}
					case zapcore.ErrorType:
						if value != field.Interface.(error).Error() {
							t.Errorf("Expected field %q to have value %v, got %v", field.Key, field.Interface.(error).Error(), value)
						}
					default:
						t.Errorf("Unsupported field type for %q", field.Key)
					}
				}
			}
		})
	}
}

func TestFatal(t *testing.T) {
	type args struct {
		msg    string
		fields []zap.Field
	}
	tests := []struct {
		name string
		args func(t *testing.T) args
		want string
	}{
		{
			name: "Simple fatal message",
			args: func(_ *testing.T) args {
				return args{
					msg: "This is a fatal error",
				}
			},
			want: "This is a fatal error",
		},
		{
			name: "Fatal message with fields",
			args: func(_ *testing.T) args {
				return args{
					msg: "Fatal error with fields",
					fields: []zap.Field{
						zap.String("error_code", "F001"),
						zap.Int("exit_code", 1),
						zap.Error(errors.New("critical error")),
					},
				}
			},
			want: "Fatal error with fields",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			customLogger := zap.New(
				zapcore.NewCore(
					zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
					zapcore.AddSync(&buf),
					zapcore.FatalLevel,
				),
				zap.WithFatalHook(zapcore.WriteThenPanic),
			)

			originalLogger := logger
			logger = customLogger
			defer func() {
				logger = originalLogger
				if r := recover(); r == nil {
					t.Errorf("The code did not panic")
				}
			}()

			tArgs := tt.args(t)
			Fatal(tArgs.msg, tArgs.fields...)

			t.Logf("Raw log output: %s", buf.String())

			if buf.Len() == 0 {
				t.Fatal("No log output captured")
			}

			var logEntry map[string]interface{}
			err := json.Unmarshal(buf.Bytes(), &logEntry)
			if err != nil {
				t.Fatalf("Failed to parse log output: %v", err)
			}

			// Check if the message is correct
			if msg, ok := logEntry["M"]; !ok || msg != tt.want {
				t.Errorf("Expected message %q, got %v", tt.want, msg)
			}

			// Check if the log level is fatal
			if level, ok := logEntry["L"]; !ok || level != "FATAL" {
				t.Errorf("Expected log level 'FATAL', got %v", level)
			}

			// Check if fields are present (if any)
			for _, field := range tArgs.fields {
				if value, exists := logEntry[field.Key]; !exists {
					t.Errorf("Expected field %q not found in log output", field.Key)
				} else {
					switch field.Type {
					case zapcore.StringType:
						if value != field.String {
							t.Errorf("Expected field %q to have value %v, got %v", field.Key, field.String, value)
						}
					case zapcore.Int64Type:
						if int64(value.(float64)) != field.Integer {
							t.Errorf("Expected field %q to have value %v, got %v", field.Key, field.Integer, value)
						}
					case zapcore.ErrorType:
						if value != field.Interface.(error).Error() {
							t.Errorf("Expected field %q to have value %v, got %v", field.Key, field.Interface.(error).Error(), value)
						}
					default:
						t.Errorf("Unsupported field type for %q", field.Key)
					}
				}
			}
		})
	}
}

func Test_sync(t *testing.T) {
	tests := []struct {
		name        string
		setupLogger func() *zap.Logger
		wantErr     bool
		inspectErr  func(error) bool
	}{
		{
			name: "Successful sync",
			setupLogger: func() *zap.Logger {
				return zap.NewNop()
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalLogger := logger
			logger = tt.setupLogger()
			defer func() { logger = originalLogger }()

			err := sync()

			t.Logf("sync() returned error: %v", err)

			if tt.wantErr != (err != nil) {
				t.Errorf("sync() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.inspectErr != nil && !tt.inspectErr(err) {
				t.Errorf("sync() unexpected error: %v", err)
			}
		})
	}
}
