package responses

type SingleItemResponse struct {
	Data interface{} `json:"data"`
}

type MultiItemResponse struct {
	Data interface{} `json:"data"`
	PaginationType
}

type PaginationType struct {
	CurrentPage int `json:"currentPage"`
	PerPage     int `json:"perPage"`
	PageCount   int `json:"pageCount"`
	TotalCount  int `json:"totalCount"`
}

type ValidationErrorResponse struct {
	Name    string        `json:"name"`
	Message string        `json:"message"`
	Code    int           `json:"code"`
	Status  string        `json:"status"`
	Type    string        `json:"type"`
	Errors  ErrorMessages `json:"errors"`
}

type ErrorMessages map[string][]string
