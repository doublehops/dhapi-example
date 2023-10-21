package resp

type SingleItemResp struct {
	Data interface{} `json:"data"`
}

type ListResp struct {
	Data interface{} `json:"data"`
	Pagination
}

type Pagination struct {
	CurrentPage int `json:"currentPage"`
	PerPage     int `json:"perPage"`
	PageCount   int `json:"pageCount"`
	TotalCount  int `json:"totalCount"`
}

type ValidateErrResp struct {
	Name    string  `json:"name"`
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Status  string  `json:"status"`
	Type    string  `json:"type"`
	Errors  ErrMsgs `json:"errors"`
}

type CustomErrorResp struct {
	Message string `json:"message"`
}

type ErrMsgs map[string][]string
