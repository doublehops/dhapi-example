package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/doublehops/dh-go-framework/internal/app"
	"github.com/doublehops/dh-go-framework/internal/logga"
)

type BaseHandler struct {
	Log *logga.Logga
}

func (bh *BaseHandler) GetUser(ctx context.Context) int32 {
	var intValue int32
	var ok bool

	val := ctx.Value(app.UserIDKey)
	if intValue, ok = val.(int32); !ok {
		bh.Log.Error(ctx, "unable to convert userID to int32", nil)
	}

	return intValue
}

func (bh *BaseHandler) WriteJSON(ctx context.Context, w http.ResponseWriter, statusCode int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		bh.Log.Error(ctx, "unable to marshal to JSON. "+err.Error(), nil)
	}
}
