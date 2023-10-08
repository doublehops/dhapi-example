package logga

import (
	"errors"
	"github.com/doublehops/dhapi-example/internal/config"
	"github.com/doublehops/dhapi-example/test/testbuffer"
	"io"
	"log/slog"
	"os"
)

var (
	InvalidLogWriter     = errors.New("a valid writer was not defined in configuration")
	InvalidLogLevelValue = errors.New("a valid log level was not defined in configuration")
)

type Logga struct {
	Log *slog.Logger
}

// New will return the log handler with the options defined in config.
func New(cfg *config.Logging) (*Logga, error) {
	level, err := getLogLevelFromConfig(cfg.LogLevel)
	if err != nil {
		return nil, err
	}

	logLevel := &slog.LevelVar{}
	logLevel.Set(level)

	writer, err := getWriterFromConfig(cfg.Writer)

	var logger *slog.Logger

	switch cfg.OutputFormat {
	case "json":
		logger = slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: logLevel}))
	case "text":
		logger = slog.New(slog.NewTextHandler(writer, &slog.HandlerOptions{Level: logLevel}))
	default:
		logger = slog.New(slog.NewTextHandler(writer, &slog.HandlerOptions{Level: logLevel}))
	}

	return &Logga{Log: logger}, nil
}

func getWriterFromConfig(configuredWriter string) (io.Writer, error) {
	switch configuredWriter {
	case "stdout":
		return os.Stdout, nil
	case "": // Default to stdout if none is defined.
		return os.Stdout, nil
	case "testwriter": // Used for testing.
		return testbuffer.TestBuffer{}, nil
	}

	return nil, InvalidLogWriter
}

func getLogLevelFromConfig(configuredLevel string) (slog.Level, error) {
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
