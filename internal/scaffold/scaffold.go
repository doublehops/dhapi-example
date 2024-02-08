package scaffold

import (
	"database/sql"
	"github.com/doublehops/dhapi-example/internal/logga"
)

type Scaffold struct {
	db    *sql.DB
	logga *logga.Logga

	ModelName   string
	ConfigPaths `json:"paths"`
}

type ConfigPaths struct {
	Handlers   string `json:"handlers"`
	Model      string `json:"model"`
	Repository string `json:"repository"`
	Service    string `json:"service"`
}

func New(cfg ConfigPaths, modelName string, db *sql.DB, logga *logga.Logga) *Scaffold {
	return &Scaffold{
		db:          db,
		logga:       logga,
		ModelName:   modelName,
		ConfigPaths: cfg,
	}
}

func (s *Scaffold) Run() {

}
