package db

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"

	"github.com/doublehops/dhapi-example/internal/config"
)

func New(l *slog.Logger, cfg config.DB) (*sql.DB, error) {
	l.Info("opening database connection")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", cfg.User, cfg.Pass, cfg.Host, cfg.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		l.Error("unable to create db connection. %s", err)
		return db, err
	}

	return db, nil
}
