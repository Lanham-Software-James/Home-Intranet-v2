// Package logger is the wrapper for our log tooling
package logger

import (
	"Home-Intranet-v2-Backend/internal/platform/config"

	"go.uber.org/zap"
)

var logger, _ = newLogger()

func newLogger() (*zap.Logger, error) {
	var zapLogger *zap.Logger
	var err error

	production := config.GetProductionFlag()

	if production {
		zapLogger, err = zap.NewProduction()
	} else {
		zapLogger, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, err
	}

	return zapLogger, nil
}

// Debug log messages at the debug level
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
	sync()
}

// Info log messages at the info level
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
	sync()
}

// Warn log messages at the warn level
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
	sync()
}

// Error log messages at the error level
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
	sync()
}

// Fatal log messages at the fatal level
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
	sync()
}

func sync() error {
	return logger.Sync()
}
