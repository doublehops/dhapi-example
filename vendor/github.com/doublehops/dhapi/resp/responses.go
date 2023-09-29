package resp

import "net/http"

func GetSingleItemResp(data interface{}) SingleItemResp {
	return SingleItemResp{
		Data: data,
	}
}

func GetListResp(data interface{}, pagination Pagination) ListResp {
	return ListResp{
		Data:       data,
		Pagination: pagination,
	}
}

func GetValidateErrResp(errors ErrMsgs) ValidateErrResp {
	return ValidateErrResp{
		Name:    "Validation failed",
		Message: "One or more validation errors occurred",
		Code:    http.StatusBadRequest,
		Status:  "error",
		Errors:  errors,
	}
}
