package repository

import (
	"fmt"
	"strings"
	"unicode"

	req "github.com/doublehops/dhapi-example/internal/request"
)

func BuildQuery(query string, p *req.Request, getCount bool) (string, []any) {
	var pParams []any
	q, params := addFilters(query, p.Filters)

	if getCount {
		q, pParams = addPagination(q, p, false)
		q = replaceCount(q, true)
	} else {
		q, pParams = addPagination(q, p, true)
		params = append(params, pParams...)
		q = replaceCount(q, false)
	}

	return q, params
}

func addPagination(query string, pagination *req.Request, includePagination bool) (string, []any) {
	if !includePagination {
		return replaceLimitClause(query, "", false), nil
	}

	q := " LIMIT ?, ?"
	params := []any{pagination.Offset, pagination.PerPage}

	return replaceLimitClause(query, q, true), params
}

func addFilters(query string, filters []req.FilterRule) (string, req.Params) {
	if len(filters) == 0 {
		return replaceWhereClause(query, ""), nil
	}

	var params req.Params

	var whereClauses []string
	for _, f := range filters {
		field := ConvertStr(f.Field)

		switch f.Type {
		case req.FilterEquals:
			clause := fmt.Sprintf("%s = ?", field)
			whereClauses = append(whereClauses, clause)
			params = append(params, f.Value)
		case req.FilterLike:
			clause := field + " LIKE ?"
			whereClauses = append(whereClauses, clause)
			val := fmt.Sprintf("%%%v%%", f.Value) // Will be equivalent to '%val%'
			params = append(params, val)
		case req.FilterIsNull:
			clause := field + " IS NULL"
			whereClauses = append(whereClauses, clause)
			params = append(params, f.Value)
		}
	}

	str := " WHERE " + strings.Join(whereClauses, " AND ")

	return replaceWhereClause(query, str), params
}

func replaceWhereClause(q, whereClause string) string {
	return strings.Replace(q, "__WHERE_CLAUSE__", whereClause, 1)
}

func replaceLimitClause(q, limitClause string, includePagination bool) string {
	if !includePagination {
		return strings.Replace(q, "__PAGINATION__", "", 1)
	}

	return strings.Replace(q, "__PAGINATION__", limitClause, 1)
}

func replaceCount(q string, getCount bool) string {
	if getCount {
		return strings.Replace(q, "__COUNT__", "count(*) count, ", 1)
	}

	return strings.Replace(q, "__COUNT__", "", 1)
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
