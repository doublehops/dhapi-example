package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxValue(t *testing.T) {
	tests := []struct {
		name           string
		maxValue       int
		errorMessage   string
		required       bool
		value          int
		expectedErrMsg string
		expectedResult bool
	}{
		{
			name:           "goodValue",
			maxValue:       20,
			errorMessage:   "",
			required:       true,
			value:          10,
			expectedErrMsg: "",
			expectedResult: true,
		},
		{
			name:           "valueTooHigh",
			maxValue:       20,
			errorMessage:   "",
			required:       true,
			value:          25,
			expectedErrMsg: MaxValueDefaultMessage,
			expectedResult: false,
		},
		{
			name:           "valueTooHighCustomErrorMessage",
			maxValue:       20,
			errorMessage:   "Value is too low -- test",
			required:       true,
			value:          25,
			expectedErrMsg: "Value is too low -- test",
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := MaxValue(tt.maxValue, tt.errorMessage)
			res, validationErrMsg := f(tt.required, tt.value)
			assert.Equal(t, tt.expectedErrMsg, validationErrMsg, "error not as expected")
			assert.Equal(t, tt.expectedResult, res, "result not as expected")
		})
	}
}

func TestMinValue(t *testing.T) {
	tests := []struct {
		name           string
		minValue       int
		errorMessage   string
		required       bool
		value          int
		expectedErrMsg string
		expectedResult bool
	}{
		{
			name:           "goodValue",
			minValue:       20,
			errorMessage:   "",
			required:       true,
			value:          21,
			expectedErrMsg: "",
			expectedResult: true,
		},
		{
			name:           "valueTooLow",
			minValue:       20,
			errorMessage:   "",
			required:       true,
			value:          10,
			expectedErrMsg: MinValueDefaultMessage,
			expectedResult: false,
		},
		{
			name:           "valueTooLowCustomErrorMessage",
			minValue:       20,
			errorMessage:   "Value is too low -- test",
			required:       true,
			value:          10,
			expectedErrMsg: "Value is too low -- test",
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := MinValue(tt.minValue, tt.errorMessage)
			res, validationErrMsg := f(tt.required, tt.value)
			assert.Equal(t, tt.expectedErrMsg, validationErrMsg, "error not as expected")
			assert.Equal(t, tt.expectedResult, res, "result not as expected")
		})
	}
}
