package db

import (
	"database/sql"
	"fmt"

	// Import sql driver to register itself with the database/sql package.
	_ "github.com/go-sql-driver/mysql"

	"github.com/doublehops/dh-go-framework/internal/config"
	"github.com/doublehops/dh-go-framework/internal/logga"
)

func New(l *logga.Logga, cfg config.DB) (*sql.DB, error) {
	l.Log.Info("opening database connection")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", cfg.User, cfg.Pass, cfg.Host, cfg.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		l.Log.Error(fmt.Sprintf("unable to create db connection. %s", err))
		return db, err
	}

	return db, nil
}
