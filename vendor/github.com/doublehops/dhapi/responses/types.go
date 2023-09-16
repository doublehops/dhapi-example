package responses

type singleItemResponseType struct {
	Data interface{} `json:"data"`
}

type multiItemResponseType struct {
	Data interface{} `json:"data"`
	PaginationType
}

type PaginationType struct {
	CurrentPage int `json:"currentPage"`
	PerPage     int `json:"perPage"`
	PageCount   int `json:"pageCount"`
	TotalCount  int `json:"totalCount"`
}

type ValidationErrorResponseType struct {
	Name    string            `json:"name"`
	Message string            `json:"message"`
	Code    int               `json:"code"`
	Status  string            `json:"status"`
	Type    string            `json:"type"`
	Errors  []ValidationField `json:"errors"`
}

type ValidationField map[string][]string

type CustomerErrorMessages map[string]map[string]string
