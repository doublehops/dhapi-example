package repository

import (
	"database/sql"
	"fmt"
)

// GetRecordCount will retrieve the number of records for a given query for pagination responses.
// The function expects only one column in the query. An example would be `SELECT COUNT(*) count FROM {table}1`.
func GetRecordCount(DB *sql.DB, q string, params []any) (int32, error) {
	var (
		err error
		row *sql.Rows
	)
	if params == nil {
		row, err = DB.Query(q)
	} else {
		row, err = DB.Query(q, params...)
	}
	defer row.Close()

	if err != nil {
		return 0, fmt.Errorf("unable to run count query. %s", err)
	}

	var count int32

	for row.Next() {
		err = row.Scan(&count)
		if err != nil {
			return 0, fmt.Errorf("unable to scan query result. %s", err)
		}
	}

	return count, err
}
