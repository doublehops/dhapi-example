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
	var cs struct {
		Count int32
	}

	if params == nil {
		row, err = DB.Query(q)
	} else {
		row, err = DB.Query(q, params...)
	}
	defer row.Close()

	if err != nil {
		return 0, fmt.Errorf("unable to run count query. %s", err)
	}

	// from https://stackoverflow.com/questions/59629440/sql-expected-3-destination-arguments-in-scan-not-1-in-golang
	TRY THIS - https://stackoverflow.com/questions/17845619/how-to-call-the-scan-variadic-function-using-reflection
	columns, err := row.Columns()
	//receiver := make([]interface{}, len(columns))
	receiver := make([]*interface{}, len(columns))

	if row.Next() {
		err = row.Scan(receiver...)
		if err != nil {
			return 0, fmt.Errorf("unable to scan query into var")
		}
		fmt.Printf("receiver: %+v\n", receiver)
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
