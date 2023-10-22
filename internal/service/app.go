package service

import (
	"database/sql"
	"github.com/doublehops/dhapi-example/internal/model"

	"github.com/doublehops/dhapi-example/internal/logga"
)

type App struct {
	DB  *sql.DB
	Log *logga.Logga
}

func (a *App) HasPermission(ID int32, record model.Model) bool {
	return ID == record.GetUserID()
}
