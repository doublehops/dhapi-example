package resp

import (
	"errors"
	"net/http"
)

var (
	ValidationError = errors.New("one or more validation errors occurred")
)

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

// GetValidateErrResp will prepare the error response. It will default to a predefined error for Message but
// will override it if one is supplied.
func GetValidateErrResp(errors ErrMsgs, errs ...string) ValidateErrResp {
	err := ValidationError.Error()
	if len(errs) > 0 {
		err = errs[0]
	}

	return ValidateErrResp{
		Name:    "Validation failed",
		Message: err,
		Code:    http.StatusBadRequest,
		Status:  "error",
		Errors:  errors,
	}
}
