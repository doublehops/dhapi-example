package responses

import "net/http"

func GetSingleItemResponse(data interface{}) SingleItemResponse {
	return SingleItemResponse{
		Data: data,
	}
}

func GetMultiItemResponse(data interface{}, pagination PaginationType) MultiItemResponse {
	return MultiItemResponse{
		Data:           data,
		PaginationType: pagination,
	}
}

func GetValidationErrorResponse(errors ErrorMessages) ValidationErrorResponse {
	return ValidationErrorResponse{
		Name:    "Validation failed",
		Message: "One or more validation errors occurred",
		Code:    http.StatusBadRequest,
		Status:  "error",
		Errors:  errors,
	}
}
