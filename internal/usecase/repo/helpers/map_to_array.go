package helpers

func MapToArray[T any, V comparable](keys []V, input map[V]T) []T {
	result := make([]T, 0)
	for _, key := range keys {
		result = append(result, input[key])
	}

	return result
}
