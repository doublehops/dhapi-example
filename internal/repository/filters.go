package repository

import (
	"fmt"
	"strings"
	"unicode"

	req "github.com/doublehops/dhapi-example/internal/request"
)

func BuildQuery(query string, filters []req.FilterRule) string {
	q := addFilters(query, filters)

	return q
}

func addFilters(query string, filters []req.FilterRule) string {
	if len(filters) == 0 {
		return replaceWhereClause(query, "")
	}
	
	var whereClauses []string
	for _, f := range filters {
		field := ConvertStr(f.Field)

		switch f.Type {
		case req.FilterEquals:
			clause := fmt.Sprintf("%s = ?", field)
			whereClauses = append(whereClauses, clause)
		case req.FilterLike:
			clause := field + " LIKE %?%"
			whereClauses = append(whereClauses, clause)
		case req.FilterIsNull:
			clause := field + " IS NULL"
			whereClauses = append(whereClauses, clause)
		}
	}

	str := " WHERE " + strings.Join(whereClauses, " AND ")

	return replaceWhereClause(query, str)
}

func replaceWhereClause(q, whereClause string) string {
	return strings.Replace(q, "__WHERE_CLAUSE__", whereClause, 1)
}

//func getFieldValue(field string, instance any) (any, error) {
//	rv := reflect.ValueOf(instance)
//
//	val := rv.FieldByName(field)
//	if !val.IsValid() {
//		return nil, fmt.Errorf("unable to find value of %s", field)
//	}
//
//	return val.Interface(), nil
//}

// ConvertStr will convert camelcase string to snake case for SQL query.
func ConvertStr(field string) string {
	var str string

	// Iterate through the input string
	for i, runeValue := range field {
		// If the character is an uppercase letter (but not the first character in the string)
		// add an underscore before it
		if i > 0 && unicode.IsUpper(runeValue) {
			str += "_"
		}
		// Then, whether or not we added an underscore, append the lowercase of the current character
		str += string(unicode.ToLower(runeValue))
	}
	return str
}
