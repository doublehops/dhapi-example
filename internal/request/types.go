package resp

type SingleItemResp struct {
	Data interface{} `json:"data"`
}

type CollResp struct {
	Data interface{} `json:"data"`
	Request
}

type Request struct {
	Page       int    `json:"Page"`
	PerPage    int    `json:"perPage"`
	PageCount  int    `json:"pageCount"`
	TotalCount int32  `json:"totalCount"`
	Offset     int    `json:"-"`
	Sort       string `json:"-"`
	Order      string `json:"-"`
}

type GeneralErrorResp struct {
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
