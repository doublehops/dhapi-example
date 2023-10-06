package logga

import (
	"errors"
	"github.com/doublehops/dhapi-example/internal/config"
	"log/slog"
	"testing"
)

func TestNew(t *testing.T) {

	tests := []struct {
		name          string
		config        *config.Logging
		expectedError error
	}{
		{
			name: "success",
			config: &config.Logging{
				OutputFormat: "text",
				LogLevel:     "DEBUG",
			},
			expectedError: nil,
		},
		{
			name: "success",
			config: &config.Logging{
				OutputFormat: "json",
				LogLevel:     "DEBUG",
			},
			expectedError: nil,
		},
		{
			name: "success",
			config: &config.Logging{
				OutputFormat: "",
				LogLevel:     "DEBUG",
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.config)
			if !errors.Is(err, tt.expectedError) {
				t.Errorf("error not as expected. Wanted: %s; got: %s", tt.expectedError, err)
			}
		})
	}
}

func TestGetLogLevel(t *testing.T) {

	tests := []struct {
		name            string
		configuredLevel string
		expectedLevel   slog.Level
		expectedError   error
	}{
		{
			name:            "successDEBUG",
			configuredLevel: "DEBUG",
			expectedLevel:   slog.LevelDebug,
			expectedError:   nil,
		},
		{
			name:            "successINFO",
			configuredLevel: "INFO",
			expectedLevel:   slog.LevelInfo,
			expectedError:   nil,
		},
		{
			name:            "successWARN",
			configuredLevel: "WARN",
			expectedLevel:   slog.LevelWarn,
			expectedError:   nil,
		},
		{
			name:            "successERROR",
			configuredLevel: "ERROR",
			expectedLevel:   slog.LevelError,
			expectedError:   nil,
		},
		{
			name:            "failNotConfigured",
			configuredLevel: "",
			expectedLevel:   slog.LevelInfo,
			expectedError:   InvalidLogLevelValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			level, err := getLogLevel(tt.configuredLevel)
			if !errors.Is(err, tt.expectedError) {
				t.Errorf("error not as expected. Wanted: %s; got: %s", tt.expectedError, err)
			}
			if tt.expectedLevel != level {
				t.Errorf("level not as expected. Expected: %s; got: %s", tt.expectedLevel, level)
			}
		})
	}
}
