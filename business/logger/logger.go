// package logger handles structural logging for the application.
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.Logger
}

// func Init creates a new zap logger.
func Init() *Logger {
	config := zap.NewDevelopmentConfig()

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return &Logger{logger: logger}
}

// Info logs message on the info level.
func (l *Logger) Info(message string, fields ...zap.Field) {
	l.logger.Info(message, fields...)
}

// Error logs message on the error level.
func (l *Logger) Error(message string, fields ...zap.Field) {
	l.logger.Error(message, fields...)
}

// Panic logs message on the error level.
func (l *Logger) Panic(message string, fields ...zap.Field) {
	l.logger.Panic(message, fields...)
}

// Debug logs message on the error level.
func (l *Logger) Debug(message string, fields ...zap.Field) {
	l.logger.Debug(message, fields...)
}
