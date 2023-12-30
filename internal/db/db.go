package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/doublehops/dhapi-example/internal/config"
	"github.com/doublehops/dhapi-example/internal/logga"
)

func New(l *logga.Logga, cfg config.DB) (*sql.DB, error) {
	l.Log.Info("opening database connection")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", cfg.User, cfg.Pass, cfg.Host, cfg.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		l.Log.Error(fmt.Sprintf("unable to create db connection. %s", err))
		return db, err
	}

	return db, nil
}
