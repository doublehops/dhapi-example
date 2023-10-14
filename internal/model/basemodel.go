package model

import (
	"context"
	"github.com/doublehops/dhapi-example/internal/app"
	"time"
)

const DateFormat = "2006-1-2 15:4:5"

type BaseModel struct {
	ID        int32     `json:"id"`
	CreatedBy int32     `json:"createdBy"`
	UpdatedBy int32     `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

func (bm *BaseModel) SetCreated(ctx context.Context) {
	userID := ctx.Value(app.UserIDKey)
	if userID != nil {
		bm.CreatedBy = int32(userID.(int))
		bm.UpdatedBy = int32(userID.(int))
	}

	bm.CreatedAt = time.Now()
	bm.UpdatedAt = time.Now()
}

func (bm *BaseModel) SetUpdated(ctx context.Context) {
	userID := ctx.Value(app.UserIDKey)
	if userID != nil {
		bm.UpdatedBy = int32(userID.(int))
	}

	bm.UpdatedAt = time.Now()
}
