package gologs

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// LogLevel represents the severity of a log message.
type LogLevel int

// Log levels.
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// Logger represents a simple logger with different log levels.
type Logger struct {
	logLevel LogLevel
	logger   *log.Logger
	output   io.Writer
}

// NewLogger creates a new Logger instance with the given log level and output.
func NewLogger(logLevel LogLevel, output io.Writer) *Logger {
	return &Logger{
		logLevel: logLevel,
		logger:   log.New(output, "", 0),
		output:   output,
	}
}

// setLogLevel sets the log level for the logger.
func (l *Logger) SetLogLevel(logLevel LogLevel) {
	l.logLevel = logLevel
}

func (l *Logger) log(level LogLevel, message interface{}) {
	if level < l.logLevel {
		return
	}

	entry := LogEntry{
		Level:     logLevelString(level),
		Timestamp: time.Now(),
		Message:   message,
	}

	entryJSON, err := json.Marshal(entry)
	if err != nil {
		log.Printf("Failed to marshal log entry: %v", err)
		return
	}

	_, err = l.output.Write(entryJSON)
	if err != nil {
		log.Printf("Failed to write log entry: %v", err)
		return
	}

	_, err = l.output.Write([]byte("\n"))
	if err != nil {
		log.Printf("Failed to write newline after log entry: %v", err)
	}
}

// Info logs an informational message.
func (l *Logger) Info(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	l.log(INFO, message)
}

// Debug logs a debug message.
func (l *Logger) Debug(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	l.log(DEBUG, message)
}

// Warn logs a warning message.
func (l *Logger) Warn(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	l.log(WARN, message)
}

// Error logs an error message.
func (l *Logger) Error(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	l.log(ERROR, message)
}

// Fatal logs a fatal message and exits the program.
func (l *Logger) Fatal(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	l.log(FATAL, message)
	os.Exit(1)
}

// CustomLogEntry represents a log entry that can be chained with level methods
type CustomLogEntry struct {
	logger  *Logger
	message interface{}
}

// Log accepts a message and returns a CustomLogEntry for method chaining
func (l *Logger) Log(message interface{}) *CustomLogEntry {
	return &CustomLogEntry{
		logger:  l,
		message: message,
	}
}

// Info logs the message at INFO level
func (c *CustomLogEntry) Info() {
	c.logger.log(INFO, c.message)
}

// Debug logs the message at DEBUG level
func (c *CustomLogEntry) Debug() {
	c.logger.log(DEBUG, c.message)
}

// Warn logs the message at WARN level
func (c *CustomLogEntry) Warn() {
	c.logger.log(WARN, c.message)
}

// Error logs the message at ERROR level
func (c *CustomLogEntry) Error() {
	c.logger.log(ERROR, c.message)
}

// Fatal logs the message at FATAL level and exits the program
func (c *CustomLogEntry) Fatal() {
	c.logger.log(FATAL, c.message)
	os.Exit(1)
}

// logLevelString converts a LogLevel to a string representation.
func logLevelString(logLevel LogLevel) string {
	switch logLevel {
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// LogLevelFromString converts a string to a LogLevel.
func LogLevelFromString(level string) LogLevel {
	switch level {
	case "INFO":
		return INFO
	case "DEBUG":
		return DEBUG
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return DEBUG
	}
}

type LogEntry struct {
	Level     string      `json:"level"`
	Timestamp time.Time   `json:"timestamp"`
	Message   interface{} `json:"message"`
}
