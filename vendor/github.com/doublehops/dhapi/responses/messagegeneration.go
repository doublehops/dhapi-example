package responses

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// getStandardErrors contains a list of standard errors for the given validation type. Eg. `required`.
func getStandardErrors(tag string) string {
	errorMap := map[string]string{
		"required": "this is a required field",
	}

	if _, ok := errorMap[tag]; ok {
		return errorMap[tag]
	}

	return ""
}

// getErrorMessage will return a custom response if the field/tag match is found. It will otherwise look for a
// standard defined one before resorting to the Gin related one.
func getErrorMessage(field, tag, ginError string, messages CustomerErrorMessages) string {
	if _, ok := messages[field]; ok {
		if _, ok = messages[field][tag]; ok {
			return messages[field][tag]
		}
	}

	stdError := getStandardErrors(tag)
	if stdError != "" {
		return stdError
	}

	return parseGinError(ginError)
}

// parseGinError will strip the superfluous text from the Gin generated error and return it as the error.
func parseGinError(str string) string {
	pos := strings.Index(str, "Error:")
	substr := str[pos+6:]

	return substr
}

// Validate will use Gin's v10 validator to determine whether the requested payload meets the specified requirements.
func Validate[T any](c *gin.Context, obj *T, messages CustomerErrorMessages) []ValidationField {
	var errs []ValidationField

	err := c.BindQuery(&obj)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				thisError := ValidationField{}
				field, _ := reflect.TypeOf(obj).Elem().FieldByName(fe.Field())
				fieldName, _ := field.Tag.Lookup("json")
				thisError[fieldName] = []string{getErrorMessage(fieldName, fe.Tag(), fe.Error(), messages)}
				errs = append(errs, thisError)
			}
		}
	}

	return errs
}
