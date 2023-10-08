package handlers

import (
	"database/sql"

	"github.com/doublehops/dhapi-example/internal/logga"
)

type App struct {
	DB  *sql.DB
	Log *logga.Logga
}
