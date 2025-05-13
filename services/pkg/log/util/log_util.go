package util

import (
	"runtime"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a wrapper around zap logger
var Logger *zap.Logger

// Initialize sets up the zap logger with development configuration
func Initialize() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	logger, _ := config.Build()
	Logger = logger
}

// LogError logs an error message with a timestamp, file, and line number
func LogError(str string, location string, processID string) {
	// Get file, line number of where LogError was called
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}

	// Create a new logger if it hasn't been initialized
	if Logger == nil {
		Initialize()
	}

	// Add contextual fields
	fields := []zapcore.Field{
		zap.String("file", file),
		zap.Int("line", line),
		zap.String("timestamp", time.Now().Format("2006-01-02 15:04:05")),
	}

	// Add location and processID if provided
	if location != "" {
		fields = append(fields, zap.String("location", location))
		fields = append(fields, zap.String("processID", processID))
	}

	// Log the error with context
	Logger.Error(str, fields...)
}

// LogSuccess logs a success message with a timestamp
func LogSuccess(str string, location string, processID string) {
	// Create a new logger if it hasn't been initialized
	if Logger == nil {
		Initialize()
	}

	// Add contextual fields
	fields := []zapcore.Field{
		zap.String("timestamp", time.Now().Format("2006-01-02 15:04:05")),
	}

	// Add location and processID if provided
	if location != "" {
		fields = append(fields, zap.String("location", location))
		fields = append(fields, zap.String("processID", processID))
	}

	// Log the success message with context
	Logger.Info(str, fields...)
}

// LogWarn logs a warning message with a timestamp
func LogWarn(body string, location string, processID string) {
	// Create a new logger if it hasn't been initialized
	if Logger == nil {
		Initialize()
	}

	// Add contextual fields
	fields := []zapcore.Field{
		zap.String("timestamp", time.Now().Format("2006-01-02 15:04:05")),
	}

	// Add location and processID if provided
	if location != "" {
		fields = append(fields, zap.String("location", location))
		fields = append(fields, zap.String("processID", processID))
	}

	// Log the warning message with context
	Logger.Warn(body, fields...)
}

// LogTask logs a task message with a timestamp
func LogTask(str string, location string, processID string) {
	// Create a new logger if it hasn't been initialized
	if Logger == nil {
		Initialize()
	}

	// Add contextual fields
	fields := []zapcore.Field{
		zap.String("timestamp", time.Now().Format("2006-01-02 15:04:05")),
	}

	// Add location and processID if provided
	if location != "" {
		fields = append(fields, zap.String("location", location))
		fields = append(fields, zap.String("processID", processID))
	}

	// Log the task message with context (using Debug level)
	Logger.Debug(str, fields...)
}

// LogInfo logs an informational message with a timestamp
func LogInfo(str string, location string, processID string) {
	// Create a new logger if it hasn't been initialized
	if Logger == nil {
		Initialize()
	}

	// Add contextual fields
	fields := []zapcore.Field{
		zap.String("timestamp", time.Now().Format("2006-01-02 15:04:05")),
	}

	// Add location and processID if provided
	if location != "" {
		fields = append(fields, zap.String("location", location))
		fields = append(fields, zap.String("processID", processID))
	}

	// Log the info message with context
	Logger.Info(str, fields...)
}
