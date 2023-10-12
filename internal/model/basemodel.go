package model

import "time"

type BaseModel struct {
	ID        int       `json:"id"`
	CreatedBy int       `json:"createdBy"`
	UpdatedBy int       `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
