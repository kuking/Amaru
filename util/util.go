package util

// SafeSlice returns a slice of the input slice from offset to offset+limit,
// ensuring that it does not exceed the slice's bounds.
func SafeSlice[T any](slice []T, offset, limit int) []T {
	// If offset is beyond the end of the slice, return an empty slice.
	if offset >= len(slice) {
		return []T{}
	}

	// Calculate the end index, adjusting it to be within bounds.
	end := offset + limit
	if end > len(slice) {
		end = len(slice)
	}

	return slice[offset:end]
}

// RemoveDuplicates removes duplicates from a slice.
func RemoveDuplicates[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := []T{}
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}
