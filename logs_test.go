package gologs

import (
	"bytes"
	"strings"
	"testing"
)

var logger *Logger
var buf bytes.Buffer // call buf.Reset() to clear buffer at the end of each test!

// test initilizes the logger
func TestInitLogger(t *testing.T) {
	logger = NewLogger(DEBUG, &buf)
	if logger == nil {
		t.Error("Expected logger to be initialized")
	}
	buf.Reset()
}

// tests debug log level
func TestDebug(t *testing.T) {
	logger.Debug("This is a debug message")
	output := buf.String()
	if !strings.Contains(output, "This is a debug message") {
		t.Errorf("Expected 'This is a debug message', got %v", output)
	}
	buf.Reset()
}

// tests info log level
func TestInfo(t *testing.T) {
	logger.Info("This is an info message")
	output := buf.String()
	if !strings.Contains(output, "This is an info message") {
		t.Errorf("Expected 'This is an info message', got %v", output)
	}
	buf.Reset()
}

// tests warn log level
func TestWarn(t *testing.T) {
	logger.Warn("This is a warning message")
	output := buf.String()
	if !strings.Contains(output, "This is a warning message") {
		t.Errorf("Expected 'This is a warning message', got %v", output)
	}
	buf.Reset()
}

// tests error log level
func TestError(t *testing.T) {
	logger.Error("This is an error message")
	output := buf.String()
	if !strings.Contains(output, "This is an error message") {
		t.Errorf("Expected 'This is an error message', got %v", output)
	}
	buf.Reset()
}

// tests fatal log level
func TestFatal(t *testing.T) {
	// We can't test os.Exit(1) directly in a unit test, so we'll test a modified version
	// For now, let's just test that Fatal logs with FATAL level
	// Note: This would need a more sophisticated test setup to properly test os.Exit()
	// For this test, we'll create a separate logger method that doesn't exit
	testLogger := NewLogger(DEBUG, &buf)
	testLogger.log(FATAL, "This is a fatal message")
	output := buf.String()
	if !strings.Contains(output, "This is a fatal message") {
		t.Errorf("Expected 'This is a fatal message', got %v", output)
	}
	if !strings.Contains(output, `"level":"FATAL"`) {
		t.Errorf("Expected FATAL level in output, got %v", output)
	}
	buf.Reset()
}

// tests debug log level with formatting
func TestDebugFormatting(t *testing.T) {
	logger.Debug("User %s has %d points", "John", 42)
	output := buf.String()
	if !strings.Contains(output, "User John has 42 points") {
		t.Errorf("Expected 'User John has 42 points', got %v", output)
	}
	buf.Reset()
}

// tests info log level with formatting
func TestInfoFormatting(t *testing.T) {
	logger.Info("Processing request %d from %s", 123, "192.168.1.1")
	output := buf.String()
	if !strings.Contains(output, "Processing request 123 from 192.168.1.1") {
		t.Errorf("Expected formatted message, got %v", output)
	}
	buf.Reset()
}

// tests warn log level with formatting
func TestWarnFormatting(t *testing.T) {
	logger.Warn("Memory usage at %.1f%% (threshold: %d%%)", 85.7, 80)
	output := buf.String()
	if !strings.Contains(output, "Memory usage at 85.7% (threshold: 80%)") {
		t.Errorf("Expected formatted warning, got %v", output)
	}
	buf.Reset()
}

// tests error log level with formatting
func TestErrorFormatting(t *testing.T) {
	logger.Error("Connection failed to %s:%d - %v", "localhost", 5432, "timeout")
	output := buf.String()
	if !strings.Contains(output, "Connection failed to localhost:5432 - timeout") {
		t.Errorf("Expected formatted error, got %v", output)
	}
	buf.Reset()
}

// test if log level filtering works
func TestLogLevelFilter(t *testing.T) {
	logger.SetLogLevel(INFO)
	logger.Debug("This is a debug message")
	output := buf.String()
	if strings.Contains(output, "This is a debug message") {
		t.Errorf("Expected 'This is a debug message' to be filtered out")
	}
	buf.Reset()
}

// tests setting log level to string
func TestLogLevelString(t *testing.T) {
	if logLevelString(DEBUG) != "DEBUG" {
		t.Errorf("Expected 'DEBUG', got %v", logLevelString(DEBUG))
	}
	if logLevelString(INFO) != "INFO" {
		t.Errorf("Expected 'INFO', got %v", logLevelString(INFO))
	}
	if logLevelString(WARN) != "WARN" {
		t.Errorf("Expected 'WARN', got %v", logLevelString(WARN))
	}
	if logLevelString(ERROR) != "ERROR" {
		t.Errorf("Expected 'ERROR', got %v", logLevelString(ERROR))
	}
	if logLevelString(FATAL) != "FATAL" {
		t.Errorf("Expected 'FATAL', got %v", logLevelString(FATAL))
	}
}

// tests setting log level from string
func TestLogLevelFromString(t *testing.T) {
	if LogLevelFromString("DEBUG") != DEBUG {
		t.Errorf("Expected Debug, got %v", LogLevelFromString("DEBUG"))
	}
	if LogLevelFromString("INFO") != INFO {
		t.Errorf("Expected Info, got %v", LogLevelFromString("INFO"))
	}
	if LogLevelFromString("WARN") != WARN {
		t.Errorf("Expected Warn, got %v", LogLevelFromString("WARN"))
	}
	if LogLevelFromString("ERROR") != ERROR {
		t.Errorf("Expected Error, got %v", LogLevelFromString("ERROR"))
	}
	if LogLevelFromString("FATAL") != FATAL {
		t.Errorf("Expected Fatal, got %v", LogLevelFromString("FATAL"))
	}
}
