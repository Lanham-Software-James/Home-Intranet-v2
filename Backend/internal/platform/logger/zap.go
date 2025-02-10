// Package logger is the wrapper for our log tooling
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zap *zap.Logger
}

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

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.zap.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.zap.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.zap.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.zap.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.zap.Fatal(msg, fields...)
}

func (l *Logger) WithFields(fields ...zap.Field) *Logger {
	return &Logger{zap: l.zap.With(fields...)}
}

func (l *Logger) SetLevel(level zapcore.Level) {
	l.zap.Core().Enabled(level)
}

func (l *Logger) Sync() error {
	return l.zap.Sync()
}
