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
	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
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
