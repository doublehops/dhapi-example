package model

import (
	"context"
	"fmt"
	"time"

	"github.com/doublehops/dhapi-example/internal/app"
)

type Model interface {
	GetUserID() int32
	SetCreated(context.Context)
}

type BaseModel struct {
	ID        int32      `json:"id"`
	UserID    int32      `json:"userId"`
	CreatedBy int32      `json:"createdBy"`
	UpdatedBy int32      `json:"updatedBy"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// Deprecated - remove
func (bm *BaseModel) GetUserID() int32 {
	return bm.UserID
}

func (bm *BaseModel) SetCreated(ctx context.Context) {
	userID := ctx.Value(app.UserIDKey)
	if userID != nil {

		uID, ok := userID.(int)
		if !ok {
			// todo - check what to do here. logging would be best.
		}

		bm.CreatedBy = int32(uID)
		bm.UpdatedBy = int32(uID)
		bm.UserID = int32(uID)
	}

	t := time.Now()

	bm.CreatedAt = &t
	bm.UpdatedAt = &t
}

func (bm *BaseModel) SetUpdated(ctx context.Context) {
	userID := bm.getRequestUserID(ctx)
	if userID > 0 {
		bm.UpdatedBy = userID
	}

	t := time.Now()

	bm.UpdatedAt = &t
}

func (bm *BaseModel) SetDeleted(ctx context.Context) {
	userID := bm.getRequestUserID(ctx)
	if userID > 0 {
		bm.UpdatedBy = userID
	}

	t := time.Now()

	bm.UpdatedAt = &t
	bm.DeletedAt = &t
}

// getRequestUserID will retrieve userID from context.
func (bm *BaseModel) getRequestUserID(ctx context.Context) int32 {
	val := ctx.Value(app.UserIDKey)
	var intValue int32
	var ok bool

	if intValue, ok = val.(int32); !ok {
		// @todo - log something here
		fmt.Println("unable to get userID.")
	}

	return intValue
}
