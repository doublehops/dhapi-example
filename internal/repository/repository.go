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
	//var cs struct {
	//	Count int32
	//}

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
	//TRY THIS - https://stackoverflow.com/questions/17845619/how-to-call-the-scan-variadic-function-using-reflection
	columns, err := row.Columns()
	colCount := len(columns)
	values := make([]interface{}, colCount)
	valuesPtrs := make([]interface{}, colCount)
	//receiver := make([]interface{}, len(columns))
	//receiver := make([]*interface{}, len(columns))

	if row.Next() {
		for i := range columns {
			valuesPtrs[i] = &values[i]

		}

		row.Scan(valuesPtrs...)

		rowCount := columns[0]
		fmt.Printf(">>>> RowCount: %s\n", rowCount)

		//count := columns[0]
		//fmt.Printf("count type: %d\n", reflect.TypeOf(count))

		count := values[0].(int64)

		return int32(count), nil

		//for i, _ := range columns {
		//	val := values[i]

		//b, ok := val.([]byte)
		//var v interface{}
		//if ok {
		//	v = string(b)
		//} else {
		//	v = val
		//}
		//
		//fmt.Printf("v type: %d\n", reflect.TypeOf(v))
		//vl, ok := v.(int)
		//fmt.Printf("vl type: %d\n", reflect.TypeOf(vl))
		//if !ok {
		//	return 0, fmt.Errorf("unable to convert count to int32")
		//}
		//
		//return int32(vl), nil
	}
	//}
	//
	return 0, fmt.Errorf("unable to find count")
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
