package app

import (
	"database/sql"

	"github.com/doublehops/dhapi-example/internal/logga"
)

const (
	UserIDKey = "userID"
)

type App struct {
	DB  *sql.DB
	Log *logga.Logga
}
