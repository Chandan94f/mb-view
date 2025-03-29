package logger

import (
	"log"
)

// Logger interface for logging
type Logger interface {
	Info(message string)
	Error(message string)
	Debug(message string)
}

// BasicLogger is a simple implementation of the Logger interface
type BasicLogger struct{}

// Info logs informational messages
func (l *BasicLogger) Info(message string) {
	log.Printf("[INFO]: %s", message)
}

// Error logs error messages
func (l *BasicLogger) Error(message string) {
	log.Printf("[ERROR]: %s", message)
}

// Debug logs debug messages
func (l *BasicLogger) Debug(message string) {
	log.Printf("[DEBUG]: %s", message)
}

// NewBasicLogger creates and returns a new BasicLogger
func NewBasicLogger() Logger {
	return &BasicLogger{}
}
