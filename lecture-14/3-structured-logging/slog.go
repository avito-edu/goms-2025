package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// ColorHandler for colored output
type ColorHandler struct {
	slog.Handler
}

func (h *ColorHandler) Handle(ctx context.Context, r slog.Record) error {
	// ANSI colors
	const (
		reset   = "\033[0m"
		red     = "\033[31m"
		green   = "\033[32m"
		yellow  = "\033[33m"
		blue    = "\033[34m"
		magenta = "\033[35m"
		cyan    = "\033[36m"
		gray    = "\033[90m"
		white   = "\033[97m"
	)

	// Colors for levels
	var levelColor, levelText string
	switch r.Level {
	case slog.LevelDebug:
		levelColor = magenta
		levelText = "DEBG"
	case slog.LevelInfo:
		levelColor = green
		levelText = "INFO"
	case slog.LevelWarn:
		levelColor = yellow
		levelText = "WARN"
	case slog.LevelError:
		levelColor = red
		levelText = "ERR "
	default:
		levelColor = white
		levelText = "????"
	}

	// Format time
	timeStr := r.Time.Format("15:04:05")

	// Output main line
	fmt.Printf("%s%s%s %s%s%s %s%s%s",
		gray, timeStr, reset,
		levelColor, levelText, reset,
		white, r.Message, reset)

	// Output attributes
	r.Attrs(func(attr slog.Attr) bool {
		fmt.Printf(" %s%s%s=%s%v%s",
			cyan, attr.Key, reset,
			blue, attr.Value.Any(), reset)
		return true
	})

	fmt.Println() // New line
	return nil
}

func main() {
	// Create colored handler
	handler := &ColorHandler{
		Handler: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
	}

	logger := slog.New(handler)

	// Demo of different levels
	logger.Debug("Debug information",
		"user_id", 123,
		"component", "auth")

	logger.Info("User logged in",
		"username", "ivan",
		"ip", "192.168.1.100",
		"timestamp", time.Now())

	logger.Warn("High memory usage",
		"memory_used", "85%",
		"threshold", "80%")

	logger.Error("Database error",
		"error", "connection timeout",
		"database", "postgresql",
		"retry_count", 3)
}

func main() {
	tracer, closer := initJaeger("my-service")
	defer closer.Close()

	span := tracer.StartSpan("my-operation")
	defer span.Finish()

	ext.SpanKindRPCClient.Set(span)
	span.SetTag("custom-tag", "tag-value")
	span.LogFields(log.String("event", "my-event"))

	// implementation
}
