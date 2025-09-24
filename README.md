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

    // Log some messages
    logger.Info("Application started")
    logger.Warn("This is a warning")
    logger.Error("Something went wrong")
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

```go
logger.Debug("Detailed debug information")
logger.Info("General information")
logger.Warn("Warning message")
logger.Error("Error occurred")
logger.Fatal("Critical error - application will exit")


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

## Example Output

```json
{"level":"INFO","timestamp":"2023-10-15T14:30:45.123456Z","message":"Application started"}
{"level":"WARN","timestamp":"2023-10-15T14:30:45.124567Z","message":"This is a warning"}
{"level":"ERROR","timestamp":"2023-10-15T14:30:45.125678Z","message":"Something went wrong"}
```
