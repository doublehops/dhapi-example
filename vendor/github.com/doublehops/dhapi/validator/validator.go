package validator

import "github.com/doublehops/dhapi/resp"

const (
	RequiredPropertyError   = "this is a required property"
	ProcessingPropertyError = "unable to process property"
)

type ValidateFuncs func(bool, interface{}) (bool, string)

type Rule struct {
	VariableName string
	Value        interface{}
	Required     bool
	Function     []ValidateFuncs
}

type Error string

func RunValidation(rules []Rule) resp.ErrMsgs {
	errorMessages := make(resp.ErrMsgs)

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
