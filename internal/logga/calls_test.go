package logga

import (
	"context"
	"reflect"
	"testing"
)

func TestGetArguments(t *testing.T) {

	tests := []struct {
		name         string
		args         []interface{}
		ctxVars      map[string]any
		expectedArgs interface{}
	}{
		{
			name:         "successWithAll",
			args:         []any{"hello", "world"},
			ctxVars:      map[string]any{"traceID": "ABCD-1234", "userID": 123},
			expectedArgs: []interface{}{"hello", "world", "traceID", "ABCD-1234", "userID", 123},
		},
		{
			name:         "SuccessOnlyArgs",
			args:         []any{"hello", "world"},
			ctxVars:      map[string]any{},
			expectedArgs: []interface{}{"hello", "world"},
		},
		{
			name:         "SuccessOnlyCtxArgs",
			args:         []any{},
			ctxVars:      map[string]any{"traceID": "ABCD-1234", "userID": 123},
			expectedArgs: []interface{}{"traceID", "ABCD-1234", "userID", 123},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			for key, value := range tt.ctxVars {
				ctx = context.WithValue(ctx, key, value)
			}
			args := getArguments(ctx, tt.args...)
			if !reflect.DeepEqual(tt.expectedArgs, args) {
				t.Errorf("args not as expected. Expected: %v; got: %v", tt.expectedArgs, args)
			}
		})
	}
}
