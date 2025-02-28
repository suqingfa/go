package test

import (
	"log/slog"
	"os"
	"testing"
)

func TestSlog(t *testing.T) {
	// Default text handler
	slog.Info("Hello, structured logging!")
	slog.Warn("This is a warning", "user", "Alice")

	// JSON handler (structured output)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("User logged in", "id", 123, "username", "alice")
}
