package validator

import (
	req "github.com/doublehops/dhapi-example/internal/request"
)

const (
	RequiredPropertyError   = "this is a required property"
	ProcessingPropertyError = "unable to process property"
)

type ValidationFuncs func(bool, interface{}) (bool, string)

type Rule struct {
	VariableName string
	Value        interface{}
	Required     bool
	Function     []ValidationFuncs
}

type Error string

func RunValidation(rules []Rule) req.ErrMsgs {
	errorMessages := make(req.ErrMsgs)

	for _, prop := range rules {
		var errors []string
		for _, rule := range prop.Function {
			valid, errMsg := rule(prop.Required, prop.Value)
			if !valid {
				errors = append(errors, errMsg)
			}
		}

		if errors != nil {
			errorMessages[prop.VariableName] = errors
		}
	}

	return errorMessages
}
