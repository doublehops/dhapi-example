package logga

import (
	"context"
	"reflect"
	"testing"

	"github.com/doublehops/dhapi-example/internal/config"
)

func TestGetArguments(t *testing.T) {
	tests := []struct {
		name         string
		context      context.Context
		args         KVPs
		ctxVars      map[string]any
		expectedArgs interface{}
	}{
		{
			name:         "successWithAll",
			context:      context.Background(),
			args:         KVPs{"hello": "world"},
			ctxVars:      map[string]any{"traceID": "ABCD-1234", "userID": 123},
			expectedArgs: []interface{}{"hello", "world", "traceID", "ABCD-1234", "userID", 123},
		},
		{
			name:         "SuccessOnlyArgs",
			context:      context.Background(),
			args:         KVPs{"hello": "world"},
			ctxVars:      map[string]any{},
			expectedArgs: []interface{}{"hello", "world"},
		},
		{
			name:         "SuccessOnlyCtxArgs",
			context:      context.Background(),
			args:         KVPs{},
			ctxVars:      map[string]any{"traceID": "ABCD-1234", "userID": 123},
			expectedArgs: []interface{}{"traceID", "ABCD-1234", "userID", 123},
		},
		{
			name:         "SuccessOnlyCtxArgs",
			context:      nil,
			args:         KVPs{"hello": "world"},
			ctxVars:      map[string]any{},
			expectedArgs: []interface{}{"hello", "world"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.context
			for key, value := range tt.ctxVars {
				ctx = context.WithValue(ctx, key, value)
			}
			args := getArguments(ctx, tt.args)
			if !reflect.DeepEqual(tt.expectedArgs, args) {
				t.Errorf("args not as expected. Expected: %v; got: %v", tt.expectedArgs, args)
			}
		})
	}
}

func TestCalls(t *testing.T) {
	ctx := context.Background()

	cfg := &config.Logging{
		LogLevel:     "DEBUG",
		OutputFormat: "text",
	}

	l, _ := New(cfg)

	l.Info(ctx, "my message", nil)
	l.Debug(ctx, "my message", nil)
	l.Warn(ctx, "my message", nil)
	l.Error(ctx, "my message", nil)
}
