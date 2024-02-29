package validator

const (
	keyNotInSlice = "value was not found in available items"
)

func In(slice []interface{}, errorMessage string) ValidationFuncs {
	return func(required bool, value interface{}) (bool, string) {
		if errorMessage == "" {
			errorMessage = keyNotInSlice
		}

		// We want to convert the slice values into strings.
		var strSlice []string
		for _, item := range slice {
			i, ok := item.(string)
			if !ok {
				return false, ProcessingPropertyError
			}
			strSlice = append(strSlice, i)
		}

		v, ok := value.(string)
		if !ok {
			return false, ProcessingPropertyError
		}

		if v == "" && required {
			return false, RequiredPropertyError
		}

		if v == "" && !required {
			return true, ""
		}

		for _, item := range strSlice {
			if item == v {
				return true, ""
			}
		}

		return false, errorMessage
	}
}
