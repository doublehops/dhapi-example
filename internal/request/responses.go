package request

import (
	"errors"
	"net/http"
)

var (
	ValidationError        = errors.New("one or more validation errors occurred")
	RecordNotFound         = errors.New("record not found")
	CouldNotSaveRecord     = errors.New("could not save record")
	CouldNotParseRequest   = errors.New("could not parse request")
	ErrorProcessingRequest = errors.New("there was an error processing the request")
	ErrorNotAuthorised     = errors.New("not authorised")

	NotFoundMsg = ErrorMessage{"message": "not found"}
)

type ErrorMessage map[string]string

func GetSingleItemResp(data interface{}) SingleItemResp {
	return SingleItemResp{
		Data: data,
	}
}

func GeneralErrResp(msg string) GeneralErrorResp {
	return GeneralErrorResp{
		Name:    "there was an error processing request",
		Message: msg,
		Code:    http.StatusInternalServerError,
		Status:  "error",
		Errors:  nil,
	}
}

func GetNotAuthorisedResp() GeneralErrorResp {
	return GeneralErrorResp{
		Name:    "there was an error processing request",
		Message: ErrorNotAuthorised.Error(),
		Code:    http.StatusForbidden,
		Status:  "error",
		Errors:  nil,
	}
}

// GetNotFoundResp for returning 404 messages.
func GetNotFoundResp() GeneralErrorResp {
	return GeneralErrorResp{
		Name:    RecordNotFound.Error(),
		Message: RecordNotFound.Error(),
		Code:    http.StatusNotFound,
		Status:  "error",
		Errors:  nil,
	}
}

func GetListResp(data interface{}, pagination *Request) CollResp {
	return CollResp{
		Data:    data,
		Request: *pagination,
	}
}

// GetValidateErrResp will prepare the error response. It will default to a predefined error for Message but
// will override it if one is supplied.
func GetValidateErrResp(errors ErrMsgs, errs ...string) GeneralErrorResp {
	err := ValidationError.Error()
	if len(errs) > 0 {
		err = errs[0]
	}

	return GeneralErrorResp{
		Name:    "Validation failed",
		Message: err,
		Code:    http.StatusBadRequest,
		Status:  "error",
		Errors:  errors,
	}
}

// UnableToParseResp will return a message indicating that the JSON request could not be parsed.
func UnableToParseResp() GeneralErrorResp {
	return GeneralErrorResp{
		Name:    "Parsing error",
		Message: CouldNotParseRequest.Error(),
		Code:    http.StatusBadRequest,
		Status:  "error",
		Errors:  nil,
	}
}

// ErrorProcessingRequestResp will return a message indicating that there was an error processing request.
func ErrorProcessingRequestResp() GeneralErrorResp {
	return GeneralErrorResp{
		Name:    "Parsing error",
		Message: ErrorProcessingRequest.Error(),
		Code:    http.StatusInternalServerError,
		Status:  "error",
		Errors:  nil,
	}
}
