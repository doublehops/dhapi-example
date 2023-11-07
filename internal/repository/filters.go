package repository

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	req "github.com/doublehops/dhapi-example/internal/request"
)

func BuildQuery(query string, p *req.Request, getCount bool) (string, []any) {
	var pParams []any
	q, params := addFilters(query, p.Filters)

	if getCount {
		//q, pParams = addPagination(q, p, false)
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
		log.Fatal("we should not be here")
	}

	params := []any{pagination.Offset, pagination.PerPage}

	return query + " LIMIT ?, ?", params
}

func addFilters(query string, filters []req.FilterRule) (string, req.Params) {

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
		}
	}

	str := " WHERE " + strings.Join(whereClauses, " AND ")

	return query + str, params
}

func replaceCount(q string, getCount bool) string {
	if getCount {
		return strings.Replace(q, "__COUNT__", "count(*) count, ", 1)
	}

	return strings.Replace(q, "__COUNT__", "", 1)
}

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
