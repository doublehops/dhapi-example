package service

import (
	"database/sql"

	"github.com/doublehops/dh-go-framework/internal/logga"
	"github.com/doublehops/dh-go-framework/internal/model"
)

type App struct {
	DB  *sql.DB
	Log *logga.Logga
}

// HasPermission will check whether the authenticated user has authorisation for the requested record. This function
// can be overwritten in each service.
func (a *App) HasPermission(ID int32, record model.Model) bool {
	return ID == record.GetUserID()
}
