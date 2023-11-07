package str

// SliceContains will take the string and check that it exists
// within the slice given and return boolean
func SliceContains[T comparable](key T, slice []T) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}

	return false
}

// KeyStringExists will check that a key exists in the given map
func KeyStringExists(key string, m map[string]interface{}) bool {

	_, ok := m[key]
	return ok
}
