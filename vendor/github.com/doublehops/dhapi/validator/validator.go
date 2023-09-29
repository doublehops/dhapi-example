package validator

import "github.com/doublehops/dhapi/responses"

const (
	RequiredPropertyError   = "this is a required property"
	ProcessingPropertyError = "unable to process property"
)

type ValidationFunctions func(bool, interface{}) (bool, string)

type Rule struct {
	VariableName string
	Value        interface{} // WHAT WAS THIS?
	Required     bool
	Function     []ValidationFunctions // AND WHAT WAS THIS?
}

type Error string

//type ErrorMessages map[string][]Error

func RunValidation(rules []Rule) responses.ErrorMessages {
	errorMessages := make(responses.ErrorMessages)

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
