package repository

import (
	"database/sql"
	"fmt"
)

// GetRecordCount will retrieve the number of records for a given query for pagination responses.
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

	columns, err := row.Columns()
	colCount := len(columns)
	values := make([]interface{}, colCount)
	valuesPtrs := make([]interface{}, colCount)

	if row.Next() {
		for i := range columns {
			valuesPtrs[i] = &values[i]

		}

		err = row.Scan(valuesPtrs...)
		if err != nil {
			return 0, fmt.Errorf("unable to scan query result. %s", err)
		}
		count := values[0].(int64) // We only want the first column.

		return int32(count), nil
	}

	return 0, fmt.Errorf("unable to find count")
}
