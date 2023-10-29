package request

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"math"
	"net/http"
	"strconv"
)

var (
	defaultPage    = 1
	defaultPerPage = 10
)

// GetRequestParams - find and return pagination request vars.
func GetRequestParams(r *http.Request, ps httprouter.Params, filters []FilterRule) *Request {
	query := r.URL.Query()
	page, perPage, offset := getVars(query)

	f := getQueryParams(r.Context(), ps, filters)

	pg := &Request{
		Page:    page,
		PerPage: perPage,
		Offset:  offset,
		Filters: f,
	}

	return pg
}

func getQueryParams(ctx context.Context, ps httprouter.Params, filters []FilterRule) []FilterRule {
	var newFilters []FilterRule
	for _, f := range filters {
		val := ps.ByName(ps.ByName(f.Field))
		if val == "" {
			continue
		}

		f.Value = val
		newFilters = append(newFilters, f)
	}

	return newFilters
}

// getVars - Search request query for wanted var and return value, if not found, return default value.
func getVars(query map[string][]string) (int, int, int) {
	page := defaultPage
	perPage := defaultPerPage
	offset := 0

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "perPage":
			perPage, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		}
	}

	if page != 1 {
		offset = (page - 1) * perPage
	}

	return page, perPage, offset
}

// SetRecordCount - Set pagination data to return in API response.
func (r *Request) SetRecordCount(count int32) {
	pageApprox := float64(count) / float64(r.PerPage)
	totalPages := math.Ceil(pageApprox)
	tp := int(totalPages)

	r.PageCount = tp
	r.TotalCount = count
}
