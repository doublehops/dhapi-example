package service

import (
	"github.com/doublehops/dh-go-framework/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApp_HasPermission(t *testing.T) {
	tests := []struct {
		name             string
		userID           int32
		recordUserID     int32
		expectedResponse bool
	}{
		{
			name:             "permissionGranted",
			userID:           1,
			recordUserID:     1,
			expectedResponse: true,
		},
		{
			name:             "permissionNotGranted",
			userID:           1,
			recordUserID:     2,
			expectedResponse: false,
		},
	}

	for _, tt := range tests {
		app := App{}
		m := &model.BaseModel{
			UserID: tt.recordUserID,
		}

		hasPermission := app.HasPermission(tt.userID, m)
		assert.Equal(t, tt.expectedResponse, hasPermission, "permission not as expected")
	}
}
