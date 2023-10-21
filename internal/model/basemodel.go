package model

import (
	"context"
	"time"

	"github.com/doublehops/dhapi-example/internal/app"
)

type Model interface {
	Unmarshal(data []byte) error
}

type BaseModel struct {
	ID        int32      `json:"id"`
	UserID    int32      `json:"id"`
	CreatedBy int32      `json:"createdBy"`
	UpdatedBy int32      `json:"updatedBy"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func (bm *BaseModel) SetCreated(ctx context.Context) {
	userID := ctx.Value(app.UserIDKey)
	if userID != nil {
		bm.CreatedBy = int32(userID.(int))
		bm.UpdatedBy = int32(userID.(int))
		bm.UserID = int32(userID.(int))
	}

	t := time.Now()

	bm.CreatedAt = &t
	bm.UpdatedAt = &t
}

func (bm *BaseModel) SetUpdated(ctx context.Context) {
	userID := ctx.Value(app.UserIDKey)
	if userID != nil {
		bm.UpdatedBy = int32(userID.(int))
	}

	t := time.Now()

	bm.UpdatedAt = &t
}

func (bm *BaseModel) SetDeleted(ctx context.Context) {
	userID := ctx.Value(app.UserIDKey)
	if userID != nil {
		bm.UpdatedBy = int32(userID.(int))
	}

	t := time.Now()

	bm.UpdatedAt = &t
	bm.DeletedAt = &t
}
