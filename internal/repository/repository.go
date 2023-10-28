package repository

import (
	"database/sql"
	"fmt"
)

func GetRecordCount(DB *sql.DB, q string, params ...any) (int32, error) {
	var (
		err error
		row *sql.Rows
	)
	var cs struct{ Count int32 }

	if params == nil {
		row, err = DB.Query(q)
	} else {
		row, err = DB.Query(q, params)
	}

	if err != nil {
		return 0, fmt.Errorf("unable to run count query. %s", err)
	}

	if row.Next() {
		row.Scan(&cs.Count)
	}

	return cs.Count, nil
}
