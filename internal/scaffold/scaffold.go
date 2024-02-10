package scaffold

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/doublehops/dhapi-example/internal/logga"
	"io"
	"log"
	"os"
	"strings"
)

type Scaffold struct {
	DB *sql.DB
	l  *logga.Logga

	tableName   string
	ConfigPaths `json:"paths"`
}

type ConfigPaths struct {
	Handlers   string `json:"handlers"`
	Model      string `json:"model"`
	Repository string `json:"repository"`
	Service    string `json:"service"`
}

type modelStrings struct {
	Name         string
	FirstInitial string
	CamelCase    string
	PascalCase   string
	SnakeCase    string
	LowerCase    string
	ModuleName   string
}

func New(cfg ConfigPaths, tableName string, db *sql.DB, logga *logga.Logga) *Scaffold {
	return &Scaffold{
		DB:          db,
		l:           logga,
		tableName:   tableName,
		ConfigPaths: cfg,
	}
}

func (s *Scaffold) Run(path string) error {
	ctx := context.Background()

	columns, err := s.getTableDefinition()
	if err != nil {
		s.l.Error(ctx, "error getting column. "+err.Error(), nil)
		return errors.New("failed to run. " + err.Error())
	}

	moduleName, err := getModuleName()
	if err != nil {
		s.l.Error(ctx, err.Error(), nil)
		return errors.New("failed to run. " + err.Error())
	}

	ms := modelStrings{
		Name:         s.tableName,
		FirstInitial: GetFirstRune(s.tableName),
		CamelCase:    ToCamelCase(s.tableName),
		PascalCase:   ToPascalCase(s.tableName),
		SnakeCase:    s.tableName,
		LowerCase:    ToLowerCase(s.tableName),
		ModuleName:   moduleName,
	}

	s.buildStruct(ms, columns)

	return nil
}

// getModuleName will get the module name from go.mod to use to populate the templates.
func getModuleName() (string, error) {
	f, err := os.Open("../../../go.mod")
	if err != nil {
		return "", errors.New("Opening go.mod failed. " + err.Error())
	}
	rawBytes, err := io.ReadAll(f)
	lines := strings.Split(string(rawBytes), "\n")

	module := strings.Replace(lines[0], "module ", "", 0)

	return module, nil
}

func (s *Scaffold) getTableDefinition() (map[string]string, error) {
	cols := map[string]string{}

	q := "DESCRIBE " + s.tableName
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

		columnName := fmt.Sprintf("%s", values[0])
		columnType := fmt.Sprintf("%s", values[1])
		cols[columnName] = columnType
	}

	return cols, nil
}
