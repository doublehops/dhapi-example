package repository

import (
	"database/sql"
	"fmt"
)

// GetRecordCount will retrieve the number of records for a given query for pagination responses.
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
		err = row.Scan(&cs.Count)
		if err != nil {
			return 0, fmt.Errorf("unable to scan query into var")
		}
	}

	return cs.Count, nil
}

//func FormatWhereClauses(r []req.FilterRules) (string, []any) {
//	var vars []any
//	clause := " WHERE "
//	for _, rule := range r {
//
//	}
//
//	return clause, vars
//}
