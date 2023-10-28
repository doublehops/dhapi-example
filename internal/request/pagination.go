package resp

import (
	"math"
	"net/http"
	"strconv"
)

type XXXPaginate struct {
	Page         int   `json:"page"`
	PerPage      int   `json:"perPage"`
	Offset       int   `json:"offset"`
	TotalPages   int   `json:"totalPages"`
	TotalRecords int32 `json:"totalRecords"`
}

var (
	defaultPage    = 1
	defaultPerPage = 10
)

// GetPaginationReq - find and return pagination request vars.
func GetPaginationReq(r *http.Request) *Request {
	query := r.URL.Query()
	page, perPage, offset := getVars(query)

	pg := &Request{
		Page:    page,
		PerPage: perPage,
		Offset:  offset,
	}

	return pg
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
