package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SingleItemResponse(c *gin.Context, data interface{}) {
	resp := singleItemResponseType{
		Data: data,
	}

	c.JSON(http.StatusOK, resp)
}

func MultiItemResponse(c *gin.Context, data interface{}, pagination PaginationType) {
	resp := multiItemResponseType{
		Data:           data,
		PaginationType: pagination,
	}

	c.JSON(http.StatusOK, resp)
}

func ValidationErrorResponse(c *gin.Context, code int, errors []ValidationField) {
	resp := ValidationErrorResponseType{
		Name:    "Validation failed",
		Message: "One or more validation errors occurred",
		Code:    code,
		Status:  "error",
		Errors:  errors,
	}

	c.JSON(code, resp)
}
