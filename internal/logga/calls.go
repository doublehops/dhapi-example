package logga

import (
	"context"
	"github.com/doublehops/dhapi-example/internal/app"
)

// Debug - args should be key/value pairs separated by a space. Example: "file", "migrate.go"
func (l *Logga) Debug(ctx context.Context, msg string, KVPs KVPs) {
	args := getArguments(ctx, KVPs)
	l.Log.Debug(msg, args...)
}

// Info - args should be key/value pairs separated by a space. Example: "file", "migrate.go"
func (l *Logga) Info(ctx context.Context, msg string, KVPs KVPs) {
	args := getArguments(ctx, KVPs)
	l.Log.Info(msg, args...)
}

// Warn - args should be key/value pairs separated by a space. Example: "file", "migrate.go"
func (l *Logga) Warn(ctx context.Context, msg string, KVPs KVPs) {
	args := getArguments(ctx, KVPs)
	l.Log.Warn(msg, args...)
}

// Error - args should be key/value pairs separated by a space. Example: "file", "migrate.go"
func (l *Logga) Error(ctx context.Context, msg string, KVPs KVPs) {
	args := getArguments(ctx, KVPs)
	l.Log.Error(msg, args...)
}

func getArguments(ctx context.Context, KVPs KVPs) []any {
	ctxArgs := getContextAtts(ctx)

	var args []any
	for key, value := range KVPs {
		args = append(args, key, value)
	}

	if len(ctxArgs) > 0 {
		args = append(args, ctxArgs...)
	}

	return args
}

func getContextAtts(ctx context.Context) []any {
	var args []interface{}

	if ctx == nil {
		return args
	}

	if traceID := ctx.Value("traceID"); traceID != nil {
		args = append(args, "traceID", traceID)
	}

	if userID := ctx.Value(app.UserIDKey); userID != nil {
		args = append(args, app.UserIDKey, userID)
	}

	return args
}
