package filter

// AllowedKeys receives a map (from query params) and a list of allowed keys.
// It returns a new map containing only the keys that are present in the allowed list.
// This acts as a HashSet to prevent SQL Column Injection errors and
// ensures safety when passing the map directly into the database's Where condition.
func AllowedKeys(input map[string]interface{}, allowedKeys []string) map[string]interface{} {
	validMap := make(map[string]interface{})

	// Use HashSet pattern (map[string]struct{}) for O(1) lookup.
	// struct{} takes up no additional memory.
	allowedSet := make(map[string]struct{})
	for _, k := range allowedKeys {
		allowedSet[k] = struct{}{}
	}

	// Iterate through the input data, add to validMap only if the key is valid.
	for k, v := range input {
		if _, exists := allowedSet[k]; exists {
			validMap[k] = v
		}
	}

	return validMap
}
