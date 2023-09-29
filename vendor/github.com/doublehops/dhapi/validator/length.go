package validator

const (
	MinLengthDefaultMessage     = "does not meet minimum length"
	MaxLengthDefaultMessage     = "exceeds maximum length"
	BetweenLengthDefaultMessage = "does not conform to min and max lengths"
)

func MinLength(minLength int, errorMessage string) ValidationFunctions {
	return func(required bool, value interface{}) (bool, string) {
		if errorMessage == "" {
			errorMessage = MinLengthDefaultMessage
		}

		var v string
		var ok bool

		if v, ok = value.(string); !ok {
			return false, ProcessingPropertyError
		}

		if v == "" && required {
			return false, RequiredPropertyError
		}

		if v == "" && !required {
			return true, ""
		}

		if len(v) < minLength {
			return false, errorMessage
		}

		return true, ""
	}
}

func MaxLength(maxLength int, errorMessage string) ValidationFunctions {
	return func(required bool, value interface{}) (bool, string) {
		if errorMessage == "" {
			errorMessage = BetweenLengthDefaultMessage
		}

		var v string
		var ok bool

		if v, ok = value.(string); !ok {
			return false, ProcessingPropertyError
		}

		if v == "" && required {
			return false, RequiredPropertyError
		}

		if v == "" && !required {
			return true, ""
		}

		if len(v) > maxLength {
			return false, errorMessage
		}

		return true, ""
	}
}

func Between(minLength, maxLength int, errorMessage string) ValidationFunctions {
	return func(required bool, value interface{}) (bool, string) {
		if errorMessage == "" {
			errorMessage = BetweenLengthDefaultMessage
		}

		var v string
		var ok bool

		if v, ok = value.(string); !ok {
			return false, ProcessingPropertyError
		}

		if v == "" && required {
			return false, RequiredPropertyError
		}

		if v == "" && !required {
			return true, ""
		}

		if len(v) >= minLength && len(v) <= maxLength {
			return true, ""
		}

		return false, errorMessage
	}
}
