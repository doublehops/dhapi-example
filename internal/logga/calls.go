package logga

import (
	"context"
)

func (l *Logga) Debug(ctx context.Context, msg string, args ...any) {
	args = getArguments(ctx, args...)
	l.Log.Debug(msg, args...)
}

func (l *Logga) Info(ctx context.Context, msg string, args ...any) {
	args = getArguments(ctx, args...)
	l.Log.Info(msg, args...)
}

func (l *Logga) Warn(ctx context.Context, msg string, args ...any) {
	args = getArguments(ctx, args...)
	l.Log.Warn(msg, args...)
}

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
