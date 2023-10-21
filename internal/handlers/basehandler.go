package handlers

import (
	"context"
	"github.com/doublehops/dhapi-example/internal/model"
)

type service interface {
	GetByID(ctx context.Context, author *model.Model, ID int32) error
}
