# go-logs

A simple, lightweight JSON logging library for Go applications with configurable log levels and structured output.

## Features

- **JSON output**: Structured logging with timestamps
- **Configurable output**: Write to any `io.Writer` (stdout, files, etc.)
- **Log level filtering**: Only output logs at or above the configured level
- **Type-safe**: Strongly typed log levels with string conversion utilities

## Installation

```bash
go get github.com/phasi/go-logs
```

## Quick Start

```go
package main

import (
    "os"
    gologs "github.com/phasi/go-logs"
)

func main() {
    // Create a new logger with INFO level, writing to stdout
    logger := gologs.NewLogger(gologs.INFO, os.Stdout)

    // Formatted logging
    logger.Info("Application started")
    logger.Warn("This is a warning")
    logger.Error("Something went wrong")

    // Fluent API for complex objects
    user := map[string]interface{}{
        "id": 123,
        "name": "John",
    }
    logger.Log(user).Info()
}
```

## Usage

### Creating a Logger

```go
import (
    "os"
    gologs "github.com/phasi/go-logs"
)

// Create logger with INFO level writing to stdout
logger := gologs.NewLogger(gologs.INFO, os.Stdout)

// Or write to a file
file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
if err != nil {
    panic(err)
}
defer file.Close()
logger := gologs.NewLogger(gologs.DEBUG, file)
```

### Log Levels

The library supports the following log levels (in ascending order of severity):

- `DEBUG` (0) - Detailed information for debugging
- `INFO` (1) - General information about application flow
- `WARN` (2) - Warning messages for potentially harmful situations
- `ERROR` (3) - Error messages for serious problems
- `FATAL` (4) - Critical errors that cause the application to exit

### Logging Messages

The library provides two main approaches for logging:

#### 1. Formatted Logging (Printf-style)

```go
// Simple messages
logger.Debug("Detailed debug information")
logger.Info("General information")
logger.Warn("Warning message")
logger.Error("Error occurred")
logger.Fatal("Critical error - application will exit")

// Formatted messages with variables
logger.Info("User %s logged in with ID %d", "john_doe", 12345)
logger.Debug("Processing %d items from batch %s", 150, "batch_001")
logger.Warn("Memory usage at %.1f%% (threshold: %d%%)", 85.7, 80)
logger.Error("Failed to connect to %s:%d - %v", "database.example.com", 5432, "timeout")
```

#### 2. Fluent API for Complex Objects

```go
// Simple messages
logger.Log("User logged in").Info()
logger.Log("Processing failed").Error()

// Complex objects (automatically serialized to JSON)
user := map[string]interface{}{
    "id": 123,
    "action": "login",
}
logger.Log(user).Info()
```

### Log Level Filtering

Only messages at or above the configured log level will be output:

```go
logger := gologs.NewLogger(gologs.WARN, os.Stdout)

logger.Debug("This won't be shown")  // Below WARN level
logger.Info("This won't be shown")   // Below WARN level
logger.Warn("This will be shown")    // At WARN level
logger.Error("This will be shown")   // Above WARN level
```

### Changing Log Level

```go
logger.SetLogLevel(gologs.ERROR)  // Now only ERROR and FATAL will be logged
```

### Log Level String Conversion

```go
// Convert log level to string
levelStr := gologs.logLevelString(gologs.INFO)  // Returns "INFO"

// Convert string to log level
level := gologs.LogLevelFromString("ERROR")     // Returns gologs.ERROR
```

### Output Format

All log messages are output as JSON with the following structure:

```json
{
  "level": "INFO",
  "timestamp": "2023-10-15T14:30:45.123456Z",
  "message": "Your log message here"
}
```

### Complex Messages

The logger accepts any type as a message:

```go
// String messages
logger.Info("Simple string message")

// Structured data
logger.Error(map[string]interface{}{
    "error": "database connection failed",
    "host": "localhost",
    "port": 5432,
})

// Custom types
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

user := User{ID: 123, Name: "John Doe"}
logger.Info(user)
```

### Fluent API with Log Method

The library also provides a fluent API through the `Log` method, which allows you to pass any object and then chain the log level:

```go
// Simple messages with fluent API
logger.Log("User logged in successfully").Info()
logger.Log("Processing failed").Error()
logger.Log("System shutting down").Fatal()

// Complex objects with fluent API
user := map[string]interface{}{
    "id":       123,
    "username": "john_doe",
    "action":   "login",
    "success":  true,
}
logger.Log(user).Info()

// Custom structs
type APIRequest struct {
    Method   string `json:"method"`
    Path     string `json:"path"`
    Status   int    `json:"status"`
    Duration int    `json:"duration_ms"`
}

request := APIRequest{
    Method:   "GET",
    Path:     "/api/users",
    Status:   200,
    Duration: 45,
}
logger.Log(request).Info()

// Arrays and slices
items := []string{"item1", "item2", "item3"}
logger.Log(items).Debug()

// Any data type
logger.Log(42).Debug()
logger.Log(3.14159).Warn()
logger.Log(true).Info()
```

The fluent API is particularly useful when you want to log complex objects without string formatting, as the object will be automatically serialized to JSON.

## Example Output

### Formatted Logging Output

```json
{"level":"INFO","timestamp":"2023-10-15T14:30:45.123456Z","message":"Application started"}
{"level":"WARN","timestamp":"2023-10-15T14:30:45.124567Z","message":"This is a warning"}
{"level":"ERROR","timestamp":"2023-10-15T14:30:45.125678Z","message":"Something went wrong"}
{"level":"INFO","timestamp":"2023-10-15T14:30:45.126789Z","message":"User john_doe logged in with ID 12345"}
```

### Fluent API Output

```json
{"level":"INFO","timestamp":"2023-10-15T14:30:45.127890Z","message":"User logged in"}
{"level":"INFO","timestamp":"2023-10-15T14:30:45.128901Z","message":{"id":123,"name":"John"}}
{"level":"INFO","timestamp":"2023-10-15T14:30:45.129012Z","message":{"method":"GET","path":"/api/users","status":200,"duration_ms":45}}
{"level":"DEBUG","timestamp":"2023-10-15T14:30:45.130123Z","message":["item1","item2","item3"]}
```
