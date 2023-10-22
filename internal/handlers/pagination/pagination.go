package pagination

import (
	"math"
	"net/http"
	"strconv"
)

type RequestPagination struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
	Offset  int `json:"offset"`
}

type ResponsePagination struct {
	Page         int   `json:"page"`
	PerPage      int   `json:"perPage"`
	TotalPages   int   `json:"totalPages"`
	TotalRecords int64 `json:"totalRecords"`
}

var (
	defaultPage    = 1
	defaultPerPage = 10
)

// GetPaginationReq - find and return pagination request vars.
func GetPaginationReq(r *http.Request) *RequestPagination {
	query := r.URL.Query()
	page, perPage, offset := getVars(query)

	pg := RequestPagination{
		Page:    page,
		PerPage: perPage,
		Offset:  offset,
	}

	return &pg
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

// GetPaginationResponse - Get pagination data to return in API response.
func GetPaginationResponse(pg *RequestPagination, count int64) *ResponsePagination {
	pageApprox := float64(count) / float64(pg.PerPage)
	totalPages := math.Ceil(pageApprox)
	tp := int(totalPages)

	return &ResponsePagination{
		Page:         pg.Page,
		PerPage:      pg.PerPage,
		TotalPages:   tp,
		TotalRecords: count,
	}
}
