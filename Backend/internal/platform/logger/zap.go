// Package logger is the wrapper for our log tooling
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zap *zap.Logger
}

var logger, _ = NewLogger(false)

func NewLogger(production bool) (*Logger, error) {
	var zapLogger *zap.Logger
	var err error

	if production {
		zapLogger, err = zap.NewProduction()
	} else {
		zapLogger, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, err
	}

	return &Logger{zap: zapLogger}, nil
}

func Debug(msg string, fields ...zap.Field) {
	logger.zap.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.zap.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.zap.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.zap.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.zap.Fatal(msg, fields...)
}

func WithFields(fields ...zap.Field) *Logger {
	return &Logger{zap: logger.zap.With(fields...)}
}

func SetLevel(level zapcore.Level) {
	logger.zap.Core().Enabled(level)
}

func Sync() error {
	return logger.zap.Sync()
}
