package logga

import (
	"github.com/doublehops/dhapi-example/internal/config"
	"log/slog"
	"os"
)

// New will return the log handler.
// It might be nice to allow the output location to be configured also.
func New(cfg *config.Logging) *slog.Logger {
	switch cfg.OutputFormat {
	case "json":
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	case "text":
		return slog.New(slog.NewTextHandler(os.Stdout, nil))
	default:
		return slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
}
