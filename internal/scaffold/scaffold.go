package scaffold

import (
	"context"
	"database/sql"
	"errors"
	"github.com/doublehops/dhapi-example/internal/logga"
	"log"
)

type Scaffold struct {
	DB *sql.DB
	l  *logga.Logga

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
		DB:          db,
		l:           logga,
		ModelName:   modelName,
		ConfigPaths: cfg,
	}
}

func (s *Scaffold) Run(path string) {
	ctx := context.Background()

	columns, err := s.getTableDefinition()
	if err != nil {
		s.l.Error(ctx, "error getting column. "+err.Error(), nil)
	}

}

func (s *Scaffold) getTableDefinition() (map[string]string, error) {
	cols := map[string]string{}

	q := "DESCRIBE " + s.ModelName
	rows, err := s.DB.Query(q)
	if err != nil {
		return nil, errors.New("error executing query. " + err.Error())
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, errors.New("Error getting columns. " + err.Error())
	}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			log.Fatal(err)
		}

		columnName := values[0]
		columnType := values[1]
		cols[columnName.(string)] = columnType.(string)
	}

	return cols, nil
}
