package logga

import (
	"context"
)

// Debug - args should be key/value pairs separated by a space. Example: "file", "main.go"
func (l *Logga) Debug(ctx context.Context, msg string, args ...any) {
	args = getArguments(ctx, args...)
	l.Log.Debug(msg, args...)
}

// Info - args should be key/value pairs separated by a space. Example: "file", "main.go"
func (l *Logga) Info(ctx context.Context, msg string, args ...any) {
	args = getArguments(ctx, args...)
	l.Log.Info(msg, args...)
}

// Warn - args should be key/value pairs separated by a space. Example: "file", "main.go"
func (l *Logga) Warn(ctx context.Context, msg string, args ...any) {
	args = getArguments(ctx, args...)
	l.Log.Warn(msg, args...)
}

// Error - args should be key/value pairs separated by a space. Example: "file", "main.go"
func (l *Logga) Error(ctx context.Context, msg string, args ...any) {
	args = getArguments(ctx, args...)
	l.Log.Error(msg, args...)
}

func getArguments(ctx context.Context, args ...any) []any {
	ctxArgs := getContextAtts(ctx)

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

	if userID := ctx.Value("userID"); userID != nil {
		args = append(args, "userID", userID)
	}

	return args
}
