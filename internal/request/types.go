package resp

// CustomErrorResp for custom errors
type CustomErrorResp struct {
	Message string `json:"message"`
}

// ErrMsgs for mapping the error responses for validation issues.
type ErrMsgs map[string][]string

// SingleItemResp for single record responses.
type SingleItemResp struct {
	Data any `json:"data"`
}

// CollResp - Collection response.
type CollResp struct {
	Data any `json:"data"`
	Request
}

// Request will be populated with request data and used as paginated data in the response.
type Request struct {
	Page       int    `json:"Page"`
	PerPage    int    `json:"perPage"`
	PageCount  int    `json:"pageCount"`
	TotalCount int32  `json:"totalCount"`
	Offset     int    `json:"-"`
	Sort       string `json:"-"`
	Order      string `json:"-"`
}

// GeneralErrorResp is a function to return general errors including validation.
type GeneralErrorResp struct {
	Name    string  `json:"name"`
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Status  string  `json:"status"`
	Type    string  `json:"type"`
	Errors  ErrMsgs `json:"errors"`
}
