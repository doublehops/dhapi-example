package logga

import (
	"errors"
	"github.com/doublehops/dhapi-example/internal/config"
	"log/slog"
	"os"
)

var (
	InvalidLogLevelValue = errors.New("a valid log level was not defined in configuration")
)

type Logga struct {
	Log *slog.Logger
}

// New will return the log handler with the options defined in config.
func New(cfg *config.Logging) (*Logga, error) {
	level, err := getLogLevel(cfg.LogLevel)
	if err != nil {
		return nil, err
	}

	logLevel := &slog.LevelVar{}
	logLevel.Set(level)

	var logger *slog.Logger

	switch cfg.OutputFormat {
	case "json":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	case "text":
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	default:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	}

	return &Logga{Log: logger}, nil
}

func getLogLevel(configuredLevel string) (slog.Level, error) {
	switch configuredLevel {
	case "DEBUG":
		return slog.LevelDebug, nil
	case "INFO":
		return slog.LevelInfo, nil
	case "WARN":
		return slog.LevelWarn, nil
	case "ERROR":
		return slog.LevelError, nil
	}

	return slog.LevelInfo, InvalidLogLevelValue
}
